package main

import (
	"os"

	"github.com/emicklei/artreyu/command"
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
		// TODO refactor this
		model.Verbose = applicationSettings.Verbose
		if applicationSettings.Verbose {
			dir, _ := os.Getwd()
			model.Printf("working directory = [%s]", dir)
		}
	}

	rootCmd.AddCommand(newArchive(applicationSettings))
	rootCmd.AddCommand(newFetch(applicationSettings))
	rootCmd.AddCommand(newAssembleCmd())
	rootCmd.AddCommand(newFormatCmd())
	rootCmd.AddCommand(newTreeCmd())
}
