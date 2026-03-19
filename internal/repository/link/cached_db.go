package link

import (
	"context"

	"github.com/erazorrr/go-link-shortener/internal/domain"
)

type CachedDBLinkRepository struct {
	dbRepository    *DBLinkRepository
	cacheRepository *CacheLinkRepository

	cacheTokens chan struct{}
}

func NewCachedDBLinkRepository(dbRepository *DBLinkRepository, cacheRepository *CacheLinkRepository, maxConcurrentCacheWrites int) *CachedDBLinkRepository {
	cacheTokens := make(chan struct{}, maxConcurrentCacheWrites)
	return &CachedDBLinkRepository{
		dbRepository:    dbRepository,
		cacheRepository: cacheRepository,
		cacheTokens:     cacheTokens,
	}
}

func (repository *CachedDBLinkRepository) CreateLink(ctx context.Context, link *domain.Link) error {
	err := repository.dbRepository.CreateLink(ctx, link)
	if err != nil {
		return err
	}
	return repository.cacheRepository.SaveLinkMapping(ctx, link.Code, link.URL)
}

func (repository *CachedDBLinkRepository) GetLinkURLByCode(ctx context.Context, code string) (string, error) {
	URL, err := repository.cacheRepository.ResolveLinkMapping(ctx, code)
	if err == nil {
		return URL, nil
	}

	URL, err = repository.dbRepository.GetLinkURLByCode(ctx, code)
	if err != nil {
		return "", err
	}

	select {
	case repository.cacheTokens <- struct{}{}:
		go func() {
			defer func() { <-repository.cacheTokens }()
			repository.cacheRepository.SaveLinkMapping(context.Background(), code, URL)
		}()
	default:
	}

	return URL, nil
}
