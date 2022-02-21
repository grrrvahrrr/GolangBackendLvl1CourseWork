package entities

import "github.com/google/uuid"

type UrlData struct {
	Id       uuid.UUID
	FullURL  string
	ShortURL string
	Data     map[string]string
}
