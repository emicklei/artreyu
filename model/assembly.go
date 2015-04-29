package model

import (
	"io/ioutil"
	"os"

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
