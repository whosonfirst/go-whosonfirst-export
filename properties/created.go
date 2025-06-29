package properties

import (
	"context"
	"time"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func EnsureCreated(ctx context.Context, feature []byte) ([]byte, error) {

	var err error

	now := int32(time.Now().Unix())

	created := gjson.GetBytes(feature, PATH_WOF_CREATED)

	if !created.Exists() {

		feature, err = sjson.SetBytes(feature, PATH_WOF_CREATED, now)

		if err != nil {
			return nil, SetPropertyFailed(PATH_WOF_CREATED, err)
		}
	}

	return feature, nil
}
