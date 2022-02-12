//go:build wasm
// +build wasm

package main

import (
	"console"
	"syscall/js"
)

func init() {
	console.Log("Hello Go Wasm!")
}

func main() {

	js.Global().Get("document").Call("write", "Hello world!")

	<-(make(chan bool))
}
