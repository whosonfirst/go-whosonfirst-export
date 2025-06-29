# go-whosonfirst-export

Go package for exporting Who's On First documents.

## What is this?

go-whosonfirst-export is a Go package for exporting Who's On First documents in Go. It is a port of the [py-mapzen-whosonfirst-geojson](https://github.com/whosonfirst/py-mapzen-whosonfirst-geojson) package and _mmmmmmmaybe_ some or all of the [py-mapzen-whosonfirst-export](https://github.com/whosonfirst/py-mapzen-whosonfirst-geojson) package.


## Documentation

[![Go Reference](https://pkg.go.dev/badge/github.com/whosonfirst/go-whosonfirst-export.svg)](https://pkg.go.dev/github.com/whosonfirst/go-whosonfirst-export)

## Example

_All error handling removed for the sake of brevity._

### Simple

```
import (
	"context"
	"os

	"github.com/whosonfirst/go-whosonfirst-export/v3"
)

func main() {

	ctx := context.Background()
	body, _ := os.ReadFile(path)

	has_changed, new_body, _ := export.Export(ctx, body)
}
```

This is how you would have done the same thing using the `/v2` release:

```
import (
	"context"
	"os

	"github.com/whosonfirst/go-whosonfirst-export/v2"
)

func main() {

	ctx := context.Background()

	body, _ := os.ReadFile(path)
	opts, _ := export.NewDefaultOptions(ctx)
	
	export.Export(body, opts, os.Stdout)
}
```

### Exporter

```
import (
	"context"
	"os

	"github.com/whosonfirst/go-whosonfirst-export/v3"
)

func main() {

	ctx := context.Background()
	ex, _ := export.NewExporter(ctx, "whosonfirst://")
	
	path := "some.geojson"     	
	body, _ := os.ReadFile(path)

	has_changed, body, _ = ex.Export(ctx, body)
	os.Stdout.Write(body)
}
```

...

```
import (
	"context"
	"os

	"github.com/whosonfirst/go-whosonfirst-export/v2"
)

func main() {

	ctx := context.Background()
	ex, _ := export.NewExporter(ctx, "whosonfirst://")
	
	path := "some.geojson"     	
	body, _ := os.ReadFile(path)

	body, _ = ex.Export(ctx, body)
	os.Stdout.Write(body)
}
```

## Interfaces

### Exporter

```
type Exporter interface {
	Export(context.Context, []byte) (bool, []byte, error)
}
```

## To do

This package needs to hold hands with the `go-whosonfirst-validate` package.

## See also

* https://github.com/tidwall/pretty
* https://github.com/tidwall/gjson
* https://github.com/tidwall/pretty/issues/2
* https://gist.github.com/tidwall/ca6ca1dd0cb780f0be4d134f8e4eb7bc
* https://github.com/whosonfirst/go-whosonfirst-validate