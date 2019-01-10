package export

import (
	"encoding/json"
	"github.com/aaronland/go-artisanal-integers"
	brooklyn_integers "github.com/aaronland/go-brooklynintegers-api"
	"github.com/tidwall/gjson"
	"github.com/tidwall/pretty"
	"github.com/whosonfirst/go-whosonfirst-export/properties"
)

type Feature struct {
	Type       string      `json:"type"`
	Id         int64       `json:"id"`
	Properties interface{} `json:"properties"`
	Bbox       interface{} `json:"bbox,omitempty"`
	Geometry   interface{} `json:"geometry"`
}

type ExportOptions struct {
	Strict        bool
	IntegerClient artisanalinteger.Client
}

func DefaultExportOptions() (*ExportOptions, error) {

	bi_client := brooklyn_integers.NewAPIClient()

	opts := ExportOptions{
		Strict:        true,
		IntegerClient: bi_client,
	}

	return &opts, nil
}

type Exporter interface {
	Export([]byte) ([]byte, error)
	ExportFeature(interface{}) ([]byte, error)
}

type FeatureExporter struct {
	Exporter
	options *ExportOptions
}

func NewExporter(opts *ExportOptions) (Exporter, error) {

	ex := FeatureExporter{
		options: opts,
	}

	return &ex, nil
}

func (ex *FeatureExporter) ExportFeature(feature interface{}) ([]byte, error) {

	body, err := json.Marshal(feature)

	if err != nil {
		return nil, err
	}

	return ex.Export(body)
}

func (ex *FeatureExporter) Export(feature []byte) ([]byte, error) {

	var err error

	feature, err = Prepare(feature, ex.options)

	if err != nil {
		return nil, err
	}

	feature, err = Format(feature, ex.options)

	if err != nil {
		return nil, err
	}

	return feature, nil
}

func Prepare(feature []byte, opts *ExportOptions) ([]byte, error) {

	var err error

	feature, err = properties.EnsureCreated(feature)

	if err != nil {
		return nil, err
	}

	feature, err = properties.EnsureLastModified(feature)

	if err != nil {
		return nil, err
	}

	feature, err = properties.EnsureWOFId(feature, opts.IntegerClient)

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

	return feature, nil
}

func Format(feature []byte, opts *ExportOptions) ([]byte, error) {

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
