package client

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func getPresignedURL(filename string, token string) string {
	url := "https://www.noye.xyz/api/v1/image"

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	q := req.URL.Query()
	q.Add("kind", "post")
	q.Add("filename", filename)

	req.URL.RawQuery = q.Encode()
	req.Header.Add("token", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return string(data)
}

// Upload upload to BoN
func putS3(url string, path string) {
	data, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	req, err := http.NewRequest("PUT", url, data)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fmt.Println(resp.StatusCode)
}
