package gui

//#include "qaccessiblestatechangeevent.h"
import "C"
import (
	"unsafe"
)

type QAccessibleStateChangeEvent struct {
	QAccessibleEvent
}

type QAccessibleStateChangeEventITF interface {
	QAccessibleEventITF
	QAccessibleStateChangeEventPTR() *QAccessibleStateChangeEvent
}

func PointerFromQAccessibleStateChangeEvent(ptr QAccessibleStateChangeEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAccessibleStateChangeEventPTR().Pointer()
	}
	return nil
}

func QAccessibleStateChangeEventFromPointer(ptr unsafe.Pointer) *QAccessibleStateChangeEvent {
	var n = new(QAccessibleStateChangeEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAccessibleStateChangeEvent) QAccessibleStateChangeEventPTR() *QAccessibleStateChangeEvent {
	return ptr
}
