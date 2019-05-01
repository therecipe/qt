// +build android android_emulator

package main

import (
	"C"

	"unsafe"
)

//export Java_org_ftylitak_qzxing_NativeFunctions_onPermissionsGranted
func Java_org_ftylitak_qzxing_NativeFunctions_onPermissionsGranted(_, _ unsafe.Pointer) {
	println("permission granted callback")
	Application.OnPermissionsGranted()
}

//export Java_org_ftylitak_qzxing_NativeFunctions_onPermissionsDenied
func Java_org_ftylitak_qzxing_NativeFunctions_onPermissionsDenied(_, _ unsafe.Pointer) {
	println("permission denied callback")
	Application.OnPermissionsDenied()
}
