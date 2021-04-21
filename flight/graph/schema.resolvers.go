package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"gatewayCore/flight/graph/generated"
	"gatewayCore/flight/graph/model"
)

func (r *queryResolver) GetFlight(ctx context.Context) (*model.Flight, error) {
	myFlight := model.Flight{
		Name:        "SpiceJet",
		Source:      "Delhi",
		Destination: "Mumbai",
	}

	return &myFlight, nil
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type queryResolver struct{ *Resolver }
