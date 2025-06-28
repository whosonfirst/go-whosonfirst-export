package export

import (
	"fmt"

	"github.com/whosonfirst/go-whosonfirst-export/v3/properties"
	"github.com/whosonfirst/go-whosonfirst-feature/alt"
	"github.com/whosonfirst/go-whosonfirst-id"
)

type PrepareOptions struct {
	IdProvider id.Provider
}

func Prepare(feature []byte, opts *PrepareOptions) ([]byte, error) {

	var err error

	feature, err = prepareWithoutTimestamps(feature, opts)

	if err != nil {
		return nil, fmt.Errorf("Failed to prepare without timestamps, %w", err)
	}

	feature, err = prepareTimestamps(feature, opts)

	if err != nil {
		return nil, fmt.Errorf("Failed to prepare with timestamps, %w", err)
	}

	return feature, nil
}

func prepareWithoutTimestamps(feature []byte, opts *PrepareOptions) ([]byte, error) {

	if alt.IsAlt(feature) {
		return prepareWithoutTimestampsAsAlternateGeometry(feature, opts)
	}

	var err error

	feature, err = properties.EnsureWOFId(feature, opts.IdProvider)

	if err != nil {
		return nil, fmt.Errorf("Failed to ensure wof:id, %w", err)
	}

	feature, err = properties.EnsureRequired(feature)

	if err != nil {
		return nil, fmt.Errorf("Failed to ensure required properties, %w", err)
	}

	feature, err = properties.EnsureEDTF(feature)

	if err != nil {
		return nil, fmt.Errorf("Failed to ensure EDTF properties, %w", err)
	}

	feature, err = properties.EnsureParentId(feature)

	if err != nil {
		return nil, fmt.Errorf("Failed to ensure parent ID, %w", err)
	}

	feature, err = properties.EnsureHierarchy(feature)

	if err != nil {
		return nil, fmt.Errorf("Failed to ensure hierarchy, %w", err)
	}

	feature, err = properties.EnsureBelongsTo(feature)

	if err != nil {
		return nil, fmt.Errorf("Failed to ensure belongs to, %w", err)
	}

	feature, err = properties.EnsureSupersedes(feature)

	if err != nil {
		return nil, fmt.Errorf("Failed to ensure supersedes, %w", err)
	}

	feature, err = properties.EnsureSupersededBy(feature)

	if err != nil {
		return nil, fmt.Errorf("Failed to ensure superseded by, %w", err)
	}

	return feature, nil
}

func prepareWithoutTimestampsAsAlternateGeometry(feature []byte, opts *PrepareOptions) ([]byte, error) {

	var err error

	feature, err = properties.EnsureWOFId(feature, opts.IdProvider)

	if err != nil {
		return nil, fmt.Errorf("Failed to ensure wof:id, %w", err)
	}

	feature, err = properties.EnsureRequired(feature)

	if err != nil {
		return nil, fmt.Errorf("Failed to ensure required properties, %w", err)
	}

	feature, err = properties.EnsureSourceAltLabel(feature)

	if err != nil {
		return nil, fmt.Errorf("Failed to ensure src:alt_label, %w", err)
	}

	return feature, nil
}

func prepareTimestamps(feature []byte, opts *PrepareOptions) ([]byte, error) {
	return properties.EnsureTimestamps(feature)
}
