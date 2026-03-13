package handlers

import (
	"bytes"
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

	link, err := handler.linkCommandService.CreateLink(r.Context(), body.Link.URL, body.Link.ExpiresAt)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var buf bytes.Buffer
	response := NewLinkCreatedDTO(link)
	err = json.NewEncoder(&buf).Encode(response)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write(buf.Bytes())
}
