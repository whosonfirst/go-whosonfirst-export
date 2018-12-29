package server

import (
	"github.com/aaronland/go-artisanal-integers"
	"github.com/aaronland/go-artisanal-integers/http"
	_ "log"
	gohttp "net/http"
	gourl "net/url"
)

type HTTPServer struct {
	artisanalinteger.Server
	url *gourl.URL
}

func NewHTTPServer(u *gourl.URL, args ...interface{}) (*HTTPServer, error) {

	u.Scheme = "http"

	server := HTTPServer{
		url: u,
	}

	return &server, nil
}

func (s *HTTPServer) Address() string {
	return s.url.String()
}

func (s *HTTPServer) ListenAndServe(service artisanalinteger.Service) error {

	mux, err := http.NewServeMux(service, s.url)

	if err != nil {
		return err
	}

	return gohttp.ListenAndServe(s.url.Host, mux)
}
