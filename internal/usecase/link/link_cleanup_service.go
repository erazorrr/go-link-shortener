package link

import "context"

type linkCleanupRepository interface {
	DeleteExpired(ctx context.Context) error
}

type LinkCleanupService struct {
	linkRepository linkCleanupRepository
}

func NewLinkCleanupService(linkRepository linkCleanupRepository) *LinkCleanupService {
	return &LinkCleanupService{
		linkRepository: linkRepository,
	}
}

func (service *LinkCleanupService) Cleanup(ctx context.Context) error {
	return service.linkRepository.DeleteExpired(ctx)
}
