# Go REST API Starter 🚀

Template REST API standar kami menggunakan [Go Fiber](https://gofiber.io) dan MongoDB.  
Didesain agar developer dapat membuat backend dengan struktur modular, bersih, dan siap dikembangkan ke skala besar.

---

## 🔧 Fitur
- Go Fiber web framework
- MongoDB integration
- Clean architecture (handler → service → repository)
- Environment variable support
- Simple starter for production-ready APIs

---

## ⚙️ Cara Menjalankan
1. Clone repository: 
- git clone https://github.com/withrizky/go-restapi-mongodb-starter.git
- cd go-restapi-mongodb-starter
2. Install dependencies:
- go mod tidy
3. Buat file .env berdasarkan .env.example
4. Jalankan server : go run cmd/server/main.go
5. Endpoint example : 
- POST http://localhost:8080/api/users
- body :
{
  "name": "Rizky",
  "email": "rizky@example.com"
}
- respon diharapkan :
{
  "message": "user created successfully"
}