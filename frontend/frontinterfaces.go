package frontend

import "github.com/google/uuid"

//This should all probably be in JS, but for now it is in GO

//Responsible for printing text to the end-user
type FullData struct {
	Id       uuid.UUID         `json:"id"`
	FullURL  string            `json:"fullurl"`
	ShortURL string            `json:"shorturl"`
	Data     map[string]string `json:"data"`
}

func (fd FullData) PrintURL(fullData FullData) error {
	return nil
}

func (fd FullData) PrintData(fullData FullData) error {
	return nil
}

//Sends data to backend part
func (fd FullData) SendFullURL(fullData FullData) error {
	return nil
}

func (fd FullData) SendShortURL(fullData FullData) error {
	return nil
}

//Gets data from backend part, values go to FullData struct
//Instead of straight up values it is going to accept a recieved struct made from an unmarshaled JSON
//TODO: ADD Unmarshal JSON func

func NewFullData(Id uuid.UUID, FullURL string, ShortURL string, Data map[string]string) *FullData {
	return &FullData{
		Id:       Id,
		FullURL:  FullURL,
		ShortURL: ShortURL,
		Data:     Data,
	}
}

//Gets input from user, values go to Sender
func GetInputFullURL() (fullURL string, err error) {
	return fullURL, nil
}

func GetInputShortUrl() (shortURL string, err error) {
	return shortURL, nil
}
