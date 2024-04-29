package bar

import (
    "fmt"
    "math/rand"
    "net/http"
    "encoding/json"
    "slijterij/db"
    "slijterij/api/base/bar/model"
    "slijterij/api/generic"
)

const REGULAR = ""
const LOGIN = "login"
const tokenLetters = "abcdefghijklmnopqrstuvwxyz"
const tokenMaxLength = 16

type BarHandler struct {
    store *db.DataStore
    URL string
}

func (h *BarHandler) CreateBar(w http.ResponseWriter, r *http.Request) {
    bar := &model.Bar{}
    reqJsonErr := json.NewDecoder(r.Body).Decode(bar)

    if reqJsonErr != nil {
        SendBadRequest(w)
        fmt.Printf("ERROR:\t%s\n", reqJsonErr)
        return
	}

    tokenSlice := make([]byte, tokenMaxLength)
    for i := range tokenSlice {
        tokenSlice[i] = tokenLetters[rand.Intn(len(tokenLetters))]
    }

    entity := &model.BarEntity{bar.Id, bar.Password, string(tokenSlice)}
    rowId, sqlErr := h.store.CreateBar(entity)
    if sqlErr != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write(generic.JSONError("API <--> Database communication error"))
        fmt.Printf("ERROR:\t%s\n", sqlErr)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write([]byte(fmt.Sprintf("Row %d created", rowId)))
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
        w.Write(generic.JSONError("API <--> Database communication error"))
        fmt.Printf("Error:\t%s\n", sqlErr)
        return
    }

    if bar.Password != found.Password {
        w.WriteHeader(http.StatusUnauthorized)
        w.Write(generic.JSONError("Incorrect password"))
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
    w.Write(generic.JSONError("Bad Request"))
}

