package bluetooth

//#include "bluetooth.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QBluetoothLocalDevice struct {
	core.QObject
}

type QBluetoothLocalDevice_ITF interface {
	core.QObject_ITF
	QBluetoothLocalDevice_PTR() *QBluetoothLocalDevice
}

func PointerFromQBluetoothLocalDevice(ptr QBluetoothLocalDevice_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothLocalDevice_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothLocalDeviceFromPointer(ptr unsafe.Pointer) *QBluetoothLocalDevice {
	var n = new(QBluetoothLocalDevice)
	n.SetPointer(ptr)
	for len(n.ObjectName()) < len("QBluetoothLocalDevice_") {
		n.SetObjectName("QBluetoothLocalDevice_" + qt.Identifier())
	}
	return n
}

func (ptr *QBluetoothLocalDevice) QBluetoothLocalDevice_PTR() *QBluetoothLocalDevice {
	return ptr
}

//QBluetoothLocalDevice::Error
type QBluetoothLocalDevice__Error int64

const (
	QBluetoothLocalDevice__NoError      = QBluetoothLocalDevice__Error(0)
	QBluetoothLocalDevice__PairingError = QBluetoothLocalDevice__Error(1)
	QBluetoothLocalDevice__UnknownError = QBluetoothLocalDevice__Error(100)
)

//QBluetoothLocalDevice::HostMode
type QBluetoothLocalDevice__HostMode int64

const (
	QBluetoothLocalDevice__HostPoweredOff                 = QBluetoothLocalDevice__HostMode(0)
	QBluetoothLocalDevice__HostConnectable                = QBluetoothLocalDevice__HostMode(1)
	QBluetoothLocalDevice__HostDiscoverable               = QBluetoothLocalDevice__HostMode(2)
	QBluetoothLocalDevice__HostDiscoverableLimitedInquiry = QBluetoothLocalDevice__HostMode(3)
)

//QBluetoothLocalDevice::Pairing
type QBluetoothLocalDevice__Pairing int64

const (
	QBluetoothLocalDevice__Unpaired         = QBluetoothLocalDevice__Pairing(0)
	QBluetoothLocalDevice__Paired           = QBluetoothLocalDevice__Pairing(1)
	QBluetoothLocalDevice__AuthorizedPaired = QBluetoothLocalDevice__Pairing(2)
)

func (ptr *QBluetoothLocalDevice) ConnectError(f func(error QBluetoothLocalDevice__Error)) {
	defer qt.Recovering("connect QBluetoothLocalDevice::error")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_ConnectError(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error", f)
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectError() {
	defer qt.Recovering("disconnect QBluetoothLocalDevice::error")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DisconnectError(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error")
	}
}

//export callbackQBluetoothLocalDeviceError
func callbackQBluetoothLocalDeviceError(ptr unsafe.Pointer, ptrName *C.char, error C.int) {
	defer qt.Recovering("callback QBluetoothLocalDevice::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error"); signal != nil {
		signal.(func(QBluetoothLocalDevice__Error))(QBluetoothLocalDevice__Error(error))
	}

}

func (ptr *QBluetoothLocalDevice) Error(error QBluetoothLocalDevice__Error) {
	defer qt.Recovering("QBluetoothLocalDevice::error")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_Error(ptr.Pointer(), C.int(error))
	}
}

func (ptr *QBluetoothLocalDevice) ConnectHostModeStateChanged(f func(state QBluetoothLocalDevice__HostMode)) {
	defer qt.Recovering("connect QBluetoothLocalDevice::hostModeStateChanged")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_ConnectHostModeStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "hostModeStateChanged", f)
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectHostModeStateChanged() {
	defer qt.Recovering("disconnect QBluetoothLocalDevice::hostModeStateChanged")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DisconnectHostModeStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "hostModeStateChanged")
	}
}

//export callbackQBluetoothLocalDeviceHostModeStateChanged
func callbackQBluetoothLocalDeviceHostModeStateChanged(ptr unsafe.Pointer, ptrName *C.char, state C.int) {
	defer qt.Recovering("callback QBluetoothLocalDevice::hostModeStateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "hostModeStateChanged"); signal != nil {
		signal.(func(QBluetoothLocalDevice__HostMode))(QBluetoothLocalDevice__HostMode(state))
	}

}

func (ptr *QBluetoothLocalDevice) HostModeStateChanged(state QBluetoothLocalDevice__HostMode) {
	defer qt.Recovering("QBluetoothLocalDevice::hostModeStateChanged")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_HostModeStateChanged(ptr.Pointer(), C.int(state))
	}
}

func (ptr *QBluetoothLocalDevice) IsValid() bool {
	defer qt.Recovering("QBluetoothLocalDevice::isValid")

	if ptr.Pointer() != nil {
		return C.QBluetoothLocalDevice_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothLocalDevice) DestroyQBluetoothLocalDevice() {
	defer qt.Recovering("QBluetoothLocalDevice::~QBluetoothLocalDevice")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DestroyQBluetoothLocalDevice(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func NewQBluetoothLocalDevice(parent core.QObject_ITF) *QBluetoothLocalDevice {
	defer qt.Recovering("QBluetoothLocalDevice::QBluetoothLocalDevice")

	return NewQBluetoothLocalDeviceFromPointer(C.QBluetoothLocalDevice_NewQBluetoothLocalDevice(core.PointerFromQObject(parent)))
}

func NewQBluetoothLocalDevice2(address QBluetoothAddress_ITF, parent core.QObject_ITF) *QBluetoothLocalDevice {
	defer qt.Recovering("QBluetoothLocalDevice::QBluetoothLocalDevice")

	return NewQBluetoothLocalDeviceFromPointer(C.QBluetoothLocalDevice_NewQBluetoothLocalDevice2(PointerFromQBluetoothAddress(address), core.PointerFromQObject(parent)))
}

func (ptr *QBluetoothLocalDevice) HostMode() QBluetoothLocalDevice__HostMode {
	defer qt.Recovering("QBluetoothLocalDevice::hostMode")

	if ptr.Pointer() != nil {
		return QBluetoothLocalDevice__HostMode(C.QBluetoothLocalDevice_HostMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothLocalDevice) Name() string {
	defer qt.Recovering("QBluetoothLocalDevice::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothLocalDevice_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothLocalDevice) PairingConfirmation(accept bool) {
	defer qt.Recovering("QBluetoothLocalDevice::pairingConfirmation")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_PairingConfirmation(ptr.Pointer(), C.int(qt.GoBoolToInt(accept)))
	}
}

func (ptr *QBluetoothLocalDevice) PairingStatus(address QBluetoothAddress_ITF) QBluetoothLocalDevice__Pairing {
	defer qt.Recovering("QBluetoothLocalDevice::pairingStatus")

	if ptr.Pointer() != nil {
		return QBluetoothLocalDevice__Pairing(C.QBluetoothLocalDevice_PairingStatus(ptr.Pointer(), PointerFromQBluetoothAddress(address)))
	}
	return 0
}

func (ptr *QBluetoothLocalDevice) PowerOn() {
	defer qt.Recovering("QBluetoothLocalDevice::powerOn")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_PowerOn(ptr.Pointer())
	}
}

func (ptr *QBluetoothLocalDevice) RequestPairing(address QBluetoothAddress_ITF, pairing QBluetoothLocalDevice__Pairing) {
	defer qt.Recovering("QBluetoothLocalDevice::requestPairing")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_RequestPairing(ptr.Pointer(), PointerFromQBluetoothAddress(address), C.int(pairing))
	}
}

func (ptr *QBluetoothLocalDevice) SetHostMode(mode QBluetoothLocalDevice__HostMode) {
	defer qt.Recovering("QBluetoothLocalDevice::setHostMode")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_SetHostMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QBluetoothLocalDevice) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QBluetoothLocalDevice::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QBluetoothLocalDevice::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

//export callbackQBluetoothLocalDeviceTimerEvent
func callbackQBluetoothLocalDeviceTimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothLocalDevice::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQBluetoothLocalDeviceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QBluetoothLocalDevice) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QBluetoothLocalDevice::timerEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QBluetoothLocalDevice) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QBluetoothLocalDevice::timerEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QBluetoothLocalDevice) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QBluetoothLocalDevice::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QBluetoothLocalDevice::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

//export callbackQBluetoothLocalDeviceChildEvent
func callbackQBluetoothLocalDeviceChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothLocalDevice::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQBluetoothLocalDeviceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QBluetoothLocalDevice) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QBluetoothLocalDevice::childEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QBluetoothLocalDevice) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QBluetoothLocalDevice::childEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QBluetoothLocalDevice) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QBluetoothLocalDevice::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QBluetoothLocalDevice::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

//export callbackQBluetoothLocalDeviceCustomEvent
func callbackQBluetoothLocalDeviceCustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothLocalDevice::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQBluetoothLocalDeviceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QBluetoothLocalDevice) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QBluetoothLocalDevice::customEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QBluetoothLocalDevice) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QBluetoothLocalDevice::customEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}
