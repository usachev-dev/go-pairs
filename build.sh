GOOS=js GOARCH=wasm go build -o dist/main.wasm ./entrypoints/wasm-client/main.go && cp -rf ./assets/ ./dist
