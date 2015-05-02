package nexus

import (
	"fmt"
	"log"
	"os/exec"
	"path/filepath"

	"github.com/emicklei/typhoon/model"
)

type Repository struct {
	config model.Config
}

func NewRepository(config model.Config) Repository {
	return Repository{config}
}

func (r Repository) Store(a model.Artifact, source string) error {
	repo := "releases"
	if a.IsSnapshot() {
		repo = "snapshots"
	}
	destination := fmt.Sprintf("%s/%s/%s", r.config.URL, repo, a.StorageLocation(r.config.OSname))
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
	source := fmt.Sprintf("%s/%s/%s", r.config.URL, repo, a.StorageLocation(r.config.OSname))
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
