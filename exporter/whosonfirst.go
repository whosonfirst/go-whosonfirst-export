package exporter

import (
	"encoding/json"
	"github.com/whosonfirst/go-whosonfirst-export"
	"github.com/whosonfirst/go-whosonfirst-export/options"
)

type WhosOnFirstExporter struct {
	Exporter
	options options.Options
}

func NewWhosOnFirstExporter(opts options.Options) (Exporter, error) {

	ex := WhosOnFirstExporter{
		options: opts,
	}

	return &ex, nil
}

func (ex *WhosOnFirstExporter) ExportFeature(feature interface{}) ([]byte, error) {

	body, err := json.Marshal(feature)

	if err != nil {
		return nil, err
	}

	return ex.Export(body)
}

func (ex *WhosOnFirstExporter) Export(feature []byte) ([]byte, error) {

	var err error

	feature, err = export.Prepare(feature, ex.options)

	if err != nil {
		return nil, err
	}

	feature, err = export.Format(feature, ex.options)

	if err != nil {
		return nil, err
	}

	return feature, nil
}
