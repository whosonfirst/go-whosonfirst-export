package utils

import (
	_ "log"
)

// no, really...
// https://revdancatt.com/2012/08/23/london-artisan-integers-distribution-hotel-infinity-punk-an-excuse-explanation-of-sorts/

func IsLondonInteger(i int64) bool {

	test := i % 9

	if test == 0 {
		return true
	}

	return false
}
