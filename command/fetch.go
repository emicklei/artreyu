package command

import (
	"path/filepath"

	"github.com/emicklei/artreyu/model"
	"github.com/emicklei/artreyu/transport"
	"github.com/spf13/cobra"
)

// Fetch represents the actual action that can be performed on an Artifact using a Repository.
type Fetch struct {
	Artifact    model.Artifact
	Repository  model.Repository
	Destination string
	AutoExtract bool
	ExitOnError bool
}

// Perform will fetch the artifact from the repository
func (f Fetch) Perform() bool {
	// check if destination is directory
	var regular string = f.Destination
	if transport.IsDirectory(f.Destination) {
		regular = filepath.Join(f.Destination, f.Artifact.StorageBase())
	}

	err := f.Repository.Fetch(f.Artifact, regular)
	if err != nil {
		if f.ExitOnError {
			model.Fatalf("fetch failed: %v", err)
		}
		return false
	}

	if f.AutoExtract && transport.IsTargz(regular) {
		if err := transport.Untargz(regular, filepath.Dir(regular)); err != nil {
			if f.ExitOnError {
				model.Fatalf("fetch failed, unable to extract artifact: %v", err)
			}
			return false
		}
		if err := transport.FileRemove(regular); err != nil {
			if f.ExitOnError {
				model.Fatalf("fetch failed, unable to remove compressed artifact: %v", err)
			}
			return false
		}
	}
	return true
}

// AutoExtract is to capture the x flag value
var AutoExtract bool

// NewCommandForFetch returns a new Command for the fetch action, without the Run function.
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

// NewFetchCommand returns a new Command for the fetch action, with a Run function using the
// Artifact and Repository providing functions.
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
			ExitOnError: true,
		}
		fetch.Perform()
	}
	return cmd
}
