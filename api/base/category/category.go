package category

import (
	"net/http"
	"slijterij/api/generic"
	"slijterij/db"
)

func NewHandler(s *db.DataStore) *CategoryHandler {
	return &CategoryHandler{s}
}

func (h *CategoryHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	generic.EnableCors(&w)
	switch r.Method {
	case http.MethodGet:
		h.GetCategories(w, r)

	case http.MethodPost:
		h.PostCategory(w, r)

	case http.MethodPut:
		h.PutCategory(w, r)

	case http.MethodDelete:
		h.DeleteCategory(w, r)
	}
}
