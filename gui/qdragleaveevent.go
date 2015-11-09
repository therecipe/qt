package gui

//#include "qdragleaveevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QDragLeaveEvent struct {
	core.QEvent
}

type QDragLeaveEvent_ITF interface {
	core.QEvent_ITF
	QDragLeaveEvent_PTR() *QDragLeaveEvent
}

func PointerFromQDragLeaveEvent(ptr QDragLeaveEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDragLeaveEvent_PTR().Pointer()
	}
	return nil
}

func NewQDragLeaveEventFromPointer(ptr unsafe.Pointer) *QDragLeaveEvent {
	var n = new(QDragLeaveEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDragLeaveEvent) QDragLeaveEvent_PTR() *QDragLeaveEvent {
	return ptr
}

func NewQDragLeaveEvent() *QDragLeaveEvent {
	return NewQDragLeaveEventFromPointer(C.QDragLeaveEvent_NewQDragLeaveEvent())
}
