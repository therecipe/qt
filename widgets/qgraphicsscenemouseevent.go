package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
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
	defer qt.Recovering("QGraphicsSceneMouseEvent::button")

	if ptr.Pointer() != nil {
		return core.Qt__MouseButton(C.QGraphicsSceneMouseEvent_Button(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsSceneMouseEvent) ButtonDownScreenPos(button core.Qt__MouseButton) *core.QPoint {
	defer qt.Recovering("QGraphicsSceneMouseEvent::buttonDownScreenPos")

	if ptr.Pointer() != nil {
		return core.NewQPointFromPointer(C.QGraphicsSceneMouseEvent_ButtonDownScreenPos(ptr.Pointer(), C.int(button)))
	}
	return nil
}

func (ptr *QGraphicsSceneMouseEvent) Buttons() core.Qt__MouseButton {
	defer qt.Recovering("QGraphicsSceneMouseEvent::buttons")

	if ptr.Pointer() != nil {
		return core.Qt__MouseButton(C.QGraphicsSceneMouseEvent_Buttons(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsSceneMouseEvent) Flags() core.Qt__MouseEventFlag {
	defer qt.Recovering("QGraphicsSceneMouseEvent::flags")

	if ptr.Pointer() != nil {
		return core.Qt__MouseEventFlag(C.QGraphicsSceneMouseEvent_Flags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsSceneMouseEvent) LastScreenPos() *core.QPoint {
	defer qt.Recovering("QGraphicsSceneMouseEvent::lastScreenPos")

	if ptr.Pointer() != nil {
		return core.NewQPointFromPointer(C.QGraphicsSceneMouseEvent_LastScreenPos(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsSceneMouseEvent) Modifiers() core.Qt__KeyboardModifier {
	defer qt.Recovering("QGraphicsSceneMouseEvent::modifiers")

	if ptr.Pointer() != nil {
		return core.Qt__KeyboardModifier(C.QGraphicsSceneMouseEvent_Modifiers(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsSceneMouseEvent) ScreenPos() *core.QPoint {
	defer qt.Recovering("QGraphicsSceneMouseEvent::screenPos")

	if ptr.Pointer() != nil {
		return core.NewQPointFromPointer(C.QGraphicsSceneMouseEvent_ScreenPos(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsSceneMouseEvent) Source() core.Qt__MouseEventSource {
	defer qt.Recovering("QGraphicsSceneMouseEvent::source")

	if ptr.Pointer() != nil {
		return core.Qt__MouseEventSource(C.QGraphicsSceneMouseEvent_Source(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsSceneMouseEvent) DestroyQGraphicsSceneMouseEvent() {
	defer qt.Recovering("QGraphicsSceneMouseEvent::~QGraphicsSceneMouseEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSceneMouseEvent_DestroyQGraphicsSceneMouseEvent(ptr.Pointer())
	}
}
