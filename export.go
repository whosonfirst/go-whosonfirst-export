package export

import (
	"bytes"
	"encoding/json"
	"github.com/tidwall/gjson"
	"github.com/tidwall/pretty"
	"github.com/whosonfirst/go-whosonfirst-export/options"
	"github.com/whosonfirst/go-whosonfirst-export/properties"
	"io"
)

type Feature struct {
	Type       string      `json:"type"`
	Id         int64       `json:"id"`
	Properties interface{} `json:"properties"`
	Bbox       interface{} `json:"bbox,omitempty"`
	Geometry   interface{} `json:"geometry"`
}

func Export(feature []byte, opts options.Options, wr io.Writer) error {

	var err error

	feature, err = Prepare(feature, opts)

	if err != nil {
		return err
	}

	feature, err = Format(feature, opts)

	if err != nil {
		return err
	}

	r := bytes.NewReader(feature)
	_, err = io.Copy(wr, r)

	return err
}

func Prepare(feature []byte, opts options.Options) ([]byte, error) {

	var err error

	feature, err = properties.EnsureWOFId(feature, opts.UIDProvider())

	if err != nil {
		return nil, err
	}

	feature, err = properties.EnsureRequired(feature)

	if err != nil {
		return nil, err
	}

	feature, err = properties.EnsureEDTF(feature)

	if err != nil {
		return nil, err
	}

	feature, err = properties.EnsureParentId(feature)

	if err != nil {
		return nil, err
	}

	feature, err = properties.EnsureBelongsTo(feature)

	if err != nil {
		return nil, err
	}

	feature, err = properties.EnsureSupersedes(feature)

	if err != nil {
		return nil, err
	}

	feature, err = properties.EnsureSupersededBy(feature)

	if err != nil {
		return nil, err
	}

	feature, err = properties.EnsureTimestamps(feature)

	if err != nil {
		return nil, err
	}

	return feature, nil
}

func Format(feature []byte, opts options.Options) ([]byte, error) {

	// see also:
	// https://github.com/tidwall/pretty/issues/2
	// https://gist.github.com/tidwall/ca6ca1dd0cb780f0be4d134f8e4eb7bc

	// the first thing we need to do is ensure that top-level keys
	// are sorted properly (see Feature definition above) specifically
	// so that the bbox and geometry properties are at the end of the
	// file

	var f Feature
	json.Unmarshal(feature, &f)

	// this has the side-effect of ensuring that all the keys in the
	// properties dictionary are sorted automagically

	feature, err := json.Marshal(f)

	if err != nil {
		return nil, err
	}

	// get the geometry (that will be mutated)

	geom_m := gjson.GetBytes(feature, "geometry")

	// sanity checks

	if geom_m.Index == 0 || len(geom_m.Raw) == 0 || geom_m.Raw[0] != '{' || geom_m.Raw[len(geom_m.Raw)-1] != '}' {
		// probably a generic geometry, make it ugly
		return pretty.Ugly(feature), nil
	}

	// make an ugly copy of just the geometry segment
	geom := pretty.Ugly(feature[geom_m.Index : geom_m.Index+len(geom_m.Raw)])

	// now for the tricky part.

	// empty the geometry object.
	// Note that this mutates the original feature.

	for i := 1; i < len(geom_m.Raw)-1; i++ {
		feature[geom_m.Index+i] = ' '
	}

	// make the json pretty

	feature = pretty.Pretty(feature)

	// find the new location of the geometry

	geom_m = gjson.GetBytes(feature, "geometry")

	// allocate the space for the final json.

	final := make([]byte, len(feature)+len(geom)-2)

	copy(final, feature[:geom_m.Index])
	copy(final[geom_m.Index:], geom)
	copy(final[geom_m.Index+len(geom):], feature[geom_m.Index+len(geom_m.Raw):])

	return final, nil
}
