package entities

import "time"

type URL struct {
	LongURL     string    `json:"url"`
	ShortURL    string    `json:"short_url"`
	AccessCount int       `json:"access_count"`
	ExpiresAt   time.Time `json:"expire_at"`
}
