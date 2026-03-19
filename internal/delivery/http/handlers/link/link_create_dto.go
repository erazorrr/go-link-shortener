package link

import "time"

type LinkCreateDTO struct {
	Link struct {
		URL       string     `json:"url"`
		ExpiresAt *time.Time `json:"expires_at"`
	} `json:"link"`
}
