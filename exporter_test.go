package export

import (
	"context"
	"testing"
)

func TestExporterSchemes(t *testing.T) {

	ctx := context.Background()

	for _, s := range ExporterSchemes() {

		_, err := NewExporter(ctx, s)

		if err != nil {
			t.Fatalf("Failed to create '%s' exporter, %v", s, err)
		}
	}
}
