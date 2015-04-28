package main

import (
	"log"
	"path"
	"strings"

	"github.com/spf13/cobra"
)

type fetchCmd struct {
	*artifactCmd
}

func newFetchCmd() *cobra.Command {
	cmd := newArtifactCmd(&cobra.Command{
		Use:   "fetch [destination]",
		Short: "copy an artifact from the typhoon repository to [destination]",
	})
	fetch := new(fetchCmd)
	fetch.artifactCmd = cmd
	cmd.Command.Run = fetch.doFetch
	return cmd.Command
}

func (f *fetchCmd) doFetch(cmd *cobra.Command, args []string) {
	g := path.Join(strings.Split(f.group, ".")...)
	src := path.Join(getRepo(), g, f.artifact, f.version, f.artifact)
	if len(args) == 0 {
		log.Fatalf("missing destination")
	}
	if !Exists(src) {
		log.Fatalf("unable to copy artifact: %s because: no such artifact", src)
	}
	destination := args[len(args)-1]
	log.Printf("copying %s to %s\n", src, destination)
	if err := Cp(destination, src); err != nil {
		log.Fatalf("unable to copy artifact: %s to: %s because:%v", src, destination, err)
	}
}
