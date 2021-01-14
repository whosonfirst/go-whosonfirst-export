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

func TestExportEDTF(t *testing.T) {
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
		"properties.date:inception_lower",
		"properties.date:inception_upper",
		"properties.date:cessation_lower",
		"properties.date:cessation_upper",
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

func TestExportWithMissingBelongstoElement(t *testing.T) {
	body := readFeature(t, "missing-belongsto-element.geojson")
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
	updatedBody := buf.Bytes()

	if bytes.Equal(body, updatedBody) {
		t.Error("Body was identical")
	}

	newBelongsto := gjson.GetBytes(updatedBody, "properties.wof:belongsto").Array()
	if len(newBelongsto) != 6 {
		t.Error("belongsto has incorrect number of elements")
	}

	lastBelongsto := newBelongsto[len(newBelongsto)-1].Int()
	if lastBelongsto != 404227469 {
		t.Error("belongsto has incorrect added element")
	}
}

func TestExportWithMissingDateDerived(t *testing.T) {
	body := readFeature(t, "missing-date-derived.geojson")
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
	updatedBody := buf.Bytes()

	if bytes.Equal(body, updatedBody) {
		t.Error("Body was identical")
	}

	ensureProps := []string{
		"properties.date:inception_lower",
		"properties.date:inception_upper",
		"properties.date:cessation_lower",
		"properties.date:cessation_upper",
	}

	for _, prop := range ensureProps {
		propRsp := gjson.GetBytes(body, prop)

		if !propRsp.Exists() {
			t.Fatalf("Missing property '%s'", prop)
		}

		fmt.Printf("%s: %s\n", prop, propRsp.String())
	}

	inceptionLowerRsp := gjson.GetBytes(body, "properties.date:inception_lower")
	cessationUpperRsp := gjson.GetBytes(body, "properties.date:cessation_upper")

	inceptionLowerStr := inceptionLowerRsp.String()
	cessationUpperStr := cessationUpperRsp.String()

	fmt.Println("INCEPTION", inceptionLowerStr)
	fmt.Println("CESSATION", cessationUpperStr)
}

func TestExportWithExtraBelongstoElement(t *testing.T) {
	body := readFeature(t, "extra-belongsto-element.geojson")
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
	updatedBody := buf.Bytes()

	if bytes.Equal(body, updatedBody) {
		t.Error("Body was identical")
	}

	newBelongsto := gjson.GetBytes(updatedBody, "properties.wof:belongsto").Array()
	if len(newBelongsto) != 6 {
		t.Error("belongsto has incorrect number of elements")
	}

	lastBelongsto := newBelongsto[len(newBelongsto)-1].Int()
	if lastBelongsto != 1360698877 {
		t.Error("belongsto has incorrect added element")
	}
}

func TestExportWithMissingBelongstoKey(t *testing.T) {
	body := readFeature(t, "missing-belongsto-key.geojson")
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
	updatedBody := buf.Bytes()

	if bytes.Equal(body, updatedBody) {
		t.Error("Body was identical")
	}

	newBelongsto := gjson.GetBytes(updatedBody, "properties.wof:belongsto").Array()
	if len(newBelongsto) != 6 {
		t.Error("belongsto has incorrect number of elements")
	}
}

func TestExportChangedWithUnchangedFile(t *testing.T) {
	body := readFeature(t, "no-changes.geojson")
	opts, err := options.NewDefaultOptions()
	if err != nil {
		t.Fatal(err)
	}

	var buf bytes.Buffer
	wr := bufio.NewWriter(&buf)

	changed, err := ExportChanged(body, body, opts, wr)
	if err != nil {
		t.Fatal(err)
	}

	if changed {
		t.Error("Exported file should not be changed")
	}

	wr.Flush()
	updatedBody := buf.Bytes()

	if len(updatedBody) > 0 {
		t.Error("Writer should not have written to file")
	}
}

func TestExportChangedWithChanges(t *testing.T) {
	body := readFeature(t, "changes-required.geojson")
	opts, err := options.NewDefaultOptions()
	if err != nil {
		t.Fatal(err)
	}

	originalLastModified := gjson.GetBytes(body, "properties.wof:lastmodified").Int()

	var buf bytes.Buffer
	wr := bufio.NewWriter(&buf)

	changed, err := ExportChanged(body, body, opts, wr)
	if err != nil {
		t.Fatal(err)
	}

	if !changed {
		t.Error("Exported file should have be changed")
	}

	wr.Flush()
	updatedBody := buf.Bytes()

	if bytes.Equal(body, updatedBody) {
		t.Error("Body was identical")
	}

	newLastModified := gjson.GetBytes(updatedBody, "properties.wof:lastmodified").Int()

	if newLastModified <= originalLastModified {
		t.Error("Last modified timestamp should have increased")
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
