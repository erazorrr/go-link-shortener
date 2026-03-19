package routes

import (
	"github.com/erazorrr/go-link-shortener/internal/delivery/http/handlers"
	"github.com/erazorrr/go-link-shortener/internal/delivery/http/handlers/cleanup"
	"github.com/go-chi/chi/v5"
)

func RegisterCleanupRoutes(router chi.Router, apiKey string, handler *cleanup.CleanupHandler) {
	router.Group(func(r chi.Router) {
		r.Use(handlers.CreateInternalOnlyMiddleware(apiKey))
		r.Post("/internal/cleanup", handler.Cleanup)
	})
}
