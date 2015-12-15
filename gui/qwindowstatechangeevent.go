package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QWindowStateChangeEvent struct {
	core.QEvent
}

type QWindowStateChangeEvent_ITF interface {
	core.QEvent_ITF
	QWindowStateChangeEvent_PTR() *QWindowStateChangeEvent
}

func PointerFromQWindowStateChangeEvent(ptr QWindowStateChangeEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWindowStateChangeEvent_PTR().Pointer()
	}
	return nil
}

func NewQWindowStateChangeEventFromPointer(ptr unsafe.Pointer) *QWindowStateChangeEvent {
	var n = new(QWindowStateChangeEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QWindowStateChangeEvent) QWindowStateChangeEvent_PTR() *QWindowStateChangeEvent {
	return ptr
}

func (ptr *QWindowStateChangeEvent) OldState() core.Qt__WindowState {
	defer qt.Recovering("QWindowStateChangeEvent::oldState")

	if ptr.Pointer() != nil {
		return core.Qt__WindowState(C.QWindowStateChangeEvent_OldState(ptr.Pointer()))
	}
	return 0
}
