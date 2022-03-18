package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"time"
)

const network = "tcp"
const port = 8080
const connPoolSize = 2

func main() {
	listener, err := net.Listen(network, fmt.Sprintf(":%d", port))
	connPool := make(chan struct{}, connPoolSize)

	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
		}
		go handleConnection(conn, connPool)
	}
}

func handleConnection(conn net.Conn, connPool chan struct{}) {
	connPool <- struct{}{}
	log.Printf("connection incomming - %s \n", conn.RemoteAddr().String())
	defer conn.Close()
	defer log.Printf("connection processed - %s\n", conn.RemoteAddr().String())
	defer func() { <-connPool }()
	for {
		_, err := io.WriteString(conn, fmt.Sprintf("Hi - %s\n", time.Now().Format("2006-01-02T15:04:05Z07:00")))
		if err != nil {
			fmt.Println("message not sent, client have close the connection")
			return
		}
		time.Sleep(5 * time.Second)

	}

}
