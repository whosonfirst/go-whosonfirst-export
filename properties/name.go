package properties

import (
	"context"
	"errors"

	"github.com/tidwall/gjson"
	_ "github.com/tidwall/sjson"
)

func EnsureName(ctx context.Context, feature []byte) ([]byte, error) {

	rsp := gjson.GetBytes(feature, "properties.wof:name")

	if !rsp.Exists() {
		return feature, errors.New("missing wof:name")
	}

	return feature, nil
}
