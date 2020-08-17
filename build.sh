GOOS=js GOARCH=wasm go build -o docs/main.wasm ./entrypoints/wasm-client/main.go && cp -rf ./assets/ ./docs
