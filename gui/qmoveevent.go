package gui

//#include "qmoveevent.h"
import "C"
import (
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
	return NewQMoveEventFromPointer(C.QMoveEvent_NewQMoveEvent(core.PointerFromQPoint(pos), core.PointerFromQPoint(oldPos)))
}
