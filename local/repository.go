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

// copy the file (source) to a local directory and name indicated by the artifact
func (r Repository) Store(a model.Artifact, source string) error {
	dest := a.StorageLocation(r.osName(a.AnyOS))
	return model.Cp(filepath.Join(r.config.URL, dest), source)
}

func (r Repository) Fetch(a model.Artifact, destination string) error {
	src := a.StorageLocation(r.osName(a.AnyOS))
	return model.Cp(destination, filepath.Join(r.config.URL, src))
}

func (r Repository) Assemble(a model.Assembly, source string) error {
	if len(a.Parts) == 0 {
		return fmt.Errorf("assemble has not parts listed")
	}
	return nil
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
