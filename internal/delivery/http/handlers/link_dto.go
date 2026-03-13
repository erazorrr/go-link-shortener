package handlers

import (
	"time"

	"github.com/erazorrr/go-link-shortener/internal/domain"
)

type LinkDTO struct {
	ID        int64      `json:"id" required:"true"`
	Code      string     `json:"code" required:"true"`
	URL       string     `json:"url" required:"true"`
	CreatedAt *time.Time `json:"created_at" required:"true"`
	ExpiresAt *time.Time `json:"expires_at"`
}

type LinkCreatedDTO struct {
	Link LinkDTO `json:"link" required:"true"`
}

func NewLinkCreatedDTO(link *domain.Link) *LinkCreatedDTO {
	return &LinkCreatedDTO{
		Link: LinkDTO{
			ID:        link.ID,
			Code:      link.Code,
			URL:       link.URL,
			CreatedAt: link.CreatedAt,
			ExpiresAt: link.ExpiresAt,
		},
	}
}
