package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/0x6969766f/attendant/config"
	"github.com/0x6969766f/attendant/database"
	"github.com/0x6969766f/attendant/migrations"
	"github.com/0x6969766f/attendant/router"
	"github.com/go-chi/chi/v5"
)

func Run(config config.Config) error {
	// setup database
	db, err := database.Connect(config.PSQL)
	if err != nil {
		return err
	}
	defer db.Close()

	// setup migrations
	err = database.MigrateFS(db, migrations.FS, ".")
	if err != nil {
		return err
	}

	// setup routes
	r := router.Setup()
	if err := chi.Walk(r, walk); err != nil {
		log.Panicf("Logging err: %s\n", err.Error()) // panic if there is an error
	}

	// start server
	fmt.Println("Server started on port", config.Server.Address)
	return http.ListenAndServe(config.Server.Address, r)
}

func walk(
	method, route string,
	handler http.Handler,
	middlewares ...func(http.Handler) http.Handler,
) error {
	log.Printf("%s %s\n", method, route) // Walk and print out all routes
	return nil
}
