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

		go func() {
			// ประกาศตัวแปรตรงนี้ เนื่องจากนำไปใช้ข้างนอก for loop
			var method string
			var uri string
			var f *os.File

			i := 0
			scanner := bufio.NewScanner(conn)
			for scanner.Scan() {
				ln := scanner.Text()
				fmt.Println(ln)

				if i == 0 {
					words := strings.Fields(ln)
					method = words[0]
					uri = words[1]

					// router
					if method == "GET" && uri == "/cat.jpg" {
						f, _ = os.Open("cat.jpg")
					}
				}
				if ln == "" {
					break
				}
				i++
			}

			body := `<img src="/cat.jpg">`
			fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
			if method == "GET" && uri == "/cat.jpg" {
				// ใส่ header เพื่อบอกให้ browser รู้ว่า นี่คือ image
				fmt.Fprint(conn, "Content-Type: image/jpeg\r\n\r\n")

				// เอา content ในไฟล์ เขียนลงไปที่ connection
				io.Copy(conn, f)
			} else {
				fmt.Fprint(conn, "Content-Type: text/html\r\n\r\n")
				fmt.Fprint(conn, body)
			}

			// Close file เพื่อป้องกัน memory leak
			f.Close()
			conn.Close()
		}()
	}
}
