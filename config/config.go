package config

import (
	"github.com/joho/godotenv"
)

type Config struct {
	//Add config fields here

}

func (c *Config) loadConfigFile(file string) error {
	err := godotenv.Load(file)
	if err != nil {
		return err
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
