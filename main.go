	// WRITE POSTS FROM GET REQUEST

package main

import (
	"io"
	"net/http"
	"os"
	"strconv"
	"time"
)

var url string = "https://jsonplaceholder.typicode.com/posts/"

func main()  {
	go makeRequestConcurrent()
	time.Sleep(2 * time.Second)
	
}

func makeRequestConcurrent() {
	for i := 1; i <= 100; i++ {
		resp, err := http.Get(url + strconv.Itoa(i))
		if err != nil {panic(err)}
		body,err := io.ReadAll(resp.Body)
		if err != nil {panic(err)}
		result := string(body)
		f, err := os.Create("post" + strconv.Itoa(i) + ".txt")
		if err != nil {panic(err)}
		f.WriteString(result)
	}
	return
}
