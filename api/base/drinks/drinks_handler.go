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
	body := &drinksmodel.DrinkByBar{}
	jsonErr := json.NewDecoder(r.Body).Decode(body)
	if jsonErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(generic.JSONError("Invalid JSON"))
		return
	}

    drinks, sqlErr := h.store.GetAllDrinks(body.BarId)
    if sqlErr != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write(generic.JSONError("Internal server error"))
		return
    }

    jsonRes, jsonErr := json.Marshal(drinks)
    if jsonErr != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write(generic.JSONError("Something went wrong when mapping the response to valid JSON"))
		return
    }

    w.WriteHeader(http.StatusOK)
    w.Write(jsonRes)
}

func (h *DrinksHandler) CreateDrink(w http.ResponseWriter, r *http.Request) {
    drink := &drinksmodel.Drink{}
    reqJsonErr := json.NewDecoder(r.Body).Decode(drink)
    if reqJsonErr != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write(generic.JSONError("Bad Request: Invalid JSON"))
		return
    }

    id, sqlErr := h.store.CreateDrink(drink)
	entity := &drinksmodel.DrinkEntity{
		Id: id,
		Name: drink.Name,
		BarId: drink.BarId,
		StartPrice: drink.StartPrice,
		CurrentPrice: drink.CurrentPrice,
		RiseMultiplier: drink.RiseMultiplier,
		Tag: drink.Tag,
		CategoryId: drink.CategoryId,
		DropMultiplier: drink.DropMultiplier,
	}
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

    jsonResponse, resJsonErr := json.Marshal(entity)
    if resJsonErr != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write(generic.JSONError("Could not create JSON response but item is created"))
		return
    }

    w.WriteHeader(http.StatusCreated)
    w.Write(jsonResponse)
}

func (h *DrinksHandler) UpdateDrink(w http.ResponseWriter, r *http.Request) {
	drink := &drinksmodel.DrinkEntity{}
    reqJsonErr := json.NewDecoder(r.Body).Decode(drink)
    if reqJsonErr != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write(generic.JSONError("Bad Request: Invalid JSON"))
		return
    }

	result, sqlErr := h.store.UpdateDrink(drink)
	if sqlErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(generic.JSONError("Internal server error"))
		return
	}

	resJson, jsonErr := json.Marshal(result)
	if jsonErr!= nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(generic.JSONError("Error while converting drink to JSON"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(resJson)
}

func (h *DrinksHandler) DeleteDrink(w http.ResponseWriter, r *http.Request) {
	idBody := &drinksmodel.DrinkId{}
	reqJsonErr := json.NewDecoder(r.Body).Decode(idBody)
	if reqJsonErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(generic.JSONError("Invalid Body"))
		return
	}

	dbErr := h.store.DeleteDrink(idBody.Id)
	if dbErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(generic.JSONError("Internal server error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(nil)
}
