package dart

import "C"
import (
	"unsafe"
)

//export asyncCallIntoRemoteRegister
func asyncCallIntoRemoteRegister(port int64, fun unsafe.Pointer) {
	_asyncCallIntoRemoteRegister(port, fun)
}
