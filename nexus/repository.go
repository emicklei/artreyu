package nexus

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/emicklei/typhoon/model"
)

var UserRepository = NewRepository(
	os.Getenv("NEXUS_USER"),
	os.Getenv("NEXUS_PASSWORD"),
	os.Getenv("NEXUS_REPOS"))

type Repository struct {
	user, password, repositoriesUrl string
}

func NewRepository(user, password, repositoriesUrl string) Repository {
	return Repository{
		user:            user,
		password:        password,
		repositoriesUrl: repositoriesUrl,
	}
}

func (r Repository) Store(a model.Artifact, source string) error {
	repo := "releases"
	if a.IsSnapshot() {
		repo = "snapshots"
	}
	destination := fmt.Sprintf("%s/%s/%s", r.repositoriesUrl, repo, a.StorageLocation())
	log.Printf("uploading %s to %s\n", source, destination)
	cmd := exec.Command(
		"curl",
		"-u",
		fmt.Sprintf("%s:%s", r.user, r.password),
		"--upload-file",
		source,
		destination)
	return cmd.Run()
}
