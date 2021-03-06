package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"

	"github.com/alcastic/assay-golang/tree/main/concurrency-memcache/v5/memo"
)

func fetchBody(url string) ([]byte, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	return data, nil
}

func fetchBodyAdapter(url string) (interface{}, error) {
	return fetchBody(url)
}

func main() {
	startTime := time.Now()
	var wg sync.WaitGroup
	fmt.Println("Init")
	urls := []string{
		"https://golang.org",
		"https://golang.org",
		"https://godoc.org",
		"https://godoc.org",
		"https://play.golang.org",
		"https://play.golang.org",
		"https://gopl.io",
		"https://gopl.io",
	}

	mc := memo.New(fetchBodyAdapter)

	for _, url := range urls {
		wg.Add(1)
		go func(url string) {
			mc.Call(url)
			wg.Done()
		}(url)
	}
	wg.Wait()
	memo.Close(mc)
	fmt.Printf("End - elapse: %v\n", time.Since(startTime))

}
