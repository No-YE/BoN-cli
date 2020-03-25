package markdown

import (
	"io/ioutil"
	"os"
)

// Read read markdown file
func Read(name string) string {
	path, err := os.Getwd()
	if err != nil {
		println(err)
		os.Exit(1)
	}

	dat, err := ioutil.ReadFile(path + "/" + name)
	if err != nil {
		println(err)
		os.Exit(1)
	}

	return string(dat)
}
