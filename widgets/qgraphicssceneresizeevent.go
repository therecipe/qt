package widgets

//#include "widgets.h"
import "C"
import (
	"log"
	"unsafe"
)

type QGraphicsSceneResizeEvent struct {
	QGraphicsSceneEvent
}

type QGraphicsSceneResizeEvent_ITF interface {
	QGraphicsSceneEvent_ITF
	QGraphicsSceneResizeEvent_PTR() *QGraphicsSceneResizeEvent
}

func PointerFromQGraphicsSceneResizeEvent(ptr QGraphicsSceneResizeEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsSceneResizeEvent_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsSceneResizeEventFromPointer(ptr unsafe.Pointer) *QGraphicsSceneResizeEvent {
	var n = new(QGraphicsSceneResizeEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGraphicsSceneResizeEvent) QGraphicsSceneResizeEvent_PTR() *QGraphicsSceneResizeEvent {
	return ptr
}

func NewQGraphicsSceneResizeEvent() *QGraphicsSceneResizeEvent {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsSceneResizeEvent::QGraphicsSceneResizeEvent")
		}
	}()

	return NewQGraphicsSceneResizeEventFromPointer(C.QGraphicsSceneResizeEvent_NewQGraphicsSceneResizeEvent())
}

func (ptr *QGraphicsSceneResizeEvent) DestroyQGraphicsSceneResizeEvent() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsSceneResizeEvent::~QGraphicsSceneResizeEvent")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsSceneResizeEvent_DestroyQGraphicsSceneResizeEvent(ptr.Pointer())
	}
}
