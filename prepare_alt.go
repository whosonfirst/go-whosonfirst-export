package export

import (
	"context"
	"errors"
	"fmt"
	"log/slog"

	"github.com/whosonfirst/go-whosonfirst-export/v3/properties"
)

func PrepareAlt(ctx context.Context, feature []byte) ([]byte, error) {

	slog.Info("YO")

	feature, err := prepareWithoutTimestampsAlt(ctx, feature)

	if err != nil {
		return nil, fmt.Errorf("Failed to prepare without timestamps, %w", err)
	}

	feature, err = prepareTimestamps(ctx, feature)

	if err != nil {
		return nil, fmt.Errorf("Failed to prepare with timestamps, %w", err)
	}

	slog.Info("PEW PEW")
	return feature, nil
}

func prepareWithoutTimestampsAlt(ctx context.Context, feature []byte) ([]byte, error) {

	feature, err := properties.EnsureWOFIdAlt(ctx, feature)

	if err != nil {
		return nil, fmt.Errorf("Failed to ensure wof:id, %w", err)
	}

	feature, err = properties.EnsureName(ctx, feature)

	if err != nil {

		var missing *properties.MissingPropertyError

		if !errors.As(err, &missing) {
			return nil, fmt.Errorf("Failed to ensure name, %w", err)
		}

		slog.Warn("Alt geometry feature is missing wof:name property")
	}

	feature, err = properties.EnsurePlacetype(ctx, feature)

	if err != nil {
		return nil, fmt.Errorf("Failed to ensure placetype, %w", err)
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
