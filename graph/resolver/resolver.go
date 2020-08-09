package resolver

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
	"github.com/PatrickKvartsaniy/image-processing-service/graph/validator"
	"github.com/PatrickKvartsaniy/image-processing-service/model"
	"io"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type (
	Resolver struct {
		storage   Storage
		processor Processor
		repo      Repository
		validator Validator
	}
	Repository interface {
		GetImage(ctx context.Context, id string) (*model.Image, error)
		GetImageVariety(ctx context.Context, id string) ([]*model.Resized, error)
		GetMultipleImages(ctx context.Context, limit, offset int) ([]*model.Image, error)
		SaveImage(ctx context.Context, version int, image model.Image) error
	}

	Storage interface {
		Read(ctx context.Context, path string) (io.Reader, error)
		Upload(ctx context.Context, data io.Reader) (string, error)
	}

	Processor interface {
		Resize(ctx context.Context, data io.Reader, output io.Writer, width, height int) error
	}

	Validator interface {
		ValidateFile(in graphql.Upload) error
		ValidateInput(input model.SizeInput) error
	}
)

func NewGraphqlResolver(s Storage, p Processor, r Repository) *Resolver {
	return &Resolver{
		storage:   s,
		processor: p,
		repo:      r,
		validator: validator.NewFileValidator(),
	}
}
