package dbbackend

//Package for all interactions with a database

//Generates random short strings
type Generater interface {
	GenerateRandomStrings() (string, error)
}

//Sends user requests to database.Reader
type DbReadRequestSender interface {
	SendDbReadURLRequest(shortURL string) (string, error)
	SendDbReadDataRequest(shortURL string) (string, error)
}

//Sends user requests to database.Writer
type DbWriteRequestSender interface {
	SendDbWriteDataRequest(data string) (string, error)
	SendDbWriteURLRequest(fullURL string, shortURL string) (string, string, error)
}

//Get information from database.Reader to send to give to FrontRequestSender interface
type DbResponseGetter interface {
	GetDbURLResponse(shortURL string) (string, error)
	GetDbDataResponse(data string) (string, error)
}
