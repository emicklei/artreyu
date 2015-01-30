package main

import (
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/spf13/cobra"
)

// go run *.go archive --artifact=main.go --group=com.ubanita --version=1.0-SNAPSHOT
// go run *.go fetch --artifact=main.go --group=com.ubanita --version=1.0-SNAPSHOT hier

var VERSION string = "dev"
var BUILDDATE string = "now"

var RootCmd = &cobra.Command{
	Use:   "typhoon",
	Short: "typhoon a is a tool for artifact management",
	Run:   func(cmd *cobra.Command, args []string) {},
}

type artifactCmd struct {
	*cobra.Command
	artifact string
	group    string
	version  string
}

type archiveCmd struct {
	*artifactCmd
	overwrite bool
}

func newArtifactCmd(cobraCmd *cobra.Command) *artifactCmd {
	cmd := new(artifactCmd)
	cmd.Command = cobraCmd
	cmd.PersistentFlags().StringVar(&cmd.artifact, "artifact", ".", "file location of artifact to copy")
	cmd.PersistentFlags().StringVar(&cmd.group, "group", ".", "store the articat under this group")
	cmd.PersistentFlags().StringVar(&cmd.version, "version", ".", "store the articat under this version")
	return cmd
}

func newArchiveCmd() *cobra.Command {
	cmd := newArtifactCmd(&cobra.Command{
		Use:   "archive [artifact]",
		Short: "copy an artifact to the typhoon repository",
	})
	archiveCmd := new(archiveCmd)
	archiveCmd.artifactCmd = cmd
	archiveCmd.PersistentFlags().BoolVar(&archiveCmd.overwrite, "force", false, "force overwrite if version exists")
	cmd.Command.Run = archiveCmd.doArchive
	return cmd.Command
}

func getRepo() string {
	repo := os.Getenv("TYPHOON_REPO")
	if len(repo) == 0 {
		log.Fatal("missing TYPHOON_REPO environment setting")
	}
	return repo
}

func (a *archiveCmd) doArchive(cmd *cobra.Command, args []string) {
	if len(args) == 0 {
		log.Fatalf("missing artifact")
	}
	a.artifact = args[len(args)-1]
	g := path.Join(strings.Split(a.group, ".")...)
	regular := path.Base(path.Clean(a.artifact))
	p := path.Join(getRepo(), g, regular, a.version)

	log.Printf("copying %s into folder %s\n", regular, p)
	if err := os.MkdirAll(p, os.ModePerm); err != nil {
		log.Fatalf("unable to create dirs: %s cause: %v", p, err)
	}
	dest := path.Join(p, regular)

	// SNAPSHOT can be overwritten
	if strings.HasSuffix(a.version, "SNAPSHOT") {
		log.Println("will overwrite|create SNAPSHOT version")
		a.overwrite = true
	}
	if !a.overwrite && Exists(dest) {
		log.Fatalf("unable to copy artifact: %s to: %s cause: it already exists and --force=false", regular, p)
	}
	if err := Cp(dest, a.artifact); err != nil {
		log.Fatalf("unable to copy artifact: %s to: %s cause:%v", regular, p, err)
	}
}

type fetchCmd struct {
	*artifactCmd
	destination string
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
	destination := args[len(args)-1]
	log.Printf("copying %s to %s\n", src, destination)
	if err := Cp(destination, src); err != nil {
		log.Fatalf("unable to copy artifact: %s to: %s cause:%v", src, f.destination, err)
	}
}

func main() {
	log.Println("_/^\\_")
	log.Println(" | | typhoon - the artifact tool [commit=", VERSION, "build=", BUILDDATE, "]")
	log.Println("-\\_/-")
	RootCmd.AddCommand(newArchiveCmd())
	RootCmd.AddCommand(newFetchCmd())
	RootCmd.Execute()
}

func Exists(dest string) bool {
	_, err := os.Stat(dest)
	return err == nil
}

func Cp(dst, src string) error {
	return exec.Command("cp", src, dst).Run()
}

// Copy does what is says. Ignores errors on Close though.
func Copy(dst, src string) error {
	in, err := os.Open(src)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, in)
	cerr := out.Close()
	if err != nil {
		return err
	}
	return cerr
}
