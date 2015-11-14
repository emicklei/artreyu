package main

import (
	"strconv"

	"github.com/emicklei/artreyu/command"
	"github.com/emicklei/artreyu/local"
	"github.com/emicklei/artreyu/model"
	"github.com/spf13/cobra"
)

func newFetch(applicationSettings *model.Settings) *cobra.Command {
	fetch := command.NewCommandForFetch()
	fetch.Run = func(cmd *cobra.Command, args []string) {
		artifact, err := model.LoadArtifact(applicationSettings.ArtifactConfigLocation)
		if err != nil {
			model.Fatalf("fetch failed, unable to load artifact: %v", err)
		}
		repoName := applicationSettings.TargetRepository
		var destination = "."
		if len(args) > 0 {
			destination = args[0]
		}
		// versions may be in local store
		// snapshots are in local store if target is set to local
		fetched := false
		if !artifact.IsSnapshot() || "local" == repoName {
			fetch := command.Fetch{
				Artifact:    artifact,
				Repository:  local.NewRepository(model.RepositoryConfigNamed(applicationSettings, "local"), applicationSettings.OS),
				Destination: destination,
				AutoExtract: command.AutoExtract,
				ExitOnError: false,
			}
			if fetch.Perform() {
				model.Printf("... fetched artifact from local cache")
				fetched = fetch.Perform()
			}
		}
		// done if target is set to local or local fetch of version was ok
		if "local" == repoName || fetched {
			return
		}

		// extend args with fetch specific flags
		extendedArgs := append(args, "--extract="+strconv.FormatBool(command.AutoExtract))

		// not local
		if err := command.RunPluginWithArtifact("artreyu-"+repoName, "fetch", artifact, *applicationSettings, extendedArgs); err != nil {
			model.Fatalf("fetch failed, could not run plugin:  %v", err)
		} else {
			// remote fetching succeeded, store a copy of a version in local
			if !artifact.IsSnapshot() {
				archive := command.Archive{
					Artifact:    artifact,
					Repository:  local.NewRepository(model.RepositoryConfigNamed(applicationSettings, "local"), applicationSettings.OS),
					Source:      destination,
					ExitOnError: false,
				}
				if archive.Perform() {
					model.Printf("... stored copy in local cache")
				}
			}
		}
	}

	return fetch
}
