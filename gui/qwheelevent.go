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

type QWheelEventITF interface {
	QInputEventITF
	QWheelEventPTR() *QWheelEvent
}

func PointerFromQWheelEvent(ptr QWheelEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QWheelEventPTR().Pointer()
	}
	return nil
}

func QWheelEventFromPointer(ptr unsafe.Pointer) *QWheelEvent {
	var n = new(QWheelEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QWheelEvent) QWheelEventPTR() *QWheelEvent {
	return ptr
}

func NewQWheelEvent(pos core.QPointFITF, globalPos core.QPointFITF, pixelDelta core.QPointITF, angleDelta core.QPointITF, qt4Delta int, qt4Orientation core.Qt__Orientation, buttons core.Qt__MouseButton, modifiers core.Qt__KeyboardModifier) *QWheelEvent {
	return QWheelEventFromPointer(unsafe.Pointer(C.QWheelEvent_NewQWheelEvent(C.QtObjectPtr(core.PointerFromQPointF(pos)), C.QtObjectPtr(core.PointerFromQPointF(globalPos)), C.QtObjectPtr(core.PointerFromQPoint(pixelDelta)), C.QtObjectPtr(core.PointerFromQPoint(angleDelta)), C.int(qt4Delta), C.int(qt4Orientation), C.int(buttons), C.int(modifiers))))
}

func NewQWheelEvent4(pos core.QPointFITF, globalPos core.QPointFITF, pixelDelta core.QPointITF, angleDelta core.QPointITF, qt4Delta int, qt4Orientation core.Qt__Orientation, buttons core.Qt__MouseButton, modifiers core.Qt__KeyboardModifier, phase core.Qt__ScrollPhase) *QWheelEvent {
	return QWheelEventFromPointer(unsafe.Pointer(C.QWheelEvent_NewQWheelEvent4(C.QtObjectPtr(core.PointerFromQPointF(pos)), C.QtObjectPtr(core.PointerFromQPointF(globalPos)), C.QtObjectPtr(core.PointerFromQPoint(pixelDelta)), C.QtObjectPtr(core.PointerFromQPoint(angleDelta)), C.int(qt4Delta), C.int(qt4Orientation), C.int(buttons), C.int(modifiers), C.int(phase))))
}

func NewQWheelEvent5(pos core.QPointFITF, globalPos core.QPointFITF, pixelDelta core.QPointITF, angleDelta core.QPointITF, qt4Delta int, qt4Orientation core.Qt__Orientation, buttons core.Qt__MouseButton, modifiers core.Qt__KeyboardModifier, phase core.Qt__ScrollPhase, source core.Qt__MouseEventSource) *QWheelEvent {
	return QWheelEventFromPointer(unsafe.Pointer(C.QWheelEvent_NewQWheelEvent5(C.QtObjectPtr(core.PointerFromQPointF(pos)), C.QtObjectPtr(core.PointerFromQPointF(globalPos)), C.QtObjectPtr(core.PointerFromQPoint(pixelDelta)), C.QtObjectPtr(core.PointerFromQPoint(angleDelta)), C.int(qt4Delta), C.int(qt4Orientation), C.int(buttons), C.int(modifiers), C.int(phase), C.int(source))))
}

func (ptr *QWheelEvent) Buttons() core.Qt__MouseButton {
	if ptr.Pointer() != nil {
		return core.Qt__MouseButton(C.QWheelEvent_Buttons(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWheelEvent) GlobalX() int {
	if ptr.Pointer() != nil {
		return int(C.QWheelEvent_GlobalX(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWheelEvent) GlobalY() int {
	if ptr.Pointer() != nil {
		return int(C.QWheelEvent_GlobalY(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWheelEvent) Phase() core.Qt__ScrollPhase {
	if ptr.Pointer() != nil {
		return core.Qt__ScrollPhase(C.QWheelEvent_Phase(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWheelEvent) Source() core.Qt__MouseEventSource {
	if ptr.Pointer() != nil {
		return core.Qt__MouseEventSource(C.QWheelEvent_Source(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWheelEvent) X() int {
	if ptr.Pointer() != nil {
		return int(C.QWheelEvent_X(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QWheelEvent) Y() int {
	if ptr.Pointer() != nil {
		return int(C.QWheelEvent_Y(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}
