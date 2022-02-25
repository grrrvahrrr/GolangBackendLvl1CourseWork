package openapichi

import (
	"CourseWork/internal/apichi"
	"encoding/json"
	"html/template"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

type OpenApiChi struct {
	*chi.Mux
	hs *apichi.Handlers
}

type PageVars struct {
	ShortURL string
	Data     string
}

func NewOpenApiRouter(hs *apichi.Handlers) *OpenApiChi {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	ret := &OpenApiChi{
		hs: hs,
	}

	r.Mount("/", Handler(ret))
	swg, err := GetSwagger()
	if err != nil {
		log.Fatal("swagger fail")
	}

	r.Get("/swagger.json", func(w http.ResponseWriter, r *http.Request) {
		enc := json.NewEncoder(w)
		_ = enc.Encode(swg)
	})

	ret.Mux = r

	return ret
}

type UrlData apichi.ApiUrlData

func (UrlData) Bind(r *http.Request) error {
	return nil
}

func (UrlData) Render(w http.ResponseWriter, r *http.Request) error {
	return nil
}

// (POST /getData)
func (rt *OpenApiChi) GetData(w http.ResponseWriter, r *http.Request) {
	urldata := UrlData{}
	if err := render.Bind(r, &urldata); err != nil {
		//Fix errors!
		err = render.Render(w, r, apichi.ErrInvalidRequest(err))
		if err != nil {
			log.Println(err)
		}
		return
	}
	nud, err := rt.hs.GetDataHandle(r.Context(), apichi.ApiUrlData(urldata))
	if err != nil {
		err = render.Render(w, r, apichi.ErrRender(err))
		if err != nil {
			log.Println(err)
		}
	}

	var s string

	for k, v := range nud.Data {
		s += k + v + "\n"
	}

	DataURLVars := PageVars{
		Data: s,
	}

	t, err := template.ParseFiles("../internal/frontend/getData.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, DataURLVars)
	if err != nil {
		log.Print("template executing error: ", err)
	}

	err = render.Render(w, r, nud)
	if err != nil {
		log.Println(err)
	}

}

// (POST /shortenURL)
func (rt *OpenApiChi) GenShortURL(w http.ResponseWriter, r *http.Request) {
	urldata := UrlData{}
	if err := render.Bind(r, &urldata); err != nil {
		//Fix errors!
		err = render.Render(w, r, apichi.ErrInvalidRequest(err))
		if err != nil {
			log.Println(err)
		}
		return
	}
	nud, err := rt.hs.GenShortUrlHandle(r.Context(), apichi.ApiUrlData(urldata))
	if err != nil {
		err = render.Render(w, r, apichi.ErrRender(err))
		if err != nil {
			log.Println(err)
		}
	}

	shortenURLVars := PageVars{
		ShortURL: nud.ShortURL,
	}

	t, err := template.ParseFiles("../internal/frontend/shortenURL.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, shortenURLVars)
	if err != nil {
		log.Print("template executing error: ", err)
	}

	err = render.Render(w, r, nud)
	if err != nil {
		log.Println(err)
	}

}

// (GET /su/{shortURL})
func (rt *OpenApiChi) Redirect(w http.ResponseWriter, r *http.Request, shortURL string) {

	if shortURL == "" {
		err := render.Render(w, r, apichi.ErrInvalidRequest(http.ErrNotSupported))
		if err != nil {
			log.Println(err)
		}
		return
	}

	nud, err := rt.hs.RedirectionHandle(r.Context(), shortURL)
	if err != nil {
		err = render.Render(w, r, apichi.ErrRender(err))
		if err != nil {
			log.Println(err)
		}
		return
	}

	http.Redirect(w, r, nud.FullURL, http.StatusMovedPermanently)

	err = render.Render(w, r, nud)
	if err != nil {
		log.Println(err)
	}
}

// (GET /home)
func (rt *OpenApiChi) GetUserURLs(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../internal/frontend/homepage.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}

	data := PageVars{}
	err = t.Execute(w, data)
	if err != nil {
		log.Print("template execute error: ", err)
	}
}
