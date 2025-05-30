# 📚 Library Management API

This is a simple RESTful API built in Go using the [Fiber](https://gofiber.io/) framework and [GORM](https://gorm.io/) ORM. It provides basic CRUD operations to manage books in a library system.

## 🚀 Features

- Add new books
- View all books
- View a single book by ID
- Update book details
- Delete books
- Auto-migration with GORM
- JSON API responses

## 🧱 Technologies Used

- Go (Golang)
- Fiber (HTTP web framework)
- GORM (ORM for database interaction)
- MySQL or SQLite (Database)

## 🛠️ Setup Instructions

### 1. Clone the repository

```bash
git clone https://github.com/Mmah-Zombo/Library-Management-API.git
cd Library-Management-API
```

### 2. Configure the database

Edit your database connection string in database/database.go:

```go
dsn := "username:password@tcp(127.0.0.1:3306)/library_db?parseTime=true"
```

You can also use SQLite by adjusting the gorm.Open(...) call accordingly.

### 3. Run the project

go mod tidy
go run main.go

The API will start at:
http://localhost:3000

## 📘 API Endpoints

Get all books

```GET /api/books```

Get a single book

```GET /api/books/:id```

Add a new book

```POST /api/books```

```http
Content-Type: application/json

{
  "title": "Book Title",
  "author": "Author Name",
  "publish_date": "2024-05-01T00:00:00Z"
}
```

Update a book

```PUT /api/books/:id```

```http
Content-Type: application/json

{
  "title": "Updated Title",
  "author": "Updated Author",
  "publish_date": "2025-01-01T00:00:00Z"
}
```

Delete a book

```DELETE /api/books/:id```

## 🧪 Sample Test with cURL

```bash
curl -X POST http://localhost:3000/api/books \
-H "Content-Type: application/json" \
-d '{"title":"Clean Code", "author":"Robert Martin", "publish_date":"2023-01-01T00:00:00Z"}'
```

## 🤝 Contributing

Pull requests are welcome! For major changes, please open an issue first to discuss what you’d like to change.

## 📄 License

MIT License

⸻

Built with ❤️ in Go
