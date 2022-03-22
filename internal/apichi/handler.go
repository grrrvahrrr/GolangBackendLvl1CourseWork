package apichi

import (
	"CourseWork/internal/dbbackend"
	"CourseWork/internal/entities"
	"context"
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
	FullURL  string `json:"fullurl"`
	ShortURL string `json:"shorturl"`
	AdminURL string `json:"adminurl"`
	Data     string `json:"data"`
	IP       string `json:"ip"`
	IPData   string `json:"ipdata"`
}

func (rt *Handlers) RedirectionHandle(ctx context.Context, surl string, ip string) (ApiUrlData, error) {
	//handle redirection to FullUrl here
	//Update short url data usage
	bud := entities.UrlData{
		ShortURL: surl,
		IP:       ip,
	}
	newdub, err := rt.ds.WriteData(ctx, bud)
	if err != nil {
		return ApiUrlData{}, err
	}
	return ApiUrlData{
		FullURL:  newdub.FullURL,
		ShortURL: newdub.ShortURL,
		Data:     newdub.Data,
	}, nil

}

func (rt *Handlers) GetDataHandle(ctx context.Context, ud ApiUrlData) (ApiUrlData, string, error) {
	//Handle request to get Data for ShortURL
	bud := entities.UrlData{
		AdminURL: ud.AdminURL,
	}
	newdub, err := rt.ds.ReadURL(ctx, bud)
	if err != nil {
		return ApiUrlData{}, "", err
	}

	bud.ShortURL = newdub.ShortURL

	ipdata, err := rt.ds.GetIPData(ctx, bud)
	if err != nil {
		return ApiUrlData{}, "", err
	}

	return ApiUrlData(*newdub), ipdata, nil
}

func (rt *Handlers) GenShortUrlHandle(ctx context.Context, ud ApiUrlData) (ApiUrlData, error) {
	//Handle request to generate Short URL
	bud := entities.UrlData{
		FullURL: ud.FullURL,
	}

	newdub, err := rt.ds.WriteURL(ctx, bud)
	if err != nil {
		return ApiUrlData{}, err
	}

	return ApiUrlData(*newdub), nil
}
