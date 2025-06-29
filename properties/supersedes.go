package properties

import (
	"context"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func EnsureSupersedes(ctx context.Context, feature []byte) ([]byte, error) {

	supersedes := make([]int64, 0)

	rsp := gjson.GetBytes(feature, "properties.wof:supersedes")

	if rsp.Exists() {
		return feature, nil
	}

	return sjson.SetBytes(feature, "properties.wof:supersedes", supersedes)
}

func EnsureSupersededBy(ctx context.Context, feature []byte) ([]byte, error) {

	superseded_by := make([]int64, 0)

	rsp := gjson.GetBytes(feature, "properties.wof:superseded_by")

	if rsp.Exists() {
		return feature, nil
	}

	return sjson.SetBytes(feature, "properties.wof:superseded_by", superseded_by)
}
