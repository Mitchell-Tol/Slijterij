package order

import (
	"net/http"
	"slijterij/api/generic"
	"slijterij/db"
)

func NewHandler(s *db.DataStore) *OrderHandler {
	return &OrderHandler{s}
}

func (h *OrderHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	generic.EnableCors(&w)
	switch r.Method {
	case http.MethodGet:
		h.GetOrders(w, r)

	case http.MethodPost:
		h.PostOrder(w, r)

	case http.MethodPut:
		h.UpdateOrder(w, r)

	case http.MethodDelete:
		h.DeleteOrder(w, r)
	}
}
