package gui

//#include "qtabletevent.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QTabletEvent struct {
	QInputEvent
}

type QTabletEventITF interface {
	QInputEventITF
	QTabletEventPTR() *QTabletEvent
}

func PointerFromQTabletEvent(ptr QTabletEventITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTabletEventPTR().Pointer()
	}
	return nil
}

func QTabletEventFromPointer(ptr unsafe.Pointer) *QTabletEvent {
	var n = new(QTabletEvent)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTabletEvent) QTabletEventPTR() *QTabletEvent {
	return ptr
}

//QTabletEvent::PointerType
type QTabletEvent__PointerType int

var (
	QTabletEvent__UnknownPointer = QTabletEvent__PointerType(0)
	QTabletEvent__Pen            = QTabletEvent__PointerType(1)
	QTabletEvent__Cursor         = QTabletEvent__PointerType(2)
	QTabletEvent__Eraser         = QTabletEvent__PointerType(3)
)

//QTabletEvent::TabletDevice
type QTabletEvent__TabletDevice int

var (
	QTabletEvent__NoDevice       = QTabletEvent__TabletDevice(0)
	QTabletEvent__Puck           = QTabletEvent__TabletDevice(1)
	QTabletEvent__Stylus         = QTabletEvent__TabletDevice(2)
	QTabletEvent__Airbrush       = QTabletEvent__TabletDevice(3)
	QTabletEvent__FourDMouse     = QTabletEvent__TabletDevice(4)
	QTabletEvent__XFreeEraser    = QTabletEvent__TabletDevice(5)
	QTabletEvent__RotationStylus = QTabletEvent__TabletDevice(6)
)

func (ptr *QTabletEvent) Button() core.Qt__MouseButton {
	if ptr.Pointer() != nil {
		return core.Qt__MouseButton(C.QTabletEvent_Button(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTabletEvent) Buttons() core.Qt__MouseButton {
	if ptr.Pointer() != nil {
		return core.Qt__MouseButton(C.QTabletEvent_Buttons(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTabletEvent) Device() QTabletEvent__TabletDevice {
	if ptr.Pointer() != nil {
		return QTabletEvent__TabletDevice(C.QTabletEvent_Device(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTabletEvent) GlobalX() int {
	if ptr.Pointer() != nil {
		return int(C.QTabletEvent_GlobalX(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTabletEvent) GlobalY() int {
	if ptr.Pointer() != nil {
		return int(C.QTabletEvent_GlobalY(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTabletEvent) PointerType() QTabletEvent__PointerType {
	if ptr.Pointer() != nil {
		return QTabletEvent__PointerType(C.QTabletEvent_PointerType(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTabletEvent) X() int {
	if ptr.Pointer() != nil {
		return int(C.QTabletEvent_X(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTabletEvent) XTilt() int {
	if ptr.Pointer() != nil {
		return int(C.QTabletEvent_XTilt(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTabletEvent) Y() int {
	if ptr.Pointer() != nil {
		return int(C.QTabletEvent_Y(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTabletEvent) YTilt() int {
	if ptr.Pointer() != nil {
		return int(C.QTabletEvent_YTilt(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTabletEvent) Z() int {
	if ptr.Pointer() != nil {
		return int(C.QTabletEvent_Z(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}
