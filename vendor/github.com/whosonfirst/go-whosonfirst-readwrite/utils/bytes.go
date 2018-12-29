package utils

import (
	"bytes"
	"io"
	"io/ioutil"
)

func ReadCloserFromBytes(b []byte) (io.ReadCloser, error) {
	body := bytes.NewReader(b)
	return ioutil.NopCloser(body), nil
}
