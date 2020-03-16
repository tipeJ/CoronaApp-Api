package models

import "time"

type News struct {
	Source       string    `json:"source"`
	Url          string    `json:"url"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Thumbnailurl string    `json:"thumbnailurl"`
	PubDate      time.Time `json:"pubDate"`
}
