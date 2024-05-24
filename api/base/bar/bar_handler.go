package bar

import (
    "fmt"
    "net/http"
    "encoding/json"
    "slijterij/db"
    "slijterij/api/base/bar/barmodel"
    "slijterij/api/generic"
	"github.com/google/uuid"
)

const REGULAR = ""
const LOGIN = "login"

type BarHandler struct {
    store *db.DataStore
    URL string
}

func (h *BarHandler) GetAllBars(w http.ResponseWriter, r *http.Request) {
    bars, retrieveErr := h.store.GetAllBars()
    if retrieveErr != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write(generic.JSONError("Internal server error"))
        return
    }

    var tokenized []barmodel.TokenizedBar
    for i := 0; i < len(bars); i++ {
        mapped := barmodel.MapEntityToTokenized(bars[i])
        tokenized = append(tokenized, mapped)
    }

    jsonResponse, jsonErr := json.Marshal(tokenized)
    if jsonErr != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write(generic.JSONError("An error occurred while mapping to JSON"))
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write(jsonResponse)
}

func (h *BarHandler) CreateBar(w http.ResponseWriter, r *http.Request) {
    bar := &barmodel.Bar{}
    reqJsonErr := json.NewDecoder(r.Body).Decode(bar)

    if reqJsonErr != nil {
        SendBadRequest(w)
        fmt.Printf("ERROR:\t%s\n", reqJsonErr)
        return
	}

    entity := &barmodel.BarEntity{bar.Id, bar.Password, uuid.New().String()}
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
    bar := &barmodel.Bar{}
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

    tokenized := &barmodel.TokenizedBar{
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

