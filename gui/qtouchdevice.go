package gui

//#include "gui.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QTouchDevice struct {
	ptr unsafe.Pointer
}

type QTouchDevice_ITF interface {
	QTouchDevice_PTR() *QTouchDevice
}

func (p *QTouchDevice) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTouchDevice) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTouchDevice(ptr QTouchDevice_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTouchDevice_PTR().Pointer()
	}
	return nil
}

func NewQTouchDeviceFromPointer(ptr unsafe.Pointer) *QTouchDevice {
	var n = new(QTouchDevice)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTouchDevice) QTouchDevice_PTR() *QTouchDevice {
	return ptr
}

//QTouchDevice::CapabilityFlag
type QTouchDevice__CapabilityFlag int64

const (
	QTouchDevice__Position           = QTouchDevice__CapabilityFlag(0x0001)
	QTouchDevice__Area               = QTouchDevice__CapabilityFlag(0x0002)
	QTouchDevice__Pressure           = QTouchDevice__CapabilityFlag(0x0004)
	QTouchDevice__Velocity           = QTouchDevice__CapabilityFlag(0x0008)
	QTouchDevice__RawPositions       = QTouchDevice__CapabilityFlag(0x0010)
	QTouchDevice__NormalizedPosition = QTouchDevice__CapabilityFlag(0x0020)
	QTouchDevice__MouseEmulation     = QTouchDevice__CapabilityFlag(0x0040)
)

//QTouchDevice::DeviceType
type QTouchDevice__DeviceType int64

const (
	QTouchDevice__TouchScreen = QTouchDevice__DeviceType(0)
	QTouchDevice__TouchPad    = QTouchDevice__DeviceType(1)
)

func NewQTouchDevice() *QTouchDevice {
	defer qt.Recovering("QTouchDevice::QTouchDevice")

	return NewQTouchDeviceFromPointer(C.QTouchDevice_NewQTouchDevice())
}

func (ptr *QTouchDevice) Capabilities() QTouchDevice__CapabilityFlag {
	defer qt.Recovering("QTouchDevice::capabilities")

	if ptr.Pointer() != nil {
		return QTouchDevice__CapabilityFlag(C.QTouchDevice_Capabilities(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTouchDevice) MaximumTouchPoints() int {
	defer qt.Recovering("QTouchDevice::maximumTouchPoints")

	if ptr.Pointer() != nil {
		return int(C.QTouchDevice_MaximumTouchPoints(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTouchDevice) Name() string {
	defer qt.Recovering("QTouchDevice::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QTouchDevice_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QTouchDevice) SetCapabilities(caps QTouchDevice__CapabilityFlag) {
	defer qt.Recovering("QTouchDevice::setCapabilities")

	if ptr.Pointer() != nil {
		C.QTouchDevice_SetCapabilities(ptr.Pointer(), C.int(caps))
	}
}

func (ptr *QTouchDevice) SetMaximumTouchPoints(max int) {
	defer qt.Recovering("QTouchDevice::setMaximumTouchPoints")

	if ptr.Pointer() != nil {
		C.QTouchDevice_SetMaximumTouchPoints(ptr.Pointer(), C.int(max))
	}
}

func (ptr *QTouchDevice) SetName(name string) {
	defer qt.Recovering("QTouchDevice::setName")

	if ptr.Pointer() != nil {
		C.QTouchDevice_SetName(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QTouchDevice) SetType(devType QTouchDevice__DeviceType) {
	defer qt.Recovering("QTouchDevice::setType")

	if ptr.Pointer() != nil {
		C.QTouchDevice_SetType(ptr.Pointer(), C.int(devType))
	}
}

func (ptr *QTouchDevice) Type() QTouchDevice__DeviceType {
	defer qt.Recovering("QTouchDevice::type")

	if ptr.Pointer() != nil {
		return QTouchDevice__DeviceType(C.QTouchDevice_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QTouchDevice) DestroyQTouchDevice() {
	defer qt.Recovering("QTouchDevice::~QTouchDevice")

	if ptr.Pointer() != nil {
		C.QTouchDevice_DestroyQTouchDevice(ptr.Pointer())
	}
}
