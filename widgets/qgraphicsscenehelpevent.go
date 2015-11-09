package widgets

//#include "qgraphicsscenehelpevent.h"
import "C"
import (
	"unsafe"
)

type QGraphicsSceneHelpEvent struct {
	QGraphicsSceneEvent
}

type QGraphicsSceneHelpEvent_ITF interface {
	QGraphicsSceneEvent_ITF
	QGraphicsSceneHelpEvent_PTR() *QGraphicsSceneHelpEvent
}

func PointerFromQGraphicsSceneHelpEvent(ptr QGraphicsSceneHelpEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsSceneHelpEvent_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsSceneHelpEventFromPointer(ptr unsafe.Pointer) *QGraphicsSceneHelpEvent {
	var n = new(QGraphicsSceneHelpEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGraphicsSceneHelpEvent) QGraphicsSceneHelpEvent_PTR() *QGraphicsSceneHelpEvent {
	return ptr
}

func (ptr *QGraphicsSceneHelpEvent) DestroyQGraphicsSceneHelpEvent() {
	if ptr.Pointer() != nil {
		C.QGraphicsSceneHelpEvent_DestroyQGraphicsSceneHelpEvent(ptr.Pointer())
	}
}
