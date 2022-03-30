package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
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
	var wg sync.WaitGroup
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
		wg.Add(1)
		st := time.Now()
		go func(url string) {
			fetchBody(url)
			fmt.Printf("url: %v, fetched: %v\n", url, time.Since(st))
			wg.Done()
		}(url)
	}
	wg.Wait()
	fmt.Printf("End - elapse: %v\n", time.Since(startTime))
}
