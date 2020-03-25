package markdown

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Read read markdown file
func Read(path string) string {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return string(dat)
}
