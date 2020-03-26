package main

import "fmt"

func main() {
	if err := Execute(); err != nil {
		fmt.Println(err)
	}
}
