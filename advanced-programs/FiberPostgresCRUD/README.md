# Fiber Postgres CRUD

A simple PostgresSQL CRUD application using the Fiber webframework.

## Requirements

- Local golang installation
- PostgresSQL

## Getting started

### Without Docker

Installing the dependencies:

```bash
go mod download
```

Running the application:

```bash
go run main.go
```

## With Docker

Running the application

```bash
docker-compose up
```

After starting the application the server should be started on port 3000 and you can start sending requests to the endpoints.

```bash
# Endpoints
"/api/v1/item"
"/api/v1/item/:id"
"/api/v1/item"
"/api/v1/item/:id"
```