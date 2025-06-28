package export

import (
	"context"

	"github.com/whosonfirst/go-whosonfirst-id"
)

type Options struct {
	IdProvider id.Provider
}

func DefaultOptions(ctx context.Context) (*Options, error) {

	provider, err := id.NewProvider(ctx)

	if err != nil {
		return nil, err
	}

	opts := &Options{
		IdProvider: provider,
	}

	return opts, err
}
