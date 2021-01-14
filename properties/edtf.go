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

		d, err := parser.ParseString(edtf_str)

		if err != nil {
			return nil, err
		}

		lower_ts, err := d.Lower()

		if err != nil {
			return nil, err
		}

		if lower_ts != nil {
			return sjson.SetBytes(feature, "properties.date:inception_lower", lower_ts.Unix())
		}

		upper_ts, err := d.Upper()

		if err != nil {
			return nil, err
		}

		if upper_ts != nil {
			return sjson.SetBytes(feature, "properties.date:inception_upper", upper_ts.Unix())
		}

		return feature, nil
	}

	return sjson.SetBytes(feature, path, edtf.UNKNOWN)
}

func EnsureCessation(feature []byte) ([]byte, error) {

	path := "properties.edtf:cessation"

	rsp := gjson.GetBytes(feature, path)

	if rsp.Exists() {

		edtf_str := rsp.String()

		d, err := parser.ParseString(edtf_str)

		if err != nil {
			return nil, err
		}

		lower_ts, err := d.Lower()

		if err != nil {
			return nil, err
		}

		if lower_ts != nil {
			return sjson.SetBytes(feature, "properties.date:cessation_lower", lower_ts.Unix())
		}

		upper_ts, err := d.Upper()

		if err != nil {
			return nil, err
		}

		if upper_ts != nil {
			return sjson.SetBytes(feature, "properties.date:cessation_upper", upper_ts.Unix())
		}

		return feature, nil
	}

	return sjson.SetBytes(feature, path, edtf.UNKNOWN)
}
