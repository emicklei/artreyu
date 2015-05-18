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

// Validate will inspect all required fields and that of its parts.
// Fail on the first error detected.
func (a Assembly) Validate() error {
	if err := a.Artifact.Validate(); err != nil {
		return err
	}
	for _, each := range a.Parts {
		if err := each.Validate(); err != nil {
			return err
		}
	}
	return nil
}
