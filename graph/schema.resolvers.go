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

func (r *droidResolver) Friends(ctx context.Context, obj *model.Droid) ([]model.Character, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *droidResolver) FriendsConnection(ctx context.Context, obj *model.Droid, first *int, after *string) (*model.FriendsConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *friendsConnectionResolver) Edges(ctx context.Context, obj *model.FriendsConnection) ([]*model.FriendsEdge, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *friendsConnectionResolver) Friends(ctx context.Context, obj *model.FriendsConnection) ([]model.Character, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *humanResolver) Friends(ctx context.Context, obj *model.Human) ([]model.Character, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *humanResolver) FriendsConnection(ctx context.Context, obj *model.Human, first *int, after *string) (*model.FriendsConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *humanResolver) Starships(ctx context.Context, obj *model.Human) ([]*model.Starship, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) CreateReview(ctx context.Context, episode model.Episode, review model.Review) (*model.Review, error) {
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
	human, found := r.datasource.humans[id]
	if !found {
		return nil, fmt.Errorf("not found")
	}

	return &human, nil
}

func (r *queryResolver) Starship(ctx context.Context, id string) (*model.Starship, error) {
	starship, found := r.datasource.starships[id]
	if !found {
		return nil, fmt.Errorf("not found")
	}

	return &starship, nil
}

func (r *starshipResolver) Length(ctx context.Context, obj *model.Starship, unit *model.LengthUnit) (float64, error) {
	if unit != nil && *unit == model.LengthUnitFoot {
		return obj.Length * 3.28084, nil
	}

	return obj.Length, nil
}

// Droid returns generated.DroidResolver implementation.
func (r *Resolver) Droid() generated.DroidResolver { return &droidResolver{r} }

// FriendsConnection returns generated.FriendsConnectionResolver implementation.
func (r *Resolver) FriendsConnection() generated.FriendsConnectionResolver {
	return &friendsConnectionResolver{r}
}

// Human returns generated.HumanResolver implementation.
func (r *Resolver) Human() generated.HumanResolver { return &humanResolver{r} }

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

// Starship returns generated.StarshipResolver implementation.
func (r *Resolver) Starship() generated.StarshipResolver { return &starshipResolver{r} }

type droidResolver struct{ *Resolver }
type friendsConnectionResolver struct{ *Resolver }
type humanResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
type starshipResolver struct{ *Resolver }
