package bar

import (
	"net/http"
	"slijterij/db"
)

func NewHandler(s *db.DataStore, url string) *BarHandler {
    return &BarHandler{s, url}
}

func (h *BarHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    switch r.Method {
    case http.MethodGet:
        params := r.URL.Query()["id"]
        if len(params) > 0 {
            h.GetBarById(w, r)
        } else {
            h.GetAllBars(w, r)
        }
        
    case http.MethodPost:
        switch h.URL {
        case REGULAR:
            h.CreateBar(w, r)

        case LOGIN:
            h.HandleLogin(w, r)
        }

	case http.MethodPut:
		h.UpdateBar(w, r)

	case http.MethodDelete:
		h.DeleteBar(w, r)
    }
}
