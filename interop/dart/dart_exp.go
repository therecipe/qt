package dart

import "C"
import (
	"unsafe"
)

//export syncCallIntoRemoteRegister
func syncCallIntoRemoteRegister(fun unsafe.Pointer) {
	_syncCallIntoRemoteRegister(fun)
}

//export asyncCallIntoRemoteRegister
func asyncCallIntoRemoteRegister(port int64, fun unsafe.Pointer) {
	_asyncCallIntoRemoteRegister(port, fun)
}
