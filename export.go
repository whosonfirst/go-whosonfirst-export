package export

import (
	"github.com/tidwall/gjson"
	"github.com/tidwall/pretty"
)

// https://gist.github.com/tidwall/ca6ca1dd0cb780f0be4d134f8e4eb7bc

func ExportGeoJSON(json []byte) []byte {

	// get the geometry.
	res := gjson.GetBytes(json, "geometry")

	// sanity checks
	if res.Index == 0 || len(res.Raw) == 0 ||
		res.Raw[0] != '{' || res.Raw[len(res.Raw)-1] != '}' {
		// probably a generic geometry, make it ugly
		return pretty.Ugly(json)
	}

	// make an ugly copy of just the geometry segment
	geom := pretty.Ugly(json[res.Index : res.Index+len(res.Raw)])

	// now for the tricky part.

	// empty the geometry object.
	// Note that this mutates the original json.
	for i := 1; i < len(res.Raw)-1; i++ {
		json[res.Index+i] = ' '
	}

	// make the json pretty
	json = pretty.Pretty(json)

	// find the new location of the geometry
	res = gjson.GetBytes(json, "geometry")

	// allocate the space for the final json.
	final := make([]byte, len(json)+len(geom)-2)
	copy(final, json[:res.Index])
	copy(final[res.Index:], geom)
	copy(final[res.Index+len(geom):], json[res.Index+len(res.Raw):])

	return final
}
