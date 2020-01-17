// +build js,wasm

package qt

import (
	"syscall/js"
	"unsafe"
)

func init() {
	Module.Set("_callbackReleaseTypedArray", js.FuncOf(func(_ js.Value, args []js.Value) interface{} {
		ReleaseTypedArray(unsafe.Pointer(uintptr(args[0].Int())))
		return nil
	}))
}

var Global = js.Global()
var Module = Global.Call("eval", "Module")

//TODO: func MakeWrapper(i interface{}) *js.Value
