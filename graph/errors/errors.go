package errors

import "fmt"

var (
	NotFound        = fmt.Errorf("not found")
	UnsupportedFile = fmt.Errorf("unsupported file type")
	TooLarge        = fmt.Errorf("image is too large")
	InvalidInput    = fmt.Errorf("invalid input parameters")
)
