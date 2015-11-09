package gui

//#include "qwheelevent.h"
import "C"
import (
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
	return NewQWheelEventFromPointer(C.QWheelEvent_NewQWheelEvent(core.PointerFromQPointF(pos), core.PointerFromQPointF(globalPos), core.PointerFromQPoint(pixelDelta), core.PointerFromQPoint(angleDelta), C.int(qt4Delta), C.int(qt4Orientation), C.int(buttons), C.int(modifiers)))
}

func NewQWheelEvent4(pos core.QPointF_ITF, globalPos core.QPointF_ITF, pixelDelta core.QPoint_ITF, angleDelta core.QPoint_ITF, qt4Delta int, qt4Orientation core.Qt__Orientation, buttons core.Qt__MouseButton, modifiers core.Qt__KeyboardModifier, phase core.Qt__ScrollPhase) *QWheelEvent {
	return NewQWheelEventFromPointer(C.QWheelEvent_NewQWheelEvent4(core.PointerFromQPointF(pos), core.PointerFromQPointF(globalPos), core.PointerFromQPoint(pixelDelta), core.PointerFromQPoint(angleDelta), C.int(qt4Delta), C.int(qt4Orientation), C.int(buttons), C.int(modifiers), C.int(phase)))
}

func NewQWheelEvent5(pos core.QPointF_ITF, globalPos core.QPointF_ITF, pixelDelta core.QPoint_ITF, angleDelta core.QPoint_ITF, qt4Delta int, qt4Orientation core.Qt__Orientation, buttons core.Qt__MouseButton, modifiers core.Qt__KeyboardModifier, phase core.Qt__ScrollPhase, source core.Qt__MouseEventSource) *QWheelEvent {
	return NewQWheelEventFromPointer(C.QWheelEvent_NewQWheelEvent5(core.PointerFromQPointF(pos), core.PointerFromQPointF(globalPos), core.PointerFromQPoint(pixelDelta), core.PointerFromQPoint(angleDelta), C.int(qt4Delta), C.int(qt4Orientation), C.int(buttons), C.int(modifiers), C.int(phase), C.int(source)))
}

func (ptr *QWheelEvent) Buttons() core.Qt__MouseButton {
	if ptr.Pointer() != nil {
		return core.Qt__MouseButton(C.QWheelEvent_Buttons(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWheelEvent) GlobalX() int {
	if ptr.Pointer() != nil {
		return int(C.QWheelEvent_GlobalX(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWheelEvent) GlobalY() int {
	if ptr.Pointer() != nil {
		return int(C.QWheelEvent_GlobalY(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWheelEvent) Phase() core.Qt__ScrollPhase {
	if ptr.Pointer() != nil {
		return core.Qt__ScrollPhase(C.QWheelEvent_Phase(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWheelEvent) Source() core.Qt__MouseEventSource {
	if ptr.Pointer() != nil {
		return core.Qt__MouseEventSource(C.QWheelEvent_Source(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWheelEvent) X() int {
	if ptr.Pointer() != nil {
		return int(C.QWheelEvent_X(ptr.Pointer()))
	}
	return 0
}

func (ptr *QWheelEvent) Y() int {
	if ptr.Pointer() != nil {
		return int(C.QWheelEvent_Y(ptr.Pointer()))
	}
	return 0
}
