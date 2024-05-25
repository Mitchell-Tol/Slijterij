package device

import (
    "encoding/json"
    "net/http"
    "slijterij/api/base/device/devicemodel"
    "slijterij/api/generic"
    "slijterij/db"
)

type DeviceHandler struct {
    store *db.DataStore
}

func (h *DeviceHandler) CreateDevice(w http.ResponseWriter, r *http.Request) {
    body := &devicemodel.Device{}
    jsonParseErr := json.NewDecoder(r.Body).Decode(body)
    if jsonParseErr != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write(generic.JSONError("Invalid JSON"))
        return
    }

    entity, queryErr := h.store.CreateDevice(body)
    if queryErr != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write(generic.JSONError("Internal server error"))
        return
    }

    jsonResp, jsonErr := json.Marshal(entity)
    if jsonErr != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write(generic.JSONError("Error while parsing created item to JSON"))
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write(jsonResp)
}
