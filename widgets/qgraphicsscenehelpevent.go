package widgets

//#include "qgraphicsscenehelpevent.h"
import "C"
import (
	"unsafe"
)

type QGraphicsSceneHelpEvent struct {
	QGraphicsSceneEvent
}

type QGraphicsSceneHelpEventITF interface {
	QGraphicsSceneEventITF
	QGraphicsSceneHelpEventPTR() *QGraphicsSceneHelpEvent
}

func PointerFromQGraphicsSceneHelpEvent(ptr QGraphicsSceneHelpEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsSceneHelpEventPTR().Pointer()
	}
	return nil
}

func QGraphicsSceneHelpEventFromPointer(ptr unsafe.Pointer) *QGraphicsSceneHelpEvent {
	var n = new(QGraphicsSceneHelpEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGraphicsSceneHelpEvent) QGraphicsSceneHelpEventPTR() *QGraphicsSceneHelpEvent {
	return ptr
}

func (ptr *QGraphicsSceneHelpEvent) DestroyQGraphicsSceneHelpEvent() {
	if ptr.Pointer() != nil {
		C.QGraphicsSceneHelpEvent_DestroyQGraphicsSceneHelpEvent(C.QtObjectPtr(ptr.Pointer()))
	}
}
