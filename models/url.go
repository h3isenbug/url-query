package models

import "time"

type URL struct {
	ShortPath string    `json:"shortPath" db:"short_path"`
	LongURL   string    `json:"longURL" db:"long_url"`
	Username  string    `db:"username"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type AnalyticsItem struct {
	ShortPath string `json:"shortPath"`
	UserAgent string `json:"userAgent"`
	ETag      string `json:"etag"`
}
