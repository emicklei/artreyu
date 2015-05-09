package main

import (
	"os"
	"path/filepath"
	"testing"

	"github.com/emicklei/artreyu/model"
)

func setup(t *testing.T) {
	initRootCommand()
	model.Fatalf = t.Fatalf
	model.Printf = t.Logf
}

func TestArchive(t *testing.T) {
	setup(t)
	RootCmd.SetArgs([]string{
		"archive",
		"testdata/LoremIpsum.txt",
		"--config=testdata/local-config.yaml",
		"--descriptor=testdata/LoremIpsum.yaml",
		"--verbose=true",
		"--os=test"})
	RootCmd.Execute()

	if !model.Exists("/tmp/artreyu/com/company/LoremIpsum/1.0-SNAPSHOT/test/LoremIpsum-1.0-SNAPSHOT.txt") {
		t.Fail()
	}
}

func TestFetch(t *testing.T) {
	setup(t)
	RootCmd.SetArgs([]string{
		"archive",
		"testdata/LoremIpsum.txt",
		"--config=testdata/local-config.yaml",
		"--descriptor=testdata/LoremIpsum.yaml",
		"--verbose=true",
		"--os=test"})
	RootCmd.Execute()

	setup(t)
	tmp := os.TempDir()
	RootCmd.SetArgs([]string{
		"fetch",
		tmp,
		"--config=testdata/local-config.yaml",
		"--descriptor=testdata/LoremIpsum.yaml",
		"--verbose=true",
		"--os=test"})
	RootCmd.Execute()

	if !model.Exists(filepath.Join(tmp, "LoremIpsum-1.0-SNAPSHOT.txt")) {
		t.Fail()
	}
}

func TestAssemble(t *testing.T) {
	setup(t)

	RootCmd.SetArgs([]string{
		"archive",
		"testdata/LoremIpsum.txt",
		"--config=testdata/local-config.yaml",
		"--descriptor=testdata/LoremIpsum.yaml",
		"--verbose=true",
		"--os=test"})
	RootCmd.Execute()

	setup(t)
	RootCmd.SetArgs([]string{
		"archive",
		"testdata/doc.tgz",
		"--config=testdata/local-config.yaml",
		"--descriptor=testdata/artreyu.yaml",
		"--verbose=true",
		"--os=test"})
	RootCmd.Execute()

	setup(t)
	tmp := filepath.Join(os.TempDir(), "artreyu")
	RootCmd.SetArgs([]string{
		"assemble",
		tmp,
		"--config=testdata/local-config.yaml",
		"--descriptor=testdata/assembly.yaml",
		"--verbose=true",
		"--os=test"})
	RootCmd.Execute()

	if !model.IsRegular("/tmp/artreyu/com/company/assembly/2/test/assembly-2.tgz") {
		t.Fail()
	}
}
