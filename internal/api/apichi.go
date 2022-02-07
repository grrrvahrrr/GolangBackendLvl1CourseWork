package apichi

import (
	"CourseWork/internal/dbbackend"
	"net/http"
)

//Package responsible for primary sevice usage - allow user to input shortURL into browser and be redirected to fullURL
//For front page handle request to generate ShortURL from FullUrl and get Data from ShortURL
//Use go-chi router here to process requests that come to the server

type ChiRouter struct {
	//chi Mux
	ds *dbbackend.DataStorage
}

func NewRouter(ds *dbbackend.DataStorage) *ChiRouter {
	return &ChiRouter{
		ds: ds,
	}
}

func (rt *ChiRouter) RedirectionHandle(w http.ResponseWriter, r *http.Request) {
	//handle redirection to FullUrl here
}

func (rt *ChiRouter) GetDataHandle(w http.ResponseWriter, r *http.Request) {
	//Handle request to get Data for ShortURL
}

func (rt *ChiRouter) GenShortUrlHandle(w http.ResponseWriter, r *http.Request) {
	//Handle request to generate Short URL
}
