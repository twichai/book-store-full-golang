# Book Store API

A simple RESTful Book Store API built with **Go (Golang)**.

This project includes:

* **Goose**: for database migrations
* **Air**: for live reload during development
* **Swagger**: for interactive API documentation

---

## Requirements

* Go 1.20 or later
* `make` (optional, for simplified commands)

---

## Setup Instructions

### 1. Clone the Repository

```bash
git clone git@github.com:twichai/book-store-full-golang.git
cd book-store-full-golang
```

### 2. Install Dependencies

```bash
go mod tidy
```

```bash
go install github.com/swaggo/swag/cmd/swag@latest
go install github.com/air-verse/air@latest
go install github.com/pressly/goose/v3/cmd/goose@latest
```

### 3. Configure Environment Variables

Create a `.env` file in the root folder:

```env
DB_DRIVER=postgres
DB_SOURCE=host=localhost user=postgres password=yourpass dbname=bookstore port=5432 sslmode=disable
PORT=8080
```

### 4. Run Database Migrations (Goose)

Install Goose:

```bash
go install github.com/pressly/goose/v3/cmd/goose@latest
```

Run migrations:

```bash
goose -dir ./migrations postgres "$DB_SOURCE" up
```

To create a new migration:

```bash
goose create create_books_table sql
```

---

## Development with Air

Install Air:

```bash
go install github.com/cosmtrek/air@latest
```

Start development server with hot reload:

```bash
air
```

---

## API Documentation with Swagger

Install Swag CLI:

```bash
go install github.com/swaggo/swag/cmd/swag@latest
```

Generate Swagger docs:

```bash
swag init
```

Then visit:

```
http://localhost:8080/swagger/index.html
```

---


## Project Structure

```
book-store/
├── cmd/                # Main application entry point
├── configs/            # Configuration files
├── internal/
│   ├── controller/     # Delivery layer (HTTP handlers)
│   ├── domain/         # Domain entities and interfaces
│   ├── repository/     # Implementation of repositories
│   └── usecase/        # Business logic
├── migrations/         # SQL migration files
├── docs/               # Swagger documentation
├── .env                # Environment variables
├── main.go             # App bootstrap
├── go.mod / go.sum
```

---

## License

MIT License

---

## Author

* Your Name [@twichai](https://github.com/twichai)
