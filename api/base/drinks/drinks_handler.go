package drinks

import (
    "fmt"
    "net/http"
    "encoding/json"
    "slijterij/db"
    "slijterij/api/generic"
    "slijterij/api/base/drinks/drinksmodel"
    "github.com/go-mysql/errors"
)

type DrinksHandler struct {
    store *db.DataStore
}

func (h *DrinksHandler) GetAllDrinks(w http.ResponseWriter, r *http.Request) {
    query := r.URL.Query()
    fmt.Println(query)
    w.Write(nil)
}

func (h *DrinksHandler) CreateDrink(w http.ResponseWriter, r *http.Request) {
    drink := &drinksmodel.DrinkEntity{}
    reqJsonErr := json.NewDecoder(r.Body).Decode(drink)
    if reqJsonErr != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write(generic.JSONError("Bad Request: Invalid JSON"))
    }

    _, sqlErr := h.store.CreateDrink(drink)
    if sqlErr != nil {
        code := errors.MySQLErrorCode(sqlErr)
        if code == 1062 {
            w.WriteHeader(http.StatusConflict)
            w.Write(generic.JSONError("Item already exists"))
            return
        }
        w.WriteHeader(http.StatusInternalServerError)
        w.Write(generic.JSONError("Internal server error"))
        fmt.Printf("ERROR:\t%s\n", sqlErr)
        return
    }

    jsonResponse, resJsonErr := json.Marshal(drink)
    if resJsonErr != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write(generic.JSONError("Could not create JSON response but item is created"))
    }

    w.WriteHeader(http.StatusCreated)
    w.Write(jsonResponse)
}

