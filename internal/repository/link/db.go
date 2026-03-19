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
	SELECT id, url, expires_at, created_at
	FROM link
	WHERE
	code = $1
	AND (expires_at IS NULL OR expires_at > now())
`

func (repository *DBLinkRepository) GetLinkByCode(ctx context.Context, code string) (*domain.Link, error) {
	link := domain.Link{
		Code: code,
	}
	err := repository.dbPool.QueryRow(ctx, findLinkByCodeQuery, code).Scan(&link.ID, &link.URL, &link.ExpiresAt, &link.CreatedAt)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return nil, domain.ErrNotFound
		}
		return nil, err
	}
	return &link, nil
}

const deleteExpiredQuery = `
	DELETE FROM link
	WHERE expires_at IS NOT NULL AND expires_at <= now()
`

func (repository *DBLinkRepository) DeleteExpired(ctx context.Context) error {
	_, err := repository.dbPool.Exec(ctx, deleteExpiredQuery)
	return err
}
