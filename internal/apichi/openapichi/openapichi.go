package openapichi

import (
	"CourseWork/internal/apichi"
	"encoding/json"
	"html/template"
	"log"
	"net"
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
		err = render.Render(w, r, apichi.ErrRender(err))
		if err != nil {
			log.Println(err)
		}
	}

	DataURLVars := PageVars{
		Data:     nud.Data,
		ShortURL: nud.ShortURL,
		IPData:   ipdata,
	}

	t, err := template.ParseFiles("../internal/frontend/getData.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, DataURLVars)
	if err != nil {
		log.Print("template executing error: ", err)
	}

}

// (POST /shortenURL)
func (rt *OpenApiChi) GenShortURL(w http.ResponseWriter, r *http.Request) {

	err := r.ParseForm()
	if err != nil {
		log.Println("error parsing form")
		return
	}

	fullurl := r.Form.Get("fullurl")
	if fullurl == "" {
		log.Println("Search query not found!")
		return
	}

	urldata := UrlData{
		FullURL: fullurl,
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
		AdminURL: nud.AdminURL,
		FullURL:  nud.FullURL,
	}

	t, err := template.ParseFiles("../internal/frontend/shortenURL.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}
	err = t.Execute(w, shortenURLVars)
	if err != nil {
		log.Print("template executing error: ", err)
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

	ip, _, err := net.SplitHostPort(r.RemoteAddr)
	if err != nil {
		log.Println(err)
		return
	}

	nud, err := rt.hs.RedirectionHandle(r.Context(), shortURL, ip)
	if err != nil {
		err = render.Render(w, r, apichi.ErrRender(err))
		if err != nil {
			log.Println(err)
		}
		return
	}

	http.Redirect(w, r, nud.FullURL, http.StatusSeeOther)

}

// (GET /home)
func (rt *OpenApiChi) GetUserFullURL(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("../internal/frontend/homepage.html")
	if err != nil {
		log.Print("template parsing error: ", err)
	}

	err = t.Execute(w, nil)
	if err != nil {
		log.Print("template execute error: ", err)
	}

}
