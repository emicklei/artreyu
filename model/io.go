package model

import (
	"io"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
)

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
