# Websockets Chat

A simple Websockets realtime chat application.

## Requirements

- Local golang installation

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

Building the image:

```bash
docker build -t websocketschat .
```

Running the application:

```bash
docker run -p 5000:5000 websocketschat
```

After starting the application the server should be started on port 5000 and you can start sending and receiving chat messages.