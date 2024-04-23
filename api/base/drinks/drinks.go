package drinks

import (
    "net/http"
)

type DrinksHandler struct{}

func NewHandler() *DrinksHandler {
    return &DrinksHandler{}
}

func (h *DrinksHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        w.Write([]byte("Retrieving drink list"))

    case http.MethodPost:
        w.Write([]byte("Created new type of drink"))

    case http.MethodPut:
        w.Write([]byte("Updated drink"))

    case http.MethodDelete:
        w.Write([]byte("Deleted drink"))
    }
}
