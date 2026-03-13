package usecase

import (
	"context"
	"crypto/rand"
	"math/big"
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

const (
	codeLength = 10
	charset    = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func randomCode() string {
	result := make([]byte, codeLength)
	for i := range codeLength {
		index, err := rand.Int(rand.Reader, big.NewInt(int64(len(charset))))
		if err != nil {
			panic(err)
		}
		result[i] = charset[int(index.Int64())]
	}
	return string(result)
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
