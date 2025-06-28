package export

import (
	"bytes"
	"context"
	"fmt"
	"io"
)

func Export(ctx context.Context, feature []byte) (bool, []byte, error) {

	ex, err := NewExporter(ctx, "whosonfirst://")

	if err != nil {
		return false, nil, fmt.Errorf("Failed to create exporter, %w", err)
	}

	return ex.Export(ctx, feature)
}

func WriteExportIfChanged(ctx context.Context, feature []byte, wr io.Writer) (bool, error) {

	has_changed, body, err := Export(ctx, feature)

	if err != nil {
		return has_changed, fmt.Errorf("Failed to export feature, %w", err)
	}

	if !has_changed {
		return false, nil
	}

	r := bytes.NewReader(body)
	_, err = io.Copy(wr, r)

	if err != nil {
		return true, fmt.Errorf("Failed to copy feature to writer, %w", err)
	}

	return true, nil
}
