package gui

//#include "qaccessibletextselectionevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAccessibleTextSelectionEvent struct {
	QAccessibleTextCursorEvent
}

type QAccessibleTextSelectionEventITF interface {
	QAccessibleTextCursorEventITF
	QAccessibleTextSelectionEventPTR() *QAccessibleTextSelectionEvent
}

func PointerFromQAccessibleTextSelectionEvent(ptr QAccessibleTextSelectionEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccessibleTextSelectionEventPTR().Pointer()
	}
	return nil
}

func QAccessibleTextSelectionEventFromPointer(ptr unsafe.Pointer) *QAccessibleTextSelectionEvent {
	var n = new(QAccessibleTextSelectionEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAccessibleTextSelectionEvent) QAccessibleTextSelectionEventPTR() *QAccessibleTextSelectionEvent {
	return ptr
}

func NewQAccessibleTextSelectionEvent2(iface QAccessibleInterfaceITF, start int, end int) *QAccessibleTextSelectionEvent {
	return QAccessibleTextSelectionEventFromPointer(unsafe.Pointer(C.QAccessibleTextSelectionEvent_NewQAccessibleTextSelectionEvent2(C.QtObjectPtr(PointerFromQAccessibleInterface(iface)), C.int(start), C.int(end))))
}

func NewQAccessibleTextSelectionEvent(object core.QObjectITF, start int, end int) *QAccessibleTextSelectionEvent {
	return QAccessibleTextSelectionEventFromPointer(unsafe.Pointer(C.QAccessibleTextSelectionEvent_NewQAccessibleTextSelectionEvent(C.QtObjectPtr(core.PointerFromQObject(object)), C.int(start), C.int(end))))
}

func (ptr *QAccessibleTextSelectionEvent) SelectionEnd() int {
	if ptr.Pointer() != nil {
		return int(C.QAccessibleTextSelectionEvent_SelectionEnd(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAccessibleTextSelectionEvent) SelectionStart() int {
	if ptr.Pointer() != nil {
		return int(C.QAccessibleTextSelectionEvent_SelectionStart(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAccessibleTextSelectionEvent) SetSelection(start int, end int) {
	if ptr.Pointer() != nil {
		C.QAccessibleTextSelectionEvent_SetSelection(C.QtObjectPtr(ptr.Pointer()), C.int(start), C.int(end))
	}
}
