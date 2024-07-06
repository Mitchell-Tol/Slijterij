package drinks

import (
	"net/http"
	"slijterij/db"
)

func NewHandler(s *db.DataStore) *DrinksHandler {
    return &DrinksHandler{s}
}

func (h *DrinksHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        h.GetAllDrinks(w, r)

    case http.MethodPost:
        h.CreateDrink(w, r)

    case http.MethodPut:
		h.UpdateDrink(w, r)

    case http.MethodDelete:
        h.DeleteDrink(w, r)
    }
}
