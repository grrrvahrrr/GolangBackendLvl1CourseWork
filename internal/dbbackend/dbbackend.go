package dbbackend

import (
	"CourseWork/internal/entities"
	"CourseWork/internal/process"
	"context"
	"fmt"
	"log"
)

//Port to use in data storage
type DataStore interface {
	WriteURL(ctx context.Context, url entities.UrlData) (*entities.UrlData, error)
	WriteData(ctx context.Context, url entities.UrlData) (*entities.UrlData, error)
	ReadURL(ctx context.Context, url entities.UrlData) (*entities.UrlData, error)
}

type DataStorage struct {
	dstore DataStore
}

func NewDataStorage(dstore DataStore) *DataStorage {
	return &DataStorage{
		dstore: dstore,
	}
}

//Write data to data storage
func (ds *DataStorage) WriteURL(ctx context.Context, url entities.UrlData) (*entities.UrlData, error) {
	err := process.ValidateURL(url.FullURL)
	if err != nil {
		return nil, fmt.Errorf("validate url error: %w", err)
	}

	url.ShortURL = process.GenerateRandomString()
	url.AdminURL = process.GenerateRandomString()
	newurldata, err := ds.dstore.WriteURL(ctx, url)
	if err != nil {
		//Log it
		return nil, fmt.Errorf("write url error: %w", err)
	}

	return newurldata, nil
}

func (ds *DataStorage) WriteData(ctx context.Context, url entities.UrlData) (*entities.UrlData, error) {
	var err error
	u, err := ds.dstore.ReadURL(ctx, url)
	if err != nil {
		//Log it
		return nil, fmt.Errorf("read data error: %w", err)
	}

	//Process data -- with more functions can be separated into independent function
	u.Data, err = process.UpdateNumOfUses(u.Data)
	if err != nil {
		//Log it
		log.Println(err)
	}

	newurldata, err := ds.dstore.WriteData(ctx, *u)
	if err != nil {
		//Log it
		return nil, fmt.Errorf("write data error: %w", err)
	}

	return newurldata, nil
}

//Read data from data storage
func (ds *DataStorage) ReadURL(ctx context.Context, url entities.UrlData) (*entities.UrlData, error) {
	u, err := ds.dstore.ReadURL(ctx, url)
	if err != nil {
		//Log it
		return nil, fmt.Errorf("read data error: %w", err)
	}
	return u, nil
}
