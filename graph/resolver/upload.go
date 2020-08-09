package resolver

import (
	"bytes"
	"io"
)

func copyReader(in io.Reader) (io.Reader, io.Reader) {
	var b bytes.Buffer
	cc := io.TeeReader(in, &b)
	return cc, &b
}
