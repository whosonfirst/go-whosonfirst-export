package application

import (
	"flag"
	"fmt"
	"github.com/aaronland/go-artisanal-integers"
	"github.com/aaronland/go-artisanal-integers/server"
	"github.com/aaronland/go-artisanal-integers/service"
	"log"
	"net/url"
)

func NewServerApplicationFlags() *flag.FlagSet {

	fs := NewFlagSet("server")

	AssignCommonFlags(fs)

	fs.Int("set-last-int", 0, "Set the last known integer.")
	fs.Int("set-offset", 0, "Set the offset used to mint integers.")
	fs.Int("set-increment", 0, "Set the increment used to mint integers.")

	return fs
}

type ServerApplication struct {
	Application
	engine artisanalinteger.Engine
}

func NewServerApplication(eng artisanalinteger.Engine) (Application, error) {

	a := ServerApplication{
		engine: eng,
	}

	return &a, nil
}

func (s *ServerApplication) Run(fl *flag.FlagSet) error {

	if !fl.Parsed() {
		ParseFlags(fl)
	}

	proto, _ := StringVar(fl, "protocol")
	host, _ := StringVar(fl, "host")
	port, _ := IntVar(fl, "port")
	path, _ := StringVar(fl, "path")

	last, _ := IntVar(fl, "set-last-int")
	offset, _ := IntVar(fl, "set-last-offset")
	increment, _ := IntVar(fl, "set-last-increment")

	if last != 0 {

		err := s.engine.SetLastInt(int64(last))

		if err != nil {
			return err
		}
	}

	if increment != 0 {

		err := s.engine.SetIncrement(int64(increment))

		if err != nil {
			return err
		}
	}

	if offset != 0 {

		err := s.engine.SetOffset(int64(offset))

		if err != nil {
			return err
		}
	}

	svc, err := service.NewArtisanalService("simple", s.engine)

	if err != nil {
		return err
	}

	u := new(url.URL)

	u.Scheme = proto
	u.Host = fmt.Sprintf("%s:%d", host, port)
	u.Path = path

	_, err = url.Parse(u.String())

	if err != nil {
		return err
	}

	svr, err := server.NewArtisanalServer(proto, u)

	if err != nil {
		return err
	}

	log.Println("Listen on", svr.Address())

	return svr.ListenAndServe(svc)
}
