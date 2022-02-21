package config

import (
	"log"
	"os"
	"strconv"

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
		//Log it
		log.Println(err)
	}
	c.WriteTimeout, err = strconv.Atoi(os.Getenv("WRITETIMEOUT"))
	if err != nil {
		//Log it
		log.Println(err)
	}
	c.ReadHeaderTimeout, err = strconv.Atoi(os.Getenv("READHEADERTIMEOUT"))
	if err != nil {
		//Log it
		log.Println(err)
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
