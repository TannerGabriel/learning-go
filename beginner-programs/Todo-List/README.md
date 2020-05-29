# Todo App

A golang todo list application using MongoDB as its database and Vue.js for the frontend.

## Requirements

- Local golang installation
- Nodejs and npm
- MongoDB

## Getting started

### Backend

Installing the dependencies:

```bash
go get github.com/gorilla/mux
go get go.mongodb.org/mongo-driver/bson
go get go.mongodb.org/mongo-driver/mongo
go get github.com/joho/godotenv
```

Running the application:

```bash
go run main.go
```

### Frontend

Installing the dependencies:

```bash
npm install
```

Running the application

```bash
npm run serve
```

The application can then be accessed using the browser on http://localhost:8080/.

### Using Docker

The external IP is required for the website to access the API from other PCs. For that, I included a shell script that automatically gets your IP Address on Linux, or lets you add a custom one if you are on Windows.

Start the application

```bash
# Building the application using the external IP as an Argument
./build-docker.sh

# Running the application in Docker
docker-compose up
```