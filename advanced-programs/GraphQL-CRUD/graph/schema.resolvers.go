package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/tannergabriel/learning-go/advanced-programs/GraphQL-CRUD/graph/generated"
	"github.com/tannergabriel/learning-go/advanced-programs/GraphQL-CRUD/graph/model"
)

func (r *mutationResolver) CreateItems(ctx context.Context, input model.NewItem) (*model.Item, error) {
	item := &model.Item{
		Title:  input.Title,
		Owner:  input.Owner,
		Rating: input.Rating,
	}
	r.items = append(r.items, item)
	return item, nil
}

func (r *queryResolver) Items(ctx context.Context) ([]*model.Item, error) {
	return r.items, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
