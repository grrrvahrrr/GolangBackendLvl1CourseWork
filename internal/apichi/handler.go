package apichi

import (
	"CourseWork/internal/dbbackend"
	"CourseWork/internal/entities"
	"context"
	"fmt"

	"github.com/google/uuid"
)

type Handlers struct {
	ds *dbbackend.DataStorage
}

func NewHandlers(ds *dbbackend.DataStorage) *Handlers {
	return &Handlers{
		ds: ds,
	}
}

type ApiUrlData struct {
	Id       uuid.UUID         `json:"id"`
	FullURL  string            `json:"fullurl"`
	ShortURL string            `json:"shorturl"`
	Data     map[string]string `json:"data"`
}

func (rt *Handlers) RedirectionHandle(ctx context.Context, ud ApiUrlData) (ApiUrlData, error) {
	//handle redirection to FullUrl here
	//Update short url data usage
	bud := entities.UrlData{
		ShortURL: ud.ShortURL,
	}
	newdub, err := rt.ds.WriteData(ctx, bud)
	if err != nil {
		return ApiUrlData{}, fmt.Errorf("error when creating: %w", err)
	}
	return ApiUrlData{
		Id:       newdub.Id,
		FullURL:  newdub.FullURL,
		ShortURL: newdub.ShortURL,
		Data:     newdub.Data,
	}, nil

}

func (rt *Handlers) GetDataHandle(ctx context.Context, ud ApiUrlData) (ApiUrlData, error) {
	//Handle request to get Data for ShortURL
	bud := entities.UrlData{
		ShortURL: ud.ShortURL,
	}
	newdub, err := rt.ds.ReadURL(ctx, bud)
	if err != nil {
		return ApiUrlData{}, fmt.Errorf("error when creating: %w", err)
	}
	return ApiUrlData(*newdub), nil
}

func (rt *Handlers) GenShortUrlHandle(ctx context.Context, ud ApiUrlData) (ApiUrlData, error) {
	//Handle request to generate Short URL
	bud := entities.UrlData{
		FullURL: ud.FullURL,
	}

	newdub, err := rt.ds.WriteURL(ctx, bud)
	if err != nil {
		return ApiUrlData{}, fmt.Errorf("error when creating: %w", err)
	}

	return ApiUrlData(*newdub), nil
}
