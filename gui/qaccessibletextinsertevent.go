package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAccessibleTextInsertEvent struct {
	QAccessibleTextCursorEvent
}

type QAccessibleTextInsertEvent_ITF interface {
	QAccessibleTextCursorEvent_ITF
	QAccessibleTextInsertEvent_PTR() *QAccessibleTextInsertEvent
}

func PointerFromQAccessibleTextInsertEvent(ptr QAccessibleTextInsertEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccessibleTextInsertEvent_PTR().Pointer()
	}
	return nil
}

func NewQAccessibleTextInsertEventFromPointer(ptr unsafe.Pointer) *QAccessibleTextInsertEvent {
	var n = new(QAccessibleTextInsertEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAccessibleTextInsertEvent) QAccessibleTextInsertEvent_PTR() *QAccessibleTextInsertEvent {
	return ptr
}

func NewQAccessibleTextInsertEvent2(iface QAccessibleInterface_ITF, position int, text string) *QAccessibleTextInsertEvent {
	defer qt.Recovering("QAccessibleTextInsertEvent::QAccessibleTextInsertEvent")

	return NewQAccessibleTextInsertEventFromPointer(C.QAccessibleTextInsertEvent_NewQAccessibleTextInsertEvent2(PointerFromQAccessibleInterface(iface), C.int(position), C.CString(text)))
}

func NewQAccessibleTextInsertEvent(object core.QObject_ITF, position int, text string) *QAccessibleTextInsertEvent {
	defer qt.Recovering("QAccessibleTextInsertEvent::QAccessibleTextInsertEvent")

	return NewQAccessibleTextInsertEventFromPointer(C.QAccessibleTextInsertEvent_NewQAccessibleTextInsertEvent(core.PointerFromQObject(object), C.int(position), C.CString(text)))
}

func (ptr *QAccessibleTextInsertEvent) ChangePosition() int {
	defer qt.Recovering("QAccessibleTextInsertEvent::changePosition")

	if ptr.Pointer() != nil {
		return int(C.QAccessibleTextInsertEvent_ChangePosition(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccessibleTextInsertEvent) TextInserted() string {
	defer qt.Recovering("QAccessibleTextInsertEvent::textInserted")

	if ptr.Pointer() != nil {
		return C.GoString(C.QAccessibleTextInsertEvent_TextInserted(ptr.Pointer()))
	}
	return ""
}
