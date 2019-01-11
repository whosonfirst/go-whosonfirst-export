package properties

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func EnsureSrcGeom(feature []byte) ([]byte, error) {

	path := "src:geom"

	rsp := gjson.GetBytes(feature, path)

	if rsp.Exists() {
		return feature, nil
	}

	return sjson.SetBytes(feature, path, "unknown")
}

func EnsureGeomHash(feature []byte) ([]byte, error) {

	rsp := gjson.GetBytes(feature, "geometry")

	if !rsp.Exists() {
		return nil, errors.New("missing geometry!")
	}

	enc, err := json.Marshal(rsp.Value())

	if err != nil {
		return nil, err
	}

	hash := md5.Sum(enc)
	geom_hash := hex.EncodeToString(hash[:])

	return sjson.SetBytes(feature, "wof:geomhash", geom_hash)
}
