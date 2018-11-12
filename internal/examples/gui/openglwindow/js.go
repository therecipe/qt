// +build js

package main

import (
	"unsafe"

	"github.com/therecipe/qt"
)

var (
	colors = func() unsafe.Pointer {
		BYTES_PER_ELEMENT := 4 //Float32Array.BYTES_PER_ELEMENT
		data := []float32{1, 0, 0, 0, 1, 0, 0, 0, 1}

		alloc := qt.WASM.Call("_malloc", len(data)*BYTES_PER_ELEMENT, "float")
		for i, v := range data {
			qt.WASM.Call("setValue", alloc.Int()+(i*BYTES_PER_ELEMENT), v, "float")
		}
		return unsafe.Pointer(uintptr(alloc.Int()))
	}()

	vertices = func() unsafe.Pointer {
		BYTES_PER_ELEMENT := 4 //Float32Array.BYTES_PER_ELEMENT
		data := []float32{0, 0.707, -0.5, -0.5, 0.5, -0.5}

		alloc := qt.WASM.Call("_malloc", len(data)*BYTES_PER_ELEMENT, "float")
		for i, v := range data {
			qt.WASM.Call("setValue", alloc.Int()+(i*BYTES_PER_ELEMENT), v, "float")
		}
		return unsafe.Pointer(uintptr(alloc.Int()))
	}()
)
