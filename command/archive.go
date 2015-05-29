package command

import (
	"github.com/emicklei/artreyu/model"
	"github.com/spf13/cobra"
)

// Archive represents the actual action that can be performed on an Artifact using a Repository.
type Archive struct {
	Artifact    model.Artifact
	Repository  model.Repository
	Source      string
	ExitOnError bool
}

// Perform will store the artifact into the repository
func (a Archive) Perform() bool {
	if len(a.Source) == 0 {
		model.Fatalf("missing source")
	}
	err := a.Repository.Store(a.Artifact, a.Source)
	if err != nil {
		if a.ExitOnError {
			model.Fatalf("unable to upload artifact:%v", err)
		}
		return false
	}
	return true
}

// NewCommandForArchive returns a new Command for the archive action, without the Run function.
func NewCommandForArchive() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "archive [artifact]",
		Short: "upload an artifact to the repository",
		Run:   nil,
	}
	return cmd
}

// NewArchiveCommand returns a new Command for the archive action, with a Run function using the
// Artifact and Repository providing functions.
func NewArchiveCommand(af ArtifactFunc, rf RepositoryFunc) *cobra.Command {
	cmd := NewCommandForArchive()
	cmd.Run = func(cmd *cobra.Command, args []string) {
		archive := Archive{
			Artifact:   af(),
			Repository: rf(),
			Source:     args[0],
		}
		archive.Perform()
	}
	return cmd
}
