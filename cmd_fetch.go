package main

import (
	"log"

	"github.com/emicklei/artreyu/model"
	"github.com/emicklei/artreyu/nexus"
	"github.com/spf13/cobra"
)

type fetchCmd struct {
	*cobra.Command
}

func newFetchCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "fetch [destination]",
		Short: "download an artifact from the repository to [destination]",
	}
	fetch := new(fetchCmd)
	fetch.Command = cmd
	fetch.Command.Run = fetch.doFetch
	return fetch.Command
}

func (f *fetchCmd) doFetch(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		log.Fatalf("missing destination")
	}
	destination := args[len(args)-1]

	a, err := model.LoadArtifact("artreyu.yaml")
	if err != nil {
		log.Fatalf("unable to load artifact descriptor:%v", err)
	}

	r := nexus.NewRepository(appConfig.Repositories[1], OSName())
	err = r.Fetch(a, destination)
	if err != nil {
		log.Fatalf("unable to download artifact:%v", err)
	}
}
