package gui

//#include "qaccessibletextremoveevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QAccessibleTextRemoveEvent struct {
	QAccessibleTextCursorEvent
}

type QAccessibleTextRemoveEventITF interface {
	QAccessibleTextCursorEventITF
	QAccessibleTextRemoveEventPTR() *QAccessibleTextRemoveEvent
}

func PointerFromQAccessibleTextRemoveEvent(ptr QAccessibleTextRemoveEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccessibleTextRemoveEventPTR().Pointer()
	}
	return nil
}

func QAccessibleTextRemoveEventFromPointer(ptr unsafe.Pointer) *QAccessibleTextRemoveEvent {
	var n = new(QAccessibleTextRemoveEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAccessibleTextRemoveEvent) QAccessibleTextRemoveEventPTR() *QAccessibleTextRemoveEvent {
	return ptr
}

func NewQAccessibleTextRemoveEvent2(iface QAccessibleInterfaceITF, position int, text string) *QAccessibleTextRemoveEvent {
	return QAccessibleTextRemoveEventFromPointer(unsafe.Pointer(C.QAccessibleTextRemoveEvent_NewQAccessibleTextRemoveEvent2(C.QtObjectPtr(PointerFromQAccessibleInterface(iface)), C.int(position), C.CString(text))))
}

func NewQAccessibleTextRemoveEvent(object core.QObjectITF, position int, text string) *QAccessibleTextRemoveEvent {
	return QAccessibleTextRemoveEventFromPointer(unsafe.Pointer(C.QAccessibleTextRemoveEvent_NewQAccessibleTextRemoveEvent(C.QtObjectPtr(core.PointerFromQObject(object)), C.int(position), C.CString(text))))
}

func (ptr *QAccessibleTextRemoveEvent) ChangePosition() int {
	if ptr.Pointer() != nil {
		return int(C.QAccessibleTextRemoveEvent_ChangePosition(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAccessibleTextRemoveEvent) TextRemoved() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAccessibleTextRemoveEvent_TextRemoved(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}
