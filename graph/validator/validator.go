package validator

import (
	"fmt"
	"github.com/99designs/gqlgen/graphql"
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
		return fmt.Errorf("unsupported file type")
	}
	if in.Size > maxSize {
		return fmt.Errorf("image size too large")
	}
	return nil
}

func (v Validator) ValidateInput(in model.SizeInput) error {
	if !in.Valid() {
		return fmt.Errorf("invalid input parameters")
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
