package gui

//#include "qaccessibletextupdateevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAccessibleTextUpdateEvent struct {
	QAccessibleTextCursorEvent
}

type QAccessibleTextUpdateEvent_ITF interface {
	QAccessibleTextCursorEvent_ITF
	QAccessibleTextUpdateEvent_PTR() *QAccessibleTextUpdateEvent
}

func PointerFromQAccessibleTextUpdateEvent(ptr QAccessibleTextUpdateEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccessibleTextUpdateEvent_PTR().Pointer()
	}
	return nil
}

func NewQAccessibleTextUpdateEventFromPointer(ptr unsafe.Pointer) *QAccessibleTextUpdateEvent {
	var n = new(QAccessibleTextUpdateEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAccessibleTextUpdateEvent) QAccessibleTextUpdateEvent_PTR() *QAccessibleTextUpdateEvent {
	return ptr
}

func NewQAccessibleTextUpdateEvent2(iface QAccessibleInterface_ITF, position int, oldText string, text string) *QAccessibleTextUpdateEvent {
	return NewQAccessibleTextUpdateEventFromPointer(C.QAccessibleTextUpdateEvent_NewQAccessibleTextUpdateEvent2(PointerFromQAccessibleInterface(iface), C.int(position), C.CString(oldText), C.CString(text)))
}

func NewQAccessibleTextUpdateEvent(object core.QObject_ITF, position int, oldText string, text string) *QAccessibleTextUpdateEvent {
	return NewQAccessibleTextUpdateEventFromPointer(C.QAccessibleTextUpdateEvent_NewQAccessibleTextUpdateEvent(core.PointerFromQObject(object), C.int(position), C.CString(oldText), C.CString(text)))
}

func (ptr *QAccessibleTextUpdateEvent) ChangePosition() int {
	if ptr.Pointer() != nil {
		return int(C.QAccessibleTextUpdateEvent_ChangePosition(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccessibleTextUpdateEvent) TextInserted() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAccessibleTextUpdateEvent_TextInserted(ptr.Pointer()))
	}
	return ""
}

func (ptr *QAccessibleTextUpdateEvent) TextRemoved() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAccessibleTextUpdateEvent_TextRemoved(ptr.Pointer()))
	}
	return ""
}
