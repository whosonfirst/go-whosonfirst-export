package export

import (
	"context"
	"fmt"

	"github.com/whosonfirst/go-whosonfirst-export/v3/properties"
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

	feature, err := properties.EnsureWOFId(ctx, feature)

	if err != nil {
		return nil, fmt.Errorf("Failed to ensure wof:id, %w", err)
	}

	feature, err = properties.EnsureRequired(ctx, feature)

	if err != nil {
		return nil, fmt.Errorf("Failed to ensure required properties, %w", err)
	}

	feature, err = properties.EnsureSourceAltLabel(ctx, feature)

	if err != nil {
		return nil, fmt.Errorf("Failed to ensure src:alt_label, %w", err)
	}

	return feature, nil
}
