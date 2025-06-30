package properties

import (
	"context"
	"time"

	"github.com/tidwall/sjson"
	wof_properties "github.com/whosonfirst/go-whosonfirst-feature/properties"
)

func EnsureLastModified(ctx context.Context, feature []byte) ([]byte, error) {

	now := int32(time.Now().Unix())

	feature, err := sjson.SetBytes(feature, wof_properties.PATH_WOF_LASTMODIFIED, now)

	if err != nil {
		return nil, SetPropertyFailed(wof_properties.PATH_WOF_LASTMODIFIED, err)
	}

	return feature, nil
}
