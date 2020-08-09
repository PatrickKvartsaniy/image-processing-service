package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/PatrickKvartsaniy/image-processing-service/graph/generated"
	"github.com/PatrickKvartsaniy/image-processing-service/model"
)

func (r *mutationResolver) Upload(ctx context.Context, image graphql.Upload, parameters model.SizeInput) (*model.Image, error) {
	if err := r.validator.ValidateFile(image); err != nil {
		return nil, err
	}
	if err := r.validator.ValidateInput(parameters); err != nil {
		return nil, err
	}

	return nil, nil
}

func (r *mutationResolver) Resize(ctx context.Context, id string, parameters model.SizeInput) (*model.Image, error) {
	if err := r.validator.ValidateInput(parameters); err != nil {
		return nil, err
	}
	return nil, nil
}

func (r *queryResolver) Images(ctx context.Context, limit int, offset int) ([]*model.Image, error) {
	return r.repo.GetMultipleImages(ctx, limit, offset)
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
