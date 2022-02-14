package main

import (
	"fmt"
	"log"
	"net/http"
)

type stockHandler struct{}

func (h stockHandler) ServeHTTP(w http.ResponseWriter, rq *http.Request) { // with this method, stockHandler is a http.Handler interface
	switch rq.URL.Path {
	case "/index":
		fmt.Fprintln(w, "Index")
		return
	case "/hi":
		fmt.Fprintln(w, "Hello, world!")
		return
	default:
		http.Error(w, "Not Found", http.StatusNotFound)
	}
}

func main() {
	handler := stockHandler{}
	log.Fatal(http.ListenAndServe(":8080", handler))
}
