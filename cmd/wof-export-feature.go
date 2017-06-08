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

		pretty := export.ExportGeoJSON(body)
		fmt.Printf("%s", pretty)

	}

	os.Exit(0)
}
