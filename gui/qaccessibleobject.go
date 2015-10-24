package gui

//#include "qaccessibleobject.h"
import "C"
import (
	"unsafe"
)

type QAccessibleObject struct {
	QAccessibleInterface
}

type QAccessibleObjectITF interface {
	QAccessibleInterfaceITF
	QAccessibleObjectPTR() *QAccessibleObject
}

func PointerFromQAccessibleObject(ptr QAccessibleObjectITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccessibleObjectPTR().Pointer()
	}
	return nil
}

func QAccessibleObjectFromPointer(ptr unsafe.Pointer) *QAccessibleObject {
	var n = new(QAccessibleObject)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAccessibleObject) QAccessibleObjectPTR() *QAccessibleObject {
	return ptr
}
