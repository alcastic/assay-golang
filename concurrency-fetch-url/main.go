package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	startTime := time.Now()
	fmt.Println("START")

	urls := os.Args[1:] // obtaion urls from args, os.Args[0] contains the program's name
	ch := make(chan string)

	// fetching urls concurrently
	for _, url := range urls {
		go fetchUrl(url, ch)
	}

	//collecting url fech results from channel
	for range urls {
		fmt.Println(<-ch)
	}

	fmt.Println("FINISH - elapsed:", time.Since(startTime).Seconds())
}

func fetchUrl(url string, ch chan<- string) {
	fetchStartTime := time.Now()

	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	bodyNroOfBytes, err := io.Copy(io.Discard, resp.Body)
	defer resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprint(err)
		return
	}

	ch <- fmt.Sprintf("duration: %.2f, body_nro_bytes: %10d, url: %s",
		time.Since(fetchStartTime).Seconds(), bodyNroOfBytes, url)
}
