// GET REQUEST FROM https://jsonplaceholder.typicode.com/

package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"
)

var url string = "https://jsonplaceholder.typicode.com/posts/"
var c chan string
var result string

func makeRequest(url string) {
	resp, err := http.Get(url)
	if err != nil {panic(err)}

	body, err := io.ReadAll(resp.Body)
	if err != nil {panic(err)}

	defer resp.Body.Close()

	result = string(body)

	fmt.Println(result)
}

func main()  {
	// makeRequest(url)
	// GET REQUEST FROM https://jsonplaceholder.typicode.com/ USING CONCURRENCY (GOROUTINES)
	for i := 1; i <= 100; i++ {
		go makeRequest(url + strconv.Itoa(i))
	}
	time.Sleep(2 * time.Second)	
}
