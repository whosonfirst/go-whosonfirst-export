package export

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"testing"

	"github.com/tidwall/gjson"
	wof_properties "github.com/whosonfirst/go-whosonfirst-feature/properties"
)

func TestExportAlt(t *testing.T) {

	ctx := context.Background()

	body := readFeature(t, "101736545-alt-quattroshapes.geojson")

	has_changed, _, err := Export(ctx, body)

	if err != nil {
		t.Fatal(err)
	}

	if has_changed {
		t.Fatal("Did not expect alt file to change")
	}
}

func TestHasChanges(t *testing.T) {

	ctx := context.Background()

	body := readFeature(t, "no-changes.geojson")

	has_changes, err := HasChanges(ctx, body, body)

	if err != nil {
		t.Fatal(err)
	}

	if has_changes {
		t.Fatal("Did not expect file to have changes")
	}

	other_body := readFeature(t, "custom-placetype.geojson")

	has_changes, err = HasChanges(ctx, body, other_body)

	if err != nil {
		t.Fatal(err)
	}

	if !has_changes {
		t.Fatal("Expected files to have changes")
	}

}

func TestCustomPlacetype(t *testing.T) {

	ctx := context.Background()

	body := readFeature(t, "custom-placetype.geojson")

	_, new_body, err := Export(ctx, body)

	if err != nil {
		t.Fatal(err)
	}

	path := fmt.Sprintf("%s.0.runway_id", wof_properties.PATH_WOF_HIERARCHY)

	rsp := gjson.GetBytes(new_body, path)

	if !rsp.Exists() {
		t.Fatalf("Unable to find %s property", path)
	}

	has_id := rsp.Int()
	expected_id := int64(1730008747)

	if has_id != expected_id {
		t.Fatalf("Result has unexpected ID. Expected '%d' but got '%d'", expected_id, has_id)
	}
}

func TestExportEDTF(t *testing.T) {

	ctx := context.Background()

	body := readFeature(t, "1159159407.geojson")

	_, new_body, err := Export(ctx, body)

	if err != nil {
		t.Fatal(err)
	}

	ensureProps := []string{
		wof_properties.PATH_WOF_ID,
		wof_properties.PATH_GEOM_BBOX,
		wof_properties.PATH_BBOX,
		wof_properties.PATH_DATE_INCEPTION_LOWER,
		wof_properties.PATH_DATE_INCEPTION_UPPER,
		wof_properties.PATH_DATE_CESSATION_LOWER,
		wof_properties.PATH_DATE_CESSATION_UPPER,
	}

	for _, prop := range ensureProps {
		propRsp := gjson.GetBytes(new_body, prop)

		if !propRsp.Exists() {
			t.Fatalf("Missing property '%s'", prop)
		}
	}

	bboxRsp := gjson.GetBytes(new_body, wof_properties.PATH_GEOM_BBOX)
	bboxStr := bboxRsp.String()

	if bboxStr != "-122.384119,37.615457,-122.384119,37.615457" {
		t.Fatal("Unexpected geom:bbox")
	}
}

func TestExportWithOldStyleEDTFUnknownDates(t *testing.T) {

	ctx := context.Background()
	body := readFeature(t, "old-edtf-uuuu-dates.geojson")

	_, new_body, err := Export(ctx, body)

	if err != nil {
		t.Fatal(err)
	}

	cessationProp := gjson.GetBytes(new_body, wof_properties.PATH_EDTF_CESSATION)

	if !cessationProp.Exists() {
		t.Fatalf("missing edtf:cessation property")
	}

	if cessationProp.String() != "" {
		t.Fatalf("edtf:cessation not set to new style format")
	}

	inceptionProp := gjson.GetBytes(new_body, wof_properties.PATH_EDTF_INCEPTION)

	if !inceptionProp.Exists() {
		t.Fatalf("missing edtf:inception property")
	}

	if inceptionProp.String() != "" {
		t.Fatalf("edtf:inception not set to new style format")
	}

	rejectProps := []string{
		wof_properties.PATH_DATE_INCEPTION_LOWER,
		wof_properties.PATH_DATE_INCEPTION_UPPER,
		wof_properties.PATH_DATE_CESSATION_LOWER,
		wof_properties.PATH_DATE_CESSATION_UPPER,
	}

	for _, prop := range rejectProps {
		propRsp := gjson.GetBytes(new_body, prop)

		if propRsp.Exists() {
			t.Fatalf("unexpected property '%s'", prop)
		}
	}
}

func TestMissingUpperLowerDates(t *testing.T) {

	ctx := context.Background()
	body := readFeature(t, "missing-upper-lower-dates.geojson")

	_, new_body, err := Export(ctx, body)

	if err != nil {
		t.Fatal(err)
	}

	cessationProp := gjson.GetBytes(new_body, wof_properties.PATH_EDTF_CESSATION)

	if !cessationProp.Exists() {
		t.Fatalf("missing edtf:cessation property")
	}

	inceptionProp := gjson.GetBytes(new_body, wof_properties.PATH_EDTF_INCEPTION)

	if !inceptionProp.Exists() {
		t.Fatalf("missing edtf:inception property")
	}

	requiredProps := []string{
		wof_properties.PATH_DATE_INCEPTION_LOWER,
		wof_properties.PATH_DATE_INCEPTION_UPPER,
		wof_properties.PATH_DATE_CESSATION_LOWER,
		wof_properties.PATH_DATE_CESSATION_UPPER,
	}

	for _, prop := range requiredProps {
		propRsp := gjson.GetBytes(new_body, prop)

		if !propRsp.Exists() {
			t.Fatalf("missing property '%s'", prop)
		}
	}
}

func TestExportWithMissingBelongstoElement(t *testing.T) {

	ctx := context.Background()

	body := readFeature(t, "missing-belongsto-element.geojson")

	_, new_body, err := Export(ctx, body)

	if err != nil {
		t.Fatal(err)
	}

	if bytes.Equal(body, new_body) {
		t.Error("Body was identical")
	}

	newBelongsto := gjson.GetBytes(new_body, wof_properties.PATH_WOF_BELONGSTO).Array()

	if len(newBelongsto) != 6 {
		t.Error("belongsto has incorrect number of elements")
	}

	lastBelongsto := newBelongsto[len(newBelongsto)-1].Int()

	if lastBelongsto != 404227469 {
		t.Error("belongsto has incorrect added element")
	}
}

func TestExportWithMissingDateDerived(t *testing.T) {

	ctx := context.Background()

	body := readFeature(t, "missing-date-derived.geojson")

	_, new_body, err := Export(ctx, body)

	if err != nil {
		t.Fatal(err)
	}

	if bytes.Equal(body, new_body) {
		t.Error("Body was identical")
	}

	ensureProps := []string{
		wof_properties.PATH_DATE_INCEPTION_LOWER,
		wof_properties.PATH_DATE_INCEPTION_UPPER,
		wof_properties.PATH_DATE_CESSATION_LOWER,
		wof_properties.PATH_DATE_CESSATION_UPPER,
	}

	for _, prop := range ensureProps {
		propRsp := gjson.GetBytes(new_body, prop)

		if !propRsp.Exists() {
			t.Fatalf("Missing property '%s'", prop)
		}
	}

	inceptionLowerRsp := gjson.GetBytes(new_body, wof_properties.PATH_DATE_INCEPTION_LOWER)
	cessationUpperRsp := gjson.GetBytes(new_body, wof_properties.PATH_DATE_CESSATION_UPPER)

	inceptionLowerStr := inceptionLowerRsp.String()
	cessationUpperStr := cessationUpperRsp.String()

	inceptionExpectedStr := "1996-07-01"
	cessationExpectedStr := "1997-02-10"

	if inceptionLowerStr != inceptionExpectedStr {
		t.Fatalf("Invalid date:inception_lower. Expected '%s' but got '%s'", inceptionExpectedStr, inceptionLowerStr)
	}

	if cessationUpperStr != cessationExpectedStr {
		t.Fatalf("Invalid date:cessation_upper. Expected '%s' but got '%s'", cessationExpectedStr, cessationUpperStr)
	}
}

func TestExportWithExtraBelongstoElement(t *testing.T) {

	ctx := context.Background()

	body := readFeature(t, "extra-belongsto-element.geojson")

	_, new_body, err := Export(ctx, body)

	if err != nil {
		t.Fatal(err)
	}

	if bytes.Equal(body, new_body) {
		t.Error("Body was identical")
	}

	newBelongsto := gjson.GetBytes(new_body, wof_properties.PATH_WOF_BELONGSTO).Array()

	if len(newBelongsto) != 6 {
		t.Error("belongsto has incorrect number of elements")
	}

	lastBelongsto := newBelongsto[len(newBelongsto)-1].Int()

	if lastBelongsto != 1360698877 {
		t.Error("belongsto has incorrect added element")
	}
}

func TestExportWithMissingBelongstoKey(t *testing.T) {

	ctx := context.Background()

	body := readFeature(t, "missing-belongsto-key.geojson")

	_, new_body, err := Export(ctx, body)

	if err != nil {
		t.Fatal(err)
	}

	if bytes.Equal(body, new_body) {
		t.Error("Body was identical")
	}

	newBelongsto := gjson.GetBytes(new_body, wof_properties.PATH_WOF_BELONGSTO).Array()

	if len(newBelongsto) != 6 {
		t.Error("belongsto has incorrect number of elements")
	}
}

func TestExportChangedWithUnchangedFile(t *testing.T) {

	ctx := context.Background()

	body := readFeature(t, "no-changes.geojson")

	changed, _, err := Export(ctx, body)

	if err != nil {
		t.Fatal(err)
	}

	if changed {
		t.Error("Exported file should not be changed")
	}
}

func TestExportChangedWithChanges(t *testing.T) {

	ctx := context.Background()

	body := readFeature(t, "changes-required.geojson")

	originalLastModified := gjson.GetBytes(body, wof_properties.PATH_WOF_LASTMODIFIED).Int()

	changed, new_body, err := Export(ctx, body)

	if err != nil {
		t.Fatal(err)
	}

	if !changed {
		t.Error("Exported file should have be changed")
	}

	if bytes.Equal(body, new_body) {
		t.Error("Body was identical")
	}

	newLastModified := gjson.GetBytes(new_body, wof_properties.PATH_WOF_LASTMODIFIED).Int()

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
	if err != nil {
		t.Fatal(err)
	}

	defer fh.Close()

	body, err := io.ReadAll(fh)

	if err != nil {
		t.Fatal(err)
	}

	return body
}
