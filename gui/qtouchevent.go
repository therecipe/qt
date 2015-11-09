package gui

//#include "qtouchevent.h"
import "C"
import (
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
	if ptr.Pointer() != nil {
		return NewQTouchDeviceFromPointer(C.QTouchEvent_Device(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTouchEvent) Target() *core.QObject {
	if ptr.Pointer() != nil {
		return core.NewQObjectFromPointer(C.QTouchEvent_Target(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTouchEvent) TouchPointStates() core.Qt__TouchPointState {
	if ptr.Pointer() != nil {
		return core.Qt__TouchPointState(C.QTouchEvent_TouchPointStates(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTouchEvent) Window() *QWindow {
	if ptr.Pointer() != nil {
		return NewQWindowFromPointer(C.QTouchEvent_Window(ptr.Pointer()))
	}
	return nil
}

func (ptr *QTouchEvent) DestroyQTouchEvent() {
	if ptr.Pointer() != nil {
		C.QTouchEvent_DestroyQTouchEvent(ptr.Pointer())
	}
}
