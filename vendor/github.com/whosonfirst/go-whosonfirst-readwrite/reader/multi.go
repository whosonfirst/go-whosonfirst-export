package reader

import (
	"errors"
	"fmt"
	"io"
	_ "log"
	"sync"
)

type MultiReader struct {
	Reader
	readers []Reader
	lookup  map[string]int
	mu      *sync.RWMutex
}

// something something something a callback function or
// map to invoke a dsn specific NewSomethingReader method
// but not today... today it's only FS readers...
// (20180807/thisisaaronland)

func NewMultiReaderFromStrings(dsn_strings ...string) (Reader, error) {

	readers := make([]Reader, 0)

	for _, dsn := range dsn_strings {

		r, err := NewFSReader(dsn)

		if err != nil {
			return nil, err
		}

		readers = append(readers, r)
	}

	return NewMultiReader(readers...)
}

func NewMultiReader(readers ...Reader) (Reader, error) {

	lookup := make(map[string]int)

	mu := new(sync.RWMutex)

	mr := MultiReader{
		readers: readers,
		lookup:  lookup,
		mu:      mu,
	}

	return &mr, nil
}

func (mr *MultiReader) Read(uri string) (io.ReadCloser, error) {

	missing := errors.New("Unable to read URI")

	mr.mu.RLock()

	idx, ok := mr.lookup[uri]

	mr.mu.RUnlock()

	if ok {

		// log.Printf("READ MULTIREADER LOOKUP INDEX FOR %s AS %d\n", uri, idx)

		if idx == -1 {
			return nil, missing
		}

		r := mr.readers[idx]
		return r.Read(uri)
	}

	var fh io.ReadCloser
	idx = -1

	for i, r := range mr.readers {

		rsp, err := r.Read(uri)

		if err == nil {

			fh = rsp
			idx = i

			break
		}
	}

	// log.Printf("SET MULTIREADER LOOKUP INDEX FOR %s AS %d\n", uri, idx)

	mr.mu.Lock()
	mr.lookup[uri] = idx
	mr.mu.Unlock()

	if fh == nil {
		return nil, missing
	}

	return fh, nil
}

func (mr *MultiReader) URI(uri string) string {

	mr.mu.RLock()

	idx, ok := mr.lookup[uri]

	mr.mu.RUnlock()

	if ok {
		return mr.readers[idx].URI(uri)
	}

	_, err := mr.Read(uri)

	if err != nil {
		return fmt.Sprintf("x-urn:go-whosonfirst-readwrite:reader:multi#%s", uri)
	}

	return mr.URI(uri)
}
