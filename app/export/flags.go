package export

import (
	"flag"

	"github.com/sfomuseum/go-flags/flagset"
)

var exporter_uri string

func DefaultFlagSet() *flag.FlagSet {

	fs := flagset.NewFlagSet("export")

	fs.StringVar(&exporter_uri, "exporter-uri", "whosonfirst://", "A valid whosonfirst/go-whosonfirst-export/v3.Exporter URI")

	return fs
}
