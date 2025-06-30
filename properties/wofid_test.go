package properties

import (
	"context"
	"testing"

	"github.com/whosonfirst/go-whosonfirst-id"
)

type testIdProvider struct {
	id.Provider
}

func (pr *testIdProvider) NewID(ctx context.Context) (int64, error) {
	return int64(999), nil
}

func newTestIdProvider() id.Provider {
	return &testIdProvider{}
}

func TestDefaultIdProvider(t *testing.T) {

	_, err := idProvider(context.Background())

	if err != nil {
		t.Fatalf("Failed to derive default Id provider, %v", err)
	}
}

func TestCustomIdProvider(t *testing.T) {

	ctx := context.Background()

	ctx = context.WithValue(ctx, ID_PROVIDER, newTestIdProvider())

	pr, err := idProvider(ctx)

	if err != nil {
		t.Fatalf("Failed to derive default Id provider, %v", err)
	}

	id, err := pr.NewID(ctx)

	if err != nil {
		t.Fatalf("Failed to derive new ID from custom provider, %v", err)
	}

	if id != 999 {
		t.Fatalf("Unexpected ID")
	}
}

func TestEnsureWOFId(t *testing.T) {

	ctx := context.Background()

	with_id := []byte(`{ "id": 1234, "properties": { "wof:id": 1234 }}`)

	_, err := EnsureWOFId(ctx, with_id)

	if err != nil {
		t.Fatalf("Failed to ensure ID (with ID), %v", err)
	}
}

func TestEnsureWOFIdWithout(t *testing.T) {

	without_id := []byte(`{}`)
	ctx := context.Background()

	_, err := EnsureWOFId(ctx, without_id)

	if err != nil {
		t.Fatalf("Failed to ensure ID (without ID), %v", err)
	}
}
