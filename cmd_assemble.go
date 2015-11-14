package main

import (
	"os"
	"path/filepath"

	"github.com/emicklei/artreyu/command"
	"github.com/emicklei/artreyu/local"
	"github.com/emicklei/artreyu/model"
	"github.com/emicklei/artreyu/transport"
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

	a, err := model.LoadAssembly(applicationSettings.ArtifactConfigLocation)
	if err != nil {
		model.Fatalf("unable to load assembly descriptor:%v", err)
	}

	if len(a.Parts) == 0 {
		model.Fatalf("assemble has no parts listed")
		return
	}

	// Download artifacts and decompress each
	for _, each := range a.Parts {
		targetFilename := filepath.Join(destination, each.StorageBase())
		repoName := applicationSettings.TargetRepository
		if "local" == repoName {
			if err := localRepository(applicationSettings).Fetch(each, targetFilename); err != nil {
				model.Fatalf("fetching artifact failed, aborted because:%v", err)
				return
			}
		} else {
			fetched := false
			// if version then try fetch from local
			if !a.IsSnapshot() {
				if err := localRepository(applicationSettings).Fetch(each, targetFilename); err == nil {
					model.Printf("copied artifact from local cache")
					fetched = true
				}
			}
			// snapshot or not local
			if !fetched {
				if err := command.RunPluginWithArtifact("artreyu-"+repoName, "fetch", each, *applicationSettings, args); err != nil {
					model.Fatalf("fetching artifact failed, aborted because:%v", err)
					return
				}
			}
			// if version then put in local
			if !a.IsSnapshot() {
				if err := localRepository(applicationSettings).Store(each, targetFilename); err != nil {
					model.Printf("unable to cache fetched artifact version")
				}
			}
		}
		if transport.IsZip(targetFilename) {
			if err := transport.Unzip(targetFilename, destination); err != nil {
				model.Fatalf("zip decompress failed, aborted because:%v", err)
				return
			}
			if err := transport.FileRemove(targetFilename); err != nil {
				model.Fatalf("remove failed, aborted because:%v", err)
				return
			}
		} else if transport.IsTargz(targetFilename) {
			if err := transport.Untargz(targetFilename, destination); err != nil {
				model.Fatalf("tar extract failed, aborted because:%v", err)
				return
			}
			if err := transport.FileRemove(targetFilename); err != nil {
				model.Fatalf("remove failed, aborted because:%v", err)
				return
			}
		}
	}
	// Compress into new artifact
	location := filepath.Join(destination, a.StorageBase())
	if transport.IsTargz("." + a.Type) {
		if err := transport.Targz(destination, location); err != nil {
			model.Fatalf("tar compress failed, aborted because:%v", err)
			return
		}
	} else if transport.IsZip("." + a.Type) {
		if err := transport.Zip(destination, location); err != nil {
			model.Fatalf("zip compress failed, aborted because:%v", err)
			return
		}
	}

	// Archive new artifact
	target := applicationSettings.TargetRepository
	if "local" == target {
		if err := localRepository(applicationSettings).Store(a.Artifact, location); err != nil {
			model.Fatalf("archiving new artifact failed, aborted because:%v", err)
			return
		}
	} else {
		if err := command.RunPluginWithArtifact("artreyu-"+target, "archive", a.Artifact, *applicationSettings, []string{location}); err != nil {
			model.Fatalf("archiving new artifact failed, aborted because:%v", err)
			return
		}
	}
}

func localRepository(settings *model.Settings) model.Repository {
	return local.NewRepository(model.RepositoryConfigNamed(settings, "local"), settings.OS)
}
