package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/emicklei/artreyu/model"
	"github.com/spf13/cobra"
)

var VERSION string = "dev"
var BUILDDATE string = "now"
var appConfig model.Config
var osOverride string
var RootCmd = &cobra.Command{
	Use:   "artreyu",
	Short: "artreyu a is an artifact assembly tool",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func main() {
	log.Printf("artreyu - artifact assembly tool (version:%s, build:%s)\n", VERSION, BUILDDATE)
	config, err := model.LoadConfig(filepath.Join(os.Getenv("HOME"), ".artreyu"))
	if err != nil {
		log.Fatalf("unable to load config from ~/.artreyu:%v", err)
	}
	// share config
	appConfig = config

	RootCmd.PersistentFlags().StringVar(&osOverride, "os", "", "overwrite if assembling for different OS")
	RootCmd.AddCommand(newArchiveCmd())
	RootCmd.AddCommand(newFetchCmd())
	RootCmd.AddCommand(newAssembleCmd())
	RootCmd.Execute()
}

func OSName() string {
	if len(osOverride) > 0 {
		return osOverride
	}
	return appConfig.OSname
}
