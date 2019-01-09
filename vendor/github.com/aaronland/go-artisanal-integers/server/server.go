package server

import (
	"errors"
	"github.com/aaronland/go-artisanal-integers"
	_ "log"
	"net/url"
	"strings"
)

func NewArtisanalServer(proto string, u *url.URL, args ...interface{}) (artisanalinteger.Server, error) {

	var svr artisanalinteger.Server
	var err error

	switch strings.ToUpper(proto) {

	case "HTTP":

		svr, err = NewHTTPServer(u, args...)

	case "TCP":

		svr, err = NewTCPServer(u, args...)

	default:
		return nil, errors.New("Invalid server protocol")
	}

	if err != nil {
		return nil, err
	}

	return svr, nil
}
