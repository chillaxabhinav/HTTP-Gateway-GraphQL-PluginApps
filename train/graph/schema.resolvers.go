package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"gatewayCore/train/graph/generated"
	"gatewayCore/train/graph/model"
)

func (r *queryResolver) GetName(ctx context.Context) (*model.Name, error) {
	myName := model.Name{
		FirstName: "Abhinav",
		LastName:  "Singh",
	}

	return &myName, nil
}

func (r *queryResolver) GetTrain(ctx context.Context) (*model.Train, error) {
	myTrain := model.Train{
		Name: "Exp",
		Num:  2429,
	}

	return &myTrain, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
