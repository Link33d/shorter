# Link Shortener Service

This is a simple URL shortener service built using **Go**. It provides basic functionality for creating short links, redirecting to the original links, and hosting a main webpage.

This is my first project in Go, so I'm open to any constructive feedback or suggestions. 😊

---

## Features

- **POST `/`**: Creates a new shortened link.
- **GET `/`**: Serves the main web page.
- **GET `/<code>`**: Redirects to the original URL associated with the provided code.

---

## Prerequisites

### Verify Go Installation

Make sure you have **Go** installed on your system. You can verify this by running:

```bash
go version
```

If Go is not installed, follow the [official installation guide](https://go.dev/doc/install).

---

## Setting Up the Database (Optional)

If you do not already have a PostgreSQL database configured, you can set one up using Docker:

### 1. Start PostgreSQL with Docker

Run the following command to create a PostgreSQL container:

```bash
docker run --name postgres-shortener -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=password -e POSTGRES_DB=shorter -p 5432:5432 -d postgres
```

This will:
- Set the PostgreSQL username to `postgres`.
- Set the password to `password`.
- Create a database named `shorter`.
- Expose the database on port `5432`.

### 2. Update the Database Configuration

Ensure the file `config/ConnectDB.go` has the following settings:

```go
const (
    host     = "localhost" // Replace with your database host, e.g., localhost if Docker is local
    port     = "5432"
    user     = "postgres"
    password = "password"
    dbname   = "shorter"
)
```

If the database is not running on the same machine as the service, replace `localhost` in the `host` variable with the appropriate IP address of your database server.

---

## Running the Service

### 1. Install Dependencies

Run the following command to install all required Go modules:

```bash
go mod tidy
```

### 2. Start the Service

Run the service with:

```bash
go run main.go
```

The service will now be running and available at `http://localhost:8080` by default.

---

## Considerations

1. If you're running the service behind a proxy, ensure that ports are forwarded correctly.
2. This project is designed to be lightweight and simple, so additional features can be added as needed.
3. Constructive feedback and suggestions are highly appreciated!

Thank you for checking out my project! 🎉