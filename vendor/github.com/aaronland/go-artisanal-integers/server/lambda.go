package server

import (
	"github.com/aaronland/go-artisanal-integers"
	"github.com/aaronland/go-artisanal-integers/http"
	"github.com/whosonfirst/algnhsa"
	_ "log"
	gourl "net/url"
)

type LambdaServer struct {
	artisanalinteger.Server
	url *gourl.URL
}

func NewLambdaServer(u *gourl.URL, args ...interface{}) (*LambdaServer, error) {

	server := LambdaServer{
		url: u,
	}

	return &server, nil
}

func (s *LambdaServer) Address() string {
	return s.url.String()
}

func (s *LambdaServer) ListenAndServe(service artisanalinteger.Service) error {

	mux, err := http.NewServeMux(service, s.url)

	if err != nil {
		return err
	}

	algnhsa.ListenAndServe(mux, nil)
	return nil
}
