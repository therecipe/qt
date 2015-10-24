package gui

//#include "qaccessiblevaluechangeevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAccessibleValueChangeEvent struct {
	QAccessibleEvent
}

type QAccessibleValueChangeEventITF interface {
	QAccessibleEventITF
	QAccessibleValueChangeEventPTR() *QAccessibleValueChangeEvent
}

func PointerFromQAccessibleValueChangeEvent(ptr QAccessibleValueChangeEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccessibleValueChangeEventPTR().Pointer()
	}
	return nil
}

func QAccessibleValueChangeEventFromPointer(ptr unsafe.Pointer) *QAccessibleValueChangeEvent {
	var n = new(QAccessibleValueChangeEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAccessibleValueChangeEvent) QAccessibleValueChangeEventPTR() *QAccessibleValueChangeEvent {
	return ptr
}

func NewQAccessibleValueChangeEvent2(iface QAccessibleInterfaceITF, val string) *QAccessibleValueChangeEvent {
	return QAccessibleValueChangeEventFromPointer(unsafe.Pointer(C.QAccessibleValueChangeEvent_NewQAccessibleValueChangeEvent2(C.QtObjectPtr(PointerFromQAccessibleInterface(iface)), C.CString(val))))
}

func NewQAccessibleValueChangeEvent(object core.QObjectITF, value string) *QAccessibleValueChangeEvent {
	return QAccessibleValueChangeEventFromPointer(unsafe.Pointer(C.QAccessibleValueChangeEvent_NewQAccessibleValueChangeEvent(C.QtObjectPtr(core.PointerFromQObject(object)), C.CString(value))))
}

func (ptr *QAccessibleValueChangeEvent) SetValue(value string) {
	if ptr.Pointer() != nil {
		C.QAccessibleValueChangeEvent_SetValue(C.QtObjectPtr(ptr.Pointer()), C.CString(value))
	}
}

func (ptr *QAccessibleValueChangeEvent) Value() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAccessibleValueChangeEvent_Value(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}
