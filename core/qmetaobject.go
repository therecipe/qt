package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
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
	defer qt.Recovering("QMetaObject::connectSlotsByName")

	C.QMetaObject_QMetaObject_ConnectSlotsByName(PointerFromQObject(object))
}

func QMetaObject_CheckConnectArgs2(signal QMetaMethod_ITF, method QMetaMethod_ITF) bool {
	defer qt.Recovering("QMetaObject::checkConnectArgs")

	return C.QMetaObject_QMetaObject_CheckConnectArgs2(PointerFromQMetaMethod(signal), PointerFromQMetaMethod(method)) != 0
}

func QMetaObject_CheckConnectArgs(signal string, method string) bool {
	defer qt.Recovering("QMetaObject::checkConnectArgs")

	return C.QMetaObject_QMetaObject_CheckConnectArgs(C.CString(signal), C.CString(method)) != 0
}

func (ptr *QMetaObject) ClassInfoCount() int {
	defer qt.Recovering("QMetaObject::classInfoCount")

	if ptr.Pointer() != nil {
		return int(C.QMetaObject_ClassInfoCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMetaObject) ClassInfoOffset() int {
	defer qt.Recovering("QMetaObject::classInfoOffset")

	if ptr.Pointer() != nil {
		return int(C.QMetaObject_ClassInfoOffset(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMetaObject) ConstructorCount() int {
	defer qt.Recovering("QMetaObject::constructorCount")

	if ptr.Pointer() != nil {
		return int(C.QMetaObject_ConstructorCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMetaObject) EnumeratorCount() int {
	defer qt.Recovering("QMetaObject::enumeratorCount")

	if ptr.Pointer() != nil {
		return int(C.QMetaObject_EnumeratorCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMetaObject) EnumeratorOffset() int {
	defer qt.Recovering("QMetaObject::enumeratorOffset")

	if ptr.Pointer() != nil {
		return int(C.QMetaObject_EnumeratorOffset(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMetaObject) IndexOfClassInfo(name string) int {
	defer qt.Recovering("QMetaObject::indexOfClassInfo")

	if ptr.Pointer() != nil {
		return int(C.QMetaObject_IndexOfClassInfo(ptr.Pointer(), C.CString(name)))
	}
	return 0
}

func (ptr *QMetaObject) IndexOfConstructor(constructor string) int {
	defer qt.Recovering("QMetaObject::indexOfConstructor")

	if ptr.Pointer() != nil {
		return int(C.QMetaObject_IndexOfConstructor(ptr.Pointer(), C.CString(constructor)))
	}
	return 0
}

func (ptr *QMetaObject) IndexOfEnumerator(name string) int {
	defer qt.Recovering("QMetaObject::indexOfEnumerator")

	if ptr.Pointer() != nil {
		return int(C.QMetaObject_IndexOfEnumerator(ptr.Pointer(), C.CString(name)))
	}
	return 0
}

func (ptr *QMetaObject) IndexOfMethod(method string) int {
	defer qt.Recovering("QMetaObject::indexOfMethod")

	if ptr.Pointer() != nil {
		return int(C.QMetaObject_IndexOfMethod(ptr.Pointer(), C.CString(method)))
	}
	return 0
}

func (ptr *QMetaObject) IndexOfProperty(name string) int {
	defer qt.Recovering("QMetaObject::indexOfProperty")

	if ptr.Pointer() != nil {
		return int(C.QMetaObject_IndexOfProperty(ptr.Pointer(), C.CString(name)))
	}
	return 0
}

func (ptr *QMetaObject) IndexOfSignal(signal string) int {
	defer qt.Recovering("QMetaObject::indexOfSignal")

	if ptr.Pointer() != nil {
		return int(C.QMetaObject_IndexOfSignal(ptr.Pointer(), C.CString(signal)))
	}
	return 0
}

func (ptr *QMetaObject) IndexOfSlot(slot string) int {
	defer qt.Recovering("QMetaObject::indexOfSlot")

	if ptr.Pointer() != nil {
		return int(C.QMetaObject_IndexOfSlot(ptr.Pointer(), C.CString(slot)))
	}
	return 0
}

func QMetaObject_InvokeMethod4(obj QObject_ITF, member string, val0 QGenericArgument_ITF, val1 QGenericArgument_ITF, val2 QGenericArgument_ITF, val3 QGenericArgument_ITF, val4 QGenericArgument_ITF, val5 QGenericArgument_ITF, val6 QGenericArgument_ITF, val7 QGenericArgument_ITF, val8 QGenericArgument_ITF, val9 QGenericArgument_ITF) bool {
	defer qt.Recovering("QMetaObject::invokeMethod")

	return C.QMetaObject_QMetaObject_InvokeMethod4(PointerFromQObject(obj), C.CString(member), PointerFromQGenericArgument(val0), PointerFromQGenericArgument(val1), PointerFromQGenericArgument(val2), PointerFromQGenericArgument(val3), PointerFromQGenericArgument(val4), PointerFromQGenericArgument(val5), PointerFromQGenericArgument(val6), PointerFromQGenericArgument(val7), PointerFromQGenericArgument(val8), PointerFromQGenericArgument(val9)) != 0
}

func QMetaObject_InvokeMethod2(obj QObject_ITF, member string, ret QGenericReturnArgument_ITF, val0 QGenericArgument_ITF, val1 QGenericArgument_ITF, val2 QGenericArgument_ITF, val3 QGenericArgument_ITF, val4 QGenericArgument_ITF, val5 QGenericArgument_ITF, val6 QGenericArgument_ITF, val7 QGenericArgument_ITF, val8 QGenericArgument_ITF, val9 QGenericArgument_ITF) bool {
	defer qt.Recovering("QMetaObject::invokeMethod")

	return C.QMetaObject_QMetaObject_InvokeMethod2(PointerFromQObject(obj), C.CString(member), PointerFromQGenericReturnArgument(ret), PointerFromQGenericArgument(val0), PointerFromQGenericArgument(val1), PointerFromQGenericArgument(val2), PointerFromQGenericArgument(val3), PointerFromQGenericArgument(val4), PointerFromQGenericArgument(val5), PointerFromQGenericArgument(val6), PointerFromQGenericArgument(val7), PointerFromQGenericArgument(val8), PointerFromQGenericArgument(val9)) != 0
}

func QMetaObject_InvokeMethod3(obj QObject_ITF, member string, ty Qt__ConnectionType, val0 QGenericArgument_ITF, val1 QGenericArgument_ITF, val2 QGenericArgument_ITF, val3 QGenericArgument_ITF, val4 QGenericArgument_ITF, val5 QGenericArgument_ITF, val6 QGenericArgument_ITF, val7 QGenericArgument_ITF, val8 QGenericArgument_ITF, val9 QGenericArgument_ITF) bool {
	defer qt.Recovering("QMetaObject::invokeMethod")

	return C.QMetaObject_QMetaObject_InvokeMethod3(PointerFromQObject(obj), C.CString(member), C.int(ty), PointerFromQGenericArgument(val0), PointerFromQGenericArgument(val1), PointerFromQGenericArgument(val2), PointerFromQGenericArgument(val3), PointerFromQGenericArgument(val4), PointerFromQGenericArgument(val5), PointerFromQGenericArgument(val6), PointerFromQGenericArgument(val7), PointerFromQGenericArgument(val8), PointerFromQGenericArgument(val9)) != 0
}

func QMetaObject_InvokeMethod(obj QObject_ITF, member string, ty Qt__ConnectionType, ret QGenericReturnArgument_ITF, val0 QGenericArgument_ITF, val1 QGenericArgument_ITF, val2 QGenericArgument_ITF, val3 QGenericArgument_ITF, val4 QGenericArgument_ITF, val5 QGenericArgument_ITF, val6 QGenericArgument_ITF, val7 QGenericArgument_ITF, val8 QGenericArgument_ITF, val9 QGenericArgument_ITF) bool {
	defer qt.Recovering("QMetaObject::invokeMethod")

	return C.QMetaObject_QMetaObject_InvokeMethod(PointerFromQObject(obj), C.CString(member), C.int(ty), PointerFromQGenericReturnArgument(ret), PointerFromQGenericArgument(val0), PointerFromQGenericArgument(val1), PointerFromQGenericArgument(val2), PointerFromQGenericArgument(val3), PointerFromQGenericArgument(val4), PointerFromQGenericArgument(val5), PointerFromQGenericArgument(val6), PointerFromQGenericArgument(val7), PointerFromQGenericArgument(val8), PointerFromQGenericArgument(val9)) != 0
}

func (ptr *QMetaObject) MethodCount() int {
	defer qt.Recovering("QMetaObject::methodCount")

	if ptr.Pointer() != nil {
		return int(C.QMetaObject_MethodCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMetaObject) MethodOffset() int {
	defer qt.Recovering("QMetaObject::methodOffset")

	if ptr.Pointer() != nil {
		return int(C.QMetaObject_MethodOffset(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMetaObject) NewInstance(val0 QGenericArgument_ITF, val1 QGenericArgument_ITF, val2 QGenericArgument_ITF, val3 QGenericArgument_ITF, val4 QGenericArgument_ITF, val5 QGenericArgument_ITF, val6 QGenericArgument_ITF, val7 QGenericArgument_ITF, val8 QGenericArgument_ITF, val9 QGenericArgument_ITF) *QObject {
	defer qt.Recovering("QMetaObject::newInstance")

	if ptr.Pointer() != nil {
		return NewQObjectFromPointer(C.QMetaObject_NewInstance(ptr.Pointer(), PointerFromQGenericArgument(val0), PointerFromQGenericArgument(val1), PointerFromQGenericArgument(val2), PointerFromQGenericArgument(val3), PointerFromQGenericArgument(val4), PointerFromQGenericArgument(val5), PointerFromQGenericArgument(val6), PointerFromQGenericArgument(val7), PointerFromQGenericArgument(val8), PointerFromQGenericArgument(val9)))
	}
	return nil
}

func QMetaObject_NormalizedSignature(method string) *QByteArray {
	defer qt.Recovering("QMetaObject::normalizedSignature")

	return NewQByteArrayFromPointer(C.QMetaObject_QMetaObject_NormalizedSignature(C.CString(method)))
}

func QMetaObject_NormalizedType(ty string) *QByteArray {
	defer qt.Recovering("QMetaObject::normalizedType")

	return NewQByteArrayFromPointer(C.QMetaObject_QMetaObject_NormalizedType(C.CString(ty)))
}

func (ptr *QMetaObject) PropertyCount() int {
	defer qt.Recovering("QMetaObject::propertyCount")

	if ptr.Pointer() != nil {
		return int(C.QMetaObject_PropertyCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMetaObject) PropertyOffset() int {
	defer qt.Recovering("QMetaObject::propertyOffset")

	if ptr.Pointer() != nil {
		return int(C.QMetaObject_PropertyOffset(ptr.Pointer()))
	}
	return 0
}

func (ptr *QMetaObject) SuperClass() *QMetaObject {
	defer qt.Recovering("QMetaObject::superClass")

	if ptr.Pointer() != nil {
		return NewQMetaObjectFromPointer(C.QMetaObject_SuperClass(ptr.Pointer()))
	}
	return nil
}
