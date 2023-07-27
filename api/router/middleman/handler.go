package middleman

import "net/http"

type apiHandler func(http.ResponseWriter, *http.Request) error

type APIError struct {
	Error string `json:"error"`
}

func NewHandler(f apiHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := f(w, r); err != nil {
			JSONResponse(w, http.StatusBadRequest, APIError{Error: err.Error()})
		}
	}
}
