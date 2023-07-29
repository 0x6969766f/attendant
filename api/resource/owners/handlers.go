package owners

import (
	"database/sql"
	"errors"
	"net/http"
	"strconv"

	"github.com/0x6969766f/attendant/api/router/middleman"
	"github.com/go-chi/chi/v5"
)

func (a *API) HandleGetOwners(
	w http.ResponseWriter,
	r *http.Request,
) error {
	owners, err := a.GetOwners()
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return nil
	}
	return middleman.JSONResponse(w, http.StatusOK, owners)
}

func (a *API) HandleGetOwner(
	w http.ResponseWriter,
	r *http.Request,
) error {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return nil
	}
	owner, err := a.GetOwner(id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			http.Error(w, "Owner not found", http.StatusUnprocessableEntity)
			return nil
		}
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return nil
	}
	return middleman.JSONResponse(w, http.StatusOK, owner)
}

func (a *API) HandleCreateOwner(
	w http.ResponseWriter,
	r *http.Request,
) error {
	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return nil
	}

	owner, err := a.CreateOwner(name)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return nil
	}
	return middleman.JSONResponse(w, http.StatusOK, owner)
}

func (a *API) HandleUpdateOwner(
	w http.ResponseWriter,
	r *http.Request,
) error {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return nil
	}

	name := r.FormValue("name")
	if name == "" {
		http.Error(w, "Name is required", http.StatusBadRequest)
		return nil
	}

	owner, err := a.UpdateOwner(id, name)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return nil
	}
	return middleman.JSONResponse(w, http.StatusOK, owner)
}

func (a *API) HandleDeleteOwner(
	w http.ResponseWriter,
	r *http.Request,
) error {
	id, err := strconv.Atoi(chi.URLParam(r, "id"))
	if err != nil {
		http.Error(w, "Invalid id", http.StatusBadRequest)
		return nil
	}
	err = a.DeleteOwner(id)
	if err != nil {
		http.Error(w, "Something went wrong", http.StatusInternalServerError)
		return nil
	}
	return middleman.JSONResponse(w, http.StatusOK, id)
}
