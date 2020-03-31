package keystore

import (
	"fmt"
	"os"
	"path"

	"github.com/peterbourgon/diskv"
)

var (
	d = diskv.New(diskv.Options{
		BasePath:     path.Join(os.Getenv("GOPATH"), "bon-store"),
		CacheSizeMax: 1 * 1,
	})
	key = "token"
)

// SaveToken save token as file
func SaveToken(token string) {
	d.Write(key, []byte(token))
}

// GetToken get token from file
func GetToken() string {
	token, err := d.Read(key)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return string(token)
}
