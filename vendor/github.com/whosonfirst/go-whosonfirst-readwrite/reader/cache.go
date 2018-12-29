package reader

import (
	"github.com/whosonfirst/go-whosonfirst-cache"
	"io"
	"log"
)

type CacheReader struct {
	Reader
	reader  Reader
	cache   cache.Cache
	options *CacheReaderOptions
}

type CacheReaderOptions struct {
	Debug  bool
	Strict bool
}

func NewDefaultCacheReaderOptions() (*CacheReaderOptions, error) {

	opts := CacheReaderOptions{
		Debug:  false,
		Strict: false,
	}

	return &opts, nil
}

func NewCacheReader(r Reader, c cache.Cache, opts *CacheReaderOptions) (Reader, error) {

	cr := CacheReader{
		reader:  r,
		cache:   c,
		options: opts,
	}

	return &cr, nil
}

func (r *CacheReader) Read(key string) (io.ReadCloser, error) {

	fh, err := r.cache.Get(key)

	if r.options.Debug {
		log.Println("GET", key, err)
	}

	if err == nil {

		if r.options.Debug {
			log.Println("HIT", key)
		}

		return fh, nil
	}

	if err != nil && !cache.IsCacheMiss(err) {
		return nil, err
	}

	if r.options.Debug {
		log.Println("MISS", key)
	}

	fh, err = r.reader.Read(key)

	if r.options.Debug {
		log.Println("READ", key, err)
	}

	if err != nil {
		return nil, err
	}

	fh, err = r.cache.Set(key, fh)

	if r.options.Debug {
		log.Println("SET", key, err)
	}

	if err != nil && r.options.Strict {
		return nil, err
	}

	return fh, nil
}
