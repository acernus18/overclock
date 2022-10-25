package overclock

import (
	"fmt"
	"testing"
)

func TestDownload(t *testing.T) {
	//Download("https://cdn.nhentai.xxx/g/2062987/1.jpg", "/Users/maples/Downloads/")

	for i := 0; i < 16; i++ {
		if err := Download(fmt.Sprintf("https://cdn.nhentai.xxx/g/1816280/%d.jpg", i+1), "/Users/maples/Downloads/111/"); err != nil {
			panic(err)
		}
	}
}
