package application

import (
	"flag"
	"fmt"
	"github.com/aaronland/go-artisanal-integers/client"
	"log"
	"net/url"
)

func NewClientApplicationFlags() *flag.FlagSet {

	fs := NewFlagSet("client")

	AssignCommonFlags(fs)

	return fs
}

type ClientApplication struct {
	Application
}

func NewClientApplication() (Application, error) {

	c := ClientApplication{}
	return &c, nil
}

func (c *ClientApplication) Run(fl *flag.FlagSet) error {

	if !fl.Parsed() {
		ParseFlags(fl)
	}

	proto, _ := StringVar(fl, "protocol")
	host, _ := StringVar(fl, "host")
	port, _ := IntVar(fl, "port")
	path, _ := StringVar(fl, "path")

	u := new(url.URL)

	u.Scheme = proto
	u.Host = fmt.Sprintf("%s:%d", host, port)
	u.Path = path

	_, err := url.Parse(u.String())

	if err != nil {
		return err
	}

	cl, err := client.NewArtisanalClient(proto, u)

	if err != nil {
		return err
	}

	i, err := cl.NextInt()

	if err != nil {
		return err
	}

	log.Println(i)
	return nil
}
