package consoles

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/0x6969766f/attendant/api/router/middleman"
	"github.com/go-chi/chi/v5"
)

func (a *API) HandleGetConsoles(
	w http.ResponseWriter,
	r *http.Request,
) error {
	consoles, err := a.GetConsoles()
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return nil
	}
	return middleman.JSONResponse(w, http.StatusOK, consoles)
}

func (a *API) HandleGetConsole(
	w http.ResponseWriter,
	r *http.Request,
) error {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return nil
	}
	console, err := a.GetConsole(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Console not found", http.StatusUnprocessableEntity)
			return nil
		}
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return nil
	}
	return middleman.JSONResponse(w, http.StatusOK, console)
}

func (a *API) HandleCreateConsole(
	w http.ResponseWriter,
	r *http.Request,
) error {
	brand := r.FormValue("brand")
	if brand == "" {
		http.Error(w, "Brand is required", http.StatusBadRequest)
		return nil
	}

	model := r.FormValue("model")
	if model == "" {
		http.Error(w, "Model is required", http.StatusBadRequest)
		return nil
	}

	console, err := a.CreateConsole(brand, model)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return nil
	}
	return middleman.JSONResponse(w, http.StatusOK, console)
}

func (a *API) HandleUpdateConsole(
	w http.ResponseWriter,
	r *http.Request,
) error {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return nil
	}

	brand := r.FormValue("brand")
	model := r.FormValue("model")
	if brand == "" && model == "" {
		http.Error(
			w,
			"Required properties (brand or model) missing",
			http.StatusBadRequest,
		)
		return nil
	}

	console, err := a.UpdateConsole(id, brand, model)
	fmt.Println(err)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return nil
	}
	return middleman.JSONResponse(w, http.StatusOK, console)
}

func (a *API) HandleDeleteConsole(
	w http.ResponseWriter,
	r *http.Request,
) error {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return nil
	}
	err = a.DeleteConsole(id)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return nil
	}
	return middleman.JSONResponse(w, http.StatusOK, id)
}
