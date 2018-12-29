package client

import (
	"errors"
	"github.com/aaronland/go-artisanal-integers"
	"net/url"
	"strings"
)

func NewArtisanalClient(proto string, u *url.URL) (artisanalinteger.Client, error) {

	var cl artisanalinteger.Client
	var err error

	switch strings.ToUpper(proto) {

	case "HTTP":
		cl, err = NewHTTPClient(u)
	case "HTTPS":
		cl, err = NewHTTPClient(u)
	case "TCP":
		cl, err = NewTCPClient(u)
	default:
		return nil, errors.New("Invalid client protocol")
	}

	if err != nil {
		return nil, err
	}

	return cl, nil
}
