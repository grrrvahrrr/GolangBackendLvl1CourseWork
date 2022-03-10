package openapichi

import (
	"CourseWork/internal/apichi"
	"embed"
	"encoding/json"
	"errors"
	"html/template"
	"net"
	"net/http"

	log "github.com/sirupsen/logrus"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/render"
)

//go:embed pages
var tpls embed.FS

type OpenApiChi struct {
	*chi.Mux
	hs *apichi.Handlers
}

type PageVars struct {
	ShortURL string
	AdminURL string
	FullURL  string
	Data     string
	IPData   string
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

// (GET /getData/{adminURL})
func (rt *OpenApiChi) AdminRedirect(w http.ResponseWriter, r *http.Request, adminURL string) {
	urldata := UrlData{
		AdminURL: adminURL,
	}

	nud, ipdata, err := rt.hs.GetDataHandle(r.Context(), apichi.ApiUrlData(urldata))
	if err != nil {
		log.Error(errors.Unwrap(err))
		err = render.Render(w, r, apichi.ErrRender(errors.Unwrap(err)))
		if err != nil {
			log.Error(err)
		}
	}

	DataURLVars := PageVars{
		Data:     nud.Data,
		ShortURL: nud.ShortURL,
		IPData:   ipdata,
	}

	t, err := template.ParseFS(tpls, "pages/getData.html")
	if err != nil {
		log.Error("template parsing error: ", err)
	}
	err = t.Execute(w, DataURLVars)
	if err != nil {
		log.Error("template executing error: ", err)
	}

}

// (POST /shortenURL)
func (rt *OpenApiChi) GenShortURL(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Error("error parsing form:", err)
	}

	fullurl := r.Form.Get("fullurl")
	if fullurl == "" {
		log.Error("search query not found:", err)
	}

	urldata := UrlData{
		FullURL: fullurl,
	}

	nud, err := rt.hs.GenShortUrlHandle(r.Context(), apichi.ApiUrlData(urldata))
	if err != nil {
		log.Error(errors.Unwrap(err))
		err = render.Render(w, r, apichi.ErrRender(errors.Unwrap(err)))
		if err != nil {
			log.Error(err)
		}
	}

	shortenURLVars := PageVars{
		ShortURL: nud.ShortURL,
		AdminURL: nud.AdminURL,
		FullURL:  nud.FullURL,
	}

	t, err := template.ParseFS(tpls, "pages/shortenURL.html")
	if err != nil {
		log.Error("template parsing error: ", err)
	}
	err = t.Execute(w, shortenURLVars)
	if err != nil {
		log.Error("template executing error: ", err)
	}

}

// (GET /su/{shortURL})
func (rt *OpenApiChi) Redirect(w http.ResponseWriter, r *http.Request, shortURL string) {

	if shortURL == "" {
		err := render.Render(w, r, apichi.ErrInvalidRequest(http.ErrNotSupported))
		log.Error(err)
		if err != nil {
			log.Error(err)
		}
		return
	}

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Error(err)
	}

	nud, err := rt.hs.RedirectionHandle(r.Context(), shortURL, ip)
	if err != nil {
		log.Error(errors.Unwrap(err))
		err = render.Render(w, r, apichi.ErrRender(errors.Unwrap(err)))
		if err != nil {
			log.Error(err)
		}
		return
	}

	http.Redirect(w, r, nud.FullURL, http.StatusSeeOther)

}

// (GET /home)
func (rt *OpenApiChi) GetUserFullURL(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFS(tpls, "pages/homepage.html")
	if err != nil {
		log.Error("template parsing error: ", err)
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Error("template execute error: ", err)
	}
}
