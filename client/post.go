package client

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// Post -
type Post struct {
	Title      string   `json:"title"`
	Categories []string `json:"categoryNames"`
	Content    string   `json:"content"`
}

// Upload upload to BoN
func Upload(post Post, token string) {
	url := "https://www.noye.xyz/api/v1/post"

	pByte, _ := json.Marshal(post)
	buff := bytes.NewBuffer(pByte)

	req, err := http.NewRequest("POST", url, buff)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("token", token)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	println(resp.StatusCode)
}
