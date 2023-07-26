package router

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type User struct {
	ID           int
	Email        string
	PasswordHash string
}

func UserRouter() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/", handleRoute(GetUser))
	return router
}

func GetUser(w http.ResponseWriter, r *http.Request) error {
	user := &User{
		ID:           1,
		Email:        "pope@nor.way",
		PasswordHash: "asxmm21o12xiAH=)ASDSIJHs",
	}
	return JSONResponse(w, http.StatusOK, user)
}
