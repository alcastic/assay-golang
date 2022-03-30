package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"sync"
	"time"
)

type result struct {
	value []byte
	err   error
}

func fetchBody(url string) *result {
	res, err := http.Get(url)
	if err != nil {
		return &result{nil, err}
	}

	defer res.Body.Close()
	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return &result{nil, err}
	}
	return &result{data, nil}
}

type MemCache struct {
	f     func(string) *result
	cache map[string]*result
}

func (m *MemCache) call(s string) *result {
	r, ok := m.cache[s]
	if !ok {
		r = m.f(s)
		m.cache[s] = r
	}
	return r
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

	mc := &MemCache{
		f:     fetchBody,
		cache: make(map[string]result),
	}

	for _, url := range urls {
		wg.Add(1)
		st := time.Now()
		go func(url string) {
			mc.f(url)
			fmt.Printf("url: %v, fetched: %v\n", url, time.Since(st))
			wg.Done()
		}(url)
	}
	wg.Wait()
	fmt.Printf("End - elapse: %v\n", time.Since(startTime))
}
