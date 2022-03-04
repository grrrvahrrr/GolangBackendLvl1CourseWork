package config

import (
	"os"
	"strconv"

	log "github.com/sirupsen/logrus"

	"github.com/joho/godotenv"
)

type Config struct {
	ReadTimeout       int
	WriteTimeout      int
	ReadHeaderTimeout int
}

func (c *Config) loadConfigFile(file string) error {
	var err error
	err = godotenv.Load(file)
	if err != nil {
		return err
	}
	c.ReadTimeout, err = strconv.Atoi(os.Getenv("READTIMEOUT"))
	if err != nil {
		log.Error(err)
	}
	c.WriteTimeout, err = strconv.Atoi(os.Getenv("WRITETIMEOUT"))
	if err != nil {
		log.Error(err)
	}
	c.ReadHeaderTimeout, err = strconv.Atoi(os.Getenv("READHEADERTIMEOUT"))
	if err != nil {
		log.Error(err)
	}

	return nil
}

func LoadConfig(file string) (Config, error) {
	var c Config
	err := c.loadConfigFile(file)
	if err != nil {
		return c, err
	}
	return c, nil
}
