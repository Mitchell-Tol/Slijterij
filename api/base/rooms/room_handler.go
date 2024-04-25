package rooms

import (
    "fmt"
    "net/http"
    "encoding/json"
)

const REGULAR = ""
const LOGIN = "login"

type RoomsHandler struct {
    URL string
}

func (h *RoomsHandler) CreateRoom(w http.ResponseWriter, r *http.Request) {
    room := &Room{}
    reqJsonErr := json.NewDecoder(r.Body).Decode(room)

    if reqJsonErr != nil {
        SendBadRequest(w)
        fmt.Printf("ERROR:\t%s\n", reqJsonErr)
        return
	}

    // TODO: Store room with generated token

    w.WriteHeader(http.StatusOK)
}

func (h *RoomsHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
    room := &Room{}
    reqJsonErr := json.NewDecoder(r.Body).Decode(room)

    if reqJsonErr != nil {
        SendBadRequest(w)
        fmt.Printf("ERROR:\t%s\n", reqJsonErr)
        return
    }

    tokenized := &TokenizedRoom{
        Id: room.Id,
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
