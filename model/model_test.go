package model_test

import (
	"github.com/PatrickKvartsaniy/image-processing-service/model"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestImage_IncreaseVersion(t *testing.T) {
	var img model.Image
	img.IncreaseVersion()
	assert.Equal(t, int64(1), img.Version)
	img.IncreaseVersion()
	assert.Equal(t, int64(2), img.Version)
}

func TestImage_NewSize(t *testing.T) {
	tc := struct {
		newPath  string
		newSize  model.SizeInput
		expected []model.Resized
	}{
		newPath: "drive.google.com/drive/u/5/my-drive",
		newSize: model.SizeInput{
			Width:  100,
			Height: 120,
		},
		expected: []model.Resized{
			{
				Path:   "drive.google.com/drive/u/5/my-drive",
				Width:  100,
				Height: 120,
			},
		},
	}
	var img model.Image
	img.NewSize(tc.newPath, "", tc.newSize)
	assert.Equal(t, tc.expected, img.Variety)
}

func TestSizeInput_Valid(t *testing.T) {
	tcs := []struct {
		in       model.SizeInput
		expected bool
	}{
		{
			in: model.SizeInput{
				Width:  10,
				Height: 20,
			},
			expected: true,
		},
		{
			in: model.SizeInput{
				Width:  -1,
				Height: 11,
			},
		},
	}
	for _, tc := range tcs {
		assert.Equal(t, tc.expected, tc.in.Valid())
	}
}
