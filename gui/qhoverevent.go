package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QHoverEvent struct {
	QInputEvent
}

type QHoverEvent_ITF interface {
	QInputEvent_ITF
	QHoverEvent_PTR() *QHoverEvent
}

func PointerFromQHoverEvent(ptr QHoverEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHoverEvent_PTR().Pointer()
	}
	return nil
}

func NewQHoverEventFromPointer(ptr unsafe.Pointer) *QHoverEvent {
	var n = new(QHoverEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QHoverEvent) QHoverEvent_PTR() *QHoverEvent {
	return ptr
}

func NewQHoverEvent(ty core.QEvent__Type, pos core.QPointF_ITF, oldPos core.QPointF_ITF, modifiers core.Qt__KeyboardModifier) *QHoverEvent {
	defer qt.Recovering("QHoverEvent::QHoverEvent")

	return NewQHoverEventFromPointer(C.QHoverEvent_NewQHoverEvent(C.int(ty), core.PointerFromQPointF(pos), core.PointerFromQPointF(oldPos), C.int(modifiers)))
}

func (ptr *QHoverEvent) OldPos() *core.QPoint {
	defer qt.Recovering("QHoverEvent::oldPos")

	if ptr.Pointer() != nil {
		return core.NewQPointFromPointer(C.QHoverEvent_OldPos(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHoverEvent) Pos() *core.QPoint {
	defer qt.Recovering("QHoverEvent::pos")

	if ptr.Pointer() != nil {
		return core.NewQPointFromPointer(C.QHoverEvent_Pos(ptr.Pointer()))
	}
	return nil
}
