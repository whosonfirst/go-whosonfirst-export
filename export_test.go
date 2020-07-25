package export

import (
	"bufio"
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"github.com/tidwall/gjson"
	"github.com/whosonfirst/go-whosonfirst-export/options"
)

func TestExport(t *testing.T) {
	body := readFeature(t, "1159159407.geojson")

	opts, err := options.NewDefaultOptions()
	if err != nil {
		t.Fatal(err)
	}

	var buf bytes.Buffer
	wr := bufio.NewWriter(&buf)

	err = Export(body, opts, wr)
	if err != nil {
		t.Fatal(err)
	}

	wr.Flush()
	body = buf.Bytes()

	ensureProps := []string{
		"properties.wof:id",
		"properties.geom:bbox",
		"bbox",
	}

	for _, prop := range ensureProps {

		propRsp := gjson.GetBytes(body, prop)

		if !propRsp.Exists() {
			t.Fatalf("Missing property '%s'", prop)
		}

		fmt.Printf("%s: %s\n", prop, propRsp.String())
	}

	bboxRsp := gjson.GetBytes(body, "properties.geom:bbox")
	bboxStr := bboxRsp.String()

	if bboxStr != "-122.384119,37.615457,-122.384119,37.615457" {
		t.Fatal("Unexpected geom:bbox")
	}

}

func readFeature(t *testing.T, filename string) []byte {
	cwd, err := os.Getwd()

	if err != nil {
		t.Fatal(err)
	}

	fixtures := filepath.Join(cwd, "fixtures")
	featurePath := filepath.Join(fixtures, filename)

	fh, err := os.Open(featurePath)
	defer fh.Close()

	if err != nil {
		t.Fatal(err)
	}

	body, err := ioutil.ReadAll(fh)

	if err != nil {
		t.Fatal(err)
	}

	return body
}
