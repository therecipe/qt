package dart

/*
#include "stdlib.h"
#include "dart_native_api.h"

typedef char* (*syncCallIntoRemote) (char*);
char* postSync(syncCallIntoRemote postFunc, char* work_addr){
	return postFunc(work_addr);
}

typedef const bool (*asyncCallIntoRemote) (Dart_Port, Dart_CObject*);
void postSomething(asyncCallIntoRemote postFunc, int64_t send_port, char* work_addr){
	Dart_CObject dart_object;
	dart_object.type = Dart_CObject_kString;
	dart_object.value.as_string = work_addr;
	postFunc(send_port, &dart_object);
}

*/
import "C"
import (
	"unsafe"

	"github.com/therecipe/qt/interop"
)

var (
	syncCallIntoRemote  C.syncCallIntoRemote
	asyncCallIntoRemote C.asyncCallIntoRemote
	port                int64
)

func init() {
	interop.SyncCallIntoRemote = func(s string) string {
		i := C.CString(s)
		o := C.postSync(syncCallIntoRemote, i)

		ret := C.GoString(o)

		C.free(unsafe.Pointer(i))
		C.free(unsafe.Pointer(o))
		return ret
	}

	interop.AsyncCallIntoRemote = func(s string) {
		i := C.CString(s)
		C.postSomething(asyncCallIntoRemote, C.int64_t(port), i)
		C.free(unsafe.Pointer(i))
	}
}

//

func _syncCallIntoRemoteRegister(fp unsafe.Pointer) {
	syncCallIntoRemote = C.syncCallIntoRemote(fp)
}

func _asyncCallIntoRemoteRegister(p int64, fp unsafe.Pointer) {
	port = p
	asyncCallIntoRemote = C.asyncCallIntoRemote(fp)
}
