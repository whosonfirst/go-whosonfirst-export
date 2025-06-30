package export

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/whosonfirst/go-whosonfirst-export/v3/properties"
	wof_properties "github.com/whosonfirst/go-whosonfirst-feature/properties"
)

func PrepareAlt(ctx context.Context, feature []byte) ([]byte, error) {

	feature, err := prepareWithoutTimestampsAlt(ctx, feature)

	if err != nil {
		return nil, fmt.Errorf("Failed to prepare without timestamps, %w", err)
	}

	feature, err = prepareTimestamps(ctx, feature)

	if err != nil {
		return nil, fmt.Errorf("Failed to prepare with timestamps, %w", err)
	}

	return feature, nil
}

func prepareWithoutTimestampsAlt(ctx context.Context, feature []byte) ([]byte, error) {

	feature, err := properties.EnsureWOFIdAlt(ctx, feature)

	if err != nil {
		return nil, fmt.Errorf("Failed to ensure wof:id, %w", err)
	}

	_, err = wof_properties.Name(feature)

	if err != nil {
		slog.Warn("Failed to derive name for alternate geometry", "error", err)
	}

	feature, err = properties.EnsurePlacetype(ctx, feature)

	if err != nil {

		return nil, fmt.Errorf("Failed to ensure placetype, %w", err)
	}

	feature, err = properties.EnsureRepo(ctx, feature)

	if err != nil {
		return nil, fmt.Errorf("Failed to ensure repo, %w", err)
	}

	feature, err = properties.EnsureGeom(ctx, feature)

	if err != nil {
		return nil, fmt.Errorf("Failed to ensure geometry, %w", err)
	}

	feature, err = properties.EnsureSourceAltLabel(ctx, feature)

	if err != nil {
		return nil, fmt.Errorf("Failed to ensure src:alt_label, %w", err)
	}

	return feature, nil
}
