package entities

type UrlData struct {
	FullURL  string `json:"fullurl"`
	ShortURL string `json:"shorturl"`
	AdminURL string `json:"adminurl"`
	Data     string `json:"data"`
	IP       string `json:"ip"`
	IPData   string `json:"ipdata"`
}
