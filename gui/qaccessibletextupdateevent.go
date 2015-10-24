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

type QAccessibleTextUpdateEventITF interface {
	QAccessibleTextCursorEventITF
	QAccessibleTextUpdateEventPTR() *QAccessibleTextUpdateEvent
}

func PointerFromQAccessibleTextUpdateEvent(ptr QAccessibleTextUpdateEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccessibleTextUpdateEventPTR().Pointer()
	}
	return nil
}

func QAccessibleTextUpdateEventFromPointer(ptr unsafe.Pointer) *QAccessibleTextUpdateEvent {
	var n = new(QAccessibleTextUpdateEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAccessibleTextUpdateEvent) QAccessibleTextUpdateEventPTR() *QAccessibleTextUpdateEvent {
	return ptr
}

func NewQAccessibleTextUpdateEvent2(iface QAccessibleInterfaceITF, position int, oldText string, text string) *QAccessibleTextUpdateEvent {
	return QAccessibleTextUpdateEventFromPointer(unsafe.Pointer(C.QAccessibleTextUpdateEvent_NewQAccessibleTextUpdateEvent2(C.QtObjectPtr(PointerFromQAccessibleInterface(iface)), C.int(position), C.CString(oldText), C.CString(text))))
}

func NewQAccessibleTextUpdateEvent(object core.QObjectITF, position int, oldText string, text string) *QAccessibleTextUpdateEvent {
	return QAccessibleTextUpdateEventFromPointer(unsafe.Pointer(C.QAccessibleTextUpdateEvent_NewQAccessibleTextUpdateEvent(C.QtObjectPtr(core.PointerFromQObject(object)), C.int(position), C.CString(oldText), C.CString(text))))
}

func (ptr *QAccessibleTextUpdateEvent) ChangePosition() int {
	if ptr.Pointer() != nil {
		return int(C.QAccessibleTextUpdateEvent_ChangePosition(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAccessibleTextUpdateEvent) TextInserted() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAccessibleTextUpdateEvent_TextInserted(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QAccessibleTextUpdateEvent) TextRemoved() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAccessibleTextUpdateEvent_TextRemoved(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}
