package export

import (
	"context"
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/sfomuseum/go-flags/flagset"
	"github.com/whosonfirst/go-whosonfirst-export/v3"
)

func Run(ctx context.Context) error {
	fs := DefaultFlagSet()
	return RunWithFlagSet(ctx, fs)
}

func RunWithFlagSet(ctx context.Context, fs *flag.FlagSet) error {

	flagset.Parse(fs)

	ex, err := export.NewExporter(ctx, exporter_uri)

	if err != nil {
		return fmt.Errorf("Failed to create exporter for '%s', %v", exporter_uri, err)
	}

	for _, path := range flag.Args() {

		r, err := os.Open(path)

		if err != nil {
			return fmt.Errorf("Failed to open %s for reading, %w", path, err)
		}

		defer r.Close()

		body, err := io.ReadAll(r)

		if err != nil {
			return fmt.Errorf("Failed to read %s, %w", path, err)
		}

		_, new_body, err := ex.Export(ctx, body)

		if err != nil {
			return fmt.Errorf("Failed to export %s, %w", path, err)
		}

		os.Stdout.Write(new_body)

	}

	return nil
}
