package bar

import (
    "fmt"
    "net/http"
    "encoding/json"
    "slijterij/db"
    "slijterij/api/base/bar/model"
)

const REGULAR = ""
const LOGIN = "login"

type BarHandler struct {
    store   *db.DataStore
    URL     string
}

func (h *BarHandler) CreateBar(w http.ResponseWriter, r *http.Request) {
    bar := &model.Bar{}
    reqJsonErr := json.NewDecoder(r.Body).Decode(bar)

    if reqJsonErr != nil {
        SendBadRequest(w)
        fmt.Printf("ERROR:\t%s\n", reqJsonErr)
        return
	}

    // TODO: Store bar with generated token

    w.WriteHeader(http.StatusOK)
}

func (h *BarHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
    bar := &model.Bar{}
    reqJsonErr := json.NewDecoder(r.Body).Decode(bar)

    if reqJsonErr != nil {
        SendBadRequest(w)
        fmt.Printf("ERROR:\t%s\n", reqJsonErr)
        return
    }

    found, sqlErr := h.store.RetrieveBar(bar.Id)
    if sqlErr != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write(nil)
        fmt.Printf("Error:\t%s\n", sqlErr)
        return
    }

    tokenized := &model.TokenizedBar{
        Id: found.Id,
        Token: found.Token,
    }
    jsonResponse, resJsonErr := json.Marshal(tokenized)

    if resJsonErr != nil {
        SendBadRequest(w)
        fmt.Printf("ERROR:\t%s\n", resJsonErr)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write(jsonResponse)
}

func SendBadRequest(w http.ResponseWriter) {
    w.WriteHeader(http.StatusBadRequest)
    w.Write(nil)
}
