package model

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Config struct {
	OSname       string
	Repositories []RepositoryConfig
}

func (c Config) Named(name string) RepositoryConfig {
	for _, each := range c.Repositories {
		if each.Name == name {
			return each
		}
	}
	panic("no such repository in config:" + name)
}

type RepositoryConfig struct {
	Name     string
	URL      string
	Path     string
	User     string
	Password string
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
