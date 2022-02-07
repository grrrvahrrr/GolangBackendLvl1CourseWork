package logging

import (
	"os"

	log "github.com/sirupsen/logrus"
)

func LogErrors() {
	log.SetFormatter(&log.JSONFormatter{})
	f, err := os.OpenFile("logs/error.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Error opening file: %v", err)
	}
	defer f.Close()
	log.SetOutput(f)
}
