package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/pangaunn/gqlgen-workshop/graph/generated"
	"github.com/pangaunn/gqlgen-workshop/graph/model"
)

func (r *mutationResolver) CreateReview(ctx context.Context, episode model.Episode, review model.ReviewInput) (*model.Review, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Hero(ctx context.Context, episode *model.Episode) (model.Character, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Reviews(ctx context.Context, episode model.Episode, since *time.Time) ([]*model.Review, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Search(ctx context.Context, text string) ([]model.SearchResult, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Character(ctx context.Context, id string) (model.Character, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Droid(ctx context.Context, id string) (*model.Droid, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Human(ctx context.Context, id string) (*model.Human, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Starship(ctx context.Context, id string) (*model.Starship, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
