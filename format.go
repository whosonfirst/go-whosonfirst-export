package export

import (
	"github.com/whosonfirst/go-whosonfirst-format"
)

func Format(body []byte) ([]byte, error) {
	return format.FormatBytes(body)
}
