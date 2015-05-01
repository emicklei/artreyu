package main

import (
	"log"
	"os"

	"github.com/emicklei/typhoon/model"
	"github.com/emicklei/typhoon/nexus"
	"github.com/spf13/cobra"
)

type archiveCmd struct {
	*artifactCmd
	overwrite bool
}

func newArchiveCmd() *cobra.Command {
	cmd := newArtifactCmd(&cobra.Command{
		Use:   "archive [artifact]",
		Short: "copy an artifact to the repository",
	})
	archiveCmd := new(archiveCmd)
	archiveCmd.artifactCmd = cmd
	archiveCmd.PersistentFlags().BoolVar(&archiveCmd.overwrite, "force", false, "force overwrite if version exists")
	cmd.Command.Run = archiveCmd.doArchive
	return cmd.Command
}

func getRepo() string {
	repo := os.Getenv("TYPHOON_REPO")
	if len(repo) == 0 {
		log.Fatal("missing TYPHOON_REPO environment setting")
	}
	return repo
}

func (c *archiveCmd) doArchive(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		log.Fatalf("missing artifact")
	}
	source := args[len(args)-1]

	a, err := model.LoadArtifact("typhoon.yaml")
	a.Uname = c.artifactCmd.uname
	if err != nil {
		log.Fatal("unable to load artifact:%v", err)
	}

	r := nexus.UserRepository
	err = r.Store(a, source)
	if err != nil {
		log.Fatal("unable to store artifact:%v", err)
	}
}
