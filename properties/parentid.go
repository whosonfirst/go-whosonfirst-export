package properties

import (
	"context"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func EnsureParentId(ctx context.Context, feature []byte) ([]byte, error) {

	rsp := gjson.GetBytes(feature, PATH_WOF_PARENTID)

	if rsp.Exists() {
		return feature, nil
	}

	feature, err := sjson.SetBytes(feature, PATH_WOF_PARENTID, -1)

	if err != nil {
		return nil, SetPropertyFailed(PATH_WOF_PARENTID, err)
	}

	return feature, nil
}
