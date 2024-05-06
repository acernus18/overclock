package overclock

import (
	"fmt"
	"regexp"
	"strconv"
	"testing"
)

func downloadComic(url string, quantity int) error {
	handler := func(name string) string {
		pattern := regexp.MustCompile(`(\d+)\.jpg`)
		if pattern.MatchString(name) {
			result := pattern.FindStringSubmatch(name)
			index, _ := strconv.ParseInt(result[1], 10, 32)
			return fmt.Sprintf("%d.jpg", index+155)
		}
		return name
	}
	for i := 0; i < quantity; i++ {
		if err := Download(fmt.Sprintf(url, i+1), "/Users/maples/Downloads/Temp/", handler); err != nil {
			return err
		}
	}
	return nil
}

func TestDownload(t *testing.T) {
	if err := downloadComic("https://cdn.hentaibomb.com/g/1803074/%d.jpg", 155); err != nil {
		panic(err)
	}
}

func TestDownloadHLSFile(t *testing.T) {
	results, err := DownloadHLSFile("https://zippey-gogin.mushroomtrack.com/hls/v7t0gokZ9b3L9C1sFHX_Hw/1714571214/35000/35390/35390.m3u8")
	if err != nil {
		t.Fatal(err)
	}
	//t.Log(results)

	for i := range results {
		url := fmt.Sprintf("https://zippey-gogin.mushroomtrack.com/hls/v7t0gokZ9b3L9C1sFHX_Hw/1714571214/35000/35390/%s", results[i])
		if err := Download(url, "/Users/maples/Downloads/Temp2/", func(s string) string {
			return s
		}); err != nil {
			t.Fatal(err)
		}
	}
}
