package main

import (
	"fmt"
	"log"
)

import (
	"dcard-2022-backend-intern/internal/config"
	"dcard-2022-backend-intern/internal/server/http"
	"dcard-2022-backend-intern/internal/storage/sqlite"
)

func main() {
	config, err := config.From("./config.json")
	if err != nil {
		log.Printf("cannot read ./config.json: %v, using default options", err)
	}
	db, err := sqlite.New(config)
	if err != nil {
		log.Fatalf("cannot initialize database: %v", err)
	}

	app := http.New(db, config)
	log.Fatal(app.Listen(fmt.Sprintf(":%d", config.Port)))
}
