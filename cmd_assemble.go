package main

import (
	"log"
	"os"

	"github.com/emicklei/artreyu/model"
	"github.com/emicklei/artreyu/nexus"
	"github.com/spf13/cobra"
)

type assembleCmd struct {
	*cobra.Command
	osname string
}

func newAssembleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "assemble [destination]",
		Short: "download parts of an artifact from the repository to [destination]",
	}
	assemble := new(assembleCmd)
	assemble.Command = cmd
	assemble.PersistentFlags().StringVar(&assemble.osname, "osname", "", "overwrite if assembling for different OS")
	assemble.Command.Run = assemble.doAssemble
	return assemble.Command
}

func (s *assembleCmd) doAssemble(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		log.Fatalf("missing destination")
	}
	destination := args[len(args)-1]

	if err := os.MkdirAll(destination, os.ModePerm); err != nil {
		log.Fatalf("unable to create destination folder:%v", err)
	}

	a, err := model.LoadAssembly("artreyu.yaml")
	if err != nil {
		log.Fatalf("unable to load assembly descriptor:%v", err)
	}

	r := nexus.NewRepository(appConfig.Repositories[1], appConfig.OSname)
	err = r.Assemble(a, destination)
	if err != nil {
		log.Fatalf("unable to assemble from artifacts:%v", err)
	}
}
