package device

import(
	"net/http"
	"slijterij/db"
)

func NewHandler(s *db.DataStore) *DeviceHandler {
	return &DeviceHandler{s}
}

func (h *DeviceHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.CreateDevice(w, r)
	}
}
