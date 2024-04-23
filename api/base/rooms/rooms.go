package rooms

import (
	"net/http"
)

func NewHandler(url string) *RoomsHandler {
	return &RoomsHandler{url}
}

func (h *RoomsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		switch h.URL {
		case REGULAR:
			h.CreateRoom(w, r)

		case LOGIN:
			h.HandleLogin(w, r)
		}
	}
}
