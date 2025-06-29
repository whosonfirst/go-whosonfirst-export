package export

import (
	"context"
	"fmt"

	"github.com/whosonfirst/go-whosonfirst-export/v3/properties"
)

func Prepare(ctx context.Context, feature []byte) ([]byte, error) {

	feature, err := prepareWithoutTimestamps(ctx, feature)

	if err != nil {
		return nil, fmt.Errorf("Failed to prepare without timestamps, %w", err)
	}

	feature, err = prepareTimestamps(ctx, feature)

	if err != nil {
		return nil, fmt.Errorf("Failed to prepare with timestamps, %w", err)
	}

	return feature, nil
}

func prepareWithoutTimestamps(ctx context.Context, feature []byte) ([]byte, error) {

	feature, err := properties.EnsureWOFId(ctx, feature)

	if err != nil {
		return nil, fmt.Errorf("Failed to ensure wof:id, %w", err)
	}

	feature, err = properties.EnsureRequired(ctx, feature)

	if err != nil {
		return nil, fmt.Errorf("Failed to ensure required properties, %w", err)
	}

	feature, err = properties.EnsureEDTF(ctx, feature)

	if err != nil {
		return nil, fmt.Errorf("Failed to ensure EDTF properties, %w", err)
	}

	feature, err = properties.EnsureParentId(ctx, feature)

	if err != nil {
		return nil, fmt.Errorf("Failed to ensure parent ID, %w", err)
	}

	feature, err = properties.EnsureHierarchy(ctx, feature)

	if err != nil {
		return nil, fmt.Errorf("Failed to ensure hierarchy, %w", err)
	}

	feature, err = properties.EnsureBelongsTo(ctx, feature)

	if err != nil {
		return nil, fmt.Errorf("Failed to ensure belongs to, %w", err)
	}

	feature, err = properties.EnsureSupersedes(ctx, feature)

	if err != nil {
		return nil, fmt.Errorf("Failed to ensure supersedes, %w", err)
	}

	feature, err = properties.EnsureSupersededBy(ctx, feature)

	if err != nil {
		return nil, fmt.Errorf("Failed to ensure superseded by, %w", err)
	}

	return feature, nil
}

func prepareTimestamps(ctx context.Context, feature []byte) ([]byte, error) {
	return properties.EnsureTimestamps(ctx, feature)
}
