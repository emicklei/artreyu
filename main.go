package main

import "github.com/spf13/cobra"

var VERSION string = "dev"
var BUILDDATE string = "now"

func main() {
	RootCmd.AddCommand(newArchiveCmd())
	RootCmd.AddCommand(newFetchCmd())
	RootCmd.AddCommand(newAssembleCmd())
	RootCmd.Execute()
}

var RootCmd = &cobra.Command{
	Use:   "typhoon",
	Short: "typhoon a is a tool for artifact management",
	Run:   func(cmd *cobra.Command, args []string) {},
}
