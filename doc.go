// go-whosonfirst-export is a Go package for exporting Who's On First documents in Go.
//
// Example
//
//	import (
//		"context"
//		"github.com/whosonfirst/go-whosonfirst-export/v2"
//		"io/ioutil"
//		"os
//	)
//	
//	func main() {
//	
//		ctx := context.Background()
//	
//		ex, _ := export.NewExporter(ctx, "whosonfirst://")
//		
//		path := "some.geojson"     	
//		fh, _ := os.Open(path)
//		defer fh.Close()
//	
//		body, _ := ioutil.ReadAll(fh)
//	
//		body, _ = ex.Export(ctx, body)
//		os.Stdout.Write(body)
//	}
package export
