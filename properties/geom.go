package properties

import (
	"geom"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)
	
func EnsureGeom(feature []byte) ([]byte, error) {

	rsp := gjson.GetBytes(feature, "geometry")

	if !rsp.Exists() {
		return nil, errors.New("missing geometry!")
	}


}
