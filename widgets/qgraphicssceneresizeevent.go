package widgets

//#include "qgraphicssceneresizeevent.h"
import "C"
import (
	"unsafe"
)

type QGraphicsSceneResizeEvent struct {
	QGraphicsSceneEvent
}

type QGraphicsSceneResizeEventITF interface {
	QGraphicsSceneEventITF
	QGraphicsSceneResizeEventPTR() *QGraphicsSceneResizeEvent
}

func PointerFromQGraphicsSceneResizeEvent(ptr QGraphicsSceneResizeEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsSceneResizeEventPTR().Pointer()
	}
	return nil
}

func QGraphicsSceneResizeEventFromPointer(ptr unsafe.Pointer) *QGraphicsSceneResizeEvent {
	var n = new(QGraphicsSceneResizeEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGraphicsSceneResizeEvent) QGraphicsSceneResizeEventPTR() *QGraphicsSceneResizeEvent {
	return ptr
}

func NewQGraphicsSceneResizeEvent() *QGraphicsSceneResizeEvent {
	return QGraphicsSceneResizeEventFromPointer(unsafe.Pointer(C.QGraphicsSceneResizeEvent_NewQGraphicsSceneResizeEvent()))
}

func (ptr *QGraphicsSceneResizeEvent) DestroyQGraphicsSceneResizeEvent() {
	if ptr.Pointer() != nil {
		C.QGraphicsSceneResizeEvent_DestroyQGraphicsSceneResizeEvent(C.QtObjectPtr(ptr.Pointer()))
	}
}
