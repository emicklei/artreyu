package model

import (
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

func IsDirectory(loc string) bool {
	s, err := os.Stat(loc)
	return err == nil && s.IsDir()
}

func IsRegular(loc string) bool {
	s, err := os.Stat(loc)
	return err == nil && !s.IsDir()
}
func Exists(loc string) bool {
	_, err := os.Stat(loc)
	return err == nil
}

// Cp copy a file (src) to new file (dst).
// Dst is the full directory path and file name
// Src can be a relative directory path and file name
func Cp(dst, src string) error {
	cleanSrc := path.Clean(src)
	cleanDst := path.Clean(dst)
	log.Printf("copy %s to %s", cleanSrc, cleanDst)
	if err := os.MkdirAll(filepath.Dir(cleanDst), os.ModePerm); err != nil {
		return err
	}
	return exec.Command("cp", cleanSrc, cleanDst).Run()
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

func Targz(sourceDir, destinationFile string) error {
	log.Printf("creating tape archive %s from %s\n", destinationFile, sourceDir)
	// and exclude the archive created
	cmd := exec.Command(
		"tar",
		"-czvf",
		destinationFile,
		"--exclude",
		filepath.Base(destinationFile),
		"-C",
		sourceDir,
		".")
	data, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(string(data))
	}
	return err
}

func Untargz(sourceFile, destinationDir string) error {
	log.Printf("extracting tape archive %s to %s\n", sourceFile, destinationDir)
	cmd := exec.Command(
		"tar",
		"-xvf",
		sourceFile,
		"-C",
		destinationDir)
	data, err := cmd.CombinedOutput()
	if err != nil {
		log.Println(string(data))
	}
	return err
}

func FileRemove(source string) error {
	log.Printf("removing %s\n", source)
	return os.Remove(source)
}
