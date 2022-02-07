package frontend

//This should all probably be in JS, but for now it is in GO

//Responsible for printing text to the end-user
type FullData struct {
	FullUrl  string
	ShortURL string
	Data     string //probably will be changed to a slice of string or a map[string]string
}

func (fd FullData) PrintURL(fullURL string, shortURL string) error {
	return nil
}

func (fd FullData) PrintData(data string) error {
	return nil
}

//Sends data to backend part
func (fd FullData) SendFullURL(fullURL string) error {
	return nil
}

func (fd FullData) SendShortURL(shortURL string) error {
	return nil
}

//Gets data from backend part, values go to FullData struct
//Instead of straight up values it is going to accept a recieved struct made from an unmarshaled JSON
//TODO: ADD Unmarshal JSON func

func NewFullData(FullURL string, ShortURL string, Data string) *FullData {
	return &FullData{
		FullUrl:  FullURL,
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
