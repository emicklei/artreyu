package main

import (
	"bytes"
	"fmt"

	"github.com/emicklei/artreyu/model"
	"github.com/spf13/cobra"
)

func newTreeCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "tree",
		Short: "reads artifact descriptors recursively and prints the dependency hierarchy",
		Run:   doPrintTree,
	}
}

type node struct {
	label string
	parts []node
}

func (n *node) add(part node) {
	n.parts = append(n.parts, part)
}

func (n *node) String() string {
	buf := new(bytes.Buffer)
	n.printOn(0, buf)
	return buf.String()
}

func (n *node) printOn(level int, buf *bytes.Buffer) {

}

// TODO fetch the descriptors of each component
func doPrintTree(cmd *cobra.Command, args []string) {
	assembly, err := model.LoadAssembly(applicationSettings.ArtifactConfigLocation)
	if err != nil {
		model.Fatalf("tree failed, unable to load assembly descriptor:%v", err)
	}
	fmt.Println("|", assembly.StorageBase())
	for _, each := range assembly.Parts {
		fmt.Println("|-- ", each.StorageBase())
	}
}
