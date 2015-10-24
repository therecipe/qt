package gui

//#include "qtouchdevice.h"
import "C"
import (
	"unsafe"
)

type QTouchDevice struct {
	ptr unsafe.Pointer
}

type QTouchDeviceITF interface {
	QTouchDevicePTR() *QTouchDevice
}

func (p *QTouchDevice) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QTouchDevice) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQTouchDevice(ptr QTouchDeviceITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QTouchDevicePTR().Pointer()
	}
	return nil
}

func QTouchDeviceFromPointer(ptr unsafe.Pointer) *QTouchDevice {
	var n = new(QTouchDevice)
	n.SetPointer(ptr)
	return n
}

func (ptr *QTouchDevice) QTouchDevicePTR() *QTouchDevice {
	return ptr
}

//QTouchDevice::CapabilityFlag
type QTouchDevice__CapabilityFlag int

var (
	QTouchDevice__Position           = QTouchDevice__CapabilityFlag(0x0001)
	QTouchDevice__Area               = QTouchDevice__CapabilityFlag(0x0002)
	QTouchDevice__Pressure           = QTouchDevice__CapabilityFlag(0x0004)
	QTouchDevice__Velocity           = QTouchDevice__CapabilityFlag(0x0008)
	QTouchDevice__RawPositions       = QTouchDevice__CapabilityFlag(0x0010)
	QTouchDevice__NormalizedPosition = QTouchDevice__CapabilityFlag(0x0020)
	QTouchDevice__MouseEmulation     = QTouchDevice__CapabilityFlag(0x0040)
)

//QTouchDevice::DeviceType
type QTouchDevice__DeviceType int

var (
	QTouchDevice__TouchScreen = QTouchDevice__DeviceType(0)
	QTouchDevice__TouchPad    = QTouchDevice__DeviceType(1)
)

func NewQTouchDevice() *QTouchDevice {
	return QTouchDeviceFromPointer(unsafe.Pointer(C.QTouchDevice_NewQTouchDevice()))
}

func (ptr *QTouchDevice) Capabilities() QTouchDevice__CapabilityFlag {
	if ptr.Pointer() != nil {
		return QTouchDevice__CapabilityFlag(C.QTouchDevice_Capabilities(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTouchDevice) MaximumTouchPoints() int {
	if ptr.Pointer() != nil {
		return int(C.QTouchDevice_MaximumTouchPoints(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTouchDevice) Name() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QTouchDevice_Name(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QTouchDevice) SetCapabilities(caps QTouchDevice__CapabilityFlag) {
	if ptr.Pointer() != nil {
		C.QTouchDevice_SetCapabilities(C.QtObjectPtr(ptr.Pointer()), C.int(caps))
	}
}

func (ptr *QTouchDevice) SetMaximumTouchPoints(max int) {
	if ptr.Pointer() != nil {
		C.QTouchDevice_SetMaximumTouchPoints(C.QtObjectPtr(ptr.Pointer()), C.int(max))
	}
}

func (ptr *QTouchDevice) SetName(name string) {
	if ptr.Pointer() != nil {
		C.QTouchDevice_SetName(C.QtObjectPtr(ptr.Pointer()), C.CString(name))
	}
}

func (ptr *QTouchDevice) SetType(devType QTouchDevice__DeviceType) {
	if ptr.Pointer() != nil {
		C.QTouchDevice_SetType(C.QtObjectPtr(ptr.Pointer()), C.int(devType))
	}
}

func (ptr *QTouchDevice) Type() QTouchDevice__DeviceType {
	if ptr.Pointer() != nil {
		return QTouchDevice__DeviceType(C.QTouchDevice_Type(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QTouchDevice) DestroyQTouchDevice() {
	if ptr.Pointer() != nil {
		C.QTouchDevice_DestroyQTouchDevice(C.QtObjectPtr(ptr.Pointer()))
	}
}
