package local

import (
	"fmt"
	"path/filepath"

	"github.com/emicklei/artreyu/model"
	"github.com/emicklei/artreyu/transport"
)

// Repository implements github.com/emicklei/artreyu/model/Repository
type Repository struct {
	osname string
	config model.RepositoryConfig
}

// NewRepository returns a new local Repository using the configuration and current OS name.
func NewRepository(config model.RepositoryConfig, operatingSystemName string) Repository {
	return Repository{operatingSystemName, config}
}

func (r Repository) ID() string { return "local" }

// Store copies the file (source) to a local directory and name indicated by the artifact
func (r Repository) Store(a model.Artifact, source string) error {
	dest := a.StorageLocation(r.osname, a.AnyOS)
	return transport.Cp(filepath.Join(r.config.Path, dest), source)
}

// Fetch copies the file indicated by the artifact to a destination file/directory.
func (r Repository) Fetch(a model.Artifact, destination string) error {
	src := filepath.Join(r.config.Path, a.StorageLocation(r.osname, a.AnyOS))
	if !transport.Exists(src) {
		return fmt.Errorf("no such file [%s]", src)
	}
	return transport.Cp(destination, src)
}

// Exists returns true if the repository holds the artifact file.
func (r Repository) Exists(a model.Artifact) bool {
	return transport.Exists(a.StorageLocation(r.osname, a.AnyOS))
}
