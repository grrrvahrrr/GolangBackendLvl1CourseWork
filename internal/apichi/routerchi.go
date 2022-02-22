package apichi

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

//Package responsible for primary sevice usage - allow user to input shortURL into browser and be redirected to fullURL
//For front page handle request to generate ShortURL from FullUrl and get Data from ShortURL
//Use go-chi router here to process requests that come to the server

type ChiRouter struct {
	*chi.Mux
	hs *Handlers
}

func NewRouter(hs *Handlers) *ChiRouter {
	r := chi.NewRouter()
	ret := &ChiRouter{
		hs: hs,
	}

	r.Use(middleware.Logger)
	r.Get("/{shortURL}", ret.Redirection)
	r.Post("/shortenURL", ret.GenShortUrl)
	r.Post("/getData", ret.GetData)

	ret.Mux = r

	return ret
}

func (ApiUrlData) Bind(r *http.Request) error {
	return nil
}

func (ApiUrlData) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (rt *ChiRouter) Redirection(w http.ResponseWriter, r *http.Request) {
	surl := chi.URLParam(r, "shortURL")

	ud := ApiUrlData{
		ShortURL: surl,
	}
	nud, err := rt.hs.RedirectionHandle(r.Context(), ud)
	if err != nil {
		err = render.Render(w, r, ErrRender(err))
		if err != nil {
			log.Println(err)
		}
		return
	}

	err = render.Render(w, r, nud)
	if err != nil {
		log.Println(err)
	}
}

func (rt *ChiRouter) GenShortUrl(w http.ResponseWriter, r *http.Request) {
	urldata := ApiUrlData{}
	if err := render.Bind(r, &urldata); err != nil {
		//Fix errors!
		err = render.Render(w, r, ErrInvalidRequest(err))
		if err != nil {
			log.Println(err)
		}
		return
	}
	nud, err := rt.hs.GenShortUrlHandle(r.Context(), urldata)
	if err != nil {
		err = render.Render(w, r, ErrRender(err))
		if err != nil {
			log.Println(err)
		}
	}
	err = render.Render(w, r, nud)
	if err != nil {
		log.Println(err)
	}
}

func (rt *ChiRouter) GetData(w http.ResponseWriter, r *http.Request) {
	urldata := ApiUrlData{}
	if err := render.Bind(r, &urldata); err != nil {
		//Fix errors!
		err = render.Render(w, r, ErrInvalidRequest(err))
		if err != nil {
			log.Println(err)
		}
		return
	}
	nud, err := rt.hs.GetDataHandle(r.Context(), urldata)
	if err != nil {
		err = render.Render(w, r, ErrRender(err))
		if err != nil {
			log.Println(err)
		}
	}
	err = render.Render(w, r, nud)
	if err != nil {
		log.Println(err)
	}
}
