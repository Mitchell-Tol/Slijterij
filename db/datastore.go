package db

import (
    "database/sql"
    "fmt"
    "os"
    "log"
    "github.com/go-sql-driver/mysql"
    "slijterij/api/base/bar/barmodel"
    "slijterij/api/base/drinks/drinksmodel"
    "slijterij/api/base/device/devicemodel"
	"github.com/google/uuid"
)

var db *sql.DB

type DataStore struct {}

func NewStore() *DataStore {
    cfg := mysql.Config{
        User:                   os.Getenv("DBUSER"),
        Passwd:                 os.Getenv("DBPASS"),
        Net:                    "tcp",
        Addr:                   "127.0.0.1:3306",
        DBName:                 "drankbeurs",
        AllowNativePasswords:   true,
    }

    var err error
    db, err = sql.Open("mysql", cfg.FormatDSN())
    if err != nil {
        log.Fatal(err)
    }

    pingErr := db.Ping()
    if pingErr != nil {
        log.Fatal(pingErr)
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

func (s *DataStore) CreateBar(entity *barmodel.BarEntity) (int64, error) {
    result, queryErr := db.Exec("INSERT INTO bar VALUES (?, ?, ?, ?)", entity.Id, entity.Name, entity.Password, entity.Token)
    if queryErr != nil {
        return -1, fmt.Errorf("CreateBar: %v", queryErr)
    }

    id, idErr := result.LastInsertId()
    if idErr != nil {
        return -1, fmt.Errorf("CreateBar: %v", idErr)
    }

    return id, nil
}

func (s *DataStore) UpdateBar(entity *barmodel.BarEntity) (*barmodel.BarEntity, error) {
    _, queryErr := db.Exec("UPDATE bar SET name = ?, password = ?, token = ? WHERE id = ?", entity.Name, entity.Password, entity.Token, entity.Id)
    if queryErr != nil {
        return nil, queryErr
    }

    return entity, nil
}

func (s *DataStore) DeleteBar(id string) (error) {
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
        Id: newId,
        BarId: model.BarId,
        Name: model.Name,
    }
    return result, nil
}

func (s *DataStore) UpdateDevice(updated *devicemodel.UpdatedDevice) (*devicemodel.UpdatedDevice, error) {
	_, queryErr := db.Exec("UPDATE device SET name = ?", updated.Name)
	if queryErr != nil {
		fmt.Println("UpdateDevice: %v", queryErr)
		return nil, queryErr
	}

	return updated, nil
}

func (s *DataStore) DeleteDevice(id string) (error) {
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
        if convertErr := rows.Scan(&drink.Id, &drink.Name, &drink.BarId, &drink.StartPrice, &drink.CurrentPrice, &drink.Multiplier, &drink.Tag); convertErr != nil {
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

func (s *DataStore) CreateDrink(entity *drinksmodel.DrinkEntity) (string, error) {
	generatedId := uuid.New().String()
    result, queryErr := db.Exec("INSERT INTO product VALUES (?, ?, ?, ?, ?, ?, ?)", generatedId, entity.Name, entity.BarId, entity.StartPrice, entity.CurrentPrice, entity.Multiplier, entity.Tag)
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
    _, queryErr := db.Exec("UPDATE product SET name = ?, start_price = ?, current_price = ?, multiplier = ?, tag = ? WHERE id = ?", entity.Name, entity.StartPrice, entity.CurrentPrice, entity.Multiplier, entity.Tag, entity.Id)
    if queryErr != nil {
        return nil, queryErr
    }

    return entity, nil
}

func (s *DataStore) DeleteDrink(id string) (error) {
    _, sqlErr := db.Exec("DELETE FROM product WHERE id = ?", id)
    return sqlErr
}
