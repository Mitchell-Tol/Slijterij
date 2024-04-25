package bar

import (
    "fmt"
    "net/http"
    "encoding/json"
)

const REGULAR = ""
const LOGIN = "login"

type BarHandler struct {
    URL string
}

func (h *BarHandler) CreateBar(w http.ResponseWriter, r *http.Request) {
    bar := &Bar{}
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
    bar := &Bar{}
    reqJsonErr := json.NewDecoder(r.Body).Decode(bar)

    if reqJsonErr != nil {
        SendBadRequest(w)
        fmt.Printf("ERROR:\t%s\n", reqJsonErr)
        return
    }

    tokenized := &TokenizedBar{
        Id: bar.Id,
        Token: "ThisIsAGeneratedToken", // TODO: Retrieve token
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
