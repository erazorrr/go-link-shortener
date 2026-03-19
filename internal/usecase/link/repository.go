package link

import (
	"context"

	"github.com/erazorrr/go-link-shortener/internal/domain"
)

type linkRepository interface {
	CreateLink(context.Context, *domain.Link) error
	GetLinkURLByCode(context.Context, string) (string, error)
}
