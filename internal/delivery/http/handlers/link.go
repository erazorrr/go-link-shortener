package handlers

import (
	"encoding/json"
	"net/http"

	"github.com/erazorrr/go-link-shortener/internal/usecase"
)

type LinkHandler struct {
	linkCommandService *usecase.LinkCommandService
}

func NewLinkHandler(linkCommandService *usecase.LinkCommandService) *LinkHandler {
	return &LinkHandler{linkCommandService: linkCommandService}
}

func (handler *LinkHandler) Create(w http.ResponseWriter, r *http.Request) {
	var body LinkCreateDTO
	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	link, err := handler.linkCommandService.CreateLink(r.Context(), body.Link.Url, body.Link.ExpiresAt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	response := NewLinkCreatedDTO(link)
	json.NewEncoder(w).Encode(response)
}
