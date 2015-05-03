package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/emicklei/artreyu/model"
	"github.com/emicklei/artreyu/nexus"
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
	archiveCmd.PersistentFlags().BoolVar(&archiveCmd.overwrite, "force", false, "force overwrite if version exists")
	archiveCmd.Command.Run = archiveCmd.doArchive
	return archiveCmd.Command
}

func (c *archiveCmd) doArchive(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		log.Fatalf("missing artifact")
	}
	source := args[len(args)-1]

	cfg, err := model.LoadConfig(filepath.Join(os.Getenv("HOME"), ".artreyu"))
	if err != nil {
		log.Fatalf("unable to load config from ~/.artreyu:%v", err)
	}
	a, err := model.LoadArtifact("artreyu.yaml")
	if err != nil {
		log.Fatalf("unable to load artifact descriptor:%v", err)
	}

	r := nexus.NewRepository(cfg.Servers["nexus"]) // TODO how to specify nexus
	err = r.Store(a, source)
	if err != nil {
		log.Fatalf("unable to upload artifact:%v", err)
	}
}
