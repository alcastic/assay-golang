package main

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
)

const network = "tcp"
const serverPort = 8080

func main() {
	conn, err := net.Dial(network, fmt.Sprintf(":%d", serverPort))
	defer conn.Close()
	if err != nil {
		log.Fatalln(err)
		return
	}
	io.Copy(os.Stdout, conn)
}
