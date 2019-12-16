package options

import (
	"context"
	"github.com/aaronland/go-artisanal-integers"
	"github.com/whosonfirst/go-whosonfirst-export/uid"
	"github.com/whosonfirst/go-whosonfirst-id"	
)

type DefaultOptions struct {
	Options
	uid_provider uid.Provider
}

func NewDefaultOptions() (Options, error) {

	id_client, err := id.NewIdClient(context.Background())

	if err != nil {
		return nil, err
	}
	
	return NewDefaultOptionsWithArtisanalIntegerClient(id_client)
}

func NewDefaultOptionsWithArtisanalIntegerClient(client artisanalinteger.Client) (Options, error) {

	provider, err := uid.NewArtisanalUIDProvider(client)

	if err != nil {
		return nil, err
	}

	opts := DefaultOptions{
		uid_provider: provider,
	}

	return &opts, nil
}

func (o *DefaultOptions) UIDProvider() uid.Provider {
	return o.uid_provider
}
