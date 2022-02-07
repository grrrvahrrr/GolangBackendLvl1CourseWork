package http

//Package responsible for primary sevice usage - allow user to input shortURL into browser and be redirected to fullURL

type HttpRequesterGetter interface {
	//Get request when a user inputs short url into browser
	GetHttpRequest(string) (shortURL string, err error)
}

type HttpRequesterSender interface {
	//Sends short URL to databse.Reader to get full URL
	SendRequestToDB(shortURL string) (string, error)
}

type Responser interface {
	//Gets response from database.Reader ReadURL function for fullURL
	GetResponse(fullURL string) (string, error)
	//Uses response - redirects user from shortURL to fullURL
	UseResponse(fullURL string) error
}
