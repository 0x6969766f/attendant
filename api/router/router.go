package router

import (
	"github.com/0x6969766f/attendant/api/resource/game"
	"github.com/0x6969766f/attendant/api/resource/health"
	"github.com/0x6969766f/attendant/api/router/middleman"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func New() *chi.Mux {
	router := chi.NewRouter()

	router.Use(
		// middleman.JSONResponse,     // Set Content-Type headers (application/json)
		middleware.Logger,          // Log API request calls
		middleware.RedirectSlashes, // Redirect slashes to no slash URL versions
		middleware.Recoverer,       // Recover from panics without crashing
	)

	router.Get("/pulse", health.Check)

	router.Route("/game", func(r chi.Router) {
		gameAPI := game.New()
		r.Method("GET", "/", middleman.NewHandler(gameAPI.GetGames))
		// r.Method("GET", "/{gameID}", game.GetGame)
		// r.Method("POST", "/", game.CreateGame)
		// r.Method("PUT", "/{gameID}", game.UpdateGame)
		// r.Method("DELETE", "/{gameID}", game.DeleteGame)
	})

	return router
}
