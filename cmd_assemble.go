package main

import (
	"os"
	"path/filepath"

	"github.com/emicklei/artreyu/command"
	"github.com/emicklei/artreyu/local"
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
		model.Fatalf("missing destination")
	}
	var destination = os.TempDir()
	if len(args) > 0 {
		destination = args[0]
	}

	if err := os.MkdirAll(destination, os.ModePerm); err != nil {
		model.Fatalf("unable to create destination folder:%v", err)
	}

	a, err := model.LoadAssembly(ApplicationSettings.ArtifactConfigLocation)
	if err != nil {
		model.Fatalf("unable to load assembly descriptor:%v", err)
	}

	if len(a.Parts) == 0 {
		model.Fatalf("assemble has no parts listed")
		return
	}

	// Download artifacts and decompress each
	for _, each := range a.Parts {
		where := filepath.Join(destination, each.StorageBase())
		target := ApplicationSettings.TargetRepository
		if "local" == target {
			if err := localRepository(ApplicationSettings).Fetch(a.Artifact, where); err != nil {
				model.Fatalf("fetching artifact failed, aborted because:%v", err)
				return
			}
		} else {
			if err := command.RunPluginWithArtifact("artreyu-"+target, "fetch", each, *ApplicationSettings, args); err != nil {
				model.Fatalf("fetching artifact failed, aborted because:%v", err)
				return
			}
		}
		// TODO .tar.gz, .zip, .gz
		if "tgz" == each.Type {
			if err := model.Untargz(where, destination); err != nil {
				model.Fatalf("tar extract failed, aborted because:%v", err)
				return
			}
			if err := model.FileRemove(where); err != nil {
				model.Fatalf("remove failed, aborted because:%v", err)
				return
			}
		}
	}
	// Compress into new artifact
	location := filepath.Join(destination, a.StorageBase())
	if "tgz" == a.Type {
		if err := model.Targz(destination, location); err != nil {
			model.Fatalf("tar compress failed, aborted because:%v", err)
			return
		}
	}
	// Archive new artifact
	target := ApplicationSettings.TargetRepository
	if "local" == target {
		if err := localRepository(ApplicationSettings).Store(a.Artifact, location); err != nil {
			model.Fatalf("archiving new artifact failed, aborted because:%v", err)
			return
		}
	} else {
		if err := command.RunPluginWithArtifact("artreyu-"+target, "archive", a.Artifact, *ApplicationSettings, args); err != nil {
			model.Fatalf("archiving new artifact failed, aborted because:%v", err)
			return
		}
	}
}

func localRepository(settings *model.Settings) model.Repository {
	return local.NewRepository(model.RepositoryConfigNamed(settings, "local"), settings.OS)
}
