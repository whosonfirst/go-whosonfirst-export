package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"

	export "github.com/whosonfirst/go-whosonfirst-export"
	"github.com/whosonfirst/go-whosonfirst-export/exporter"
	"github.com/whosonfirst/go-whosonfirst-export/options"
)

func main() {
	useExporter := flag.Bool("exporter", false, "...")
	flag.Parse()

	opts, err := options.NewDefaultOptions()

	if err != nil {
		log.Fatal(err)
	}

	for _, path := range flag.Args() {

		fh, err := os.Open(path)

		if err != nil {
			log.Fatal(err)
		}

		defer fh.Close()

		body, err := ioutil.ReadAll(fh)

		if err != nil {
			log.Fatal(err)
		}

		if !*useExporter {
			err = export.Export(body, opts, os.Stdout)

			if err != nil {
				log.Fatal(err)
			}

		} else {
			ex, err := exporter.NewWhosOnFirstExporter(opts)

			if err != nil {
				log.Fatal(err)
			}

			pretty, err := ex.Export(body)

			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("%s", pretty)
		}

	}

	os.Exit(0)
}