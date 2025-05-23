package order

import (
	"encoding/json"
	"fmt"
	"math"
	"net/http"
	"slijterij/api/base/drinks/drinksmodel"
	"slijterij/api/base/order/ordermodel"
	"slijterij/api/generic"
	"slijterij/db"
)

type OrderHandler struct {
	store *db.DataStore
}

func (h *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()["deviceId"]
	if len(params) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(generic.JSONError("No deviceId parameter provided"))
		return
	}

	result, queryErr := h.store.GetAllOrders(params[0])
	if queryErr != nil {
		fmt.Printf("%v", queryErr)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(generic.JSONError("An error occurred while retrieving orders"))
		return
	}

	jsonResponse, jsonErr := json.Marshal(result)
	if jsonErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(generic.JSONError("An error occurred while mapping response to JSON"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (h *OrderHandler) PostOrder(w http.ResponseWriter, r *http.Request) {
    body := &ordermodel.Order{}
    jsonParseErr := json.NewDecoder(r.Body).Decode(body)
    if jsonParseErr != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write(generic.JSONError("Invalid JSON"))
        return
    }

    entity, queryErr := h.store.CreateOrder(body)
    if queryErr != nil {
		fmt.Println("Error in CreateOrder: %v", queryErr)
        w.WriteHeader(http.StatusInternalServerError)
        w.Write(generic.JSONError("Internal server error"))
        return
    }

	increaseErr := h.IncreaseDrink(entity.ProductId, body.Amount)
	if increaseErr != nil {
		fmt.Println("%v", increaseErr)
	}

    jsonResp, jsonErr := json.Marshal(entity)
    if jsonErr != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write(generic.JSONError("Error while parsing created item to JSON"))
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write(jsonResp)
}

func (h *OrderHandler) UpdateOrder(w http.ResponseWriter, r *http.Request) {
	body := &ordermodel.UpdatedOrder{}
	reqErr := json.NewDecoder(r.Body).Decode(body)
	if reqErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(generic.JSONError("Invalid JSON"))
		return
	}

	result, sqlErr := h.store.UpdateOrder(body)
	if sqlErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(generic.JSONError("An error occurred while updating the order"))
		return
	}

	jsonResponse, jsonErr := json.Marshal(result)
	if jsonErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(generic.JSONError("An error occurred while mapping the response to JSON"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)
}

func (h *OrderHandler) DeleteOrder(w http.ResponseWriter, r *http.Request) {
	body := &ordermodel.OrderId{}
	reqErr := json.NewDecoder(r.Body).Decode(body)
	if reqErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(generic.JSONError("Invalid JSON"))
		return
	}

	queryErr := h.store.DeleteOrder(body.Id)
	if queryErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(generic.JSONError("An error occurred while deleting the item:"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(nil)
}

func (h *OrderHandler) IncreaseDrink(id string, amount int) (error) {
	drink, drinkQueryErr := h.store.GetDrink(id)
	if drinkQueryErr != nil {
		return fmt.Errorf("IncreaseDrink %s: %v", id, drinkQueryErr)
	}

	allDrinks, allDrinksErr := h.store.GetDrinksByCategory(drink.CategoryId)
	if allDrinksErr != nil {
		return fmt.Errorf("IncreaseDrink %s: %v", id, allDrinksErr)
	}

	for i := 0; i < len(allDrinks); i++ {
		var newPrice float32
		if allDrinks[i].Id == drink.Id {
			modifier := drink.StartPrice * drink.RiseMultiplier + 1
			byAmount := math.Pow(float64(modifier), float64(amount))
			newPrice = drink.CurrentPrice * float32(byAmount)
		} else {
			modifier := 1 - allDrinks[i].StartPrice * allDrinks[i].RiseMultiplier
			byAmount := math.Pow(float64(modifier), float64(amount))
			newPrice = allDrinks[i].CurrentPrice * float32(byAmount)
		}
		h.store.UpdateDrink(
			&drinksmodel.DrinkEntity{
				Id: allDrinks[i].Id,
				Name: allDrinks[i].Name,
				BarId: allDrinks[i].BarId,
				StartPrice: allDrinks[i].StartPrice,
				CurrentPrice: newPrice,
				RiseMultiplier: allDrinks[i].RiseMultiplier,
				Tag: allDrinks[i].Tag,
				CategoryId: allDrinks[i].CategoryId,
				DropMultiplier: allDrinks[i].DropMultiplier,
			},
		)
	}

	return nil
}
