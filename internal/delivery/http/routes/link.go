package routes

import (
	"github.com/erazorrr/go-link-shortener/internal/delivery/http/handlers"
	"github.com/go-chi/chi/v5"
)

func RegisterLinkRoutes(router chi.Router, handler *handlers.LinkHandler) {
	router.Route("/links", func(r chi.Router) {
		r.Post("/", handler.Create)
	})
}
