package main

import (
	"os"
	"text/template"

	"github.com/emicklei/artreyu/model"
	"github.com/spf13/cobra"
)

func newFormatCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "format [template]",
		Short: "process the artifact descriptor with the given template and write to stdout",
		Run:   doFormat,
	}
}

func doFormat(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		model.Fatalf("missing template")
	}

	assembly, err := model.LoadAssembly(applicationSettings.ArtifactConfigLocation)
	if err != nil {
		model.Fatalf("format failed, unable to load assembly descriptor:%v", err)
	}

	tmp, err := template.New("temporary").Parse(args[0])
	if err != nil {
		model.Fatalf("format failed, cannot parse template: %v", err)
	}
	var value interface{}
	value = assembly
	// unless no parts, use the artifact instead
	if len(assembly.Parts) == 0 {
		value = assembly.Artifact
	}

	if err := tmp.Execute(os.Stdout, value); err != nil {
		model.Fatalf("format failed, cannot execute template: %v", err)
	}
}
