# Go + Gin Web API

Proyek ini adalah contoh web API menggunakan Golang dengan framework [Gin](https://github.com/gin-gonic/gin). Di dalamnya terdapat middleware untuk logging request **dengan timezone WIB (Asia/Jakarta)** dan **request ID** per request.

## ✨ Fitur Utama

- Middleware `RequestID`: menghasilkan UUID unik untuk setiap request dan menambahkannya ke header `X-Request-ID`.
- Middleware `LogStartEnd`: menampilkan log ketika request masuk dan selesai diproses.
- Log disesuaikan ke zona waktu **WIB (+07:00)** meskipun server berada di zona waktu lain.
- Output log bersih, tanpa duplikasi timestamp.

## 📁 Struktur Direktori