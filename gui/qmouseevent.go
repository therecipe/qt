package gui

//#include "qmouseevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QMouseEvent struct {
	QInputEvent
}

type QMouseEventITF interface {
	QInputEventITF
	QMouseEventPTR() *QMouseEvent
}

func PointerFromQMouseEvent(ptr QMouseEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMouseEventPTR().Pointer()
	}
	return nil
}

func QMouseEventFromPointer(ptr unsafe.Pointer) *QMouseEvent {
	var n = new(QMouseEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMouseEvent) QMouseEventPTR() *QMouseEvent {
	return ptr
}

func NewQMouseEvent(ty core.QEvent__Type, localPos core.QPointFITF, button core.Qt__MouseButton, buttons core.Qt__MouseButton, modifiers core.Qt__KeyboardModifier) *QMouseEvent {
	return QMouseEventFromPointer(unsafe.Pointer(C.QMouseEvent_NewQMouseEvent(C.int(ty), C.QtObjectPtr(core.PointerFromQPointF(localPos)), C.int(button), C.int(buttons), C.int(modifiers))))
}

func NewQMouseEvent2(ty core.QEvent__Type, localPos core.QPointFITF, screenPos core.QPointFITF, button core.Qt__MouseButton, buttons core.Qt__MouseButton, modifiers core.Qt__KeyboardModifier) *QMouseEvent {
	return QMouseEventFromPointer(unsafe.Pointer(C.QMouseEvent_NewQMouseEvent2(C.int(ty), C.QtObjectPtr(core.PointerFromQPointF(localPos)), C.QtObjectPtr(core.PointerFromQPointF(screenPos)), C.int(button), C.int(buttons), C.int(modifiers))))
}

func NewQMouseEvent3(ty core.QEvent__Type, localPos core.QPointFITF, windowPos core.QPointFITF, screenPos core.QPointFITF, button core.Qt__MouseButton, buttons core.Qt__MouseButton, modifiers core.Qt__KeyboardModifier) *QMouseEvent {
	return QMouseEventFromPointer(unsafe.Pointer(C.QMouseEvent_NewQMouseEvent3(C.int(ty), C.QtObjectPtr(core.PointerFromQPointF(localPos)), C.QtObjectPtr(core.PointerFromQPointF(windowPos)), C.QtObjectPtr(core.PointerFromQPointF(screenPos)), C.int(button), C.int(buttons), C.int(modifiers))))
}

func (ptr *QMouseEvent) Button() core.Qt__MouseButton {
	if ptr.Pointer() != nil {
		return core.Qt__MouseButton(C.QMouseEvent_Button(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMouseEvent) Buttons() core.Qt__MouseButton {
	if ptr.Pointer() != nil {
		return core.Qt__MouseButton(C.QMouseEvent_Buttons(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMouseEvent) Flags() core.Qt__MouseEventFlag {
	if ptr.Pointer() != nil {
		return core.Qt__MouseEventFlag(C.QMouseEvent_Flags(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMouseEvent) GlobalX() int {
	if ptr.Pointer() != nil {
		return int(C.QMouseEvent_GlobalX(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMouseEvent) GlobalY() int {
	if ptr.Pointer() != nil {
		return int(C.QMouseEvent_GlobalY(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMouseEvent) Source() core.Qt__MouseEventSource {
	if ptr.Pointer() != nil {
		return core.Qt__MouseEventSource(C.QMouseEvent_Source(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMouseEvent) X() int {
	if ptr.Pointer() != nil {
		return int(C.QMouseEvent_X(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QMouseEvent) Y() int {
	if ptr.Pointer() != nil {
		return int(C.QMouseEvent_Y(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}
