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
	AnyOS bool `yaml:"any-os"`
	// if not empty use this value for the actual storage file
	OverrideStorageBase string
	// optionally specify the repository to archive/fetch this artifact
	RepositoryName string `yaml:"repository`
}

// StorageBase returns the file name to which the artifact is stored.
func (a Artifact) StorageBase() string {
	if len(a.OverrideStorageBase) > 0 {
		return a.OverrideStorageBase
	}
	if len(a.Version) == 0 {
		return fmt.Sprintf("%s.%s", a.Name, a.Type)
	}
	return fmt.Sprintf("%s-%s.%s", a.Name, a.Version, a.Type)
}

// UseStorageBase is used to override the default format of the actual file to archive/fetch.
func (a *Artifact) UseStorageBase(actualFilename string) {
	a.OverrideStorageBase = actualFilename
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
		return fmt.Errorf("empty group of descriptor [%#v]", a)
	}
	if strings.Contains(a.Group, "/") {
		return fmt.Errorf("path separator in group of descriptor [%s]", a.Group)
	}
	if len(a.Name) == 0 {
		return fmt.Errorf("empty artifact (name) in descriptor [%#v]", a)
	}
	if strings.HasSuffix(a.Name, a.Type) {
		return fmt.Errorf("unexpected extension (type) in descriptor (name) [%#v]", a)
	}
	if strings.Contains(a.Name, "/") {
		return fmt.Errorf("path separator in descriptor (name) [%s]", a.Name)
	}
	if len(a.Version) == 0 {
		return fmt.Errorf("empty version in descriptor [%#v]", a)
	}
	if len(a.Type) == 0 {
		return fmt.Errorf("empty type (extension) in descriptor [%#v]", a)
	}
	if strings.HasPrefix(a.Type, ".") {
		return fmt.Errorf("unexpected dot in type (extension) of descriptor [%s]", a.Type)
	}
	return nil
}

func (a Artifact) PluginParameters() (params []string) {
	return append(params,
		"--group="+a.Group,
		"--artifact="+a.Name,
		"--version="+a.Version,
		"--type="+a.Type,
		"--base="+a.OverrideStorageBase,
		"--any-os="+strconv.FormatBool(a.AnyOS))
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
