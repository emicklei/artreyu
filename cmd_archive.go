package main

import (
	"github.com/emicklei/artreyu/command"
	"github.com/emicklei/artreyu/local"
	"github.com/emicklei/artreyu/model"
	"github.com/spf13/cobra"
)

func newArchive(applicationSettings *model.Settings) *cobra.Command {
	archive := command.NewCommandForArchive()
	archive.Run = func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			model.Fatalf("archive failed, missing source argument")
		}
		artifact, err := model.LoadArtifact(applicationSettings.ArtifactConfigLocation)
		if err != nil {
			model.Fatalf("archive failed, invalid artifact: %v", err)
		}
		descriptorArtifact := model.Artifact{
			Api:     artifact.Api,
			AnyOS:   artifact.AnyOS,
			Group:   artifact.Group,
			Name:    artifact.Name,
			Version: artifact.Version,
			Type:    artifact.Type,
		}
		descriptorArtifact.UseStorageBase("artreyu.yaml")
		repoName := applicationSettings.TargetRepository
		// put versions in local repo.
		// put snapshots in local store if local is target
		if !artifact.IsSnapshot() || "local" == repoName {
			archive := command.Archive{
				Artifact:    artifact,
				Repository:  local.NewRepository(model.RepositoryConfigNamed(applicationSettings, "local"), applicationSettings.OS),
				Source:      args[0],
				ExitOnError: false,
			}
			ok := archive.Perform()
			if ok {
				model.Printf("... stored artifact in local cache")
				// now store the descriptor
				descArchive := command.Archive{
					Artifact:    descriptorArtifact,
					Repository:  local.NewRepository(model.RepositoryConfigNamed(applicationSettings, "local"), applicationSettings.OS),
					Source:      applicationSettings.ArtifactConfigLocation,
					ExitOnError: false,
				}
				if descArchive.Perform() {
					model.Printf("... stored descriptor in local cache")
				}
			} else {
				model.Printf("[WARN] unable to store artifact in local cache")
			}
		}
		// done if local is target
		if "local" == repoName {
			return
		}
		targetRepo := model.RepositoryConfigNamed(applicationSettings, repoName)
		pluginName := targetRepo.Plugin

		// not local, no archive specific flags to add
		if err := command.RunPluginWithArtifact("artreyu-"+pluginName, "archive", artifact, *applicationSettings, args); err != nil {
			model.Fatalf("archive failed, could not run plugin: %v", err)
		} else {
			// now store the descriptor. replace the source in args
			args[0] = applicationSettings.ArtifactConfigLocation
			command.RunPluginWithArtifact("artreyu-"+pluginName, "archive", descriptorArtifact, *applicationSettings, args)
		}
	}
	return archive
}
