package transport

import (
	"io"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"

	"github.com/emicklei/artreyu/model"
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

// Copy does what is says. Ignores errors on Close though.
func Copy(dst, src string) error {
	os.MkdirAll(filepath.Dir(dst), os.ModePerm)
	cleanSrc, err := filepath.Abs(path.Clean(src))
	if err != nil {
		return err
	}
	cleanDst, err := filepath.Abs(path.Clean(dst))
	if err != nil {
		return err
	}

	in, err := os.Open(cleanSrc)
	if err != nil {
		return err
	}
	defer in.Close()
	out, err := os.Create(cleanDst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, in)
	cerr := out.Close()
	if err != nil {
		return err
	}
	model.Printf("copy [%s] to [%s]", cleanSrc, cleanDst)
	return cerr
}

func IsTargz(filenameOrExtension string) bool {
	return strings.HasSuffix(filenameOrExtension, ".tar.gz") ||
		strings.HasSuffix(filenameOrExtension, ".tgz")
}

func IsZip(filenameOrExtension string) bool {
	return strings.HasSuffix(filenameOrExtension, ".zip")
}

func Targz(sourceDir, destinationFile string) error {
	model.Printf("compress into tape archive [%s] from [%s]\n", destinationFile, sourceDir)
	tmp := filepath.Join(os.TempDir(), filepath.Base(destinationFile))
	cmd, _ := asCommand(
		"tar",
		"czvf",
		tmp,
		"-C",
		sourceDir,
		".")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	if err := cmd.Run(); err != nil {
		return err
	}
	// now move to destination
	return os.Rename(tmp, destinationFile)
}

func Untargz(sourceFile, destinationDir string) error {
	model.Printf("extract from tape archive [%s] to [%s]\n", sourceFile, destinationDir)
	cmd, _ := asCommand(
		"tar",
		"xvf",
		sourceFile,
		"-C",
		destinationDir)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func Zip(sourceDir, destinationFile string) error {
	model.Printf("compress into zip archive [%s] from [%s]\n", destinationFile, sourceDir)
	tmp := filepath.Join(os.TempDir(), filepath.Base(destinationFile))
	if err := createZip(sourceDir, tmp); err != nil {
		return err
	}
	// now move to destination
	return os.Rename(tmp, destinationFile)
}

func Unzip(sourceFile, destinationDir string) error {
	model.Printf("decompress from zip archive [%s] to [%s]\n", sourceFile, destinationDir)
	return extractZip(sourceFile, destinationDir)
}

func FileRemove(source string) error {
	model.Printf("removing [%s]\n", source)
	return os.Remove(source)
}

func asCommand(params ...string) (*exec.Cmd, string) {
	return exec.Command(params[0], params[1:]...), strings.Join(params, " ")
}
