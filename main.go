package main

import (
	"github.com/emicklei/artreyu/command"
	"github.com/emicklei/artreyu/model"
	"github.com/spf13/cobra"
)

var VERSION string = "dev"
var BUILDDATE string = "now"

var ApplicationSettings *model.Settings

var RootCmd *cobra.Command

func main() {
	model.Printf("artreyu - artifact assembly tool (build:%s, commit:%s)\n", BUILDDATE, VERSION)
	initRootCommand()
	RootCmd.Execute()
}

func initRootCommand() {
	RootCmd = &cobra.Command{
		Use:   "artreyu",
		Short: "archives, fetches and assembles build artifacts",
		Long: `A tool for handling versioned, platform dependent artifacts.
Its primary purpose is to create assembly artifacts from build artifacts archived in a (remote) repository.

See https://github.com/emicklei/artreyu for more details.

(c)2015 ernestmicklei.com, MIT license`,
		Run: func(cmd *cobra.Command, args []string) {},
	}
	ApplicationSettings = command.NewSettingsBoundToFlags(RootCmd)

	archive := command.NewCommandForArchive()
	archive.Run = func(cmd *cobra.Command, args []string) {
		command.RunPlugin("artreyu-nexus", "archive", *ApplicationSettings, args)
	}
	RootCmd.AddCommand(archive)

	fetch := command.NewCommandForFetch()
	fetch.Run = func(cmd *cobra.Command, args []string) {
		command.RunPlugin("artreyu-nexus", "fetch", *ApplicationSettings, args)
	}
	RootCmd.AddCommand(fetch)

	RootCmd.AddCommand(newAssembleCmd())
}
