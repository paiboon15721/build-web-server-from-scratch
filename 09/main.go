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

		go func() {
			// ตัวแปร title เพื่อใช้แสดงว่ากำลังอยู่ที่ route ไหน
			var title string
			// ตัวแปร form ใช้สำหรับแสดง add form
			var form string

			i := 0
			scanner := bufio.NewScanner(conn)
			for scanner.Scan() {
				ln := scanner.Text()
				fmt.Println(ln)
				if i == 0 {
					words := strings.Fields(ln)
					method := words[0]
					uri := words[1]

					// router
					if method == "GET" && uri == "/" {
						title = "INDEX"
					}
					if method == "GET" && uri == "/profile" {
						title = "PROFILE"
					}
					if method == "GET" && uri == "/about" {
						title = "ABOUT"
					}
					if method == "GET" && uri == "/contact" {
						title = "CONTACT"
					}
					if method == "GET" && uri == "/add" {
						title = "ADD DATA"
						form = `<form method="POST" action="/add">
						<input type="submit" value="add">
						</form>`
					}
					if method == "POST" && uri == "/add" {
						title = "ADD DATA SUCCESS"
					}
				}
				if ln == "" {
					break
				}
				i++
			}

			body := fmt.Sprintf(`<!DOCTYPE html>
				<html lang="en">
					<head>
						<meta charet="UTF-8">
						<title></title>
					</head>
					<body>
						<strong>%s</strong><br>
						<a href="/">index</a><br>
						<a href="/profile">profile</a><br>
						<a href="/about">about</a><br>
						<a href="/contact">contact</a><br>
						<a href="/add">add</a><br>
						%s
					</body>
				</html>`, title, form)
			fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
			fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
			fmt.Fprint(conn, "Content-Type: text/html\r\n")
			fmt.Fprint(conn, "\r\n")
			fmt.Fprint(conn, body)
			conn.Close()
		}()
	}
}
