package export

import (
	"context"

	"github.com/whosonfirst/go-whosonfirst-export/v3/properties"
)

func PrepareTimestamps(ctx context.Context, feature []byte) ([]byte, error) {
	return properties.EnsureTimestamps(ctx, feature)
}
