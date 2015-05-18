package main

import (
	"os"

	"github.com/emicklei/artreyu/command"
	"github.com/emicklei/artreyu/local"
	"github.com/emicklei/artreyu/model"
	"github.com/spf13/cobra"
)

var (
	VERSION   string = "dev"
	BUILDDATE string = "now"

	applicationSettings *model.Settings
	rootCmd             *cobra.Command
)

func main() {
	model.Printf("artreyu - artifact assembly tool (build:%s, commit:%s)\n", BUILDDATE, VERSION)
	initRootCommand()
	rootCmd.Execute()
}

func initRootCommand() {
	rootCmd = &cobra.Command{
		Use:   "artreyu",
		Short: "archives, fetches and assembles build artifacts",
		Long: `A tool for handling versioned, platform dependent artifacts.
Its primary purpose is to create assembly artifacts from build artifacts archived in a (remote) repository.

See https://github.com/emicklei/artreyu for more details.

(c)2015 http://ernestmicklei.com, MIT license`,
		Run: func(cmd *cobra.Command, args []string) {
			cmd.Help()
		},
	}
	applicationSettings = command.NewSettingsBoundToFlags(rootCmd)
	rootCmd.PersistentPreRun = func(cmd *cobra.Command, args []string) {
		if applicationSettings.Verbose {
			dir, _ := os.Getwd()
			model.Printf("working directory = [%s]", dir)
		}
	}

	archive := command.NewCommandForArchive()
	archive.Run = func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			model.Fatalf("archive failed, missing source argument")
		}
		artifact, err := model.LoadArtifact(applicationSettings.ArtifactConfigLocation)
		if err != nil {
			model.Fatalf("archive failed, could not load artifact: %v", err)
		}
		target := applicationSettings.TargetRepository
		if "local" == target {
			archive := command.Archive{
				Artifact:   artifact,
				Repository: local.NewRepository(model.RepositoryConfigNamed(applicationSettings, "local"), applicationSettings.OS),
				Source:     args[0],
			}
			archive.Perform()
			return
		}
		// not local
		if err := command.RunPluginWithArtifact("artreyu-"+target, "archive", artifact, *applicationSettings, args); err != nil {
			model.Fatalf("archive failed, could not run plugin: %v", err)
		}
	}
	rootCmd.AddCommand(archive)

	fetch := command.NewCommandForFetch()
	fetch.Run = func(cmd *cobra.Command, args []string) {
		artifact, err := model.LoadArtifact(applicationSettings.ArtifactConfigLocation)
		if err != nil {
			model.Fatalf("fetch failed, unable to load artifact: %v", err)
		}
		target := applicationSettings.TargetRepository
		if "local" == target {
			var destination = "."
			if len(args) > 0 {
				destination = args[0]
			}
			fetch := command.Fetch{
				Artifact:    artifact,
				Repository:  local.NewRepository(model.RepositoryConfigNamed(applicationSettings, "local"), applicationSettings.OS),
				Destination: destination,
				AutoExtract: command.AutoExtract,
			}
			fetch.Perform()
			return
		}
		// not local
		if err := command.RunPluginWithArtifact("artreyu-"+target, "fetch", artifact, *applicationSettings, args); err != nil {
			model.Fatalf("fetch failed, could not run plugin:  %v", err)
		}
	}
	rootCmd.AddCommand(fetch)
	rootCmd.AddCommand(newAssembleCmd())
}
