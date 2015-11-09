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

type QGraphicsSceneMouseEvent_ITF interface {
	QGraphicsSceneEvent_ITF
	QGraphicsSceneMouseEvent_PTR() *QGraphicsSceneMouseEvent
}

func PointerFromQGraphicsSceneMouseEvent(ptr QGraphicsSceneMouseEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsSceneMouseEvent_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsSceneMouseEventFromPointer(ptr unsafe.Pointer) *QGraphicsSceneMouseEvent {
	var n = new(QGraphicsSceneMouseEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGraphicsSceneMouseEvent) QGraphicsSceneMouseEvent_PTR() *QGraphicsSceneMouseEvent {
	return ptr
}

func (ptr *QGraphicsSceneMouseEvent) Button() core.Qt__MouseButton {
	if ptr.Pointer() != nil {
		return core.Qt__MouseButton(C.QGraphicsSceneMouseEvent_Button(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsSceneMouseEvent) Buttons() core.Qt__MouseButton {
	if ptr.Pointer() != nil {
		return core.Qt__MouseButton(C.QGraphicsSceneMouseEvent_Buttons(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsSceneMouseEvent) Flags() core.Qt__MouseEventFlag {
	if ptr.Pointer() != nil {
		return core.Qt__MouseEventFlag(C.QGraphicsSceneMouseEvent_Flags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsSceneMouseEvent) Modifiers() core.Qt__KeyboardModifier {
	if ptr.Pointer() != nil {
		return core.Qt__KeyboardModifier(C.QGraphicsSceneMouseEvent_Modifiers(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsSceneMouseEvent) Source() core.Qt__MouseEventSource {
	if ptr.Pointer() != nil {
		return core.Qt__MouseEventSource(C.QGraphicsSceneMouseEvent_Source(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsSceneMouseEvent) DestroyQGraphicsSceneMouseEvent() {
	if ptr.Pointer() != nil {
		C.QGraphicsSceneMouseEvent_DestroyQGraphicsSceneMouseEvent(ptr.Pointer())
	}
}
