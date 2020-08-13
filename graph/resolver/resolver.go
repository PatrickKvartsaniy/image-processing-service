package resolver

import (
	"context"
	"github.com/99designs/gqlgen/graphql"
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
		GetMultipleImages(ctx context.Context, limit, offset int64) ([]*model.Image, error)
		SaveImage(ctx context.Context, image *model.Image) error
		UpdateImage(ctx context.Context, image *model.Image) error
	}
	Storage interface {
		Read(ctx context.Context, path string) (io.Reader, error)
		Upload(ctx context.Context, data io.Reader) (string, error)
	}
	Processor interface {
		Resize(data io.Reader, output io.Writer, parameters model.SizeInput) error
	}
	Validator interface {
		ValidateFile(in graphql.Upload) error
		ValidateInput(input model.SizeInput) error
	}
)

func NewGraphqlResolver(s Storage, p Processor, r Repository, v Validator) *Resolver {
	return &Resolver{
		storage:   s,
		processor: p,
		repo:      r,
		validator: v,
	}
}
