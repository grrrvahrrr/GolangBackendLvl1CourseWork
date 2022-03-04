package process

//This package is goint to hold business processes and import only urldata.go

import (
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
func UpdateNumOfUses(data string) (string, error) {
	iv, err := strconv.Atoi(data)
	if err != nil {
		return "", err
	}
	iv++
	data = strconv.Itoa(iv)

	return data, nil
}
