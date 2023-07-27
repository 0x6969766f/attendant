package middleman

import (
	"encoding/json"
	"net/http"
)

func JSONResponse(w http.ResponseWriter, status int, value any) error {
	w.Header().Add("Content-Type", "application/json;charset=utf-8")
	w.WriteHeader(status)
	return json.NewEncoder(w).Encode(value)
}
