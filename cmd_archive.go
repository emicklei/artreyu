package main

import (
	"log"

	"github.com/emicklei/artreyu/model"
	"github.com/spf13/cobra"
)

type archiveCmd struct {
	*cobra.Command
	overwrite bool
}

func newArchiveCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "archive [artifact]",
		Short: "upload an artifact to the repository",
	}
	archiveCmd := new(archiveCmd)
	archiveCmd.Command = cmd
	archiveCmd.Command.Run = archiveCmd.doArchive
	return archiveCmd.Command
}

func (c *archiveCmd) doArchive(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		log.Fatalf("missing artifact")
	}
	source := args[len(args)-1]

	a, err := model.LoadArtifact("artreyu.yaml")
	if err != nil {
		log.Fatalf("unable to load artifact descriptor:%v", err)
	}

	err = mainRepo.Store(a, source)
	if err != nil {
		log.Fatalf("unable to upload artifact:%v", err)
	}
}
