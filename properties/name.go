package properties

import (
	"context"
	"fmt"

	"github.com/tidwall/gjson"
)

func EnsureName(ctx context.Context, feature []byte) ([]byte, error) {

	rsp := gjson.GetBytes(feature, PATH_WOF_NAME)

	if !rsp.Exists() {
		return nil, MissingProperty(PATH_WOF_NAME)
	}

	if rsp.String() == "" {
		return nil, fmt.Errorf("%s property is empty", PATH_WOF_NAME)
	}

	return feature, nil
}
