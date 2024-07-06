package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"
	"slijterij/api/base/bar/barmodel"
	"slijterij/api/base/category/categorymodel"
	"slijterij/api/base/device/devicemodel"
	"slijterij/api/base/drinks/drinksmodel"
	"slijterij/api/base/order/ordermodel"
	"github.com/google/uuid"
)

var db *sql.DB

type DataStore struct{}

func NewStore() *DataStore {
	user := os.Getenv("DBUSER")
	password := os.Getenv("DBPASS")
	database := os.Getenv("DBNAME")

	connectSuccess := false
	for !connectSuccess {
		log.Println("Connecting to MYSQL")
		var err error
		db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(db:3306)/%s", user, password, database))
		if err != nil {
			log.Println(err)
		}
	
		pingErr := db.Ping()
		if pingErr != nil {
			log.Println(pingErr)
			log.Println("Trying again in 5s")
			time.Sleep(5000000000)
		} else {
			connectSuccess = true
		}
	}

	return &DataStore{}
}

// BARS
func (s *DataStore) GetAllBars() ([]barmodel.BarEntity, error) {
	bars := []barmodel.BarEntity{}

	rows, queryErr := db.Query("SELECT * FROM bar")
	if queryErr != nil {
		return nil, queryErr
	}

	defer rows.Close()
	for rows.Next() {
		var bar barmodel.BarEntity
		convertErr := rows.Scan(&bar.Id, &bar.Name, &bar.Password, &bar.Token)
		if convertErr != nil {
			fmt.Errorf("GetAllBars %v", convertErr)
			continue
		}
		bars = append(bars, bar)
	}

	if rowsErr := rows.Err(); rowsErr != nil {
		return nil, rowsErr
	}
	return bars, nil
}

func (s *DataStore) RetrieveBar(name string) (*barmodel.BarEntity, error) {
	bar := &barmodel.BarEntity{}

	row := db.QueryRow("SELECT * FROM bar WHERE name = ?", name)
	if err := row.Scan(&bar.Id, &bar.Name, &bar.Password, &bar.Token); err != nil {
		if err == sql.ErrNoRows {
			return bar, fmt.Errorf("RetrieveBar %s: no such bar", name)
		}
		return bar, fmt.Errorf("RetrieveBar %s: %v", name, err)
	}
	return bar, nil
}

func (s *DataStore) CreateBar(entity *barmodel.BarEntity) (*barmodel.BarEntity, error) {
	_, queryErr := db.Exec("INSERT INTO bar VALUES (?, ?, ?, ?)", entity.Id, entity.Name, entity.Password, entity.Token)
	if queryErr != nil {
		return nil, fmt.Errorf("CreateBar: %v", queryErr)
	}

	return entity, nil
}

func (s *DataStore) UpdateBar(entity *barmodel.BarEntity) (*barmodel.BarEntity, error) {
	_, queryErr := db.Exec("UPDATE bar SET name = ?, password = ?, token = ? WHERE id = ?", entity.Name, entity.Password, entity.Token, entity.Id)
	if queryErr != nil {
		return nil, queryErr
	}

	return entity, nil
}

func (s *DataStore) DeleteBar(id string) error {
	_, queryErr := db.Exec("DELETE FROM bar WHERE id = ?", id)
	return queryErr
}

// DEVICES
func (s *DataStore) GetDevices(barId string) ([]devicemodel.DeviceEntity, error) {
	devices := []devicemodel.DeviceEntity{}
	rows, queryErr := db.Query("SELECT * FROM device WHERE bar_id = ?", barId)
	if queryErr != nil {
		fmt.Println("GetAllDevices: %v", queryErr)
		return nil, queryErr
	}
	defer rows.Close()

	for rows.Next() {
		var device devicemodel.DeviceEntity
		convertErr := rows.Scan(&device.Id, &device.BarId, &device.Name)
		if convertErr != nil {
			fmt.Println("GetAllDevices: %v", convertErr)
			continue
		}
		devices = append(devices, device)
	}

	if rowsErr := rows.Err(); rowsErr != nil {
		return nil, rowsErr
	}
	return devices, nil
}

func (s *DataStore) CreateDevice(model *devicemodel.Device) (*devicemodel.DeviceEntity, error) {
	newId := uuid.New().String()
	_, queryErr := db.Exec("INSERT INTO device VALUES (?, ?, ?)", newId, model.BarId, model.Name)
	if queryErr != nil {
		return nil, queryErr
	}

	result := &devicemodel.DeviceEntity{
		Id:    newId,
		BarId: model.BarId,
		Name:  model.Name,
	}
	return result, nil
}

func (s *DataStore) UpdateDevice(updated *devicemodel.UpdatedDevice) (*devicemodel.UpdatedDevice, error) {
	_, queryErr := db.Exec("UPDATE device SET name = ? WHERE id = ?", updated.Name, updated.Id)
	if queryErr != nil {
		fmt.Println("UpdateDevice: %v", queryErr)
		return nil, queryErr
	}

	return updated, nil
}

func (s *DataStore) DeleteDevice(id string) error {
	_, queryError := db.Exec("DELETE FROM device WHERE id = ?", id)
	if queryError != nil {
		fmt.Println("DeleteDevice: %v", queryError)
	}
	return queryError
}

// DRINKS
func (s *DataStore) GetAllDrinks(barId string) ([]drinksmodel.DrinkEntity, error) {
	drinks := []drinksmodel.DrinkEntity{}

	rows, queryErr := db.Query("SELECT * FROM product WHERE bar_id = ?", barId)
	if queryErr != nil {
		return nil, queryErr
	}
	defer rows.Close()

	for rows.Next() {
		var drink drinksmodel.DrinkEntity
		if convertErr := rows.Scan(&drink.Id, &drink.Name, &drink.BarId, &drink.StartPrice, &drink.CurrentPrice, &drink.RiseMultiplier, &drink.Tag, &drink.CategoryId, &drink.DropMultiplier); convertErr != nil {
			fmt.Errorf("GetAllDrinks %s: %v", barId, convertErr)
			continue
		}
		drinks = append(drinks, drink)
	}

	if rowsErr := rows.Err(); rowsErr != nil {
		return nil, rowsErr
	}
	return drinks, nil
}

func (s *DataStore) GetDrinksByCategory(categoryId string) ([]drinksmodel.DrinkEntity, error) {
	drinks := []drinksmodel.DrinkEntity{}

	rows, queryErr := db.Query("SELECT * FROM product WHERE category_id = ?", categoryId)
	if queryErr != nil {
		return nil, queryErr
	}
	defer rows.Close()

	for rows.Next() {
		var drink drinksmodel.DrinkEntity
		if convertErr := rows.Scan(&drink.Id, &drink.Name, &drink.BarId, &drink.StartPrice, &drink.CurrentPrice, &drink.RiseMultiplier, &drink.Tag, &drink.CategoryId, &drink.DropMultiplier); convertErr != nil {
			fmt.Errorf("GetAllDrinks %s: %v", categoryId, convertErr)
			continue
		}
		drinks = append(drinks, drink)
	}

	if rowsErr := rows.Err(); rowsErr != nil {
		return nil, rowsErr
	}
	return drinks, nil
}

func (s *DataStore) GetDrink(id string) (*drinksmodel.DrinkEntity, error) {
	drink := &drinksmodel.DrinkEntity{}

	row := db.QueryRow("SELECT * FROM product WHERE id = ?", id)
	queryErr := row.Scan(&drink.Id, &drink.Name, &drink.BarId, &drink.StartPrice, &drink.CurrentPrice, &drink.RiseMultiplier, &drink.Tag, &drink.CategoryId, &drink.DropMultiplier)
	if queryErr != nil {
		return drink, fmt.Errorf("GetDrink %s: %v", id, queryErr)
	}

	return drink, nil
}

func (s *DataStore) CreateDrink(entity *drinksmodel.Drink) (string, error) {
	generatedId := uuid.New().String()
	result, queryErr := db.Exec("INSERT INTO product VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)", generatedId, entity.Name, entity.BarId, entity.StartPrice, entity.CurrentPrice, entity.RiseMultiplier, entity.Tag, entity.CategoryId, entity.DropMultiplier)
	if queryErr != nil {
		return "", queryErr
	}

	_, idErr := result.LastInsertId()
	if idErr != nil {
		return "", fmt.Errorf("CreateDrink: %v", idErr)
	}

	return generatedId, nil
}

func (s *DataStore) UpdateDrink(entity *drinksmodel.DrinkEntity) (*drinksmodel.DrinkEntity, error) {
	_, queryErr := db.Exec("UPDATE product SET name = ?, start_price = ?, current_price = ?, rise_multiplier = ?, tag = ?, category_id = ?, drop_multiplier = ? WHERE id = ?", entity.Name, entity.StartPrice, entity.CurrentPrice, entity.RiseMultiplier, entity.Tag, entity.CategoryId, entity.DropMultiplier, entity.Id)
	if queryErr != nil {
		fmt.Println("UpdateDrink: %v", queryErr)
		return nil, queryErr
	}

	return entity, nil
}

func (s *DataStore) DeleteDrink(id string) error {
	_, sqlErr := db.Exec("DELETE FROM product WHERE id = ?", id)
	return sqlErr
}

// CATEGORIES
func (s *DataStore) GetAllCategories(barId string) ([]categorymodel.CategoryEntity, error) {
	categories := []categorymodel.CategoryEntity{}

	rows, queryErr := db.Query("SELECT * FROM category WHERE bar_id = ?", barId)
	if queryErr != nil {
		return nil, queryErr
	}
	defer rows.Close()

	for rows.Next() {
		var category categorymodel.CategoryEntity
		if convertErr := rows.Scan(&category.Id, &category.Name, &category.BarId); convertErr != nil {
			fmt.Errorf("GetAllCategories %s: %v", barId, convertErr)
			continue
		}
		categories = append(categories, category)
	}

	if rowsErr := rows.Err(); rowsErr != nil {
		return nil, rowsErr
	}
	return categories, nil
}

func (s *DataStore) CreateCategory(model *categorymodel.Category) (*categorymodel.CategoryEntity, error) {
	newId := uuid.New().String()
	_, queryErr := db.Exec("INSERT INTO category VALUES (?, ?, ?)", newId, model.Name, model.BarId)
	if queryErr != nil {
		return nil, queryErr
	}

	result := &categorymodel.CategoryEntity{
		Id:    newId,
		BarId: model.BarId,
		Name:  model.Name,
	}
	return result, nil
}

func (s *DataStore) UpdateCategory(model *categorymodel.UpdatedCategory) (*categorymodel.UpdatedCategory, error) {
	_, queryErr := db.Exec("UPDATE category SET name = ? WHERE id = ?", model.Name, model.Id)
	if queryErr != nil {
		return nil, queryErr
	}

	return model, nil
}

func (s *DataStore) DeleteCategory(id string) error {
	_, sqlErr := db.Exec("DELETE FROM category WHERE id = ?", id)
	return sqlErr
}

// ORDERS
func (s *DataStore) GetAllOrders(deviceId string) ([]ordermodel.OrderEntity, error) {
	orders := []ordermodel.OrderEntity{}

	rows, queryErr := db.Query("SELECT * FROM `order` WHERE device_id = ?", deviceId)
	if queryErr != nil {
		return nil, fmt.Errorf("GetAllOrders %s: %v\n", deviceId, queryErr)
	}
	defer rows.Close()

	for rows.Next() {
		var order ordermodel.OrderEntity
		if convertErr := rows.Scan(&order.Id, &order.DeviceId, &order.ProductId, &order.Timestamp, &order.Amount, &order.PricePerProduct); convertErr != nil {
			fmt.Errorf("GetAllOrders %s: %v\n", deviceId, convertErr)
			continue
		}
		orders = append(orders, order)
	}

	if rowsErr := rows.Err(); rowsErr != nil {
		return nil, rowsErr
	}
	return orders, nil
}

func (s *DataStore) CreateOrder(model *ordermodel.Order) (*ordermodel.OrderEntity, error) {
	newId := uuid.New().String()
	_, queryErr := db.Exec("INSERT INTO `order` VALUES (?, ?, ?, ?, ?, ?)", newId, model.DeviceId, model.ProductId, model.Timestamp, model.Amount, model.PricePerProduct)
	if queryErr != nil {
		return nil, fmt.Errorf("GetAllOrders: %v\n", queryErr)
	}

	result := &ordermodel.OrderEntity{
		Id:              newId,
		DeviceId:        model.DeviceId,
		ProductId:       model.ProductId,
		Timestamp:       model.Timestamp,
		Amount:          model.Amount,
		PricePerProduct: model.PricePerProduct,
	}
	return result, nil
}

func (s *DataStore) UpdateOrder(updated *ordermodel.UpdatedOrder) (*ordermodel.UpdatedOrder, error) {
	_, queryErr := db.Exec("UPDATE `order` SET device_id = ?, product_id = ?, amount = ?, price_per_product = ? WHERE id = ?", updated.DeviceId, updated.ProductId, updated.Amount, updated.PricePerProduct, updated.Id)
	if queryErr != nil {
		fmt.Println("UpdateOrder: %v", queryErr)
		return nil, queryErr
	}

	return updated, nil
}

func (s *DataStore) DeleteOrder(id string) error {
	_, queryError := db.Exec("DELETE FROM `order` WHERE id = ?", id)
	if queryError != nil {
		fmt.Println("DeleteOrder: %v", queryError)
	}
	return queryError
}
