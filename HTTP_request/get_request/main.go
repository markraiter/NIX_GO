package main

import (
	"fmt"
	"io"
	"net/http"
)

func makeRequest() {
	url := "https://jsonplaceholder.typicode.com/posts/"

	resp, err := http.Get(url)
	if err != nil {panic(err)}
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Println(string(body))
}

func main() {
	makeRequest()
}