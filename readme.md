# 🚀 Belajar Go REST API (Gin + GORM)

Project ini adalah backend REST API sederhana menggunakan **Golang**, dengan arsitektur scalable berbasis:

* Gin (HTTP framework)
* GORM (ORM)
* PostgreSQL (database)
* Clean Architecture (Controller → Service → Repository)

---

## 📂 Struktur Project

```
├── config/         # Konfigurasi (DB connection, env, dll)
├── controllers/    # Handler / controller (entry point request)
├── dto/            # Data Transfer Object (request & response)
├── middleware/     # Middleware (auth, logging, dll)
├── migrations/     # Migration database
├── models/         # Struct model (representasi tabel)
├── providers/      # Dependency injection (container, provider)
├── repositories/   # Layer akses database
├── routes/         # Routing endpoint
├── services/       # Business logic
├── tmp/            # Build output (air / binary)
├── .air.toml       # Config hot reload (Air)
├── .env            # Environment variables
├── go.mod
├── go.sum
└── main.go         # Entry point aplikasi
```

---

## ⚙️ Teknologi yang Digunakan

* Golang
* Gin Gonic
* GORM
* PostgreSQL
* Air (hot reload)

---

## 🔄 Arsitektur

```
Request → Controller → Service → Repository → Database
```

### 📌 Penjelasan

* **Controller**: handle HTTP request & response
* **Service**: business logic
* **Repository**: query database
* **DTO**: validasi & format request

---

## 🧪 Endpoint API

### 👤 Users

| Method | Endpoint   | Deskripsi      |
| ------ | ---------- | -------------- |
| GET    | /users     | Get all users  |
| GET    | /users/:id | Get user by ID |
| POST   | /users     | Create user    |
| PATCH  | /users/:id | Update user    |
| DELETE | /users/:id | Delete user    |

---

## 📥 Contoh Request

### Create User

```json
POST /users

{
  "name": "John Doe",
  "email": "john@mail.com"
}
```

### Update User (Partial)

```json
PATCH /users/:id

{
  "name": "Updated Name"
}
```

---

## 🛠️ Cara Menjalankan Project

### 1. Clone repository

```bash
git clone https://github.com/helkass/belajar-go.git
cd belajar-go
```

---

### 2. Install dependency

```bash
go mod tidy
```

---

### 3. Setup environment

Buat file `.env`:

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASS=postgres
DB_NAME=belajar
```

---

### 4. Jalankan aplikasi

```bash
go run main.go
```

---

### 5. Jalankan dengan Hot Reload

Menggunakan Air:

```bash
air
```

---

## 🔐 UUID sebagai Primary Key

Model menggunakan UUID:

```go
ID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
```

Pastikan extension di PostgreSQL sudah aktif:

```sql
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";
```

---

## ⚡ Best Practice yang Digunakan

* Separation of concerns (Controller, Service, Repository)
* DTO untuk validasi request
* Dependency Injection (Provider & Container)
* UUID sebagai primary key
* Clean & scalable structure

---

## 🧹 .gitignore yang Disarankan

```
# binary
/tmp
*.exe
*.out

# env
.env

# logs
*.log

# air
tmp/

# OS
.DS_Store
```

---

## 🚀 Catatan

* Gunakan **query params** untuk filter
  `/users?name=andi`

* Gunakan **path param** untuk ID
  `/users/:id`

* Gunakan **PATCH** untuk partial update

---

## 🤝 Kontribusi

Pull request terbuka untuk improvement 🚀

---

## 📄 License

MIT License
