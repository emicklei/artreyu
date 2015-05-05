package main

import (
	"log"
	"path/filepath"

	"github.com/emicklei/artreyu/model"
	"github.com/spf13/cobra"
)

type fetchCmd struct {
	*cobra.Command
}

func newFetchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fetch [destination]",
		Short: "download an artifact from the repository",
		Long: `destination can be a directory or regular file.
Parent directories will be created if absent.`,
	}
	fetch := new(fetchCmd)
	fetch.Command = cmd
	fetch.Command.Run = fetch.doFetch
	return fetch.Command
}

func (f *fetchCmd) doFetch(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		log.Fatalf("missing destination")
	}
	destination := args[len(args)-1]

	a, err := model.LoadArtifact("artreyu.yaml")
	if err != nil {
		log.Fatalf("unable to load artifact descriptor:%v", err)
	}

	// check if destination is directory
	var regular string = destination
	if model.IsDirectory(destination) {
		regular = filepath.Join(destination, a.StorageBase())
	}

	err = mainRepo().Fetch(a, regular)
	if err != nil {
		log.Fatalf("unable to download artifact:%v", err)
	}
}
