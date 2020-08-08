package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"

	"github.com/99designs/gqlgen/graphql"
	"github.com/PatrickKvartsaniy/image-processing-service/graph/generated"
	"github.com/PatrickKvartsaniy/image-processing-service/model"
)

func (r *imageResolver) Variety(ctx context.Context, obj *model.Image) ([]*model.Resized, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) UploadImage(ctx context.Context, image graphql.Upload, parameters []*model.SizeInput) (*model.Image, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *mutationResolver) ResizeImage(ctx context.Context, id string, parameters []*model.SizeInput) (*model.Image, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *queryResolver) Images(ctx context.Context, limit int, offset int) ([]*model.Image, error) {
	panic(fmt.Errorf("not implemented"))
}

// Image returns generated1.ImageResolver implementation.
func (r *Resolver) Image() generated.ImageResolver { return &imageResolver{r} }

// Mutation returns generated1.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated1.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type imageResolver struct{ *Resolver }
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
