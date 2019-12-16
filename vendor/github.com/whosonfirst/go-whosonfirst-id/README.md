# go-whosonfirst-id

Go package for generating valid Who's On First IDs.

## What is this?

This is a common Go package for generating valid Who's On First (WOF) identifiers. It implements the [go-artisanal-integers.Client](https://github.com/aaronland/go-artisanal-integers#client) interface for generating new IDs.

Under the hood it uses a [go-uid.Provider](https://github.com/aaronland/go-uid) for generating those IDs, specifically a [go-uid-artisanal](https://github.com/aaronland/go-uid-artisanal) provider.

This allows you to specify alternative and/or multiple artisanal integer providers (the default provider for WOF is [Brooklyn Integers](https://brooklynintegers.com/)) as well as a customizable pool of pre-generated and caches IDs using the [go-artisanal-integers-proxy](https://github.com/aaronland/go-artisanal-integers-proxy) and [go-pool](https://github.com/aaronland?utf8=%E2%9C%93&q=go-pool&type=&language=) packages.

Note: The use of the [go-uid.Provider](https://github.com/aaronland/go-uid) interface might be overkill. We'll see.

## Example

_Error handling omitted for the sake of brevity._

### Simple

```
package main

import (
	"context"
	"fmt"
	"testing"
)

func main() {

	ctx := context.Background()
	cl, _ := NewIdClient(ctx)

	id, _ := cl.NextInt()
	fmt.Println(id)
}
```

### Fancy

The default `IdClient` does not pre-generate or cache IDs. To do so create a new `IdClient` using the handy `NewIdClientWithURI` method:

```
package main

import (
	"context"
	"fmt"
	"testing"
	_ "github.com/aaronland/go-missionintegers-api"	
)

func main() {

	ctx := context.Background()

	uri := "artisanal:///?client=missionintegers%3A%2F%2F&minimum=5&pool=memory%3A%2F%2F"
	cl, _ := NewIdClientWithURI(ctx, uri)

	id, _ := cl.NextInt()
	fmt.Println(id)
}
```

This expects a valid [go-uid-artisanal](https://github.com/aaronland/go-uid-artisanal) URI string.

## See also

* https://github.com/aaronland/go-artisanal-integers
* https://github.com/aaronland/go-artisanal-integers-proxy
* https://github.com/aaronland/go-uid-artisanal
* https://github.com/aaronland/go-brooklynintegers-api