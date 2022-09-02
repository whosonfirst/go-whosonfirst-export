package main

// https://github.com/jnyfah/outreachy/tree/main/Jennifer/Golang-to-WASI
// https://wasmtime.dev/
// https://tinygo.org/docs/reference/lang-support/stdlib/

import (
	"context"
	"flag"
	"fmt"
	"github.com/whosonfirst/go-whosonfirst-export/v2"
	"log"
	"strings"
)

func main() {

	ctx := context.Background()

	// START OF this is a profoundly dumb way to read from STDIN
	// but I can't figure out how else to do it in tinygo...

	flag.Parse()
	body := strings.Join(flag.Args(), "")

	// END OF this is a profoundly dumb way to read from STDIN

	/*

		> wasmtime export.wasm
		Error: failed to run main module `export.wasm`

		Caused by:
		    0: failed to instantiate "export.wasm"
		    1: unknown import: `env::time.resetTimer` has not been defined

		TinyGo does not support time.Tickers:

		https://github.com/tinygo-org/tinygo/issues/1037

		These are used by the following depedencies:

		vendor/github.com/andres-erbsen/clock/clock.go
		vendor/github.com/cenkalti/backoff/v4/backoff.go

		Even if we got this far though none of the net/* packages are supported
		which means any calls to the Brooklyn Integers API would fail.

	*/

	ex, err := export.NewExporter(ctx, "whosonfirst://")

	if err != nil {
		log.Fatalf("Failed to create exporter, %v", err)
	}

	fmt.Println(ex)
}
