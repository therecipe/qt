package multimedia

//#include "multimedia.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
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
func callbackQVideoDeviceSelectorControlDevicesChanged(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QVideoDeviceSelectorControl::devicesChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "devicesChanged"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QVideoDeviceSelectorControl) DevicesChanged() {
	defer qt.Recovering("QVideoDeviceSelectorControl::devicesChanged")

	if ptr.Pointer() != nil {
		C.QVideoDeviceSelectorControl_DevicesChanged(ptr.Pointer())
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
func callbackQVideoDeviceSelectorControlSelectedDeviceChanged2(ptr unsafe.Pointer, ptrName *C.char, name *C.char) {
	defer qt.Recovering("callback QVideoDeviceSelectorControl::selectedDeviceChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "selectedDeviceChanged2"); signal != nil {
		signal.(func(string))(C.GoString(name))
	}

}

func (ptr *QVideoDeviceSelectorControl) SelectedDeviceChanged2(name string) {
	defer qt.Recovering("QVideoDeviceSelectorControl::selectedDeviceChanged")

	if ptr.Pointer() != nil {
		C.QVideoDeviceSelectorControl_SelectedDeviceChanged2(ptr.Pointer(), C.CString(name))
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
func callbackQVideoDeviceSelectorControlSelectedDeviceChanged(ptr unsafe.Pointer, ptrName *C.char, index C.int) {
	defer qt.Recovering("callback QVideoDeviceSelectorControl::selectedDeviceChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "selectedDeviceChanged"); signal != nil {
		signal.(func(int))(int(index))
	}

}

func (ptr *QVideoDeviceSelectorControl) SelectedDeviceChanged(index int) {
	defer qt.Recovering("QVideoDeviceSelectorControl::selectedDeviceChanged")

	if ptr.Pointer() != nil {
		C.QVideoDeviceSelectorControl_SelectedDeviceChanged(ptr.Pointer(), C.int(index))
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

func (ptr *QVideoDeviceSelectorControl) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QVideoDeviceSelectorControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QVideoDeviceSelectorControl) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QVideoDeviceSelectorControl::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQVideoDeviceSelectorControlTimerEvent
func callbackQVideoDeviceSelectorControlTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoDeviceSelectorControl::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQVideoDeviceSelectorControlFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QVideoDeviceSelectorControl) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QVideoDeviceSelectorControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QVideoDeviceSelectorControl_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QVideoDeviceSelectorControl) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QVideoDeviceSelectorControl::timerEvent")

	if ptr.Pointer() != nil {
		C.QVideoDeviceSelectorControl_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QVideoDeviceSelectorControl) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QVideoDeviceSelectorControl::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QVideoDeviceSelectorControl) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QVideoDeviceSelectorControl::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQVideoDeviceSelectorControlChildEvent
func callbackQVideoDeviceSelectorControlChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoDeviceSelectorControl::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQVideoDeviceSelectorControlFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QVideoDeviceSelectorControl) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QVideoDeviceSelectorControl::childEvent")

	if ptr.Pointer() != nil {
		C.QVideoDeviceSelectorControl_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QVideoDeviceSelectorControl) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QVideoDeviceSelectorControl::childEvent")

	if ptr.Pointer() != nil {
		C.QVideoDeviceSelectorControl_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QVideoDeviceSelectorControl) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QVideoDeviceSelectorControl::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QVideoDeviceSelectorControl) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QVideoDeviceSelectorControl::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQVideoDeviceSelectorControlCustomEvent
func callbackQVideoDeviceSelectorControlCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QVideoDeviceSelectorControl::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQVideoDeviceSelectorControlFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QVideoDeviceSelectorControl) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QVideoDeviceSelectorControl::customEvent")

	if ptr.Pointer() != nil {
		C.QVideoDeviceSelectorControl_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QVideoDeviceSelectorControl) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QVideoDeviceSelectorControl::customEvent")

	if ptr.Pointer() != nil {
		C.QVideoDeviceSelectorControl_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
