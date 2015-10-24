package gui

//#include "qfocusevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QFocusEvent struct {
	core.QEvent
}

type QFocusEventITF interface {
	core.QEventITF
	QFocusEventPTR() *QFocusEvent
}

func PointerFromQFocusEvent(ptr QFocusEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QFocusEventPTR().Pointer()
	}
	return nil
}

func QFocusEventFromPointer(ptr unsafe.Pointer) *QFocusEvent {
	var n = new(QFocusEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QFocusEvent) QFocusEventPTR() *QFocusEvent {
	return ptr
}

func NewQFocusEvent(ty core.QEvent__Type, reason core.Qt__FocusReason) *QFocusEvent {
	return QFocusEventFromPointer(unsafe.Pointer(C.QFocusEvent_NewQFocusEvent(C.int(ty), C.int(reason))))
}

func (ptr *QFocusEvent) GotFocus() bool {
	if ptr.Pointer() != nil {
		return C.QFocusEvent_GotFocus(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QFocusEvent) LostFocus() bool {
	if ptr.Pointer() != nil {
		return C.QFocusEvent_LostFocus(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QFocusEvent) Reason() core.Qt__FocusReason {
	if ptr.Pointer() != nil {
		return core.Qt__FocusReason(C.QFocusEvent_Reason(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}
