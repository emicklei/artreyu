package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/emicklei/typhoon/model"
	"github.com/emicklei/typhoon/nexus"
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

	cfg, err := model.LoadConfig(filepath.Join(os.Getenv("HOME"), ".typhoon"))
	if err != nil {
		log.Fatalf("unable to load config from ~/.typhoon:%v", err)
	}
	a, err := model.LoadAssembly("typhoon.yaml")
	if err != nil {
		log.Fatalf("unable to load assembly descriptor:%v", err)
	}

	r := nexus.NewRepository(cfg)
	err = r.Assemble(a, destination)
	if err != nil {
		log.Fatalf("unable to assemble from artifacts:%v", err)
	}
}
