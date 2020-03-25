package keystore

import (
	"os"

	"github.com/peterbourgon/diskv"
)

var (
	d = diskv.New(diskv.Options{
		BasePath:     "bon",
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
		println(err)
		os.Exit(1)
	}

	return string(token)
}
