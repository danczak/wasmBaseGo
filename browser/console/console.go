//go:build wasm
// +build wasm

package console

/*
	Basic wrapper for js console
*/

import "syscall/js"

// Print log message to the js console
func Log(fmt ...interface{}) {
	js.Global().Get("console").Call("log", fmt...)
}

// Print error message to the js console
func Error(fmt ...interface{}) {
	js.Global().Get("console").Call("error", fmt...)
}

// Print warning message to the js console
func Warn(fmt ...interface{}) {
	js.Global().Get("console").Call("warn", fmt...)
}
