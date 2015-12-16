package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QVideoDeviceSelectorControl struct {
	QMediaControl
}

type QVideoDeviceSelectorControl_ITF interface {
	QMediaControl_ITF
	QVideoDeviceSelectorControl_PTR() *QVideoDeviceSelectorControl
}

func PointerFromQVideoDeviceSelectorControl(ptr QVideoDeviceSelectorControl_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QVideoDeviceSelectorControl_PTR().Pointer()
	}
	return nil
}

func NewQVideoDeviceSelectorControlFromPointer(ptr unsafe.Pointer) *QVideoDeviceSelectorControl {
	var n = new(QVideoDeviceSelectorControl)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QVideoDeviceSelectorControl_") {
		n.SetObjectName("QVideoDeviceSelectorControl_" + qt.Identifier())
	}
	return n
}

func (ptr *QVideoDeviceSelectorControl) QVideoDeviceSelectorControl_PTR() *QVideoDeviceSelectorControl {
	return ptr
}

func (ptr *QVideoDeviceSelectorControl) DefaultDevice() int {
	defer qt.Recovering("QVideoDeviceSelectorControl::defaultDevice")

	if ptr.Pointer() != nil {
		return int(C.QVideoDeviceSelectorControl_DefaultDevice(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoDeviceSelectorControl) DeviceCount() int {
	defer qt.Recovering("QVideoDeviceSelectorControl::deviceCount")

	if ptr.Pointer() != nil {
		return int(C.QVideoDeviceSelectorControl_DeviceCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoDeviceSelectorControl) DeviceDescription(index int) string {
	defer qt.Recovering("QVideoDeviceSelectorControl::deviceDescription")

	if ptr.Pointer() != nil {
		return C.GoString(C.QVideoDeviceSelectorControl_DeviceDescription(ptr.Pointer(), C.int(index)))
	}
	return ""
}

func (ptr *QVideoDeviceSelectorControl) DeviceName(index int) string {
	defer qt.Recovering("QVideoDeviceSelectorControl::deviceName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QVideoDeviceSelectorControl_DeviceName(ptr.Pointer(), C.int(index)))
	}
	return ""
}

func (ptr *QVideoDeviceSelectorControl) ConnectDevicesChanged(f func()) {
	defer qt.Recovering("connect QVideoDeviceSelectorControl::devicesChanged")

	if ptr.Pointer() != nil {
		C.QVideoDeviceSelectorControl_ConnectDevicesChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "devicesChanged", f)
	}
}

func (ptr *QVideoDeviceSelectorControl) DisconnectDevicesChanged() {
	defer qt.Recovering("disconnect QVideoDeviceSelectorControl::devicesChanged")

	if ptr.Pointer() != nil {
		C.QVideoDeviceSelectorControl_DisconnectDevicesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "devicesChanged")
	}
}

//export callbackQVideoDeviceSelectorControlDevicesChanged
func callbackQVideoDeviceSelectorControlDevicesChanged(ptrName *C.char) {
	defer qt.Recovering("callback QVideoDeviceSelectorControl::devicesChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "devicesChanged")
	if signal != nil {
		signal.(func())()
	}

}

func (ptr *QVideoDeviceSelectorControl) SelectedDevice() int {
	defer qt.Recovering("QVideoDeviceSelectorControl::selectedDevice")

	if ptr.Pointer() != nil {
		return int(C.QVideoDeviceSelectorControl_SelectedDevice(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoDeviceSelectorControl) ConnectSelectedDeviceChanged2(f func(name string)) {
	defer qt.Recovering("connect QVideoDeviceSelectorControl::selectedDeviceChanged")

	if ptr.Pointer() != nil {
		C.QVideoDeviceSelectorControl_ConnectSelectedDeviceChanged2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "selectedDeviceChanged2", f)
	}
}

func (ptr *QVideoDeviceSelectorControl) DisconnectSelectedDeviceChanged2() {
	defer qt.Recovering("disconnect QVideoDeviceSelectorControl::selectedDeviceChanged")

	if ptr.Pointer() != nil {
		C.QVideoDeviceSelectorControl_DisconnectSelectedDeviceChanged2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "selectedDeviceChanged2")
	}
}

//export callbackQVideoDeviceSelectorControlSelectedDeviceChanged2
func callbackQVideoDeviceSelectorControlSelectedDeviceChanged2(ptrName *C.char, name *C.char) {
	defer qt.Recovering("callback QVideoDeviceSelectorControl::selectedDeviceChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "selectedDeviceChanged2")
	if signal != nil {
		signal.(func(string))(C.GoString(name))
	}

}

func (ptr *QVideoDeviceSelectorControl) ConnectSelectedDeviceChanged(f func(index int)) {
	defer qt.Recovering("connect QVideoDeviceSelectorControl::selectedDeviceChanged")

	if ptr.Pointer() != nil {
		C.QVideoDeviceSelectorControl_ConnectSelectedDeviceChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "selectedDeviceChanged", f)
	}
}

func (ptr *QVideoDeviceSelectorControl) DisconnectSelectedDeviceChanged() {
	defer qt.Recovering("disconnect QVideoDeviceSelectorControl::selectedDeviceChanged")

	if ptr.Pointer() != nil {
		C.QVideoDeviceSelectorControl_DisconnectSelectedDeviceChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "selectedDeviceChanged")
	}
}

//export callbackQVideoDeviceSelectorControlSelectedDeviceChanged
func callbackQVideoDeviceSelectorControlSelectedDeviceChanged(ptrName *C.char, index C.int) {
	defer qt.Recovering("callback QVideoDeviceSelectorControl::selectedDeviceChanged")

	var signal = qt.GetSignal(C.GoString(ptrName), "selectedDeviceChanged")
	if signal != nil {
		signal.(func(int))(int(index))
	}

}

func (ptr *QVideoDeviceSelectorControl) SetSelectedDevice(index int) {
	defer qt.Recovering("QVideoDeviceSelectorControl::setSelectedDevice")

	if ptr.Pointer() != nil {
		C.QVideoDeviceSelectorControl_SetSelectedDevice(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QVideoDeviceSelectorControl) DestroyQVideoDeviceSelectorControl() {
	defer qt.Recovering("QVideoDeviceSelectorControl::~QVideoDeviceSelectorControl")

	if ptr.Pointer() != nil {
		C.QVideoDeviceSelectorControl_DestroyQVideoDeviceSelectorControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
