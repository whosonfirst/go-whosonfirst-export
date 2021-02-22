package export

import (
	"context"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func EnsureProperties(ctx context.Context, body []byte, to_ensure map[string]interface{}) ([]byte, error) {

	to_assign := make(map[string]interface{})

	for path, v := range to_ensure {

		rsp := gjson.GetBytes(body, path)

		if rsp.Exists() {
			continue
		}

		to_assign[path] = v
	}

	return AssignProperties(ctx, body, to_assign)
}

func AssignProperties(ctx context.Context, body []byte, to_assign map[string]interface{}) ([]byte, error) {

	var err error

	for path, v := range to_assign {

		body, err = sjson.SetBytes(body, path, v)

		if err != nil {
			return nil, err
		}
	}

	return body, nil
}

func RemoveProperties(ctx context.Context, body []byte, to_remove []string) ([]byte, error) {

	var err error

	for _, path := range to_remove {

		body, err = sjson.DeleteBytes(body, path)

		if err != nil {
			return nil, err
		}
	}

	return body, nil
}
