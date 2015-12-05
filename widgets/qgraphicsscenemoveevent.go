package widgets

//#include "widgets.h"
import "C"
import (
	"log"
	"unsafe"
)

type QGraphicsSceneMoveEvent struct {
	QGraphicsSceneEvent
}

type QGraphicsSceneMoveEvent_ITF interface {
	QGraphicsSceneEvent_ITF
	QGraphicsSceneMoveEvent_PTR() *QGraphicsSceneMoveEvent
}

func PointerFromQGraphicsSceneMoveEvent(ptr QGraphicsSceneMoveEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsSceneMoveEvent_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsSceneMoveEventFromPointer(ptr unsafe.Pointer) *QGraphicsSceneMoveEvent {
	var n = new(QGraphicsSceneMoveEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGraphicsSceneMoveEvent) QGraphicsSceneMoveEvent_PTR() *QGraphicsSceneMoveEvent {
	return ptr
}

func NewQGraphicsSceneMoveEvent() *QGraphicsSceneMoveEvent {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsSceneMoveEvent::QGraphicsSceneMoveEvent")
		}
	}()

	return NewQGraphicsSceneMoveEventFromPointer(C.QGraphicsSceneMoveEvent_NewQGraphicsSceneMoveEvent())
}

func (ptr *QGraphicsSceneMoveEvent) DestroyQGraphicsSceneMoveEvent() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsSceneMoveEvent::~QGraphicsSceneMoveEvent")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsSceneMoveEvent_DestroyQGraphicsSceneMoveEvent(ptr.Pointer())
	}
}
