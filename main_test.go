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
		filepath.Join("testdata", "LoremIpsum.txt"),
		"--config=testdata/local-config.yaml",
		"--descriptor=testdata/LoremIpsum.yaml",
		"--verbose=true",
		"--os=testOS"})
	rootCmd.Execute()

	if !transport.Exists("testdata/com/company/LoremIpsum/1.0-SNAPSHOT/testOS/LoremIpsum-1.0-SNAPSHOT.txt") {
		t.Fail()
	}
}

func TestFetch(t *testing.T) {
	setup(t)
	rootCmd.SetArgs([]string{
		"archive",
		filepath.Join("testdata", "LoremIpsum.txt"),
		"--config=testdata/local-config.yaml",
		"--descriptor=testdata/LoremIpsum.yaml",
		"--verbose=true",
		"--os=testOS"})
	rootCmd.Execute()

	setup(t)
	tmp := os.TempDir()
	rootCmd.SetArgs([]string{
		"fetch",
		tmp,
		"--config=testdata/local-config.yaml",
		"--descriptor=testdata/LoremIpsum.yaml",
		"--verbose=true",
		"--os=testOS"})
	rootCmd.Execute()

	if !transport.Exists(filepath.Join(tmp, "LoremIpsum-1.0-SNAPSHOT.txt")) {
		t.Fail()
	}
}

func TestAssemble(t *testing.T) {
	setup(t)
	t.Log("---\narchive LoremIpsum.txt")
	rootCmd.SetArgs([]string{
		"archive",
		filepath.Join("testdata", "LoremIpsum.txt"),
		"--config=testdata/local-config.yaml",
		"--descriptor=testdata/LoremIpsum.yaml",
		"--verbose=true",
		"--os=testOS"})
	rootCmd.Execute()

	setup(t)
	t.Log("---\narchive doc.tgz")
	rootCmd.SetArgs([]string{
		"archive",
		filepath.Join("testdata", "doc.tgz"),
		"--config=testdata/local-config.yaml",
		"--descriptor=testdata/artreyu.yaml",
		"--verbose=true",
		"--os=testOS"})
	rootCmd.Execute()

	t.Log("---\nassemble assembly-2.tgz")
	setup(t)
	tmp := filepath.Join(os.TempDir(), "artreyu")
	rootCmd.SetArgs([]string{
		"assemble",
		tmp,
		"--config=testdata/local-config.yaml",
		"--descriptor=testdata/assembly.yaml",
		"--verbose=true",
		"--os=testOS"})
	rootCmd.Execute()

	if !transport.IsRegular("testdata/com/company/assembly/2/testOS/assembly-2.tgz") {
		t.Error("Expected to see file:", "testdata/com/company/assembly/2/testOS/assembly-2.tgz")
	}
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
		"--os=testOS"})
	rootCmd.Execute()

	setup(t)
	rootCmd.SetArgs([]string{
		"archive",
		"testdata/doc.tgz",
		"--config=testdata/local-config.yaml",
		"--descriptor=testdata/artreyu.yaml",
		"--verbose=true",
		"--os=testOS"})
	rootCmd.Execute()

	setup(t)
	tmp := filepath.Join(os.TempDir(), "artreyu")
	rootCmd.SetArgs([]string{
		"assemble",
		tmp,
		"--config=testdata/local-config.yaml",
		"--descriptor=testdata/assembly-zip.yaml",
		"--verbose=true",
		"--os=testOS"})
	rootCmd.Execute()

	if !transport.IsRegular("testdata/com/company/assembly/2/testOS/assembly-2.zip") {
		t.Error("Expected to see file:", "testdata/com/company/assembly/2/testOS/assembly-2.zip")
	}
}
