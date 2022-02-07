package uibackend

//Package for interactions with the front end part
type FrontRequestGetter interface {
	//Gets information from frontend.Sender
	GetFrontURLRequest(fullURL string) (string, error)
	GetFrontDataRequest(shortURL string) (string, error)
}

type FrontRequestSender interface {
	//Sends information to frontend.Getter
	SendFrontURLRequest(shortURL string) (string, error)
	SendFrontDataRequest(data string) (string, error)
}
