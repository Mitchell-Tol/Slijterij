package crash

import (
    "encoding/json"
    "fmt"
    "net/http"
    "slijterij/api/base/drinks/crash/crashmodel"
    "slijterij/api/generic"
    "slijterij/db"
)

type CrashHandler struct {
    store *db.DataStore
}

func (h *CrashHandler) CrashProducts(w http.ResponseWriter, r *http.Request) {
    request := &crashmodel.DrinkCrash{}
    reqJsonErr := json.NewDecoder(r.Body).Decode(request)
    if reqJsonErr != nil {
        w.WriteHeader(http.StatusBadRequest)
        w.Write(generic.JSONError("Bad Request: Invalid JSON"))
        return
    }

    updated := []string{}
    for i := 0; i < len(request.Ids); i++ {
        entity, error := h.store.GetDrink(request.Ids[i])
        if error == nil {
            updated = append(updated, entity.Id)
            multiplier := 1 - (request.DropPercentage / 100)
            entity.CurrentPrice = entity.CurrentPrice * multiplier
            h.store.UpdateDrink(entity)
        }
    }

    responseModel := crashmodel.CrashResponse{Updated: updated}
    response, resJsonErr := json.Marshal(responseModel)
    if resJsonErr != nil {
        w.WriteHeader(http.StatusInternalServerError)
        w.Write(generic.JSONError(fmt.Sprintf("%v", resJsonErr)))
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Write(response)
}
