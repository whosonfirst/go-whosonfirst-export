package main

// https://github.com/jnyfah/outreachy/tree/main/Jennifer/Golang-to-WASI
// https://wasmtime.dev/

import (
	"context"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-export/v2"
	"log"
)

func main() {

	ctx := context.Background()

	ex, err := export.NewExporter(ctx, "whosonfirst://")

	if err != nil {
		log.Fatalf("Failed to create exporter , %v", err)
	}

	var body []byte
	
	pretty, err := ex.Export(ctx, body)

	if err != nil {
		log.Fatalf("Failed to export body, %v", err)
	}

	fmt.Println(pretty)
}
