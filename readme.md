# Project Ramadhan TODO LIST 

Project ini dibuat sebagai bagian sharing saya dengan tajuk tantangan belajar coding selama Ramadhan untuk mempelajari microservices, pemanfaatan redis ttl, dan event carrier state transfer.
Aplikasi ini adalah **To do List** berbasis Golang dengan **PostgreSQL, Redis, dan NATS**, yang memungkinkan pengguna untuk mencatat tugas, memantau progress, dan mengatur waktu kadaluarsa tugas secara otomatis.
Pada Service ini, ini service yang akan listen ke subject yang dikirim secara event carrier state transfer menggunakan NATS (implementasi event driven).

## **Tech Stack yang Digunakan**  
- **Golang** → Bahasa pemrograman utama  
- **PostgreSQL** → Database utama untuk menyimpan data tugas  
- **Redis** → Digunakan untuk caching dan TTL tugas  
- **NATS** → Event-driven system untuk komunikasi antar service  
- **JWT** → Digunakan untuk autentikasi pengguna  
- **Docker** → Untuk menjalankan layanan dengan lebih mudah  