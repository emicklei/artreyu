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
	Parts    []Artifact `json:"parts" yaml:"parts"`
}

type Artifact struct {
	// Descriptor api version. 1 is the default.
	Api int `json:"artreyu-api" yaml:"artreyu-api"`

	// Name of the group of artifacts. Cannot contain "/" or whitespace characters.
	Group string
	// Name of the artifact. Cannot contain "/" or whitespace characters.
	Name string `yaml:"artifact"`
	// Use semantic versioning or include the "SNAPSHOT" keyword for non-fixed versions.
	Version string
	// Represents the file extension.
	Type string
	// If true then use "any" for the operating system name when archiving/fetching.
	AnyOS bool `yaml:"anyos"`
}

func (a Artifact) StorageBase() string {
	return fmt.Sprintf("%s-%s.%s", a.Name, a.Version, a.Type)
}

func (a Artifact) StorageLocation(osname string) string {
	return filepath.Join(strings.Replace(a.Group, ".", "/", -1), a.Name, a.Version, osname, a.StorageBase())
}

func (a Artifact) IsSnapshot() bool {
	return strings.Index(a.Version, "SNAPSHOT") != -1
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
