// +build js,wasm
// +build go1.13

package qt

import (
	"syscall/js"
	"unsafe"
)

func TypedArrayOf(src []byte) js.Value {
	dst := Global.Get("Uint8Array").New(len(src))
	js.CopyBytesToJS(dst, src)
	return dst
}

func ReleaseTypedArray(unsafe.Pointer) {}
