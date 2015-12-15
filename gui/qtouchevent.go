package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QTouchEvent struct {
	QInputEvent
}

type QTouchEvent_ITF interface {
	QInputEvent_ITF
	QTouchEvent_PTR() *QTouchEvent
}

func PointerFromQTouchEvent(ptr QTouchEvent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTouchEvent_PTR().Pointer()
	}
	return nil
}

func NewQTouchEventFromPointer(ptr unsafe.Pointer) *QTouchEvent {
	var n = new(QTouchEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTouchEvent) QTouchEvent_PTR() *QTouchEvent {
	return ptr
}

func (ptr *QTouchEvent) Device() *QTouchDevice {
	defer qt.Recovering("QTouchEvent::device")

	if ptr.Pointer() != nil {
		return NewQTouchDeviceFromPointer(C.QTouchEvent_Device(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTouchEvent) Target() *core.QObject {
	defer qt.Recovering("QTouchEvent::target")

	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QTouchEvent_Target(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTouchEvent) TouchPointStates() core.Qt__TouchPointState {
	defer qt.Recovering("QTouchEvent::touchPointStates")

	if ptr.Pointer() != nil {
		return core.Qt__TouchPointState(C.QTouchEvent_TouchPointStates(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTouchEvent) Window() *QWindow {
	defer qt.Recovering("QTouchEvent::window")

	if ptr.Pointer() != nil {
		return NewQWindowFromPointer(C.QTouchEvent_Window(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTouchEvent) DestroyQTouchEvent() {
	defer qt.Recovering("QTouchEvent::~QTouchEvent")

	if ptr.Pointer() != nil {
		C.QTouchEvent_DestroyQTouchEvent(ptr.Pointer())
	}
}
