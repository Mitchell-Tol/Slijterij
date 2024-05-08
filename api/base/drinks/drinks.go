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
        w.Write([]byte("Retrieving drink list"))

    case http.MethodPost:
        h.CreateDrink(w, r)

    case http.MethodPut:
        w.Write([]byte("Updated drink"))

    case http.MethodDelete:
        w.Write([]byte("Deleted drink"))
    }
}

