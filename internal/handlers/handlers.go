package handlers

import (
	"net/http"

	"github.com/gofiber/fiber/v2"
)

import (
	"dcard-2022-backend-intern/internal/config"
	"dcard-2022-backend-intern/internal/storage"
)

type Handlers struct {
	Db     storage.Storage
	Config *config.Config
}

type Handler struct {
	Path       string
	Method     string
	HandleFunc fiber.Handler
}

func (h *Handlers) Routes() []Handler {
	return []Handler{
		{
			Path:       "/api/v1/urls",
			Method:     http.MethodPost,
			HandleFunc: h.postNewUrl,
		},
		{
			Path:       "/:url",
			Method:     http.MethodGet,
			HandleFunc: h.getRoot,
		},
	}
}
