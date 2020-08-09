package resolver

import (
	"bytes"
	"io"
)

func copyReader(in io.Reader) (io.Reader, io.Reader) {
	var b bytes.Buffer
	r := io.TeeReader(in, &b)
	return r, &b
}
