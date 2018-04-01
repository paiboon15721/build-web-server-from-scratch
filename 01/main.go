package main

import (
	"io"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalln(err)
	}
	conn, err := li.Accept()
	if err != nil {
		log.Println(err)
	}
	io.WriteString(conn, "This message from TCP server")
}
