package base

import (
    "net/http"
)

type BaseHandler struct{}

func NewHandler() *BaseHandler {
    return &BaseHandler{}
}

func (h *BaseHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.WriteHeader(http.StatusOK)
    w.Write([]byte("Success"))
}
