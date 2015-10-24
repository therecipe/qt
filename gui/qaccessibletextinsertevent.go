package gui

//#include "qaccessibletextinsertevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAccessibleTextInsertEvent struct {
	QAccessibleTextCursorEvent
}

type QAccessibleTextInsertEventITF interface {
	QAccessibleTextCursorEventITF
	QAccessibleTextInsertEventPTR() *QAccessibleTextInsertEvent
}

func PointerFromQAccessibleTextInsertEvent(ptr QAccessibleTextInsertEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccessibleTextInsertEventPTR().Pointer()
	}
	return nil
}

func QAccessibleTextInsertEventFromPointer(ptr unsafe.Pointer) *QAccessibleTextInsertEvent {
	var n = new(QAccessibleTextInsertEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAccessibleTextInsertEvent) QAccessibleTextInsertEventPTR() *QAccessibleTextInsertEvent {
	return ptr
}

func NewQAccessibleTextInsertEvent2(iface QAccessibleInterfaceITF, position int, text string) *QAccessibleTextInsertEvent {
	return QAccessibleTextInsertEventFromPointer(unsafe.Pointer(C.QAccessibleTextInsertEvent_NewQAccessibleTextInsertEvent2(C.QtObjectPtr(PointerFromQAccessibleInterface(iface)), C.int(position), C.CString(text))))
}

func NewQAccessibleTextInsertEvent(object core.QObjectITF, position int, text string) *QAccessibleTextInsertEvent {
	return QAccessibleTextInsertEventFromPointer(unsafe.Pointer(C.QAccessibleTextInsertEvent_NewQAccessibleTextInsertEvent(C.QtObjectPtr(core.PointerFromQObject(object)), C.int(position), C.CString(text))))
}

func (ptr *QAccessibleTextInsertEvent) ChangePosition() int {
	if ptr.Pointer() != nil {
		return int(C.QAccessibleTextInsertEvent_ChangePosition(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAccessibleTextInsertEvent) TextInserted() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAccessibleTextInsertEvent_TextInserted(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}
