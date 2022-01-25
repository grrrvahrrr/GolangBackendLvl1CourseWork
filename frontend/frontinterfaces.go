package frontend

//Responsible for printing text to the end-user
type Printer interface {
	PrintURL(fullURL string, shortURL string) error
	PrintData(data string) error
}

//Sends data to backend part
type Sender interface {
	SendFullURL(fullURL string) error
	SendShortURL(shortURL string) error
}

//Gets data from backend part, values go to Printer
type BackendGetter interface {
	GetURL(string) (shortURL string, err error)
	GetData(string) (data string, err error)
}

//Gets input from user, values go to Sender
type UserGetter interface {
	GetInputFullURL() (fullURL string, err error)
	GetInputShortUrl() (shortURL string, err error)
}
