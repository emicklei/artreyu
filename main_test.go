package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/emicklei/artreyu/model"
	"github.com/emicklei/artreyu/transport"
)

func setup(t *testing.T) {
	initRootCommand()
	model.Fatalf = t.Fatalf
	model.Printf = t.Logf
}

func TestArchive(t *testing.T) {
	setup(t)
	rootCmd.SetArgs([]string{
		"archive",
		"testdata/LoremIpsum.txt",
		"--config=testdata/local-config.yaml",
		"--descriptor=testdata/LoremIpsum.yaml",
		"--verbose=true",
		"--os=test"})
	rootCmd.Execute()

	if !transport.Exists("testdata/com/company/LoremIpsum/1.0-SNAPSHOT/test/LoremIpsum-1.0-SNAPSHOT.txt") {
		t.Fail()
	}
}

func TestFetch(t *testing.T) {
	setup(t)
	rootCmd.SetArgs([]string{
		"archive",
		"testdata/LoremIpsum.txt",
		"--config=testdata/local-config.yaml",
		"--descriptor=testdata/LoremIpsum.yaml",
		"--verbose=true",
		"--os=test"})
	rootCmd.Execute()

	setup(t)
	tmp := os.TempDir()
	rootCmd.SetArgs([]string{
		"fetch",
		tmp,
		"--config=testdata/local-config.yaml",
		"--descriptor=testdata/LoremIpsum.yaml",
		"--verbose=true",
		"--os=test"})
	rootCmd.Execute()

	if !transport.Exists(filepath.Join(tmp, "LoremIpsum-1.0-SNAPSHOT.txt")) {
		t.Fail()
	}
}

func TestAssemble(t *testing.T) {
	setup(t)

	rootCmd.SetArgs([]string{
		"archive",
		"testdata/LoremIpsum.txt",
		"--config=testdata/local-config.yaml",
		"--descriptor=testdata/LoremIpsum.yaml",
		"--verbose=true",
		"--os=test"})
	rootCmd.Execute()

	setup(t)
	rootCmd.SetArgs([]string{
		"archive",
		"testdata/doc.tgz",
		"--config=testdata/local-config.yaml",
		"--descriptor=testdata/artreyu.yaml",
		"--verbose=true",
		"--os=test"})
	rootCmd.Execute()

	setup(t)
	tmp := filepath.Join(os.TempDir(), "artreyu")
	rootCmd.SetArgs([]string{
		"assemble",
		tmp,
		"--config=testdata/local-config.yaml",
		"--descriptor=testdata/assembly.yaml",
		"--verbose=true",
		"--os=test"})
	rootCmd.Execute()

	if !transport.IsRegular("/tmp/artreyu/com/company/assembly/2/test/assembly-2.tgz") {
		t.Fail()
	}
	transport.FileRemove("/tmp/artreyu/com/company/assembly/2/test/assembly-2.tgz")
}

// clear && go test -v -test.run=TestAssembleZip
func TestAssembleZip(t *testing.T) {
	setup(t)

	rootCmd.SetArgs([]string{
		"archive",
		"testdata/LoremIpsum.txt",
		"--config=testdata/local-config.yaml",
		"--descriptor=testdata/LoremIpsum.yaml",
		"--verbose=true",
		"--os=test"})
	rootCmd.Execute()

	setup(t)
	rootCmd.SetArgs([]string{
		"archive",
		"testdata/doc.tgz",
		"--config=testdata/local-config.yaml",
		"--descriptor=testdata/artreyu.yaml",
		"--verbose=true",
		"--os=test"})
	rootCmd.Execute()

	setup(t)
	tmp := filepath.Join(os.TempDir(), "artreyu")
	rootCmd.SetArgs([]string{
		"assemble",
		tmp,
		"--config=testdata/local-config.yaml",
		"--descriptor=testdata/assembly-zip.yaml",
		"--verbose=true",
		"--os=test"})
	rootCmd.Execute()

	if !transport.IsRegular("/tmp/artreyu/com/company/assembly/2/test/assembly-2.zip") {
		t.Fail()
	}
	//transport.FileRemove("/tmp/artreyu/com/company/assembly/2/test/assembly-2.zip")
}
