package handlers

import "time"

type LinkCreateDTO struct {
	Link struct {
		Url       string     `json:"url" required:"true"`
		ExpiresAt *time.Time `json:"expires_at"`
	} `json:"link" required:"true"`
}
