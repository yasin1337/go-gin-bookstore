# Go Bookstore API

Small learning project written in Go using the Gin framework. It exposes a handful of read/write endpoints for managing an in-memory list of books while demonstrating JSON binding, validation, and structured HTTP responses.

## Features
- List all books or fetch a single title by ID.
- Add new books with simple validation for required fields.
- Consistent JSON response types for both success and error scenarios.
- Minimal, readable structure split across `main.go`, `api.go`, `types.go`, and `sampleData.go`.

## Requirements
- Go 1.21+ (earlier versions may work but were not tested).

## Getting Started
```bash
git clone current project
cd go-bookstore
go run .
```

The server starts on `localhost:1337` by default.

## API

| Method | Path          | Description               |
|--------|---------------|---------------------------|
| GET    | `/books`      | Return all books.         |
| GET    | `/books/:id`  | Return a single book.     |
| POST   | `/books`      | Create a new book entry.  |

### Sample Payload
```json
{
  "id": "4",
  "title": "The Obelisk Gate",
  "price": 39.99
}
```

### Responses
- `200 OK` for successful GET requests.
- `201 Created` when a book is added.
- `400 Bad Request` if required fields are missing or invalid.
- `404 Not Found` when a book ID does not exist.

## Next Steps
- Swap the in-memory slice for persistent storage.
- Add authentication and richer validation.
- Write handler tests using `net/http/httptest`.
# go-bookstore
