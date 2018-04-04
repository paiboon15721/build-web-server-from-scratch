package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}

	for {
		conn, err := li.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go func(conn net.Conn) {
			scanner := bufio.NewScanner(conn)
			for scanner.Scan() {
				ln := scanner.Text()
				fmt.Println(ln)
				if ln == "" {
					break
				}
			}

			// utf-8 byte slice
			body := string([]byte{224, 184, 151, 224, 184, 148, 224, 184, 170, 224, 184, 173, 224, 184, 154, 224, 184, 160, 224, 184, 178, 224, 184, 169, 224, 184, 178, 224, 185, 132, 224, 184, 151, 224, 184, 162})
			// tis-620 byte slice
			// body := string([]byte{183, 180, 202, 205, 186, 192, 210, 201, 210, 228, 183, 194, 32})
			fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
			fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
			fmt.Fprint(conn, "Content-Type: text/plain; charset=utf-8\r\n")
			// fmt.Fprint(conn, "Content-Type: text/plain; charset=tis-620\r\n")
			fmt.Fprint(conn, "\r\n")
			fmt.Fprint(conn, body)
			conn.Close()
		}(conn)
	}
}
