package command

import (
	"github.com/emicklei/artreyu/model"
	"github.com/spf13/cobra"
)

type Archive struct {
	Artifact   model.Artifact
	Repository model.Repository
	Source     string
}

func (a Archive) Perform() {
	if len(a.Source) == 0 {
		model.Fatalf("missing source")
	}
	err := a.Repository.Store(a.Artifact, a.Source)
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
		archive := Archive{
			Artifact:   af(),
			Repository: rf(),
			Source:     args[0],
		}
		archive.Perform()
	}
	return cmd
}
