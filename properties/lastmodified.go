package properties

import (
	"time"
	"github.com/tidwall/sjson"
	"github.com/whosonfirst/go-whosonfirst-export"	// this was cause a circular import...
)

func EnsureLastModified(feature []byte, opts *export.ExportOptions) ([]byte, error) {

	var err error

	now := int32(time.Now().Unix())

	feature, err = sjson.SetBytes(feature, "properties.wof:lastmodified", now)

	if err != nil {
		return nil, err
	}

	return feature, nil
}
