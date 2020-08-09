package processor

import (
	"github.com/PatrickKvartsaniy/image-processing-service/model"
	"github.com/disintegration/imaging"
	"io"
)

type Processor struct {}

func NewImageProcessor()  *Processor{
	return &Processor{}
}

func (p Processor) Resize(in io.Reader, out io.Writer, params model.SizeInput)  error {
	original, err := imaging.Decode(in, imaging.AutoOrientation(true))
	if err != nil{
		return err
	}
	resized := imaging.Resize(original, params.Width, params.Height, imaging.Lanczos)
	return imaging.Encode(out, resized, imaging.PNG)
}

