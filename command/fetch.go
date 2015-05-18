package command

import (
	"path/filepath"

	"github.com/emicklei/artreyu/model"
	"github.com/spf13/cobra"
)

type Fetch struct {
	Artifact    model.Artifact
	Repository  model.Repository
	Destination string
	AutoExtract bool
}

func (f Fetch) Perform() {
	// check if destination is directory
	var regular string = f.Destination
	if model.IsDirectory(f.Destination) {
		regular = filepath.Join(f.Destination, f.Artifact.StorageBase())
	}

	err := f.Repository.Fetch(f.Artifact, regular)
	if err != nil {
		model.Fatalf("unable to download artifact:%v", err)
	}

	if f.AutoExtract && model.IsTargz(regular) {
		if err := model.Untargz(regular, filepath.Dir(regular)); err != nil {
			model.Fatalf("unable to extract artifact:%v", err)
			return
		}
		if err := model.FileRemove(regular); err != nil {
			model.Fatalf("unable to remove compressed artifact:%v", err)
			return
		}
	}
}

var AutoExtract bool

func NewCommandForFetch() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fetch [optional:destination]",
		Short: "download an artifact from the repository",
		Long: `destination can be a directory or regular file.
Parent directories will be created if absent.`,
	}
	cmd.Flags().BoolVarP(&AutoExtract, "extract", "x", false, "extract the content of the compressed artifact")
	return cmd
}

func NewFetchCommand(af ArtifactFunc, rf RepositoryFunc) *cobra.Command {
	cmd := NewCommandForFetch()
	cmd.Run = func(cmd *cobra.Command, args []string) {
		var destination = "."
		if len(args) > 0 {
			destination = args[0]
		}
		fetch := Fetch{
			Artifact:    af(),
			Repository:  rf(),
			Destination: destination,
			AutoExtract: AutoExtract,
		}
		fetch.Perform()
	}
	return cmd
}
