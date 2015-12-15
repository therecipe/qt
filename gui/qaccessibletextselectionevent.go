package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAccessibleTextSelectionEvent struct {
	QAccessibleTextCursorEvent
}

type QAccessibleTextSelectionEvent_ITF interface {
	QAccessibleTextCursorEvent_ITF
	QAccessibleTextSelectionEvent_PTR() *QAccessibleTextSelectionEvent
}

func PointerFromQAccessibleTextSelectionEvent(ptr QAccessibleTextSelectionEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccessibleTextSelectionEvent_PTR().Pointer()
	}
	return nil
}

func NewQAccessibleTextSelectionEventFromPointer(ptr unsafe.Pointer) *QAccessibleTextSelectionEvent {
	var n = new(QAccessibleTextSelectionEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAccessibleTextSelectionEvent) QAccessibleTextSelectionEvent_PTR() *QAccessibleTextSelectionEvent {
	return ptr
}

func NewQAccessibleTextSelectionEvent2(iface QAccessibleInterface_ITF, start int, end int) *QAccessibleTextSelectionEvent {
	defer qt.Recovering("QAccessibleTextSelectionEvent::QAccessibleTextSelectionEvent")

	return NewQAccessibleTextSelectionEventFromPointer(C.QAccessibleTextSelectionEvent_NewQAccessibleTextSelectionEvent2(PointerFromQAccessibleInterface(iface), C.int(start), C.int(end)))
}

func NewQAccessibleTextSelectionEvent(object core.QObject_ITF, start int, end int) *QAccessibleTextSelectionEvent {
	defer qt.Recovering("QAccessibleTextSelectionEvent::QAccessibleTextSelectionEvent")

	return NewQAccessibleTextSelectionEventFromPointer(C.QAccessibleTextSelectionEvent_NewQAccessibleTextSelectionEvent(core.PointerFromQObject(object), C.int(start), C.int(end)))
}

func (ptr *QAccessibleTextSelectionEvent) SelectionEnd() int {
	defer qt.Recovering("QAccessibleTextSelectionEvent::selectionEnd")

	if ptr.Pointer() != nil {
		return int(C.QAccessibleTextSelectionEvent_SelectionEnd(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccessibleTextSelectionEvent) SelectionStart() int {
	defer qt.Recovering("QAccessibleTextSelectionEvent::selectionStart")

	if ptr.Pointer() != nil {
		return int(C.QAccessibleTextSelectionEvent_SelectionStart(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccessibleTextSelectionEvent) SetSelection(start int, end int) {
	defer qt.Recovering("QAccessibleTextSelectionEvent::setSelection")

	if ptr.Pointer() != nil {
		C.QAccessibleTextSelectionEvent_SetSelection(ptr.Pointer(), C.int(start), C.int(end))
	}
}
