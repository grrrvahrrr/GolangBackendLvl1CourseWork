package database

import (
	"CourseWork/internal/dbbackend"
	"CourseWork/internal/entities"
	"context"
	"log"
	"strings"

	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/util"
)

var _ dbbackend.DataStore = &FullDataFile{}

type FullDataFile struct {
	URLData    entities.UrlData
	shorturldb *leveldb.DB
	adminurldb *leveldb.DB
	datadb     *leveldb.DB
	ipdb       *leveldb.DB
}

func NewFullDataFile(shorturldbfn string, adminurldbfn string, datadbfn string, ipdbfn string) (*FullDataFile, error) {
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

	ipdb, err := leveldb.OpenFile(ipdbfn, nil)
	if err != nil {
		log.Fatal(err)
	}

	fd := &FullDataFile{
		shorturldb: shorturldb,
		adminurldb: adminurldb,
		datadb:     datadb,
		ipdb:       ipdb,
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

	d := strings.Join([]string{fd.URLData.ShortURL, fd.URLData.IP}, ":")

	err = fd.ipdb.Put([]byte(d), []byte(url.IPData), nil)
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
		fd.URLData.IP = url.IP
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

	d := strings.Join([]string{fd.URLData.ShortURL, fd.URLData.IP}, ":")

	ipdata, err := fd.ipdb.Get([]byte(d), nil)
	if err != nil && err != leveldb.ErrNotFound {
		log.Println(err)
		return nil, err
	} else if err == leveldb.ErrNotFound {
		fd.URLData.IPData = "0"
	} else {
		fd.URLData.IPData = string(ipdata)
	}

	return &entities.UrlData{
		FullURL:  fd.URLData.FullURL,
		ShortURL: fd.URLData.ShortURL,
		AdminURL: fd.URLData.AdminURL,
		Data:     fd.URLData.Data,
		IP:       fd.URLData.IP,
		IPData:   fd.URLData.IPData,
	}, nil

}

func (fd *FullDataFile) GetIPData(ctx context.Context, url entities.UrlData) (string, error) {
	var ipdata string
	iter := fd.ipdb.NewIterator(util.BytesPrefix([]byte(fd.URLData.ShortURL)), nil)
	for iter.Next() {
		key := iter.Key()
		value := iter.Value()

		ipdata += "IP: " + strings.TrimLeft(string(key), fd.URLData.ShortURL) + " # Redirects: " + string(value) + "\n"
	}
	iter.Release()
	err := iter.Error()
	if err != nil {
		log.Println(err)
		return "", err
	}
	return ipdata, nil
}

func (fd *FullDataFile) Close() {
	//Add check
	fd.shorturldb.Close()
	fd.adminurldb.Close()
	fd.datadb.Close()

}
