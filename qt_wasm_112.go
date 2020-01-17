// +build js,wasm
// +build go1.12,!go1.13,!go1.14,!go1.15

package qt

import (
	"syscall/js"
	"unsafe"
)

func TypedArrayOf(src []byte) js.TypedArray { return js.TypedArrayOf(src) }

func ReleaseTypedArray(p unsafe.Pointer) { (*js.TypedArray)(p).Release() }
