package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/PatrickKvartsaniy/image-processing-service/graph/generated"
	"github.com/PatrickKvartsaniy/image-processing-service/graph/model"
)

func (r *mutationResolver) UploadImage(ctx context.Context, image graphql.Upload, patameters []*model.SizeInput) (*model.Image, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ResizeImage(ctx context.Context, id string, patameters []*model.SizeInput) (*model.Image, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Images(ctx context.Context, limit int, offset int) ([]*model.Image, error) {
	panic(fmt.Errorf("not implemented"))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

// !!! WARNING !!!
// The code below was going to be deleted when updating resolvers. It has been copied here so you have
// one last chance to move it out of harms way if you want. There are two reasons this happens:
//  - When renaming or deleting a resolver the old code will be put in here. You can safely delete
//    it when you're done.
//  - You have helper methods in this file. Move them out to keep these resolver files clean.
func (r *mutationResolver) Upload(ctx context.Context, image graphql.Upload, patameters []*model.SizeInput) (*model.Image, error) {
	panic(fmt.Errorf("not implemented"))
}
func (r *mutationResolver) Resize(ctx context.Context, imageID string, patameters []*model.SizeInput) (*model.Image, error) {
	panic(fmt.Errorf("not implemented"))
}
