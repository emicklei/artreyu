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

func main() {
	config, err := model.LoadConfig(filepath.Join(os.Getenv("HOME"), ".artreyu"))
	if err != nil {
		log.Fatalf("unable to load config from ~/.artreyu:%v", err)
	}
	// share config
	appConfig = config

	RootCmd.AddCommand(newArchiveCmd())
	RootCmd.AddCommand(newFetchCmd())
	RootCmd.AddCommand(newAssembleCmd())
	RootCmd.Execute()
}

var RootCmd = &cobra.Command{
	Use:   "artreyu",
	Short: "artreyu a is a tool for artifact management",
	Run:   func(cmd *cobra.Command, args []string) {},
}
