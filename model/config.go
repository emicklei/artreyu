package model

import (
	"errors"
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

const AnyOSDirectoryName = "any-os"

var ErrNoSuchRepository = errors.New("no configuration for this repository name")

type Config struct {
	Api          int
	Repositories []RepositoryConfig

	verbose      bool
	osOverride   string
	fileLocation string
}

func (c Config) Named(name string) (RepositoryConfig, error) {
	for _, each := range c.Repositories {
		if each.Name == name {
			return each, nil
		}
	}
	return RepositoryConfig{}, ErrNoSuchRepository
}

type RepositoryConfig struct {
	Name      string
	URL       string
	Path      string
	User      string
	Password  string
	Snapshots bool
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

func RepositoryConfigNamed(settings *Settings, name string) RepositoryConfig {
	config, err := LoadConfig(settings.MainConfigLocation)
	if err != nil {
		Fatalf("unable to load config %v", err)
	}
	repoconfig, err := config.Named(name)
	if err != nil {
		Fatalf("no configuration for repository named [%s], %v", name, err)
	}
	return repoconfig
}
