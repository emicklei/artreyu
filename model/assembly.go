package model

import (
	"io/ioutil"
	"os"

	"gopkg.in/yaml.v2"
)

type Assembly struct {
	Artifact `yaml:",inline"`
	Parts    []Artifact `json:"parts" yaml:"parts"`
}

// LoadAssembly parses an Assembly by reading the src file.
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

// Verify will inspect all required fields and that of its parts. Exit if one does.
func (a Assembly) Verify() {
	a.Artifact.Verify()
	for _, each := range a.Parts {
		each.Verify()
	}
}
