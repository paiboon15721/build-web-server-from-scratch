package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"os"
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
			// ประกาศตัวแปรตรงนี้ เนื่องจากนำไปใช้ข้างนอก for loop
			var method string
			var uri string

			i := 0
			scanner := bufio.NewScanner(conn)
			for scanner.Scan() {
				ln := scanner.Text()
				fmt.Println(ln)

				if i == 0 {
					words := strings.Fields(ln)
					method = words[0]
					uri = words[1]
				}
				if ln == "" {
					break
				}
				i++
			}

			if method == "GET" && uri == "/cat.jpg" {
				fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")

				// ใส่ header เพื่อบอกให้ browser รู้ว่า นี่คือ image
				fmt.Fprint(conn, "Content-Type: image/jpeg\r\n\r\n")

				// เปิดไฟล์ cat.jpg
				f, err := os.Open("cat.jpg")
				if err != nil {
					log.Println(err)
				}

				// เอา content ในไฟล์ เขียนลงไปที่ connection
				io.Copy(conn, f)

				// Close file เพื่อป้องกัน memory leak
				f.Close()
			}
			if method == "GET" && uri == "/" {
				fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
				fmt.Fprint(conn, "Content-Type: text/html\r\n\r\n")
				fmt.Fprint(conn, `<img src="/cat.jpg">`)
			}

			conn.Close()
		}(conn)
	}
}
