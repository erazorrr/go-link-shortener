package handlers

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/erazorrr/go-link-shortener/internal/usercase"
)

type LinkHandler struct {
	linkCommandService *usercase.LinkCommandService
}

func NewLinkHandler(linkCommandService *usercase.LinkCommandService) *LinkHandler {
	return &LinkHandler{linkCommandService: linkCommandService}
}

func (handler *LinkHandler) Create(w http.ResponseWriter, r *http.Request) {
	var body LinkCreateDTO
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	link, err := handler.linkCommandService.CreateLink(context.Background(), body.Link.Url, body.Link.ExpiresAt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response := NewLinkCreatedDTO(link)
	json.NewEncoder(w).Encode(response)
}
