package properties

import (
	"context"
	"fmt"
)

func EnsureGeom(ctx context.Context, feature []byte) ([]byte, error) {

	feature, err := EnsureSrcGeom(ctx, feature)

	if err != nil {
		return nil, fmt.Errorf("Failed to ensure src:geom, %w", err)
	}

	feature, err = EnsureGeomHash(ctx, feature)

	if err != nil {
		return nil, fmt.Errorf("Failed to ensure geom:hash, %w", err)
	}

	feature, err = EnsureGeomCoords(ctx, feature)

	if err != nil {
		return nil, fmt.Errorf("Failed to ensure geometry coordinates, %w", err)
	}

	return feature, nil
}

func EnsureTimestamps(ctx context.Context, feature []byte) ([]byte, error) {

	feature, err := EnsureCreated(ctx, feature)

	if err != nil {
		return nil, fmt.Errorf("Failed to ensure wof:created, %w", err)
	}

	feature, err = EnsureLastModified(ctx, feature)

	if err != nil {
		return nil, fmt.Errorf("Failed to ensure wof:lastmodified, %w", err)
	}

	return feature, nil
}
