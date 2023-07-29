package router

import (
	"database/sql"

	"github.com/0x6969766f/attendant/api/resource/consoles"
	"github.com/0x6969766f/attendant/api/resource/games"
	"github.com/0x6969766f/attendant/api/resource/health"
	"github.com/0x6969766f/attendant/api/resource/image"
	"github.com/0x6969766f/attendant/api/resource/owners"
	"github.com/0x6969766f/attendant/api/router/middleman"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func New(db *sql.DB) *chi.Mux {
	router := chi.NewRouter()

	router.Use(
		middleware.Logger,          // Log API request calls
		middleware.RedirectSlashes, // Redirect slashes to no slash URL versions
		middleware.Recoverer,       // Recover from panics without crashing
	)

	router.Get("/pulse", health.Check)

	router.Route("/games", func(r chi.Router) {
		gamesAPI := games.New(db)
		r.Method(
			"GET", "/",
			middleman.NewHandler(gamesAPI.HandleGetGames),
		)
		r.Method(
			"POST", "/",
			middleman.NewHandler(gamesAPI.HandleCreateGame),
		)
		r.Method(
			"GET", "/{id}",
			middleman.NewHandler(gamesAPI.HandleGetGame),
		)
		r.Method(
			"PATCH", "/{id}",
			middleman.NewHandler(gamesAPI.HandleUpdateGame),
		)
		r.Method(
			"DELETE", "/{id}",
			middleman.NewHandler(gamesAPI.HandleDeleteGame),
		)
	})

	router.Route("/owners", func(r chi.Router) {
		ownersAPI := owners.New(db)
		r.Method(
			"GET", "/",
			middleman.NewHandler(ownersAPI.HandleGetOwners),
		)
		r.Method(
			"POST", "/",
			middleman.NewHandler(ownersAPI.HandleCreateOwner),
		)
		r.Method(
			"GET", "/{id}",
			middleman.NewHandler(ownersAPI.HandleGetOwner),
		)
		r.Method(
			"PATCH", "/{id}",
			middleman.NewHandler(ownersAPI.HandleUpdateOwner),
		)
		r.Method(
			"DELETE", "/{id}",
			middleman.NewHandler(ownersAPI.HandleDeleteOwner),
		)
	})

	router.Route("/consoles", func(r chi.Router) {
		consolesAPI := consoles.New(db)
		r.Method(
			"GET", "/",
			middleman.NewHandler(consolesAPI.HandleGetConsoles),
		)
		r.Method(
			"POST", "/",
			middleman.NewHandler(consolesAPI.HandleCreateConsole),
		)
		r.Method(
			"GET", "/{id}",
			middleman.NewHandler(consolesAPI.HandleGetConsole),
		)
		r.Method(
			"PATCH", "/{id}",
			middleman.NewHandler(consolesAPI.HandleUpdateConsole),
		)
		r.Method(
			"DELETE", "/{id}",
			middleman.NewHandler(consolesAPI.HandleDeleteConsole),
		)
	})

	router.Route("/images", func(r chi.Router) {
		imageAPI := image.New()
		r.Method(
			"GET", "/{filename}",
			middleman.NewHandler(imageAPI.GetImage),
		)
	})

	return router
}
