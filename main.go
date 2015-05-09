package main

import (
	"os"
	"path/filepath"
	"runtime"

	"github.com/emicklei/artreyu/local"
	"github.com/emicklei/artreyu/model"
	"github.com/emicklei/artreyu/nexus"
	"github.com/spf13/cobra"
)

var VERSION string = "dev"
var BUILDDATE string = "now"

var appSettings Settings

var repo model.Repository

var RootCmd *cobra.Command

type Settings struct {
	Verbose                bool
	OS                     string
	MainConfigLocation     string
	ArtifactConfigLocation string
}

func main() {
	initRootCommand()
	RootCmd.Execute()
}

func initRootCommand() {
	RootCmd = &cobra.Command{
		Use:   "artreyu",
		Short: "archives, fetches and assembles build artifacts",
		Long: `A tool for handling versioned, platform dependent artifacts.
Its primary purpose it to create assembly artifacts from build artifacts archived in a (remote) repository.

Currently, it supports a local (filesystem) and Sonatype Nexus repository.

See https://github.com/emicklei/artreyu for more details.

(c)2015 ernestmicklei.com, MIT license`,
		Run: func(cmd *cobra.Command, args []string) {},
	}

	RootCmd.PersistentFlags().StringVarP(&appSettings.MainConfigLocation,
		"config",
		"c",
		filepath.Join(os.Getenv("HOME"), ".artreyu"),
		"location of the artreyu repositories configuration")
	RootCmd.PersistentFlags().StringVarP(&appSettings.ArtifactConfigLocation,
		"descriptor",
		"d",
		"artreyu.yaml",
		"overwrite if the artifact descriptor has a different name or location")
	RootCmd.PersistentFlags().StringVarP(&appSettings.OS,
		"os",
		"o",
		runtime.GOOS,
		"overwrite if assembling for different OS")
	RootCmd.PersistentFlags().BoolVarP(&appSettings.Verbose,
		"verbose",
		"v",
		false,
		"set to true for more execution details")
	RootCmd.AddCommand(newArchiveCmd())
	RootCmd.AddCommand(newFetchCmd())
	RootCmd.AddCommand(newAssembleCmd())
}

func mainRepo() model.Repository {
	// lazy initialize
	if repo != nil {
		return repo
	}
	model.Printf("artreyu - artifact assembly tool (version:%s, build:%s)\n", VERSION, BUILDDATE)
	if appSettings.Verbose {
		model.Printf("loading configuration from [%s]\n", appSettings.MainConfigLocation)
	}
	config, err := model.LoadConfig(appSettings.MainConfigLocation)
	if err != nil {
		model.Fatalf("unable to load config %v", err)
	}
	localconfig, err := config.Named("local")
	if err != nil {
		if appSettings.Verbose {
			model.Printf("no local repository available\n")
		}
		nexusconfig, err := config.Named("nexus")
		if err != nil {
			if appSettings.Verbose {
				model.Printf("no nexus repository available\n")
			}
			model.Fatalf("no repository available")
		}
		return nexus.NewRepository(nexusconfig, appSettings.OS)
	}
	local := local.NewRepository(localconfig, appSettings.OS)
	nexusconfig, err := config.Named("nexus")
	if err != nil {
		if appSettings.Verbose {
			model.Printf("no nexus repository available, using local\n")
		}
		return local
	}
	nexus := nexus.NewRepository(nexusconfig, appSettings.OS)
	// share
	repo = model.NewCachingRepository(nexus, local)
	return repo
}
