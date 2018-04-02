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
				// ถ้าเจอบันทัดไหนเป็นว่างๆ ให้ถือว่าหมดส่วนของ header และ ออกจาก loop เพื่อหยุด scan
				if ln == "" {
					break
				}
			}

			// return html response to client
			body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body><strong>Hello World</strong></body></html>`
			fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
			fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
			fmt.Fprint(conn, "Content-Type: text/html; charset=utf-8\r\n")
			fmt.Fprint(conn, "\r\n")
			fmt.Fprint(conn, body)
			conn.Close()
		}()
	}
}
