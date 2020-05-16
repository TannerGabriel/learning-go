# gRPC Server-Client example

A simple gRPC example demonstration server to client communication.

## Requirements

- Protocol Buffers v3 installation

## Installation

The Protocol Buffer can be installed using the following command.

```bash
go get -u github.com/golang/protobuf/protoc-gen-go
```

You will have to ensure the protoc is added to your path environment variable so you can use it in your console.

## Getting started

Installing the dependencies:

```bash
go mod download
```

Starting the server:

```bash
go run server.go
```

Starting the client:

```bash
go run client.go
```