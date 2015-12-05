package gui

//#include "gui.h"
import "C"
import (
	"unsafe"
)

type QAccessibleStateChangeEvent struct {
	QAccessibleEvent
}

type QAccessibleStateChangeEvent_ITF interface {
	QAccessibleEvent_ITF
	QAccessibleStateChangeEvent_PTR() *QAccessibleStateChangeEvent
}

func PointerFromQAccessibleStateChangeEvent(ptr QAccessibleStateChangeEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccessibleStateChangeEvent_PTR().Pointer()
	}
	return nil
}

func NewQAccessibleStateChangeEventFromPointer(ptr unsafe.Pointer) *QAccessibleStateChangeEvent {
	var n = new(QAccessibleStateChangeEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAccessibleStateChangeEvent) QAccessibleStateChangeEvent_PTR() *QAccessibleStateChangeEvent {
	return ptr
}
