//go:build js && wasm

package main

import (
	"syscall/js"
)

func main() {
	ch := make(chan struct{}, 0)
	js.Global().Set("list_connectors", js.FuncOf(list_connectors))
	js.Global().Get("console").Call("log", "WASM Initialized and Ready")
	<-ch
}
