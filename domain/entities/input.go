package entities

import (
	"errors"
	"net/url"
)

type GetShortURLInput struct {
	ShortURL string `json:"short_url" uri:"short_url" binding:"required"`
}

type CreateShortURLInput struct {
	LongURL        string `form:"long_url" json:"long_url" binding:"required"`
	CustomShortURL string `form:"custom_short_url" json:"custom_short_url"`
	TTL            int64  `form:"ttl" json:"ttl"` // TTL in seconds (optional)
}

func (input *CreateShortURLInput) Validate() error {
	_, err := url.ParseRequestURI(input.LongURL)
	if err != nil {
		return errors.New("invalid URL")
	}

	return nil
}
