package gui

//#include "qaccessibletextcursorevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAccessibleTextCursorEvent struct {
	QAccessibleEvent
}

type QAccessibleTextCursorEventITF interface {
	QAccessibleEventITF
	QAccessibleTextCursorEventPTR() *QAccessibleTextCursorEvent
}

func PointerFromQAccessibleTextCursorEvent(ptr QAccessibleTextCursorEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccessibleTextCursorEventPTR().Pointer()
	}
	return nil
}

func QAccessibleTextCursorEventFromPointer(ptr unsafe.Pointer) *QAccessibleTextCursorEvent {
	var n = new(QAccessibleTextCursorEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAccessibleTextCursorEvent) QAccessibleTextCursorEventPTR() *QAccessibleTextCursorEvent {
	return ptr
}

func NewQAccessibleTextCursorEvent2(iface QAccessibleInterfaceITF, cursorPos int) *QAccessibleTextCursorEvent {
	return QAccessibleTextCursorEventFromPointer(unsafe.Pointer(C.QAccessibleTextCursorEvent_NewQAccessibleTextCursorEvent2(C.QtObjectPtr(PointerFromQAccessibleInterface(iface)), C.int(cursorPos))))
}

func NewQAccessibleTextCursorEvent(object core.QObjectITF, cursorPos int) *QAccessibleTextCursorEvent {
	return QAccessibleTextCursorEventFromPointer(unsafe.Pointer(C.QAccessibleTextCursorEvent_NewQAccessibleTextCursorEvent(C.QtObjectPtr(core.PointerFromQObject(object)), C.int(cursorPos))))
}

func (ptr *QAccessibleTextCursorEvent) CursorPosition() int {
	if ptr.Pointer() != nil {
		return int(C.QAccessibleTextCursorEvent_CursorPosition(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAccessibleTextCursorEvent) SetCursorPosition(position int) {
	if ptr.Pointer() != nil {
		C.QAccessibleTextCursorEvent_SetCursorPosition(C.QtObjectPtr(ptr.Pointer()), C.int(position))
	}
}
