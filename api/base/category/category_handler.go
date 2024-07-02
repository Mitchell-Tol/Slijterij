package category

import (
	"encoding/json"
	"net/http"
	"slijterij/db"
	"slijterij/api/base/category/categorymodel"
    "slijterij/api/generic"
)

type CategoryHandler struct {
	store *db.DataStore
}

func (h *CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	body := &categorymodel.CategoryByBar{}
	jsonErr := json.NewDecoder(r.Body).Decode(body)
	if jsonErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(generic.JSONError("Invalid JSON"))
		return
	}

	result, retrieveErr := h.store.GetAllCategories(body.BarId)
	if retrieveErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(generic.JSONError("An error occurred while retrieving categories"))
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

func (h *CategoryHandler) PostCategory(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (h *CategoryHandler) PutCategory(w http.ResponseWriter, r *http.Request) {
	// TODO
}

func (h *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	// TODO
}
