package domain

import "time"

type Link struct {
	Id        int64
	Code      string
	Url       string
	CreatedAt *time.Time
	ExpiresAt *time.Time
}
