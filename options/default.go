package options

import (
	brooklyn_integers "github.com/aaronland/go-brooklynintegers-api"
	"github.com/whosonfirst/go-whosonfirst-export/uid"
)

type DefaultOptions struct {
	Options
	uid_provider uid.Provider
}

func NewDefaultOptions() (Options, error) {

	bi_client := brooklyn_integers.NewAPIClient()
	provider, err := uid.NewArtisanalUIDProvider(bi_client)

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
