package nexus

import (
	"fmt"
	"log"
	"os/exec"
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

func (r Repository) osName(isAny bool) string {
	if isAny {
		return "any"
	}
	return r.osname
}

func (r Repository) Store(a model.Artifact, source string) error {
	repo := "releases"
	if a.IsSnapshot() {
		repo = "snapshots"
	}
	destination := r.config.URL + filepath.Join(r.config.Path, repo, a.StorageLocation(r.osName(a.AnyOS)))
	log.Printf("uploading %s to %s\n", source, destination)
	cmd := exec.Command(
		"curl",
		"-u",
		fmt.Sprintf("%s:%s", r.config.User, r.config.Password),
		"--upload-file",
		source,
		destination)
	data, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(string(data))
	}
	return err
}

func (r Repository) Fetch(a model.Artifact, destination string) error {
	repo := "releases"
	if a.IsSnapshot() {
		repo = "snapshots"
	}
	source := r.config.URL + filepath.Join(r.config.Path, repo, a.StorageLocation(r.osName(a.AnyOS)))
	log.Printf("downloading %s to %s\n", source, destination)
	cmd := exec.Command(
		"curl",
		"-u",
		fmt.Sprintf("%s:%s", r.config.User, r.config.Password),
		source,
		"-o",
		destination)
	data, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(string(data))
	}
	return err
}

func (r Repository) Assemble(a model.Assembly, destination string) error {
	if len(a.Parts) == 0 {
		return fmt.Errorf("assemble has not parts listed")
	}
	for _, each := range a.Parts {
		where := filepath.Join(destination, each.StorageBase())
		if err := r.Fetch(each, where); err != nil {
			return fmt.Errorf("aborted because:%v", err)
		}
		if "tgz" == each.Type {
			if err := model.Untargz(where, destination); err != nil {
				return fmt.Errorf("untargz failed, aborted because:%v", err)
			}
			if err := model.FileRemove(where); err != nil {
				return fmt.Errorf("remove failed, aborted because:%v", err)
			}
		}
		if "tgz" == a.Type {
			if err := model.Targz(destination, filepath.Join(destination, a.StorageBase())); err != nil {
				return fmt.Errorf("targz failed, aborted because:%v", err)
			}
		}
	}
	return nil
}

func (r Repository) Exists(a model.Artifact) bool {
	repo := "releases"
	if a.IsSnapshot() {
		repo = "snapshots"
	}
	source := r.config.URL + filepath.Join(r.config.Path, repo, a.StorageLocation(r.osName(a.AnyOS)))
	fmt.Println(source)
	// TODO
	return false
}
