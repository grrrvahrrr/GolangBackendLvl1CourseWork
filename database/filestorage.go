package database

import "context"

type FullDataFile struct {
	FullURL  string
	ShortURL string
	Data     string //to be remade into slice of strings of a map[string]string
}

//Add constructors so we could accept data as a struct into functions
//For now methods are skeletons of what they are

//Writing to a file/db what we got from Backend after validating URLs and Updating Data
func (fd *FullDataFile) WriteURL(ctx context.Context, fullURL string, shortURL string) error {
	return nil
}

func (fd *FullDataFile) WriteData(ctx context.Context, shortURL string, data string) error {
	return nil
}

//Reads info from file/DB
func (fd *FullDataFile) ReadURL(ctx context.Context, ShortURL string) (FullURL string, err error) {
	return FullURL, nil
}

func (fd *FullDataFile) ReadData(ctx context.Context, ShortURL string) (Data string, err error) {
	return Data, nil
}

//Delete record if neccessary
func (fd *FullDataFile) Delete(ctx context.Context, FullURL string, ShortURL string) error {
	return nil
}

//Some kind of Sorting
func Sort(filename string) error {
	return nil
}
