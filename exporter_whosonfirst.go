package export

import (
	"bytes"
	"context"
	"net/url"

	"github.com/whosonfirst/go-whosonfirst-id"
)

type WhosOnFirstExporter struct {
	Exporter
	id_provider id.Provider
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

	provider, err := id.NewProvider(ctx)

	if err != nil {
		return nil, err
	}

	ex := WhosOnFirstExporter{
		id_provider: provider,
	}

	return &ex, nil
}

func (ex *WhosOnFirstExporter) Export(ctx context.Context, feature []byte) (bool, []byte, error) {

	prepare_opts := &PrepareOptions{
		IdProvider: ex.id_provider,
	}

	tmp_feature, err := prepareWithoutTimestamps(feature, prepare_opts)

	if err != nil {
		return false, nil, err
	}

	tmp_feature, err = Format(tmp_feature)

	if err != nil {
		return false, nil, err
	}

	changed := !bytes.Equal(tmp_feature, feature)

	if !changed {
		return false, nil, nil
	}

	new_feature, err := prepareTimestamps(feature, prepare_opts)

	if err != nil {
		return true, nil, err
	}

	new_feature, err = Format(new_feature)

	if err != nil {
		return true, nil, err
	}

	return true, new_feature, nil
}
