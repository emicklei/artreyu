package main

import (
	"log"

	"github.com/spf13/cobra"
)

func newVersionCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "version",
		Short: "show build info",
		Run: func(cmd *cobra.Command, args []string) {
			log.Println("_/^\\_")
			log.Println(" | | typhoon - artifact tool [commit=", VERSION, "build=", BUILDDATE, "]")
			log.Println("-\\_/-")
		},
	}
}
