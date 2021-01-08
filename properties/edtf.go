package properties

import (
	"github.com/sfomuseum/go-edtf"
	"github.com/sfomuseum/go-edtf/parser"
	"github.com/tidwall/gjson"
	"github.com/tidwall/sjson"
)

func EnsureEDTF(feature []byte) ([]byte, error) {

	var err error

	feature, err = EnsureInception(feature)

	if err != nil {
		return nil, err
	}

	feature, err = EnsureCessation(feature)

	if err != nil {
		return nil, err
	}

	return feature, nil
}

func EnsureInception(feature []byte) ([]byte, error) {

	path := "properties.edtf:inception"

	rsp := gjson.GetBytes(feature, path)

	if rsp.Exists() {

		edtf_str := rsp.String()

		_, err := parser.ParseString(edtf_str)

		if err != nil {
			return nil, err
		}

		// set date:FOO here

		return feature, nil
	}

	return sjson.SetBytes(feature, path, edtf.UNKNOWN)
}

func EnsureCessation(feature []byte) ([]byte, error) {

	path := "properties.edtf:cessation"

	rsp := gjson.GetBytes(feature, path)

	if rsp.Exists() {

		edtf_str := rsp.String()

		_, err := parser.ParseString(edtf_str)

		if err != nil {
			return nil, err
		}

		// set date:BAR here

		return feature, nil
	}

	return sjson.SetBytes(feature, path, edtf.UNKNOWN)
}
