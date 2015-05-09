package main

import (
	"github.com/emicklei/artreyu/model"
	"github.com/spf13/cobra"
)

func newArchiveCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "archive [artifact]",
		Short: "upload an artifact to the repository",
		Run:   doArchive,
	}
}

func doArchive(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		model.Fatalf("missing artifact")
	}
	source := args[len(args)-1]

	a, err := model.LoadArtifact(appSettings.ArtifactConfigLocation)
	if err != nil {
		model.Fatalf("unable to load artifact descriptor:%v\n", err)
	}

	err = mainRepo().Store(a, source)
	if err != nil {
		model.Fatalf("unable to upload artifact:%v", err)
	}
}
