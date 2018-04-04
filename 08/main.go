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
			// ประกาศตัวแปรเพื่อเก็บ content ในแต่ละ route
			var content string

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
						content = "Hello World"
					}
					if method == "GET" && uri == "/page1" {
						content = "Page 1"
					}
					if method == "GET" && uri == "/page2" {
						content = "Page 2"
					}

					// method สามารถเป็นอะไรก็ได้ แต่ควรจะทำตาม convention
					// สามารถทดลองยิงด้วย nc ได้
					if method == "WHATEVER" && uri == "/" {
						content = "This is Easter Egg"
					}
				}
				if ln == "" {
					break
				}
				i++
			}

			// ใช้ตัวแปร content แทน fix text
			body := fmt.Sprintf(`<!DOCTYPE html>
				<html lang="en">
					<head>
						<meta charet="UTF-8">
						<title></title>
					</head>
					<body>
						<strong>%s</strong>
					</body>
				</html>`, content)
			fmt.Fprint(conn, "HTTP/1.1 200 OK\r\n")
			fmt.Fprintf(conn, "Content-Length: %d\r\n", len(body))
			fmt.Fprint(conn, "Content-Type: text/html\r\n")
			fmt.Fprint(conn, "\r\n")
			fmt.Fprint(conn, body)
			conn.Close()
		}(conn)
	}
}
