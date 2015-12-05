package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QGraphicsSceneHoverEvent struct {
	QGraphicsSceneEvent
}

type QGraphicsSceneHoverEvent_ITF interface {
	QGraphicsSceneEvent_ITF
	QGraphicsSceneHoverEvent_PTR() *QGraphicsSceneHoverEvent
}

func PointerFromQGraphicsSceneHoverEvent(ptr QGraphicsSceneHoverEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsSceneHoverEvent_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsSceneHoverEventFromPointer(ptr unsafe.Pointer) *QGraphicsSceneHoverEvent {
	var n = new(QGraphicsSceneHoverEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGraphicsSceneHoverEvent) QGraphicsSceneHoverEvent_PTR() *QGraphicsSceneHoverEvent {
	return ptr
}

func (ptr *QGraphicsSceneHoverEvent) Modifiers() core.Qt__KeyboardModifier {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsSceneHoverEvent::modifiers")
		}
	}()

	if ptr.Pointer() != nil {
		return core.Qt__KeyboardModifier(C.QGraphicsSceneHoverEvent_Modifiers(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsSceneHoverEvent) DestroyQGraphicsSceneHoverEvent() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QGraphicsSceneHoverEvent::~QGraphicsSceneHoverEvent")
		}
	}()

	if ptr.Pointer() != nil {
		C.QGraphicsSceneHoverEvent_DestroyQGraphicsSceneHoverEvent(ptr.Pointer())
	}
}
