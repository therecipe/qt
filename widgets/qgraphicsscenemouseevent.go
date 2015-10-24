package widgets

//#include "qgraphicsscenemouseevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QGraphicsSceneMouseEvent struct {
	QGraphicsSceneEvent
}

type QGraphicsSceneMouseEventITF interface {
	QGraphicsSceneEventITF
	QGraphicsSceneMouseEventPTR() *QGraphicsSceneMouseEvent
}

func PointerFromQGraphicsSceneMouseEvent(ptr QGraphicsSceneMouseEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsSceneMouseEventPTR().Pointer()
	}
	return nil
}

func QGraphicsSceneMouseEventFromPointer(ptr unsafe.Pointer) *QGraphicsSceneMouseEvent {
	var n = new(QGraphicsSceneMouseEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGraphicsSceneMouseEvent) QGraphicsSceneMouseEventPTR() *QGraphicsSceneMouseEvent {
	return ptr
}

func (ptr *QGraphicsSceneMouseEvent) Button() core.Qt__MouseButton {
	if ptr.Pointer() != nil {
		return core.Qt__MouseButton(C.QGraphicsSceneMouseEvent_Button(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsSceneMouseEvent) Buttons() core.Qt__MouseButton {
	if ptr.Pointer() != nil {
		return core.Qt__MouseButton(C.QGraphicsSceneMouseEvent_Buttons(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsSceneMouseEvent) Flags() core.Qt__MouseEventFlag {
	if ptr.Pointer() != nil {
		return core.Qt__MouseEventFlag(C.QGraphicsSceneMouseEvent_Flags(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsSceneMouseEvent) Modifiers() core.Qt__KeyboardModifier {
	if ptr.Pointer() != nil {
		return core.Qt__KeyboardModifier(C.QGraphicsSceneMouseEvent_Modifiers(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsSceneMouseEvent) Source() core.Qt__MouseEventSource {
	if ptr.Pointer() != nil {
		return core.Qt__MouseEventSource(C.QGraphicsSceneMouseEvent_Source(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QGraphicsSceneMouseEvent) DestroyQGraphicsSceneMouseEvent() {
	if ptr.Pointer() != nil {
		C.QGraphicsSceneMouseEvent_DestroyQGraphicsSceneMouseEvent(C.QtObjectPtr(ptr.Pointer()))
	}
}
