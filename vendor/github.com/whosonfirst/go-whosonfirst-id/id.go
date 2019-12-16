package id

import (
	"context"
	"github.com/aaronland/go-artisanal-integers"
	_ "github.com/aaronland/go-brooklynintegers-api"
	"github.com/aaronland/go-uid"
	"github.com/aaronland/go-uid-artisanal"
	"strconv"
)

type IdClient struct {
	artisanalinteger.Client
	provider uid.Provider
}

func NewIdClient(ctx context.Context) (artisanalinteger.Client, error) {

	opts := &artisanal.ArtisanalProviderURIOptions{
		Pool:    "memory://",
		Minimum: 0,
		Clients: []string{
			"brooklynintegers://",
		},
	}

	uri, err := artisanal.NewArtisanalProviderURI(opts)

	if err != nil {
		return nil, err
	}

	// str_uri ends up looking like this:
	// artisanal:?client=brooklynintegers%3A%2F%2F&minimum=5&pool=memory%3A%2F%2F

	str_uri := uri.String()

	return NewIdClientWithURI(ctx, str_uri)
}

func NewIdClientWithURI(ctx context.Context, uri string) (artisanalinteger.Client, error) {

	pr, err := uid.NewProvider(ctx, uri)

	if err != nil {
		return nil, err
	}

	cl := &IdClient{
		provider: pr,
	}

	return cl, nil
}

func (cl *IdClient) NextInt() (int64, error) {

	uid, err := cl.provider.UID()

	if err != nil {
		return -1, err
	}

	str_id := uid.String()
	return strconv.ParseInt(str_id, 10, 64)
}
