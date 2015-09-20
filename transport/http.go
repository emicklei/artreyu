package transport

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func HttpGetFile(sourceURL, destinationFilename string) error {
	req, err := http.NewRequest("GET", sourceURL, nil)
	if err != nil {
		return err
	}
	resp, err := new(http.Client).Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		return fmt.Errorf("GET failed:%s", resp.Status)
	}
	destination, err := os.Create(destinationFilename)
	if err != nil {
		return err
	}
	defer destination.Close()
	_, err = io.Copy(destination, resp.Body)
	return err
}

func HttpPostFile(sourceFilename, destinationURL string) error {
	source, err := os.Open(sourceFilename)
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", destinationURL, source)
	//req.Header.Set("Expect", "100-continue")

	fmt.Println(req.Header)

	if err != nil {
		return err
	}
	resp, err := new(http.Client).Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	if resp.StatusCode != 201 {
		return fmt.Errorf("POST failed:%s", resp.Status)
	}
	return err

	/**
		// Because of https://github.com/golang/go/issues/3665, we need curl to do the job
		cmd := exec.Command(
			"curl",
			"--upload-file",
			sourceFilename,
			destinationURL)
		if model.Verbose {
			model.Printf("%v", cmd.Args)
		}
		data, err := cmd.CombinedOutput()
		if err != nil || model.Verbose {
			model.Printf("%s", string(data))
		}
		return err
	**/
}
