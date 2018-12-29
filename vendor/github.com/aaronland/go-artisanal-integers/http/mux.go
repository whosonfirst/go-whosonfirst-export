package http

import (
	"fmt"
	"github.com/aaronland/go-artisanal-integers"
	gohttp "net/http"
	gourl "net/url"
	"strings"
)

func NewServeMux(s artisanalinteger.Service, u *gourl.URL) (*gohttp.ServeMux, error) {

	integer_handler, err := IntegerHandler(s)

	if err != nil {
		return nil, err
	}

	integer_path := u.Path

	if !strings.HasPrefix(integer_path, "/") {
		integer_path = fmt.Sprintf("/%s", integer_path)
	}

	ping_handler, err := PingHandler()

	if err != nil {
		return nil, err
	}

	ping_path := "/ping"

	mux := gohttp.NewServeMux()

	mux.Handle(integer_path, integer_handler)
	mux.Handle(ping_path, ping_handler)

	return mux, nil
}
