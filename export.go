package export

import (
	"bytes"
	"fmt"
	"io"
)

func Export(feature []byte, opts *Options, wr io.Writer) error {

	var err error

	feature, err = Prepare(feature, opts)

	if err != nil {
		return fmt.Errorf("Failed to prepare feature, %w", err)
	}

	feature, err = Format(feature, opts)

	if err != nil {
		return fmt.Errorf("Failed to format feature, %w", err)
	}

	r := bytes.NewReader(feature)
	_, err = io.Copy(wr, r)

	if err != nil {
		return fmt.Errorf("Failed to copy feature to writer, %w", err)
	}

	return nil
}

// ExportChanged returns a boolean which indicates whether the file was changed
// by comparing it to the `existingFeature` byte slice, before the lastmodified
// timestamp is incremented. If the `feature` is identical to `existingFeature`
// it doesn't write to the `io.Writer`.
func ExportChanged(feature []byte, existingFeature []byte, opts *Options, wr io.Writer) (changed bool, err error) {

	changed = false

	feature, err = prepareWithoutTimestamps(feature, opts)

	if err != nil {
		return
	}

	feature, err = Format(feature, opts)

	if err != nil {
		return
	}

	changed = !bytes.Equal(feature, existingFeature)

	if !changed {
		return
	}

	feature, err = prepareTimestamps(feature, opts)

	if err != nil {
		return
	}

	feature, err = Format(feature, opts)

	if err != nil {
		return
	}

	r := bytes.NewReader(feature)
	_, err = io.Copy(wr, r)

	return
}
