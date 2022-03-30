package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// Ignore fetch error
func fetchBody(url string) []byte {
	res, err := http.Get(url)
	if err != nil {
		return nil
	}

	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil
	}
	return data
}

func main() {
	startTime := time.Now()
	fmt.Println("Init")
	urls := []string{
		"https://golang.org",
		"https://godoc.org",
		"https://play.golang.org",
		"https://gopl.io",
		"https://golang.org",
		"https://godoc.org",
		"https://play.golang.org",
		"https://gopl.io",
	}
	for _, url := range urls {
		st := time.Now()
		fetchBody(url)
		fmt.Printf("url: %v, fetched: %v\n", url, time.Since(st))
	}
	fmt.Printf("End - elapse: %v\n", time.Since(startTime))
}
