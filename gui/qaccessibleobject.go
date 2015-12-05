package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QAccessibleObject struct {
	QAccessibleInterface
}

type QAccessibleObject_ITF interface {
	QAccessibleInterface_ITF
	QAccessibleObject_PTR() *QAccessibleObject
}

func PointerFromQAccessibleObject(ptr QAccessibleObject_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccessibleObject_PTR().Pointer()
	}
	return nil
}

func NewQAccessibleObjectFromPointer(ptr unsafe.Pointer) *QAccessibleObject {
	var n = new(QAccessibleObject)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAccessibleObject) QAccessibleObject_PTR() *QAccessibleObject {
	return ptr
}

func (ptr *QAccessibleObject) ChildAt(x int, y int) *QAccessibleInterface {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleObject::childAt")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQAccessibleInterfaceFromPointer(C.QAccessibleObject_ChildAt(ptr.Pointer(), C.int(x), C.int(y)))
	}
	return nil
}

func (ptr *QAccessibleObject) IsValid() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleObject::isValid")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QAccessibleObject_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QAccessibleObject) Object() *core.QObject {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleObject::object")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QAccessibleObject_Object(ptr.Pointer()))
	}
	return nil
}

func (ptr *QAccessibleObject) SetText(t QAccessible__Text, text string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAccessibleObject::setText")
		}
	}()

	if ptr.Pointer() != nil {
		C.QAccessibleObject_SetText(ptr.Pointer(), C.int(t), C.CString(text))
	}
}
