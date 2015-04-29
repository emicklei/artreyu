package model

import (
	"os"
	"os/exec"
)

func Exists(loc string) bool {
	_, err := os.Stat(loc)
	return err == nil
}

func Cp(dst, src string) error {
	return exec.Command("cp", src, dst).Run()
}
