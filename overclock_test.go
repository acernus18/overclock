package overclock

import (
	"fmt"
	"testing"
)

func downloadComic(url string, quantity int) error {
	for i := 0; i < quantity; i++ {
		if err := Download(fmt.Sprintf(url, i+1), "/Users/maples/Downloads/111/"); err != nil {
			return err
		}
	}
	return nil
}

func TestDownload(t *testing.T) {
	if err := downloadComic("https://cdn.nhentai.xxx/g/1710910/%d.jpg", 23); err != nil {
		panic(err)
	}
}
