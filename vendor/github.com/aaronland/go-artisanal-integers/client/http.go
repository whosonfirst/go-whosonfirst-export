package client

import (
	"github.com/aaronland/go-artisanal-integers"
	"io/ioutil"
	_ "log"
	"net/http"
	"net/url"
	"strconv"
)

type HTTPClient struct {
	artisanalinteger.Client
	url *url.URL
}

func NewHTTPClient(u *url.URL) (*HTTPClient, error) {

	cl := HTTPClient{
		url: u,
	}

	return &cl, nil
}

func (cl *HTTPClient) NextInt() (int64, error) {

	rsp, err := http.Get(cl.url.String())

	if err != nil {
		return -1, err
	}

	defer rsp.Body.Close()

	byte_i, err := ioutil.ReadAll(rsp.Body)

	if err != nil {
		return -1, err
	}

	str_i := string(byte_i)

	i, err := strconv.ParseInt(str_i, 10, 64)

	if err != nil {
		return -1, err
	}

	return i, err
}
