package order

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slijterij/api/base/order/ordermodel"
	"slijterij/api/generic"
	"slijterij/db"
)

type OrderHandler struct {
	store *db.DataStore
}

func (h *OrderHandler) GetOrders(w http.ResponseWriter, r *http.Request) {
	body := &ordermodel.OrderByDevice{}
	jsonErr := json.NewDecoder(r.Body).Decode(body)
	if jsonErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(generic.JSONError("Invalid JSON"))
		return
	}

	result, queryErr := h.store.GetAllOrders(body.DeviceId)
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
        w.WriteHeader(http.StatusInternalServerError)
        w.Write(generic.JSONError("Internal server error"))
        return
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
