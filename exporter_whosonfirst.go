package export

import (
	"bytes"
	"context"
	"fmt"
	"net/url"

	"log/slog"
	"os"

	"github.com/whosonfirst/go-whosonfirst-feature/alt"
	"github.com/whosonfirst/go-whosonfirst-validate"
)

type WhosOnFirstExporter struct {
	Exporter
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

	ex := WhosOnFirstExporter{}

	return &ex, nil
}

func (ex *WhosOnFirstExporter) Export(ctx context.Context, feature []byte) (bool, []byte, error) {

	if alt.IsAlt(feature) {
		return ex.exportAlt(ctx, feature)
	}

	return ex.export(ctx, feature)
}

func (ex *WhosOnFirstExporter) export(ctx context.Context, feature []byte) (bool, []byte, error) {

	tmp_feature, err := PrepareFeatureWithoutTimestamps(ctx, feature)

	if err != nil {
		return false, nil, fmt.Errorf("Failed to prepare record (without timestamps), %w", err)
	}

	tmp_feature, err = Format(ctx, tmp_feature)

	if err != nil {
		return false, nil, fmt.Errorf("Failed to format tmp record, %w", err)
	}

	if bytes.Equal(tmp_feature, feature) {
		return false, feature, nil
	}

	new_feature, err := PrepareTimestamps(ctx, tmp_feature)

	if err != nil {
		return true, nil, fmt.Errorf("Failed to prepare record, %w", err)
	}

	err = validate.Validate(new_feature)

	if err != nil {
		return true, nil, fmt.Errorf("Failed to validate record, %w", err)
	}

	new_feature, err = Format(ctx, new_feature)

	if err != nil {
		return true, nil, fmt.Errorf("Failed to format record, %w", err)
	}

	return true, new_feature, nil
}

func (ex *WhosOnFirstExporter) exportAlt(ctx context.Context, feature []byte) (bool, []byte, error) {

	os.Stdout.Write(feature)

	tmp_feature, err := PrepareAltFeatureWithoutTimestamps(ctx, feature)

	if err != nil {
		return false, nil, fmt.Errorf("Failed to prepare input record (without timestamps), %w", err)
	}

	tmp_feature, err = Format(ctx, tmp_feature)

	if err != nil {
		return false, nil, fmt.Errorf("Failed to format tmp record, %w", err)
	}

	os.Stdout.Write(tmp_feature)

	if bytes.Equal(feature, tmp_feature) {
		return false, feature, nil
	}

	new_feature, err := PrepareTimestamps(ctx, tmp_feature)

	if err != nil {
		slog.Info("NOPE 4")
		return true, nil, fmt.Errorf("Failed to prepare record, %w", err)
	}

	err = validate.ValidateAlt(new_feature)

	if err != nil {
		return true, nil, fmt.Errorf("Failed to validate record, %w", err)
	}

	new_feature, err = Format(ctx, new_feature)

	if err != nil {
		slog.Info("NOPE 5")
		return true, nil, fmt.Errorf("Failed to format record, %w", err)
	}

	slog.Info("WUT")
	return true, new_feature, nil
}
