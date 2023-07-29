package games

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/0x6969766f/attendant/api/router/middleman"
	"github.com/go-chi/chi/v5"
)

type Params map[string]string

// params := Params{}
// keys := []string{"console", "incollection", "name"}
// for _, key := range keys {
// 	if p := r.URL.Query().Get(key); p != "" {
// 		params[key] = p
// 	}
// }

func (a *API) HandleGetGames(
	w http.ResponseWriter,
	r *http.Request,
) error {
	games, err := a.GetGames()
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return nil
	}
	return middleman.JSONResponse(w, http.StatusOK, games)
}

func (a *API) HandleGetGame(
	w http.ResponseWriter,
	r *http.Request,
) error {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return nil
	}

	game, err := a.GetGame(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Game not found", http.StatusUnprocessableEntity)
			return nil
		}
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return nil
	}

	return middleman.JSONResponse(w, http.StatusOK, game)
}

func (a *API) HandleCreateGame(
	w http.ResponseWriter,
	r *http.Request,
) error {
	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "Missing name", http.StatusBadRequest)
		return nil
	}

	link := r.FormValue("link")
	image := r.FormValue("image")

	ownerID, err := strconv.Atoi(r.FormValue("owner_id"))
	if err != nil {
		http.Error(w, "Invalid owner_id", http.StatusBadRequest)
		return nil
	}

	consoleID, err := strconv.Atoi(r.FormValue("console_id"))
	if err != nil {
		http.Error(w, "Invalid console_id", http.StatusBadRequest)
		return nil
	}

	game, err := a.CreateGame(name, link, image, ownerID, consoleID)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return nil
	}
	return middleman.JSONResponse(w, http.StatusOK, game)
}

func (a *API) HandleUpdateGame(
	w http.ResponseWriter,
	r *http.Request,
) error {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return nil
	}

	name := r.FormValue("name")
	link := r.FormValue("link")
	image := r.FormValue("image")
	ownerID := r.FormValue("owner_id")
	consoleID := r.FormValue("console_id")

	game, err := a.UpdateGame(id, name, link, image, ownerID, consoleID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Game not found", http.StatusUnprocessableEntity)
			return nil
		}
		fmt.Println(err)
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return nil
	}
	return middleman.JSONResponse(w, http.StatusOK, game)
}

func (a *API) HandleDeleteGame(
	w http.ResponseWriter,
	r *http.Request,
) error {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return nil
	}

	err = a.DeleteGame(id)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return nil
	}

	return middleman.JSONResponse(w, http.StatusOK, id)
}
