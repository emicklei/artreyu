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
func (a Artifact) StorageLocation(osname string, isAnyOS bool) string {
	osdir := osname
	if isAnyOS {
		osdir = AnyOSDirectoryName
	}
	return filepath.Join(strings.Replace(a.Group, ".", "/", -1), a.Name, a.Version, osdir, a.StorageBase())
}

// IsSnapshot returns true if the version has the substring "SNAPSHOT".
func (a Artifact) IsSnapshot() bool {
	return strings.Index(a.Version, "SNAPSHOT") != -1
}

// Validate will inspect all required fields.
func (a Artifact) Validate() error {
	if len(a.Group) == 0 {
		return fmt.Errorf("group of descriptor [%#v] cannot be empty", a)
	}
	if strings.Contains(a.Group, "/") {
		return fmt.Errorf("group of descriptor cannot be have path separator [%s]", a.Group)
	}
	if len(a.Name) == 0 {
		return fmt.Errorf("artifact (name) of descriptor [%#v] cannot be empty", a)
	}
	if strings.Contains(a.Name, "/") {
		return fmt.Errorf("artifact (name) of descriptor cannot be have path separator [%s]", a.Name)
	}
	if len(a.Version) == 0 {
		return fmt.Errorf("version of descriptor [%#v] cannot be empty", a)
	}
	if len(a.Type) == 0 {
		return fmt.Errorf("type (extension) of descriptor [%#v] cannot be empty", a)
	}
	if strings.HasPrefix(a.Type, ".") {
		return fmt.Errorf("type (extension) of descriptor cannot have the dot prefix [%s]", a.Type)
	}
	return nil
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
	data, err := ioutil.ReadAll(f)
	if err != nil {
		return a, err
	}
	err = yaml.Unmarshal(data, &a)
	if err != nil {
		return a, err
	}
	return a, a.Validate()
}
