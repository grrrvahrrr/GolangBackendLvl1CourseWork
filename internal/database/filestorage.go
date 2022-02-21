package database

import (
	"CourseWork/internal/dbbackend"
	"CourseWork/internal/entities"
	"context"
	"encoding/json"
	"io"
	"log"
	"os"

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
func (fd *FullDataFile) WriteURL(ctx context.Context, url entities.UrlData) error {
	fd.URLData = url
	err := fd.enc.Encode(fd.URLData)
	if err != nil {
		//Log it
		log.Println("err")
		return err
	}
	return nil
}

func (fd *FullDataFile) WriteData(ctx context.Context, id uuid.UUID, data map[string]string) error {
	for {
		if err := fd.dec.Decode(&fd.URLData); err != nil {
			if err == io.EOF {
				log.Println("URL not found")
				return nil
			}
			return err
		}
		if id == fd.URLData.Id {

			fd.URLData.Data = data
			err := fd.enc.Encode(fd.URLData)
			if err != nil {
				//Log it
				return err
			}
			return nil
		}
	}
	// input, err := ioutil.ReadFile(fd.filename)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// lines := strings.Split(string(input), "\n")
	// for i, line := range lines {
	// 	if strings.Contains(line, url.Id.String()) {
	// 		lines[i] = lines[len(lines)-1]
	// 		lines[len(lines)-1] = ""
	// 		lines = lines[:len(lines)-1]
	// 	}
	// }
	// output := strings.Join(lines, "\n")
	// err = ioutil.WriteFile(fd.filename, []byte(output), 0644)
	// if err != nil {
	// 	log.Fatalln(err)
	// }

	// url.Data = data
	// err = fd.enc.Encode(url)
	// if err != nil {
	// 	//Log it
	// 	return err
	// }
	// return nil
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

		if url.ShortURL == fd.URLData.ShortURL {
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
