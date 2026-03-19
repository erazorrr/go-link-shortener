package link

import (
	"bytes"
	"encoding/json"
	"net/http"
)

const maxRequestBodySizeBytes = 1 << 20

func (handler *LinkHandler) Create(w http.ResponseWriter, r *http.Request) {
	var body LinkCreateDTO
	decoder := json.NewDecoder(http.MaxBytesReader(w, r.Body, maxRequestBodySizeBytes))
	decoder.DisallowUnknownFields()
	err := decoder.Decode(&body)
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

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(buf.Bytes())
}
