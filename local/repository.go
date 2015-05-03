package local

import (
	"fmt"
	"path/filepath"

	"github.com/emicklei/artreyu/model"
)

type Repository struct {
	config model.Config
}

func NewRepository(config model.Config) Repository {
	return Repository{config}
}

// copy the file (source) to a local directory and name indicated by the artifact
func (r Repository) Store(a model.Artifact, source string) error {
	dest := a.StorageLocation(r.config.OSname)
	return model.Cp(filepath.Join(r.config.URL, dest), source)
}

func (r Repository) Fetch(a model.Artifact, destination string) error {
	src := a.StorageLocation(r.config.OSname)
	return model.Cp(destination, filepath.Join(r.config.URL, src))
}

func (r Repository) Assemble(a Assembly, source string) error {
	if len(a.Parts) == 0 {
		return fmt.Errorf("assemble has not parts listed")
	}
	return nil
}
