# go-whosonfirst-export

Go package for exporting Who's On First documents.

## Important

Too soon. Move along.

## What is this?

This is an experimental package to format Who's On First documents in Go. It is meant to be a port of the [py-mapzen-whosonfirst-geojson](https://github.com/whosonfirst/py-mapzen-whosonfirst-geojson) package and _mmmmmmmaybe_ some or all of the [py-mapzen-whosonfirst-export](https://github.com/whosonfirst/py-mapzen-whosonfirst-geojson) package.

It is also in flux and you should assume anything you see or read now _will_ change.

## Example

### Simple

```
import (
	"github.com/whosonfirst/go-whosonfirst-export"
	"github.com/whosonfirst/go-whosonfirst-export/options"	
	"io/ioutil"
	"os
)

func main() {

	path := "some.geojson"     	
	fh, _ := os.Open(path)
	defer fh.Close()

	body, _ := ioutil.ReadAll(fh)

	opts, _ := options.NewDefaultOptions()
	export.Export(body, opts, os.Stdout)
}
```

_Error handling removed for the sake of brevity._

## See also

* https://github.com/tidwall/pretty
* https://github.com/tidwall/gjson
* https://github.com/tidwall/pretty/issues/2
* https://gist.github.com/tidwall/ca6ca1dd0cb780f0be4d134f8e4eb7bc