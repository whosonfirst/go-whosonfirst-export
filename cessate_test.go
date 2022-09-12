package export

import (
	"context"
	"github.com/tidwall/gjson"
	"io"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestCessateRecord(t *testing.T) {

	ctx := context.Background()

	rel_path := "fixtures/no-changes.geojson"
	abs_path, err := filepath.Abs(rel_path)

	if err != nil {
		t.Fatalf("Failed to derive absolute path for %s, %v", rel_path, err)
	}

	r, err := os.Open(abs_path)

	if err != nil {
		t.Fatalf("Failed to open %s for reading, %v", abs_path, err)
	}

	defer r.Close()

	body, err := io.ReadAll(r)

	if err != nil {
		t.Fatalf("Failed to read %s, %v", abs_path, err)
	}

	ex, err := NewExporter(ctx, "whosonfirst://")

	if err != nil {
		t.Fatalf("Failed to create exporter, %v", err)
	}

	now := time.Now()

	new_body, err := CessateRecordWithTime(ctx, ex, now, body)

	if err != nil {
		t.Fatalf("Failed to cessate record, %v", err)
	}

	expected := map[string]string{
		"properties.edtf:cessation": now.Format("2006-01-02"),
		"properties.mz:is_current":  "0",
	}

	for path, value := range expected {

		rsp := gjson.GetBytes(new_body, path)

		if rsp.String() != value {
			t.Fatalf("Unexpected value for %s: %s (expected %s)", path, rsp.String(), value)
		}
	}

}
