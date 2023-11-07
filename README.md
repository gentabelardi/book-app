Certainly! Here's a README.md template for a REST API CRUD Book Store project using Golang, Mux, GORM, and PostgreSQL:

# REST API CRUD Book Store

This is a sample project for a REST API that allows you to perform CRUD (Create, Read, Update, Delete) operations on book data using the Golang programming language, the Mux router, GORM ORM, and the PostgreSQL database.

## Features

- Add new books.
- List all books.
- Retrieve book details by ID.
- Update existing books.
- Delete books by ID.

## Prerequisites

Before you get started with this project, make sure you have the following prerequisites:

- Golang is installed on your computer.
- PostgreSQL is installed and running.
- You have a basic understanding of RESTful APIs, Golang, Mux, and GORM.

## Initial Setup

1. Clone this repository to your computer.

2. Create a PostgreSQL database for the project or create one using docker execute the following command:

```bash
docker compose up -d
```
Make sure you have updated the contents of the `compose.yml` field

## Running the Application

To run the application, execute the following command:

```bash
go run main.go
```

The application will run at `localhost:8080`. You can access the API endpoints using this address.

## Usage

Here are some example API requests you can make:

- **Create a New Book**:

  ```http
  POST /books
  Content-Type: application/json

  {
    "title": "Name book",
    "description": "Description book.",
    "genre": "Genre book",
    "author": "Author Name",
    "price": 250.00,
    "stock_quantity": 150,
    "published": "1999-10-10T00:00:00Z"
  }
  ```

- **List All Books**:

  ```http
  GET /books
  ```

- **Retrieve Book Details**:

  ```http
  GET /books/{id}
  ```

- **Update a Book**:

  ```http
  PUT /books/{id}
  Content-Type: application/json

  {
    "title": "Update title",
    "description": "Description book.",
    "genre": "Genre book",
    "author": "Author Name",
    "price": 250.00,
    "stock_quantity": 150,
    "published": "1999-10-10T00:00:00Z"
  }
  ```

- **Delete a Book**:

  ```http
  DELETE /books/{id}
  ```

Make sure to replace `{id}` with the actual book ID in your requests.

We hope this README helps you get started with the REST API CRUD Book Store project using Golang, Mux, GORM, and PostgreSQL. Happy coding!
