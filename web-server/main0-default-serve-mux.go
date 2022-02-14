package main

import (
	"fmt"
	"log"
	"net/http"
)

// Using http.DefaultServeMux
func main() {
	http.HandleFunc("/index", indexHandleFunc)
	http.HandleFunc("/hi", hiHandleFunc)
	log.Fatal(http.ListenAndServe(":8080", nil)) // nil means http.DefaultServeMux to be used
}

func indexHandleFunc(rw http.ResponseWriter, rq *http.Request) {
	fmt.Fprintln(rw, "Index")
}
func hiHandleFunc(rw http.ResponseWriter, rq *http.Request) {
	fmt.Fprintln(rw, "Hello, world!")
}
