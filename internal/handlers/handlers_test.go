package handlers_test

import "testing"

import (
	"dcard-2022-backend-intern/internal/config"
	"dcard-2022-backend-intern/internal/handlers"
	"dcard-2022-backend-intern/internal/storage/sqlite"
)

func TestHandlers_Routes(t *testing.T) {
	db, err := sqlite.New(config.Default())
	if err != nil {
		t.Fatalf("cannot initialize database: %v\n", err)
	}
	h := handlers.Handlers{Db: db, Config: config.Default()}
	expected := 2
	if got := len(h.Routes()); got != expected {
		t.Fatalf("Handlers.Routes(): expected %d, got %d\n", expected, got)
	}
}
