package main

import (
	"io"
	"log"
	"net"
)

func main() {
	// Listen ที่ localhost port 8080
	li, err := net.Listen("tcp", ":8080")

	// ถ้าหากมี error ให้จบโปรแกรม
	if err != nil {
		panic(err)
	}

	for {
		// code จะหยุดรอที่บันทัดนี้ จนกว่าจะมี client request เข้ามา
		conn, err := li.Accept()

		// ถ้าหากมี error ให้จบโปรแกรม
		if err != nil {
			// เปลี่ยนเป็น log error ออกมาแทน เนื่องจากถ้าใช้ panic web server จะหยุดทำงานทันที
			// ซึ่งจะกระทบกับผู้ใช้งานทั้งหมด แม้จะ error แค่ request เดียว
			log.Println(err)
			continue
		}

		// write string กลับไปที่ connection ที่ request เข้ามา
		io.WriteString(conn, "This message from TCP server")

		// close connection เมื่อทำงานเสร็จ
		conn.Close()
	}
}
