package main

import (
	"io"
	"net"
)

func main() {
	// Listen ที่ localhost port 8080
	li, err := net.Listen("tcp", ":8080")

	// ถ้าหากมี error ให้จบโปรแกรม
	if err != nil {
		panic(err)
	}

	// เพิ่ม infinite loop เพื่อให้สามารถรอรับ request ถัดไปได้
	for {
		// code จะหยุดรอที่บันทัดนี้ จนกว่าจะมี client request เข้ามา
		conn, err := li.Accept()

		// ถ้าหากมี error ให้จบโปรแกรม
		if err != nil {
			panic(err)
		}

		// write string กลับไปที่ connection ที่ request เข้ามา
		io.WriteString(conn, "This message from TCP server")
	}
}
