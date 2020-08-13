package resolver

import (
	"bytes"
	"io"
)

func getImageExtension(contentType string) string {
	if e, ok := extensions[contentType]; ok {
		return e
	}
	return ""
}

var extensions = map[string]string{
	"image/jpeg": ".jpeg",
	"image/gif":  ".gif",
	"image/png":  ".png",
	"image/pdf":  ".pdf",
	"image/ico":  ".ico",
}

func copyReader(in io.Reader) (io.Reader, io.Reader) {
	var b bytes.Buffer
	r := io.TeeReader(in, &b)
	return r, &b
}
