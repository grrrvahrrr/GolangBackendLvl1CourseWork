package entities

import "github.com/google/uuid"

type UrlData struct {
	Id       uuid.UUID         `json:"id"`
	FullURL  string            `json:"fullurl"`
	ShortURL string            `json:"shorturl"`
	Data     map[string]string `json:"data"`
}
