package link

import (
	"context"
	"errors"

	"github.com/erazorrr/go-link-shortener/internal/domain"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type DBLinkRepository struct {
	dbPool *pgxpool.Pool
}

func NewDBLinkRepository(dbPool *pgxpool.Pool) *DBLinkRepository {
	return &DBLinkRepository{dbPool: dbPool}
}

const createLinkQuery = `
	INSERT INTO link
	(code, url, expires_at)
	VALUES
	($1, $2, $3)
	RETURNING id, created_at
`

func (repository *DBLinkRepository) CreateLink(ctx context.Context, link *domain.Link) error {
	return repository.dbPool.QueryRow(ctx, createLinkQuery, link.Code, link.URL, link.ExpiresAt).Scan(&link.ID, &link.CreatedAt)
}

const findLinkByCodeQuery = `
	SELECT url
	FROM link
	WHERE
	code = $1
`

func (repository *DBLinkRepository) GetLinkURLByCode(ctx context.Context, code string) (string, error) {
	var url string
	err := repository.dbPool.QueryRow(ctx, findLinkByCodeQuery, code).Scan(&url)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return "", domain.ErrNotFound
		}
		return "", err
	}
	return url, nil
}
