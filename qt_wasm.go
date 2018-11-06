// +build js,wasm

package qt

import (
	"syscall/js"
	"unsafe"
)

var WASM = js.Global().Call("eval", "Module")

func init() {
	WASM.Set("_callbackReleaseTypedArray", js.NewCallback(func(_ js.Value, args []js.Value) interface{} {
		(*js.TypedArray)(unsafe.Pointer(uintptr(args[0].Int()))).Release()
		return nil
	}))
}
