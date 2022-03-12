package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func handleConn(conn net.Conn) {
	defer conn.Close()
	io.WriteString(conn, "Hi\n")
}

const port = 8080

func main() {
	log.Printf("Starting tcp server at port %d\n", port)
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		log.Fatal("server unable to start")
	}
	log.Printf("Server started. Try to connect with command like: '$ nc localhost %d'\n", port)
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Print("connection not accepted")
			continue
		}
		go handleConn(conn)
	}

}
