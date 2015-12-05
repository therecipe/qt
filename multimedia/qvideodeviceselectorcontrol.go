package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"log"
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
		n.SetObjectName("QVideoDeviceSelectorControl_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QVideoDeviceSelectorControl) QVideoDeviceSelectorControl_PTR() *QVideoDeviceSelectorControl {
	return ptr
}

func (ptr *QVideoDeviceSelectorControl) DefaultDevice() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoDeviceSelectorControl::defaultDevice")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QVideoDeviceSelectorControl_DefaultDevice(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoDeviceSelectorControl) DeviceCount() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoDeviceSelectorControl::deviceCount")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QVideoDeviceSelectorControl_DeviceCount(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoDeviceSelectorControl) DeviceDescription(index int) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoDeviceSelectorControl::deviceDescription")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QVideoDeviceSelectorControl_DeviceDescription(ptr.Pointer(), C.int(index)))
	}
	return ""
}

func (ptr *QVideoDeviceSelectorControl) DeviceName(index int) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoDeviceSelectorControl::deviceName")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QVideoDeviceSelectorControl_DeviceName(ptr.Pointer(), C.int(index)))
	}
	return ""
}

func (ptr *QVideoDeviceSelectorControl) ConnectDevicesChanged(f func()) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoDeviceSelectorControl::devicesChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoDeviceSelectorControl_ConnectDevicesChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "devicesChanged", f)
	}
}

func (ptr *QVideoDeviceSelectorControl) DisconnectDevicesChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoDeviceSelectorControl::devicesChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoDeviceSelectorControl_DisconnectDevicesChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "devicesChanged")
	}
}

//export callbackQVideoDeviceSelectorControlDevicesChanged
func callbackQVideoDeviceSelectorControlDevicesChanged(ptrName *C.char) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoDeviceSelectorControl::devicesChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "devicesChanged").(func())()
}

func (ptr *QVideoDeviceSelectorControl) SelectedDevice() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoDeviceSelectorControl::selectedDevice")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QVideoDeviceSelectorControl_SelectedDevice(ptr.Pointer()))
	}
	return 0
}

func (ptr *QVideoDeviceSelectorControl) ConnectSelectedDeviceChanged(f func(index int)) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoDeviceSelectorControl::selectedDeviceChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoDeviceSelectorControl_ConnectSelectedDeviceChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "selectedDeviceChanged", f)
	}
}

func (ptr *QVideoDeviceSelectorControl) DisconnectSelectedDeviceChanged() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoDeviceSelectorControl::selectedDeviceChanged")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoDeviceSelectorControl_DisconnectSelectedDeviceChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "selectedDeviceChanged")
	}
}

//export callbackQVideoDeviceSelectorControlSelectedDeviceChanged
func callbackQVideoDeviceSelectorControlSelectedDeviceChanged(ptrName *C.char, index C.int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoDeviceSelectorControl::selectedDeviceChanged")
		}
	}()

	qt.GetSignal(C.GoString(ptrName), "selectedDeviceChanged").(func(int))(int(index))
}

func (ptr *QVideoDeviceSelectorControl) SetSelectedDevice(index int) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoDeviceSelectorControl::setSelectedDevice")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoDeviceSelectorControl_SetSelectedDevice(ptr.Pointer(), C.int(index))
	}
}

func (ptr *QVideoDeviceSelectorControl) DestroyQVideoDeviceSelectorControl() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QVideoDeviceSelectorControl::~QVideoDeviceSelectorControl")
		}
	}()

	if ptr.Pointer() != nil {
		C.QVideoDeviceSelectorControl_DestroyQVideoDeviceSelectorControl(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
