package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QMoveEvent struct {
	core.QEvent
}

type QMoveEvent_ITF interface {
	core.QEvent_ITF
	QMoveEvent_PTR() *QMoveEvent
}

func PointerFromQMoveEvent(ptr QMoveEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMoveEvent_PTR().Pointer()
	}
	return nil
}

func NewQMoveEventFromPointer(ptr unsafe.Pointer) *QMoveEvent {
	var n = new(QMoveEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMoveEvent) QMoveEvent_PTR() *QMoveEvent {
	return ptr
}

func NewQMoveEvent(pos core.QPoint_ITF, oldPos core.QPoint_ITF) *QMoveEvent {
	defer qt.Recovering("QMoveEvent::QMoveEvent")

	return NewQMoveEventFromPointer(C.QMoveEvent_NewQMoveEvent(core.PointerFromQPoint(pos), core.PointerFromQPoint(oldPos)))
}

func (ptr *QMoveEvent) OldPos() *core.QPoint {
	defer qt.Recovering("QMoveEvent::oldPos")

	if ptr.Pointer() != nil {
		return core.NewQPointFromPointer(C.QMoveEvent_OldPos(ptr.Pointer()))
	}
	return nil
}

func (ptr *QMoveEvent) Pos() *core.QPoint {
	defer qt.Recovering("QMoveEvent::pos")

	if ptr.Pointer() != nil {
		return core.NewQPointFromPointer(C.QMoveEvent_Pos(ptr.Pointer()))
	}
	return nil
}
