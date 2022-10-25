package overclock

import (
	"errors"
	"io"
	"log"
	"net/http"
	"os"
	"path"
	"regexp"
)

func Download(url, destination string) error {
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
	file, err := os.Create(path.Join(destination, matchResult[0]))
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
	log.Printf("downloaded \"%s\" from %s to %s, length = %d bytes", matchResult[0], url, destination, bytesLen)
	return nil
}
