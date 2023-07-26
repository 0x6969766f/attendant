package main

import (
	"log"

	"github.com/0x6969766f/attendant/app"
	"github.com/0x6969766f/attendant/config"
)

func main() {
	config, err := config.Load()
	if err != nil {
		panic(err)
	}
	log.Fatal(app.Run(config))
}
