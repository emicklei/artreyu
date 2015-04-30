package main

import (
	"log"
	"os"

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

func (a *archiveCmd) doArchive(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		log.Fatalf("missing artifact")
	}
	source := args[len(args)-1]
	log.Println(source)
}
