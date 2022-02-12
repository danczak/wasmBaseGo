GOROOT = $(shell go env GOROOT)

browser: *.go
	cp $(GOROOT)/misc/wasm/wasm_exec.js ../wwwroot/js/
	GOOS=js GOARCH=wasm go build -ldflags="-s -w" -o ../wwwroot/main.wasm

clean:
	rm -f ../wwwroot/js/wasm_exec.js
	rm -f ../wwwroot/main.wasm
