package handlers

import (
	"time"

	"github.com/erazorrr/go-link-shortener/internal/domain"
)

type LinkDTO struct {
	Id        int64     `json:"id" required:"true"`
	Code      string    `json:"code" required:"true"`
	Url       string    `json:"url" required:"true"`
	CreatedAt time.Time `json:"created_at" required:"true"`
	ExpiresAt time.Time `json:"expires_at" required:"true"`
}

type LinkCreatedDTO struct {
	Link LinkDTO `json:"link" required:"true"`
}

func NewLinkCreatedDTO(link *domain.Link) *LinkCreatedDTO {
	return &LinkCreatedDTO{
		Link: LinkDTO{
			Id:        link.Id,
			Code:      link.Code,
			Url:       link.Url,
			CreatedAt: link.CreatedAt,
			ExpiresAt: link.ExpiresAt,
		},
	}
}
