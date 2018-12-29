package service

import (
	"errors"
	"github.com/aaronland/go-artisanal-integers"
	"strings"
)

func NewArtisanalService(proto string, eng artisanalinteger.Engine) (artisanalinteger.Service, error) {

	var svc artisanalinteger.Service
	var err error

	switch strings.ToUpper(proto) {

	case "SIMPLE":
		svc, err = NewSimpleService(eng)
	default:
		return nil, errors.New("Invalid service protocol")
	}

	if err != nil {
		return nil, err
	}

	return svc, nil
}
