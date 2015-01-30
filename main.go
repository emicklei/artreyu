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

// go run *.go archive --artifact=main.go --group=com.ubanita --version=1.0-SNAPSHOT --force=true

var RootCmd = &cobra.Command{
	Use:   "typhoon",
	Short: "typhoon a is a tool for artifact management",
	Run:   func(cmd *cobra.Command, args []string) {},
}

type archiveCmd struct {
	*cobra.Command
	artifact  string
	group     string
	version   string
	overwrite bool
}

func newArchiveCmd() *cobra.Command {
	cmd := new(archiveCmd)
	cmd.Command = &cobra.Command{
		Use:   "archive",
		Short: "copy an artifact to the typhoon repository",
		Run:   cmd.doArchive,
	}
	cmd.PersistentFlags().StringVar(&cmd.artifact, "artifact", ".", "file location of artifact to copy")
	cmd.PersistentFlags().StringVar(&cmd.group, "group", ".", "store the articat under this group")
	cmd.PersistentFlags().StringVar(&cmd.version, "version", ".", "store the articat under this version")
	cmd.PersistentFlags().BoolVar(&cmd.overwrite, "force", false, "force overwrite if version exists")
	return cmd.Command
}

func (a *archiveCmd) doArchive(cmd *cobra.Command, args []string) {
	repo := os.Getenv("TYPHOON_REPO")
	if len(repo) == 0 {
		log.Fatal("missing TYPHOON_REPO environment setting")
	}
	g := path.Join(strings.Split(a.group, ".")...)
	regular := path.Base(path.Clean(a.artifact))
	p := path.Join(repo, g, regular, a.version)

	log.Printf("copying %s to %s", regular, p)
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

func main() {
	RootCmd.AddCommand(newArchiveCmd())
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
