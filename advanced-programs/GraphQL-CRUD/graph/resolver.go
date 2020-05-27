package graph

//go:generate go run github.com/99designs/gqlgen

import "github.com/tannergabriel/advanced-programs/GraphQL-CRUD/graph/model"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	todos []*model.Todo
}
