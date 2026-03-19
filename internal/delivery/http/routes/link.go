package routes

import (
	"github.com/erazorrr/go-link-shortener/internal/delivery/http/handlers/link"
	"github.com/go-chi/chi/v5"
)

func RegisterLinkRoutes(router chi.Router, handler *link.LinkHandler) {
	router.Route("/links", func(r chi.Router) {
		r.Post("/", handler.Create)
	})
}

func RegisterRedirectRoute(router chi.Router, handler *link.LinkHandler) {
	router.Get("/{code}", handler.Redirect)
}
