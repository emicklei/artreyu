package main

import (
	"path/filepath"

	"github.com/emicklei/artreyu/model"
	"github.com/spf13/cobra"
)

func newFetchCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "fetch [optional:destination]",
		Short: "download an artifact from the repository",
		Long: `destination can be a directory or regular file.
Parent directories will be created if absent.`,
		Run: doFetch,
	}
}

func doFetch(cmd *cobra.Command, args []string) {
	var destination = "."
	if len(args) > 0 {
		destination = args[0]
	}

	a, err := model.LoadArtifact(appSettings.ArtifactConfigLocation)
	if err != nil {
		model.Fatalf("unable to load artifact descriptor:%v", err)
	}

	// check if destination is directory
	var regular string = destination
	if model.IsDirectory(destination) {
		regular = filepath.Join(destination, a.StorageBase())
	}

	err = mainRepo().Fetch(a, regular)
	if err != nil {
		model.Fatalf("unable to download artifact:%v", err)
	}
}
