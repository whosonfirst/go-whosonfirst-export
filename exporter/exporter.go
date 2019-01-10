package exporter

type Exporter interface {
	Export([]byte) ([]byte, error)
	ExportFeature(interface{}) ([]byte, error)
}
