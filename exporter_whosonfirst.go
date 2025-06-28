package export

import (
	"bytes"
	"context"
	"fmt"
	"net/url"

	"github.com/whosonfirst/go-whosonfirst-id"
	"github.com/whosonfirst/go-whosonfirst-validate"	
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
		return nil, fmt.Errorf("Failed to parse URI, %w", err)
	}

	provider, err := id.NewProvider(ctx)

	if err != nil {
		return nil, fmt.Errorf("Failed to create new Id provider, %w", err)
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
		return false, nil, fmt.Errorf("Failed to prepare record (without timestamps), %w", err)
	}

	tmp_feature, err = Format(tmp_feature)

	if err != nil {
		return false, nil, fmt.Errorf("Failed to format tmp record, %w", err)
	}

	changed := !bytes.Equal(tmp_feature, feature)

	if !changed {
		return false, nil, nil
	}

	new_feature, err := prepareTimestamps(feature, prepare_opts)

	if err != nil {
		return true, nil, fmt.Errorf("Failed to prepare record, %w", err)
	}

	err = validate.Validate(new_feature)

	if err != nil {
	   return true, nil, fmt.Errorf("Failed to validate record, %w", err)
	}
	
	new_feature, err = Format(new_feature)

	if err != nil {
		return true, nil, fmt.Errorf("Failed to format record, %w", err)
	}

	return true, new_feature, nil
}
