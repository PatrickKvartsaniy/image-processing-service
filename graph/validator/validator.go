package validator

import (
	"github.com/99designs/gqlgen/graphql"
	"github.com/PatrickKvartsaniy/image-processing-service/graph/errors"
	"github.com/PatrickKvartsaniy/image-processing-service/model"
)

var allowedTypes = [5]string{"image/jpeg", "image/gif", "image/png", "image/pdf", "image/ico"}

const maxSize = 10 * 1024 * 1024 // 5mb

type Validator struct{}

func NewFileValidator() *Validator {
	return &Validator{}
}

func (v Validator) ValidateFile(in graphql.Upload) error {
	if !isImage(in.ContentType) {
		return errors.UnsupportedFile
	}
	if in.Size > maxSize {
		return errors.TooLarge
	}
	return nil
}

func (v Validator) ValidateInput(in model.SizeInput) error {
	if !in.Valid() {
		return errors.InvalidInput
	}
	return nil
}

func isImage(in string) bool {
	for _, t := range allowedTypes {
		if in == t {
			return true
		}
	}
	return false
}
