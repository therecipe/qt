// +build !js

package main

import "unsafe"

var (
	colors   = unsafe.Pointer(&([]float32{1, 0, 0, 0, 1, 0, 0, 0, 1}[0]))
	vertices = unsafe.Pointer(&([]float32{0, 0.707, -0.5, -0.5, 0.5, -0.5}[0]))
)
