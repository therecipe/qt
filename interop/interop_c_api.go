package interop

import "C"

//export syncCallIntoLocal
func syncCallIntoLocal(i *C.char) *C.char {
	return C.CString(_syncCallIntoLocal(C.GoString(i)))
}

//export asyncCallIntoRemoteResponse
func asyncCallIntoRemoteResponse(i *C.char) {
	_asyncCallIntoRemoteResponse(C.GoString(i))
}
