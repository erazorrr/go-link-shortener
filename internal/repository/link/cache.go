package link

import (
	"context"
	"time"

	"github.com/redis/go-redis/v9"
)

type CacheLinkRepository struct {
	cache *redis.Client
}

func NewCacheLinkRepository(cache *redis.Client) *CacheLinkRepository {
	return &CacheLinkRepository{
		cache: cache,
	}
}

func (repository *CacheLinkRepository) getKey(code string) string {
	return "link:" + code
}

func (repository *CacheLinkRepository) SaveLinkMapping(ctx context.Context, code string, URL string) error {
	return repository.cache.Set(ctx, repository.getKey(code), URL, 24*time.Hour).Err()
}

func (repository *CacheLinkRepository) ResolveLinkMapping(ctx context.Context, code string) (string, error) {
	return repository.cache.Get(ctx, repository.getKey(code)).Result()
}
