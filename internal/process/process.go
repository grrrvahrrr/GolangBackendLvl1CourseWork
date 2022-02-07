package process

//This package is goint to hold business processes and import only urldata.go

import "math/rand"

//Generates random short strings
func GenerateRandomString(n int) string {
	var symbols = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789")
	b := make([]rune, n)
	for i := range b {
		b[i] = symbols[rand.Intn(len(symbols))]
	}
	return string(b)
}

//Parse if full URL that we got is a URL
func ParseURL(fullURL string) bool {
	return true
}

//Update neccessary data to send it to DB
func UpdateData(shortURL string, data string) (newData string, err error) {
	return newData, nil
}
