package link

import (
	"errors"
	"net/http"

	"github.com/erazorrr/go-link-shortener/internal/domain"
	"github.com/go-chi/chi/v5"
)

func (handler *LinkHandler) Redirect(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	url, err := handler.linkQueryService.GetLinkURL(r.Context(), code)
	if err != nil {
		if errors.Is(err, domain.ErrNotFound) {
			http.Error(w, err.Error(), http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}
