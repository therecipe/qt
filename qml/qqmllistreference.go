package qml

//#include "qqmllistreference.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QQmlListReference struct {
	ptr unsafe.Pointer
}

type QQmlListReferenceITF interface {
	QQmlListReferencePTR() *QQmlListReference
}

func (p *QQmlListReference) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QQmlListReference) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQQmlListReference(ptr QQmlListReferenceITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlListReferencePTR().Pointer()
	}
	return nil
}

func QQmlListReferenceFromPointer(ptr unsafe.Pointer) *QQmlListReference {
	var n = new(QQmlListReference)
	n.SetPointer(ptr)
	return n
}

func (ptr *QQmlListReference) QQmlListReferencePTR() *QQmlListReference {
	return ptr
}

func NewQQmlListReference() *QQmlListReference {
	return QQmlListReferenceFromPointer(unsafe.Pointer(C.QQmlListReference_NewQQmlListReference()))
}

func NewQQmlListReference2(object core.QObjectITF, property string, engine QQmlEngineITF) *QQmlListReference {
	return QQmlListReferenceFromPointer(unsafe.Pointer(C.QQmlListReference_NewQQmlListReference2(C.QtObjectPtr(core.PointerFromQObject(object)), C.CString(property), C.QtObjectPtr(PointerFromQQmlEngine(engine)))))
}

func (ptr *QQmlListReference) Append(object core.QObjectITF) bool {
	if ptr.Pointer() != nil {
		return C.QQmlListReference_Append(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQObject(object))) != 0
	}
	return false
}

func (ptr *QQmlListReference) At(index int) *core.QObject {
	if ptr.Pointer() != nil {
		return core.QObjectFromPointer(unsafe.Pointer(C.QQmlListReference_At(C.QtObjectPtr(ptr.Pointer()), C.int(index))))
	}
	return nil
}

func (ptr *QQmlListReference) CanAppend() bool {
	if ptr.Pointer() != nil {
		return C.QQmlListReference_CanAppend(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlListReference) CanAt() bool {
	if ptr.Pointer() != nil {
		return C.QQmlListReference_CanAt(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlListReference) CanClear() bool {
	if ptr.Pointer() != nil {
		return C.QQmlListReference_CanClear(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlListReference) CanCount() bool {
	if ptr.Pointer() != nil {
		return C.QQmlListReference_CanCount(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlListReference) Clear() bool {
	if ptr.Pointer() != nil {
		return C.QQmlListReference_Clear(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlListReference) Count() int {
	if ptr.Pointer() != nil {
		return int(C.QQmlListReference_Count(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QQmlListReference) IsManipulable() bool {
	if ptr.Pointer() != nil {
		return C.QQmlListReference_IsManipulable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlListReference) IsReadable() bool {
	if ptr.Pointer() != nil {
		return C.QQmlListReference_IsReadable(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlListReference) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QQmlListReference_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QQmlListReference) ListElementType() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.QMetaObjectFromPointer(unsafe.Pointer(C.QQmlListReference_ListElementType(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QQmlListReference) Object() *core.QObject {
	if ptr.Pointer() != nil {
		return core.QObjectFromPointer(unsafe.Pointer(C.QQmlListReference_Object(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}
