package gui

//#include "qenterevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QEnterEvent struct {
	core.QEvent
}

type QEnterEvent_ITF interface {
	core.QEvent_ITF
	QEnterEvent_PTR() *QEnterEvent
}

func PointerFromQEnterEvent(ptr QEnterEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QEnterEvent_PTR().Pointer()
	}
	return nil
}

func NewQEnterEventFromPointer(ptr unsafe.Pointer) *QEnterEvent {
	var n = new(QEnterEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QEnterEvent) QEnterEvent_PTR() *QEnterEvent {
	return ptr
}

func NewQEnterEvent(localPos core.QPointF_ITF, windowPos core.QPointF_ITF, screenPos core.QPointF_ITF) *QEnterEvent {
	return NewQEnterEventFromPointer(C.QEnterEvent_NewQEnterEvent(core.PointerFromQPointF(localPos), core.PointerFromQPointF(windowPos), core.PointerFromQPointF(screenPos)))
}

func (ptr *QEnterEvent) GlobalX() int {
	if ptr.Pointer() != nil {
		return int(C.QEnterEvent_GlobalX(ptr.Pointer()))
	}
	return 0
}

func (ptr *QEnterEvent) GlobalY() int {
	if ptr.Pointer() != nil {
		return int(C.QEnterEvent_GlobalY(ptr.Pointer()))
	}
	return 0
}

func (ptr *QEnterEvent) X() int {
	if ptr.Pointer() != nil {
		return int(C.QEnterEvent_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QEnterEvent) Y() int {
	if ptr.Pointer() != nil {
		return int(C.QEnterEvent_Y(ptr.Pointer()))
	}
	return 0
}
