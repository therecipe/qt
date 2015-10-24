package core

//#include "qsharedmemory.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QSharedMemory struct {
	QObject
}

type QSharedMemoryITF interface {
	QObjectITF
	QSharedMemoryPTR() *QSharedMemory
}

func PointerFromQSharedMemory(ptr QSharedMemoryITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSharedMemoryPTR().Pointer()
	}
	return nil
}

func QSharedMemoryFromPointer(ptr unsafe.Pointer) *QSharedMemory {
	var n = new(QSharedMemory)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QSharedMemory_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QSharedMemory) QSharedMemoryPTR() *QSharedMemory {
	return ptr
}

//QSharedMemory::AccessMode
type QSharedMemory__AccessMode int

var (
	QSharedMemory__ReadOnly  = QSharedMemory__AccessMode(0)
	QSharedMemory__ReadWrite = QSharedMemory__AccessMode(1)
)

//QSharedMemory::SharedMemoryError
type QSharedMemory__SharedMemoryError int

var (
	QSharedMemory__NoError          = QSharedMemory__SharedMemoryError(0)
	QSharedMemory__PermissionDenied = QSharedMemory__SharedMemoryError(1)
	QSharedMemory__InvalidSize      = QSharedMemory__SharedMemoryError(2)
	QSharedMemory__KeyError         = QSharedMemory__SharedMemoryError(3)
	QSharedMemory__AlreadyExists    = QSharedMemory__SharedMemoryError(4)
	QSharedMemory__NotFound         = QSharedMemory__SharedMemoryError(5)
	QSharedMemory__LockError        = QSharedMemory__SharedMemoryError(6)
	QSharedMemory__OutOfResources   = QSharedMemory__SharedMemoryError(7)
	QSharedMemory__UnknownError     = QSharedMemory__SharedMemoryError(8)
)

func NewQSharedMemory2(parent QObjectITF) *QSharedMemory {
	return QSharedMemoryFromPointer(unsafe.Pointer(C.QSharedMemory_NewQSharedMemory2(C.QtObjectPtr(PointerFromQObject(parent)))))
}

func NewQSharedMemory(key string, parent QObjectITF) *QSharedMemory {
	return QSharedMemoryFromPointer(unsafe.Pointer(C.QSharedMemory_NewQSharedMemory(C.CString(key), C.QtObjectPtr(PointerFromQObject(parent)))))
}

func (ptr *QSharedMemory) Attach(mode QSharedMemory__AccessMode) bool {
	if ptr.Pointer() != nil {
		return C.QSharedMemory_Attach(C.QtObjectPtr(ptr.Pointer()), C.int(mode)) != 0
	}
	return false
}

func (ptr *QSharedMemory) ConstData() {
	if ptr.Pointer() != nil {
		C.QSharedMemory_ConstData(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSharedMemory) Create(size int, mode QSharedMemory__AccessMode) bool {
	if ptr.Pointer() != nil {
		return C.QSharedMemory_Create(C.QtObjectPtr(ptr.Pointer()), C.int(size), C.int(mode)) != 0
	}
	return false
}

func (ptr *QSharedMemory) Data() {
	if ptr.Pointer() != nil {
		C.QSharedMemory_Data(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSharedMemory) Data2() {
	if ptr.Pointer() != nil {
		C.QSharedMemory_Data2(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QSharedMemory) Detach() bool {
	if ptr.Pointer() != nil {
		return C.QSharedMemory_Detach(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSharedMemory) Error() QSharedMemory__SharedMemoryError {
	if ptr.Pointer() != nil {
		return QSharedMemory__SharedMemoryError(C.QSharedMemory_Error(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSharedMemory) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSharedMemory_ErrorString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSharedMemory) IsAttached() bool {
	if ptr.Pointer() != nil {
		return C.QSharedMemory_IsAttached(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSharedMemory) Key() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSharedMemory_Key(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSharedMemory) Lock() bool {
	if ptr.Pointer() != nil {
		return C.QSharedMemory_Lock(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSharedMemory) NativeKey() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSharedMemory_NativeKey(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QSharedMemory) SetKey(key string) {
	if ptr.Pointer() != nil {
		C.QSharedMemory_SetKey(C.QtObjectPtr(ptr.Pointer()), C.CString(key))
	}
}

func (ptr *QSharedMemory) SetNativeKey(key string) {
	if ptr.Pointer() != nil {
		C.QSharedMemory_SetNativeKey(C.QtObjectPtr(ptr.Pointer()), C.CString(key))
	}
}

func (ptr *QSharedMemory) Size() int {
	if ptr.Pointer() != nil {
		return int(C.QSharedMemory_Size(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QSharedMemory) Unlock() bool {
	if ptr.Pointer() != nil {
		return C.QSharedMemory_Unlock(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSharedMemory) DestroyQSharedMemory() {
	if ptr.Pointer() != nil {
		C.QSharedMemory_DestroyQSharedMemory(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
