package qml

//#include "qml.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QQmlListReference struct {
	ptr unsafe.Pointer
}

type QQmlListReference_ITF interface {
	QQmlListReference_PTR() *QQmlListReference
}

func (p *QQmlListReference) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QQmlListReference) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQQmlListReference(ptr QQmlListReference_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QQmlListReference_PTR().Pointer()
	}
	return nil
}

func NewQQmlListReferenceFromPointer(ptr unsafe.Pointer) *QQmlListReference {
	var n = new(QQmlListReference)
	n.SetPointer(ptr)
	return n
}

func (ptr *QQmlListReference) QQmlListReference_PTR() *QQmlListReference {
	return ptr
}

func NewQQmlListReference() *QQmlListReference {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlListReference::QQmlListReference")
		}
	}()

	return NewQQmlListReferenceFromPointer(C.QQmlListReference_NewQQmlListReference())
}

func NewQQmlListReference2(object core.QObject_ITF, property string, engine QQmlEngine_ITF) *QQmlListReference {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlListReference::QQmlListReference")
		}
	}()

	return NewQQmlListReferenceFromPointer(C.QQmlListReference_NewQQmlListReference2(core.PointerFromQObject(object), C.CString(property), PointerFromQQmlEngine(engine)))
}

func (ptr *QQmlListReference) Append(object core.QObject_ITF) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlListReference::append")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QQmlListReference_Append(ptr.Pointer(), core.PointerFromQObject(object)) != 0
	}
	return false
}

func (ptr *QQmlListReference) At(index int) *core.QObject {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlListReference::at")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QQmlListReference_At(ptr.Pointer(), C.int(index)))
	}
	return nil
}

func (ptr *QQmlListReference) CanAppend() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlListReference::canAppend")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QQmlListReference_CanAppend(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlListReference) CanAt() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlListReference::canAt")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QQmlListReference_CanAt(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlListReference) CanClear() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlListReference::canClear")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QQmlListReference_CanClear(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlListReference) CanCount() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlListReference::canCount")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QQmlListReference_CanCount(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlListReference) Clear() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlListReference::clear")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QQmlListReference_Clear(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlListReference) Count() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlListReference::count")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QQmlListReference_Count(ptr.Pointer()))
	}
	return 0
}

func (ptr *QQmlListReference) IsManipulable() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlListReference::isManipulable")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QQmlListReference_IsManipulable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlListReference) IsReadable() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlListReference::isReadable")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QQmlListReference_IsReadable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlListReference) IsValid() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlListReference::isValid")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QQmlListReference_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QQmlListReference) ListElementType() *core.QMetaObject {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlListReference::listElementType")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QQmlListReference_ListElementType(ptr.Pointer()))
	}
	return nil
}

func (ptr *QQmlListReference) Object() *core.QObject {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QQmlListReference::object")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QQmlListReference_Object(ptr.Pointer()))
	}
	return nil
}
