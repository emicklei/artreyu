package model

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v2"
)

type Assembly struct {
	Artifact `yaml:",inline"`
	Api      int        `json:"typhoon-api" yaml:"typhoon-api"`
	Parts    []Artifact `json:"parts" yaml:"parts"`
}

type Artifact struct {
	Group     string
	Name      string `yaml:"artifact"`
	Version   string
	Extension string
}

func (a Artifact) StorageLocation() string {
	return filepath.Join(strings.Replace(a.Group, ".", "/", -1), a.Name, a.Version, fmt.Sprintf("%s-%s.%s", a.Name, a.Version, a.Extension))
}

func LoadAssembly(src string) (a Assembly, e error) {
	f, err := os.Open(src)
	if err != nil {
		return a, err
	}
	defer f.Close()

	data, _ := ioutil.ReadAll(f)
	err = yaml.Unmarshal(data, &a)
	if err != nil {
		return a, err
	}
	return a, nil
}

func LoadArtifact(src string) (a Artifact, e error) {
	f, err := os.Open(src)
	if err != nil {
		return a, err
	}
	defer f.Close()

	data, _ := ioutil.ReadAll(f)
	err = yaml.Unmarshal(data, &a)
	if err != nil {
		return a, err
	}
	return a, nil
}
