package model

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Servers map[string]ServerConfig
}
type ServerConfig struct {
	URL      string
	Path     string
	User     string
	Password string
	OSname   string
}

func LoadConfig(source string) (c Config, err error) {
	f, err := os.Open(source)
	if err != nil {
		return c, err
	}
	defer f.Close()

	data, _ := ioutil.ReadAll(f)
	err = yaml.Unmarshal(data, &c)
	if err != nil {
		return c, err
	}
	return c, nil
}
