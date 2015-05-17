package main

import (
	"github.com/emicklei/artreyu/command"
	"github.com/emicklei/artreyu/local"
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
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	ApplicationSettings = command.NewSettingsBoundToFlags(RootCmd)

	archive := command.NewCommandForArchive()
	archive.Run = func(cmd *cobra.Command, args []string) {
		artifact, err := model.LoadArtifact(ApplicationSettings.ArtifactConfigLocation)
		if err != nil {
			model.Fatalf("load artifact failed, archive aborted: %v", err)
		}
		target := ApplicationSettings.TargetRepository
		if "local" == target {
			archive := command.Archive{
				Artifact:   artifact,
				Repository: local.NewRepository(model.RepositoryConfigNamed(ApplicationSettings, "local"), ApplicationSettings.OS),
				Source:     args[0],
			}
			archive.Perform()
			return
		}
		// not local
		if err := command.RunPluginWithArtifact("artreyu-"+target, "archive", artifact, *ApplicationSettings, args); err != nil {
			model.Fatalf("archive failed, %v", err)
		}
	}
	RootCmd.AddCommand(archive)

	fetch := command.NewCommandForFetch()
	fetch.Run = func(cmd *cobra.Command, args []string) {
		artifact, err := model.LoadArtifact(ApplicationSettings.ArtifactConfigLocation)
		if err != nil {
			model.Fatalf("load artifact failed, fetch aborted: %v", err)
		}
		target := ApplicationSettings.TargetRepository
		if "local" == target {
			var destination = "."
			if len(args) > 0 {
				destination = args[0]
			}
			fetch := command.Fetch{
				Artifact:    artifact,
				Repository:  local.NewRepository(model.RepositoryConfigNamed(ApplicationSettings, "local"), ApplicationSettings.OS),
				Destination: destination,
				AutoExtract: command.AutoExtract,
			}
			fetch.Perform()
			return
		}
		// not local
		if err := command.RunPluginWithArtifact("artreyu-"+target, "fetch", artifact, *ApplicationSettings, args); err != nil {
			model.Fatalf("fetch failed, %v", err)
		}
	}
	RootCmd.AddCommand(fetch)
	RootCmd.AddCommand(newAssembleCmd())
}
