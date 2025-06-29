package properties

import (
	"context"

	"github.com/tidwall/gjson"
)

func EnsurePlacetype(ctx context.Context, feature []byte) ([]byte, error) {

	rsp := gjson.GetBytes(feature, PATH_WOF_PLACETYPE)

	if !rsp.Exists() {
		return feature, MissingProperty(PATH_WOF_PLACETYPE)
	}

	// Validate placetype?

	return feature, nil
}
