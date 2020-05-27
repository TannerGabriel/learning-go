package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"

	"github.com/tannergabriel/learning-go/advanced-programs/GraphQL-CRUD/database"
	"github.com/tannergabriel/learning-go/advanced-programs/GraphQL-CRUD/graph/generated"
	"github.com/tannergabriel/learning-go/advanced-programs/GraphQL-CRUD/graph/model"
)

func (r *mutationResolver) CreateItems(ctx context.Context, input model.NewItem) (*model.Item, error) {
	item := &model.Item{
		Title:  input.Title,
		Owner:  input.Owner,
		Rating: input.Rating,
	}

	db := database.DBConn

	db.Create(&item)

	return item, nil
}

func (r *mutationResolver) Delete(ctx context.Context, id int) (string, error) {
	db := database.DBConn

	var item model.Item
	db.First(&item, id)
	if item.Title == "" {
		return "No item found with given ID", nil
	}

	db.Delete(&item)
	return "Item successfully deleted", nil
}

func (r *queryResolver) Items(ctx context.Context) ([]*model.Item, error) {
	db := database.DBConn
	var items []*model.Item
	db.Find(&items)
	return items, nil
}

func (r *queryResolver) Item(ctx context.Context, id int) (*model.Item, error) {
	db := database.DBConn
	var item model.Item
	db.Find(&item, id)

	returnItem := &model.Item{
		ID:     item.ID,
		Title:  item.Title,
		Owner:  item.Owner,
		Rating: item.Rating,
	}

	return returnItem, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
