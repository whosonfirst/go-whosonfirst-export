package main

import (
	"errors"
	"github.com/aaronland/go-artisanal-integers"
	"github.com/aaronland/go-artisanal-integers/application"
	"github.com/aaronland/go-artisanal-integers/engine"
	"log"
	"os"
)

func main() {

	flags := application.NewServerApplicationFlags()

	var engine_name string
	var dsn string

	flags.StringVar(&engine_name, "engine", "memory", "...")
	flags.StringVar(&dsn, "dsn", "example", "The data source name (dsn) for connecting to the artisanal integer engine.")

	application.ParseFlags(flags)

	var eng artisanalinteger.Engine
	var err error

	switch engine_name {
	case "fs":
		eng, err = engine.NewFSEngine(dsn)
	case "memory":
		eng, err = engine.NewMemoryEngine("")
	default:
		err = errors.New("Invalid engine")
	}

	if err != nil {
		log.Fatal(err)
	}

	app, err := application.NewServerApplication(eng)

	if err != nil {
		log.Fatal(err)
	}

	err = app.Run(flags)

	if err != nil {
		log.Fatal(err)
	}

	os.Exit(0)
}
