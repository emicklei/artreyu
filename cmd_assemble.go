package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/emicklei/artreyu/model"
	"github.com/spf13/cobra"
)

func newAssembleCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "assemble [optional:destination]",
		Short: "upload a new artifact by assembling fetched parts as specified in the descriptor",
		Run:   doAssemble,
	}
}

func doAssemble(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		log.Fatalf("missing destination")
	}
	var destination = os.TempDir()
	if len(args) > 0 {
		destination = args[0]
	}

	if err := os.MkdirAll(destination, os.ModePerm); err != nil {
		log.Fatalf("unable to create destination folder:%v", err)
	}

	a, err := model.LoadAssembly(appSettings.ArtifactConfigLocation)
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
	location := filepath.Join(destination, a.StorageBase())
	if "tgz" == a.Type {
		if err := model.Targz(destination, location); err != nil {
			log.Fatalf("targz failed, aborted because:%v", err)
			return
		}
	}
	// Archive new artifact
	if err := mainRepo().Store(a.Artifact, location); err != nil {
		log.Fatalf("archiving new artifact failed, aborted because:%v", err)
		return
	}
}
