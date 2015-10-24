package multimedia

//#include "qvideodeviceselectorcontrol.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QVideoDeviceSelectorControl struct {
	QMediaControl
}

type QVideoDeviceSelectorControlITF interface {
	QMediaControlITF
	QVideoDeviceSelectorControlPTR() *QVideoDeviceSelectorControl
}

func PointerFromQVideoDeviceSelectorControl(ptr QVideoDeviceSelectorControlITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVideoDeviceSelectorControlPTR().Pointer()
	}
	return nil
}

func QVideoDeviceSelectorControlFromPointer(ptr unsafe.Pointer) *QVideoDeviceSelectorControl {
	var n = new(QVideoDeviceSelectorControl)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QVideoDeviceSelectorControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QVideoDeviceSelectorControl) QVideoDeviceSelectorControlPTR() *QVideoDeviceSelectorControl {
	return ptr
}

func (ptr *QVideoDeviceSelectorControl) DefaultDevice() int {
	if ptr.Pointer() != nil {
		return int(C.QVideoDeviceSelectorControl_DefaultDevice(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QVideoDeviceSelectorControl) DeviceCount() int {
	if ptr.Pointer() != nil {
		return int(C.QVideoDeviceSelectorControl_DeviceCount(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QVideoDeviceSelectorControl) DeviceDescription(index int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QVideoDeviceSelectorControl_DeviceDescription(C.QtObjectPtr(ptr.Pointer()), C.int(index)))
	}
	return ""
}

func (ptr *QVideoDeviceSelectorControl) DeviceName(index int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QVideoDeviceSelectorControl_DeviceName(C.QtObjectPtr(ptr.Pointer()), C.int(index)))
	}
	return ""
}

func (ptr *QVideoDeviceSelectorControl) ConnectDevicesChanged(f func()) {
	if ptr.Pointer() != nil {
		C.QVideoDeviceSelectorControl_ConnectDevicesChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "devicesChanged", f)
	}
}

func (ptr *QVideoDeviceSelectorControl) DisconnectDevicesChanged() {
	if ptr.Pointer() != nil {
		C.QVideoDeviceSelectorControl_DisconnectDevicesChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "devicesChanged")
	}
}

//export callbackQVideoDeviceSelectorControlDevicesChanged
func callbackQVideoDeviceSelectorControlDevicesChanged(ptrName *C.char) {
	qt.GetSignal(C.GoString(ptrName), "devicesChanged").(func())()
}

func (ptr *QVideoDeviceSelectorControl) SelectedDevice() int {
	if ptr.Pointer() != nil {
		return int(C.QVideoDeviceSelectorControl_SelectedDevice(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QVideoDeviceSelectorControl) ConnectSelectedDeviceChanged(f func(index int)) {
	if ptr.Pointer() != nil {
		C.QVideoDeviceSelectorControl_ConnectSelectedDeviceChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.ConnectSignal(ptr.ObjectName(), "selectedDeviceChanged", f)
	}
}

func (ptr *QVideoDeviceSelectorControl) DisconnectSelectedDeviceChanged() {
	if ptr.Pointer() != nil {
		C.QVideoDeviceSelectorControl_DisconnectSelectedDeviceChanged(C.QtObjectPtr(ptr.Pointer()))
		qt.DisconnectSignal(ptr.ObjectName(), "selectedDeviceChanged")
	}
}

//export callbackQVideoDeviceSelectorControlSelectedDeviceChanged
func callbackQVideoDeviceSelectorControlSelectedDeviceChanged(ptrName *C.char, index C.int) {
	qt.GetSignal(C.GoString(ptrName), "selectedDeviceChanged").(func(int))(int(index))
}

func (ptr *QVideoDeviceSelectorControl) SetSelectedDevice(index int) {
	if ptr.Pointer() != nil {
		C.QVideoDeviceSelectorControl_SetSelectedDevice(C.QtObjectPtr(ptr.Pointer()), C.int(index))
	}
}

func (ptr *QVideoDeviceSelectorControl) DestroyQVideoDeviceSelectorControl() {
	if ptr.Pointer() != nil {
		C.QVideoDeviceSelectorControl_DestroyQVideoDeviceSelectorControl(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
