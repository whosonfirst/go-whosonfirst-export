//go:build wasmjs
package wasm

import (
	"context"
        "fmt"
        "syscall/js"

        "github.com/whosonfirst/go-whosonfirst-export/v3"
	"github.com/whosonfirst/go-whosonfirst-feature/alt"	
)

func PrepareWithoutTimestampsFunc() js.Func {

        return js.FuncOf(func(this js.Value, args []js.Value) interface{} {

                feature_str := args[0].String()

                handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			ctx := context.Background()
			
                        resolve := args[0]
                        reject := args[1]

			feature := []byte(feature_str)
			
			is_alt := alt.IsAlt(feature)

			var feature_prepped []byte
			var err error
			
			if is_alt {
				feature_prepped, err = export.PrepareAltFeatureWithoutTimestamps(ctx, []byte(feature_str))				
			} else {
				feature_prepped, err = export.PrepareFeatureWithoutTimestamps(ctx, []byte(feature_str))
			}
			
                        if err != nil {
                                reject.Invoke(fmt.Printf("Failed to prepare feature, %v\n", err))
                                return nil
                        }

                        resolve.Invoke(string(feature_prepped))
                        return nil
                })

                promiseConstructor := js.Global().Get("Promise")
                return promiseConstructor.New(handler)
        })
}
