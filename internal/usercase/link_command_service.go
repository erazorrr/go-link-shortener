package usercase

import (
	"context"
	"math/rand"
	"strings"
	"time"

	"github.com/erazorrr/go-link-shortener/internal/domain"
)

type linkRepository interface {
	CreateLink(context.Context, *domain.Link) error
}

type LinkCommandService struct {
	linkRepository linkRepository
}

func NewLinkCommandService(linkRepository linkRepository) *LinkCommandService {
	return &LinkCommandService{linkRepository: linkRepository}
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
const codeLength = 10

func randomCode() string {
	b := make([]byte, codeLength)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	var sb strings.Builder
	sb.Write(b)
	return sb.String()
}

func (service *LinkCommandService) CreateLink(ctx context.Context, url string, expiresAt time.Time) (*domain.Link, error) {
	link := domain.Link{
		Code:      randomCode(),
		Url:       url,
		ExpiresAt: expiresAt,
	}
	err := service.linkRepository.CreateLink(ctx, &link)
	return &link, err
}
