package generic

import (
    "encoding/json"
)

type Error struct {
    Message string `json:"message"`
}

func JSONError(text string) []byte {
    jsonErr, _ := json.Marshal(Error{text})
    return jsonErr
}

