package main

import "github.com/spf13/cobra"

func newListCmd() *cobra.Command {
	cmd := newArtifactCmd(&cobra.Command{
		Use:   "list",
		Short: "list all available artifacts from the typhoon repository",
	})
	cmd.Command.Run = cmd.doList
	return cmd.Command
}

func (c *artifactCmd) doList(cmd *cobra.Command, args []string) {
	//	g := path.Join(strings.Split(c.group, ".")...)
	//	group := path.Join(getRepo(), g, c.artifact, c.version)
	//	files, _ := ioutil.ReadDir(group)
	//	for _, f := range files {
	//		fmt.Println(f.Name())
	//	}
}
