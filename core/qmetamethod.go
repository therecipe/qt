package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QMetaMethod struct {
	ptr unsafe.Pointer
}

type QMetaMethod_ITF interface {
	QMetaMethod_PTR() *QMetaMethod
}

func (p *QMetaMethod) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMetaMethod) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMetaMethod(ptr QMetaMethod_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMetaMethod_PTR().Pointer()
	}
	return nil
}

func NewQMetaMethodFromPointer(ptr unsafe.Pointer) *QMetaMethod {
	var n = new(QMetaMethod)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMetaMethod) QMetaMethod_PTR() *QMetaMethod {
	return ptr
}

//QMetaMethod::Access
type QMetaMethod__Access int64

const (
	QMetaMethod__Private   = QMetaMethod__Access(0)
	QMetaMethod__Protected = QMetaMethod__Access(1)
	QMetaMethod__Public    = QMetaMethod__Access(2)
)

//QMetaMethod::MethodType
type QMetaMethod__MethodType int64

const (
	QMetaMethod__Method      = QMetaMethod__MethodType(0)
	QMetaMethod__Signal      = QMetaMethod__MethodType(1)
	QMetaMethod__Slot        = QMetaMethod__MethodType(2)
	QMetaMethod__Constructor = QMetaMethod__MethodType(3)
)

func (ptr *QMetaMethod) Access() QMetaMethod__Access {
	defer qt.Recovering("QMetaMethod::access")

	if ptr.Pointer() != nil {
		return QMetaMethod__Access(C.QMetaMethod_Access(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMetaMethod) Invoke4(object QObject_ITF, val0 QGenericArgument_ITF, val1 QGenericArgument_ITF, val2 QGenericArgument_ITF, val3 QGenericArgument_ITF, val4 QGenericArgument_ITF, val5 QGenericArgument_ITF, val6 QGenericArgument_ITF, val7 QGenericArgument_ITF, val8 QGenericArgument_ITF, val9 QGenericArgument_ITF) bool {
	defer qt.Recovering("QMetaMethod::invoke")

	if ptr.Pointer() != nil {
		return C.QMetaMethod_Invoke4(ptr.Pointer(), PointerFromQObject(object), PointerFromQGenericArgument(val0), PointerFromQGenericArgument(val1), PointerFromQGenericArgument(val2), PointerFromQGenericArgument(val3), PointerFromQGenericArgument(val4), PointerFromQGenericArgument(val5), PointerFromQGenericArgument(val6), PointerFromQGenericArgument(val7), PointerFromQGenericArgument(val8), PointerFromQGenericArgument(val9)) != 0
	}
	return false
}

func (ptr *QMetaMethod) Invoke2(object QObject_ITF, returnValue QGenericReturnArgument_ITF, val0 QGenericArgument_ITF, val1 QGenericArgument_ITF, val2 QGenericArgument_ITF, val3 QGenericArgument_ITF, val4 QGenericArgument_ITF, val5 QGenericArgument_ITF, val6 QGenericArgument_ITF, val7 QGenericArgument_ITF, val8 QGenericArgument_ITF, val9 QGenericArgument_ITF) bool {
	defer qt.Recovering("QMetaMethod::invoke")

	if ptr.Pointer() != nil {
		return C.QMetaMethod_Invoke2(ptr.Pointer(), PointerFromQObject(object), PointerFromQGenericReturnArgument(returnValue), PointerFromQGenericArgument(val0), PointerFromQGenericArgument(val1), PointerFromQGenericArgument(val2), PointerFromQGenericArgument(val3), PointerFromQGenericArgument(val4), PointerFromQGenericArgument(val5), PointerFromQGenericArgument(val6), PointerFromQGenericArgument(val7), PointerFromQGenericArgument(val8), PointerFromQGenericArgument(val9)) != 0
	}
	return false
}

func (ptr *QMetaMethod) Invoke3(object QObject_ITF, connectionType Qt__ConnectionType, val0 QGenericArgument_ITF, val1 QGenericArgument_ITF, val2 QGenericArgument_ITF, val3 QGenericArgument_ITF, val4 QGenericArgument_ITF, val5 QGenericArgument_ITF, val6 QGenericArgument_ITF, val7 QGenericArgument_ITF, val8 QGenericArgument_ITF, val9 QGenericArgument_ITF) bool {
	defer qt.Recovering("QMetaMethod::invoke")

	if ptr.Pointer() != nil {
		return C.QMetaMethod_Invoke3(ptr.Pointer(), PointerFromQObject(object), C.int(connectionType), PointerFromQGenericArgument(val0), PointerFromQGenericArgument(val1), PointerFromQGenericArgument(val2), PointerFromQGenericArgument(val3), PointerFromQGenericArgument(val4), PointerFromQGenericArgument(val5), PointerFromQGenericArgument(val6), PointerFromQGenericArgument(val7), PointerFromQGenericArgument(val8), PointerFromQGenericArgument(val9)) != 0
	}
	return false
}

func (ptr *QMetaMethod) Invoke(object QObject_ITF, connectionType Qt__ConnectionType, returnValue QGenericReturnArgument_ITF, val0 QGenericArgument_ITF, val1 QGenericArgument_ITF, val2 QGenericArgument_ITF, val3 QGenericArgument_ITF, val4 QGenericArgument_ITF, val5 QGenericArgument_ITF, val6 QGenericArgument_ITF, val7 QGenericArgument_ITF, val8 QGenericArgument_ITF, val9 QGenericArgument_ITF) bool {
	defer qt.Recovering("QMetaMethod::invoke")

	if ptr.Pointer() != nil {
		return C.QMetaMethod_Invoke(ptr.Pointer(), PointerFromQObject(object), C.int(connectionType), PointerFromQGenericReturnArgument(returnValue), PointerFromQGenericArgument(val0), PointerFromQGenericArgument(val1), PointerFromQGenericArgument(val2), PointerFromQGenericArgument(val3), PointerFromQGenericArgument(val4), PointerFromQGenericArgument(val5), PointerFromQGenericArgument(val6), PointerFromQGenericArgument(val7), PointerFromQGenericArgument(val8), PointerFromQGenericArgument(val9)) != 0
	}
	return false
}

func (ptr *QMetaMethod) IsValid() bool {
	defer qt.Recovering("QMetaMethod::isValid")

	if ptr.Pointer() != nil {
		return C.QMetaMethod_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QMetaMethod) MethodIndex() int {
	defer qt.Recovering("QMetaMethod::methodIndex")

	if ptr.Pointer() != nil {
		return int(C.QMetaMethod_MethodIndex(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMetaMethod) MethodSignature() *QByteArray {
	defer qt.Recovering("QMetaMethod::methodSignature")

	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QMetaMethod_MethodSignature(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMetaMethod) MethodType() QMetaMethod__MethodType {
	defer qt.Recovering("QMetaMethod::methodType")

	if ptr.Pointer() != nil {
		return QMetaMethod__MethodType(C.QMetaMethod_MethodType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMetaMethod) Name() *QByteArray {
	defer qt.Recovering("QMetaMethod::name")

	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QMetaMethod_Name(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMetaMethod) ParameterCount() int {
	defer qt.Recovering("QMetaMethod::parameterCount")

	if ptr.Pointer() != nil {
		return int(C.QMetaMethod_ParameterCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMetaMethod) ParameterType(index int) int {
	defer qt.Recovering("QMetaMethod::parameterType")

	if ptr.Pointer() != nil {
		return int(C.QMetaMethod_ParameterType(ptr.Pointer(), C.int(index)))
	}
	return 0
}

func (ptr *QMetaMethod) ReturnType() int {
	defer qt.Recovering("QMetaMethod::returnType")

	if ptr.Pointer() != nil {
		return int(C.QMetaMethod_ReturnType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMetaMethod) Revision() int {
	defer qt.Recovering("QMetaMethod::revision")

	if ptr.Pointer() != nil {
		return int(C.QMetaMethod_Revision(ptr.Pointer()))
	}
	return 0
}
