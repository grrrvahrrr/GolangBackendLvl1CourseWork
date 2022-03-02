package database

import (
	"CourseWork/internal/dbbackend"
	"CourseWork/internal/entities"
	"context"
	"log"

	"github.com/syndtr/goleveldb/leveldb"
)

var _ dbbackend.DataStore = &FullDataFile{}

type FullDataFile struct {
	URLData    entities.UrlData
	shorturldb *leveldb.DB
	adminurldb *leveldb.DB
	datadb     *leveldb.DB
}

func NewFullDataFile(shorturldbfn string, adminurldbfn string, datadbfn string) (*FullDataFile, error) {
	var err error
	shorturldb, err := leveldb.OpenFile(shorturldbfn, nil)
	if err != nil {
		log.Fatal(err)
	}
	adminurldb, err := leveldb.OpenFile(adminurldbfn, nil)
	if err != nil {
		log.Fatal(err)
	}
	datadb, err := leveldb.OpenFile(datadbfn, nil)
	if err != nil {
		log.Fatal(err)
	}

	fd := &FullDataFile{
		shorturldb: shorturldb,
		adminurldb: adminurldb,
		datadb:     datadb,
	}

	return fd, nil
}

//Writing to a file/db what we got from Backend after validating URLs and Updating Data
func (fd *FullDataFile) WriteURL(ctx context.Context, url entities.UrlData) (*entities.UrlData, error) {
	fd.URLData = url
	err := fd.shorturldb.Put([]byte(url.ShortURL), []byte(url.FullURL), nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = fd.adminurldb.Put([]byte(url.AdminURL), []byte(url.ShortURL), nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	fd.URLData.Data = "0"
	err = fd.datadb.Put([]byte(url.ShortURL), []byte(fd.URLData.Data), nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &entities.UrlData{
		FullURL:  fd.URLData.FullURL,
		ShortURL: fd.URLData.ShortURL,
		AdminURL: fd.URLData.AdminURL,
		Data:     fd.URLData.Data,
	}, nil
}

func (fd *FullDataFile) WriteData(ctx context.Context, url entities.UrlData) (*entities.UrlData, error) {
	err := fd.datadb.Put([]byte(url.ShortURL), []byte(url.Data), nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &url, nil
}

//Reads info from file/DB
func (fd *FullDataFile) ReadURL(ctx context.Context, url entities.UrlData) (*entities.UrlData, error) {

	if url.AdminURL != "" {
		fd.URLData.AdminURL = url.AdminURL
		data, err := fd.adminurldb.Get([]byte(fd.URLData.AdminURL), nil)
		fd.URLData.ShortURL = string(data)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	} else if url.ShortURL != "" {
		fd.URLData.ShortURL = url.ShortURL
	} else {
		//Change that with error handling!
		log.Println("Couldn't find URL in DB")
		return nil, nil
	}

	data, err := fd.shorturldb.Get([]byte(fd.URLData.ShortURL), nil)
	fd.URLData.FullURL = string(data)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	data, err = fd.datadb.Get([]byte(fd.URLData.ShortURL), nil)
	fd.URLData.Data = string(data)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return &entities.UrlData{
		FullURL:  fd.URLData.FullURL,
		ShortURL: fd.URLData.ShortURL,
		AdminURL: fd.URLData.AdminURL,
		Data:     fd.URLData.Data,
	}, nil

}

func (fd *FullDataFile) Close() {
	//Add check
	fd.shorturldb.Close()
	fd.adminurldb.Close()
	fd.datadb.Close()

}
