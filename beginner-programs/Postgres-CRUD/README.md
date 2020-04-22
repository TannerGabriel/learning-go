# Postgres CRUD

This is an example project that shows the basics of working with the PostgresSQL database in golang. Go provides an abstract datalayer using the database/sql standard library.

## Prerequisites

* Local Postgres instance
* Valid Golang installation

## Getting started

First you need to install the Postgres Go drivers.

```bash
go get -u github.com/lib/pq
```

You can then start the application using go run.

```bash
go run postgres.go
```

## Implementation

* [Postrgres CRUD](https://github.com/TannerGabriel/learning-go/blob/master/beginner-programs/Postgres-CRUD/postgres.go)