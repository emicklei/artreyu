package local

import (
	"fmt"
	"path/filepath"

	"github.com/emicklei/artreyu/model"
)

type Repository struct {
	osname string
	config model.RepositoryConfig
}

func NewRepository(config model.RepositoryConfig, operatingSystemName string) Repository {
	return Repository{operatingSystemName, config}
}

func (r Repository) ID() string { return "local" }

// copy the file (source) to a local directory and name indicated by the artifact
func (r Repository) Store(a model.Artifact, source string) error {
	dest := a.StorageLocation(r.osName(a.AnyOS))
	return model.Cp(filepath.Join(r.config.Path, dest), source)
}

func (r Repository) Fetch(a model.Artifact, destination string) error {
	src := filepath.Join(r.config.Path, a.StorageLocation(r.osName(a.AnyOS)))
	if !model.Exists(src) {
		return fmt.Errorf("nu such file [%s]", src)
	}
	return model.Cp(destination, src)
}

func (r Repository) Exists(a model.Artifact) bool {
	src := a.StorageLocation(r.osName(a.AnyOS))
	return model.Exists(src)
}

func (r Repository) osName(isAny bool) string {
	if isAny {
		return "any"
	}
	return r.osname
}
