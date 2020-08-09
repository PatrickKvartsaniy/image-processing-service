package processor_test

import (
	"bytes"
	"github.com/PatrickKvartsaniy/image-processing-service/model"
	"github.com/PatrickKvartsaniy/image-processing-service/processor"
	"github.com/stretchr/testify/assert"
	"image"
	"os"
	"testing"
)

func TestProcessor_Resize(t *testing.T) {
	proc := processor.NewImageProcessor()
	src, err := os.Open("testdata/gopher.png")
	assert.NoError(t, err)

	var b bytes.Buffer
	params := model.SizeInput{
		Width:  100,
		Height: 100,
	}

	err = proc.Resize(src, &b, params)
	assert.NoError(t, err)

	img, _, err := image.DecodeConfig(&b)
	assert.NoError(t, err)
	assert.Equal(t, 100, img.Height)
	assert.Equal(t, 100, img.Width)
}
