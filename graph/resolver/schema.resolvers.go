package resolver

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"bytes"
	"context"
	"fmt"
	"time"

	"github.com/99designs/gqlgen/graphql"
	"github.com/PatrickKvartsaniy/image-processing-service/graph/generated"
	"github.com/PatrickKvartsaniy/image-processing-service/model"
)

func (r *mutationResolver) Upload(ctx context.Context, image graphql.Upload, parameters model.SizeInput) (*model.Image, error) {
	if err := r.validator.ValidateFile(image); err != nil {
		return nil, fmt.Errorf("validating file: %w", err)
	}

	if err := r.validator.ValidateInput(parameters); err != nil {
		return nil, fmt.Errorf("validating input parameters: %w", err)
	}

	cp, orig := copyReader(image.File)
	path, err := r.storage.Upload(ctx, orig)
	if err != nil {
		return nil, fmt.Errorf("uploading original image: %w", err)
	}

	var resized bytes.Buffer
	if err = r.processor.Resize(ctx, cp, &resized, parameters); err != nil {
		return nil, fmt.Errorf("resizing image: %w", err)
	}

	resizedPath, err := r.storage.Upload(ctx, &resized)
	if err != nil {
		return nil, fmt.Errorf("uploading resized image: %w", err)
	}

	img := &model.Image{
		Path: path,
		Type: image.ContentType,
		Size: image.Size,
		Ts:   time.Now(),
	}
	img.NewSize(resizedPath, parameters)

	if err := r.repo.SaveImage(ctx, img); err != nil {
		return nil, fmt.Errorf("saving image: %w", err)
	}
	return img, nil
}

func (r *mutationResolver) Resize(ctx context.Context, id string, parameters model.SizeInput) (*model.Image, error) {
	if err := r.validator.ValidateInput(parameters); err != nil {
		return nil, fmt.Errorf("validating input parameters: %w", err)
	}

	image, err := r.repo.GetImage(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("getting image: %w", err)
	}

	orig, err := r.storage.Read(ctx, image.Path)
	if err != nil {
		return nil, fmt.Errorf("reading image from storage: %w", err)
	}

	var resized bytes.Buffer
	if err = r.processor.Resize(ctx, orig, &resized, parameters); err != nil {
		return nil, fmt.Errorf("resizing image: %w", err)
	}

	resizedPath, err := r.storage.Upload(ctx, &resized)
	if err != nil {
		return nil, fmt.Errorf("uploading resized image: %w", err)
	}

	image.NewSize(resizedPath, parameters)
	image.IncreaseVersion()

	if err = r.repo.UpdateImage(ctx, image); err != nil {
		return nil, fmt.Errorf("saving image: %w", err)
	}

	return image, nil
}

func (r *queryResolver) Images(ctx context.Context, limit int, offset int) ([]*model.Image, error) {
	return r.repo.GetMultipleImages(ctx, int64(limit), int64(offset))
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
