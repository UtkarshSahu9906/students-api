# 🎓 Students API

A beginner-friendly REST API built in Go for managing student records. Uses SQLite for storage, structured logging, and clean project layout.

---

## 📁 Project Structure

```
students-api/
├── cmd/
│   └── students-api/
│       └── main.go              # Entry point — starts the server
├── config/
│   └── local.yaml               # Local config (port, DB path, env)
├── internal/
│   ├── config/
│   │   └── config.go            # Loads config from YAML or env vars
│   ├── http/
│   │   └── handlers/
│   │       └── student/
│   │           └── student.go   # HTTP handlers (Create, GetByID, GetList)
│   ├── storage/
│   │   ├── storage.go           # Storage interface (abstraction)
│   │   └── sqlite/
│   │       └── sqlite.go        # SQLite implementation
│   ├── types/
│   │   └── types.go             # Student struct definition
│   └── utils/
│       └── response/
│           └── response.go      # JSON response helpers
├── storage/
│   └── storage.db               # SQLite database file (auto-created)
├── go.mod
└── go.sum
```

---

## 🚀 Getting Started

### Prerequisites

- [Go 1.21+](https://go.dev/dl/) installed
- GCC (required for `go-sqlite3` — a CGo package)
  - Windows: Install [TDM-GCC](https://jmeubank.github.io/tdm-gcc/)
  - Linux: `sudo apt install gcc`
  - Mac: `xcode-select --install`

### Run the Server

```bash
# 1. Clone the repo
git clone https://github.com/UtkarshSahu9906/students-api.git
cd students-api

# 2. Download dependencies
go mod tidy

# 3. Run the server
go run ./cmd/students-api --config config/local.yaml
```

The server will start at `http://localhost:8082`

---

## ⚙️ Configuration

Config is loaded from a YAML file (or environment variables).

**`config/local.yaml`**
```yaml
env: "dev"
storage_path: "storage/storage.db"
http_server:
  address: "localhost:8082"
```

| Field          | Description                        | Env Variable         |
|----------------|------------------------------------|----------------------|
| `env`          | Environment name (dev/prod)        | `ENV`                |
| `storage_path` | Path to SQLite database file       | —                    |
| `http_server.address` | Host and port for the server | `HTTP_SERVER_ADDR`   |

You can also set `CONFIG_PATH` environment variable to point to your config file instead of using the `--config` flag.

---

## 📡 API Endpoints

### Create a Student
```
POST /api/students
```
**Request Body (JSON):**
```json
{
  "name": "Utkarsh Sahu",
  "email": "utkarsh@example.com",
  "age": 21
}
```
**Response:**
```json
{ "id": 1 }
```

---

### Get a Student by ID
```
GET /api/students/{id}
```
**Response:**
```json
{
  "id": 1,
  "name": "Utkarsh Sahu",
  "email": "utkarsh@example.com",
  "age": 21
}
```

---

### Get All Students
```
GET /api/students
```
**Response:**
```json
[
  {
    "id": 1,
    "name": "Utkarsh Sahu",
    "email": "utkarsh@example.com",
    "age": 21
  }
]
```

---

## 🧱 How It Works

### 1. Config (`internal/config/config.go`)
Reads the YAML config file using the `cleanenv` library. If `CONFIG_PATH` env var is set, it uses that path; otherwise it falls back to the `--config` command-line flag.

### 2. Storage Interface (`internal/storage/storage.go`)
Defines a `Storage` interface with three methods:
```go
type Storage interface {
    CreateStudent(name, email string, age int) (int64, error)
    GetStudentById(id int64) (types.Student, error)
    GetStudents() ([]types.Student, error)
}
```
This makes the code flexible — you could swap SQLite for PostgreSQL without changing the handlers.

### 3. SQLite Implementation (`internal/storage/sqlite/sqlite.go`)
Implements the `Storage` interface using SQLite. On startup, it auto-creates the `students` table if it doesn't exist.

### 4. Handlers (`internal/http/handlers/student/student.go`)
Each handler is a function that returns an `http.HandlerFunc` — a common Go pattern for injecting dependencies like storage.

### 5. Graceful Shutdown (`cmd/students-api/main.go`)
The server listens for OS signals (`SIGINT`, `SIGTERM`) and shuts down cleanly with a 5-second timeout, allowing in-flight requests to finish.

---

## 📦 Dependencies

| Package | Purpose |
|---|---|
| `go-sqlite3` | SQLite database driver (CGo-based) |
| `cleanenv` | Reads config from YAML + env vars |
| `validator/v10` | Validates struct fields (e.g. `required`) |

---

## 🐛 Known Issues / TODOs

- [ ] `PUT /api/students/{id}` — Update student (commented out, not implemented)
- [ ] `DELETE /api/students/{id}` — Delete student (commented out, not implemented)
- [ ] Validation error in `New` handler does not `return` after writing the error — this can cause a double-response bug
- [ ] No 404 handling when student is not found (returns 500 instead)
- [ ] Tests are not yet written

---

## 💡 What You Can Learn From This Project

- Structuring a Go project with `cmd/` and `internal/` directories
- Using interfaces for clean, testable code
- Reading config from YAML files
- Building a REST API using only Go's standard `net/http` package
- Working with SQLite in Go using prepared statements
- Structured logging with `log/slog` (Go 1.21+)
- Graceful HTTP server shutdown

---

## 🙋 Author

**Utkarsh Sahu** — [GitHub](https://github.com/UtkarshSahu9906)
