package main

import (
	"context"
	"flag"
	"fmt"
	"io"
	"log"
	"os"

	"github.com/whosonfirst/go-whosonfirst-export/v3"
)

func main() {

	exporter_uri := flag.String("exporter-uri", "whosonfirst://", "A valid whosonfirst/go-whosonfirst-export URI")

	flag.Parse()

	ctx := context.Background()

	ex, err := export.NewExporter(ctx, *exporter_uri)

	if err != nil {
		log.Fatalf("Failed to create exporter for '%s', %v", *exporter_uri, err)
	}

	for _, path := range flag.Args() {

		fh, err := os.Open(path)

		if err != nil {
			log.Fatal(err)
		}

		defer fh.Close()

		body, err := io.ReadAll(fh)

		if err != nil {
			log.Fatal(err)
		}

		pretty, err := ex.Export(ctx, body)

		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("%s", pretty)

	}

	os.Exit(0)
}
