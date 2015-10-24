package core

//#include "qmetamethod.h"
import "C"
import (
	"unsafe"
)

type QMetaMethod struct {
	ptr unsafe.Pointer
}

type QMetaMethodITF interface {
	QMetaMethodPTR() *QMetaMethod
}

func (p *QMetaMethod) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMetaMethod) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMetaMethod(ptr QMetaMethodITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMetaMethodPTR().Pointer()
	}
	return nil
}

func QMetaMethodFromPointer(ptr unsafe.Pointer) *QMetaMethod {
	var n = new(QMetaMethod)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMetaMethod) QMetaMethodPTR() *QMetaMethod {
	return ptr
}

//QMetaMethod::Access
type QMetaMethod__Access int

var (
	QMetaMethod__Private   = QMetaMethod__Access(0)
	QMetaMethod__Protected = QMetaMethod__Access(1)
	QMetaMethod__Public    = QMetaMethod__Access(2)
)

//QMetaMethod::MethodType
type QMetaMethod__MethodType int

var (
	QMetaMethod__Method      = QMetaMethod__MethodType(0)
	QMetaMethod__Signal      = QMetaMethod__MethodType(1)
	QMetaMethod__Slot        = QMetaMethod__MethodType(2)
	QMetaMethod__Constructor = QMetaMethod__MethodType(3)
)

func (ptr *QMetaMethod) Access() QMetaMethod__Access {
	if ptr.Pointer() != nil {
		return QMetaMethod__Access(C.QMetaMethod_Access(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMetaMethod) Invoke4(object QObjectITF, val0 QGenericArgumentITF, val1 QGenericArgumentITF, val2 QGenericArgumentITF, val3 QGenericArgumentITF, val4 QGenericArgumentITF, val5 QGenericArgumentITF, val6 QGenericArgumentITF, val7 QGenericArgumentITF, val8 QGenericArgumentITF, val9 QGenericArgumentITF) bool {
	if ptr.Pointer() != nil {
		return C.QMetaMethod_Invoke4(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQObject(object)), C.QtObjectPtr(PointerFromQGenericArgument(val0)), C.QtObjectPtr(PointerFromQGenericArgument(val1)), C.QtObjectPtr(PointerFromQGenericArgument(val2)), C.QtObjectPtr(PointerFromQGenericArgument(val3)), C.QtObjectPtr(PointerFromQGenericArgument(val4)), C.QtObjectPtr(PointerFromQGenericArgument(val5)), C.QtObjectPtr(PointerFromQGenericArgument(val6)), C.QtObjectPtr(PointerFromQGenericArgument(val7)), C.QtObjectPtr(PointerFromQGenericArgument(val8)), C.QtObjectPtr(PointerFromQGenericArgument(val9))) != 0
	}
	return false
}

func (ptr *QMetaMethod) Invoke2(object QObjectITF, returnValue QGenericReturnArgumentITF, val0 QGenericArgumentITF, val1 QGenericArgumentITF, val2 QGenericArgumentITF, val3 QGenericArgumentITF, val4 QGenericArgumentITF, val5 QGenericArgumentITF, val6 QGenericArgumentITF, val7 QGenericArgumentITF, val8 QGenericArgumentITF, val9 QGenericArgumentITF) bool {
	if ptr.Pointer() != nil {
		return C.QMetaMethod_Invoke2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQObject(object)), C.QtObjectPtr(PointerFromQGenericReturnArgument(returnValue)), C.QtObjectPtr(PointerFromQGenericArgument(val0)), C.QtObjectPtr(PointerFromQGenericArgument(val1)), C.QtObjectPtr(PointerFromQGenericArgument(val2)), C.QtObjectPtr(PointerFromQGenericArgument(val3)), C.QtObjectPtr(PointerFromQGenericArgument(val4)), C.QtObjectPtr(PointerFromQGenericArgument(val5)), C.QtObjectPtr(PointerFromQGenericArgument(val6)), C.QtObjectPtr(PointerFromQGenericArgument(val7)), C.QtObjectPtr(PointerFromQGenericArgument(val8)), C.QtObjectPtr(PointerFromQGenericArgument(val9))) != 0
	}
	return false
}

func (ptr *QMetaMethod) Invoke3(object QObjectITF, connectionType Qt__ConnectionType, val0 QGenericArgumentITF, val1 QGenericArgumentITF, val2 QGenericArgumentITF, val3 QGenericArgumentITF, val4 QGenericArgumentITF, val5 QGenericArgumentITF, val6 QGenericArgumentITF, val7 QGenericArgumentITF, val8 QGenericArgumentITF, val9 QGenericArgumentITF) bool {
	if ptr.Pointer() != nil {
		return C.QMetaMethod_Invoke3(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQObject(object)), C.int(connectionType), C.QtObjectPtr(PointerFromQGenericArgument(val0)), C.QtObjectPtr(PointerFromQGenericArgument(val1)), C.QtObjectPtr(PointerFromQGenericArgument(val2)), C.QtObjectPtr(PointerFromQGenericArgument(val3)), C.QtObjectPtr(PointerFromQGenericArgument(val4)), C.QtObjectPtr(PointerFromQGenericArgument(val5)), C.QtObjectPtr(PointerFromQGenericArgument(val6)), C.QtObjectPtr(PointerFromQGenericArgument(val7)), C.QtObjectPtr(PointerFromQGenericArgument(val8)), C.QtObjectPtr(PointerFromQGenericArgument(val9))) != 0
	}
	return false
}

func (ptr *QMetaMethod) Invoke(object QObjectITF, connectionType Qt__ConnectionType, returnValue QGenericReturnArgumentITF, val0 QGenericArgumentITF, val1 QGenericArgumentITF, val2 QGenericArgumentITF, val3 QGenericArgumentITF, val4 QGenericArgumentITF, val5 QGenericArgumentITF, val6 QGenericArgumentITF, val7 QGenericArgumentITF, val8 QGenericArgumentITF, val9 QGenericArgumentITF) bool {
	if ptr.Pointer() != nil {
		return C.QMetaMethod_Invoke(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQObject(object)), C.int(connectionType), C.QtObjectPtr(PointerFromQGenericReturnArgument(returnValue)), C.QtObjectPtr(PointerFromQGenericArgument(val0)), C.QtObjectPtr(PointerFromQGenericArgument(val1)), C.QtObjectPtr(PointerFromQGenericArgument(val2)), C.QtObjectPtr(PointerFromQGenericArgument(val3)), C.QtObjectPtr(PointerFromQGenericArgument(val4)), C.QtObjectPtr(PointerFromQGenericArgument(val5)), C.QtObjectPtr(PointerFromQGenericArgument(val6)), C.QtObjectPtr(PointerFromQGenericArgument(val7)), C.QtObjectPtr(PointerFromQGenericArgument(val8)), C.QtObjectPtr(PointerFromQGenericArgument(val9))) != 0
	}
	return false
}

func (ptr *QMetaMethod) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QMetaMethod_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QMetaMethod) MethodIndex() int {
	if ptr.Pointer() != nil {
		return int(C.QMetaMethod_MethodIndex(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMetaMethod) MethodType() QMetaMethod__MethodType {
	if ptr.Pointer() != nil {
		return QMetaMethod__MethodType(C.QMetaMethod_MethodType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMetaMethod) ParameterCount() int {
	if ptr.Pointer() != nil {
		return int(C.QMetaMethod_ParameterCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMetaMethod) ParameterType(index int) int {
	if ptr.Pointer() != nil {
		return int(C.QMetaMethod_ParameterType(C.QtObjectPtr(ptr.Pointer()), C.int(index)))
	}
	return 0
}

func (ptr *QMetaMethod) ReturnType() int {
	if ptr.Pointer() != nil {
		return int(C.QMetaMethod_ReturnType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMetaMethod) Revision() int {
	if ptr.Pointer() != nil {
		return int(C.QMetaMethod_Revision(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}
