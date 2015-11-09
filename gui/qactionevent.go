package gui

//#include "qactionevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QActionEvent struct {
	core.QEvent
}

type QActionEvent_ITF interface {
	core.QEvent_ITF
	QActionEvent_PTR() *QActionEvent
}

func PointerFromQActionEvent(ptr QActionEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QActionEvent_PTR().Pointer()
	}
	return nil
}

func NewQActionEventFromPointer(ptr unsafe.Pointer) *QActionEvent {
	var n = new(QActionEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QActionEvent) QActionEvent_PTR() *QActionEvent {
	return ptr
}
