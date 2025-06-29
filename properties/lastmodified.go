package properties

import (
	"context"
	"time"

	"github.com/tidwall/sjson"
)

func EnsureLastModified(ctx context.Context, feature []byte) ([]byte, error) {

	now := int32(time.Now().Unix())

	feature, err := sjson.SetBytes(feature, PATH_WOF_LASTMODIFIED, now)

	if err != nil {
		return nil, SetPropertyFailed(PATH_WOF_LASTMODIFIED, err)
	}

	return feature, nil
}
