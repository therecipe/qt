package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAccessibleTextCursorEvent struct {
	QAccessibleEvent
}

type QAccessibleTextCursorEvent_ITF interface {
	QAccessibleEvent_ITF
	QAccessibleTextCursorEvent_PTR() *QAccessibleTextCursorEvent
}

func PointerFromQAccessibleTextCursorEvent(ptr QAccessibleTextCursorEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccessibleTextCursorEvent_PTR().Pointer()
	}
	return nil
}

func NewQAccessibleTextCursorEventFromPointer(ptr unsafe.Pointer) *QAccessibleTextCursorEvent {
	var n = new(QAccessibleTextCursorEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAccessibleTextCursorEvent) QAccessibleTextCursorEvent_PTR() *QAccessibleTextCursorEvent {
	return ptr
}

func NewQAccessibleTextCursorEvent2(iface QAccessibleInterface_ITF, cursorPos int) *QAccessibleTextCursorEvent {
	defer qt.Recovering("QAccessibleTextCursorEvent::QAccessibleTextCursorEvent")

	return NewQAccessibleTextCursorEventFromPointer(C.QAccessibleTextCursorEvent_NewQAccessibleTextCursorEvent2(PointerFromQAccessibleInterface(iface), C.int(cursorPos)))
}

func NewQAccessibleTextCursorEvent(object core.QObject_ITF, cursorPos int) *QAccessibleTextCursorEvent {
	defer qt.Recovering("QAccessibleTextCursorEvent::QAccessibleTextCursorEvent")

	return NewQAccessibleTextCursorEventFromPointer(C.QAccessibleTextCursorEvent_NewQAccessibleTextCursorEvent(core.PointerFromQObject(object), C.int(cursorPos)))
}

func (ptr *QAccessibleTextCursorEvent) CursorPosition() int {
	defer qt.Recovering("QAccessibleTextCursorEvent::cursorPosition")

	if ptr.Pointer() != nil {
		return int(C.QAccessibleTextCursorEvent_CursorPosition(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAccessibleTextCursorEvent) SetCursorPosition(position int) {
	defer qt.Recovering("QAccessibleTextCursorEvent::setCursorPosition")

	if ptr.Pointer() != nil {
		C.QAccessibleTextCursorEvent_SetCursorPosition(ptr.Pointer(), C.int(position))
	}
}
