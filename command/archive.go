package command

import (
	"github.com/emicklei/artreyu/model"
	"github.com/spf13/cobra"
)

type Archive struct {
	artifact   model.Artifact
	repository model.Repository
	source     string
}

func (a Archive) Perform() {
	if len(a.source) == 0 {
		model.Fatalf("missing source")
	}
	err := a.repository.Store(a.artifact, a.source)
	if err != nil {
		model.Fatalf("unable to upload artifact:%v", err)
	}
}

func NewCommandForArchive() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "archive [artifact]",
		Short: "upload an artifact to the repository",
		Run:   nil,
	}
	return cmd
}

func NewArchiveCommand(af ArtifactFunc, rf RepositoryFunc) *cobra.Command {
	cmd := NewCommandForArchive()
	cmd.Run = func(cmd *cobra.Command, args []string) {
		source := args[0]
		archive := Archive{
			artifact:   af(),
			repository: rf(),
			source:     source,
		}
		archive.Perform()
	}
	return cmd
}
