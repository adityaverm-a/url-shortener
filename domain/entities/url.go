package entities

type URL struct {
	LongURL     string `json:"url"`
	ShortURL    string `json:"short_url"`
	AccessCount int    `json:"access_count"`
	ExpiresAt   int64  `json:"expire_at"`
}
