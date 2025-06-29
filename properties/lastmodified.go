package properties

import (
	"context"
	"time"

	"github.com/tidwall/sjson"
)

func EnsureLastModified(ctx context.Context, feature []byte) ([]byte, error) {

	var err error

	now := int32(time.Now().Unix())

	feature, err = sjson.SetBytes(feature, "properties.wof:lastmodified", now)

	if err != nil {
		return nil, err
	}

	return feature, nil
}
