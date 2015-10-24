package gui

//#include "qwindowstatechangeevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QWindowStateChangeEvent struct {
	core.QEvent
}

type QWindowStateChangeEventITF interface {
	core.QEventITF
	QWindowStateChangeEventPTR() *QWindowStateChangeEvent
}

func PointerFromQWindowStateChangeEvent(ptr QWindowStateChangeEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWindowStateChangeEventPTR().Pointer()
	}
	return nil
}

func QWindowStateChangeEventFromPointer(ptr unsafe.Pointer) *QWindowStateChangeEvent {
	var n = new(QWindowStateChangeEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QWindowStateChangeEvent) QWindowStateChangeEventPTR() *QWindowStateChangeEvent {
	return ptr
}

func (ptr *QWindowStateChangeEvent) OldState() core.Qt__WindowState {
	if ptr.Pointer() != nil {
		return core.Qt__WindowState(C.QWindowStateChangeEvent_OldState(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}
