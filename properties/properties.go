package properties

import (
	"context"
	"fmt"
)

func EnsureRequired(ctx context.Context, feature []byte) ([]byte, error) {

	var err error

	feature, err = EnsureName(ctx, feature)

	if err != nil {
		return nil, fmt.Errorf("Failed to ensure wof:name, %w", err)
	}

	feature, err = EnsurePlacetype(ctx, feature)

	if err != nil {
		return nil, fmt.Errorf("Failed to ensure placetype, %w", err)
	}

	feature, err = EnsureGeom(ctx, feature)

	if err != nil {
		return nil, fmt.Errorf("Failed to ensure geometry, %w", err)
	}

	return feature, nil
}

func EnsureGeom(ctx context.Context, feature []byte) ([]byte, error) {

	var err error

	feature, err = EnsureSrcGeom(ctx, feature)

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

	var err error

	feature, err = EnsureCreated(ctx, feature)

	if err != nil {
		return nil, fmt.Errorf("Failed to ensure wof:created, %w", err)
	}

	feature, err = EnsureLastModified(ctx, feature)

	if err != nil {
		return nil, fmt.Errorf("Failed to ensure wof:lastmodified, %w", err)
	}

	return feature, nil
}
