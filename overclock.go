package overclock

import (
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"regexp"
)

func Fetch(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	content, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

func Download(url, destination string, filenameHandler func(string) string) error {
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	resourcesMatchPattern := regexp.MustCompile(`([^/]+?\.[^/]+)$`)
	matchResult := resourcesMatchPattern.FindStringSubmatch(url)
	if len(matchResult) <= 0 {
		return errors.New("cannot find valid filename")
	}
	targetFilename := filenameHandler(matchResult[0])
	file, err := os.Create(path.Join(destination, targetFilename))
	if err != nil {
		return err
	}
	defer func() {
		err = resp.Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}()
	bytesLen, err := io.Copy(file, resp.Body)
	if err != nil {
		return err
	}
	log.Printf("downloaded \"%s\" from %s to %s, length = %d bytes", targetFilename, url, destination, bytesLen)
	return nil
}

func DownloadHLSFile(url string) ([]string, error) {
	pattern := regexp.MustCompile(`\d+.ts`)
	content, err := Fetch(url)
	if err != nil {
		return nil, err
	}
	return pattern.FindAllString(content, -1), nil
}
