package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/hi", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%v\n", err)
	}
}

func handler(rw http.ResponseWriter, rq *http.Request) {
	fmt.Fprintln(rw, "Hello, world!")
}
