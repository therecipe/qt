package dart

/*
#include "stdlib.h"
#include "dart_native_api.h"

typedef const bool (*asyncCallIntoRemote) (Dart_Port, Dart_CObject*);
void postSomething(asyncCallIntoRemote postFunc, int64_t send_port, char* work_addr) {
	Dart_CObject dart_object;
	dart_object.type = Dart_CObject_kString;
	dart_object.value.as_string = work_addr;
	postFunc(send_port, &dart_object);
}

*/
import "C"
import (
	"unsafe"

	"github.com/StarAurryon/qt/interop"
)

var (
	asyncCallIntoRemote C.asyncCallIntoRemote
	port                int64
)

func init() {
	interop.ReturnPointersAsStrings = false

	interop.AsyncCallIntoRemote = func(s string) {
		i := C.CString(s)
		C.postSomething(asyncCallIntoRemote, C.int64_t(port), i)
		C.free(unsafe.Pointer(i))
	}
}

//

func _asyncCallIntoRemoteRegister(p int64, fp unsafe.Pointer) {
	port = p
	asyncCallIntoRemote = C.asyncCallIntoRemote(fp)
}
