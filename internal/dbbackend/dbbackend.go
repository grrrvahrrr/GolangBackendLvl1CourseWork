package dbbackend

import (
	"CourseWork/internal/entities"
	"context"
	"fmt"

	"github.com/google/uuid"
)

//Package for all interactions with a database
//Sould only import urldata.go

//Port to use in data storage
type DataStore interface {
	WriteURL(ctx context.Context, url entities.UrlData) (*uuid.UUID, error)
	WriteData(ctx context.Context, id uuid.UUID) (*entities.UrlData, error)
	ReadURL(ctx context.Context, id uuid.UUID) (*entities.UrlData, error)
	Delete(ctx context.Context, id uuid.UUID) error
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
	//Check if we need to create new id here
	url.Id = uuid.New()
	id, err := ds.dstore.WriteURL(ctx, url)
	if err != nil {
		//Log it
		return nil, fmt.Errorf("write url error: %w", err)
	}
	url.Id = *id
	return &url, nil
}

func (ds *DataStorage) WriteData(ctx context.Context, id uuid.UUID) (*entities.UrlData, error) {
	u, err := ds.dstore.WriteData(ctx, id)
	if err != nil {
		//Log it
		return nil, fmt.Errorf("write data error: %w", err)
	}
	return u, nil
}

//Read data from data storage
func (ds *DataStorage) ReadURL(ctx context.Context, id uuid.UUID) (*entities.UrlData, error) {
	u, err := ds.dstore.ReadURL(ctx, id)
	if err != nil {
		//Log it
		return nil, fmt.Errorf("write data error: %w", err)
	}
	return u, nil
}

//Delete data from data storage
func (ds *DataStorage) Delete(ctx context.Context, id uuid.UUID) (*entities.UrlData, error) {
	u, err := ds.dstore.ReadURL(ctx, id)
	if err != nil {
		//Log it
		return nil, fmt.Errorf("read url error: %w", err)
	}
	return u, ds.dstore.Delete(ctx, id)
}
