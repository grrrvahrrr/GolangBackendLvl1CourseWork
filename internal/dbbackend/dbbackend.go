package dbbackend

import "context"

//Package for all interactions with a database

//Port to use in data storage
type DataStore interface {
	WriteURL(ctx context.Context, fullURL string, shortURL string) error
	WriteData(ctx context.Context, shortURL string, data string) error
	ReadURL(ctx context.Context, ShortURL string) (FullURL string, err error)
	ReadData(ctx context.Context, ShortURL string) (Data string, err error)
	Delete(ctx context.Context, FullURL string, ShortURL string) error
}

type DataStorage struct {
	ds DataStore
}

func NewDataStorage(ds DataStore) *DataStorage {
	return &DataStorage{
		ds: ds,
	}
}

//Write data to data storage
func (ds *DataStorage) WriteURL(ctx context.Context, fullURL string, shortURL string) error {
	return nil
}

func (ds *DataStorage) WriteData(ctx context.Context, shortURL string, data string) error {
	return nil
}

//Read data from data storage
func (ds *DataStorage) ReadURL(ctx context.Context, ShortURL string) (FullURL string, err error) {
	return FullURL, nil
}

func (ds *DataStorage) ReadData(ctx context.Context, ShortURL string) (Data string, err error) {
	return Data, nil
}

//Delete data from data storage
func (ds *DataStorage) Delete(ctx context.Context, FullURL string, ShortURL string) error {
	return nil
}
