package widgets

//#include "qgraphicsscenehoverevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QGraphicsSceneHoverEvent struct {
	QGraphicsSceneEvent
}

type QGraphicsSceneHoverEventITF interface {
	QGraphicsSceneEventITF
	QGraphicsSceneHoverEventPTR() *QGraphicsSceneHoverEvent
}

func PointerFromQGraphicsSceneHoverEvent(ptr QGraphicsSceneHoverEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsSceneHoverEventPTR().Pointer()
	}
	return nil
}

func QGraphicsSceneHoverEventFromPointer(ptr unsafe.Pointer) *QGraphicsSceneHoverEvent {
	var n = new(QGraphicsSceneHoverEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGraphicsSceneHoverEvent) QGraphicsSceneHoverEventPTR() *QGraphicsSceneHoverEvent {
	return ptr
}

func (ptr *QGraphicsSceneHoverEvent) Modifiers() core.Qt__KeyboardModifier {
	if ptr.Pointer() != nil {
		return core.Qt__KeyboardModifier(C.QGraphicsSceneHoverEvent_Modifiers(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsSceneHoverEvent) DestroyQGraphicsSceneHoverEvent() {
	if ptr.Pointer() != nil {
		C.QGraphicsSceneHoverEvent_DestroyQGraphicsSceneHoverEvent(C.QtObjectPtr(ptr.Pointer()))
	}
}
