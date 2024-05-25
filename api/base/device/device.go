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
	case http.MethodGet:
		h.GetDevices(w, r)

    case http.MethodPost:
        h.CreateDevice(w, r)

	case http.MethodPut:
		h.UpdateDevice(w, r)
    }
}
