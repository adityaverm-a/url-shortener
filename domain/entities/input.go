package entities

import (
	"errors"
	"net/url"
)

type GetShortURLInput struct {
	ShortURL string `json:"short_url" uri:"short_url" binding:"required"`
}

type CreateShortURLInput struct {
	CustomShortURL string `form:"custom_short_url" json:"custom_short_url"`
	ExpiresAt      int64  `form:"expires_at" json:"expires_at"`
	LongURL        string `form:"long_url" json:"long_url" binding:"required"`
}

func (input *CreateShortURLInput) Validate() error {
	_, err := url.ParseRequestURI(input.LongURL)
	if err != nil {
		return errors.New("invalid URL")
	}

	return nil
}
