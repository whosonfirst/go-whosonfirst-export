package application

import (
	"flag"
)

type Application interface {
	Run(*flag.FlagSet) error
}
