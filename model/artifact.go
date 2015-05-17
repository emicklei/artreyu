package model

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"gopkg.in/yaml.v2"
)

type Artifact struct {
	// Descriptor api version. 1 is the default.
	Api int `json:"api" yaml:"api"`

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

// StorageBase returns the file name to which the artifact is stored.
func (a Artifact) StorageBase() string {
	return fmt.Sprintf("%s-%s.%s", a.Name, a.Version, a.Type)
}

// StorageLocation returns the relative resource path to store the artifact.
func (a Artifact) StorageLocation(osname string) string {
	return filepath.Join(strings.Replace(a.Group, ".", "/", -1), a.Name, a.Version, osname, a.StorageBase())
}

// IsSnapshot returns true if the version has the substring "SNAPSHOT".
func (a Artifact) IsSnapshot() bool {
	return strings.Index(a.Version, "SNAPSHOT") != -1
}

// Verify will inspect all required fields. Exit if one does.
func (a Artifact) Verify() {
	if len(a.Group) == 0 {
		Fatalf("group of artifact [%#v] cannot be empty", a)
	}
	if len(a.Name) == 0 {
		Fatalf("name of artifact [%#v] cannot be empty", a)
	}
	if len(a.Version) == 0 {
		Fatalf("version of artifact [%#v] cannot be empty", a)
	}
	if len(a.Type) == 0 {
		Fatalf("type (extension) of artifact [%#v] cannot be empty", a)
	}
}

func (a Artifact) PluginParameters() (params []string) {
	return append(params,
		"--group="+a.Group,
		"--artifact="+a.Name,
		"--version="+a.Version,
		"--type="+a.Type,
		"--anyos="+strconv.FormatBool(a.AnyOS))
}

// LoadArtifact parses an Artifact by reading the src file.
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
