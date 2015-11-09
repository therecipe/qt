package core

//#include "qmetaobject.h"
import "C"
import (
	"unsafe"
)

type QMetaObject struct {
	ptr unsafe.Pointer
}

type QMetaObject_ITF interface {
	QMetaObject_PTR() *QMetaObject
}

func (p *QMetaObject) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMetaObject) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMetaObject(ptr QMetaObject_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMetaObject_PTR().Pointer()
	}
	return nil
}

func NewQMetaObjectFromPointer(ptr unsafe.Pointer) *QMetaObject {
	var n = new(QMetaObject)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMetaObject) QMetaObject_PTR() *QMetaObject {
	return ptr
}

func QMetaObject_ConnectSlotsByName(object QObject_ITF) {
	C.QMetaObject_QMetaObject_ConnectSlotsByName(PointerFromQObject(object))
}

func QMetaObject_CheckConnectArgs2(signal QMetaMethod_ITF, method QMetaMethod_ITF) bool {
	return C.QMetaObject_QMetaObject_CheckConnectArgs2(PointerFromQMetaMethod(signal), PointerFromQMetaMethod(method)) != 0
}

func QMetaObject_CheckConnectArgs(signal string, method string) bool {
	return C.QMetaObject_QMetaObject_CheckConnectArgs(C.CString(signal), C.CString(method)) != 0
}

func (ptr *QMetaObject) ClassInfoCount() int {
	if ptr.Pointer() != nil {
		return int(C.QMetaObject_ClassInfoCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMetaObject) ClassInfoOffset() int {
	if ptr.Pointer() != nil {
		return int(C.QMetaObject_ClassInfoOffset(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMetaObject) ConstructorCount() int {
	if ptr.Pointer() != nil {
		return int(C.QMetaObject_ConstructorCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMetaObject) EnumeratorCount() int {
	if ptr.Pointer() != nil {
		return int(C.QMetaObject_EnumeratorCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMetaObject) EnumeratorOffset() int {
	if ptr.Pointer() != nil {
		return int(C.QMetaObject_EnumeratorOffset(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMetaObject) IndexOfClassInfo(name string) int {
	if ptr.Pointer() != nil {
		return int(C.QMetaObject_IndexOfClassInfo(ptr.Pointer(), C.CString(name)))
	}
	return 0
}

func (ptr *QMetaObject) IndexOfConstructor(constructor string) int {
	if ptr.Pointer() != nil {
		return int(C.QMetaObject_IndexOfConstructor(ptr.Pointer(), C.CString(constructor)))
	}
	return 0
}

func (ptr *QMetaObject) IndexOfEnumerator(name string) int {
	if ptr.Pointer() != nil {
		return int(C.QMetaObject_IndexOfEnumerator(ptr.Pointer(), C.CString(name)))
	}
	return 0
}

func (ptr *QMetaObject) IndexOfMethod(method string) int {
	if ptr.Pointer() != nil {
		return int(C.QMetaObject_IndexOfMethod(ptr.Pointer(), C.CString(method)))
	}
	return 0
}

func (ptr *QMetaObject) IndexOfProperty(name string) int {
	if ptr.Pointer() != nil {
		return int(C.QMetaObject_IndexOfProperty(ptr.Pointer(), C.CString(name)))
	}
	return 0
}

func (ptr *QMetaObject) IndexOfSignal(signal string) int {
	if ptr.Pointer() != nil {
		return int(C.QMetaObject_IndexOfSignal(ptr.Pointer(), C.CString(signal)))
	}
	return 0
}

func (ptr *QMetaObject) IndexOfSlot(slot string) int {
	if ptr.Pointer() != nil {
		return int(C.QMetaObject_IndexOfSlot(ptr.Pointer(), C.CString(slot)))
	}
	return 0
}

func QMetaObject_InvokeMethod4(obj QObject_ITF, member string, val0 QGenericArgument_ITF, val1 QGenericArgument_ITF, val2 QGenericArgument_ITF, val3 QGenericArgument_ITF, val4 QGenericArgument_ITF, val5 QGenericArgument_ITF, val6 QGenericArgument_ITF, val7 QGenericArgument_ITF, val8 QGenericArgument_ITF, val9 QGenericArgument_ITF) bool {
	return C.QMetaObject_QMetaObject_InvokeMethod4(PointerFromQObject(obj), C.CString(member), PointerFromQGenericArgument(val0), PointerFromQGenericArgument(val1), PointerFromQGenericArgument(val2), PointerFromQGenericArgument(val3), PointerFromQGenericArgument(val4), PointerFromQGenericArgument(val5), PointerFromQGenericArgument(val6), PointerFromQGenericArgument(val7), PointerFromQGenericArgument(val8), PointerFromQGenericArgument(val9)) != 0
}

func QMetaObject_InvokeMethod2(obj QObject_ITF, member string, ret QGenericReturnArgument_ITF, val0 QGenericArgument_ITF, val1 QGenericArgument_ITF, val2 QGenericArgument_ITF, val3 QGenericArgument_ITF, val4 QGenericArgument_ITF, val5 QGenericArgument_ITF, val6 QGenericArgument_ITF, val7 QGenericArgument_ITF, val8 QGenericArgument_ITF, val9 QGenericArgument_ITF) bool {
	return C.QMetaObject_QMetaObject_InvokeMethod2(PointerFromQObject(obj), C.CString(member), PointerFromQGenericReturnArgument(ret), PointerFromQGenericArgument(val0), PointerFromQGenericArgument(val1), PointerFromQGenericArgument(val2), PointerFromQGenericArgument(val3), PointerFromQGenericArgument(val4), PointerFromQGenericArgument(val5), PointerFromQGenericArgument(val6), PointerFromQGenericArgument(val7), PointerFromQGenericArgument(val8), PointerFromQGenericArgument(val9)) != 0
}

func QMetaObject_InvokeMethod3(obj QObject_ITF, member string, ty Qt__ConnectionType, val0 QGenericArgument_ITF, val1 QGenericArgument_ITF, val2 QGenericArgument_ITF, val3 QGenericArgument_ITF, val4 QGenericArgument_ITF, val5 QGenericArgument_ITF, val6 QGenericArgument_ITF, val7 QGenericArgument_ITF, val8 QGenericArgument_ITF, val9 QGenericArgument_ITF) bool {
	return C.QMetaObject_QMetaObject_InvokeMethod3(PointerFromQObject(obj), C.CString(member), C.int(ty), PointerFromQGenericArgument(val0), PointerFromQGenericArgument(val1), PointerFromQGenericArgument(val2), PointerFromQGenericArgument(val3), PointerFromQGenericArgument(val4), PointerFromQGenericArgument(val5), PointerFromQGenericArgument(val6), PointerFromQGenericArgument(val7), PointerFromQGenericArgument(val8), PointerFromQGenericArgument(val9)) != 0
}

func QMetaObject_InvokeMethod(obj QObject_ITF, member string, ty Qt__ConnectionType, ret QGenericReturnArgument_ITF, val0 QGenericArgument_ITF, val1 QGenericArgument_ITF, val2 QGenericArgument_ITF, val3 QGenericArgument_ITF, val4 QGenericArgument_ITF, val5 QGenericArgument_ITF, val6 QGenericArgument_ITF, val7 QGenericArgument_ITF, val8 QGenericArgument_ITF, val9 QGenericArgument_ITF) bool {
	return C.QMetaObject_QMetaObject_InvokeMethod(PointerFromQObject(obj), C.CString(member), C.int(ty), PointerFromQGenericReturnArgument(ret), PointerFromQGenericArgument(val0), PointerFromQGenericArgument(val1), PointerFromQGenericArgument(val2), PointerFromQGenericArgument(val3), PointerFromQGenericArgument(val4), PointerFromQGenericArgument(val5), PointerFromQGenericArgument(val6), PointerFromQGenericArgument(val7), PointerFromQGenericArgument(val8), PointerFromQGenericArgument(val9)) != 0
}

func (ptr *QMetaObject) MethodCount() int {
	if ptr.Pointer() != nil {
		return int(C.QMetaObject_MethodCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMetaObject) MethodOffset() int {
	if ptr.Pointer() != nil {
		return int(C.QMetaObject_MethodOffset(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMetaObject) NewInstance(val0 QGenericArgument_ITF, val1 QGenericArgument_ITF, val2 QGenericArgument_ITF, val3 QGenericArgument_ITF, val4 QGenericArgument_ITF, val5 QGenericArgument_ITF, val6 QGenericArgument_ITF, val7 QGenericArgument_ITF, val8 QGenericArgument_ITF, val9 QGenericArgument_ITF) *QObject {
	if ptr.Pointer() != nil {
		return NewQObjectFromPointer(C.QMetaObject_NewInstance(ptr.Pointer(), PointerFromQGenericArgument(val0), PointerFromQGenericArgument(val1), PointerFromQGenericArgument(val2), PointerFromQGenericArgument(val3), PointerFromQGenericArgument(val4), PointerFromQGenericArgument(val5), PointerFromQGenericArgument(val6), PointerFromQGenericArgument(val7), PointerFromQGenericArgument(val8), PointerFromQGenericArgument(val9)))
	}
	return nil
}

func QMetaObject_NormalizedSignature(method string) *QByteArray {
	return NewQByteArrayFromPointer(C.QMetaObject_QMetaObject_NormalizedSignature(C.CString(method)))
}

func QMetaObject_NormalizedType(ty string) *QByteArray {
	return NewQByteArrayFromPointer(C.QMetaObject_QMetaObject_NormalizedType(C.CString(ty)))
}

func (ptr *QMetaObject) PropertyCount() int {
	if ptr.Pointer() != nil {
		return int(C.QMetaObject_PropertyCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMetaObject) PropertyOffset() int {
	if ptr.Pointer() != nil {
		return int(C.QMetaObject_PropertyOffset(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMetaObject) SuperClass() *QMetaObject {
	if ptr.Pointer() != nil {
		return NewQMetaObjectFromPointer(C.QMetaObject_SuperClass(ptr.Pointer()))
	}
	return nil
}
