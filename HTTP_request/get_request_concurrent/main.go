package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

var url string = "https://jsonplaceholder.typicode.com/posts/"

func makeRequest(url string) {
	resp, err := http.Get(url)
	if err != nil {panic(err)}
	body, err := io.ReadAll(resp.Body)
	defer resp.Body.Close()
	fmt.Println(string(body))
}

func main() {
	for i := 1; i <= 100; i++ {
		go makeRequest(url + strconv.Itoa(i))
	}
	time.Sleep(1 * time.Second)
}