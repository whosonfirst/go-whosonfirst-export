# go-whosonfirst-export

Go package for exporting Who's On First documents.

## Documentation

[![Go Reference](https://pkg.go.dev/badge/github.com/whosonfirst/go-whosonfirst-export.svg)](https://pkg.go.dev/github.com/whosonfirst/go-whosonfirst-export)

## Example

Version 3.x of this package introduce major, backward-incompatible changes from earlier releases. That said, migragting from version 2.x to 3.x should be relatively straightforward as a the _basic_ concepts are still the same but (hopefully) simplified. There are some important changes "under the hood" but the user-facing changes, while important, should be easy to update.

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

	if has_changes {
		os.Stdout.Write(new_body)
	}
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

The `export.Export` method is really just a convenience around the default Who's On First exporter package which implements the `export.Exporter` interface (described below). The goal behind the interface it to allow for custom exporters which can supplement the default export functionality with application-specific needs. To use the exporter package directly you would do this:

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

	has_changes, body, _ = ex.Export(ctx, body)

	if has_changes {
		os.Stdout.Write(body)
	}
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

* https://github.com/whosonfirst/go-whosonfirst-format
* https://github.com/whosonfirst/go-whosonfirst-validate
* https://github.com/tidwall/gjson
* https://github.com/tidwall/sjson
