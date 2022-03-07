package database

import (
	"CourseWork/internal/dbbackend"
	"CourseWork/internal/entities"
	"context"
	"fmt"
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
		return nil, err
	}
	adminurldb, err := leveldb.OpenFile(adminurldbfn, nil)
	if err != nil {
		return nil, err
	}
	datadb, err := leveldb.OpenFile(datadbfn, nil)
	if err != nil {
		return nil, err
	}

	ipdb, err := leveldb.OpenFile(ipdbfn, nil)
	if err != nil {
		return nil, err
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
		return nil, fmt.Errorf("error writing to shorturldb : %w", err)
	}
	err = fd.adminurldb.Put([]byte(url.AdminURL), []byte(url.ShortURL), nil)
	if err != nil {
		return nil, fmt.Errorf("error writing to adminurldb : %w", err)
	}

	fd.URLData.Data = "0"
	err = fd.datadb.Put([]byte(url.ShortURL), []byte(fd.URLData.Data), nil)
	if err != nil {
		return nil, fmt.Errorf("error writing to datadb : %w", err)
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
		return nil, fmt.Errorf("error writing to datadb : %w", err)
	}

	d := strings.Join([]string{fd.URLData.ShortURL, fd.URLData.IP}, ":")

	err = fd.ipdb.Put([]byte(d), []byte(url.IPData), nil)
	if err != nil {
		return nil, fmt.Errorf("error writing to ipdb : %w", err)
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
			return nil, fmt.Errorf("error reading adminurldb : %w", err)
		}
	} else if url.ShortURL != "" {
		fd.URLData.ShortURL = url.ShortURL
		fd.URLData.IP = url.IP
	} else {
		return nil, fmt.Errorf("recieved empty struct, no key to find")
	}

	data, err := fd.shorturldb.Get([]byte(fd.URLData.ShortURL), nil)
	fd.URLData.FullURL = string(data)
	if err != nil {
		return nil, fmt.Errorf("error reading shorturldb : %w", err)
	}

	data, err = fd.datadb.Get([]byte(fd.URLData.ShortURL), nil)
	fd.URLData.Data = string(data)
	if err != nil {
		return nil, fmt.Errorf("error reading datadb : %w", err)
	}

	d := strings.Join([]string{fd.URLData.ShortURL, fd.URLData.IP}, ":")

	ipdata, err := fd.ipdb.Get([]byte(d), nil)
	if err != nil && err != leveldb.ErrNotFound {
		return nil, fmt.Errorf("error reading ipdb : %w", err)
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
		return "", fmt.Errorf("error iterrating over ipdb : %w", err)
	}
	return ipdata, nil
}

func (fd *FullDataFile) Close() {
	fd.shorturldb.Close()
	fd.adminurldb.Close()
	fd.datadb.Close()
	fd.ipdb.Close()
}
