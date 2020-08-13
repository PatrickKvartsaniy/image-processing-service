package processor

import (
	"fmt"
	"github.com/PatrickKvartsaniy/image-processing-service/model"
	"github.com/disintegration/imaging"
	"io"
)

type Processor struct{}

func New() *Processor {
	return &Processor{}
}

func (p Processor) Resize(data io.Reader, output io.Writer, parameters model.SizeInput) error {
	original, err := imaging.Decode(data, imaging.AutoOrientation(true))
	if err != nil {
		return fmt.Errorf("decoding: %w", err)
	}
	resized := imaging.Resize(original, int(parameters.Width), int(parameters.Height), imaging.Lanczos)
	return imaging.Encode(output, resized, imaging.PNG)
}
