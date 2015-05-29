package command

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"github.com/emicklei/artreyu/model"
	"github.com/spf13/cobra"
)

// ArtifactFunc is a provider function to return the Artifact
type ArtifactFunc func() model.Artifact

// RepositoryFunc is a provider function to return the Repository
type RepositoryFunc func() model.Repository

// NewSettingsBoundToFlags returns a new Settings with values that are bound to command line flag values.
func NewSettingsBoundToFlags(cmd *cobra.Command) *model.Settings {
	settings := new(model.Settings)
	cmd.PersistentFlags().StringVarP(&settings.MainConfigLocation,
		"config",
		"c",
		filepath.Join(os.Getenv("HOME"), ".artreyu"),
		"location of the artreyu repositories configuration")
	cmd.PersistentFlags().StringVarP(&settings.ArtifactConfigLocation,
		"descriptor",
		"d",
		"artreyu.yaml",
		"overwrite if the artifact descriptor has a different name or location")
	cmd.PersistentFlags().StringVarP(&settings.OS,
		"os",
		"o",
		runtime.GOOS,
		"overwrite if assembling for different OS")
	cmd.PersistentFlags().BoolVarP(&settings.Verbose,
		"verbose",
		"v",
		false,
		"set to true to log more execution details")
	cmd.PersistentFlags().StringVarP(&settings.TargetRepository,
		"repository",
		"r",
		"local",
		"name of the repository as defined in the artreyu repositories configuration")
	return settings
}

// NewPluginCommand returns the values that are needed to create an Artreyu plugin program.
func NewPluginCommand() (*cobra.Command, *model.Settings, *model.Artifact) {
	cmd := new(cobra.Command)
	artifact := new(model.Artifact)
	cmd.PersistentFlags().StringVarP(&artifact.Name,
		"artifact",
		"a",
		"",
		"name of the artifact")
	cmd.PersistentFlags().StringVarP(&artifact.Group,
		"group",
		"g",
		"",
		"name of the group")
	cmd.PersistentFlags().StringVarP(&artifact.Version,
		"version",
		"s",
		"",
		"version of the artifact")
	cmd.PersistentFlags().StringVarP(&artifact.Type,
		"type",
		"t",
		"",
		"type (extension) of the artifact")
	cmd.PersistentFlags().BoolVarP(&artifact.AnyOS,
		"any-os",
		"y",
		false,
		"if true then the artifact is operating system independent")
	return cmd, NewSettingsBoundToFlags(cmd), artifact
}

// RunPlugin starts an external program with parameters from settings and artifacts loaded via settings and extra cli arguments.
func RunPlugin(programName, subCommand string, settings model.Settings, args []string) error {
	a, err := model.LoadArtifact(settings.ArtifactConfigLocation)
	if err != nil {
		model.Fatalf("invalid artifact descriptor:%v", err)
		return nil
	}
	return RunPluginWithArtifact(programName, subCommand, a, settings, args)
}

// RunPluginWithArtifact starts an external program with parameters from settings,artifacts and extra cli arguments.
func RunPluginWithArtifact(programName, subCommand string, artifact model.Artifact, settings model.Settings, args []string) error {
	params := append([]string{subCommand}, settings.PluginParameters()...)
	params = append(params, artifact.PluginParameters()...)
	params = append(params, args...)
	plugin := exec.Command(programName, params...)
	if settings.Verbose {
		model.Printf("%v", plugin.Args)
	}
	plugin.Stdout = os.Stdout
	plugin.Stderr = os.Stderr
	err := plugin.Run()
	if err != nil {
		model.Fatalf("unable to run plugin [%s], %v", programName, err)
	}
	return err
}
