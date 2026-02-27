# Go Server Exploration

Proyek ini merupakan eksperimen sederhana untuk mempelajari pengembangan _backend_ menggunakan bahasa pemrograman **Go (Golang)**. Fokus utama dari repositori ini adalah memahami cara kerja _HTTP Server_ dan manajemen _routing_.

## Struktur Proyek

Proyek ini mengikuti pola struktur standar Go untuk memisahkan logika aplikasi:

- `cmd/`: Direktori untuk aplikasi utama.
  - `api/`: Berisi _entry point_ untuk API server.
    - `main.go`: File utama untuk menjalankan server.

## Cara Menjalankan

Untuk menjalankan server, pastikan Anda sudah menginstal Go, lalu jalankan perintah berikut di terminal:

```bash
go run cmd/api/main.go
```
