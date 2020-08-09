package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"bytes"
	"context"
	"fmt"
	"time"

	uuid "github.com/satori/go.uuid"

	"github.com/99designs/gqlgen/graphql"
	"github.com/PatrickKvartsaniy/image-processing-service/graph/generated"
	"github.com/PatrickKvartsaniy/image-processing-service/model"
)

func (r *mutationResolver) Upload(ctx context.Context, image graphql.Upload, parameters model.SizeInput) (*model.Image, error) {
	if err := r.validator.ValidateFile(image); err != nil{
		return nil, fmt.Errorf("validating file: %w", err)
	}

	if err := r.validator.ValidateInput(parameters); err != nil {
		return nil, fmt.Errorf("validating input parameters: %w", err)
	}

	cp, orig := copyReader(image.File)
	path, err := r.storage.Upload(ctx, cp)
	if err != nil{
		return nil, fmt.Errorf("uploading original image: %w", err)
	}

	var resized bytes.Buffer
	if err = r.processor.Resize(ctx, orig, &resized, parameters); err != nil{
		return nil, fmt.Errorf("resizing image: %w", err)
	}

	resizedPath, err := r.storage.Upload(ctx, &resized)
	if err != nil{
		return nil, fmt.Errorf("uploading resized image: %w", err)
	}
	
	img := &model.Image{
		ID:      uuid.NewV4().String(),
		Path:    path,
		Type:    image.ContentType,
		Size:    int(image.Size),
		Ts:      time.Now(),
		Variety: []model.Resized{
			{
				Path: resizedPath,
				Width: parameters.Width,
				Height: parameters.Height,
			},
		},
	}

	if err := r.repo.SaveImage(ctx, *img); err != nil{
		return nil, fmt.Errorf("saving image: %w", err)
	}

	return img, nil
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
