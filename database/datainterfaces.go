package database

import "context"

type FullDataDB struct {
	FullURL  string
	ShortURL string
	Data     string //to be remade into slice of strings of a map[string]string
}

//Add constructors so we could accept data as a struct into functions
//For now methods are skeletons of what they are

//Writing to a file/db what we got from Backend after validating URLs and Updating Data
func (fd *FullDataDB) WriteURL(ctx context.Context, fullURL string, shortURL string) error {
	return nil
}

func (fd *FullDataDB) WriteData(ctx context.Context, shortURL string, data string) error {
	return nil
}

//Reads info from file/DB
type Reader interface {
	ReadURL(ShortURL string) (FullURL string, err error)
	ReadData(ShortURL string) (Data string, err error)
}

func (fd *FullDataDB) ReadURL(ShortURL string) (FullURL string, err error) {
	return FullURL, nil
}

func (fd *FullDataDB) ReadData(ShortURL string) (Data string, err error) {
	return Data, nil
}

//Delete record if neccessary
func (fd *FullDataDB) Delete(FullURL string, ShortURL string) error {
	return nil
}

//Some kind of Sorting
func Sort(filename string) error {
	return nil
}
