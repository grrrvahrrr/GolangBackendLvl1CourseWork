package process

//This package is goint to hold business processes and import only urldata.go

import (
	"log"
	"math/rand"
	"net/url"
	"strconv"
)

//Generates random short strings
func GenerateRandomString() string {
	var symbols = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, 10)
	for i := range b {
		b[i] = symbols[rand.Intn(len(symbols))]
	}
	return string(b)
}

//Validate full URL that we got is a URL
func ValidateURL(fullURL string) error {
	_, err := url.ParseRequestURI(fullURL)
	return err
}

//Update neccessary data to send it to DB
//UpdateNumOfUses here, in future different functions for new elements of data
func UpdateNumOfUses(data map[string]string) (map[string]string, error) {
	v, ok := data["NumOfUses"]
	if !ok {
		data["NumOfUses"] = "1"
	} else {
		iv, err := strconv.Atoi(v)
		if err != nil {
			//Log it with logrus
			log.Println(err)
		}
		iv++
		data["NumOfUses"] = strconv.Itoa(iv)
	}
	return data, nil
}
