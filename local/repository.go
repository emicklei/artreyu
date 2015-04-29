package local

import "github.com/emicklei/typhoon/model"

type Repository struct {
	basePath string
}

func NewRepository(path string) Repository {
	return Repository{basePath: path}
}

func (r Repository) Fetch(a model.Artifact, destination string) error {
	return nil
}

// copy the file (source) to a local directory and name indicated by the artifact
func (r Repository) Store(a model.Artifact, source string) error {

}
