package main

import "github.com/spf13/cobra"

type artifactCmd struct {
	*cobra.Command
	artifact string
	group    string
	version  string
}

func newArtifactCmd(cobraCmd *cobra.Command) *artifactCmd {
	cmd := new(artifactCmd)
	cmd.Command = cobraCmd
	cmd.PersistentFlags().StringVar(&cmd.artifact, "artifact", ".", "file location of artifact to copy")
	cmd.PersistentFlags().StringVar(&cmd.group, "group", ".", "folder containing the artifacts")
	cmd.PersistentFlags().StringVar(&cmd.version, "version", ".", "version of the artifact")
	return cmd
}
