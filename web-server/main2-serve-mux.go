package main

import (
	"fmt"
	"net/http"
)

func main() {
	// ServeMux is a request multiplexer, it aggregates a collection of http.Handler into a single http.Handler
	sm := new(http.ServeMux)
	sm.HandleFunc("index", indexHandleFunc)
	// Go is first-class function, so, function can also implements an interface.
	// Check the implementation of http.HandlerFunc to see a sample
	sm.Handle("/hi", http.HandlerFunc(indexHandleFunc))
	http.ListenAndServe(":8080", sm)
}

func indexHandleFunc(rw http.ResponseWriter, rq *http.Request) {
	fmt.Fprintln(rw, "Index")
}
func hiHandleFunc(rw http.ResponseWriter, rq *http.Request) {
	fmt.Fprintln(rw, "Hello, world!")
}
