package widgets

//#include "qgraphicsscenewheelevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QGraphicsSceneWheelEvent struct {
	QGraphicsSceneEvent
}

type QGraphicsSceneWheelEventITF interface {
	QGraphicsSceneEventITF
	QGraphicsSceneWheelEventPTR() *QGraphicsSceneWheelEvent
}

func PointerFromQGraphicsSceneWheelEvent(ptr QGraphicsSceneWheelEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsSceneWheelEventPTR().Pointer()
	}
	return nil
}

func QGraphicsSceneWheelEventFromPointer(ptr unsafe.Pointer) *QGraphicsSceneWheelEvent {
	var n = new(QGraphicsSceneWheelEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGraphicsSceneWheelEvent) QGraphicsSceneWheelEventPTR() *QGraphicsSceneWheelEvent {
	return ptr
}

func (ptr *QGraphicsSceneWheelEvent) Buttons() core.Qt__MouseButton {
	if ptr.Pointer() != nil {
		return core.Qt__MouseButton(C.QGraphicsSceneWheelEvent_Buttons(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsSceneWheelEvent) Delta() int {
	if ptr.Pointer() != nil {
		return int(C.QGraphicsSceneWheelEvent_Delta(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsSceneWheelEvent) Modifiers() core.Qt__KeyboardModifier {
	if ptr.Pointer() != nil {
		return core.Qt__KeyboardModifier(C.QGraphicsSceneWheelEvent_Modifiers(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsSceneWheelEvent) Orientation() core.Qt__Orientation {
	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QGraphicsSceneWheelEvent_Orientation(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsSceneWheelEvent) DestroyQGraphicsSceneWheelEvent() {
	if ptr.Pointer() != nil {
		C.QGraphicsSceneWheelEvent_DestroyQGraphicsSceneWheelEvent(C.QtObjectPtr(ptr.Pointer()))
	}
}
