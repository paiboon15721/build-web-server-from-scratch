1. ลองรันโปรแกรมผ่านทั้งทาง nc และ browser
  1.1 ลองเปิดหลาย browser ดู
  1.2 ลองเอา go func() ออก และลองเปิดหลาย browser ดู
2. ให้ลองเปลี่ยน content type จาก text/html เป็น text/plain
3. ลองเอา conn.Close() ออก แล้วลอง run ดู จะเห็นว่า browser ไม่รอเพราะเนื่องจากรู้ content length
4. ให้ลองเอา content length ออก และ ลองเพิ่ม +1 หลัง len(body) จะเห็นว่า browser จะยังรอการตอบจาก server อยู่ เนื่องจาก
    browser ไม่รู้ว่า content length มีเท่าไหร่ หรือ data ที่ได้จาก server ยังมาไม่ครบ

5. ลองเล่นกับ charset โดยการลองสลับ header และ byte slice ไปมา

fmt.Fprint(conn, "Content-Type: text/plain; charset=utf-8\r\n")
fmt.Fprint(conn, "Content-Type: text/plain; charset=tis-620\r\n")

** utf-8 byte slice
body := string([]byte{224, 184, 151, 224, 184, 148, 224, 184, 170, 224, 184, 173, 224, 184, 154, 224, 184, 160, 224, 184, 178, 224, 184, 169, 224, 184, 178, 224, 185, 132, 224, 184, 151, 224, 184, 162})

** tis-620 byte slice
body := string([]byte{183, 180, 202, 205, 186, 192, 210, 201, 210, 228, 183, 194, 32})

*** จะเห็นว่า utf-8 byte slice จะมีความยาวมากกว่า tis-620 เนื่องจากใช้พื้นที่ในการเก็บแต่ละอักษรที่ 3 byte เนื่องจากรองรับภาษาที่มากกว่านั่นเอง