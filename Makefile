tools:
	go build -o bin/wof-export-feature cmd/wof-export-feature/main.go

wasi:
	tinygo build -wasm-abi=generic -target=wasi -o export.wasm cmd/wof-export-feature-wasi/main.go
