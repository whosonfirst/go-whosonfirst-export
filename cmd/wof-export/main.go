package main

import (
	"context"
	"log"

	"github.com/whosonfirst/go-whosonfirst-export/v3/app/export"
)

func main() {

	ctx := context.Background()
	err := export.Run(ctx)

	if err != nil {
		log.Fatalf("Failed to run exporter, %v", err)
	}
}
