package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/emicklei/artreyu/model"
	"github.com/spf13/cobra"
)

type assembleCmd struct {
	*cobra.Command
}

func newAssembleCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "assemble [destination]",
		Short: "create a new artifact [destination] by assembling parts from the descriptor",
	}
	assemble := new(assembleCmd)
	assemble.Command = cmd
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

	if len(a.Parts) == 0 {
		log.Fatalf("assemble has no parts listed")
		return
	}

	// Download artifacts and decompress each
	for _, each := range a.Parts {
		where := filepath.Join(destination, each.StorageBase())
		if err := mainRepo().Fetch(each, where); err != nil {
			log.Fatalf("aborted because:%v", err)
			return
		}
		if "tgz" == each.Type {
			if err := model.Untargz(where, destination); err != nil {
				log.Fatalf("untargz failed, aborted because:%v", err)
				return
			}
			if err := model.FileRemove(where); err != nil {
				log.Fatalf("remove failed, aborted because:%v", err)
				return
			}
		}
	}
	// Compress into new artifact
	if "tgz" == a.Type {
		if err := model.Targz(destination, filepath.Join(destination, a.StorageBase())); err != nil {
			log.Fatalf("targz failed, aborted because:%v", err)
			return
		}
	}
	// Archive new artifact
	if err := mainRepo().Store(a.Artifact, destination); err != nil {
		log.Fatalf("archiving new artifact failed, aborted because:%v", err)
		return
	}
}
