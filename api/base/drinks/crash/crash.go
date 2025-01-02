package crash

import (
	"net/http"
	"slijterij/db"
)

func NewHandler(s *db.DataStore) *CrashHandler {
	return &CrashHandler{s}
}

func (h *CrashHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPut:
		h.CrashProducts(w, r)
	}
}
