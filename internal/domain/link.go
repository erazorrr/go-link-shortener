package domain

import (
	"errors"
	"time"
)

type Link struct {
	ID        int64
	Code      string
	URL       string
	CreatedAt *time.Time
	ExpiresAt *time.Time
}

var ErrNotFound = errors.New("not found")
