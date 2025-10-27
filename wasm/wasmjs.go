//go:build wasmjs
package wasm

import (
        "fmt"
        "syscall/js"

        "github.com/whosonfirst/go-whosonfirst-export/v3"
)

func PrepareFeatureFunc() js.Func {

        return js.FuncOf(func(this js.Value, args []js.Value) interface{} {

                feature_str := args[0].String()

                handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

                        resolve := args[0]
                        reject := args[1]

                        feature_prepped, err := export.PrepareFeatureWithTimestamps([]byte(feature_str))

                        if err != nil {
                                reject.Invoke(fmt.Printf("Failed to prepare feature, %v\n", err))
                                return nil
                        }

                        resolve.Invoke(string(feature_fmt))
                        return nil
                })

                promiseConstructor := js.Global().Get("Promise")
                return promiseConstructor.New(handler)
        })
}

func PrepareAltFeatureFunc() js.Func {

        return js.FuncOf(func(this js.Value, args []js.Value) interface{} {

                feature_str := args[0].String()

                handler := js.FuncOf(func(this js.Value, args []js.Value) interface{} {

                        resolve := args[0]
                        reject := args[1]

                        feature_prepped, err := export.PrepareAltFeatureWithTimestamps([]byte(feature_str))

                        if err != nil {
                                reject.Invoke(fmt.Printf("Failed to prepare alt feature, %v\n", err))
                                return nil
                        }

                        resolve.Invoke(string(feature_fmt))
                        return nil
                })

                promiseConstructor := js.Global().Get("Promise")
                return promiseConstructor.New(handler)
        })
}
