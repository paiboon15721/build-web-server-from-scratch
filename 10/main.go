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

		go func() {
			scanner := bufio.NewScanner(conn)
			for scanner.Scan() {
				ln := scanner.Text()
				fmt.Println(ln)
				if ln == "" {
					break
				}
			}

			body := `<img src="https://assets.teenvogue.com/photos/5925af0bf5c4720abcde5c0b/3:2/w_1200,h_630,c_limit/cat-fb.jpg">`
			fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
			fmt.Fprint(conn, "Content-Type: text/html\r\n\r\n")
			fmt.Fprint(conn, body)
			conn.Close()
		}()
	}
}
