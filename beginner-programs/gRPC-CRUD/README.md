# GRPC CRUD using Cobra CLI

MongoDB CRUD application showcasting how to use GRPC to communicate between a client and server. The client is build using the Cobra CLI library.

## Requirements

* MongoDB Database
* Golang installation

## Technology Stack

* GRPC as the communication protocol between the server and the client
* MongoDB as the database of the CRUD application
* Cobra for building a well structured and documented command line application

## Getting started

### Building the proto file

Building the proto file:

```bash
protoc pb/blog.proto --go_out=plugins=grpc:./pb
```

### Server

Starting the server:

```bash
go run server/server.go
```

### Client

Installing the CLI:

```bash
go install ./client/
```

Inserting a document:

```bash
client create --authorid "101010010" --content "Some great blog content" --title "Very great blog post"
```

Reading a document:

```bash
client read --blogid BLOG_ID
# For example: 
client read --blogid 5f01cf70960b42b7d962acf3
```

Updating a document:

```bash
client update --id 5f01cf70960b42b7d962acf3 --title "Some great title" --content "Some great content" --authorid "100203900392840932875"
```

Deleting a document:

```bash
client delete --blogid 5f01cf70960b42b7d962acf3
```