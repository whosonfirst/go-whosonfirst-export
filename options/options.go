package options

import (
	"github.com/whosonfirst/go-whosonfirst-export/uid"
)

type Options interface {
	UIDProvider() uid.Provider
}
