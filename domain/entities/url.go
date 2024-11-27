package entities

type URL struct {
	LongURL     string `json:"url"`
	ShortURL    string `json:"short"`
	AccessCount int    `json:"total"`
	ExpiresAt   int64  `json:"currency_unit"`
}
