package widgets

//#include "qgraphicsscenemoveevent.h"
import "C"
import (
	"unsafe"
)

type QGraphicsSceneMoveEvent struct {
	QGraphicsSceneEvent
}

type QGraphicsSceneMoveEventITF interface {
	QGraphicsSceneEventITF
	QGraphicsSceneMoveEventPTR() *QGraphicsSceneMoveEvent
}

func PointerFromQGraphicsSceneMoveEvent(ptr QGraphicsSceneMoveEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsSceneMoveEventPTR().Pointer()
	}
	return nil
}

func QGraphicsSceneMoveEventFromPointer(ptr unsafe.Pointer) *QGraphicsSceneMoveEvent {
	var n = new(QGraphicsSceneMoveEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGraphicsSceneMoveEvent) QGraphicsSceneMoveEventPTR() *QGraphicsSceneMoveEvent {
	return ptr
}

func NewQGraphicsSceneMoveEvent() *QGraphicsSceneMoveEvent {
	return QGraphicsSceneMoveEventFromPointer(unsafe.Pointer(C.QGraphicsSceneMoveEvent_NewQGraphicsSceneMoveEvent()))
}

func (ptr *QGraphicsSceneMoveEvent) DestroyQGraphicsSceneMoveEvent() {
	if ptr.Pointer() != nil {
		C.QGraphicsSceneMoveEvent_DestroyQGraphicsSceneMoveEvent(C.QtObjectPtr(ptr.Pointer()))
	}
}
