package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
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
	defer qt.Recovering("QEnterEvent::QEnterEvent")

	return NewQEnterEventFromPointer(C.QEnterEvent_NewQEnterEvent(core.PointerFromQPointF(localPos), core.PointerFromQPointF(windowPos), core.PointerFromQPointF(screenPos)))
}

func (ptr *QEnterEvent) GlobalPos() *core.QPoint {
	defer qt.Recovering("QEnterEvent::globalPos")

	if ptr.Pointer() != nil {
		return core.NewQPointFromPointer(C.QEnterEvent_GlobalPos(ptr.Pointer()))
	}
	return nil
}

func (ptr *QEnterEvent) GlobalX() int {
	defer qt.Recovering("QEnterEvent::globalX")

	if ptr.Pointer() != nil {
		return int(C.QEnterEvent_GlobalX(ptr.Pointer()))
	}
	return 0
}

func (ptr *QEnterEvent) GlobalY() int {
	defer qt.Recovering("QEnterEvent::globalY")

	if ptr.Pointer() != nil {
		return int(C.QEnterEvent_GlobalY(ptr.Pointer()))
	}
	return 0
}

func (ptr *QEnterEvent) Pos() *core.QPoint {
	defer qt.Recovering("QEnterEvent::pos")

	if ptr.Pointer() != nil {
		return core.NewQPointFromPointer(C.QEnterEvent_Pos(ptr.Pointer()))
	}
	return nil
}

func (ptr *QEnterEvent) X() int {
	defer qt.Recovering("QEnterEvent::x")

	if ptr.Pointer() != nil {
		return int(C.QEnterEvent_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QEnterEvent) Y() int {
	defer qt.Recovering("QEnterEvent::y")

	if ptr.Pointer() != nil {
		return int(C.QEnterEvent_Y(ptr.Pointer()))
	}
	return 0
}
