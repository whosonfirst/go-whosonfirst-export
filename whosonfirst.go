package export

import (
	"context"
	"net/url"
)

type WhosOnFirstExporter struct {
	Exporter
	options *Options
}

func init() {

	ctx := context.Background()

	err := RegisterExporter(ctx, "whosonfirst", NewWhosOnFirstExporter)

	if err != nil {
		panic(err)
	}
}

func NewWhosOnFirstExporter(ctx context.Context, uri string) (Exporter, error) {

	_, err := url.Parse(uri)

	if err != nil {
		return nil, err
	}

	opts, err := DefaultOptions(ctx)

	ex := WhosOnFirstExporter{
		options: opts,
	}

	return &ex, nil
}

func (ex *WhosOnFirstExporter) Export(ctx context.Context, feature []byte) ([]byte, error) {

	var err error

	feature, err = Prepare(feature, ex.options)

	if err != nil {
		return nil, err
	}

	feature, err = Format(feature, ex.options)

	if err != nil {
		return nil, err
	}

	return feature, nil
}
