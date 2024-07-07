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

    jsonResponse, jsonErr := json.Marshal(bars)
    if jsonErr != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write(generic.JSONError("An error occurred while mapping to JSON"))
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write(jsonResponse)
}

func (h *BarHandler) GetBarById(w http.ResponseWriter, r *http.Request) {
    params := r.URL.Query()["id"]
    if len(params) == 0 {
        w.WriteHeader(http.StatusBadRequest)
        w.Write(generic.JSONError("No bar id provided"))
        return
    }

    result, retrieveErr := h.store.BarById(params[0])
    if retrieveErr != nil {
        fmt.Printf("GetBarById %s: %v\n", params[0], retrieveErr)
        w.WriteHeader(http.StatusNotFound)
        w.Write(generic.JSONError(fmt.Sprintf("Bar with id %s not found", params[0])))
        return
    }

    jsonResp, jsonErr := json.Marshal(result)
    if jsonErr != nil {
        fmt.Printf("%s", jsonErr)
        w.WriteHeader(http.StatusInternalServerError)
        w.Write(generic.JSONError(fmt.Sprintf("%s", jsonErr)))
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write(jsonResp)
}

func (h *BarHandler) CreateBar(w http.ResponseWriter, r *http.Request) {
    bar := &barmodel.Bar{SuperAdmin: 0}
    reqJsonErr := json.NewDecoder(r.Body).Decode(bar)

    if reqJsonErr != nil {
        SendBadRequest(w)
        fmt.Printf("ERROR:\t%s\n", reqJsonErr)
        return
	}

    entity := &barmodel.BarEntity{
		Id: uuid.New().String(), 
		Name: bar.Name, 
		Password: bar.Password, 
		Token: uuid.New().String(),
		SuperAdmin: bar.SuperAdmin,
	}
    result, sqlErr := h.store.CreateBar(entity)
    if sqlErr != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write(generic.JSONError("API <--> Database communication error"))
        fmt.Printf("ERROR:\t%s\n", sqlErr)
        return
    }

	jsonRes, jsonErr := json.Marshal(result)
	if jsonErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(generic.JSONError("An error occurred while mapping result to JSON"))
		return
	}

    w.WriteHeader(http.StatusOK)
    w.Write(jsonRes)
}

func (h *BarHandler) UpdateBar(w http.ResponseWriter, r *http.Request) {
	entity := &barmodel.BarEntity{}
	reqJsonErr := json.NewDecoder(r.Body).Decode(entity)
	if reqJsonErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(generic.JSONError("Invalid JSON"))
		return
	}

	_, queryErr := h.store.UpdateBar(entity)
	if queryErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(generic.JSONError("Internal server error"))
		return
	}

	jsonRes, jsonMapErr := json.Marshal(entity)
	if jsonMapErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(generic.JSONError("Error while mapping response to JSON"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(jsonRes)
}

func (h *BarHandler) DeleteBar(w http.ResponseWriter, r *http.Request) {
	rBody := &barmodel.BarId{}
	jsonErr := json.NewDecoder(r.Body).Decode(rBody)
	if jsonErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write(generic.JSONError("Invalid JSON"))
		return
	}

	queryErr := h.store.DeleteBar(rBody.Id)
	if queryErr != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write(generic.JSONError("Internal server error"))
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write(nil)
}

func (h *BarHandler) HandleLogin(w http.ResponseWriter, r *http.Request) {
    bar := &barmodel.LoginBar{}
    reqJsonErr := json.NewDecoder(r.Body).Decode(bar)

    if reqJsonErr != nil {
        SendBadRequest(w)
        fmt.Printf("ERROR:\t%s\n", reqJsonErr)
        return
    }

    found, sqlErr := h.store.RetrieveBar(bar.Name)
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
		Name: found.Name,
        Token: found.Token,
		SuperAdmin: found.SuperAdmin,
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
