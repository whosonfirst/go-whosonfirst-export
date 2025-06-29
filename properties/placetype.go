package properties

import (
	"context"
	"errors"

	"github.com/tidwall/gjson"
	_ "github.com/tidwall/sjson"
)

func EnsurePlacetype(ctx context.Context, feature []byte) ([]byte, error) {

	rsp := gjson.GetBytes(feature, "properties.wof:placetype")

	if !rsp.Exists() {
		return feature, errors.New("missing wof:placetype")
	}

	return feature, nil
}
