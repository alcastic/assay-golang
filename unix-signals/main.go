package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func ipcHandler(signals <-chan os.Signal, server *http.Server) {
	signal := <-signals
	wt := 5 * time.Second
	fmt.Printf("Signal received: %v\n", signal)
	fmt.Printf("Server will be shutdown in %v\n", wt)
	time.Sleep(wt)
	if err := server.Shutdown(context.Background()); err != nil {
		log.Fatal(err)
	}
	os.Exit(0)
}

func main() {
	fmt.Printf("Process ID: %v\n", os.Getpid())

	var signals chan os.Signal = make(chan os.Signal, 1)
	signal.Notify(signals, syscall.SIGTERM)

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(rw http.ResponseWriter, rq *http.Request) {
		fmt.Fprintln(rw, "Hello World!")
	})
	server := &http.Server{Addr: ":8080", Handler: mux}

	go ipcHandler(signals, server)

	if err := server.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
}
