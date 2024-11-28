package entities

import (
	"errors"
	"net/url"
)

type CreateShortURLInput struct {
	ShortURL  string `form:"short_url" json:"short_url"`
	ExpiresAt int64  `form:"expires_at" json:"expires_at"`
	LongURL   string `form:"long_url" json:"long_url" binding:"required"`
}

func (input *CreateShortURLInput) Validate() error {
	_, err := url.ParseRequestURI(input.LongURL)
	if err != nil {
		return errors.New("invalid URL")
	}

	return nil
}
