package main

import (
	"path/filepath"

	"github.com/emicklei/artreyu/model"
	"github.com/spf13/cobra"
)

var autoExtract bool

func newFetchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fetch [optional:destination]",
		Short: "download an artifact from the repository",
		Long: `destination can be a directory or regular file.
Parent directories will be created if absent.`,
		Run: doFetch,
	}
	cmd.Flags().BoolVarP(&autoExtract, "extract", "x", false, "extract the content of the compressed artifact")
	return cmd
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

	if autoExtract && a.Type == "tgz" {
		if err := model.Untargz(regular, filepath.Dir(regular)); err != nil {
			model.Fatalf("unable to extract artifact:%v", err)
			return
		}
		if err := model.FileRemove(regular); err != nil {
			model.Fatalf("unable to remove compressed artifact:%v", err)
			return
		}
	}
}
