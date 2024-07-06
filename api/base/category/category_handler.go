package category

import (
	"encoding/json"
	"net/http"
	"slijterij/api/base/category/categorymodel"
	"slijterij/api/generic"
	"slijterij/db"
)

type CategoryHandler struct {
	store *db.DataStore
}

func (h *CategoryHandler) GetCategories(w http.ResponseWriter, r *http.Request) {
	params := r.URL.Query()["barId"]
	if len(params) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(generic.JSONError("No barId parameter provided"))
		return
	}

	result, retrieveErr := h.store.GetAllCategories(params[0])
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
	body := &categorymodel.Category{}
	jsonParseErr := json.NewDecoder(r.Body).Decode(body)
	if jsonParseErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(generic.JSONError("Invalid JSON"))
		return
	}

	entity, queryErr := h.store.CreateCategory(body)
	if queryErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(generic.JSONError("An error occurred while storing the category"))
		return
	}

	jsonResp, jsonErr := json.Marshal(entity)
	if jsonErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(generic.JSONError("An error occurred mapping the created item to JSON"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}

func (h *CategoryHandler) PutCategory(w http.ResponseWriter, r *http.Request) {
	body := &categorymodel.UpdatedCategory{}
	jsonErr := json.NewDecoder(r.Body).Decode(body)
	if jsonErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(generic.JSONError("Invalid JSON"))
		return
	}

	result, queryErr := h.store.UpdateCategory(body)
	if queryErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(generic.JSONError("An error occurred while updating category"))
		return
	}

	jsonResp, jsonErr := json.Marshal(result)
	if jsonErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(generic.JSONError("An error occurred while mapping the updated category to JSON"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonResp)
}

func (h *CategoryHandler) DeleteCategory(w http.ResponseWriter, r *http.Request) {
	body := &categorymodel.CategoryId{}
	jsonErr := json.NewDecoder(r.Body).Decode(body)
	if jsonErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(generic.JSONError("Invalid JSON"))
		return
	}

	queryErr := h.store.DeleteCategory(body.Id)
	if queryErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(generic.JSONError("An error occurred while deleting the category"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(nil)
}
