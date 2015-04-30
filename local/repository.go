package local

import (
	"path/filepath"

	"github.com/emicklei/typhoon/model"
)

type Repository struct {
	basePath string
}

func NewRepository(path string) Repository {
	return Repository{basePath: path}
}

// copy the file (source) to a local directory and name indicated by the artifact
func (r Repository) Store(a model.Artifact, source string) error {
	dest := a.StorageLocation()
	return model.Cp(filepath.Join(r.basePath, dest), source)
}

func (r Repository) Fetch(a model.Artifact, destination string) error {
	src := a.StorageLocation()
	return model.Cp(destination, filepath.Join(r.basePath, src))
}
