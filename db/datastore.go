package db

import (
    "database/sql"
    "fmt"
    "os"
    "log"
    "github.com/go-sql-driver/mysql"
    "slijterij/api/base/bar/barmodel"
    "slijterij/api/base/drinks/drinksmodel"
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

func (s *DataStore) RetrieveBar(id string) (*barmodel.BarEntity, error) {
    bar := &barmodel.BarEntity{}

    row := db.QueryRow("SELECT * FROM bar WHERE id = ?", id)
    if err := row.Scan(&bar.Id, &bar.Password, &bar.Token); err != nil {
        if err == sql.ErrNoRows {
            return bar, fmt.Errorf("RetrieveBar %s: no such bar", id)
        }
        return bar, fmt.Errorf("RetrieveBar %s: %v", id, err)
    }
    return bar, nil
}

func (s *DataStore) CreateBar(entity *barmodel.BarEntity) (int64, error) {
    result, queryErr := db.Exec("INSERT INTO bar VALUES (?, ?, ?)", entity.Id, entity.Password, entity.Token)
    if queryErr != nil {
        return -1, fmt.Errorf("CreateBar: %v", queryErr)
    }

    id, idErr := result.LastInsertId()
    if idErr != nil {
        return -1, fmt.Errorf("CreateBar: %v", idErr)
    }

    return id, nil
}

func (s *DataStore) GetAllDrinks(barId string) ([]drinksmodel.DrinkEntity, error) {
    var drinks []drinksmodel.DrinkEntity

    rows, queryErr := db.Query("SELECT * FROM product WHERE bar_id = ?", barId)
    if queryErr != nil {
        return nil, queryErr
    }
    defer rows.Close()

    for rows.Next() {
        var drink drinksmodel.DrinkEntity
        if convertErr := rows.Scan(&drink.Id, &drink.Name, &drink.BarId, &drink.StartPrice, &drink.CurrentPrice, &drink.Multiplier); convertErr != nil {
            continue
        }
        drinks = append(drinks, drink)
    }

    if rowsErr := rows.Err(); rowsErr != nil {
        return nil, rowsErr
    }
    return drinks, nil
}

func (s *DataStore) CreateDrink(entity *drinksmodel.DrinkEntity) (int64, error) {
    result, queryErr := db.Exec("INSERT INTO product VALUES (?, ?, ?, ?, ?, ?)", entity.Id, entity.Name, entity.BarId, entity.StartPrice, entity.CurrentPrice, entity.Multiplier)
    if queryErr != nil {
        return -1, queryErr
    }

    id, idErr := result.LastInsertId()
    if idErr != nil {
        return -1, fmt.Errorf("CreateDrink: %v", idErr)
    }

    return id, nil
}


