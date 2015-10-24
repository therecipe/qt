package core

//#include "qmetaobject.h"
import "C"
import (
	"unsafe"
)

type QMetaObject struct {
	ptr unsafe.Pointer
}

type QMetaObjectITF interface {
	QMetaObjectPTR() *QMetaObject
}

func (p *QMetaObject) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMetaObject) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMetaObject(ptr QMetaObjectITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMetaObjectPTR().Pointer()
	}
	return nil
}

func QMetaObjectFromPointer(ptr unsafe.Pointer) *QMetaObject {
	var n = new(QMetaObject)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMetaObject) QMetaObjectPTR() *QMetaObject {
	return ptr
}

func QMetaObject_ConnectSlotsByName(object QObjectITF) {
	C.QMetaObject_QMetaObject_ConnectSlotsByName(C.QtObjectPtr(PointerFromQObject(object)))
}

func QMetaObject_CheckConnectArgs2(signal QMetaMethodITF, method QMetaMethodITF) bool {
	return C.QMetaObject_QMetaObject_CheckConnectArgs2(C.QtObjectPtr(PointerFromQMetaMethod(signal)), C.QtObjectPtr(PointerFromQMetaMethod(method))) != 0
}

func QMetaObject_CheckConnectArgs(signal string, method string) bool {
	return C.QMetaObject_QMetaObject_CheckConnectArgs(C.CString(signal), C.CString(method)) != 0
}

func (ptr *QMetaObject) ClassInfoCount() int {
	if ptr.Pointer() != nil {
		return int(C.QMetaObject_ClassInfoCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMetaObject) ClassInfoOffset() int {
	if ptr.Pointer() != nil {
		return int(C.QMetaObject_ClassInfoOffset(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMetaObject) ConstructorCount() int {
	if ptr.Pointer() != nil {
		return int(C.QMetaObject_ConstructorCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMetaObject) EnumeratorCount() int {
	if ptr.Pointer() != nil {
		return int(C.QMetaObject_EnumeratorCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMetaObject) EnumeratorOffset() int {
	if ptr.Pointer() != nil {
		return int(C.QMetaObject_EnumeratorOffset(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMetaObject) IndexOfClassInfo(name string) int {
	if ptr.Pointer() != nil {
		return int(C.QMetaObject_IndexOfClassInfo(C.QtObjectPtr(ptr.Pointer()), C.CString(name)))
	}
	return 0
}

func (ptr *QMetaObject) IndexOfConstructor(constructor string) int {
	if ptr.Pointer() != nil {
		return int(C.QMetaObject_IndexOfConstructor(C.QtObjectPtr(ptr.Pointer()), C.CString(constructor)))
	}
	return 0
}

func (ptr *QMetaObject) IndexOfEnumerator(name string) int {
	if ptr.Pointer() != nil {
		return int(C.QMetaObject_IndexOfEnumerator(C.QtObjectPtr(ptr.Pointer()), C.CString(name)))
	}
	return 0
}

func (ptr *QMetaObject) IndexOfMethod(method string) int {
	if ptr.Pointer() != nil {
		return int(C.QMetaObject_IndexOfMethod(C.QtObjectPtr(ptr.Pointer()), C.CString(method)))
	}
	return 0
}

func (ptr *QMetaObject) IndexOfProperty(name string) int {
	if ptr.Pointer() != nil {
		return int(C.QMetaObject_IndexOfProperty(C.QtObjectPtr(ptr.Pointer()), C.CString(name)))
	}
	return 0
}

func (ptr *QMetaObject) IndexOfSignal(signal string) int {
	if ptr.Pointer() != nil {
		return int(C.QMetaObject_IndexOfSignal(C.QtObjectPtr(ptr.Pointer()), C.CString(signal)))
	}
	return 0
}

func (ptr *QMetaObject) IndexOfSlot(slot string) int {
	if ptr.Pointer() != nil {
		return int(C.QMetaObject_IndexOfSlot(C.QtObjectPtr(ptr.Pointer()), C.CString(slot)))
	}
	return 0
}

func QMetaObject_InvokeMethod4(obj QObjectITF, member string, val0 QGenericArgumentITF, val1 QGenericArgumentITF, val2 QGenericArgumentITF, val3 QGenericArgumentITF, val4 QGenericArgumentITF, val5 QGenericArgumentITF, val6 QGenericArgumentITF, val7 QGenericArgumentITF, val8 QGenericArgumentITF, val9 QGenericArgumentITF) bool {
	return C.QMetaObject_QMetaObject_InvokeMethod4(C.QtObjectPtr(PointerFromQObject(obj)), C.CString(member), C.QtObjectPtr(PointerFromQGenericArgument(val0)), C.QtObjectPtr(PointerFromQGenericArgument(val1)), C.QtObjectPtr(PointerFromQGenericArgument(val2)), C.QtObjectPtr(PointerFromQGenericArgument(val3)), C.QtObjectPtr(PointerFromQGenericArgument(val4)), C.QtObjectPtr(PointerFromQGenericArgument(val5)), C.QtObjectPtr(PointerFromQGenericArgument(val6)), C.QtObjectPtr(PointerFromQGenericArgument(val7)), C.QtObjectPtr(PointerFromQGenericArgument(val8)), C.QtObjectPtr(PointerFromQGenericArgument(val9))) != 0
}

func QMetaObject_InvokeMethod2(obj QObjectITF, member string, ret QGenericReturnArgumentITF, val0 QGenericArgumentITF, val1 QGenericArgumentITF, val2 QGenericArgumentITF, val3 QGenericArgumentITF, val4 QGenericArgumentITF, val5 QGenericArgumentITF, val6 QGenericArgumentITF, val7 QGenericArgumentITF, val8 QGenericArgumentITF, val9 QGenericArgumentITF) bool {
	return C.QMetaObject_QMetaObject_InvokeMethod2(C.QtObjectPtr(PointerFromQObject(obj)), C.CString(member), C.QtObjectPtr(PointerFromQGenericReturnArgument(ret)), C.QtObjectPtr(PointerFromQGenericArgument(val0)), C.QtObjectPtr(PointerFromQGenericArgument(val1)), C.QtObjectPtr(PointerFromQGenericArgument(val2)), C.QtObjectPtr(PointerFromQGenericArgument(val3)), C.QtObjectPtr(PointerFromQGenericArgument(val4)), C.QtObjectPtr(PointerFromQGenericArgument(val5)), C.QtObjectPtr(PointerFromQGenericArgument(val6)), C.QtObjectPtr(PointerFromQGenericArgument(val7)), C.QtObjectPtr(PointerFromQGenericArgument(val8)), C.QtObjectPtr(PointerFromQGenericArgument(val9))) != 0
}

func QMetaObject_InvokeMethod3(obj QObjectITF, member string, ty Qt__ConnectionType, val0 QGenericArgumentITF, val1 QGenericArgumentITF, val2 QGenericArgumentITF, val3 QGenericArgumentITF, val4 QGenericArgumentITF, val5 QGenericArgumentITF, val6 QGenericArgumentITF, val7 QGenericArgumentITF, val8 QGenericArgumentITF, val9 QGenericArgumentITF) bool {
	return C.QMetaObject_QMetaObject_InvokeMethod3(C.QtObjectPtr(PointerFromQObject(obj)), C.CString(member), C.int(ty), C.QtObjectPtr(PointerFromQGenericArgument(val0)), C.QtObjectPtr(PointerFromQGenericArgument(val1)), C.QtObjectPtr(PointerFromQGenericArgument(val2)), C.QtObjectPtr(PointerFromQGenericArgument(val3)), C.QtObjectPtr(PointerFromQGenericArgument(val4)), C.QtObjectPtr(PointerFromQGenericArgument(val5)), C.QtObjectPtr(PointerFromQGenericArgument(val6)), C.QtObjectPtr(PointerFromQGenericArgument(val7)), C.QtObjectPtr(PointerFromQGenericArgument(val8)), C.QtObjectPtr(PointerFromQGenericArgument(val9))) != 0
}

func QMetaObject_InvokeMethod(obj QObjectITF, member string, ty Qt__ConnectionType, ret QGenericReturnArgumentITF, val0 QGenericArgumentITF, val1 QGenericArgumentITF, val2 QGenericArgumentITF, val3 QGenericArgumentITF, val4 QGenericArgumentITF, val5 QGenericArgumentITF, val6 QGenericArgumentITF, val7 QGenericArgumentITF, val8 QGenericArgumentITF, val9 QGenericArgumentITF) bool {
	return C.QMetaObject_QMetaObject_InvokeMethod(C.QtObjectPtr(PointerFromQObject(obj)), C.CString(member), C.int(ty), C.QtObjectPtr(PointerFromQGenericReturnArgument(ret)), C.QtObjectPtr(PointerFromQGenericArgument(val0)), C.QtObjectPtr(PointerFromQGenericArgument(val1)), C.QtObjectPtr(PointerFromQGenericArgument(val2)), C.QtObjectPtr(PointerFromQGenericArgument(val3)), C.QtObjectPtr(PointerFromQGenericArgument(val4)), C.QtObjectPtr(PointerFromQGenericArgument(val5)), C.QtObjectPtr(PointerFromQGenericArgument(val6)), C.QtObjectPtr(PointerFromQGenericArgument(val7)), C.QtObjectPtr(PointerFromQGenericArgument(val8)), C.QtObjectPtr(PointerFromQGenericArgument(val9))) != 0
}

func (ptr *QMetaObject) MethodCount() int {
	if ptr.Pointer() != nil {
		return int(C.QMetaObject_MethodCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMetaObject) MethodOffset() int {
	if ptr.Pointer() != nil {
		return int(C.QMetaObject_MethodOffset(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMetaObject) NewInstance(val0 QGenericArgumentITF, val1 QGenericArgumentITF, val2 QGenericArgumentITF, val3 QGenericArgumentITF, val4 QGenericArgumentITF, val5 QGenericArgumentITF, val6 QGenericArgumentITF, val7 QGenericArgumentITF, val8 QGenericArgumentITF, val9 QGenericArgumentITF) *QObject {
	if ptr.Pointer() != nil {
		return QObjectFromPointer(unsafe.Pointer(C.QMetaObject_NewInstance(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQGenericArgument(val0)), C.QtObjectPtr(PointerFromQGenericArgument(val1)), C.QtObjectPtr(PointerFromQGenericArgument(val2)), C.QtObjectPtr(PointerFromQGenericArgument(val3)), C.QtObjectPtr(PointerFromQGenericArgument(val4)), C.QtObjectPtr(PointerFromQGenericArgument(val5)), C.QtObjectPtr(PointerFromQGenericArgument(val6)), C.QtObjectPtr(PointerFromQGenericArgument(val7)), C.QtObjectPtr(PointerFromQGenericArgument(val8)), C.QtObjectPtr(PointerFromQGenericArgument(val9)))))
	}
	return nil
}

func (ptr *QMetaObject) PropertyCount() int {
	if ptr.Pointer() != nil {
		return int(C.QMetaObject_PropertyCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMetaObject) PropertyOffset() int {
	if ptr.Pointer() != nil {
		return int(C.QMetaObject_PropertyOffset(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMetaObject) SuperClass() *QMetaObject {
	if ptr.Pointer() != nil {
		return QMetaObjectFromPointer(unsafe.Pointer(C.QMetaObject_SuperClass(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}
