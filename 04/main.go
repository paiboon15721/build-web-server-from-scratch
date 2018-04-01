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

		// scan text จาก client
		scanner := bufio.NewScanner(conn)

		// ถ้าเจอการขึ้นบันทัดใหม่ ให้พิมพ์ text ออกที่ console ของ server แล้ว loop ขึ้นมา scan ใหม่
		for scanner.Scan() {
			ln := scanner.Text()
			fmt.Println(ln)
		}
		conn.Close()
	}
}
