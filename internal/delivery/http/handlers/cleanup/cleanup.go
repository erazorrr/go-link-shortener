package cleanup

import (
	"net/http"

	"github.com/erazorrr/go-link-shortener/internal/usecase/link"
)

type CleanupHandler struct {
	linkCleanupService *link.LinkCleanupService
}

func NewCleanupHandler(linkCleanupService *link.LinkCleanupService) *CleanupHandler {
	return &CleanupHandler{
		linkCleanupService: linkCleanupService,
	}
}

func (handler *CleanupHandler) Cleanup(w http.ResponseWriter, r *http.Request) {
	err := handler.linkCleanupService.Cleanup(r.Context())
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
