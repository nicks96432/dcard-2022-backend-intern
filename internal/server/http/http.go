package http

import "github.com/gofiber/fiber/v2"

import (
	"dcard-2022-backend-intern/internal/config"
	"dcard-2022-backend-intern/internal/handlers"
	"dcard-2022-backend-intern/internal/storage"
)

func New(db storage.Storage, c *config.Config) *fiber.App {
	app := fiber.New(fiber.Config{Prefork: true})

	h := handlers.Handlers{Db: db, Config: c}
	for _, r := range h.Routes() {
		app.Add(r.Method, r.Path, r.HandleFunc)
	}

	return app
}
