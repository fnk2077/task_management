# วิธีการตั้งค่าโปรเจกต์

โปรเจกต์นี้ประกอบด้วยสองส่วนหลักคือ Backend (Go) และ Frontend (Vue.js) ซึ่งสามารถตั้งค่าและรันได้ตามขั้นตอนดังนี้:

1.เปิดเทอร์มินัลและไปที่โฟลเดอร์โปรเจกต์
2.รันคำสั่งต่อไปนี้เพื่อเปิด Docker:
   ```bash
   docker compose up -d
   ```
3.รันโปรเจกต์ Go ด้วยคำสั่ง:
   ```bash
   go run main.go
   ```
4.เปิดเทอร์มินัลใหม่เพื่อรันโปรเจกต์ Vue.jst ด้วยคำสั่ง:
   ```bash
   cd vue-project
   npm install
   npm run dev
   ```
5.ไปที่ URL http://localhost:5173 เพื่อดูหน้าเว็บที่รันอยู่

รหัสสำหรับเริ่มต้น
```bash
email: a1@a.com
password: 1
```
```bash
email: test01@a.com
password: 1234
```
