1. ลองรันโปรแกรมผ่านทั้งทาง nc และ browser
  1.1 ลองเปิดหลาย browser ดู
  1.2 ลองเอา go func() ออก และลองเปิดหลาย browser ดู
2. ให้ลองเปลี่ยน content type จาก text/html เป็น text/plain
3. ลองเอา conn.Close() ออก แล้วลอง run ดู จะเห็นว่า browser ไม่รอเพราะเนื่องจากรู้ content length
4. ให้ลองเอา content length ออก และ ลองเพิ่ม +1 หลัง len(body) จะเห็นว่า browser จะยังรอการตอบจาก server อยู่ เนื่องจาก
    browser ไม่รู้ว่า content length มีเท่าไหร่ หรือ data ที่ได้จาก server ยังมาไม่ครบ