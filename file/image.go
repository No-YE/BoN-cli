package file

import (
	"fmt"
	"io/ioutil"
	"os"
)

// Image read markdown file
func Image(path string) []byte {
	dat, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return dat
}
