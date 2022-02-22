package database

import (
	"CourseWork/internal/dbbackend"
	"CourseWork/internal/entities"
	"context"
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"os"
	"regexp"
	"strings"

	"github.com/google/uuid"
)

var _ dbbackend.DataStore = &FullDataFile{}

type FullDataFile struct {
	URLData  entities.UrlData
	filename string
	file     *os.File
	dec      *json.Decoder
	enc      *json.Encoder
}

func NewFullDataFile(filename string) (*FullDataFile, error) {
	var err error
	fd := &FullDataFile{
		filename: filename,
	}
	fd.file, err = os.OpenFile(filename, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalln(err)
	}
	fd.enc = json.NewEncoder(fd.file)
	fd.dec = json.NewDecoder(fd.file)
	return fd, nil
}

//Writing to a file/db what we got from Backend after validating URLs and Updating Data
func (fd *FullDataFile) WriteURL(ctx context.Context, url entities.UrlData) (*entities.UrlData, error) {
	fd.URLData = url
	err := fd.enc.Encode(fd.URLData)
	if err != nil {
		//Log it
		log.Println("err")
		return nil, err
	}
	return &entities.UrlData{
		Id:       fd.URLData.Id,
		FullURL:  fd.URLData.FullURL,
		ShortURL: fd.URLData.ShortURL,
		Data:     fd.URLData.Data,
	}, nil
}

func (fd *FullDataFile) WriteData(ctx context.Context, id uuid.UUID, data map[string]string) (*entities.UrlData, error) {
	for {
		if err := fd.dec.Decode(&fd.URLData); err != nil {
			if err == io.EOF {
				log.Println("URL not found")
				return nil, nil
			}
			return nil, err
		}
		if id == fd.URLData.Id {
			fd.URLData.Data = data
			err := fd.enc.Encode(fd.URLData)
			if err != nil {
				//Log it
				return nil, err
			}
			//Process file
			err = fd.Sort(id)
			if err != nil {
				return nil, err
			}

			return &entities.UrlData{
				Id:       fd.URLData.Id,
				FullURL:  fd.URLData.FullURL,
				ShortURL: fd.URLData.ShortURL,
				Data:     fd.URLData.Data,
			}, nil
		}
	}
}

//Reads info from file/DB
func (fd *FullDataFile) ReadURL(ctx context.Context, url entities.UrlData) (*entities.UrlData, error) {
	for {
		if err := fd.dec.Decode(&fd.URLData); err != nil {
			if err == io.EOF {
				log.Println("URL not found")
				return nil, nil
			}
			return nil, err
		}

		if url.FullURL == fd.URLData.FullURL || url.ShortURL == fd.URLData.ShortURL {
			return &entities.UrlData{
				Id:       fd.URLData.Id,
				FullURL:  fd.URLData.FullURL,
				ShortURL: fd.URLData.ShortURL,
				Data:     fd.URLData.Data,
			}, nil
		}
	}
}

func (fd *FullDataFile) Close() {
	if fd.file != nil {
		fd.file.Close()
	}
}

func (fd *FullDataFile) Sort(id uuid.UUID) error {
	input, err := ioutil.ReadFile(fd.filename)
	if err != nil {
		log.Fatalln(err)
	}

	lines := strings.Split(string(input), "\n")

	for i, line := range lines {
		if strings.Contains(line, id.String()) {
			lines[i] = ""
			break
		}
	}
	output := strings.Join(lines, "\n")
	regex, err := regexp.Compile("\n\n")
	if err != nil {
		return err
	}
	output = regex.ReplaceAllString(output, "\n")

	err = ioutil.WriteFile(fd.filename, []byte(output), 0644)
	if err != nil {
		return err
	}
	return nil
}
