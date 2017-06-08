package export

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"github.com/tidwall/pretty"
	_ "log"
)

type Feature struct {
	Type       string      `json:"type"`
	Id         int64       `json:"id"`
	Properties interface{} `json:"properties"`
	Bbox       interface{} `json:"bbox"`
	Geometry   interface{} `json:"geometry"`
}

func ExportFeature(feature []byte) ([]byte, error) {

	// see also:
	// https://github.com/tidwall/pretty/issues/2
	// https://gist.github.com/tidwall/ca6ca1dd0cb780f0be4d134f8e4eb7bc

	// the first thing we need to do is ensure that top-level keys
	// are sorted properly (see Feature definition above) specifically
	// so that the bbox and geometry properties are at the end of the
	// file

	var f Feature
	json.Unmarshal(feature, &f)

	feature, err := json.Marshal(f)

	if err != nil {
		return nil, err
	}

	// get the geometry.

	res := gjson.GetBytes(feature, "geometry")

	// sanity checks

	if res.Index == 0 || len(res.Raw) == 0 || res.Raw[0] != '{' || res.Raw[len(res.Raw)-1] != '}' {
		// probably a generic geometry, make it ugly
		return pretty.Ugly(feature), nil
	}

	// make an ugly copy of just the geometry segment
	geom := pretty.Ugly(feature[res.Index : res.Index+len(res.Raw)])

	// now for the tricky part.

	// empty the geometry object.
	// Note that this mutates the original feature.

	for i := 1; i < len(res.Raw)-1; i++ {
		feature[res.Index+i] = ' '
	}

	// make the json pretty

	feature = pretty.Pretty(feature)

	// find the new location of the geometry

	res = gjson.GetBytes(feature, "geometry")

	// allocate the space for the final json.

	final := make([]byte, len(feature)+len(geom)-2)

	copy(final, feature[:res.Index])
	copy(final[res.Index:], geom)
	copy(final[res.Index+len(geom):], feature[res.Index+len(res.Raw):])

	return final, nil
}
