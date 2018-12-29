package main

import (
	"flag"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-export"
	"io/ioutil"
	"log"
	"os"
)

func main() {

	flag.Parse()

	for _, path := range flag.Args() {

		fh, err := os.Open(path)

		if err != nil {
			log.Fatal(err)
		}

		body, err := ioutil.ReadAll(fh)

		if err != nil {
			log.Fatal(err)
		}

		opts, err := export.DefaultExportOptions()

		if err != nil {
			log.Fatal(err)
		}

		pretty, err := export.Export(body, opts)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s", pretty)

	}

	os.Exit(0)
}
