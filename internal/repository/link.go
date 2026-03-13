package repository

import (
	"context"

	"github.com/erazorrr/go-link-shortener/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type LinkRepository struct {
	dbPool *pgxpool.Pool
}

func NewLinkRepository(dbPool *pgxpool.Pool) *LinkRepository {
	return &LinkRepository{dbPool: dbPool}
}

const createLinkQuery = `
	INSERT INTO link
	(code, url, expires_at)
	VALUES
	($1, $2, $3)
	RETURNING id, created_at
`

func (repository *LinkRepository) CreateLink(ctx context.Context, link *domain.Link) error {
	return repository.dbPool.QueryRow(ctx, createLinkQuery, link.Code, link.URL, link.ExpiresAt).Scan(&link.ID, &link.CreatedAt)
}
