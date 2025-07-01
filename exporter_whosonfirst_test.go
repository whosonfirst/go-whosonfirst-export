package export

import (
	"context"
	"testing"
)

func TestWhosOnFirstExporter(t *testing.T) {

	ctx := context.Background()
	_, err := NewExporter(ctx, "whosonfirst://")

	if err != nil {
		t.Fatalf("Failed to create whosonfirst:// exporter, %v", err)
	}
}
