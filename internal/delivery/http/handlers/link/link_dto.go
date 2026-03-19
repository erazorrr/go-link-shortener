package link

import (
	"time"

	"github.com/erazorrr/go-link-shortener/internal/domain"
)

type LinkDTO struct {
	ID        int64      `json:"id"`
	Code      string     `json:"code"`
	URL       string     `json:"url"`
	CreatedAt *time.Time `json:"created_at"`
	ExpiresAt *time.Time `json:"expires_at"`
}

type LinkCreatedDTO struct {
	Link LinkDTO `json:"link"`
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
