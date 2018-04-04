package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
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
			// ประกาศตัวแปรเพื่อเก็บว่ากำลัง scan บันทัดไหนอยู่
			i := 0

			scanner := bufio.NewScanner(conn)
			for scanner.Scan() {
				ln := scanner.Text()
				fmt.Println(ln)

				// เช็คเงื่อนไขเมื่อ scan บันทัดแรก เพื่อ get ค่า method และ uri ที่ได้จาก browser
				if i == 0 {
					// คล้ายๆ split ด้วยช่องว่าง
					words := strings.Fields(ln)

					// เก็บค่า method และ uri เข้าตัวแปร เพื่อนำไปใช้กำหนดเงื่อนไขการทำ router
					method := words[0]
					uri := words[1]
					fmt.Println("----METHOD = ", method)
					fmt.Println("----URI    = ", uri)
				}
				if ln == "" {
					break
				}
				i++
			}

			body := `<!DOCTYPE html><html lang="en"><head><meta charet="UTF-8"><title></title></head><body><strong>Hello World</strong></body></html>`
			fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
			fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
			fmt.Fprint(conn, "Content-Type: text/html\r\n")
			fmt.Fprint(conn, "\r\n")
			fmt.Fprint(conn, body)
			conn.Close()
		}(conn)
	}
}
