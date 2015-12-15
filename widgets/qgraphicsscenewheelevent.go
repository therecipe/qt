package widgets

//#include "widgets.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QGraphicsSceneWheelEvent struct {
	QGraphicsSceneEvent
}

type QGraphicsSceneWheelEvent_ITF interface {
	QGraphicsSceneEvent_ITF
	QGraphicsSceneWheelEvent_PTR() *QGraphicsSceneWheelEvent
}

func PointerFromQGraphicsSceneWheelEvent(ptr QGraphicsSceneWheelEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QGraphicsSceneWheelEvent_PTR().Pointer()
	}
	return nil
}

func NewQGraphicsSceneWheelEventFromPointer(ptr unsafe.Pointer) *QGraphicsSceneWheelEvent {
	var n = new(QGraphicsSceneWheelEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QGraphicsSceneWheelEvent) QGraphicsSceneWheelEvent_PTR() *QGraphicsSceneWheelEvent {
	return ptr
}

func (ptr *QGraphicsSceneWheelEvent) Buttons() core.Qt__MouseButton {
	defer qt.Recovering("QGraphicsSceneWheelEvent::buttons")

	if ptr.Pointer() != nil {
		return core.Qt__MouseButton(C.QGraphicsSceneWheelEvent_Buttons(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsSceneWheelEvent) Delta() int {
	defer qt.Recovering("QGraphicsSceneWheelEvent::delta")

	if ptr.Pointer() != nil {
		return int(C.QGraphicsSceneWheelEvent_Delta(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsSceneWheelEvent) Modifiers() core.Qt__KeyboardModifier {
	defer qt.Recovering("QGraphicsSceneWheelEvent::modifiers")

	if ptr.Pointer() != nil {
		return core.Qt__KeyboardModifier(C.QGraphicsSceneWheelEvent_Modifiers(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsSceneWheelEvent) Orientation() core.Qt__Orientation {
	defer qt.Recovering("QGraphicsSceneWheelEvent::orientation")

	if ptr.Pointer() != nil {
		return core.Qt__Orientation(C.QGraphicsSceneWheelEvent_Orientation(ptr.Pointer()))
	}
	return 0
}

func (ptr *QGraphicsSceneWheelEvent) ScreenPos() *core.QPoint {
	defer qt.Recovering("QGraphicsSceneWheelEvent::screenPos")

	if ptr.Pointer() != nil {
		return core.NewQPointFromPointer(C.QGraphicsSceneWheelEvent_ScreenPos(ptr.Pointer()))
	}
	return nil
}

func (ptr *QGraphicsSceneWheelEvent) DestroyQGraphicsSceneWheelEvent() {
	defer qt.Recovering("QGraphicsSceneWheelEvent::~QGraphicsSceneWheelEvent")

	if ptr.Pointer() != nil {
		C.QGraphicsSceneWheelEvent_DestroyQGraphicsSceneWheelEvent(ptr.Pointer())
	}
}
