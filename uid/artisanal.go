package uid

import (
	"github.com/aaronland/go-artisanal-integers"
)

type ArtisanalUIDProvider struct {
	Provider
	client artisanalinteger.Client	
}

func NewArtisanalUIDProvider(client artisanalinteger.Client) (Provider, error) {

	p := ArtisanalUIDProvider{
		client: client,
	}

	return &p, nil
}

func (p *ArtisanalUIDProvider) UID() (int64, error) {
	return p.client.NextInt()
}
