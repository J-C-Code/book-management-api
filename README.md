# Book Management API

This project is a simple RESTful API for managing books using the Gin framework in Go. It allows users to retrieve a list of books, get details of a specific book, create new books, check out books, and return books.

## Features

- **Get all books**: Retrieve a list of all books in JSON format.
- **Get book by ID**: Fetch details of a specific book by its ID.
- **Create a new book**: Add a new book to the collection.
- **Checkout a book**: Check out a book, reducing its quantity.
- **Return a book**: Return a book, increasing its quantity.

## Installation

### Prerequisites

Make sure you have Go installed on your machine. You can download it from [golang.org](https://golang.org/dl/).

### Clone the Repository

```bash
git clone https://github.com/Jomma52637/book-management-api.git
cd book-management-api
```

### Install Dependencies

Make sure to have the Gin framework installed. You can do this by running:

```bash
go get -u github.com/gin-gonic/gin
```

### Running the API

You can run the API by executing:

```bash
go run main.go
```

The server will start at `http://localhost:8080`.

## API Endpoints

### Get All Books

- **Endpoint**: `/books`
- **Method**: `GET`
- **Response**: List of all books in JSON format.

**Example Response**:
```json
[
    {
        "id": "1",
        "title": "Harry Potter and the Sorceror's Stone",
        "author": "J.K Rowling",
        "qty": 2
    },
    {
        "id": "2",
        "title": "Harry Potter and the Order of the Phoenix",
        "author": "J.K Rowling",
        "qty": 3
    },
    {
        "id": "3",
        "title": "Harry Potter and the Deathly Hallows",
        "author": "J.K Rowling",
        "qty": 5
    }
]
```

### Get Book by ID

- **Endpoint**: `/books/:id`
- **Method**: `GET`
- **Response**: Details of the book with the specified ID.

**Example Request**:
```
GET /books/1
```

**Example Response**:
```json
{
    "id": "1",
    "title": "Harry Potter and the Sorceror's Apprentice",
    "author": "J.K Rowling",
    "qty": 2
}
```

### Create a New Book

- **Endpoint**: `/books`
- **Method**: `POST`
- **Request Body**: JSON object representing the new book.

**Example Request Body**:
```json
{
    "id": "4",
    "title": "Harry Potter and the Goblet of Fire",
    "author": "J.K Rowling",
    "qty": 4
}
```

**Example Response**:
```json
{
    "id": "4",
    "title": "Harry Potter and the Goblet of Fire",
    "author": "J.K Rowling",
    "qty": 4
}
```

### Checkout a Book

- **Endpoint**: `/checkout`
- **Method**: `PATCH`
- **Query Parameter**: `id` (the ID of the book to check out).

**Example Request**:
```
PATCH /checkout?id=1
```

**Example Response**:
```json
{
    "id": "1",
    "title": "Harry Potter and the Sorceror's Apprentice",
    "author": "J.K Rowling",
    "qty": 1
}
```

### Return a Book

- **Endpoint**: `/checkin`
- **Method**: `PATCH`
- **Query Parameter**: `id` (the ID of the book to return).

**Example Request**:
```
PATCH /checkin?id=1
```

**Example Response**:
```json
{
    "message": "The book: Harry Potter and the Sorceror's Apprentice has been returned."
}
```

## Contributing

Contributions are welcome! Please feel free to submit a pull request.
