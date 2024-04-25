package bar

import (
    "net/http"
)

func NewHandler(url string) *BarHandler {
    return &BarHandler{url}
}

func (h *BarHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodPost:
        switch h.URL {
        case REGULAR:
            h.CreateBar(w, r)

        case LOGIN:
            h.HandleLogin(w, r)
        }
    }
}
