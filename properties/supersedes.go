package properties

import (
	"context"

	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func EnsureSupersedes(ctx context.Context, feature []byte) ([]byte, error) {

	supersedes := make([]int64, 0)

	rsp := gjson.GetBytes(feature, PATH_WOF_SUPERSEDES)

	if rsp.Exists() {
		return feature, nil
	}

	feature, err := sjson.SetBytes(feature, PATH_WOF_SUPERSEDES, supersedes)

	if err != nil {
		return nil, SetPropertyFailed(PATH_WOF_SUPERSEDES, err)
	}

	return feature, nil
}

func EnsureSupersededBy(ctx context.Context, feature []byte) ([]byte, error) {

	superseded_by := make([]int64, 0)

	rsp := gjson.GetBytes(feature, PATH_WOF_SUPERSEDED_BY)

	if rsp.Exists() {
		return feature, nil
	}

	feature, err := sjson.SetBytes(feature, PATH_WOF_SUPERSEDED_BY, superseded_by)

	if err != nil {
		return nil, SetPropertyFailed(PATH_WOF_SUPERSEDED_BY, err)
	}

	return feature, nil
}
