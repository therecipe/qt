package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QWheelEvent struct {
	QInputEvent
}

type QWheelEvent_ITF interface {
	QInputEvent_ITF
	QWheelEvent_PTR() *QWheelEvent
}

func PointerFromQWheelEvent(ptr QWheelEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWheelEvent_PTR().Pointer()
	}
	return nil
}

func NewQWheelEventFromPointer(ptr unsafe.Pointer) *QWheelEvent {
	var n = new(QWheelEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QWheelEvent) QWheelEvent_PTR() *QWheelEvent {
	return ptr
}

func NewQWheelEvent(pos core.QPointF_ITF, globalPos core.QPointF_ITF, pixelDelta core.QPoint_ITF, angleDelta core.QPoint_ITF, qt4Delta int, qt4Orientation core.Qt__Orientation, buttons core.Qt__MouseButton, modifiers core.Qt__KeyboardModifier) *QWheelEvent {
	defer qt.Recovering("QWheelEvent::QWheelEvent")

	return NewQWheelEventFromPointer(C.QWheelEvent_NewQWheelEvent(core.PointerFromQPointF(pos), core.PointerFromQPointF(globalPos), core.PointerFromQPoint(pixelDelta), core.PointerFromQPoint(angleDelta), C.int(qt4Delta), C.int(qt4Orientation), C.int(buttons), C.int(modifiers)))
}

func NewQWheelEvent4(pos core.QPointF_ITF, globalPos core.QPointF_ITF, pixelDelta core.QPoint_ITF, angleDelta core.QPoint_ITF, qt4Delta int, qt4Orientation core.Qt__Orientation, buttons core.Qt__MouseButton, modifiers core.Qt__KeyboardModifier, phase core.Qt__ScrollPhase) *QWheelEvent {
	defer qt.Recovering("QWheelEvent::QWheelEvent")

	return NewQWheelEventFromPointer(C.QWheelEvent_NewQWheelEvent4(core.PointerFromQPointF(pos), core.PointerFromQPointF(globalPos), core.PointerFromQPoint(pixelDelta), core.PointerFromQPoint(angleDelta), C.int(qt4Delta), C.int(qt4Orientation), C.int(buttons), C.int(modifiers), C.int(phase)))
}

func NewQWheelEvent5(pos core.QPointF_ITF, globalPos core.QPointF_ITF, pixelDelta core.QPoint_ITF, angleDelta core.QPoint_ITF, qt4Delta int, qt4Orientation core.Qt__Orientation, buttons core.Qt__MouseButton, modifiers core.Qt__KeyboardModifier, phase core.Qt__ScrollPhase, source core.Qt__MouseEventSource) *QWheelEvent {
	defer qt.Recovering("QWheelEvent::QWheelEvent")

	return NewQWheelEventFromPointer(C.QWheelEvent_NewQWheelEvent5(core.PointerFromQPointF(pos), core.PointerFromQPointF(globalPos), core.PointerFromQPoint(pixelDelta), core.PointerFromQPoint(angleDelta), C.int(qt4Delta), C.int(qt4Orientation), C.int(buttons), C.int(modifiers), C.int(phase), C.int(source)))
}

func (ptr *QWheelEvent) AngleDelta() *core.QPoint {
	defer qt.Recovering("QWheelEvent::angleDelta")

	if ptr.Pointer() != nil {
		return core.NewQPointFromPointer(C.QWheelEvent_AngleDelta(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWheelEvent) Buttons() core.Qt__MouseButton {
	defer qt.Recovering("QWheelEvent::buttons")

	if ptr.Pointer() != nil {
		return core.Qt__MouseButton(C.QWheelEvent_Buttons(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWheelEvent) GlobalPos() *core.QPoint {
	defer qt.Recovering("QWheelEvent::globalPos")

	if ptr.Pointer() != nil {
		return core.NewQPointFromPointer(C.QWheelEvent_GlobalPos(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWheelEvent) GlobalX() int {
	defer qt.Recovering("QWheelEvent::globalX")

	if ptr.Pointer() != nil {
		return int(C.QWheelEvent_GlobalX(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWheelEvent) GlobalY() int {
	defer qt.Recovering("QWheelEvent::globalY")

	if ptr.Pointer() != nil {
		return int(C.QWheelEvent_GlobalY(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWheelEvent) Phase() core.Qt__ScrollPhase {
	defer qt.Recovering("QWheelEvent::phase")

	if ptr.Pointer() != nil {
		return core.Qt__ScrollPhase(C.QWheelEvent_Phase(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWheelEvent) PixelDelta() *core.QPoint {
	defer qt.Recovering("QWheelEvent::pixelDelta")

	if ptr.Pointer() != nil {
		return core.NewQPointFromPointer(C.QWheelEvent_PixelDelta(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWheelEvent) Pos() *core.QPoint {
	defer qt.Recovering("QWheelEvent::pos")

	if ptr.Pointer() != nil {
		return core.NewQPointFromPointer(C.QWheelEvent_Pos(ptr.Pointer()))
	}
	return nil
}

func (ptr *QWheelEvent) Source() core.Qt__MouseEventSource {
	defer qt.Recovering("QWheelEvent::source")

	if ptr.Pointer() != nil {
		return core.Qt__MouseEventSource(C.QWheelEvent_Source(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWheelEvent) X() int {
	defer qt.Recovering("QWheelEvent::x")

	if ptr.Pointer() != nil {
		return int(C.QWheelEvent_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWheelEvent) Y() int {
	defer qt.Recovering("QWheelEvent::y")

	if ptr.Pointer() != nil {
		return int(C.QWheelEvent_Y(ptr.Pointer()))
	}
	return 0
}
