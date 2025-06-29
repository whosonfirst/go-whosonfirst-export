package export

import (
	"context"

	"github.com/whosonfirst/go-whosonfirst-format"
)

func Format(ctx context.Context, body []byte) ([]byte, error) {
	return format.FormatBytes(body)
}
