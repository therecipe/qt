// +build !minimal

package bluetooth

//#include <stdlib.h>
//#include "bluetooth.h"
import "C"
import (
	"encoding/hex"
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/network"
	"runtime"
	"unsafe"
)

type QBluetoothAddress struct {
	ptr unsafe.Pointer
}

type QBluetoothAddress_ITF interface {
	QBluetoothAddress_PTR() *QBluetoothAddress
}

func (p *QBluetoothAddress) QBluetoothAddress_PTR() *QBluetoothAddress {
	return p
}

func (p *QBluetoothAddress) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QBluetoothAddress) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQBluetoothAddress(ptr QBluetoothAddress_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothAddress_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothAddressFromPointer(ptr unsafe.Pointer) *QBluetoothAddress {
	var n = new(QBluetoothAddress)
	n.SetPointer(ptr)
	return n
}

func newQBluetoothAddressFromPointer(ptr unsafe.Pointer) *QBluetoothAddress {
	var n = NewQBluetoothAddressFromPointer(ptr)
	return n
}

func NewQBluetoothAddress() *QBluetoothAddress {
	defer qt.Recovering("QBluetoothAddress::QBluetoothAddress")

	return newQBluetoothAddressFromPointer(C.QBluetoothAddress_NewQBluetoothAddress())
}

func NewQBluetoothAddress4(other QBluetoothAddress_ITF) *QBluetoothAddress {
	defer qt.Recovering("QBluetoothAddress::QBluetoothAddress")

	return newQBluetoothAddressFromPointer(C.QBluetoothAddress_NewQBluetoothAddress4(PointerFromQBluetoothAddress(other)))
}

func NewQBluetoothAddress3(address string) *QBluetoothAddress {
	defer qt.Recovering("QBluetoothAddress::QBluetoothAddress")

	var addressC = C.CString(address)
	defer C.free(unsafe.Pointer(addressC))
	return newQBluetoothAddressFromPointer(C.QBluetoothAddress_NewQBluetoothAddress3(addressC))
}

func (ptr *QBluetoothAddress) Clear() {
	defer qt.Recovering("QBluetoothAddress::clear")

	if ptr.Pointer() != nil {
		C.QBluetoothAddress_Clear(ptr.Pointer())
	}
}

func (ptr *QBluetoothAddress) IsNull() bool {
	defer qt.Recovering("QBluetoothAddress::isNull")

	if ptr.Pointer() != nil {
		return C.QBluetoothAddress_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothAddress) ToString() string {
	defer qt.Recovering("QBluetoothAddress::toString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothAddress_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothAddress) DestroyQBluetoothAddress() {
	defer qt.Recovering("QBluetoothAddress::~QBluetoothAddress")

	if ptr.Pointer() != nil {
		C.QBluetoothAddress_DestroyQBluetoothAddress(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//QBluetoothDeviceDiscoveryAgent::Error
type QBluetoothDeviceDiscoveryAgent__Error int64

const (
	QBluetoothDeviceDiscoveryAgent__NoError                      = QBluetoothDeviceDiscoveryAgent__Error(0)
	QBluetoothDeviceDiscoveryAgent__InputOutputError             = QBluetoothDeviceDiscoveryAgent__Error(1)
	QBluetoothDeviceDiscoveryAgent__PoweredOffError              = QBluetoothDeviceDiscoveryAgent__Error(2)
	QBluetoothDeviceDiscoveryAgent__InvalidBluetoothAdapterError = QBluetoothDeviceDiscoveryAgent__Error(3)
	QBluetoothDeviceDiscoveryAgent__UnsupportedPlatformError     = QBluetoothDeviceDiscoveryAgent__Error(4)
	QBluetoothDeviceDiscoveryAgent__UnknownError                 = QBluetoothDeviceDiscoveryAgent__Error(100)
)

//QBluetoothDeviceDiscoveryAgent::InquiryType
type QBluetoothDeviceDiscoveryAgent__InquiryType int64

const (
	QBluetoothDeviceDiscoveryAgent__GeneralUnlimitedInquiry = QBluetoothDeviceDiscoveryAgent__InquiryType(0)
	QBluetoothDeviceDiscoveryAgent__LimitedInquiry          = QBluetoothDeviceDiscoveryAgent__InquiryType(1)
)

type QBluetoothDeviceDiscoveryAgent struct {
	core.QObject
}

type QBluetoothDeviceDiscoveryAgent_ITF interface {
	core.QObject_ITF
	QBluetoothDeviceDiscoveryAgent_PTR() *QBluetoothDeviceDiscoveryAgent
}

func (p *QBluetoothDeviceDiscoveryAgent) QBluetoothDeviceDiscoveryAgent_PTR() *QBluetoothDeviceDiscoveryAgent {
	return p
}

func (p *QBluetoothDeviceDiscoveryAgent) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QBluetoothDeviceDiscoveryAgent) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQBluetoothDeviceDiscoveryAgent(ptr QBluetoothDeviceDiscoveryAgent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothDeviceDiscoveryAgent_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr unsafe.Pointer) *QBluetoothDeviceDiscoveryAgent {
	var n = new(QBluetoothDeviceDiscoveryAgent)
	n.SetPointer(ptr)
	return n
}

func newQBluetoothDeviceDiscoveryAgentFromPointer(ptr unsafe.Pointer) *QBluetoothDeviceDiscoveryAgent {
	var n = NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr)
	for len(n.ObjectName()) < len("QBluetoothDeviceDiscoveryAgent_") {
		n.SetObjectName("QBluetoothDeviceDiscoveryAgent_" + qt.Identifier())
	}
	return n
}

//export callbackQBluetoothDeviceDiscoveryAgent_Canceled
func callbackQBluetoothDeviceDiscoveryAgent_Canceled(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QBluetoothDeviceDiscoveryAgent::canceled")

	if signal := qt.GetSignal(C.GoString(ptrName), "canceled"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectCanceled(f func()) {
	defer qt.Recovering("connect QBluetoothDeviceDiscoveryAgent::canceled")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_ConnectCanceled(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "canceled", f)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectCanceled() {
	defer qt.Recovering("disconnect QBluetoothDeviceDiscoveryAgent::canceled")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_DisconnectCanceled(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "canceled")
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) Canceled() {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::canceled")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_Canceled(ptr.Pointer())
	}
}

//export callbackQBluetoothDeviceDiscoveryAgent_DeviceDiscovered
func callbackQBluetoothDeviceDiscoveryAgent_DeviceDiscovered(ptr unsafe.Pointer, ptrName *C.char, info unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothDeviceDiscoveryAgent::deviceDiscovered")

	if signal := qt.GetSignal(C.GoString(ptrName), "deviceDiscovered"); signal != nil {
		signal.(func(*QBluetoothDeviceInfo))(NewQBluetoothDeviceInfoFromPointer(info))
	}

}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectDeviceDiscovered(f func(info *QBluetoothDeviceInfo)) {
	defer qt.Recovering("connect QBluetoothDeviceDiscoveryAgent::deviceDiscovered")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_ConnectDeviceDiscovered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "deviceDiscovered", f)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectDeviceDiscovered() {
	defer qt.Recovering("disconnect QBluetoothDeviceDiscoveryAgent::deviceDiscovered")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_DisconnectDeviceDiscovered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "deviceDiscovered")
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DeviceDiscovered(info QBluetoothDeviceInfo_ITF) {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::deviceDiscovered")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_DeviceDiscovered(ptr.Pointer(), PointerFromQBluetoothDeviceInfo(info))
	}
}

//export callbackQBluetoothDeviceDiscoveryAgent_Error2
func callbackQBluetoothDeviceDiscoveryAgent_Error2(ptr unsafe.Pointer, ptrName *C.char, error C.int) {
	defer qt.Recovering("callback QBluetoothDeviceDiscoveryAgent::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(QBluetoothDeviceDiscoveryAgent__Error))(QBluetoothDeviceDiscoveryAgent__Error(error))
	}

}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectError2(f func(error QBluetoothDeviceDiscoveryAgent__Error)) {
	defer qt.Recovering("connect QBluetoothDeviceDiscoveryAgent::error")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error2", f)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectError2() {
	defer qt.Recovering("disconnect QBluetoothDeviceDiscoveryAgent::error")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error2")
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) Error2(error QBluetoothDeviceDiscoveryAgent__Error) {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::error")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_Error2(ptr.Pointer(), C.int(error))
	}
}

//export callbackQBluetoothDeviceDiscoveryAgent_Finished
func callbackQBluetoothDeviceDiscoveryAgent_Finished(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QBluetoothDeviceDiscoveryAgent::finished")

	if signal := qt.GetSignal(C.GoString(ptrName), "finished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectFinished(f func()) {
	defer qt.Recovering("connect QBluetoothDeviceDiscoveryAgent::finished")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectFinished() {
	defer qt.Recovering("disconnect QBluetoothDeviceDiscoveryAgent::finished")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) Finished() {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::finished")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_Finished(ptr.Pointer())
	}
}

func NewQBluetoothDeviceDiscoveryAgent(parent core.QObject_ITF) *QBluetoothDeviceDiscoveryAgent {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::QBluetoothDeviceDiscoveryAgent")

	return newQBluetoothDeviceDiscoveryAgentFromPointer(C.QBluetoothDeviceDiscoveryAgent_NewQBluetoothDeviceDiscoveryAgent(core.PointerFromQObject(parent)))
}

func NewQBluetoothDeviceDiscoveryAgent2(deviceAdapter QBluetoothAddress_ITF, parent core.QObject_ITF) *QBluetoothDeviceDiscoveryAgent {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::QBluetoothDeviceDiscoveryAgent")

	return newQBluetoothDeviceDiscoveryAgentFromPointer(C.QBluetoothDeviceDiscoveryAgent_NewQBluetoothDeviceDiscoveryAgent2(PointerFromQBluetoothAddress(deviceAdapter), core.PointerFromQObject(parent)))
}

func (ptr *QBluetoothDeviceDiscoveryAgent) Error() QBluetoothDeviceDiscoveryAgent__Error {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::error")

	if ptr.Pointer() != nil {
		return QBluetoothDeviceDiscoveryAgent__Error(C.QBluetoothDeviceDiscoveryAgent_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ErrorString() string {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothDeviceDiscoveryAgent_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothDeviceDiscoveryAgent) InquiryType() QBluetoothDeviceDiscoveryAgent__InquiryType {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::inquiryType")

	if ptr.Pointer() != nil {
		return QBluetoothDeviceDiscoveryAgent__InquiryType(C.QBluetoothDeviceDiscoveryAgent_InquiryType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothDeviceDiscoveryAgent) IsActive() bool {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::isActive")

	if ptr.Pointer() != nil {
		return C.QBluetoothDeviceDiscoveryAgent_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothDeviceDiscoveryAgent) SetInquiryType(ty QBluetoothDeviceDiscoveryAgent__InquiryType) {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::setInquiryType")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_SetInquiryType(ptr.Pointer(), C.int(ty))
	}
}

//export callbackQBluetoothDeviceDiscoveryAgent_Start
func callbackQBluetoothDeviceDiscoveryAgent_Start(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QBluetoothDeviceDiscoveryAgent::start")

	if signal := qt.GetSignal(C.GoString(ptrName), "start"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectStart(f func()) {
	defer qt.Recovering("connect QBluetoothDeviceDiscoveryAgent::start")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "start", f)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectStart() {
	defer qt.Recovering("disconnect QBluetoothDeviceDiscoveryAgent::start")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "start")
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) Start() {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::start")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_Start(ptr.Pointer())
	}
}

//export callbackQBluetoothDeviceDiscoveryAgent_Stop
func callbackQBluetoothDeviceDiscoveryAgent_Stop(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QBluetoothDeviceDiscoveryAgent::stop")

	if signal := qt.GetSignal(C.GoString(ptrName), "stop"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectStop(f func()) {
	defer qt.Recovering("connect QBluetoothDeviceDiscoveryAgent::stop")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "stop", f)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectStop() {
	defer qt.Recovering("disconnect QBluetoothDeviceDiscoveryAgent::stop")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "stop")
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) Stop() {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::stop")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_Stop(ptr.Pointer())
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DestroyQBluetoothDeviceDiscoveryAgent() {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::~QBluetoothDeviceDiscoveryAgent")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QBluetoothDeviceDiscoveryAgent_DestroyQBluetoothDeviceDiscoveryAgent(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQBluetoothDeviceDiscoveryAgent_TimerEvent
func callbackQBluetoothDeviceDiscoveryAgent_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothDeviceDiscoveryAgent::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QBluetoothDeviceDiscoveryAgent::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QBluetoothDeviceDiscoveryAgent::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::timerEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::timerEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQBluetoothDeviceDiscoveryAgent_ChildEvent
func callbackQBluetoothDeviceDiscoveryAgent_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothDeviceDiscoveryAgent::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QBluetoothDeviceDiscoveryAgent::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QBluetoothDeviceDiscoveryAgent::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::childEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::childEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQBluetoothDeviceDiscoveryAgent_ConnectNotify
func callbackQBluetoothDeviceDiscoveryAgent_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothDeviceDiscoveryAgent::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QBluetoothDeviceDiscoveryAgent::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QBluetoothDeviceDiscoveryAgent::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::connectNotify")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::connectNotify")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBluetoothDeviceDiscoveryAgent_CustomEvent
func callbackQBluetoothDeviceDiscoveryAgent_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothDeviceDiscoveryAgent::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QBluetoothDeviceDiscoveryAgent::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QBluetoothDeviceDiscoveryAgent::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::customEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::customEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQBluetoothDeviceDiscoveryAgent_DeleteLater
func callbackQBluetoothDeviceDiscoveryAgent_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QBluetoothDeviceDiscoveryAgent::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QBluetoothDeviceDiscoveryAgent::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QBluetoothDeviceDiscoveryAgent::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DeleteLater() {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QBluetoothDeviceDiscoveryAgent_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DeleteLaterDefault() {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QBluetoothDeviceDiscoveryAgent_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQBluetoothDeviceDiscoveryAgent_DisconnectNotify
func callbackQBluetoothDeviceDiscoveryAgent_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothDeviceDiscoveryAgent::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QBluetoothDeviceDiscoveryAgent::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QBluetoothDeviceDiscoveryAgent::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBluetoothDeviceDiscoveryAgent_Event
func callbackQBluetoothDeviceDiscoveryAgent_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QBluetoothDeviceDiscoveryAgent::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QBluetoothDeviceDiscoveryAgent::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectEvent() {
	defer qt.Recovering("disconnect QBluetoothDeviceDiscoveryAgent::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::event")

	if ptr.Pointer() != nil {
		return C.QBluetoothDeviceDiscoveryAgent_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QBluetoothDeviceDiscoveryAgent) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::event")

	if ptr.Pointer() != nil {
		return C.QBluetoothDeviceDiscoveryAgent_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQBluetoothDeviceDiscoveryAgent_EventFilter
func callbackQBluetoothDeviceDiscoveryAgent_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QBluetoothDeviceDiscoveryAgent::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QBluetoothDeviceDiscoveryAgent::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QBluetoothDeviceDiscoveryAgent::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::eventFilter")

	if ptr.Pointer() != nil {
		return C.QBluetoothDeviceDiscoveryAgent_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QBluetoothDeviceDiscoveryAgent) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::eventFilter")

	if ptr.Pointer() != nil {
		return C.QBluetoothDeviceDiscoveryAgent_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQBluetoothDeviceDiscoveryAgent_MetaObject
func callbackQBluetoothDeviceDiscoveryAgent_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QBluetoothDeviceDiscoveryAgent::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QBluetoothDeviceDiscoveryAgent::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QBluetoothDeviceDiscoveryAgent::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QBluetoothDeviceDiscoveryAgent_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBluetoothDeviceDiscoveryAgent) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QBluetoothDeviceDiscoveryAgent::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QBluetoothDeviceDiscoveryAgent_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QBluetoothDeviceInfo::CoreConfiguration
type QBluetoothDeviceInfo__CoreConfiguration int64

const (
	QBluetoothDeviceInfo__UnknownCoreConfiguration              = QBluetoothDeviceInfo__CoreConfiguration(0x0)
	QBluetoothDeviceInfo__LowEnergyCoreConfiguration            = QBluetoothDeviceInfo__CoreConfiguration(0x01)
	QBluetoothDeviceInfo__BaseRateCoreConfiguration             = QBluetoothDeviceInfo__CoreConfiguration(0x02)
	QBluetoothDeviceInfo__BaseRateAndLowEnergyCoreConfiguration = QBluetoothDeviceInfo__CoreConfiguration(0x03)
)

//QBluetoothDeviceInfo::DataCompleteness
type QBluetoothDeviceInfo__DataCompleteness int64

const (
	QBluetoothDeviceInfo__DataComplete    = QBluetoothDeviceInfo__DataCompleteness(0)
	QBluetoothDeviceInfo__DataIncomplete  = QBluetoothDeviceInfo__DataCompleteness(1)
	QBluetoothDeviceInfo__DataUnavailable = QBluetoothDeviceInfo__DataCompleteness(2)
)

//QBluetoothDeviceInfo::MajorDeviceClass
type QBluetoothDeviceInfo__MajorDeviceClass int64

const (
	QBluetoothDeviceInfo__MiscellaneousDevice = QBluetoothDeviceInfo__MajorDeviceClass(0)
	QBluetoothDeviceInfo__ComputerDevice      = QBluetoothDeviceInfo__MajorDeviceClass(1)
	QBluetoothDeviceInfo__PhoneDevice         = QBluetoothDeviceInfo__MajorDeviceClass(2)
	QBluetoothDeviceInfo__LANAccessDevice     = QBluetoothDeviceInfo__MajorDeviceClass(3)
	QBluetoothDeviceInfo__AudioVideoDevice    = QBluetoothDeviceInfo__MajorDeviceClass(4)
	QBluetoothDeviceInfo__PeripheralDevice    = QBluetoothDeviceInfo__MajorDeviceClass(5)
	QBluetoothDeviceInfo__ImagingDevice       = QBluetoothDeviceInfo__MajorDeviceClass(6)
	QBluetoothDeviceInfo__WearableDevice      = QBluetoothDeviceInfo__MajorDeviceClass(7)
	QBluetoothDeviceInfo__ToyDevice           = QBluetoothDeviceInfo__MajorDeviceClass(8)
	QBluetoothDeviceInfo__HealthDevice        = QBluetoothDeviceInfo__MajorDeviceClass(9)
	QBluetoothDeviceInfo__UncategorizedDevice = QBluetoothDeviceInfo__MajorDeviceClass(31)
)

//QBluetoothDeviceInfo::MinorAudioVideoClass
type QBluetoothDeviceInfo__MinorAudioVideoClass int64

const (
	QBluetoothDeviceInfo__UncategorizedAudioVideoDevice = QBluetoothDeviceInfo__MinorAudioVideoClass(0)
	QBluetoothDeviceInfo__WearableHeadsetDevice         = QBluetoothDeviceInfo__MinorAudioVideoClass(1)
	QBluetoothDeviceInfo__HandsFreeDevice               = QBluetoothDeviceInfo__MinorAudioVideoClass(2)
	QBluetoothDeviceInfo__Microphone                    = QBluetoothDeviceInfo__MinorAudioVideoClass(4)
	QBluetoothDeviceInfo__Loudspeaker                   = QBluetoothDeviceInfo__MinorAudioVideoClass(5)
	QBluetoothDeviceInfo__Headphones                    = QBluetoothDeviceInfo__MinorAudioVideoClass(6)
	QBluetoothDeviceInfo__PortableAudioDevice           = QBluetoothDeviceInfo__MinorAudioVideoClass(7)
	QBluetoothDeviceInfo__CarAudio                      = QBluetoothDeviceInfo__MinorAudioVideoClass(8)
	QBluetoothDeviceInfo__SetTopBox                     = QBluetoothDeviceInfo__MinorAudioVideoClass(9)
	QBluetoothDeviceInfo__HiFiAudioDevice               = QBluetoothDeviceInfo__MinorAudioVideoClass(10)
	QBluetoothDeviceInfo__Vcr                           = QBluetoothDeviceInfo__MinorAudioVideoClass(11)
	QBluetoothDeviceInfo__VideoCamera                   = QBluetoothDeviceInfo__MinorAudioVideoClass(12)
	QBluetoothDeviceInfo__Camcorder                     = QBluetoothDeviceInfo__MinorAudioVideoClass(13)
	QBluetoothDeviceInfo__VideoMonitor                  = QBluetoothDeviceInfo__MinorAudioVideoClass(14)
	QBluetoothDeviceInfo__VideoDisplayAndLoudspeaker    = QBluetoothDeviceInfo__MinorAudioVideoClass(15)
	QBluetoothDeviceInfo__VideoConferencing             = QBluetoothDeviceInfo__MinorAudioVideoClass(16)
	QBluetoothDeviceInfo__GamingDevice                  = QBluetoothDeviceInfo__MinorAudioVideoClass(18)
)

//QBluetoothDeviceInfo::MinorComputerClass
type QBluetoothDeviceInfo__MinorComputerClass int64

const (
	QBluetoothDeviceInfo__UncategorizedComputer     = QBluetoothDeviceInfo__MinorComputerClass(0)
	QBluetoothDeviceInfo__DesktopComputer           = QBluetoothDeviceInfo__MinorComputerClass(1)
	QBluetoothDeviceInfo__ServerComputer            = QBluetoothDeviceInfo__MinorComputerClass(2)
	QBluetoothDeviceInfo__LaptopComputer            = QBluetoothDeviceInfo__MinorComputerClass(3)
	QBluetoothDeviceInfo__HandheldClamShellComputer = QBluetoothDeviceInfo__MinorComputerClass(4)
	QBluetoothDeviceInfo__HandheldComputer          = QBluetoothDeviceInfo__MinorComputerClass(5)
	QBluetoothDeviceInfo__WearableComputer          = QBluetoothDeviceInfo__MinorComputerClass(6)
)

//QBluetoothDeviceInfo::MinorHealthClass
type QBluetoothDeviceInfo__MinorHealthClass int64

const (
	QBluetoothDeviceInfo__UncategorizedHealthDevice  = QBluetoothDeviceInfo__MinorHealthClass(0)
	QBluetoothDeviceInfo__HealthBloodPressureMonitor = QBluetoothDeviceInfo__MinorHealthClass(0x1)
	QBluetoothDeviceInfo__HealthThermometer          = QBluetoothDeviceInfo__MinorHealthClass(0x2)
	QBluetoothDeviceInfo__HealthWeightScale          = QBluetoothDeviceInfo__MinorHealthClass(0x3)
	QBluetoothDeviceInfo__HealthGlucoseMeter         = QBluetoothDeviceInfo__MinorHealthClass(0x4)
	QBluetoothDeviceInfo__HealthPulseOximeter        = QBluetoothDeviceInfo__MinorHealthClass(0x5)
	QBluetoothDeviceInfo__HealthDataDisplay          = QBluetoothDeviceInfo__MinorHealthClass(0x7)
	QBluetoothDeviceInfo__HealthStepCounter          = QBluetoothDeviceInfo__MinorHealthClass(0x8)
)

//QBluetoothDeviceInfo::MinorImagingClass
type QBluetoothDeviceInfo__MinorImagingClass int64

const (
	QBluetoothDeviceInfo__UncategorizedImagingDevice = QBluetoothDeviceInfo__MinorImagingClass(0)
	QBluetoothDeviceInfo__ImageDisplay               = QBluetoothDeviceInfo__MinorImagingClass(0x04)
	QBluetoothDeviceInfo__ImageCamera                = QBluetoothDeviceInfo__MinorImagingClass(0x08)
	QBluetoothDeviceInfo__ImageScanner               = QBluetoothDeviceInfo__MinorImagingClass(0x10)
	QBluetoothDeviceInfo__ImagePrinter               = QBluetoothDeviceInfo__MinorImagingClass(0x20)
)

//QBluetoothDeviceInfo::MinorMiscellaneousClass
type QBluetoothDeviceInfo__MinorMiscellaneousClass int64

const (
	QBluetoothDeviceInfo__UncategorizedMiscellaneous = QBluetoothDeviceInfo__MinorMiscellaneousClass(0)
)

//QBluetoothDeviceInfo::MinorNetworkClass
type QBluetoothDeviceInfo__MinorNetworkClass int64

const (
	QBluetoothDeviceInfo__NetworkFullService     = QBluetoothDeviceInfo__MinorNetworkClass(0x00)
	QBluetoothDeviceInfo__NetworkLoadFactorOne   = QBluetoothDeviceInfo__MinorNetworkClass(0x08)
	QBluetoothDeviceInfo__NetworkLoadFactorTwo   = QBluetoothDeviceInfo__MinorNetworkClass(0x10)
	QBluetoothDeviceInfo__NetworkLoadFactorThree = QBluetoothDeviceInfo__MinorNetworkClass(0x18)
	QBluetoothDeviceInfo__NetworkLoadFactorFour  = QBluetoothDeviceInfo__MinorNetworkClass(0x20)
	QBluetoothDeviceInfo__NetworkLoadFactorFive  = QBluetoothDeviceInfo__MinorNetworkClass(0x28)
	QBluetoothDeviceInfo__NetworkLoadFactorSix   = QBluetoothDeviceInfo__MinorNetworkClass(0x30)
	QBluetoothDeviceInfo__NetworkNoService       = QBluetoothDeviceInfo__MinorNetworkClass(0x38)
)

//QBluetoothDeviceInfo::MinorPeripheralClass
type QBluetoothDeviceInfo__MinorPeripheralClass int64

const (
	QBluetoothDeviceInfo__UncategorizedPeripheral              = QBluetoothDeviceInfo__MinorPeripheralClass(0)
	QBluetoothDeviceInfo__KeyboardPeripheral                   = QBluetoothDeviceInfo__MinorPeripheralClass(0x10)
	QBluetoothDeviceInfo__PointingDevicePeripheral             = QBluetoothDeviceInfo__MinorPeripheralClass(0x20)
	QBluetoothDeviceInfo__KeyboardWithPointingDevicePeripheral = QBluetoothDeviceInfo__MinorPeripheralClass(0x30)
	QBluetoothDeviceInfo__JoystickPeripheral                   = QBluetoothDeviceInfo__MinorPeripheralClass(0x01)
	QBluetoothDeviceInfo__GamepadPeripheral                    = QBluetoothDeviceInfo__MinorPeripheralClass(0x02)
	QBluetoothDeviceInfo__RemoteControlPeripheral              = QBluetoothDeviceInfo__MinorPeripheralClass(0x03)
	QBluetoothDeviceInfo__SensingDevicePeripheral              = QBluetoothDeviceInfo__MinorPeripheralClass(0x04)
	QBluetoothDeviceInfo__DigitizerTabletPeripheral            = QBluetoothDeviceInfo__MinorPeripheralClass(0x05)
	QBluetoothDeviceInfo__CardReaderPeripheral                 = QBluetoothDeviceInfo__MinorPeripheralClass(0x06)
)

//QBluetoothDeviceInfo::MinorPhoneClass
type QBluetoothDeviceInfo__MinorPhoneClass int64

const (
	QBluetoothDeviceInfo__UncategorizedPhone            = QBluetoothDeviceInfo__MinorPhoneClass(0)
	QBluetoothDeviceInfo__CellularPhone                 = QBluetoothDeviceInfo__MinorPhoneClass(1)
	QBluetoothDeviceInfo__CordlessPhone                 = QBluetoothDeviceInfo__MinorPhoneClass(2)
	QBluetoothDeviceInfo__SmartPhone                    = QBluetoothDeviceInfo__MinorPhoneClass(3)
	QBluetoothDeviceInfo__WiredModemOrVoiceGatewayPhone = QBluetoothDeviceInfo__MinorPhoneClass(4)
	QBluetoothDeviceInfo__CommonIsdnAccessPhone         = QBluetoothDeviceInfo__MinorPhoneClass(5)
)

//QBluetoothDeviceInfo::MinorToyClass
type QBluetoothDeviceInfo__MinorToyClass int64

const (
	QBluetoothDeviceInfo__UncategorizedToy = QBluetoothDeviceInfo__MinorToyClass(0)
	QBluetoothDeviceInfo__ToyRobot         = QBluetoothDeviceInfo__MinorToyClass(1)
	QBluetoothDeviceInfo__ToyVehicle       = QBluetoothDeviceInfo__MinorToyClass(2)
	QBluetoothDeviceInfo__ToyDoll          = QBluetoothDeviceInfo__MinorToyClass(3)
	QBluetoothDeviceInfo__ToyController    = QBluetoothDeviceInfo__MinorToyClass(4)
	QBluetoothDeviceInfo__ToyGame          = QBluetoothDeviceInfo__MinorToyClass(5)
)

//QBluetoothDeviceInfo::MinorWearableClass
type QBluetoothDeviceInfo__MinorWearableClass int64

const (
	QBluetoothDeviceInfo__UncategorizedWearableDevice = QBluetoothDeviceInfo__MinorWearableClass(0)
	QBluetoothDeviceInfo__WearableWristWatch          = QBluetoothDeviceInfo__MinorWearableClass(1)
	QBluetoothDeviceInfo__WearablePager               = QBluetoothDeviceInfo__MinorWearableClass(2)
	QBluetoothDeviceInfo__WearableJacket              = QBluetoothDeviceInfo__MinorWearableClass(3)
	QBluetoothDeviceInfo__WearableHelmet              = QBluetoothDeviceInfo__MinorWearableClass(4)
	QBluetoothDeviceInfo__WearableGlasses             = QBluetoothDeviceInfo__MinorWearableClass(5)
)

//QBluetoothDeviceInfo::ServiceClass
type QBluetoothDeviceInfo__ServiceClass int64

const (
	QBluetoothDeviceInfo__NoService             = QBluetoothDeviceInfo__ServiceClass(0x0000)
	QBluetoothDeviceInfo__PositioningService    = QBluetoothDeviceInfo__ServiceClass(0x0001)
	QBluetoothDeviceInfo__NetworkingService     = QBluetoothDeviceInfo__ServiceClass(0x0002)
	QBluetoothDeviceInfo__RenderingService      = QBluetoothDeviceInfo__ServiceClass(0x0004)
	QBluetoothDeviceInfo__CapturingService      = QBluetoothDeviceInfo__ServiceClass(0x0008)
	QBluetoothDeviceInfo__ObjectTransferService = QBluetoothDeviceInfo__ServiceClass(0x0010)
	QBluetoothDeviceInfo__AudioService          = QBluetoothDeviceInfo__ServiceClass(0x0020)
	QBluetoothDeviceInfo__TelephonyService      = QBluetoothDeviceInfo__ServiceClass(0x0040)
	QBluetoothDeviceInfo__InformationService    = QBluetoothDeviceInfo__ServiceClass(0x0080)
	QBluetoothDeviceInfo__AllServices           = QBluetoothDeviceInfo__ServiceClass(0x07ff)
)

type QBluetoothDeviceInfo struct {
	ptr unsafe.Pointer
}

type QBluetoothDeviceInfo_ITF interface {
	QBluetoothDeviceInfo_PTR() *QBluetoothDeviceInfo
}

func (p *QBluetoothDeviceInfo) QBluetoothDeviceInfo_PTR() *QBluetoothDeviceInfo {
	return p
}

func (p *QBluetoothDeviceInfo) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QBluetoothDeviceInfo) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQBluetoothDeviceInfo(ptr QBluetoothDeviceInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothDeviceInfo_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothDeviceInfoFromPointer(ptr unsafe.Pointer) *QBluetoothDeviceInfo {
	var n = new(QBluetoothDeviceInfo)
	n.SetPointer(ptr)
	return n
}

func newQBluetoothDeviceInfoFromPointer(ptr unsafe.Pointer) *QBluetoothDeviceInfo {
	var n = NewQBluetoothDeviceInfoFromPointer(ptr)
	return n
}

func NewQBluetoothDeviceInfo() *QBluetoothDeviceInfo {
	defer qt.Recovering("QBluetoothDeviceInfo::QBluetoothDeviceInfo")

	return newQBluetoothDeviceInfoFromPointer(C.QBluetoothDeviceInfo_NewQBluetoothDeviceInfo())
}

func NewQBluetoothDeviceInfo4(other QBluetoothDeviceInfo_ITF) *QBluetoothDeviceInfo {
	defer qt.Recovering("QBluetoothDeviceInfo::QBluetoothDeviceInfo")

	return newQBluetoothDeviceInfoFromPointer(C.QBluetoothDeviceInfo_NewQBluetoothDeviceInfo4(PointerFromQBluetoothDeviceInfo(other)))
}

func (ptr *QBluetoothDeviceInfo) Address() *QBluetoothAddress {
	defer qt.Recovering("QBluetoothDeviceInfo::address")

	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothAddressFromPointer(C.QBluetoothDeviceInfo_Address(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothDeviceInfo) CoreConfigurations() QBluetoothDeviceInfo__CoreConfiguration {
	defer qt.Recovering("QBluetoothDeviceInfo::coreConfigurations")

	if ptr.Pointer() != nil {
		return QBluetoothDeviceInfo__CoreConfiguration(C.QBluetoothDeviceInfo_CoreConfigurations(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothDeviceInfo) DeviceUuid() *QBluetoothUuid {
	defer qt.Recovering("QBluetoothDeviceInfo::deviceUuid")

	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothUuidFromPointer(C.QBluetoothDeviceInfo_DeviceUuid(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothDeviceInfo) IsCached() bool {
	defer qt.Recovering("QBluetoothDeviceInfo::isCached")

	if ptr.Pointer() != nil {
		return C.QBluetoothDeviceInfo_IsCached(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothDeviceInfo) IsValid() bool {
	defer qt.Recovering("QBluetoothDeviceInfo::isValid")

	if ptr.Pointer() != nil {
		return C.QBluetoothDeviceInfo_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothDeviceInfo) MajorDeviceClass() QBluetoothDeviceInfo__MajorDeviceClass {
	defer qt.Recovering("QBluetoothDeviceInfo::majorDeviceClass")

	if ptr.Pointer() != nil {
		return QBluetoothDeviceInfo__MajorDeviceClass(C.QBluetoothDeviceInfo_MajorDeviceClass(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothDeviceInfo) Name() string {
	defer qt.Recovering("QBluetoothDeviceInfo::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothDeviceInfo_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothDeviceInfo) ServiceClasses() QBluetoothDeviceInfo__ServiceClass {
	defer qt.Recovering("QBluetoothDeviceInfo::serviceClasses")

	if ptr.Pointer() != nil {
		return QBluetoothDeviceInfo__ServiceClass(C.QBluetoothDeviceInfo_ServiceClasses(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothDeviceInfo) ServiceUuidsCompleteness() QBluetoothDeviceInfo__DataCompleteness {
	defer qt.Recovering("QBluetoothDeviceInfo::serviceUuidsCompleteness")

	if ptr.Pointer() != nil {
		return QBluetoothDeviceInfo__DataCompleteness(C.QBluetoothDeviceInfo_ServiceUuidsCompleteness(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothDeviceInfo) SetCached(cached bool) {
	defer qt.Recovering("QBluetoothDeviceInfo::setCached")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceInfo_SetCached(ptr.Pointer(), C.int(qt.GoBoolToInt(cached)))
	}
}

func (ptr *QBluetoothDeviceInfo) SetCoreConfigurations(coreConfigs QBluetoothDeviceInfo__CoreConfiguration) {
	defer qt.Recovering("QBluetoothDeviceInfo::setCoreConfigurations")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceInfo_SetCoreConfigurations(ptr.Pointer(), C.int(coreConfigs))
	}
}

func (ptr *QBluetoothDeviceInfo) SetDeviceUuid(uuid QBluetoothUuid_ITF) {
	defer qt.Recovering("QBluetoothDeviceInfo::setDeviceUuid")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceInfo_SetDeviceUuid(ptr.Pointer(), PointerFromQBluetoothUuid(uuid))
	}
}

func (ptr *QBluetoothDeviceInfo) DestroyQBluetoothDeviceInfo() {
	defer qt.Recovering("QBluetoothDeviceInfo::~QBluetoothDeviceInfo")

	if ptr.Pointer() != nil {
		C.QBluetoothDeviceInfo_DestroyQBluetoothDeviceInfo(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QBluetoothHostInfo struct {
	ptr unsafe.Pointer
}

type QBluetoothHostInfo_ITF interface {
	QBluetoothHostInfo_PTR() *QBluetoothHostInfo
}

func (p *QBluetoothHostInfo) QBluetoothHostInfo_PTR() *QBluetoothHostInfo {
	return p
}

func (p *QBluetoothHostInfo) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QBluetoothHostInfo) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQBluetoothHostInfo(ptr QBluetoothHostInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothHostInfo_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothHostInfoFromPointer(ptr unsafe.Pointer) *QBluetoothHostInfo {
	var n = new(QBluetoothHostInfo)
	n.SetPointer(ptr)
	return n
}

func newQBluetoothHostInfoFromPointer(ptr unsafe.Pointer) *QBluetoothHostInfo {
	var n = NewQBluetoothHostInfoFromPointer(ptr)
	return n
}

func NewQBluetoothHostInfo() *QBluetoothHostInfo {
	defer qt.Recovering("QBluetoothHostInfo::QBluetoothHostInfo")

	return newQBluetoothHostInfoFromPointer(C.QBluetoothHostInfo_NewQBluetoothHostInfo())
}

func NewQBluetoothHostInfo2(other QBluetoothHostInfo_ITF) *QBluetoothHostInfo {
	defer qt.Recovering("QBluetoothHostInfo::QBluetoothHostInfo")

	return newQBluetoothHostInfoFromPointer(C.QBluetoothHostInfo_NewQBluetoothHostInfo2(PointerFromQBluetoothHostInfo(other)))
}

func (ptr *QBluetoothHostInfo) Address() *QBluetoothAddress {
	defer qt.Recovering("QBluetoothHostInfo::address")

	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothAddressFromPointer(C.QBluetoothHostInfo_Address(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothHostInfo) Name() string {
	defer qt.Recovering("QBluetoothHostInfo::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothHostInfo_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothHostInfo) SetAddress(address QBluetoothAddress_ITF) {
	defer qt.Recovering("QBluetoothHostInfo::setAddress")

	if ptr.Pointer() != nil {
		C.QBluetoothHostInfo_SetAddress(ptr.Pointer(), PointerFromQBluetoothAddress(address))
	}
}

func (ptr *QBluetoothHostInfo) SetName(name string) {
	defer qt.Recovering("QBluetoothHostInfo::setName")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QBluetoothHostInfo_SetName(ptr.Pointer(), nameC)
	}
}

func (ptr *QBluetoothHostInfo) DestroyQBluetoothHostInfo() {
	defer qt.Recovering("QBluetoothHostInfo::~QBluetoothHostInfo")

	if ptr.Pointer() != nil {
		C.QBluetoothHostInfo_DestroyQBluetoothHostInfo(ptr.Pointer())
		ptr.SetPointer(nil)
	}
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

type QBluetoothLocalDevice struct {
	core.QObject
}

type QBluetoothLocalDevice_ITF interface {
	core.QObject_ITF
	QBluetoothLocalDevice_PTR() *QBluetoothLocalDevice
}

func (p *QBluetoothLocalDevice) QBluetoothLocalDevice_PTR() *QBluetoothLocalDevice {
	return p
}

func (p *QBluetoothLocalDevice) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QBluetoothLocalDevice) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
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
	return n
}

func newQBluetoothLocalDeviceFromPointer(ptr unsafe.Pointer) *QBluetoothLocalDevice {
	var n = NewQBluetoothLocalDeviceFromPointer(ptr)
	for len(n.ObjectName()) < len("QBluetoothLocalDevice_") {
		n.SetObjectName("QBluetoothLocalDevice_" + qt.Identifier())
	}
	return n
}

//export callbackQBluetoothLocalDevice_DeviceConnected
func callbackQBluetoothLocalDevice_DeviceConnected(ptr unsafe.Pointer, ptrName *C.char, address unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothLocalDevice::deviceConnected")

	if signal := qt.GetSignal(C.GoString(ptrName), "deviceConnected"); signal != nil {
		signal.(func(*QBluetoothAddress))(NewQBluetoothAddressFromPointer(address))
	}

}

func (ptr *QBluetoothLocalDevice) ConnectDeviceConnected(f func(address *QBluetoothAddress)) {
	defer qt.Recovering("connect QBluetoothLocalDevice::deviceConnected")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_ConnectDeviceConnected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "deviceConnected", f)
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectDeviceConnected() {
	defer qt.Recovering("disconnect QBluetoothLocalDevice::deviceConnected")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DisconnectDeviceConnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "deviceConnected")
	}
}

func (ptr *QBluetoothLocalDevice) DeviceConnected(address QBluetoothAddress_ITF) {
	defer qt.Recovering("QBluetoothLocalDevice::deviceConnected")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DeviceConnected(ptr.Pointer(), PointerFromQBluetoothAddress(address))
	}
}

//export callbackQBluetoothLocalDevice_DeviceDisconnected
func callbackQBluetoothLocalDevice_DeviceDisconnected(ptr unsafe.Pointer, ptrName *C.char, address unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothLocalDevice::deviceDisconnected")

	if signal := qt.GetSignal(C.GoString(ptrName), "deviceDisconnected"); signal != nil {
		signal.(func(*QBluetoothAddress))(NewQBluetoothAddressFromPointer(address))
	}

}

func (ptr *QBluetoothLocalDevice) ConnectDeviceDisconnected(f func(address *QBluetoothAddress)) {
	defer qt.Recovering("connect QBluetoothLocalDevice::deviceDisconnected")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_ConnectDeviceDisconnected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "deviceDisconnected", f)
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectDeviceDisconnected() {
	defer qt.Recovering("disconnect QBluetoothLocalDevice::deviceDisconnected")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DisconnectDeviceDisconnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "deviceDisconnected")
	}
}

func (ptr *QBluetoothLocalDevice) DeviceDisconnected(address QBluetoothAddress_ITF) {
	defer qt.Recovering("QBluetoothLocalDevice::deviceDisconnected")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DeviceDisconnected(ptr.Pointer(), PointerFromQBluetoothAddress(address))
	}
}

//export callbackQBluetoothLocalDevice_Error
func callbackQBluetoothLocalDevice_Error(ptr unsafe.Pointer, ptrName *C.char, error C.int) {
	defer qt.Recovering("callback QBluetoothLocalDevice::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error"); signal != nil {
		signal.(func(QBluetoothLocalDevice__Error))(QBluetoothLocalDevice__Error(error))
	}

}

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

func (ptr *QBluetoothLocalDevice) Error(error QBluetoothLocalDevice__Error) {
	defer qt.Recovering("QBluetoothLocalDevice::error")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_Error(ptr.Pointer(), C.int(error))
	}
}

//export callbackQBluetoothLocalDevice_HostModeStateChanged
func callbackQBluetoothLocalDevice_HostModeStateChanged(ptr unsafe.Pointer, ptrName *C.char, state C.int) {
	defer qt.Recovering("callback QBluetoothLocalDevice::hostModeStateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "hostModeStateChanged"); signal != nil {
		signal.(func(QBluetoothLocalDevice__HostMode))(QBluetoothLocalDevice__HostMode(state))
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

func (ptr *QBluetoothLocalDevice) HostModeStateChanged(state QBluetoothLocalDevice__HostMode) {
	defer qt.Recovering("QBluetoothLocalDevice::hostModeStateChanged")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_HostModeStateChanged(ptr.Pointer(), C.int(state))
	}
}

//export callbackQBluetoothLocalDevice_PairingDisplayConfirmation
func callbackQBluetoothLocalDevice_PairingDisplayConfirmation(ptr unsafe.Pointer, ptrName *C.char, address unsafe.Pointer, pin *C.char) {
	defer qt.Recovering("callback QBluetoothLocalDevice::pairingDisplayConfirmation")

	if signal := qt.GetSignal(C.GoString(ptrName), "pairingDisplayConfirmation"); signal != nil {
		signal.(func(*QBluetoothAddress, string))(NewQBluetoothAddressFromPointer(address), C.GoString(pin))
	}

}

func (ptr *QBluetoothLocalDevice) ConnectPairingDisplayConfirmation(f func(address *QBluetoothAddress, pin string)) {
	defer qt.Recovering("connect QBluetoothLocalDevice::pairingDisplayConfirmation")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_ConnectPairingDisplayConfirmation(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "pairingDisplayConfirmation", f)
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectPairingDisplayConfirmation() {
	defer qt.Recovering("disconnect QBluetoothLocalDevice::pairingDisplayConfirmation")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DisconnectPairingDisplayConfirmation(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "pairingDisplayConfirmation")
	}
}

func (ptr *QBluetoothLocalDevice) PairingDisplayConfirmation(address QBluetoothAddress_ITF, pin string) {
	defer qt.Recovering("QBluetoothLocalDevice::pairingDisplayConfirmation")

	if ptr.Pointer() != nil {
		var pinC = C.CString(pin)
		defer C.free(unsafe.Pointer(pinC))
		C.QBluetoothLocalDevice_PairingDisplayConfirmation(ptr.Pointer(), PointerFromQBluetoothAddress(address), pinC)
	}
}

//export callbackQBluetoothLocalDevice_PairingDisplayPinCode
func callbackQBluetoothLocalDevice_PairingDisplayPinCode(ptr unsafe.Pointer, ptrName *C.char, address unsafe.Pointer, pin *C.char) {
	defer qt.Recovering("callback QBluetoothLocalDevice::pairingDisplayPinCode")

	if signal := qt.GetSignal(C.GoString(ptrName), "pairingDisplayPinCode"); signal != nil {
		signal.(func(*QBluetoothAddress, string))(NewQBluetoothAddressFromPointer(address), C.GoString(pin))
	}

}

func (ptr *QBluetoothLocalDevice) ConnectPairingDisplayPinCode(f func(address *QBluetoothAddress, pin string)) {
	defer qt.Recovering("connect QBluetoothLocalDevice::pairingDisplayPinCode")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_ConnectPairingDisplayPinCode(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "pairingDisplayPinCode", f)
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectPairingDisplayPinCode() {
	defer qt.Recovering("disconnect QBluetoothLocalDevice::pairingDisplayPinCode")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DisconnectPairingDisplayPinCode(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "pairingDisplayPinCode")
	}
}

func (ptr *QBluetoothLocalDevice) PairingDisplayPinCode(address QBluetoothAddress_ITF, pin string) {
	defer qt.Recovering("QBluetoothLocalDevice::pairingDisplayPinCode")

	if ptr.Pointer() != nil {
		var pinC = C.CString(pin)
		defer C.free(unsafe.Pointer(pinC))
		C.QBluetoothLocalDevice_PairingDisplayPinCode(ptr.Pointer(), PointerFromQBluetoothAddress(address), pinC)
	}
}

//export callbackQBluetoothLocalDevice_PairingFinished
func callbackQBluetoothLocalDevice_PairingFinished(ptr unsafe.Pointer, ptrName *C.char, address unsafe.Pointer, pairing C.int) {
	defer qt.Recovering("callback QBluetoothLocalDevice::pairingFinished")

	if signal := qt.GetSignal(C.GoString(ptrName), "pairingFinished"); signal != nil {
		signal.(func(*QBluetoothAddress, QBluetoothLocalDevice__Pairing))(NewQBluetoothAddressFromPointer(address), QBluetoothLocalDevice__Pairing(pairing))
	}

}

func (ptr *QBluetoothLocalDevice) ConnectPairingFinished(f func(address *QBluetoothAddress, pairing QBluetoothLocalDevice__Pairing)) {
	defer qt.Recovering("connect QBluetoothLocalDevice::pairingFinished")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_ConnectPairingFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "pairingFinished", f)
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectPairingFinished() {
	defer qt.Recovering("disconnect QBluetoothLocalDevice::pairingFinished")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DisconnectPairingFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "pairingFinished")
	}
}

func (ptr *QBluetoothLocalDevice) PairingFinished(address QBluetoothAddress_ITF, pairing QBluetoothLocalDevice__Pairing) {
	defer qt.Recovering("QBluetoothLocalDevice::pairingFinished")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_PairingFinished(ptr.Pointer(), PointerFromQBluetoothAddress(address), C.int(pairing))
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
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QBluetoothLocalDevice_DestroyQBluetoothLocalDevice(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func NewQBluetoothLocalDevice(parent core.QObject_ITF) *QBluetoothLocalDevice {
	defer qt.Recovering("QBluetoothLocalDevice::QBluetoothLocalDevice")

	return newQBluetoothLocalDeviceFromPointer(C.QBluetoothLocalDevice_NewQBluetoothLocalDevice(core.PointerFromQObject(parent)))
}

func NewQBluetoothLocalDevice2(address QBluetoothAddress_ITF, parent core.QObject_ITF) *QBluetoothLocalDevice {
	defer qt.Recovering("QBluetoothLocalDevice::QBluetoothLocalDevice")

	return newQBluetoothLocalDeviceFromPointer(C.QBluetoothLocalDevice_NewQBluetoothLocalDevice2(PointerFromQBluetoothAddress(address), core.PointerFromQObject(parent)))
}

func (ptr *QBluetoothLocalDevice) Address() *QBluetoothAddress {
	defer qt.Recovering("QBluetoothLocalDevice::address")

	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothAddressFromPointer(C.QBluetoothLocalDevice_Address(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
		return tmpValue
	}
	return nil
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

//export callbackQBluetoothLocalDevice_PairingConfirmation
func callbackQBluetoothLocalDevice_PairingConfirmation(ptr unsafe.Pointer, ptrName *C.char, accept C.int) {
	defer qt.Recovering("callback QBluetoothLocalDevice::pairingConfirmation")

	if signal := qt.GetSignal(C.GoString(ptrName), "pairingConfirmation"); signal != nil {
		signal.(func(bool))(int(accept) != 0)
	}

}

func (ptr *QBluetoothLocalDevice) ConnectPairingConfirmation(f func(accept bool)) {
	defer qt.Recovering("connect QBluetoothLocalDevice::pairingConfirmation")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "pairingConfirmation", f)
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectPairingConfirmation(accept bool) {
	defer qt.Recovering("disconnect QBluetoothLocalDevice::pairingConfirmation")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "pairingConfirmation")
	}
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

//export callbackQBluetoothLocalDevice_TimerEvent
func callbackQBluetoothLocalDevice_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothLocalDevice::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQBluetoothLocalDeviceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
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

//export callbackQBluetoothLocalDevice_ChildEvent
func callbackQBluetoothLocalDevice_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothLocalDevice::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQBluetoothLocalDeviceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
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

//export callbackQBluetoothLocalDevice_ConnectNotify
func callbackQBluetoothLocalDevice_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothLocalDevice::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBluetoothLocalDeviceFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBluetoothLocalDevice) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QBluetoothLocalDevice::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QBluetoothLocalDevice::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QBluetoothLocalDevice) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QBluetoothLocalDevice::connectNotify")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QBluetoothLocalDevice) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QBluetoothLocalDevice::connectNotify")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBluetoothLocalDevice_CustomEvent
func callbackQBluetoothLocalDevice_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothLocalDevice::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQBluetoothLocalDeviceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
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

//export callbackQBluetoothLocalDevice_DeleteLater
func callbackQBluetoothLocalDevice_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QBluetoothLocalDevice::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQBluetoothLocalDeviceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QBluetoothLocalDevice) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QBluetoothLocalDevice::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QBluetoothLocalDevice::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QBluetoothLocalDevice) DeleteLater() {
	defer qt.Recovering("QBluetoothLocalDevice::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QBluetoothLocalDevice_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothLocalDevice) DeleteLaterDefault() {
	defer qt.Recovering("QBluetoothLocalDevice::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QBluetoothLocalDevice_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQBluetoothLocalDevice_DisconnectNotify
func callbackQBluetoothLocalDevice_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothLocalDevice::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBluetoothLocalDeviceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBluetoothLocalDevice) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QBluetoothLocalDevice::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QBluetoothLocalDevice::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QBluetoothLocalDevice::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QBluetoothLocalDevice::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBluetoothLocalDevice_Event
func callbackQBluetoothLocalDevice_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QBluetoothLocalDevice::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQBluetoothLocalDeviceFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QBluetoothLocalDevice) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QBluetoothLocalDevice::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectEvent() {
	defer qt.Recovering("disconnect QBluetoothLocalDevice::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QBluetoothLocalDevice) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QBluetoothLocalDevice::event")

	if ptr.Pointer() != nil {
		return C.QBluetoothLocalDevice_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QBluetoothLocalDevice) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QBluetoothLocalDevice::event")

	if ptr.Pointer() != nil {
		return C.QBluetoothLocalDevice_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQBluetoothLocalDevice_EventFilter
func callbackQBluetoothLocalDevice_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QBluetoothLocalDevice::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQBluetoothLocalDeviceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QBluetoothLocalDevice) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QBluetoothLocalDevice::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QBluetoothLocalDevice::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QBluetoothLocalDevice) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QBluetoothLocalDevice::eventFilter")

	if ptr.Pointer() != nil {
		return C.QBluetoothLocalDevice_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QBluetoothLocalDevice) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QBluetoothLocalDevice::eventFilter")

	if ptr.Pointer() != nil {
		return C.QBluetoothLocalDevice_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQBluetoothLocalDevice_MetaObject
func callbackQBluetoothLocalDevice_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QBluetoothLocalDevice::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQBluetoothLocalDeviceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QBluetoothLocalDevice) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QBluetoothLocalDevice::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QBluetoothLocalDevice::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QBluetoothLocalDevice) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QBluetoothLocalDevice::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QBluetoothLocalDevice_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBluetoothLocalDevice) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QBluetoothLocalDevice::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QBluetoothLocalDevice_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QBluetoothServer::Error
type QBluetoothServer__Error int64

const (
	QBluetoothServer__NoError                       = QBluetoothServer__Error(0)
	QBluetoothServer__UnknownError                  = QBluetoothServer__Error(1)
	QBluetoothServer__PoweredOffError               = QBluetoothServer__Error(2)
	QBluetoothServer__InputOutputError              = QBluetoothServer__Error(3)
	QBluetoothServer__ServiceAlreadyRegisteredError = QBluetoothServer__Error(4)
	QBluetoothServer__UnsupportedProtocolError      = QBluetoothServer__Error(5)
)

type QBluetoothServer struct {
	core.QObject
}

type QBluetoothServer_ITF interface {
	core.QObject_ITF
	QBluetoothServer_PTR() *QBluetoothServer
}

func (p *QBluetoothServer) QBluetoothServer_PTR() *QBluetoothServer {
	return p
}

func (p *QBluetoothServer) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QBluetoothServer) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQBluetoothServer(ptr QBluetoothServer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothServer_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothServerFromPointer(ptr unsafe.Pointer) *QBluetoothServer {
	var n = new(QBluetoothServer)
	n.SetPointer(ptr)
	return n
}

func newQBluetoothServerFromPointer(ptr unsafe.Pointer) *QBluetoothServer {
	var n = NewQBluetoothServerFromPointer(ptr)
	for len(n.ObjectName()) < len("QBluetoothServer_") {
		n.SetObjectName("QBluetoothServer_" + qt.Identifier())
	}
	return n
}

func NewQBluetoothServer(serverType QBluetoothServiceInfo__Protocol, parent core.QObject_ITF) *QBluetoothServer {
	defer qt.Recovering("QBluetoothServer::QBluetoothServer")

	return newQBluetoothServerFromPointer(C.QBluetoothServer_NewQBluetoothServer(C.int(serverType), core.PointerFromQObject(parent)))
}

//export callbackQBluetoothServer_Error2
func callbackQBluetoothServer_Error2(ptr unsafe.Pointer, ptrName *C.char, error C.int) {
	defer qt.Recovering("callback QBluetoothServer::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(QBluetoothServer__Error))(QBluetoothServer__Error(error))
	}

}

func (ptr *QBluetoothServer) ConnectError2(f func(error QBluetoothServer__Error)) {
	defer qt.Recovering("connect QBluetoothServer::error")

	if ptr.Pointer() != nil {
		C.QBluetoothServer_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error2", f)
	}
}

func (ptr *QBluetoothServer) DisconnectError2() {
	defer qt.Recovering("disconnect QBluetoothServer::error")

	if ptr.Pointer() != nil {
		C.QBluetoothServer_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error2")
	}
}

func (ptr *QBluetoothServer) Error2(error QBluetoothServer__Error) {
	defer qt.Recovering("QBluetoothServer::error")

	if ptr.Pointer() != nil {
		C.QBluetoothServer_Error2(ptr.Pointer(), C.int(error))
	}
}

//export callbackQBluetoothServer_NewConnection
func callbackQBluetoothServer_NewConnection(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QBluetoothServer::newConnection")

	if signal := qt.GetSignal(C.GoString(ptrName), "newConnection"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBluetoothServer) ConnectNewConnection(f func()) {
	defer qt.Recovering("connect QBluetoothServer::newConnection")

	if ptr.Pointer() != nil {
		C.QBluetoothServer_ConnectNewConnection(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "newConnection", f)
	}
}

func (ptr *QBluetoothServer) DisconnectNewConnection() {
	defer qt.Recovering("disconnect QBluetoothServer::newConnection")

	if ptr.Pointer() != nil {
		C.QBluetoothServer_DisconnectNewConnection(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "newConnection")
	}
}

func (ptr *QBluetoothServer) NewConnection() {
	defer qt.Recovering("QBluetoothServer::newConnection")

	if ptr.Pointer() != nil {
		C.QBluetoothServer_NewConnection(ptr.Pointer())
	}
}

func (ptr *QBluetoothServer) Error() QBluetoothServer__Error {
	defer qt.Recovering("QBluetoothServer::error")

	if ptr.Pointer() != nil {
		return QBluetoothServer__Error(C.QBluetoothServer_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothServer) IsListening() bool {
	defer qt.Recovering("QBluetoothServer::isListening")

	if ptr.Pointer() != nil {
		return C.QBluetoothServer_IsListening(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothServer) Listen2(uuid QBluetoothUuid_ITF, serviceName string) *QBluetoothServiceInfo {
	defer qt.Recovering("QBluetoothServer::listen")

	if ptr.Pointer() != nil {
		var serviceNameC = C.CString(serviceName)
		defer C.free(unsafe.Pointer(serviceNameC))
		var tmpValue = NewQBluetoothServiceInfoFromPointer(C.QBluetoothServer_Listen2(ptr.Pointer(), PointerFromQBluetoothUuid(uuid), serviceNameC))
		runtime.SetFinalizer(tmpValue, (*QBluetoothServiceInfo).DestroyQBluetoothServiceInfo)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothServer) MaxPendingConnections() int {
	defer qt.Recovering("QBluetoothServer::maxPendingConnections")

	if ptr.Pointer() != nil {
		return int(C.QBluetoothServer_MaxPendingConnections(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothServer) ServerType() QBluetoothServiceInfo__Protocol {
	defer qt.Recovering("QBluetoothServer::serverType")

	if ptr.Pointer() != nil {
		return QBluetoothServiceInfo__Protocol(C.QBluetoothServer_ServerType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothServer) DestroyQBluetoothServer() {
	defer qt.Recovering("QBluetoothServer::~QBluetoothServer")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QBluetoothServer_DestroyQBluetoothServer(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothServer) Close() {
	defer qt.Recovering("QBluetoothServer::close")

	if ptr.Pointer() != nil {
		C.QBluetoothServer_Close(ptr.Pointer())
	}
}

func (ptr *QBluetoothServer) HasPendingConnections() bool {
	defer qt.Recovering("QBluetoothServer::hasPendingConnections")

	if ptr.Pointer() != nil {
		return C.QBluetoothServer_HasPendingConnections(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothServer) NextPendingConnection() *QBluetoothSocket {
	defer qt.Recovering("QBluetoothServer::nextPendingConnection")

	if ptr.Pointer() != nil {
		return NewQBluetoothSocketFromPointer(C.QBluetoothServer_NextPendingConnection(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBluetoothServer) ServerAddress() *QBluetoothAddress {
	defer qt.Recovering("QBluetoothServer::serverAddress")

	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothAddressFromPointer(C.QBluetoothServer_ServerAddress(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothServer) SetMaxPendingConnections(numConnections int) {
	defer qt.Recovering("QBluetoothServer::setMaxPendingConnections")

	if ptr.Pointer() != nil {
		C.QBluetoothServer_SetMaxPendingConnections(ptr.Pointer(), C.int(numConnections))
	}
}

//export callbackQBluetoothServer_TimerEvent
func callbackQBluetoothServer_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothServer::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQBluetoothServerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QBluetoothServer) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QBluetoothServer::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QBluetoothServer) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QBluetoothServer::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QBluetoothServer) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QBluetoothServer::timerEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothServer_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QBluetoothServer) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QBluetoothServer::timerEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothServer_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQBluetoothServer_ChildEvent
func callbackQBluetoothServer_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothServer::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQBluetoothServerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QBluetoothServer) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QBluetoothServer::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QBluetoothServer) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QBluetoothServer::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QBluetoothServer) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QBluetoothServer::childEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothServer_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QBluetoothServer) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QBluetoothServer::childEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothServer_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQBluetoothServer_ConnectNotify
func callbackQBluetoothServer_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothServer::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBluetoothServerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBluetoothServer) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QBluetoothServer::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QBluetoothServer) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QBluetoothServer::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QBluetoothServer) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QBluetoothServer::connectNotify")

	if ptr.Pointer() != nil {
		C.QBluetoothServer_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QBluetoothServer) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QBluetoothServer::connectNotify")

	if ptr.Pointer() != nil {
		C.QBluetoothServer_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBluetoothServer_CustomEvent
func callbackQBluetoothServer_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothServer::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQBluetoothServerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QBluetoothServer) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QBluetoothServer::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QBluetoothServer) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QBluetoothServer::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QBluetoothServer) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QBluetoothServer::customEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothServer_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QBluetoothServer) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QBluetoothServer::customEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothServer_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQBluetoothServer_DeleteLater
func callbackQBluetoothServer_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QBluetoothServer::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQBluetoothServerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QBluetoothServer) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QBluetoothServer::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QBluetoothServer) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QBluetoothServer::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QBluetoothServer) DeleteLater() {
	defer qt.Recovering("QBluetoothServer::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QBluetoothServer_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothServer) DeleteLaterDefault() {
	defer qt.Recovering("QBluetoothServer::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QBluetoothServer_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQBluetoothServer_DisconnectNotify
func callbackQBluetoothServer_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothServer::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBluetoothServerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBluetoothServer) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QBluetoothServer::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QBluetoothServer) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QBluetoothServer::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QBluetoothServer) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QBluetoothServer::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QBluetoothServer_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QBluetoothServer) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QBluetoothServer::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QBluetoothServer_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBluetoothServer_Event
func callbackQBluetoothServer_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QBluetoothServer::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQBluetoothServerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QBluetoothServer) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QBluetoothServer::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QBluetoothServer) DisconnectEvent() {
	defer qt.Recovering("disconnect QBluetoothServer::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QBluetoothServer) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QBluetoothServer::event")

	if ptr.Pointer() != nil {
		return C.QBluetoothServer_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QBluetoothServer) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QBluetoothServer::event")

	if ptr.Pointer() != nil {
		return C.QBluetoothServer_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQBluetoothServer_EventFilter
func callbackQBluetoothServer_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QBluetoothServer::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQBluetoothServerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QBluetoothServer) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QBluetoothServer::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QBluetoothServer) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QBluetoothServer::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QBluetoothServer) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QBluetoothServer::eventFilter")

	if ptr.Pointer() != nil {
		return C.QBluetoothServer_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QBluetoothServer) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QBluetoothServer::eventFilter")

	if ptr.Pointer() != nil {
		return C.QBluetoothServer_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQBluetoothServer_MetaObject
func callbackQBluetoothServer_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QBluetoothServer::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQBluetoothServerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QBluetoothServer) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QBluetoothServer::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QBluetoothServer) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QBluetoothServer::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QBluetoothServer) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QBluetoothServer::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QBluetoothServer_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBluetoothServer) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QBluetoothServer::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QBluetoothServer_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QBluetoothServiceDiscoveryAgent::DiscoveryMode
type QBluetoothServiceDiscoveryAgent__DiscoveryMode int64

const (
	QBluetoothServiceDiscoveryAgent__MinimalDiscovery = QBluetoothServiceDiscoveryAgent__DiscoveryMode(0)
	QBluetoothServiceDiscoveryAgent__FullDiscovery    = QBluetoothServiceDiscoveryAgent__DiscoveryMode(1)
)

//QBluetoothServiceDiscoveryAgent::Error
type QBluetoothServiceDiscoveryAgent__Error int64

const (
	QBluetoothServiceDiscoveryAgent__NoError                      = QBluetoothServiceDiscoveryAgent__Error(QBluetoothDeviceDiscoveryAgent__NoError)
	QBluetoothServiceDiscoveryAgent__InputOutputError             = QBluetoothServiceDiscoveryAgent__Error(QBluetoothDeviceDiscoveryAgent__InputOutputError)
	QBluetoothServiceDiscoveryAgent__PoweredOffError              = QBluetoothServiceDiscoveryAgent__Error(QBluetoothDeviceDiscoveryAgent__PoweredOffError)
	QBluetoothServiceDiscoveryAgent__InvalidBluetoothAdapterError = QBluetoothServiceDiscoveryAgent__Error(QBluetoothDeviceDiscoveryAgent__InvalidBluetoothAdapterError)
	QBluetoothServiceDiscoveryAgent__UnknownError                 = QBluetoothServiceDiscoveryAgent__Error(QBluetoothDeviceDiscoveryAgent__UnknownError)
)

type QBluetoothServiceDiscoveryAgent struct {
	core.QObject
}

type QBluetoothServiceDiscoveryAgent_ITF interface {
	core.QObject_ITF
	QBluetoothServiceDiscoveryAgent_PTR() *QBluetoothServiceDiscoveryAgent
}

func (p *QBluetoothServiceDiscoveryAgent) QBluetoothServiceDiscoveryAgent_PTR() *QBluetoothServiceDiscoveryAgent {
	return p
}

func (p *QBluetoothServiceDiscoveryAgent) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QBluetoothServiceDiscoveryAgent) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQBluetoothServiceDiscoveryAgent(ptr QBluetoothServiceDiscoveryAgent_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothServiceDiscoveryAgent_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothServiceDiscoveryAgentFromPointer(ptr unsafe.Pointer) *QBluetoothServiceDiscoveryAgent {
	var n = new(QBluetoothServiceDiscoveryAgent)
	n.SetPointer(ptr)
	return n
}

func newQBluetoothServiceDiscoveryAgentFromPointer(ptr unsafe.Pointer) *QBluetoothServiceDiscoveryAgent {
	var n = NewQBluetoothServiceDiscoveryAgentFromPointer(ptr)
	for len(n.ObjectName()) < len("QBluetoothServiceDiscoveryAgent_") {
		n.SetObjectName("QBluetoothServiceDiscoveryAgent_" + qt.Identifier())
	}
	return n
}

//export callbackQBluetoothServiceDiscoveryAgent_Canceled
func callbackQBluetoothServiceDiscoveryAgent_Canceled(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QBluetoothServiceDiscoveryAgent::canceled")

	if signal := qt.GetSignal(C.GoString(ptrName), "canceled"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectCanceled(f func()) {
	defer qt.Recovering("connect QBluetoothServiceDiscoveryAgent::canceled")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_ConnectCanceled(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "canceled", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectCanceled() {
	defer qt.Recovering("disconnect QBluetoothServiceDiscoveryAgent::canceled")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_DisconnectCanceled(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "canceled")
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) Canceled() {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::canceled")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_Canceled(ptr.Pointer())
	}
}

//export callbackQBluetoothServiceDiscoveryAgent_Error2
func callbackQBluetoothServiceDiscoveryAgent_Error2(ptr unsafe.Pointer, ptrName *C.char, error C.int) {
	defer qt.Recovering("callback QBluetoothServiceDiscoveryAgent::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(QBluetoothServiceDiscoveryAgent__Error))(QBluetoothServiceDiscoveryAgent__Error(error))
	}

}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectError2(f func(error QBluetoothServiceDiscoveryAgent__Error)) {
	defer qt.Recovering("connect QBluetoothServiceDiscoveryAgent::error")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error2", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectError2() {
	defer qt.Recovering("disconnect QBluetoothServiceDiscoveryAgent::error")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error2")
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) Error2(error QBluetoothServiceDiscoveryAgent__Error) {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::error")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_Error2(ptr.Pointer(), C.int(error))
	}
}

//export callbackQBluetoothServiceDiscoveryAgent_Finished
func callbackQBluetoothServiceDiscoveryAgent_Finished(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QBluetoothServiceDiscoveryAgent::finished")

	if signal := qt.GetSignal(C.GoString(ptrName), "finished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectFinished(f func()) {
	defer qt.Recovering("connect QBluetoothServiceDiscoveryAgent::finished")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectFinished() {
	defer qt.Recovering("disconnect QBluetoothServiceDiscoveryAgent::finished")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) Finished() {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::finished")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_Finished(ptr.Pointer())
	}
}

//export callbackQBluetoothServiceDiscoveryAgent_ServiceDiscovered
func callbackQBluetoothServiceDiscoveryAgent_ServiceDiscovered(ptr unsafe.Pointer, ptrName *C.char, info unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothServiceDiscoveryAgent::serviceDiscovered")

	if signal := qt.GetSignal(C.GoString(ptrName), "serviceDiscovered"); signal != nil {
		signal.(func(*QBluetoothServiceInfo))(NewQBluetoothServiceInfoFromPointer(info))
	}

}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectServiceDiscovered(f func(info *QBluetoothServiceInfo)) {
	defer qt.Recovering("connect QBluetoothServiceDiscoveryAgent::serviceDiscovered")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_ConnectServiceDiscovered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "serviceDiscovered", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectServiceDiscovered() {
	defer qt.Recovering("disconnect QBluetoothServiceDiscoveryAgent::serviceDiscovered")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_DisconnectServiceDiscovered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "serviceDiscovered")
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) ServiceDiscovered(info QBluetoothServiceInfo_ITF) {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::serviceDiscovered")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_ServiceDiscovered(ptr.Pointer(), PointerFromQBluetoothServiceInfo(info))
	}
}

func NewQBluetoothServiceDiscoveryAgent(parent core.QObject_ITF) *QBluetoothServiceDiscoveryAgent {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::QBluetoothServiceDiscoveryAgent")

	return newQBluetoothServiceDiscoveryAgentFromPointer(C.QBluetoothServiceDiscoveryAgent_NewQBluetoothServiceDiscoveryAgent(core.PointerFromQObject(parent)))
}

func NewQBluetoothServiceDiscoveryAgent2(deviceAdapter QBluetoothAddress_ITF, parent core.QObject_ITF) *QBluetoothServiceDiscoveryAgent {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::QBluetoothServiceDiscoveryAgent")

	return newQBluetoothServiceDiscoveryAgentFromPointer(C.QBluetoothServiceDiscoveryAgent_NewQBluetoothServiceDiscoveryAgent2(PointerFromQBluetoothAddress(deviceAdapter), core.PointerFromQObject(parent)))
}

//export callbackQBluetoothServiceDiscoveryAgent_Clear
func callbackQBluetoothServiceDiscoveryAgent_Clear(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QBluetoothServiceDiscoveryAgent::clear")

	if signal := qt.GetSignal(C.GoString(ptrName), "clear"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectClear(f func()) {
	defer qt.Recovering("connect QBluetoothServiceDiscoveryAgent::clear")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "clear", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectClear() {
	defer qt.Recovering("disconnect QBluetoothServiceDiscoveryAgent::clear")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "clear")
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) Clear() {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::clear")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_Clear(ptr.Pointer())
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) Error() QBluetoothServiceDiscoveryAgent__Error {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::error")

	if ptr.Pointer() != nil {
		return QBluetoothServiceDiscoveryAgent__Error(C.QBluetoothServiceDiscoveryAgent_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothServiceDiscoveryAgent) ErrorString() string {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothServiceDiscoveryAgent_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothServiceDiscoveryAgent) IsActive() bool {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::isActive")

	if ptr.Pointer() != nil {
		return C.QBluetoothServiceDiscoveryAgent_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothServiceDiscoveryAgent) RemoteAddress() *QBluetoothAddress {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::remoteAddress")

	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothAddressFromPointer(C.QBluetoothServiceDiscoveryAgent_RemoteAddress(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothServiceDiscoveryAgent) SetRemoteAddress(address QBluetoothAddress_ITF) bool {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::setRemoteAddress")

	if ptr.Pointer() != nil {
		return C.QBluetoothServiceDiscoveryAgent_SetRemoteAddress(ptr.Pointer(), PointerFromQBluetoothAddress(address)) != 0
	}
	return false
}

func (ptr *QBluetoothServiceDiscoveryAgent) SetUuidFilter2(uuid QBluetoothUuid_ITF) {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::setUuidFilter")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_SetUuidFilter2(ptr.Pointer(), PointerFromQBluetoothUuid(uuid))
	}
}

//export callbackQBluetoothServiceDiscoveryAgent_Start
func callbackQBluetoothServiceDiscoveryAgent_Start(ptr unsafe.Pointer, ptrName *C.char, mode C.int) {
	defer qt.Recovering("callback QBluetoothServiceDiscoveryAgent::start")

	if signal := qt.GetSignal(C.GoString(ptrName), "start"); signal != nil {
		signal.(func(QBluetoothServiceDiscoveryAgent__DiscoveryMode))(QBluetoothServiceDiscoveryAgent__DiscoveryMode(mode))
	}

}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectStart(f func(mode QBluetoothServiceDiscoveryAgent__DiscoveryMode)) {
	defer qt.Recovering("connect QBluetoothServiceDiscoveryAgent::start")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "start", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectStart(mode QBluetoothServiceDiscoveryAgent__DiscoveryMode) {
	defer qt.Recovering("disconnect QBluetoothServiceDiscoveryAgent::start")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "start")
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) Start(mode QBluetoothServiceDiscoveryAgent__DiscoveryMode) {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::start")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_Start(ptr.Pointer(), C.int(mode))
	}
}

//export callbackQBluetoothServiceDiscoveryAgent_Stop
func callbackQBluetoothServiceDiscoveryAgent_Stop(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QBluetoothServiceDiscoveryAgent::stop")

	if signal := qt.GetSignal(C.GoString(ptrName), "stop"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectStop(f func()) {
	defer qt.Recovering("connect QBluetoothServiceDiscoveryAgent::stop")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "stop", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectStop() {
	defer qt.Recovering("disconnect QBluetoothServiceDiscoveryAgent::stop")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "stop")
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) Stop() {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::stop")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_Stop(ptr.Pointer())
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DestroyQBluetoothServiceDiscoveryAgent() {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::~QBluetoothServiceDiscoveryAgent")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QBluetoothServiceDiscoveryAgent_DestroyQBluetoothServiceDiscoveryAgent(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQBluetoothServiceDiscoveryAgent_TimerEvent
func callbackQBluetoothServiceDiscoveryAgent_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothServiceDiscoveryAgent::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQBluetoothServiceDiscoveryAgentFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QBluetoothServiceDiscoveryAgent::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QBluetoothServiceDiscoveryAgent::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::timerEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::timerEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQBluetoothServiceDiscoveryAgent_ChildEvent
func callbackQBluetoothServiceDiscoveryAgent_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothServiceDiscoveryAgent::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQBluetoothServiceDiscoveryAgentFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QBluetoothServiceDiscoveryAgent::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QBluetoothServiceDiscoveryAgent::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::childEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::childEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQBluetoothServiceDiscoveryAgent_ConnectNotify
func callbackQBluetoothServiceDiscoveryAgent_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothServiceDiscoveryAgent::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBluetoothServiceDiscoveryAgentFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QBluetoothServiceDiscoveryAgent::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QBluetoothServiceDiscoveryAgent::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::connectNotify")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::connectNotify")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBluetoothServiceDiscoveryAgent_CustomEvent
func callbackQBluetoothServiceDiscoveryAgent_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothServiceDiscoveryAgent::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQBluetoothServiceDiscoveryAgentFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QBluetoothServiceDiscoveryAgent::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QBluetoothServiceDiscoveryAgent::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::customEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::customEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQBluetoothServiceDiscoveryAgent_DeleteLater
func callbackQBluetoothServiceDiscoveryAgent_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QBluetoothServiceDiscoveryAgent::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQBluetoothServiceDiscoveryAgentFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QBluetoothServiceDiscoveryAgent::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QBluetoothServiceDiscoveryAgent::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DeleteLater() {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QBluetoothServiceDiscoveryAgent_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DeleteLaterDefault() {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QBluetoothServiceDiscoveryAgent_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQBluetoothServiceDiscoveryAgent_DisconnectNotify
func callbackQBluetoothServiceDiscoveryAgent_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothServiceDiscoveryAgent::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBluetoothServiceDiscoveryAgentFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QBluetoothServiceDiscoveryAgent::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QBluetoothServiceDiscoveryAgent::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBluetoothServiceDiscoveryAgent_Event
func callbackQBluetoothServiceDiscoveryAgent_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QBluetoothServiceDiscoveryAgent::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQBluetoothServiceDiscoveryAgentFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QBluetoothServiceDiscoveryAgent::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectEvent() {
	defer qt.Recovering("disconnect QBluetoothServiceDiscoveryAgent::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::event")

	if ptr.Pointer() != nil {
		return C.QBluetoothServiceDiscoveryAgent_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QBluetoothServiceDiscoveryAgent) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::event")

	if ptr.Pointer() != nil {
		return C.QBluetoothServiceDiscoveryAgent_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQBluetoothServiceDiscoveryAgent_EventFilter
func callbackQBluetoothServiceDiscoveryAgent_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QBluetoothServiceDiscoveryAgent::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQBluetoothServiceDiscoveryAgentFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QBluetoothServiceDiscoveryAgent::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QBluetoothServiceDiscoveryAgent::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::eventFilter")

	if ptr.Pointer() != nil {
		return C.QBluetoothServiceDiscoveryAgent_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QBluetoothServiceDiscoveryAgent) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::eventFilter")

	if ptr.Pointer() != nil {
		return C.QBluetoothServiceDiscoveryAgent_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQBluetoothServiceDiscoveryAgent_MetaObject
func callbackQBluetoothServiceDiscoveryAgent_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QBluetoothServiceDiscoveryAgent::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQBluetoothServiceDiscoveryAgentFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QBluetoothServiceDiscoveryAgent::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QBluetoothServiceDiscoveryAgent::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QBluetoothServiceDiscoveryAgent_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBluetoothServiceDiscoveryAgent) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QBluetoothServiceDiscoveryAgent::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QBluetoothServiceDiscoveryAgent_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QBluetoothServiceInfo::AttributeId
type QBluetoothServiceInfo__AttributeId int64

var (
	QBluetoothServiceInfo__ServiceRecordHandle              = QBluetoothServiceInfo__AttributeId(0x0000)
	QBluetoothServiceInfo__ServiceClassIds                  = QBluetoothServiceInfo__AttributeId(0x0001)
	QBluetoothServiceInfo__ServiceRecordState               = QBluetoothServiceInfo__AttributeId(0x0002)
	QBluetoothServiceInfo__ServiceId                        = QBluetoothServiceInfo__AttributeId(0x0003)
	QBluetoothServiceInfo__ProtocolDescriptorList           = QBluetoothServiceInfo__AttributeId(0x0004)
	QBluetoothServiceInfo__BrowseGroupList                  = QBluetoothServiceInfo__AttributeId(0x0005)
	QBluetoothServiceInfo__LanguageBaseAttributeIdList      = QBluetoothServiceInfo__AttributeId(0x0006)
	QBluetoothServiceInfo__ServiceInfoTimeToLive            = QBluetoothServiceInfo__AttributeId(0x0007)
	QBluetoothServiceInfo__ServiceAvailability              = QBluetoothServiceInfo__AttributeId(0x0008)
	QBluetoothServiceInfo__BluetoothProfileDescriptorList   = QBluetoothServiceInfo__AttributeId(0x0009)
	QBluetoothServiceInfo__DocumentationUrl                 = QBluetoothServiceInfo__AttributeId(0x000A)
	QBluetoothServiceInfo__ClientExecutableUrl              = QBluetoothServiceInfo__AttributeId(0x000B)
	QBluetoothServiceInfo__IconUrl                          = QBluetoothServiceInfo__AttributeId(0x000C)
	QBluetoothServiceInfo__AdditionalProtocolDescriptorList = QBluetoothServiceInfo__AttributeId(0x000D)
	QBluetoothServiceInfo__PrimaryLanguageBase              = QBluetoothServiceInfo__AttributeId(0x0100)
	QBluetoothServiceInfo__ServiceName                      = QBluetoothServiceInfo__AttributeId(C.QBluetoothServiceInfo_ServiceName_Type())
	QBluetoothServiceInfo__ServiceDescription               = QBluetoothServiceInfo__AttributeId(C.QBluetoothServiceInfo_ServiceDescription_Type())
	QBluetoothServiceInfo__ServiceProvider                  = QBluetoothServiceInfo__AttributeId(C.QBluetoothServiceInfo_ServiceProvider_Type())
)

//QBluetoothServiceInfo::Protocol
type QBluetoothServiceInfo__Protocol int64

const (
	QBluetoothServiceInfo__UnknownProtocol = QBluetoothServiceInfo__Protocol(0)
	QBluetoothServiceInfo__L2capProtocol   = QBluetoothServiceInfo__Protocol(1)
	QBluetoothServiceInfo__RfcommProtocol  = QBluetoothServiceInfo__Protocol(2)
)

type QBluetoothServiceInfo struct {
	ptr unsafe.Pointer
}

type QBluetoothServiceInfo_ITF interface {
	QBluetoothServiceInfo_PTR() *QBluetoothServiceInfo
}

func (p *QBluetoothServiceInfo) QBluetoothServiceInfo_PTR() *QBluetoothServiceInfo {
	return p
}

func (p *QBluetoothServiceInfo) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QBluetoothServiceInfo) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQBluetoothServiceInfo(ptr QBluetoothServiceInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothServiceInfo_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothServiceInfoFromPointer(ptr unsafe.Pointer) *QBluetoothServiceInfo {
	var n = new(QBluetoothServiceInfo)
	n.SetPointer(ptr)
	return n
}

func newQBluetoothServiceInfoFromPointer(ptr unsafe.Pointer) *QBluetoothServiceInfo {
	var n = NewQBluetoothServiceInfoFromPointer(ptr)
	return n
}

func (ptr *QBluetoothServiceInfo) ServiceDescription() string {
	defer qt.Recovering("QBluetoothServiceInfo::serviceDescription")

	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothServiceInfo_ServiceDescription(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothServiceInfo) ServiceName() string {
	defer qt.Recovering("QBluetoothServiceInfo::serviceName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothServiceInfo_ServiceName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothServiceInfo) ServiceProvider() string {
	defer qt.Recovering("QBluetoothServiceInfo::serviceProvider")

	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothServiceInfo_ServiceProvider(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothServiceInfo) ServiceUuid() *QBluetoothUuid {
	defer qt.Recovering("QBluetoothServiceInfo::serviceUuid")

	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothUuidFromPointer(C.QBluetoothServiceInfo_ServiceUuid(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothServiceInfo) SetServiceDescription(description string) {
	defer qt.Recovering("QBluetoothServiceInfo::setServiceDescription")

	if ptr.Pointer() != nil {
		var descriptionC = C.CString(description)
		defer C.free(unsafe.Pointer(descriptionC))
		C.QBluetoothServiceInfo_SetServiceDescription(ptr.Pointer(), descriptionC)
	}
}

func (ptr *QBluetoothServiceInfo) SetServiceName(name string) {
	defer qt.Recovering("QBluetoothServiceInfo::setServiceName")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QBluetoothServiceInfo_SetServiceName(ptr.Pointer(), nameC)
	}
}

func (ptr *QBluetoothServiceInfo) SetServiceProvider(provider string) {
	defer qt.Recovering("QBluetoothServiceInfo::setServiceProvider")

	if ptr.Pointer() != nil {
		var providerC = C.CString(provider)
		defer C.free(unsafe.Pointer(providerC))
		C.QBluetoothServiceInfo_SetServiceProvider(ptr.Pointer(), providerC)
	}
}

func (ptr *QBluetoothServiceInfo) SetServiceUuid(uuid QBluetoothUuid_ITF) {
	defer qt.Recovering("QBluetoothServiceInfo::setServiceUuid")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceInfo_SetServiceUuid(ptr.Pointer(), PointerFromQBluetoothUuid(uuid))
	}
}

func NewQBluetoothServiceInfo() *QBluetoothServiceInfo {
	defer qt.Recovering("QBluetoothServiceInfo::QBluetoothServiceInfo")

	return newQBluetoothServiceInfoFromPointer(C.QBluetoothServiceInfo_NewQBluetoothServiceInfo())
}

func NewQBluetoothServiceInfo2(other QBluetoothServiceInfo_ITF) *QBluetoothServiceInfo {
	defer qt.Recovering("QBluetoothServiceInfo::QBluetoothServiceInfo")

	return newQBluetoothServiceInfoFromPointer(C.QBluetoothServiceInfo_NewQBluetoothServiceInfo2(PointerFromQBluetoothServiceInfo(other)))
}

func (ptr *QBluetoothServiceInfo) Device() *QBluetoothDeviceInfo {
	defer qt.Recovering("QBluetoothServiceInfo::device")

	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothDeviceInfoFromPointer(C.QBluetoothServiceInfo_Device(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothDeviceInfo).DestroyQBluetoothDeviceInfo)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothServiceInfo) IsComplete() bool {
	defer qt.Recovering("QBluetoothServiceInfo::isComplete")

	if ptr.Pointer() != nil {
		return C.QBluetoothServiceInfo_IsComplete(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothServiceInfo) IsRegistered() bool {
	defer qt.Recovering("QBluetoothServiceInfo::isRegistered")

	if ptr.Pointer() != nil {
		return C.QBluetoothServiceInfo_IsRegistered(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothServiceInfo) IsValid() bool {
	defer qt.Recovering("QBluetoothServiceInfo::isValid")

	if ptr.Pointer() != nil {
		return C.QBluetoothServiceInfo_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothServiceInfo) ProtocolServiceMultiplexer() int {
	defer qt.Recovering("QBluetoothServiceInfo::protocolServiceMultiplexer")

	if ptr.Pointer() != nil {
		return int(C.QBluetoothServiceInfo_ProtocolServiceMultiplexer(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothServiceInfo) RegisterService(localAdapter QBluetoothAddress_ITF) bool {
	defer qt.Recovering("QBluetoothServiceInfo::registerService")

	if ptr.Pointer() != nil {
		return C.QBluetoothServiceInfo_RegisterService(ptr.Pointer(), PointerFromQBluetoothAddress(localAdapter)) != 0
	}
	return false
}

func (ptr *QBluetoothServiceInfo) ServerChannel() int {
	defer qt.Recovering("QBluetoothServiceInfo::serverChannel")

	if ptr.Pointer() != nil {
		return int(C.QBluetoothServiceInfo_ServerChannel(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothServiceInfo) SetDevice(device QBluetoothDeviceInfo_ITF) {
	defer qt.Recovering("QBluetoothServiceInfo::setDevice")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceInfo_SetDevice(ptr.Pointer(), PointerFromQBluetoothDeviceInfo(device))
	}
}

func (ptr *QBluetoothServiceInfo) SocketProtocol() QBluetoothServiceInfo__Protocol {
	defer qt.Recovering("QBluetoothServiceInfo::socketProtocol")

	if ptr.Pointer() != nil {
		return QBluetoothServiceInfo__Protocol(C.QBluetoothServiceInfo_SocketProtocol(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothServiceInfo) UnregisterService() bool {
	defer qt.Recovering("QBluetoothServiceInfo::unregisterService")

	if ptr.Pointer() != nil {
		return C.QBluetoothServiceInfo_UnregisterService(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothServiceInfo) DestroyQBluetoothServiceInfo() {
	defer qt.Recovering("QBluetoothServiceInfo::~QBluetoothServiceInfo")

	if ptr.Pointer() != nil {
		C.QBluetoothServiceInfo_DestroyQBluetoothServiceInfo(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//QBluetoothSocket::SocketError
type QBluetoothSocket__SocketError int64

const (
	QBluetoothSocket__NoSocketError            = QBluetoothSocket__SocketError(-2)
	QBluetoothSocket__UnknownSocketError       = QBluetoothSocket__SocketError(network.QAbstractSocket__UnknownSocketError)
	QBluetoothSocket__HostNotFoundError        = QBluetoothSocket__SocketError(network.QAbstractSocket__HostNotFoundError)
	QBluetoothSocket__ServiceNotFoundError     = QBluetoothSocket__SocketError(network.QAbstractSocket__SocketAddressNotAvailableError)
	QBluetoothSocket__NetworkError             = QBluetoothSocket__SocketError(network.QAbstractSocket__NetworkError)
	QBluetoothSocket__UnsupportedProtocolError = QBluetoothSocket__SocketError(8)
	QBluetoothSocket__OperationError           = QBluetoothSocket__SocketError(network.QAbstractSocket__OperationError)
)

//QBluetoothSocket::SocketState
type QBluetoothSocket__SocketState int64

const (
	QBluetoothSocket__UnconnectedState   = QBluetoothSocket__SocketState(network.QAbstractSocket__UnconnectedState)
	QBluetoothSocket__ServiceLookupState = QBluetoothSocket__SocketState(network.QAbstractSocket__HostLookupState)
	QBluetoothSocket__ConnectingState    = QBluetoothSocket__SocketState(network.QAbstractSocket__ConnectingState)
	QBluetoothSocket__ConnectedState     = QBluetoothSocket__SocketState(network.QAbstractSocket__ConnectedState)
	QBluetoothSocket__BoundState         = QBluetoothSocket__SocketState(network.QAbstractSocket__BoundState)
	QBluetoothSocket__ClosingState       = QBluetoothSocket__SocketState(network.QAbstractSocket__ClosingState)
	QBluetoothSocket__ListeningState     = QBluetoothSocket__SocketState(network.QAbstractSocket__ListeningState)
)

type QBluetoothSocket struct {
	core.QIODevice
}

type QBluetoothSocket_ITF interface {
	core.QIODevice_ITF
	QBluetoothSocket_PTR() *QBluetoothSocket
}

func (p *QBluetoothSocket) QBluetoothSocket_PTR() *QBluetoothSocket {
	return p
}

func (p *QBluetoothSocket) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QIODevice_PTR().Pointer()
	}
	return nil
}

func (p *QBluetoothSocket) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QIODevice_PTR().SetPointer(ptr)
	}
}

func PointerFromQBluetoothSocket(ptr QBluetoothSocket_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothSocket_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothSocketFromPointer(ptr unsafe.Pointer) *QBluetoothSocket {
	var n = new(QBluetoothSocket)
	n.SetPointer(ptr)
	return n
}

func newQBluetoothSocketFromPointer(ptr unsafe.Pointer) *QBluetoothSocket {
	var n = NewQBluetoothSocketFromPointer(ptr)
	for len(n.ObjectName()) < len("QBluetoothSocket_") {
		n.SetObjectName("QBluetoothSocket_" + qt.Identifier())
	}
	return n
}

//export callbackQBluetoothSocket_Connected
func callbackQBluetoothSocket_Connected(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QBluetoothSocket::connected")

	if signal := qt.GetSignal(C.GoString(ptrName), "connected"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBluetoothSocket) ConnectConnected(f func()) {
	defer qt.Recovering("connect QBluetoothSocket::connected")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_ConnectConnected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "connected", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectConnected() {
	defer qt.Recovering("disconnect QBluetoothSocket::connected")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DisconnectConnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "connected")
	}
}

func (ptr *QBluetoothSocket) Connected() {
	defer qt.Recovering("QBluetoothSocket::connected")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_Connected(ptr.Pointer())
	}
}

//export callbackQBluetoothSocket_Disconnected
func callbackQBluetoothSocket_Disconnected(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QBluetoothSocket::disconnected")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnected"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBluetoothSocket) ConnectDisconnected(f func()) {
	defer qt.Recovering("connect QBluetoothSocket::disconnected")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_ConnectDisconnected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "disconnected", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectDisconnected() {
	defer qt.Recovering("disconnect QBluetoothSocket::disconnected")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DisconnectDisconnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "disconnected")
	}
}

func (ptr *QBluetoothSocket) Disconnected() {
	defer qt.Recovering("QBluetoothSocket::disconnected")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_Disconnected(ptr.Pointer())
	}
}

//export callbackQBluetoothSocket_Error2
func callbackQBluetoothSocket_Error2(ptr unsafe.Pointer, ptrName *C.char, error C.int) {
	defer qt.Recovering("callback QBluetoothSocket::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(QBluetoothSocket__SocketError))(QBluetoothSocket__SocketError(error))
	}

}

func (ptr *QBluetoothSocket) ConnectError2(f func(error QBluetoothSocket__SocketError)) {
	defer qt.Recovering("connect QBluetoothSocket::error")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error2", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectError2() {
	defer qt.Recovering("disconnect QBluetoothSocket::error")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error2")
	}
}

func (ptr *QBluetoothSocket) Error2(error QBluetoothSocket__SocketError) {
	defer qt.Recovering("QBluetoothSocket::error")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_Error2(ptr.Pointer(), C.int(error))
	}
}

//export callbackQBluetoothSocket_StateChanged
func callbackQBluetoothSocket_StateChanged(ptr unsafe.Pointer, ptrName *C.char, state C.int) {
	defer qt.Recovering("callback QBluetoothSocket::stateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stateChanged"); signal != nil {
		signal.(func(QBluetoothSocket__SocketState))(QBluetoothSocket__SocketState(state))
	}

}

func (ptr *QBluetoothSocket) ConnectStateChanged(f func(state QBluetoothSocket__SocketState)) {
	defer qt.Recovering("connect QBluetoothSocket::stateChanged")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectStateChanged() {
	defer qt.Recovering("disconnect QBluetoothSocket::stateChanged")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

func (ptr *QBluetoothSocket) StateChanged(state QBluetoothSocket__SocketState) {
	defer qt.Recovering("QBluetoothSocket::stateChanged")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_StateChanged(ptr.Pointer(), C.int(state))
	}
}

func NewQBluetoothSocket(socketType QBluetoothServiceInfo__Protocol, parent core.QObject_ITF) *QBluetoothSocket {
	defer qt.Recovering("QBluetoothSocket::QBluetoothSocket")

	return newQBluetoothSocketFromPointer(C.QBluetoothSocket_NewQBluetoothSocket(C.int(socketType), core.PointerFromQObject(parent)))
}

func NewQBluetoothSocket2(parent core.QObject_ITF) *QBluetoothSocket {
	defer qt.Recovering("QBluetoothSocket::QBluetoothSocket")

	return newQBluetoothSocketFromPointer(C.QBluetoothSocket_NewQBluetoothSocket2(core.PointerFromQObject(parent)))
}

func (ptr *QBluetoothSocket) Abort() {
	defer qt.Recovering("QBluetoothSocket::abort")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_Abort(ptr.Pointer())
	}
}

//export callbackQBluetoothSocket_BytesAvailable
func callbackQBluetoothSocket_BytesAvailable(ptr unsafe.Pointer, ptrName *C.char) C.longlong {
	defer qt.Recovering("callback QBluetoothSocket::bytesAvailable")

	if signal := qt.GetSignal(C.GoString(ptrName), "bytesAvailable"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQBluetoothSocketFromPointer(ptr).BytesAvailableDefault())
}

func (ptr *QBluetoothSocket) ConnectBytesAvailable(f func() int64) {
	defer qt.Recovering("connect QBluetoothSocket::bytesAvailable")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "bytesAvailable", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectBytesAvailable() {
	defer qt.Recovering("disconnect QBluetoothSocket::bytesAvailable")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "bytesAvailable")
	}
}

func (ptr *QBluetoothSocket) BytesAvailable() int64 {
	defer qt.Recovering("QBluetoothSocket::bytesAvailable")

	if ptr.Pointer() != nil {
		return int64(C.QBluetoothSocket_BytesAvailable(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothSocket) BytesAvailableDefault() int64 {
	defer qt.Recovering("QBluetoothSocket::bytesAvailable")

	if ptr.Pointer() != nil {
		return int64(C.QBluetoothSocket_BytesAvailableDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQBluetoothSocket_BytesToWrite
func callbackQBluetoothSocket_BytesToWrite(ptr unsafe.Pointer, ptrName *C.char) C.longlong {
	defer qt.Recovering("callback QBluetoothSocket::bytesToWrite")

	if signal := qt.GetSignal(C.GoString(ptrName), "bytesToWrite"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQBluetoothSocketFromPointer(ptr).BytesToWriteDefault())
}

func (ptr *QBluetoothSocket) ConnectBytesToWrite(f func() int64) {
	defer qt.Recovering("connect QBluetoothSocket::bytesToWrite")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "bytesToWrite", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectBytesToWrite() {
	defer qt.Recovering("disconnect QBluetoothSocket::bytesToWrite")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "bytesToWrite")
	}
}

func (ptr *QBluetoothSocket) BytesToWrite() int64 {
	defer qt.Recovering("QBluetoothSocket::bytesToWrite")

	if ptr.Pointer() != nil {
		return int64(C.QBluetoothSocket_BytesToWrite(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothSocket) BytesToWriteDefault() int64 {
	defer qt.Recovering("QBluetoothSocket::bytesToWrite")

	if ptr.Pointer() != nil {
		return int64(C.QBluetoothSocket_BytesToWriteDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQBluetoothSocket_CanReadLine
func callbackQBluetoothSocket_CanReadLine(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QBluetoothSocket::canReadLine")

	if signal := qt.GetSignal(C.GoString(ptrName), "canReadLine"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewQBluetoothSocketFromPointer(ptr).CanReadLineDefault()))
}

func (ptr *QBluetoothSocket) ConnectCanReadLine(f func() bool) {
	defer qt.Recovering("connect QBluetoothSocket::canReadLine")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "canReadLine", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectCanReadLine() {
	defer qt.Recovering("disconnect QBluetoothSocket::canReadLine")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "canReadLine")
	}
}

func (ptr *QBluetoothSocket) CanReadLine() bool {
	defer qt.Recovering("QBluetoothSocket::canReadLine")

	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_CanReadLine(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothSocket) CanReadLineDefault() bool {
	defer qt.Recovering("QBluetoothSocket::canReadLine")

	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_CanReadLineDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQBluetoothSocket_Close
func callbackQBluetoothSocket_Close(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QBluetoothSocket::close")

	if signal := qt.GetSignal(C.GoString(ptrName), "close"); signal != nil {
		signal.(func())()
	} else {
		NewQBluetoothSocketFromPointer(ptr).CloseDefault()
	}
}

func (ptr *QBluetoothSocket) ConnectClose(f func()) {
	defer qt.Recovering("connect QBluetoothSocket::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "close", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectClose() {
	defer qt.Recovering("disconnect QBluetoothSocket::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "close")
	}
}

func (ptr *QBluetoothSocket) Close() {
	defer qt.Recovering("QBluetoothSocket::close")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_Close(ptr.Pointer())
	}
}

func (ptr *QBluetoothSocket) CloseDefault() {
	defer qt.Recovering("QBluetoothSocket::close")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_CloseDefault(ptr.Pointer())
	}
}

func (ptr *QBluetoothSocket) ConnectToService2(address QBluetoothAddress_ITF, uuid QBluetoothUuid_ITF, openMode core.QIODevice__OpenModeFlag) {
	defer qt.Recovering("QBluetoothSocket::connectToService")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_ConnectToService2(ptr.Pointer(), PointerFromQBluetoothAddress(address), PointerFromQBluetoothUuid(uuid), C.int(openMode))
	}
}

func (ptr *QBluetoothSocket) ConnectToService(service QBluetoothServiceInfo_ITF, openMode core.QIODevice__OpenModeFlag) {
	defer qt.Recovering("QBluetoothSocket::connectToService")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_ConnectToService(ptr.Pointer(), PointerFromQBluetoothServiceInfo(service), C.int(openMode))
	}
}

func (ptr *QBluetoothSocket) DisconnectFromService() {
	defer qt.Recovering("QBluetoothSocket::disconnectFromService")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DisconnectFromService(ptr.Pointer())
	}
}

func (ptr *QBluetoothSocket) DoDeviceDiscovery(service QBluetoothServiceInfo_ITF, openMode core.QIODevice__OpenModeFlag) {
	defer qt.Recovering("QBluetoothSocket::doDeviceDiscovery")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DoDeviceDiscovery(ptr.Pointer(), PointerFromQBluetoothServiceInfo(service), C.int(openMode))
	}
}

func (ptr *QBluetoothSocket) Error() QBluetoothSocket__SocketError {
	defer qt.Recovering("QBluetoothSocket::error")

	if ptr.Pointer() != nil {
		return QBluetoothSocket__SocketError(C.QBluetoothSocket_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothSocket) ErrorString() string {
	defer qt.Recovering("QBluetoothSocket::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothSocket_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothSocket) IsSequential() bool {
	defer qt.Recovering("QBluetoothSocket::isSequential")

	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_IsSequential(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothSocket) LocalAddress() *QBluetoothAddress {
	defer qt.Recovering("QBluetoothSocket::localAddress")

	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothAddressFromPointer(C.QBluetoothSocket_LocalAddress(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothSocket) LocalName() string {
	defer qt.Recovering("QBluetoothSocket::localName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothSocket_LocalName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothSocket) PeerAddress() *QBluetoothAddress {
	defer qt.Recovering("QBluetoothSocket::peerAddress")

	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothAddressFromPointer(C.QBluetoothSocket_PeerAddress(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothSocket) PeerName() string {
	defer qt.Recovering("QBluetoothSocket::peerName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothSocket_PeerName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothSocket) SetSocketDescriptor(socketDescriptor int, socketType QBluetoothServiceInfo__Protocol, socketState QBluetoothSocket__SocketState, openMode core.QIODevice__OpenModeFlag) bool {
	defer qt.Recovering("QBluetoothSocket::setSocketDescriptor")

	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_SetSocketDescriptor(ptr.Pointer(), C.int(socketDescriptor), C.int(socketType), C.int(socketState), C.int(openMode)) != 0
	}
	return false
}

func (ptr *QBluetoothSocket) SetSocketError(error_ QBluetoothSocket__SocketError) {
	defer qt.Recovering("QBluetoothSocket::setSocketError")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_SetSocketError(ptr.Pointer(), C.int(error_))
	}
}

func (ptr *QBluetoothSocket) SetSocketState(state QBluetoothSocket__SocketState) {
	defer qt.Recovering("QBluetoothSocket::setSocketState")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_SetSocketState(ptr.Pointer(), C.int(state))
	}
}

func (ptr *QBluetoothSocket) SocketDescriptor() int {
	defer qt.Recovering("QBluetoothSocket::socketDescriptor")

	if ptr.Pointer() != nil {
		return int(C.QBluetoothSocket_SocketDescriptor(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothSocket) SocketType() QBluetoothServiceInfo__Protocol {
	defer qt.Recovering("QBluetoothSocket::socketType")

	if ptr.Pointer() != nil {
		return QBluetoothServiceInfo__Protocol(C.QBluetoothSocket_SocketType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothSocket) State() QBluetoothSocket__SocketState {
	defer qt.Recovering("QBluetoothSocket::state")

	if ptr.Pointer() != nil {
		return QBluetoothSocket__SocketState(C.QBluetoothSocket_State(ptr.Pointer()))
	}
	return 0
}

//export callbackQBluetoothSocket_WriteData
func callbackQBluetoothSocket_WriteData(ptr unsafe.Pointer, ptrName *C.char, data *C.char, maxSize C.longlong) C.longlong {
	defer qt.Recovering("callback QBluetoothSocket::writeData")

	if signal := qt.GetSignal(C.GoString(ptrName), "writeData"); signal != nil {
		return C.longlong(signal.(func(string, int64) int64)(C.GoString(data), int64(maxSize)))
	}

	return C.longlong(NewQBluetoothSocketFromPointer(ptr).WriteDataDefault(C.GoString(data), int64(maxSize)))
}

func (ptr *QBluetoothSocket) ConnectWriteData(f func(data string, maxSize int64) int64) {
	defer qt.Recovering("connect QBluetoothSocket::writeData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "writeData", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectWriteData() {
	defer qt.Recovering("disconnect QBluetoothSocket::writeData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "writeData")
	}
}

func (ptr *QBluetoothSocket) WriteData(data string, maxSize int64) int64 {
	defer qt.Recovering("QBluetoothSocket::writeData")

	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		return int64(C.QBluetoothSocket_WriteData(ptr.Pointer(), dataC, C.longlong(maxSize)))
	}
	return 0
}

func (ptr *QBluetoothSocket) WriteDataDefault(data string, maxSize int64) int64 {
	defer qt.Recovering("QBluetoothSocket::writeData")

	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		return int64(C.QBluetoothSocket_WriteDataDefault(ptr.Pointer(), dataC, C.longlong(maxSize)))
	}
	return 0
}

func (ptr *QBluetoothSocket) DestroyQBluetoothSocket() {
	defer qt.Recovering("QBluetoothSocket::~QBluetoothSocket")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QBluetoothSocket_DestroyQBluetoothSocket(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQBluetoothSocket_AtEnd
func callbackQBluetoothSocket_AtEnd(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QBluetoothSocket::atEnd")

	if signal := qt.GetSignal(C.GoString(ptrName), "atEnd"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewQBluetoothSocketFromPointer(ptr).AtEndDefault()))
}

func (ptr *QBluetoothSocket) ConnectAtEnd(f func() bool) {
	defer qt.Recovering("connect QBluetoothSocket::atEnd")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "atEnd", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectAtEnd() {
	defer qt.Recovering("disconnect QBluetoothSocket::atEnd")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "atEnd")
	}
}

func (ptr *QBluetoothSocket) AtEnd() bool {
	defer qt.Recovering("QBluetoothSocket::atEnd")

	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_AtEnd(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothSocket) AtEndDefault() bool {
	defer qt.Recovering("QBluetoothSocket::atEnd")

	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_AtEndDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQBluetoothSocket_Open
func callbackQBluetoothSocket_Open(ptr unsafe.Pointer, ptrName *C.char, mode C.int) C.int {
	defer qt.Recovering("callback QBluetoothSocket::open")

	if signal := qt.GetSignal(C.GoString(ptrName), "open"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(core.QIODevice__OpenModeFlag) bool)(core.QIODevice__OpenModeFlag(mode))))
	}

	return C.int(qt.GoBoolToInt(NewQBluetoothSocketFromPointer(ptr).OpenDefault(core.QIODevice__OpenModeFlag(mode))))
}

func (ptr *QBluetoothSocket) ConnectOpen(f func(mode core.QIODevice__OpenModeFlag) bool) {
	defer qt.Recovering("connect QBluetoothSocket::open")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "open", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectOpen() {
	defer qt.Recovering("disconnect QBluetoothSocket::open")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "open")
	}
}

func (ptr *QBluetoothSocket) Open(mode core.QIODevice__OpenModeFlag) bool {
	defer qt.Recovering("QBluetoothSocket::open")

	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_Open(ptr.Pointer(), C.int(mode)) != 0
	}
	return false
}

func (ptr *QBluetoothSocket) OpenDefault(mode core.QIODevice__OpenModeFlag) bool {
	defer qt.Recovering("QBluetoothSocket::open")

	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_OpenDefault(ptr.Pointer(), C.int(mode)) != 0
	}
	return false
}

//export callbackQBluetoothSocket_Pos
func callbackQBluetoothSocket_Pos(ptr unsafe.Pointer, ptrName *C.char) C.longlong {
	defer qt.Recovering("callback QBluetoothSocket::pos")

	if signal := qt.GetSignal(C.GoString(ptrName), "pos"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQBluetoothSocketFromPointer(ptr).PosDefault())
}

func (ptr *QBluetoothSocket) ConnectPos(f func() int64) {
	defer qt.Recovering("connect QBluetoothSocket::pos")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "pos", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectPos() {
	defer qt.Recovering("disconnect QBluetoothSocket::pos")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "pos")
	}
}

func (ptr *QBluetoothSocket) Pos() int64 {
	defer qt.Recovering("QBluetoothSocket::pos")

	if ptr.Pointer() != nil {
		return int64(C.QBluetoothSocket_Pos(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothSocket) PosDefault() int64 {
	defer qt.Recovering("QBluetoothSocket::pos")

	if ptr.Pointer() != nil {
		return int64(C.QBluetoothSocket_PosDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQBluetoothSocket_ReadLineData
func callbackQBluetoothSocket_ReadLineData(ptr unsafe.Pointer, ptrName *C.char, data *C.char, maxSize C.longlong) C.longlong {
	defer qt.Recovering("callback QBluetoothSocket::readLineData")

	if signal := qt.GetSignal(C.GoString(ptrName), "readLineData"); signal != nil {
		return C.longlong(signal.(func(string, int64) int64)(C.GoString(data), int64(maxSize)))
	}

	return C.longlong(NewQBluetoothSocketFromPointer(ptr).ReadLineDataDefault(C.GoString(data), int64(maxSize)))
}

func (ptr *QBluetoothSocket) ConnectReadLineData(f func(data string, maxSize int64) int64) {
	defer qt.Recovering("connect QBluetoothSocket::readLineData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "readLineData", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectReadLineData() {
	defer qt.Recovering("disconnect QBluetoothSocket::readLineData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "readLineData")
	}
}

func (ptr *QBluetoothSocket) ReadLineData(data string, maxSize int64) int64 {
	defer qt.Recovering("QBluetoothSocket::readLineData")

	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		return int64(C.QBluetoothSocket_ReadLineData(ptr.Pointer(), dataC, C.longlong(maxSize)))
	}
	return 0
}

func (ptr *QBluetoothSocket) ReadLineDataDefault(data string, maxSize int64) int64 {
	defer qt.Recovering("QBluetoothSocket::readLineData")

	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		return int64(C.QBluetoothSocket_ReadLineDataDefault(ptr.Pointer(), dataC, C.longlong(maxSize)))
	}
	return 0
}

//export callbackQBluetoothSocket_Reset
func callbackQBluetoothSocket_Reset(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QBluetoothSocket::reset")

	if signal := qt.GetSignal(C.GoString(ptrName), "reset"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewQBluetoothSocketFromPointer(ptr).ResetDefault()))
}

func (ptr *QBluetoothSocket) ConnectReset(f func() bool) {
	defer qt.Recovering("connect QBluetoothSocket::reset")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "reset", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectReset() {
	defer qt.Recovering("disconnect QBluetoothSocket::reset")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "reset")
	}
}

func (ptr *QBluetoothSocket) Reset() bool {
	defer qt.Recovering("QBluetoothSocket::reset")

	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_Reset(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothSocket) ResetDefault() bool {
	defer qt.Recovering("QBluetoothSocket::reset")

	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_ResetDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQBluetoothSocket_Seek
func callbackQBluetoothSocket_Seek(ptr unsafe.Pointer, ptrName *C.char, pos C.longlong) C.int {
	defer qt.Recovering("callback QBluetoothSocket::seek")

	if signal := qt.GetSignal(C.GoString(ptrName), "seek"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(int64) bool)(int64(pos))))
	}

	return C.int(qt.GoBoolToInt(NewQBluetoothSocketFromPointer(ptr).SeekDefault(int64(pos))))
}

func (ptr *QBluetoothSocket) ConnectSeek(f func(pos int64) bool) {
	defer qt.Recovering("connect QBluetoothSocket::seek")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "seek", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectSeek() {
	defer qt.Recovering("disconnect QBluetoothSocket::seek")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "seek")
	}
}

func (ptr *QBluetoothSocket) Seek(pos int64) bool {
	defer qt.Recovering("QBluetoothSocket::seek")

	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_Seek(ptr.Pointer(), C.longlong(pos)) != 0
	}
	return false
}

func (ptr *QBluetoothSocket) SeekDefault(pos int64) bool {
	defer qt.Recovering("QBluetoothSocket::seek")

	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_SeekDefault(ptr.Pointer(), C.longlong(pos)) != 0
	}
	return false
}

//export callbackQBluetoothSocket_Size
func callbackQBluetoothSocket_Size(ptr unsafe.Pointer, ptrName *C.char) C.longlong {
	defer qt.Recovering("callback QBluetoothSocket::size")

	if signal := qt.GetSignal(C.GoString(ptrName), "size"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQBluetoothSocketFromPointer(ptr).SizeDefault())
}

func (ptr *QBluetoothSocket) ConnectSize(f func() int64) {
	defer qt.Recovering("connect QBluetoothSocket::size")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "size", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectSize() {
	defer qt.Recovering("disconnect QBluetoothSocket::size")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "size")
	}
}

func (ptr *QBluetoothSocket) Size() int64 {
	defer qt.Recovering("QBluetoothSocket::size")

	if ptr.Pointer() != nil {
		return int64(C.QBluetoothSocket_Size(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothSocket) SizeDefault() int64 {
	defer qt.Recovering("QBluetoothSocket::size")

	if ptr.Pointer() != nil {
		return int64(C.QBluetoothSocket_SizeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQBluetoothSocket_WaitForBytesWritten
func callbackQBluetoothSocket_WaitForBytesWritten(ptr unsafe.Pointer, ptrName *C.char, msecs C.int) C.int {
	defer qt.Recovering("callback QBluetoothSocket::waitForBytesWritten")

	if signal := qt.GetSignal(C.GoString(ptrName), "waitForBytesWritten"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(int) bool)(int(msecs))))
	}

	return C.int(qt.GoBoolToInt(NewQBluetoothSocketFromPointer(ptr).WaitForBytesWrittenDefault(int(msecs))))
}

func (ptr *QBluetoothSocket) ConnectWaitForBytesWritten(f func(msecs int) bool) {
	defer qt.Recovering("connect QBluetoothSocket::waitForBytesWritten")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "waitForBytesWritten", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectWaitForBytesWritten() {
	defer qt.Recovering("disconnect QBluetoothSocket::waitForBytesWritten")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "waitForBytesWritten")
	}
}

func (ptr *QBluetoothSocket) WaitForBytesWritten(msecs int) bool {
	defer qt.Recovering("QBluetoothSocket::waitForBytesWritten")

	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_WaitForBytesWritten(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QBluetoothSocket) WaitForBytesWrittenDefault(msecs int) bool {
	defer qt.Recovering("QBluetoothSocket::waitForBytesWritten")

	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_WaitForBytesWrittenDefault(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

//export callbackQBluetoothSocket_WaitForReadyRead
func callbackQBluetoothSocket_WaitForReadyRead(ptr unsafe.Pointer, ptrName *C.char, msecs C.int) C.int {
	defer qt.Recovering("callback QBluetoothSocket::waitForReadyRead")

	if signal := qt.GetSignal(C.GoString(ptrName), "waitForReadyRead"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(int) bool)(int(msecs))))
	}

	return C.int(qt.GoBoolToInt(NewQBluetoothSocketFromPointer(ptr).WaitForReadyReadDefault(int(msecs))))
}

func (ptr *QBluetoothSocket) ConnectWaitForReadyRead(f func(msecs int) bool) {
	defer qt.Recovering("connect QBluetoothSocket::waitForReadyRead")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "waitForReadyRead", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectWaitForReadyRead() {
	defer qt.Recovering("disconnect QBluetoothSocket::waitForReadyRead")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "waitForReadyRead")
	}
}

func (ptr *QBluetoothSocket) WaitForReadyRead(msecs int) bool {
	defer qt.Recovering("QBluetoothSocket::waitForReadyRead")

	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_WaitForReadyRead(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

func (ptr *QBluetoothSocket) WaitForReadyReadDefault(msecs int) bool {
	defer qt.Recovering("QBluetoothSocket::waitForReadyRead")

	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_WaitForReadyReadDefault(ptr.Pointer(), C.int(msecs)) != 0
	}
	return false
}

//export callbackQBluetoothSocket_TimerEvent
func callbackQBluetoothSocket_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothSocket::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQBluetoothSocketFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QBluetoothSocket) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QBluetoothSocket::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QBluetoothSocket::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QBluetoothSocket) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QBluetoothSocket::timerEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QBluetoothSocket) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QBluetoothSocket::timerEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQBluetoothSocket_ChildEvent
func callbackQBluetoothSocket_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothSocket::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQBluetoothSocketFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QBluetoothSocket) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QBluetoothSocket::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QBluetoothSocket::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QBluetoothSocket) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QBluetoothSocket::childEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QBluetoothSocket) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QBluetoothSocket::childEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQBluetoothSocket_ConnectNotify
func callbackQBluetoothSocket_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothSocket::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBluetoothSocketFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBluetoothSocket) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QBluetoothSocket::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QBluetoothSocket::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QBluetoothSocket) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QBluetoothSocket::connectNotify")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QBluetoothSocket) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QBluetoothSocket::connectNotify")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBluetoothSocket_CustomEvent
func callbackQBluetoothSocket_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothSocket::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQBluetoothSocketFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QBluetoothSocket) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QBluetoothSocket::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QBluetoothSocket::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QBluetoothSocket) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QBluetoothSocket::customEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QBluetoothSocket) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QBluetoothSocket::customEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQBluetoothSocket_DeleteLater
func callbackQBluetoothSocket_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QBluetoothSocket::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQBluetoothSocketFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QBluetoothSocket) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QBluetoothSocket::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QBluetoothSocket::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QBluetoothSocket) DeleteLater() {
	defer qt.Recovering("QBluetoothSocket::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QBluetoothSocket_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothSocket) DeleteLaterDefault() {
	defer qt.Recovering("QBluetoothSocket::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QBluetoothSocket_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQBluetoothSocket_DisconnectNotify
func callbackQBluetoothSocket_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothSocket::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBluetoothSocketFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBluetoothSocket) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QBluetoothSocket::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QBluetoothSocket::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QBluetoothSocket) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QBluetoothSocket::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QBluetoothSocket) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QBluetoothSocket::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBluetoothSocket_Event
func callbackQBluetoothSocket_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QBluetoothSocket::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQBluetoothSocketFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QBluetoothSocket) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QBluetoothSocket::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectEvent() {
	defer qt.Recovering("disconnect QBluetoothSocket::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QBluetoothSocket) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QBluetoothSocket::event")

	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QBluetoothSocket) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QBluetoothSocket::event")

	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQBluetoothSocket_EventFilter
func callbackQBluetoothSocket_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QBluetoothSocket::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQBluetoothSocketFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QBluetoothSocket) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QBluetoothSocket::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QBluetoothSocket::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QBluetoothSocket) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QBluetoothSocket::eventFilter")

	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QBluetoothSocket) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QBluetoothSocket::eventFilter")

	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQBluetoothSocket_MetaObject
func callbackQBluetoothSocket_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QBluetoothSocket::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQBluetoothSocketFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QBluetoothSocket) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QBluetoothSocket::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QBluetoothSocket) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QBluetoothSocket::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QBluetoothSocket) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QBluetoothSocket::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QBluetoothSocket_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBluetoothSocket) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QBluetoothSocket::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QBluetoothSocket_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QBluetoothTransferManager struct {
	core.QObject
}

type QBluetoothTransferManager_ITF interface {
	core.QObject_ITF
	QBluetoothTransferManager_PTR() *QBluetoothTransferManager
}

func (p *QBluetoothTransferManager) QBluetoothTransferManager_PTR() *QBluetoothTransferManager {
	return p
}

func (p *QBluetoothTransferManager) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QBluetoothTransferManager) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQBluetoothTransferManager(ptr QBluetoothTransferManager_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothTransferManager_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothTransferManagerFromPointer(ptr unsafe.Pointer) *QBluetoothTransferManager {
	var n = new(QBluetoothTransferManager)
	n.SetPointer(ptr)
	return n
}

func newQBluetoothTransferManagerFromPointer(ptr unsafe.Pointer) *QBluetoothTransferManager {
	var n = NewQBluetoothTransferManagerFromPointer(ptr)
	for len(n.ObjectName()) < len("QBluetoothTransferManager_") {
		n.SetObjectName("QBluetoothTransferManager_" + qt.Identifier())
	}
	return n
}

func (ptr *QBluetoothTransferManager) Put(request QBluetoothTransferRequest_ITF, data core.QIODevice_ITF) *QBluetoothTransferReply {
	defer qt.Recovering("QBluetoothTransferManager::put")

	if ptr.Pointer() != nil {
		return NewQBluetoothTransferReplyFromPointer(C.QBluetoothTransferManager_Put(ptr.Pointer(), PointerFromQBluetoothTransferRequest(request), core.PointerFromQIODevice(data)))
	}
	return nil
}

func NewQBluetoothTransferManager(parent core.QObject_ITF) *QBluetoothTransferManager {
	defer qt.Recovering("QBluetoothTransferManager::QBluetoothTransferManager")

	return newQBluetoothTransferManagerFromPointer(C.QBluetoothTransferManager_NewQBluetoothTransferManager(core.PointerFromQObject(parent)))
}

//export callbackQBluetoothTransferManager_Finished
func callbackQBluetoothTransferManager_Finished(ptr unsafe.Pointer, ptrName *C.char, reply unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothTransferManager::finished")

	if signal := qt.GetSignal(C.GoString(ptrName), "finished"); signal != nil {
		signal.(func(*QBluetoothTransferReply))(NewQBluetoothTransferReplyFromPointer(reply))
	}

}

func (ptr *QBluetoothTransferManager) ConnectFinished(f func(reply *QBluetoothTransferReply)) {
	defer qt.Recovering("connect QBluetoothTransferManager::finished")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QBluetoothTransferManager) DisconnectFinished() {
	defer qt.Recovering("disconnect QBluetoothTransferManager::finished")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

func (ptr *QBluetoothTransferManager) Finished(reply QBluetoothTransferReply_ITF) {
	defer qt.Recovering("QBluetoothTransferManager::finished")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_Finished(ptr.Pointer(), PointerFromQBluetoothTransferReply(reply))
	}
}

func (ptr *QBluetoothTransferManager) DestroyQBluetoothTransferManager() {
	defer qt.Recovering("QBluetoothTransferManager::~QBluetoothTransferManager")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QBluetoothTransferManager_DestroyQBluetoothTransferManager(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQBluetoothTransferManager_TimerEvent
func callbackQBluetoothTransferManager_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothTransferManager::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQBluetoothTransferManagerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QBluetoothTransferManager) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QBluetoothTransferManager::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QBluetoothTransferManager) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QBluetoothTransferManager::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QBluetoothTransferManager) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QBluetoothTransferManager::timerEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QBluetoothTransferManager) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QBluetoothTransferManager::timerEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQBluetoothTransferManager_ChildEvent
func callbackQBluetoothTransferManager_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothTransferManager::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQBluetoothTransferManagerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QBluetoothTransferManager) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QBluetoothTransferManager::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QBluetoothTransferManager) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QBluetoothTransferManager::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QBluetoothTransferManager) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QBluetoothTransferManager::childEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QBluetoothTransferManager) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QBluetoothTransferManager::childEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQBluetoothTransferManager_ConnectNotify
func callbackQBluetoothTransferManager_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothTransferManager::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBluetoothTransferManagerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBluetoothTransferManager) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QBluetoothTransferManager::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QBluetoothTransferManager) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QBluetoothTransferManager::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QBluetoothTransferManager) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QBluetoothTransferManager::connectNotify")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QBluetoothTransferManager) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QBluetoothTransferManager::connectNotify")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBluetoothTransferManager_CustomEvent
func callbackQBluetoothTransferManager_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothTransferManager::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQBluetoothTransferManagerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QBluetoothTransferManager) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QBluetoothTransferManager::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QBluetoothTransferManager) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QBluetoothTransferManager::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QBluetoothTransferManager) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QBluetoothTransferManager::customEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QBluetoothTransferManager) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QBluetoothTransferManager::customEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQBluetoothTransferManager_DeleteLater
func callbackQBluetoothTransferManager_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QBluetoothTransferManager::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQBluetoothTransferManagerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QBluetoothTransferManager) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QBluetoothTransferManager::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QBluetoothTransferManager) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QBluetoothTransferManager::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QBluetoothTransferManager) DeleteLater() {
	defer qt.Recovering("QBluetoothTransferManager::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QBluetoothTransferManager_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothTransferManager) DeleteLaterDefault() {
	defer qt.Recovering("QBluetoothTransferManager::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QBluetoothTransferManager_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQBluetoothTransferManager_DisconnectNotify
func callbackQBluetoothTransferManager_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothTransferManager::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBluetoothTransferManagerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBluetoothTransferManager) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QBluetoothTransferManager::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QBluetoothTransferManager) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QBluetoothTransferManager::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QBluetoothTransferManager) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QBluetoothTransferManager::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QBluetoothTransferManager) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QBluetoothTransferManager::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBluetoothTransferManager_Event
func callbackQBluetoothTransferManager_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QBluetoothTransferManager::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQBluetoothTransferManagerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QBluetoothTransferManager) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QBluetoothTransferManager::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QBluetoothTransferManager) DisconnectEvent() {
	defer qt.Recovering("disconnect QBluetoothTransferManager::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QBluetoothTransferManager) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QBluetoothTransferManager::event")

	if ptr.Pointer() != nil {
		return C.QBluetoothTransferManager_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QBluetoothTransferManager) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QBluetoothTransferManager::event")

	if ptr.Pointer() != nil {
		return C.QBluetoothTransferManager_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQBluetoothTransferManager_EventFilter
func callbackQBluetoothTransferManager_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QBluetoothTransferManager::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQBluetoothTransferManagerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QBluetoothTransferManager) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QBluetoothTransferManager::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QBluetoothTransferManager) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QBluetoothTransferManager::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QBluetoothTransferManager) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QBluetoothTransferManager::eventFilter")

	if ptr.Pointer() != nil {
		return C.QBluetoothTransferManager_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QBluetoothTransferManager) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QBluetoothTransferManager::eventFilter")

	if ptr.Pointer() != nil {
		return C.QBluetoothTransferManager_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQBluetoothTransferManager_MetaObject
func callbackQBluetoothTransferManager_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QBluetoothTransferManager::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQBluetoothTransferManagerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QBluetoothTransferManager) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QBluetoothTransferManager::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QBluetoothTransferManager) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QBluetoothTransferManager::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QBluetoothTransferManager) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QBluetoothTransferManager::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QBluetoothTransferManager_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBluetoothTransferManager) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QBluetoothTransferManager::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QBluetoothTransferManager_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QBluetoothTransferReply::TransferError
type QBluetoothTransferReply__TransferError int64

const (
	QBluetoothTransferReply__NoError                   = QBluetoothTransferReply__TransferError(0)
	QBluetoothTransferReply__UnknownError              = QBluetoothTransferReply__TransferError(1)
	QBluetoothTransferReply__FileNotFoundError         = QBluetoothTransferReply__TransferError(2)
	QBluetoothTransferReply__HostNotFoundError         = QBluetoothTransferReply__TransferError(3)
	QBluetoothTransferReply__UserCanceledTransferError = QBluetoothTransferReply__TransferError(4)
	QBluetoothTransferReply__IODeviceNotReadableError  = QBluetoothTransferReply__TransferError(5)
	QBluetoothTransferReply__ResourceBusyError         = QBluetoothTransferReply__TransferError(6)
	QBluetoothTransferReply__SessionError              = QBluetoothTransferReply__TransferError(7)
)

type QBluetoothTransferReply struct {
	core.QObject
}

type QBluetoothTransferReply_ITF interface {
	core.QObject_ITF
	QBluetoothTransferReply_PTR() *QBluetoothTransferReply
}

func (p *QBluetoothTransferReply) QBluetoothTransferReply_PTR() *QBluetoothTransferReply {
	return p
}

func (p *QBluetoothTransferReply) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QBluetoothTransferReply) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQBluetoothTransferReply(ptr QBluetoothTransferReply_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothTransferReply_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothTransferReplyFromPointer(ptr unsafe.Pointer) *QBluetoothTransferReply {
	var n = new(QBluetoothTransferReply)
	n.SetPointer(ptr)
	return n
}

func newQBluetoothTransferReplyFromPointer(ptr unsafe.Pointer) *QBluetoothTransferReply {
	var n = NewQBluetoothTransferReplyFromPointer(ptr)
	for len(n.ObjectName()) < len("QBluetoothTransferReply_") {
		n.SetObjectName("QBluetoothTransferReply_" + qt.Identifier())
	}
	return n
}

//export callbackQBluetoothTransferReply_Abort
func callbackQBluetoothTransferReply_Abort(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QBluetoothTransferReply::abort")

	if signal := qt.GetSignal(C.GoString(ptrName), "abort"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBluetoothTransferReply) ConnectAbort(f func()) {
	defer qt.Recovering("connect QBluetoothTransferReply::abort")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "abort", f)
	}
}

func (ptr *QBluetoothTransferReply) DisconnectAbort() {
	defer qt.Recovering("disconnect QBluetoothTransferReply::abort")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "abort")
	}
}

func (ptr *QBluetoothTransferReply) Abort() {
	defer qt.Recovering("QBluetoothTransferReply::abort")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_Abort(ptr.Pointer())
	}
}

func (ptr *QBluetoothTransferReply) SetManager(manager QBluetoothTransferManager_ITF) {
	defer qt.Recovering("QBluetoothTransferReply::setManager")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_SetManager(ptr.Pointer(), PointerFromQBluetoothTransferManager(manager))
	}
}

func (ptr *QBluetoothTransferReply) SetRequest(request QBluetoothTransferRequest_ITF) {
	defer qt.Recovering("QBluetoothTransferReply::setRequest")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_SetRequest(ptr.Pointer(), PointerFromQBluetoothTransferRequest(request))
	}
}

func NewQBluetoothTransferReply(parent core.QObject_ITF) *QBluetoothTransferReply {
	defer qt.Recovering("QBluetoothTransferReply::QBluetoothTransferReply")

	return newQBluetoothTransferReplyFromPointer(C.QBluetoothTransferReply_NewQBluetoothTransferReply(core.PointerFromQObject(parent)))
}

//export callbackQBluetoothTransferReply_Error2
func callbackQBluetoothTransferReply_Error2(ptr unsafe.Pointer, ptrName *C.char, errorType C.int) {
	defer qt.Recovering("callback QBluetoothTransferReply::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(QBluetoothTransferReply__TransferError))(QBluetoothTransferReply__TransferError(errorType))
	}

}

func (ptr *QBluetoothTransferReply) ConnectError2(f func(errorType QBluetoothTransferReply__TransferError)) {
	defer qt.Recovering("connect QBluetoothTransferReply::error")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error2", f)
	}
}

func (ptr *QBluetoothTransferReply) DisconnectError2() {
	defer qt.Recovering("disconnect QBluetoothTransferReply::error")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error2")
	}
}

func (ptr *QBluetoothTransferReply) Error2(errorType QBluetoothTransferReply__TransferError) {
	defer qt.Recovering("QBluetoothTransferReply::error")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_Error2(ptr.Pointer(), C.int(errorType))
	}
}

//export callbackQBluetoothTransferReply_Error
func callbackQBluetoothTransferReply_Error(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QBluetoothTransferReply::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error"); signal != nil {
		return C.int(signal.(func() QBluetoothTransferReply__TransferError)())
	}

	return C.int(0)
}

func (ptr *QBluetoothTransferReply) ConnectError(f func() QBluetoothTransferReply__TransferError) {
	defer qt.Recovering("connect QBluetoothTransferReply::error")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "error", f)
	}
}

func (ptr *QBluetoothTransferReply) DisconnectError() {
	defer qt.Recovering("disconnect QBluetoothTransferReply::error")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "error")
	}
}

func (ptr *QBluetoothTransferReply) Error() QBluetoothTransferReply__TransferError {
	defer qt.Recovering("QBluetoothTransferReply::error")

	if ptr.Pointer() != nil {
		return QBluetoothTransferReply__TransferError(C.QBluetoothTransferReply_Error(ptr.Pointer()))
	}
	return 0
}

//export callbackQBluetoothTransferReply_ErrorString
func callbackQBluetoothTransferReply_ErrorString(ptr unsafe.Pointer, ptrName *C.char) *C.char {
	defer qt.Recovering("callback QBluetoothTransferReply::errorString")

	if signal := qt.GetSignal(C.GoString(ptrName), "errorString"); signal != nil {
		return C.CString(signal.(func() string)())
	}

	return C.CString("")
}

func (ptr *QBluetoothTransferReply) ConnectErrorString(f func() string) {
	defer qt.Recovering("connect QBluetoothTransferReply::errorString")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "errorString", f)
	}
}

func (ptr *QBluetoothTransferReply) DisconnectErrorString() {
	defer qt.Recovering("disconnect QBluetoothTransferReply::errorString")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "errorString")
	}
}

func (ptr *QBluetoothTransferReply) ErrorString() string {
	defer qt.Recovering("QBluetoothTransferReply::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothTransferReply_ErrorString(ptr.Pointer()))
	}
	return ""
}

//export callbackQBluetoothTransferReply_Finished
func callbackQBluetoothTransferReply_Finished(ptr unsafe.Pointer, ptrName *C.char, reply unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothTransferReply::finished")

	if signal := qt.GetSignal(C.GoString(ptrName), "finished"); signal != nil {
		signal.(func(*QBluetoothTransferReply))(NewQBluetoothTransferReplyFromPointer(reply))
	}

}

func (ptr *QBluetoothTransferReply) ConnectFinished(f func(reply *QBluetoothTransferReply)) {
	defer qt.Recovering("connect QBluetoothTransferReply::finished")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QBluetoothTransferReply) DisconnectFinished() {
	defer qt.Recovering("disconnect QBluetoothTransferReply::finished")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

func (ptr *QBluetoothTransferReply) Finished(reply QBluetoothTransferReply_ITF) {
	defer qt.Recovering("QBluetoothTransferReply::finished")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_Finished(ptr.Pointer(), PointerFromQBluetoothTransferReply(reply))
	}
}

//export callbackQBluetoothTransferReply_IsFinished
func callbackQBluetoothTransferReply_IsFinished(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QBluetoothTransferReply::isFinished")

	if signal := qt.GetSignal(C.GoString(ptrName), "isFinished"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(false))
}

func (ptr *QBluetoothTransferReply) ConnectIsFinished(f func() bool) {
	defer qt.Recovering("connect QBluetoothTransferReply::isFinished")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "isFinished", f)
	}
}

func (ptr *QBluetoothTransferReply) DisconnectIsFinished() {
	defer qt.Recovering("disconnect QBluetoothTransferReply::isFinished")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "isFinished")
	}
}

func (ptr *QBluetoothTransferReply) IsFinished() bool {
	defer qt.Recovering("QBluetoothTransferReply::isFinished")

	if ptr.Pointer() != nil {
		return C.QBluetoothTransferReply_IsFinished(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQBluetoothTransferReply_IsRunning
func callbackQBluetoothTransferReply_IsRunning(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QBluetoothTransferReply::isRunning")

	if signal := qt.GetSignal(C.GoString(ptrName), "isRunning"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(false))
}

func (ptr *QBluetoothTransferReply) ConnectIsRunning(f func() bool) {
	defer qt.Recovering("connect QBluetoothTransferReply::isRunning")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "isRunning", f)
	}
}

func (ptr *QBluetoothTransferReply) DisconnectIsRunning() {
	defer qt.Recovering("disconnect QBluetoothTransferReply::isRunning")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "isRunning")
	}
}

func (ptr *QBluetoothTransferReply) IsRunning() bool {
	defer qt.Recovering("QBluetoothTransferReply::isRunning")

	if ptr.Pointer() != nil {
		return C.QBluetoothTransferReply_IsRunning(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothTransferReply) Manager() *QBluetoothTransferManager {
	defer qt.Recovering("QBluetoothTransferReply::manager")

	if ptr.Pointer() != nil {
		return NewQBluetoothTransferManagerFromPointer(C.QBluetoothTransferReply_Manager(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBluetoothTransferReply) Request() *QBluetoothTransferRequest {
	defer qt.Recovering("QBluetoothTransferReply::request")

	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothTransferRequestFromPointer(C.QBluetoothTransferReply_Request(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothTransferRequest).DestroyQBluetoothTransferRequest)
		return tmpValue
	}
	return nil
}

//export callbackQBluetoothTransferReply_TransferProgress
func callbackQBluetoothTransferReply_TransferProgress(ptr unsafe.Pointer, ptrName *C.char, bytesTransferred C.longlong, bytesTotal C.longlong) {
	defer qt.Recovering("callback QBluetoothTransferReply::transferProgress")

	if signal := qt.GetSignal(C.GoString(ptrName), "transferProgress"); signal != nil {
		signal.(func(int64, int64))(int64(bytesTransferred), int64(bytesTotal))
	}

}

func (ptr *QBluetoothTransferReply) ConnectTransferProgress(f func(bytesTransferred int64, bytesTotal int64)) {
	defer qt.Recovering("connect QBluetoothTransferReply::transferProgress")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_ConnectTransferProgress(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "transferProgress", f)
	}
}

func (ptr *QBluetoothTransferReply) DisconnectTransferProgress() {
	defer qt.Recovering("disconnect QBluetoothTransferReply::transferProgress")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_DisconnectTransferProgress(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "transferProgress")
	}
}

func (ptr *QBluetoothTransferReply) TransferProgress(bytesTransferred int64, bytesTotal int64) {
	defer qt.Recovering("QBluetoothTransferReply::transferProgress")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_TransferProgress(ptr.Pointer(), C.longlong(bytesTransferred), C.longlong(bytesTotal))
	}
}

func (ptr *QBluetoothTransferReply) DestroyQBluetoothTransferReply() {
	defer qt.Recovering("QBluetoothTransferReply::~QBluetoothTransferReply")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QBluetoothTransferReply_DestroyQBluetoothTransferReply(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQBluetoothTransferReply_TimerEvent
func callbackQBluetoothTransferReply_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothTransferReply::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQBluetoothTransferReplyFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QBluetoothTransferReply) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QBluetoothTransferReply::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QBluetoothTransferReply) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QBluetoothTransferReply::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QBluetoothTransferReply) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QBluetoothTransferReply::timerEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QBluetoothTransferReply) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QBluetoothTransferReply::timerEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQBluetoothTransferReply_ChildEvent
func callbackQBluetoothTransferReply_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothTransferReply::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQBluetoothTransferReplyFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QBluetoothTransferReply) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QBluetoothTransferReply::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QBluetoothTransferReply) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QBluetoothTransferReply::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QBluetoothTransferReply) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QBluetoothTransferReply::childEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QBluetoothTransferReply) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QBluetoothTransferReply::childEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQBluetoothTransferReply_ConnectNotify
func callbackQBluetoothTransferReply_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothTransferReply::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBluetoothTransferReplyFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBluetoothTransferReply) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QBluetoothTransferReply::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QBluetoothTransferReply) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QBluetoothTransferReply::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QBluetoothTransferReply) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QBluetoothTransferReply::connectNotify")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QBluetoothTransferReply) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QBluetoothTransferReply::connectNotify")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBluetoothTransferReply_CustomEvent
func callbackQBluetoothTransferReply_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothTransferReply::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQBluetoothTransferReplyFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QBluetoothTransferReply) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QBluetoothTransferReply::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QBluetoothTransferReply) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QBluetoothTransferReply::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QBluetoothTransferReply) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QBluetoothTransferReply::customEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QBluetoothTransferReply) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QBluetoothTransferReply::customEvent")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQBluetoothTransferReply_DeleteLater
func callbackQBluetoothTransferReply_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QBluetoothTransferReply::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQBluetoothTransferReplyFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QBluetoothTransferReply) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QBluetoothTransferReply::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QBluetoothTransferReply) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QBluetoothTransferReply::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QBluetoothTransferReply) DeleteLater() {
	defer qt.Recovering("QBluetoothTransferReply::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QBluetoothTransferReply_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothTransferReply) DeleteLaterDefault() {
	defer qt.Recovering("QBluetoothTransferReply::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QBluetoothTransferReply_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQBluetoothTransferReply_DisconnectNotify
func callbackQBluetoothTransferReply_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QBluetoothTransferReply::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBluetoothTransferReplyFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBluetoothTransferReply) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QBluetoothTransferReply::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QBluetoothTransferReply) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QBluetoothTransferReply::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QBluetoothTransferReply) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QBluetoothTransferReply::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QBluetoothTransferReply) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QBluetoothTransferReply::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBluetoothTransferReply_Event
func callbackQBluetoothTransferReply_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QBluetoothTransferReply::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQBluetoothTransferReplyFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QBluetoothTransferReply) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QBluetoothTransferReply::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QBluetoothTransferReply) DisconnectEvent() {
	defer qt.Recovering("disconnect QBluetoothTransferReply::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QBluetoothTransferReply) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QBluetoothTransferReply::event")

	if ptr.Pointer() != nil {
		return C.QBluetoothTransferReply_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QBluetoothTransferReply) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QBluetoothTransferReply::event")

	if ptr.Pointer() != nil {
		return C.QBluetoothTransferReply_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQBluetoothTransferReply_EventFilter
func callbackQBluetoothTransferReply_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QBluetoothTransferReply::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQBluetoothTransferReplyFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QBluetoothTransferReply) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QBluetoothTransferReply::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QBluetoothTransferReply) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QBluetoothTransferReply::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QBluetoothTransferReply) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QBluetoothTransferReply::eventFilter")

	if ptr.Pointer() != nil {
		return C.QBluetoothTransferReply_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QBluetoothTransferReply) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QBluetoothTransferReply::eventFilter")

	if ptr.Pointer() != nil {
		return C.QBluetoothTransferReply_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQBluetoothTransferReply_MetaObject
func callbackQBluetoothTransferReply_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QBluetoothTransferReply::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQBluetoothTransferReplyFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QBluetoothTransferReply) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QBluetoothTransferReply::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QBluetoothTransferReply) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QBluetoothTransferReply::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QBluetoothTransferReply) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QBluetoothTransferReply::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QBluetoothTransferReply_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QBluetoothTransferReply) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QBluetoothTransferReply::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QBluetoothTransferReply_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QBluetoothTransferRequest::Attribute
type QBluetoothTransferRequest__Attribute int64

const (
	QBluetoothTransferRequest__DescriptionAttribute = QBluetoothTransferRequest__Attribute(0)
	QBluetoothTransferRequest__TimeAttribute        = QBluetoothTransferRequest__Attribute(1)
	QBluetoothTransferRequest__TypeAttribute        = QBluetoothTransferRequest__Attribute(2)
	QBluetoothTransferRequest__LengthAttribute      = QBluetoothTransferRequest__Attribute(3)
	QBluetoothTransferRequest__NameAttribute        = QBluetoothTransferRequest__Attribute(4)
)

type QBluetoothTransferRequest struct {
	ptr unsafe.Pointer
}

type QBluetoothTransferRequest_ITF interface {
	QBluetoothTransferRequest_PTR() *QBluetoothTransferRequest
}

func (p *QBluetoothTransferRequest) QBluetoothTransferRequest_PTR() *QBluetoothTransferRequest {
	return p
}

func (p *QBluetoothTransferRequest) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QBluetoothTransferRequest) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQBluetoothTransferRequest(ptr QBluetoothTransferRequest_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothTransferRequest_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothTransferRequestFromPointer(ptr unsafe.Pointer) *QBluetoothTransferRequest {
	var n = new(QBluetoothTransferRequest)
	n.SetPointer(ptr)
	return n
}

func newQBluetoothTransferRequestFromPointer(ptr unsafe.Pointer) *QBluetoothTransferRequest {
	var n = NewQBluetoothTransferRequestFromPointer(ptr)
	return n
}

func NewQBluetoothTransferRequest(address QBluetoothAddress_ITF) *QBluetoothTransferRequest {
	defer qt.Recovering("QBluetoothTransferRequest::QBluetoothTransferRequest")

	return newQBluetoothTransferRequestFromPointer(C.QBluetoothTransferRequest_NewQBluetoothTransferRequest(PointerFromQBluetoothAddress(address)))
}

func NewQBluetoothTransferRequest2(other QBluetoothTransferRequest_ITF) *QBluetoothTransferRequest {
	defer qt.Recovering("QBluetoothTransferRequest::QBluetoothTransferRequest")

	return newQBluetoothTransferRequestFromPointer(C.QBluetoothTransferRequest_NewQBluetoothTransferRequest2(PointerFromQBluetoothTransferRequest(other)))
}

func (ptr *QBluetoothTransferRequest) Address() *QBluetoothAddress {
	defer qt.Recovering("QBluetoothTransferRequest::address")

	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothAddressFromPointer(C.QBluetoothTransferRequest_Address(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothTransferRequest) Attribute(code QBluetoothTransferRequest__Attribute, defaultValue core.QVariant_ITF) *core.QVariant {
	defer qt.Recovering("QBluetoothTransferRequest::attribute")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QBluetoothTransferRequest_Attribute(ptr.Pointer(), C.int(code), core.PointerFromQVariant(defaultValue)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothTransferRequest) SetAttribute(code QBluetoothTransferRequest__Attribute, value core.QVariant_ITF) {
	defer qt.Recovering("QBluetoothTransferRequest::setAttribute")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferRequest_SetAttribute(ptr.Pointer(), C.int(code), core.PointerFromQVariant(value))
	}
}

func (ptr *QBluetoothTransferRequest) DestroyQBluetoothTransferRequest() {
	defer qt.Recovering("QBluetoothTransferRequest::~QBluetoothTransferRequest")

	if ptr.Pointer() != nil {
		C.QBluetoothTransferRequest_DestroyQBluetoothTransferRequest(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//QBluetoothUuid::CharacteristicType
type QBluetoothUuid__CharacteristicType int64

const (
	QBluetoothUuid__DeviceName                                    = QBluetoothUuid__CharacteristicType(0x2a00)
	QBluetoothUuid__Appearance                                    = QBluetoothUuid__CharacteristicType(0x2a01)
	QBluetoothUuid__PeripheralPrivacyFlag                         = QBluetoothUuid__CharacteristicType(0x2a02)
	QBluetoothUuid__ReconnectionAddress                           = QBluetoothUuid__CharacteristicType(0x2a03)
	QBluetoothUuid__PeripheralPreferredConnectionParameters       = QBluetoothUuid__CharacteristicType(0x2a04)
	QBluetoothUuid__ServiceChanged                                = QBluetoothUuid__CharacteristicType(0x2a05)
	QBluetoothUuid__AlertLevel                                    = QBluetoothUuid__CharacteristicType(0x2a06)
	QBluetoothUuid__TxPowerLevel                                  = QBluetoothUuid__CharacteristicType(0x2a07)
	QBluetoothUuid__DateTime                                      = QBluetoothUuid__CharacteristicType(0x2a08)
	QBluetoothUuid__DayOfWeek                                     = QBluetoothUuid__CharacteristicType(0x2a09)
	QBluetoothUuid__DayDateTime                                   = QBluetoothUuid__CharacteristicType(0x2a0a)
	QBluetoothUuid__ExactTime256                                  = QBluetoothUuid__CharacteristicType(0x2a0c)
	QBluetoothUuid__DSTOffset                                     = QBluetoothUuid__CharacteristicType(0x2a0d)
	QBluetoothUuid__TimeZone                                      = QBluetoothUuid__CharacteristicType(0x2a0e)
	QBluetoothUuid__LocalTimeInformation                          = QBluetoothUuid__CharacteristicType(0x2a0f)
	QBluetoothUuid__TimeWithDST                                   = QBluetoothUuid__CharacteristicType(0x2a11)
	QBluetoothUuid__TimeAccuracy                                  = QBluetoothUuid__CharacteristicType(0x2a12)
	QBluetoothUuid__TimeSource                                    = QBluetoothUuid__CharacteristicType(0x2a13)
	QBluetoothUuid__ReferenceTimeInformation                      = QBluetoothUuid__CharacteristicType(0x2a14)
	QBluetoothUuid__TimeUpdateControlPoint                        = QBluetoothUuid__CharacteristicType(0x2a16)
	QBluetoothUuid__TimeUpdateState                               = QBluetoothUuid__CharacteristicType(0x2a17)
	QBluetoothUuid__GlucoseMeasurement                            = QBluetoothUuid__CharacteristicType(0x2a18)
	QBluetoothUuid__BatteryLevel                                  = QBluetoothUuid__CharacteristicType(0x2a19)
	QBluetoothUuid__TemperatureMeasurement                        = QBluetoothUuid__CharacteristicType(0x2a1c)
	QBluetoothUuid__TemperatureType                               = QBluetoothUuid__CharacteristicType(0x2a1d)
	QBluetoothUuid__IntermediateTemperature                       = QBluetoothUuid__CharacteristicType(0x2a1e)
	QBluetoothUuid__MeasurementInterval                           = QBluetoothUuid__CharacteristicType(0x2a21)
	QBluetoothUuid__BootKeyboardInputReport                       = QBluetoothUuid__CharacteristicType(0x2a22)
	QBluetoothUuid__SystemID                                      = QBluetoothUuid__CharacteristicType(0x2a23)
	QBluetoothUuid__ModelNumberString                             = QBluetoothUuid__CharacteristicType(0x2a24)
	QBluetoothUuid__SerialNumberString                            = QBluetoothUuid__CharacteristicType(0x2a25)
	QBluetoothUuid__FirmwareRevisionString                        = QBluetoothUuid__CharacteristicType(0x2a26)
	QBluetoothUuid__HardwareRevisionString                        = QBluetoothUuid__CharacteristicType(0x2a27)
	QBluetoothUuid__SoftwareRevisionString                        = QBluetoothUuid__CharacteristicType(0x2a28)
	QBluetoothUuid__ManufacturerNameString                        = QBluetoothUuid__CharacteristicType(0x2a29)
	QBluetoothUuid__IEEE1107320601RegulatoryCertificationDataList = QBluetoothUuid__CharacteristicType(0x2a2a)
	QBluetoothUuid__CurrentTime                                   = QBluetoothUuid__CharacteristicType(0x2a2b)
	QBluetoothUuid__MagneticDeclination                           = QBluetoothUuid__CharacteristicType(0x2a2c)
	QBluetoothUuid__ScanRefresh                                   = QBluetoothUuid__CharacteristicType(0x2a31)
	QBluetoothUuid__BootKeyboardOutputReport                      = QBluetoothUuid__CharacteristicType(0x2a32)
	QBluetoothUuid__BootMouseInputReport                          = QBluetoothUuid__CharacteristicType(0x2a33)
	QBluetoothUuid__GlucoseMeasurementContext                     = QBluetoothUuid__CharacteristicType(0x2a34)
	QBluetoothUuid__BloodPressureMeasurement                      = QBluetoothUuid__CharacteristicType(0x2a35)
	QBluetoothUuid__IntermediateCuffPressure                      = QBluetoothUuid__CharacteristicType(0x2a36)
	QBluetoothUuid__HeartRateMeasurement                          = QBluetoothUuid__CharacteristicType(0x2a37)
	QBluetoothUuid__BodySensorLocation                            = QBluetoothUuid__CharacteristicType(0x2a38)
	QBluetoothUuid__HeartRateControlPoint                         = QBluetoothUuid__CharacteristicType(0x2a39)
	QBluetoothUuid__AlertStatus                                   = QBluetoothUuid__CharacteristicType(0x2a3f)
	QBluetoothUuid__RingerControlPoint                            = QBluetoothUuid__CharacteristicType(0x2a40)
	QBluetoothUuid__RingerSetting                                 = QBluetoothUuid__CharacteristicType(0x2a41)
	QBluetoothUuid__AlertCategoryIDBitMask                        = QBluetoothUuid__CharacteristicType(0x2a42)
	QBluetoothUuid__AlertCategoryID                               = QBluetoothUuid__CharacteristicType(0x2a43)
	QBluetoothUuid__AlertNotificationControlPoint                 = QBluetoothUuid__CharacteristicType(0x2a44)
	QBluetoothUuid__UnreadAlertStatus                             = QBluetoothUuid__CharacteristicType(0x2a45)
	QBluetoothUuid__NewAlert                                      = QBluetoothUuid__CharacteristicType(0x2a46)
	QBluetoothUuid__SupportedNewAlertCategory                     = QBluetoothUuid__CharacteristicType(0x2a47)
	QBluetoothUuid__SupportedUnreadAlertCategory                  = QBluetoothUuid__CharacteristicType(0x2a48)
	QBluetoothUuid__BloodPressureFeature                          = QBluetoothUuid__CharacteristicType(0x2a49)
	QBluetoothUuid__HIDInformation                                = QBluetoothUuid__CharacteristicType(0x2a4a)
	QBluetoothUuid__ReportMap                                     = QBluetoothUuid__CharacteristicType(0x2a4b)
	QBluetoothUuid__HIDControlPoint                               = QBluetoothUuid__CharacteristicType(0x2a4c)
	QBluetoothUuid__Report                                        = QBluetoothUuid__CharacteristicType(0x2a4d)
	QBluetoothUuid__ProtocolMode                                  = QBluetoothUuid__CharacteristicType(0x2a4e)
	QBluetoothUuid__ScanIntervalWindow                            = QBluetoothUuid__CharacteristicType(0x2a4f)
	QBluetoothUuid__PnPID                                         = QBluetoothUuid__CharacteristicType(0x2a50)
	QBluetoothUuid__GlucoseFeature                                = QBluetoothUuid__CharacteristicType(0x2a51)
	QBluetoothUuid__RecordAccessControlPoint                      = QBluetoothUuid__CharacteristicType(0x2a52)
	QBluetoothUuid__RSCMeasurement                                = QBluetoothUuid__CharacteristicType(0x2a53)
	QBluetoothUuid__RSCFeature                                    = QBluetoothUuid__CharacteristicType(0x2a54)
	QBluetoothUuid__SCControlPoint                                = QBluetoothUuid__CharacteristicType(0x2a55)
	QBluetoothUuid__CSCMeasurement                                = QBluetoothUuid__CharacteristicType(0x2a5b)
	QBluetoothUuid__CSCFeature                                    = QBluetoothUuid__CharacteristicType(0x2a5c)
	QBluetoothUuid__SensorLocation                                = QBluetoothUuid__CharacteristicType(0x2a5d)
	QBluetoothUuid__CyclingPowerMeasurement                       = QBluetoothUuid__CharacteristicType(0x2a63)
	QBluetoothUuid__CyclingPowerVector                            = QBluetoothUuid__CharacteristicType(0x2a64)
	QBluetoothUuid__CyclingPowerFeature                           = QBluetoothUuid__CharacteristicType(0x2a65)
	QBluetoothUuid__CyclingPowerControlPoint                      = QBluetoothUuid__CharacteristicType(0x2a66)
	QBluetoothUuid__LocationAndSpeed                              = QBluetoothUuid__CharacteristicType(0x2a67)
	QBluetoothUuid__Navigation                                    = QBluetoothUuid__CharacteristicType(0x2a68)
	QBluetoothUuid__PositionQuality                               = QBluetoothUuid__CharacteristicType(0x2a69)
	QBluetoothUuid__LNFeature                                     = QBluetoothUuid__CharacteristicType(0x2a6a)
	QBluetoothUuid__LNControlPoint                                = QBluetoothUuid__CharacteristicType(0x2a6b)
	QBluetoothUuid__Elevation                                     = QBluetoothUuid__CharacteristicType(0x2a6c)
	QBluetoothUuid__Pressure                                      = QBluetoothUuid__CharacteristicType(0x2a6d)
	QBluetoothUuid__Temperature                                   = QBluetoothUuid__CharacteristicType(0x2a6e)
	QBluetoothUuid__Humidity                                      = QBluetoothUuid__CharacteristicType(0x2a6f)
	QBluetoothUuid__TrueWindSpeed                                 = QBluetoothUuid__CharacteristicType(0x2a70)
	QBluetoothUuid__TrueWindDirection                             = QBluetoothUuid__CharacteristicType(0x2a71)
	QBluetoothUuid__ApparentWindSpeed                             = QBluetoothUuid__CharacteristicType(0x2a72)
	QBluetoothUuid__ApparentWindDirection                         = QBluetoothUuid__CharacteristicType(0x2a73)
	QBluetoothUuid__GustFactor                                    = QBluetoothUuid__CharacteristicType(0x2a74)
	QBluetoothUuid__PollenConcentration                           = QBluetoothUuid__CharacteristicType(0x2a75)
	QBluetoothUuid__UVIndex                                       = QBluetoothUuid__CharacteristicType(0x2a76)
	QBluetoothUuid__Irradiance                                    = QBluetoothUuid__CharacteristicType(0x2a77)
	QBluetoothUuid__Rainfall                                      = QBluetoothUuid__CharacteristicType(0x2a78)
	QBluetoothUuid__WindChill                                     = QBluetoothUuid__CharacteristicType(0x2a79)
	QBluetoothUuid__HeatIndex                                     = QBluetoothUuid__CharacteristicType(0x2a7a)
	QBluetoothUuid__DewPoint                                      = QBluetoothUuid__CharacteristicType(0x2a7b)
	QBluetoothUuid__DescriptorValueChanged                        = QBluetoothUuid__CharacteristicType(0x2a7d)
	QBluetoothUuid__AerobicHeartRateLowerLimit                    = QBluetoothUuid__CharacteristicType(0x2a7e)
	QBluetoothUuid__AerobicThreshold                              = QBluetoothUuid__CharacteristicType(0x2a7f)
	QBluetoothUuid__Age                                           = QBluetoothUuid__CharacteristicType(0x2a80)
	QBluetoothUuid__AnaerobicHeartRateLowerLimit                  = QBluetoothUuid__CharacteristicType(0x2a81)
	QBluetoothUuid__AnaerobicHeartRateUpperLimit                  = QBluetoothUuid__CharacteristicType(0x2a82)
	QBluetoothUuid__AnaerobicThreshold                            = QBluetoothUuid__CharacteristicType(0x2a83)
	QBluetoothUuid__AerobicHeartRateUpperLimit                    = QBluetoothUuid__CharacteristicType(0x2a84)
	QBluetoothUuid__DateOfBirth                                   = QBluetoothUuid__CharacteristicType(0x2a85)
	QBluetoothUuid__DateOfThresholdAssessment                     = QBluetoothUuid__CharacteristicType(0x2a86)
	QBluetoothUuid__EmailAddress                                  = QBluetoothUuid__CharacteristicType(0x2a87)
	QBluetoothUuid__FatBurnHeartRateLowerLimit                    = QBluetoothUuid__CharacteristicType(0x2a88)
	QBluetoothUuid__FatBurnHeartRateUpperLimit                    = QBluetoothUuid__CharacteristicType(0x2a89)
	QBluetoothUuid__FirstName                                     = QBluetoothUuid__CharacteristicType(0x2a8a)
	QBluetoothUuid__FiveZoneHeartRateLimits                       = QBluetoothUuid__CharacteristicType(0x2a8b)
	QBluetoothUuid__Gender                                        = QBluetoothUuid__CharacteristicType(0x2a8c)
	QBluetoothUuid__HeartRateMax                                  = QBluetoothUuid__CharacteristicType(0x2a8d)
	QBluetoothUuid__Height                                        = QBluetoothUuid__CharacteristicType(0x2a8e)
	QBluetoothUuid__HipCircumference                              = QBluetoothUuid__CharacteristicType(0x2a8f)
	QBluetoothUuid__LastName                                      = QBluetoothUuid__CharacteristicType(0x2a90)
	QBluetoothUuid__MaximumRecommendedHeartRate                   = QBluetoothUuid__CharacteristicType(0x2a91)
	QBluetoothUuid__RestingHeartRate                              = QBluetoothUuid__CharacteristicType(0x2a92)
	QBluetoothUuid__SportTypeForAerobicAnaerobicThresholds        = QBluetoothUuid__CharacteristicType(0x2a93)
	QBluetoothUuid__ThreeZoneHeartRateLimits                      = QBluetoothUuid__CharacteristicType(0x2a94)
	QBluetoothUuid__TwoZoneHeartRateLimits                        = QBluetoothUuid__CharacteristicType(0x2a95)
	QBluetoothUuid__VO2Max                                        = QBluetoothUuid__CharacteristicType(0x2a96)
	QBluetoothUuid__WaistCircumference                            = QBluetoothUuid__CharacteristicType(0x2a97)
	QBluetoothUuid__Weight                                        = QBluetoothUuid__CharacteristicType(0x2a98)
	QBluetoothUuid__DatabaseChangeIncrement                       = QBluetoothUuid__CharacteristicType(0x2a99)
	QBluetoothUuid__UserIndex                                     = QBluetoothUuid__CharacteristicType(0x2a9a)
	QBluetoothUuid__BodyCompositionFeature                        = QBluetoothUuid__CharacteristicType(0x2a9b)
	QBluetoothUuid__BodyCompositionMeasurement                    = QBluetoothUuid__CharacteristicType(0x2a9c)
	QBluetoothUuid__WeightMeasurement                             = QBluetoothUuid__CharacteristicType(0x2a9d)
	QBluetoothUuid__WeightScaleFeature                            = QBluetoothUuid__CharacteristicType(0x2a9e)
	QBluetoothUuid__UserControlPoint                              = QBluetoothUuid__CharacteristicType(0x2a9f)
	QBluetoothUuid__MagneticFluxDensity2D                         = QBluetoothUuid__CharacteristicType(0x2aa0)
	QBluetoothUuid__MagneticFluxDensity3D                         = QBluetoothUuid__CharacteristicType(0x2aa1)
	QBluetoothUuid__Language                                      = QBluetoothUuid__CharacteristicType(0x2aa2)
	QBluetoothUuid__BarometricPressureTrend                       = QBluetoothUuid__CharacteristicType(0x2aa3)
)

//QBluetoothUuid::DescriptorType
type QBluetoothUuid__DescriptorType int64

const (
	QBluetoothUuid__UnknownDescriptorType              = QBluetoothUuid__DescriptorType(0x0)
	QBluetoothUuid__CharacteristicExtendedProperties   = QBluetoothUuid__DescriptorType(0x2900)
	QBluetoothUuid__CharacteristicUserDescription      = QBluetoothUuid__DescriptorType(0x2901)
	QBluetoothUuid__ClientCharacteristicConfiguration  = QBluetoothUuid__DescriptorType(0x2902)
	QBluetoothUuid__ServerCharacteristicConfiguration  = QBluetoothUuid__DescriptorType(0x2903)
	QBluetoothUuid__CharacteristicPresentationFormat   = QBluetoothUuid__DescriptorType(0x2904)
	QBluetoothUuid__CharacteristicAggregateFormat      = QBluetoothUuid__DescriptorType(0x2905)
	QBluetoothUuid__ValidRange                         = QBluetoothUuid__DescriptorType(0x2906)
	QBluetoothUuid__ExternalReportReference            = QBluetoothUuid__DescriptorType(0x2907)
	QBluetoothUuid__ReportReference                    = QBluetoothUuid__DescriptorType(0x2908)
	QBluetoothUuid__EnvironmentalSensingConfiguration  = QBluetoothUuid__DescriptorType(0x290b)
	QBluetoothUuid__EnvironmentalSensingMeasurement    = QBluetoothUuid__DescriptorType(0x290c)
	QBluetoothUuid__EnvironmentalSensingTriggerSetting = QBluetoothUuid__DescriptorType(0x290d)
)

//QBluetoothUuid::ProtocolUuid
type QBluetoothUuid__ProtocolUuid int64

const (
	QBluetoothUuid__Sdp                    = QBluetoothUuid__ProtocolUuid(0x0001)
	QBluetoothUuid__Udp                    = QBluetoothUuid__ProtocolUuid(0x0002)
	QBluetoothUuid__Rfcomm                 = QBluetoothUuid__ProtocolUuid(0x0003)
	QBluetoothUuid__Tcp                    = QBluetoothUuid__ProtocolUuid(0x0004)
	QBluetoothUuid__TcsBin                 = QBluetoothUuid__ProtocolUuid(0x0005)
	QBluetoothUuid__TcsAt                  = QBluetoothUuid__ProtocolUuid(0x0006)
	QBluetoothUuid__Att                    = QBluetoothUuid__ProtocolUuid(0x0007)
	QBluetoothUuid__Obex                   = QBluetoothUuid__ProtocolUuid(0x0008)
	QBluetoothUuid__Ip                     = QBluetoothUuid__ProtocolUuid(0x0009)
	QBluetoothUuid__Ftp                    = QBluetoothUuid__ProtocolUuid(0x000A)
	QBluetoothUuid__Http                   = QBluetoothUuid__ProtocolUuid(0x000C)
	QBluetoothUuid__Wsp                    = QBluetoothUuid__ProtocolUuid(0x000E)
	QBluetoothUuid__Bnep                   = QBluetoothUuid__ProtocolUuid(0x000F)
	QBluetoothUuid__Upnp                   = QBluetoothUuid__ProtocolUuid(0x0010)
	QBluetoothUuid__Hidp                   = QBluetoothUuid__ProtocolUuid(0x0011)
	QBluetoothUuid__HardcopyControlChannel = QBluetoothUuid__ProtocolUuid(0x0012)
	QBluetoothUuid__HardcopyDataChannel    = QBluetoothUuid__ProtocolUuid(0x0014)
	QBluetoothUuid__HardcopyNotification   = QBluetoothUuid__ProtocolUuid(0x0016)
	QBluetoothUuid__Avctp                  = QBluetoothUuid__ProtocolUuid(0x0017)
	QBluetoothUuid__Avdtp                  = QBluetoothUuid__ProtocolUuid(0x0019)
	QBluetoothUuid__Cmtp                   = QBluetoothUuid__ProtocolUuid(0x001B)
	QBluetoothUuid__UdiCPlain              = QBluetoothUuid__ProtocolUuid(0x001D)
	QBluetoothUuid__McapControlChannel     = QBluetoothUuid__ProtocolUuid(0x001E)
	QBluetoothUuid__McapDataChannel        = QBluetoothUuid__ProtocolUuid(0x001F)
	QBluetoothUuid__L2cap                  = QBluetoothUuid__ProtocolUuid(0x0100)
)

//QBluetoothUuid::ServiceClassUuid
type QBluetoothUuid__ServiceClassUuid int64

const (
	QBluetoothUuid__ServiceDiscoveryServer                = QBluetoothUuid__ServiceClassUuid(0x1000)
	QBluetoothUuid__BrowseGroupDescriptor                 = QBluetoothUuid__ServiceClassUuid(0x1001)
	QBluetoothUuid__PublicBrowseGroup                     = QBluetoothUuid__ServiceClassUuid(0x1002)
	QBluetoothUuid__SerialPort                            = QBluetoothUuid__ServiceClassUuid(0x1101)
	QBluetoothUuid__LANAccessUsingPPP                     = QBluetoothUuid__ServiceClassUuid(0x1102)
	QBluetoothUuid__DialupNetworking                      = QBluetoothUuid__ServiceClassUuid(0x1103)
	QBluetoothUuid__IrMCSync                              = QBluetoothUuid__ServiceClassUuid(0x1104)
	QBluetoothUuid__ObexObjectPush                        = QBluetoothUuid__ServiceClassUuid(0x1105)
	QBluetoothUuid__OBEXFileTransfer                      = QBluetoothUuid__ServiceClassUuid(0x1106)
	QBluetoothUuid__IrMCSyncCommand                       = QBluetoothUuid__ServiceClassUuid(0x1107)
	QBluetoothUuid__Headset                               = QBluetoothUuid__ServiceClassUuid(0x1108)
	QBluetoothUuid__AudioSource                           = QBluetoothUuid__ServiceClassUuid(0x110a)
	QBluetoothUuid__AudioSink                             = QBluetoothUuid__ServiceClassUuid(0x110b)
	QBluetoothUuid__AV_RemoteControlTarget                = QBluetoothUuid__ServiceClassUuid(0x110c)
	QBluetoothUuid__AdvancedAudioDistribution             = QBluetoothUuid__ServiceClassUuid(0x110d)
	QBluetoothUuid__AV_RemoteControl                      = QBluetoothUuid__ServiceClassUuid(0x110e)
	QBluetoothUuid__AV_RemoteControlController            = QBluetoothUuid__ServiceClassUuid(0x110f)
	QBluetoothUuid__HeadsetAG                             = QBluetoothUuid__ServiceClassUuid(0x1112)
	QBluetoothUuid__PANU                                  = QBluetoothUuid__ServiceClassUuid(0x1115)
	QBluetoothUuid__NAP                                   = QBluetoothUuid__ServiceClassUuid(0x1116)
	QBluetoothUuid__GN                                    = QBluetoothUuid__ServiceClassUuid(0x1117)
	QBluetoothUuid__DirectPrinting                        = QBluetoothUuid__ServiceClassUuid(0x1118)
	QBluetoothUuid__ReferencePrinting                     = QBluetoothUuid__ServiceClassUuid(0x1119)
	QBluetoothUuid__BasicImage                            = QBluetoothUuid__ServiceClassUuid(0x111a)
	QBluetoothUuid__ImagingResponder                      = QBluetoothUuid__ServiceClassUuid(0x111b)
	QBluetoothUuid__ImagingAutomaticArchive               = QBluetoothUuid__ServiceClassUuid(0x111c)
	QBluetoothUuid__ImagingReferenceObjects               = QBluetoothUuid__ServiceClassUuid(0x111d)
	QBluetoothUuid__Handsfree                             = QBluetoothUuid__ServiceClassUuid(0x111e)
	QBluetoothUuid__HandsfreeAudioGateway                 = QBluetoothUuid__ServiceClassUuid(0x111f)
	QBluetoothUuid__DirectPrintingReferenceObjectsService = QBluetoothUuid__ServiceClassUuid(0x1120)
	QBluetoothUuid__ReflectedUI                           = QBluetoothUuid__ServiceClassUuid(0x1121)
	QBluetoothUuid__BasicPrinting                         = QBluetoothUuid__ServiceClassUuid(0x1122)
	QBluetoothUuid__PrintingStatus                        = QBluetoothUuid__ServiceClassUuid(0x1123)
	QBluetoothUuid__HumanInterfaceDeviceService           = QBluetoothUuid__ServiceClassUuid(0x1124)
	QBluetoothUuid__HardcopyCableReplacement              = QBluetoothUuid__ServiceClassUuid(0x1125)
	QBluetoothUuid__HCRPrint                              = QBluetoothUuid__ServiceClassUuid(0x1126)
	QBluetoothUuid__HCRScan                               = QBluetoothUuid__ServiceClassUuid(0x1127)
	QBluetoothUuid__SIMAccess                             = QBluetoothUuid__ServiceClassUuid(0x112d)
	QBluetoothUuid__PhonebookAccessPCE                    = QBluetoothUuid__ServiceClassUuid(0x112e)
	QBluetoothUuid__PhonebookAccessPSE                    = QBluetoothUuid__ServiceClassUuid(0x112f)
	QBluetoothUuid__PhonebookAccess                       = QBluetoothUuid__ServiceClassUuid(0x1130)
	QBluetoothUuid__HeadsetHS                             = QBluetoothUuid__ServiceClassUuid(0x1131)
	QBluetoothUuid__MessageAccessServer                   = QBluetoothUuid__ServiceClassUuid(0x1132)
	QBluetoothUuid__MessageNotificationServer             = QBluetoothUuid__ServiceClassUuid(0x1133)
	QBluetoothUuid__MessageAccessProfile                  = QBluetoothUuid__ServiceClassUuid(0x1134)
	QBluetoothUuid__GNSS                                  = QBluetoothUuid__ServiceClassUuid(0x1135)
	QBluetoothUuid__GNSSServer                            = QBluetoothUuid__ServiceClassUuid(0x1136)
	QBluetoothUuid__Display3D                             = QBluetoothUuid__ServiceClassUuid(0x1137)
	QBluetoothUuid__Glasses3D                             = QBluetoothUuid__ServiceClassUuid(0x1138)
	QBluetoothUuid__Synchronization3D                     = QBluetoothUuid__ServiceClassUuid(0x1139)
	QBluetoothUuid__MPSProfile                            = QBluetoothUuid__ServiceClassUuid(0x113a)
	QBluetoothUuid__MPSService                            = QBluetoothUuid__ServiceClassUuid(0x113b)
	QBluetoothUuid__PnPInformation                        = QBluetoothUuid__ServiceClassUuid(0x1200)
	QBluetoothUuid__GenericNetworking                     = QBluetoothUuid__ServiceClassUuid(0x1201)
	QBluetoothUuid__GenericFileTransfer                   = QBluetoothUuid__ServiceClassUuid(0x1202)
	QBluetoothUuid__GenericAudio                          = QBluetoothUuid__ServiceClassUuid(0x1203)
	QBluetoothUuid__GenericTelephony                      = QBluetoothUuid__ServiceClassUuid(0x1204)
	QBluetoothUuid__VideoSource                           = QBluetoothUuid__ServiceClassUuid(0x1303)
	QBluetoothUuid__VideoSink                             = QBluetoothUuid__ServiceClassUuid(0x1304)
	QBluetoothUuid__VideoDistribution                     = QBluetoothUuid__ServiceClassUuid(0x1305)
	QBluetoothUuid__HDP                                   = QBluetoothUuid__ServiceClassUuid(0x1400)
	QBluetoothUuid__HDPSource                             = QBluetoothUuid__ServiceClassUuid(0x1401)
	QBluetoothUuid__HDPSink                               = QBluetoothUuid__ServiceClassUuid(0x1402)
	QBluetoothUuid__GenericAccess                         = QBluetoothUuid__ServiceClassUuid(0x1800)
	QBluetoothUuid__GenericAttribute                      = QBluetoothUuid__ServiceClassUuid(0x1801)
	QBluetoothUuid__ImmediateAlert                        = QBluetoothUuid__ServiceClassUuid(0x1802)
	QBluetoothUuid__LinkLoss                              = QBluetoothUuid__ServiceClassUuid(0x1803)
	QBluetoothUuid__TxPower                               = QBluetoothUuid__ServiceClassUuid(0x1804)
	QBluetoothUuid__CurrentTimeService                    = QBluetoothUuid__ServiceClassUuid(0x1805)
	QBluetoothUuid__ReferenceTimeUpdateService            = QBluetoothUuid__ServiceClassUuid(0x1806)
	QBluetoothUuid__NextDSTChangeService                  = QBluetoothUuid__ServiceClassUuid(0x1807)
	QBluetoothUuid__Glucose                               = QBluetoothUuid__ServiceClassUuid(0x1808)
	QBluetoothUuid__HealthThermometer                     = QBluetoothUuid__ServiceClassUuid(0x1809)
	QBluetoothUuid__DeviceInformation                     = QBluetoothUuid__ServiceClassUuid(0x180a)
	QBluetoothUuid__HeartRate                             = QBluetoothUuid__ServiceClassUuid(0x180d)
	QBluetoothUuid__PhoneAlertStatusService               = QBluetoothUuid__ServiceClassUuid(0x180e)
	QBluetoothUuid__BatteryService                        = QBluetoothUuid__ServiceClassUuid(0x180f)
	QBluetoothUuid__BloodPressure                         = QBluetoothUuid__ServiceClassUuid(0x1810)
	QBluetoothUuid__AlertNotificationService              = QBluetoothUuid__ServiceClassUuid(0x1811)
	QBluetoothUuid__HumanInterfaceDevice                  = QBluetoothUuid__ServiceClassUuid(0x1812)
	QBluetoothUuid__ScanParameters                        = QBluetoothUuid__ServiceClassUuid(0x1813)
	QBluetoothUuid__RunningSpeedAndCadence                = QBluetoothUuid__ServiceClassUuid(0x1814)
	QBluetoothUuid__CyclingSpeedAndCadence                = QBluetoothUuid__ServiceClassUuid(0x1816)
	QBluetoothUuid__CyclingPower                          = QBluetoothUuid__ServiceClassUuid(0x1818)
	QBluetoothUuid__LocationAndNavigation                 = QBluetoothUuid__ServiceClassUuid(0x1819)
	QBluetoothUuid__EnvironmentalSensing                  = QBluetoothUuid__ServiceClassUuid(0x181a)
	QBluetoothUuid__BodyComposition                       = QBluetoothUuid__ServiceClassUuid(0x181b)
	QBluetoothUuid__UserData                              = QBluetoothUuid__ServiceClassUuid(0x181c)
	QBluetoothUuid__WeightScale                           = QBluetoothUuid__ServiceClassUuid(0x181d)
	QBluetoothUuid__BondManagement                        = QBluetoothUuid__ServiceClassUuid(0x181e)
	QBluetoothUuid__ContinuousGlucoseMonitoring           = QBluetoothUuid__ServiceClassUuid(0x181f)
)

type QBluetoothUuid struct {
	core.QUuid
}

type QBluetoothUuid_ITF interface {
	core.QUuid_ITF
	QBluetoothUuid_PTR() *QBluetoothUuid
}

func (p *QBluetoothUuid) QBluetoothUuid_PTR() *QBluetoothUuid {
	return p
}

func (p *QBluetoothUuid) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QUuid_PTR().Pointer()
	}
	return nil
}

func (p *QBluetoothUuid) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QUuid_PTR().SetPointer(ptr)
	}
}

func PointerFromQBluetoothUuid(ptr QBluetoothUuid_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothUuid_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothUuidFromPointer(ptr unsafe.Pointer) *QBluetoothUuid {
	var n = new(QBluetoothUuid)
	n.SetPointer(ptr)
	return n
}

func newQBluetoothUuidFromPointer(ptr unsafe.Pointer) *QBluetoothUuid {
	var n = NewQBluetoothUuidFromPointer(ptr)
	return n
}

func NewQBluetoothUuid() *QBluetoothUuid {
	defer qt.Recovering("QBluetoothUuid::QBluetoothUuid")

	return newQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid())
}

func NewQBluetoothUuid4(uuid QBluetoothUuid__CharacteristicType) *QBluetoothUuid {
	defer qt.Recovering("QBluetoothUuid::QBluetoothUuid")

	return newQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid4(C.int(uuid)))
}

func NewQBluetoothUuid5(uuid QBluetoothUuid__DescriptorType) *QBluetoothUuid {
	defer qt.Recovering("QBluetoothUuid::QBluetoothUuid")

	return newQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid5(C.int(uuid)))
}

func NewQBluetoothUuid2(uuid QBluetoothUuid__ProtocolUuid) *QBluetoothUuid {
	defer qt.Recovering("QBluetoothUuid::QBluetoothUuid")

	return newQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid2(C.int(uuid)))
}

func NewQBluetoothUuid3(uuid QBluetoothUuid__ServiceClassUuid) *QBluetoothUuid {
	defer qt.Recovering("QBluetoothUuid::QBluetoothUuid")

	return newQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid3(C.int(uuid)))
}

func NewQBluetoothUuid10(uuid QBluetoothUuid_ITF) *QBluetoothUuid {
	defer qt.Recovering("QBluetoothUuid::QBluetoothUuid")

	return newQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid10(PointerFromQBluetoothUuid(uuid)))
}

func NewQBluetoothUuid9(uuid string) *QBluetoothUuid {
	defer qt.Recovering("QBluetoothUuid::QBluetoothUuid")

	var uuidC = C.CString(uuid)
	defer C.free(unsafe.Pointer(uuidC))
	return newQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid9(uuidC))
}

func NewQBluetoothUuid11(uuid core.QUuid_ITF) *QBluetoothUuid {
	defer qt.Recovering("QBluetoothUuid::QBluetoothUuid")

	return newQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid11(core.PointerFromQUuid(uuid)))
}

func QBluetoothUuid_CharacteristicToString(uuid QBluetoothUuid__CharacteristicType) string {
	defer qt.Recovering("QBluetoothUuid::characteristicToString")

	return C.GoString(C.QBluetoothUuid_QBluetoothUuid_CharacteristicToString(C.int(uuid)))
}

func (ptr *QBluetoothUuid) CharacteristicToString(uuid QBluetoothUuid__CharacteristicType) string {
	defer qt.Recovering("QBluetoothUuid::characteristicToString")

	return C.GoString(C.QBluetoothUuid_QBluetoothUuid_CharacteristicToString(C.int(uuid)))
}

func QBluetoothUuid_DescriptorToString(uuid QBluetoothUuid__DescriptorType) string {
	defer qt.Recovering("QBluetoothUuid::descriptorToString")

	return C.GoString(C.QBluetoothUuid_QBluetoothUuid_DescriptorToString(C.int(uuid)))
}

func (ptr *QBluetoothUuid) DescriptorToString(uuid QBluetoothUuid__DescriptorType) string {
	defer qt.Recovering("QBluetoothUuid::descriptorToString")

	return C.GoString(C.QBluetoothUuid_QBluetoothUuid_DescriptorToString(C.int(uuid)))
}

func (ptr *QBluetoothUuid) MinimumSize() int {
	defer qt.Recovering("QBluetoothUuid::minimumSize")

	if ptr.Pointer() != nil {
		return int(C.QBluetoothUuid_MinimumSize(ptr.Pointer()))
	}
	return 0
}

func QBluetoothUuid_ProtocolToString(uuid QBluetoothUuid__ProtocolUuid) string {
	defer qt.Recovering("QBluetoothUuid::protocolToString")

	return C.GoString(C.QBluetoothUuid_QBluetoothUuid_ProtocolToString(C.int(uuid)))
}

func (ptr *QBluetoothUuid) ProtocolToString(uuid QBluetoothUuid__ProtocolUuid) string {
	defer qt.Recovering("QBluetoothUuid::protocolToString")

	return C.GoString(C.QBluetoothUuid_QBluetoothUuid_ProtocolToString(C.int(uuid)))
}

func QBluetoothUuid_ServiceClassToString(uuid QBluetoothUuid__ServiceClassUuid) string {
	defer qt.Recovering("QBluetoothUuid::serviceClassToString")

	return C.GoString(C.QBluetoothUuid_QBluetoothUuid_ServiceClassToString(C.int(uuid)))
}

func (ptr *QBluetoothUuid) ServiceClassToString(uuid QBluetoothUuid__ServiceClassUuid) string {
	defer qt.Recovering("QBluetoothUuid::serviceClassToString")

	return C.GoString(C.QBluetoothUuid_QBluetoothUuid_ServiceClassToString(C.int(uuid)))
}

func (ptr *QBluetoothUuid) DestroyQBluetoothUuid() {
	defer qt.Recovering("QBluetoothUuid::~QBluetoothUuid")

	if ptr.Pointer() != nil {
		C.QBluetoothUuid_DestroyQBluetoothUuid(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//QLowEnergyAdvertisingData::Discoverability
type QLowEnergyAdvertisingData__Discoverability int64

const (
	QLowEnergyAdvertisingData__DiscoverabilityNone    = QLowEnergyAdvertisingData__Discoverability(0)
	QLowEnergyAdvertisingData__DiscoverabilityLimited = QLowEnergyAdvertisingData__Discoverability(1)
	QLowEnergyAdvertisingData__DiscoverabilityGeneral = QLowEnergyAdvertisingData__Discoverability(2)
)

type QLowEnergyAdvertisingData struct {
	ptr unsafe.Pointer
}

type QLowEnergyAdvertisingData_ITF interface {
	QLowEnergyAdvertisingData_PTR() *QLowEnergyAdvertisingData
}

func (p *QLowEnergyAdvertisingData) QLowEnergyAdvertisingData_PTR() *QLowEnergyAdvertisingData {
	return p
}

func (p *QLowEnergyAdvertisingData) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QLowEnergyAdvertisingData) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQLowEnergyAdvertisingData(ptr QLowEnergyAdvertisingData_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLowEnergyAdvertisingData_PTR().Pointer()
	}
	return nil
}

func NewQLowEnergyAdvertisingDataFromPointer(ptr unsafe.Pointer) *QLowEnergyAdvertisingData {
	var n = new(QLowEnergyAdvertisingData)
	n.SetPointer(ptr)
	return n
}

func newQLowEnergyAdvertisingDataFromPointer(ptr unsafe.Pointer) *QLowEnergyAdvertisingData {
	var n = NewQLowEnergyAdvertisingDataFromPointer(ptr)
	return n
}

func NewQLowEnergyAdvertisingData() *QLowEnergyAdvertisingData {
	defer qt.Recovering("QLowEnergyAdvertisingData::QLowEnergyAdvertisingData")

	return newQLowEnergyAdvertisingDataFromPointer(C.QLowEnergyAdvertisingData_NewQLowEnergyAdvertisingData())
}

func NewQLowEnergyAdvertisingData2(other QLowEnergyAdvertisingData_ITF) *QLowEnergyAdvertisingData {
	defer qt.Recovering("QLowEnergyAdvertisingData::QLowEnergyAdvertisingData")

	return newQLowEnergyAdvertisingDataFromPointer(C.QLowEnergyAdvertisingData_NewQLowEnergyAdvertisingData2(PointerFromQLowEnergyAdvertisingData(other)))
}

func (ptr *QLowEnergyAdvertisingData) Discoverability() QLowEnergyAdvertisingData__Discoverability {
	defer qt.Recovering("QLowEnergyAdvertisingData::discoverability")

	if ptr.Pointer() != nil {
		return QLowEnergyAdvertisingData__Discoverability(C.QLowEnergyAdvertisingData_Discoverability(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyAdvertisingData) IncludePowerLevel() bool {
	defer qt.Recovering("QLowEnergyAdvertisingData::includePowerLevel")

	if ptr.Pointer() != nil {
		return C.QLowEnergyAdvertisingData_IncludePowerLevel(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLowEnergyAdvertisingData) LocalName() string {
	defer qt.Recovering("QLowEnergyAdvertisingData::localName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QLowEnergyAdvertisingData_LocalName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLowEnergyAdvertisingData) ManufacturerData() string {
	defer qt.Recovering("QLowEnergyAdvertisingData::manufacturerData")

	if ptr.Pointer() != nil {
		return qt.HexDecodeToString(C.GoString(C.QLowEnergyAdvertisingData_ManufacturerData(ptr.Pointer())))
	}
	return ""
}

func (ptr *QLowEnergyAdvertisingData) RawData() string {
	defer qt.Recovering("QLowEnergyAdvertisingData::rawData")

	if ptr.Pointer() != nil {
		return qt.HexDecodeToString(C.GoString(C.QLowEnergyAdvertisingData_RawData(ptr.Pointer())))
	}
	return ""
}

func (ptr *QLowEnergyAdvertisingData) SetDiscoverability(mode QLowEnergyAdvertisingData__Discoverability) {
	defer qt.Recovering("QLowEnergyAdvertisingData::setDiscoverability")

	if ptr.Pointer() != nil {
		C.QLowEnergyAdvertisingData_SetDiscoverability(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QLowEnergyAdvertisingData) SetIncludePowerLevel(doInclude bool) {
	defer qt.Recovering("QLowEnergyAdvertisingData::setIncludePowerLevel")

	if ptr.Pointer() != nil {
		C.QLowEnergyAdvertisingData_SetIncludePowerLevel(ptr.Pointer(), C.int(qt.GoBoolToInt(doInclude)))
	}
}

func (ptr *QLowEnergyAdvertisingData) SetLocalName(name string) {
	defer qt.Recovering("QLowEnergyAdvertisingData::setLocalName")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QLowEnergyAdvertisingData_SetLocalName(ptr.Pointer(), nameC)
	}
}

func (ptr *QLowEnergyAdvertisingData) SetRawData(data string) {
	defer qt.Recovering("QLowEnergyAdvertisingData::setRawData")

	if ptr.Pointer() != nil {
		var dataC = C.CString(hex.EncodeToString([]byte(data)))
		defer C.free(unsafe.Pointer(dataC))
		C.QLowEnergyAdvertisingData_SetRawData(ptr.Pointer(), dataC)
	}
}

func (ptr *QLowEnergyAdvertisingData) Swap(other QLowEnergyAdvertisingData_ITF) {
	defer qt.Recovering("QLowEnergyAdvertisingData::swap")

	if ptr.Pointer() != nil {
		C.QLowEnergyAdvertisingData_Swap(ptr.Pointer(), PointerFromQLowEnergyAdvertisingData(other))
	}
}

func (ptr *QLowEnergyAdvertisingData) DestroyQLowEnergyAdvertisingData() {
	defer qt.Recovering("QLowEnergyAdvertisingData::~QLowEnergyAdvertisingData")

	if ptr.Pointer() != nil {
		C.QLowEnergyAdvertisingData_DestroyQLowEnergyAdvertisingData(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//QLowEnergyAdvertisingParameters::FilterPolicy
type QLowEnergyAdvertisingParameters__FilterPolicy int64

const (
	QLowEnergyAdvertisingParameters__IgnoreWhiteList                      = QLowEnergyAdvertisingParameters__FilterPolicy(0x00)
	QLowEnergyAdvertisingParameters__UseWhiteListForScanning              = QLowEnergyAdvertisingParameters__FilterPolicy(0x01)
	QLowEnergyAdvertisingParameters__UseWhiteListForConnecting            = QLowEnergyAdvertisingParameters__FilterPolicy(0x02)
	QLowEnergyAdvertisingParameters__UseWhiteListForScanningAndConnecting = QLowEnergyAdvertisingParameters__FilterPolicy(0x03)
)

//QLowEnergyAdvertisingParameters::Mode
type QLowEnergyAdvertisingParameters__Mode int64

const (
	QLowEnergyAdvertisingParameters__AdvInd        = QLowEnergyAdvertisingParameters__Mode(0x0)
	QLowEnergyAdvertisingParameters__AdvScanInd    = QLowEnergyAdvertisingParameters__Mode(0x2)
	QLowEnergyAdvertisingParameters__AdvNonConnInd = QLowEnergyAdvertisingParameters__Mode(0x3)
)

type QLowEnergyAdvertisingParameters struct {
	ptr unsafe.Pointer
}

type QLowEnergyAdvertisingParameters_ITF interface {
	QLowEnergyAdvertisingParameters_PTR() *QLowEnergyAdvertisingParameters
}

func (p *QLowEnergyAdvertisingParameters) QLowEnergyAdvertisingParameters_PTR() *QLowEnergyAdvertisingParameters {
	return p
}

func (p *QLowEnergyAdvertisingParameters) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QLowEnergyAdvertisingParameters) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQLowEnergyAdvertisingParameters(ptr QLowEnergyAdvertisingParameters_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLowEnergyAdvertisingParameters_PTR().Pointer()
	}
	return nil
}

func NewQLowEnergyAdvertisingParametersFromPointer(ptr unsafe.Pointer) *QLowEnergyAdvertisingParameters {
	var n = new(QLowEnergyAdvertisingParameters)
	n.SetPointer(ptr)
	return n
}

func newQLowEnergyAdvertisingParametersFromPointer(ptr unsafe.Pointer) *QLowEnergyAdvertisingParameters {
	var n = NewQLowEnergyAdvertisingParametersFromPointer(ptr)
	return n
}

func NewQLowEnergyAdvertisingParameters() *QLowEnergyAdvertisingParameters {
	defer qt.Recovering("QLowEnergyAdvertisingParameters::QLowEnergyAdvertisingParameters")

	return newQLowEnergyAdvertisingParametersFromPointer(C.QLowEnergyAdvertisingParameters_NewQLowEnergyAdvertisingParameters())
}

func NewQLowEnergyAdvertisingParameters2(other QLowEnergyAdvertisingParameters_ITF) *QLowEnergyAdvertisingParameters {
	defer qt.Recovering("QLowEnergyAdvertisingParameters::QLowEnergyAdvertisingParameters")

	return newQLowEnergyAdvertisingParametersFromPointer(C.QLowEnergyAdvertisingParameters_NewQLowEnergyAdvertisingParameters2(PointerFromQLowEnergyAdvertisingParameters(other)))
}

func (ptr *QLowEnergyAdvertisingParameters) FilterPolicy() QLowEnergyAdvertisingParameters__FilterPolicy {
	defer qt.Recovering("QLowEnergyAdvertisingParameters::filterPolicy")

	if ptr.Pointer() != nil {
		return QLowEnergyAdvertisingParameters__FilterPolicy(C.QLowEnergyAdvertisingParameters_FilterPolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyAdvertisingParameters) MaximumInterval() int {
	defer qt.Recovering("QLowEnergyAdvertisingParameters::maximumInterval")

	if ptr.Pointer() != nil {
		return int(C.QLowEnergyAdvertisingParameters_MaximumInterval(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyAdvertisingParameters) MinimumInterval() int {
	defer qt.Recovering("QLowEnergyAdvertisingParameters::minimumInterval")

	if ptr.Pointer() != nil {
		return int(C.QLowEnergyAdvertisingParameters_MinimumInterval(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyAdvertisingParameters) Mode() QLowEnergyAdvertisingParameters__Mode {
	defer qt.Recovering("QLowEnergyAdvertisingParameters::mode")

	if ptr.Pointer() != nil {
		return QLowEnergyAdvertisingParameters__Mode(C.QLowEnergyAdvertisingParameters_Mode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyAdvertisingParameters) SetMode(mode QLowEnergyAdvertisingParameters__Mode) {
	defer qt.Recovering("QLowEnergyAdvertisingParameters::setMode")

	if ptr.Pointer() != nil {
		C.QLowEnergyAdvertisingParameters_SetMode(ptr.Pointer(), C.int(mode))
	}
}

func (ptr *QLowEnergyAdvertisingParameters) Swap(other QLowEnergyAdvertisingParameters_ITF) {
	defer qt.Recovering("QLowEnergyAdvertisingParameters::swap")

	if ptr.Pointer() != nil {
		C.QLowEnergyAdvertisingParameters_Swap(ptr.Pointer(), PointerFromQLowEnergyAdvertisingParameters(other))
	}
}

func (ptr *QLowEnergyAdvertisingParameters) DestroyQLowEnergyAdvertisingParameters() {
	defer qt.Recovering("QLowEnergyAdvertisingParameters::~QLowEnergyAdvertisingParameters")

	if ptr.Pointer() != nil {
		C.QLowEnergyAdvertisingParameters_DestroyQLowEnergyAdvertisingParameters(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//QLowEnergyCharacteristic::PropertyType
type QLowEnergyCharacteristic__PropertyType int64

const (
	QLowEnergyCharacteristic__Unknown          = QLowEnergyCharacteristic__PropertyType(0x00)
	QLowEnergyCharacteristic__Broadcasting     = QLowEnergyCharacteristic__PropertyType(0x01)
	QLowEnergyCharacteristic__Read             = QLowEnergyCharacteristic__PropertyType(0x02)
	QLowEnergyCharacteristic__WriteNoResponse  = QLowEnergyCharacteristic__PropertyType(0x04)
	QLowEnergyCharacteristic__Write            = QLowEnergyCharacteristic__PropertyType(0x08)
	QLowEnergyCharacteristic__Notify           = QLowEnergyCharacteristic__PropertyType(0x10)
	QLowEnergyCharacteristic__Indicate         = QLowEnergyCharacteristic__PropertyType(0x20)
	QLowEnergyCharacteristic__WriteSigned      = QLowEnergyCharacteristic__PropertyType(0x40)
	QLowEnergyCharacteristic__ExtendedProperty = QLowEnergyCharacteristic__PropertyType(0x80)
)

type QLowEnergyCharacteristic struct {
	ptr unsafe.Pointer
}

type QLowEnergyCharacteristic_ITF interface {
	QLowEnergyCharacteristic_PTR() *QLowEnergyCharacteristic
}

func (p *QLowEnergyCharacteristic) QLowEnergyCharacteristic_PTR() *QLowEnergyCharacteristic {
	return p
}

func (p *QLowEnergyCharacteristic) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QLowEnergyCharacteristic) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQLowEnergyCharacteristic(ptr QLowEnergyCharacteristic_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLowEnergyCharacteristic_PTR().Pointer()
	}
	return nil
}

func NewQLowEnergyCharacteristicFromPointer(ptr unsafe.Pointer) *QLowEnergyCharacteristic {
	var n = new(QLowEnergyCharacteristic)
	n.SetPointer(ptr)
	return n
}

func newQLowEnergyCharacteristicFromPointer(ptr unsafe.Pointer) *QLowEnergyCharacteristic {
	var n = NewQLowEnergyCharacteristicFromPointer(ptr)
	return n
}

func NewQLowEnergyCharacteristic() *QLowEnergyCharacteristic {
	defer qt.Recovering("QLowEnergyCharacteristic::QLowEnergyCharacteristic")

	return newQLowEnergyCharacteristicFromPointer(C.QLowEnergyCharacteristic_NewQLowEnergyCharacteristic())
}

func NewQLowEnergyCharacteristic2(other QLowEnergyCharacteristic_ITF) *QLowEnergyCharacteristic {
	defer qt.Recovering("QLowEnergyCharacteristic::QLowEnergyCharacteristic")

	return newQLowEnergyCharacteristicFromPointer(C.QLowEnergyCharacteristic_NewQLowEnergyCharacteristic2(PointerFromQLowEnergyCharacteristic(other)))
}

func (ptr *QLowEnergyCharacteristic) Descriptor(uuid QBluetoothUuid_ITF) *QLowEnergyDescriptor {
	defer qt.Recovering("QLowEnergyCharacteristic::descriptor")

	if ptr.Pointer() != nil {
		var tmpValue = NewQLowEnergyDescriptorFromPointer(C.QLowEnergyCharacteristic_Descriptor(ptr.Pointer(), PointerFromQBluetoothUuid(uuid)))
		runtime.SetFinalizer(tmpValue, (*QLowEnergyDescriptor).DestroyQLowEnergyDescriptor)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyCharacteristic) IsValid() bool {
	defer qt.Recovering("QLowEnergyCharacteristic::isValid")

	if ptr.Pointer() != nil {
		return C.QLowEnergyCharacteristic_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLowEnergyCharacteristic) Name() string {
	defer qt.Recovering("QLowEnergyCharacteristic::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QLowEnergyCharacteristic_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLowEnergyCharacteristic) Properties() QLowEnergyCharacteristic__PropertyType {
	defer qt.Recovering("QLowEnergyCharacteristic::properties")

	if ptr.Pointer() != nil {
		return QLowEnergyCharacteristic__PropertyType(C.QLowEnergyCharacteristic_Properties(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyCharacteristic) Uuid() *QBluetoothUuid {
	defer qt.Recovering("QLowEnergyCharacteristic::uuid")

	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothUuidFromPointer(C.QLowEnergyCharacteristic_Uuid(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyCharacteristic) Value() string {
	defer qt.Recovering("QLowEnergyCharacteristic::value")

	if ptr.Pointer() != nil {
		return qt.HexDecodeToString(C.GoString(C.QLowEnergyCharacteristic_Value(ptr.Pointer())))
	}
	return ""
}

func (ptr *QLowEnergyCharacteristic) DestroyQLowEnergyCharacteristic() {
	defer qt.Recovering("QLowEnergyCharacteristic::~QLowEnergyCharacteristic")

	if ptr.Pointer() != nil {
		C.QLowEnergyCharacteristic_DestroyQLowEnergyCharacteristic(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QLowEnergyCharacteristicData struct {
	ptr unsafe.Pointer
}

type QLowEnergyCharacteristicData_ITF interface {
	QLowEnergyCharacteristicData_PTR() *QLowEnergyCharacteristicData
}

func (p *QLowEnergyCharacteristicData) QLowEnergyCharacteristicData_PTR() *QLowEnergyCharacteristicData {
	return p
}

func (p *QLowEnergyCharacteristicData) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QLowEnergyCharacteristicData) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQLowEnergyCharacteristicData(ptr QLowEnergyCharacteristicData_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLowEnergyCharacteristicData_PTR().Pointer()
	}
	return nil
}

func NewQLowEnergyCharacteristicDataFromPointer(ptr unsafe.Pointer) *QLowEnergyCharacteristicData {
	var n = new(QLowEnergyCharacteristicData)
	n.SetPointer(ptr)
	return n
}

func newQLowEnergyCharacteristicDataFromPointer(ptr unsafe.Pointer) *QLowEnergyCharacteristicData {
	var n = NewQLowEnergyCharacteristicDataFromPointer(ptr)
	return n
}

func NewQLowEnergyCharacteristicData() *QLowEnergyCharacteristicData {
	defer qt.Recovering("QLowEnergyCharacteristicData::QLowEnergyCharacteristicData")

	return newQLowEnergyCharacteristicDataFromPointer(C.QLowEnergyCharacteristicData_NewQLowEnergyCharacteristicData())
}

func NewQLowEnergyCharacteristicData2(other QLowEnergyCharacteristicData_ITF) *QLowEnergyCharacteristicData {
	defer qt.Recovering("QLowEnergyCharacteristicData::QLowEnergyCharacteristicData")

	return newQLowEnergyCharacteristicDataFromPointer(C.QLowEnergyCharacteristicData_NewQLowEnergyCharacteristicData2(PointerFromQLowEnergyCharacteristicData(other)))
}

func (ptr *QLowEnergyCharacteristicData) AddDescriptor(descriptor QLowEnergyDescriptorData_ITF) {
	defer qt.Recovering("QLowEnergyCharacteristicData::addDescriptor")

	if ptr.Pointer() != nil {
		C.QLowEnergyCharacteristicData_AddDescriptor(ptr.Pointer(), PointerFromQLowEnergyDescriptorData(descriptor))
	}
}

func (ptr *QLowEnergyCharacteristicData) IsValid() bool {
	defer qt.Recovering("QLowEnergyCharacteristicData::isValid")

	if ptr.Pointer() != nil {
		return C.QLowEnergyCharacteristicData_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLowEnergyCharacteristicData) MaximumValueLength() int {
	defer qt.Recovering("QLowEnergyCharacteristicData::maximumValueLength")

	if ptr.Pointer() != nil {
		return int(C.QLowEnergyCharacteristicData_MaximumValueLength(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyCharacteristicData) MinimumValueLength() int {
	defer qt.Recovering("QLowEnergyCharacteristicData::minimumValueLength")

	if ptr.Pointer() != nil {
		return int(C.QLowEnergyCharacteristicData_MinimumValueLength(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyCharacteristicData) Properties() QLowEnergyCharacteristic__PropertyType {
	defer qt.Recovering("QLowEnergyCharacteristicData::properties")

	if ptr.Pointer() != nil {
		return QLowEnergyCharacteristic__PropertyType(C.QLowEnergyCharacteristicData_Properties(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyCharacteristicData) SetProperties(properties QLowEnergyCharacteristic__PropertyType) {
	defer qt.Recovering("QLowEnergyCharacteristicData::setProperties")

	if ptr.Pointer() != nil {
		C.QLowEnergyCharacteristicData_SetProperties(ptr.Pointer(), C.int(properties))
	}
}

func (ptr *QLowEnergyCharacteristicData) SetUuid(uuid QBluetoothUuid_ITF) {
	defer qt.Recovering("QLowEnergyCharacteristicData::setUuid")

	if ptr.Pointer() != nil {
		C.QLowEnergyCharacteristicData_SetUuid(ptr.Pointer(), PointerFromQBluetoothUuid(uuid))
	}
}

func (ptr *QLowEnergyCharacteristicData) SetValue(value string) {
	defer qt.Recovering("QLowEnergyCharacteristicData::setValue")

	if ptr.Pointer() != nil {
		var valueC = C.CString(hex.EncodeToString([]byte(value)))
		defer C.free(unsafe.Pointer(valueC))
		C.QLowEnergyCharacteristicData_SetValue(ptr.Pointer(), valueC)
	}
}

func (ptr *QLowEnergyCharacteristicData) SetValueLength(minimum int, maximum int) {
	defer qt.Recovering("QLowEnergyCharacteristicData::setValueLength")

	if ptr.Pointer() != nil {
		C.QLowEnergyCharacteristicData_SetValueLength(ptr.Pointer(), C.int(minimum), C.int(maximum))
	}
}

func (ptr *QLowEnergyCharacteristicData) Swap(other QLowEnergyCharacteristicData_ITF) {
	defer qt.Recovering("QLowEnergyCharacteristicData::swap")

	if ptr.Pointer() != nil {
		C.QLowEnergyCharacteristicData_Swap(ptr.Pointer(), PointerFromQLowEnergyCharacteristicData(other))
	}
}

func (ptr *QLowEnergyCharacteristicData) Uuid() *QBluetoothUuid {
	defer qt.Recovering("QLowEnergyCharacteristicData::uuid")

	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothUuidFromPointer(C.QLowEnergyCharacteristicData_Uuid(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyCharacteristicData) Value() string {
	defer qt.Recovering("QLowEnergyCharacteristicData::value")

	if ptr.Pointer() != nil {
		return qt.HexDecodeToString(C.GoString(C.QLowEnergyCharacteristicData_Value(ptr.Pointer())))
	}
	return ""
}

func (ptr *QLowEnergyCharacteristicData) DestroyQLowEnergyCharacteristicData() {
	defer qt.Recovering("QLowEnergyCharacteristicData::~QLowEnergyCharacteristicData")

	if ptr.Pointer() != nil {
		C.QLowEnergyCharacteristicData_DestroyQLowEnergyCharacteristicData(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QLowEnergyConnectionParameters struct {
	ptr unsafe.Pointer
}

type QLowEnergyConnectionParameters_ITF interface {
	QLowEnergyConnectionParameters_PTR() *QLowEnergyConnectionParameters
}

func (p *QLowEnergyConnectionParameters) QLowEnergyConnectionParameters_PTR() *QLowEnergyConnectionParameters {
	return p
}

func (p *QLowEnergyConnectionParameters) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QLowEnergyConnectionParameters) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQLowEnergyConnectionParameters(ptr QLowEnergyConnectionParameters_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLowEnergyConnectionParameters_PTR().Pointer()
	}
	return nil
}

func NewQLowEnergyConnectionParametersFromPointer(ptr unsafe.Pointer) *QLowEnergyConnectionParameters {
	var n = new(QLowEnergyConnectionParameters)
	n.SetPointer(ptr)
	return n
}

func newQLowEnergyConnectionParametersFromPointer(ptr unsafe.Pointer) *QLowEnergyConnectionParameters {
	var n = NewQLowEnergyConnectionParametersFromPointer(ptr)
	return n
}

func NewQLowEnergyConnectionParameters() *QLowEnergyConnectionParameters {
	defer qt.Recovering("QLowEnergyConnectionParameters::QLowEnergyConnectionParameters")

	return newQLowEnergyConnectionParametersFromPointer(C.QLowEnergyConnectionParameters_NewQLowEnergyConnectionParameters())
}

func NewQLowEnergyConnectionParameters2(other QLowEnergyConnectionParameters_ITF) *QLowEnergyConnectionParameters {
	defer qt.Recovering("QLowEnergyConnectionParameters::QLowEnergyConnectionParameters")

	return newQLowEnergyConnectionParametersFromPointer(C.QLowEnergyConnectionParameters_NewQLowEnergyConnectionParameters2(PointerFromQLowEnergyConnectionParameters(other)))
}

func (ptr *QLowEnergyConnectionParameters) Latency() int {
	defer qt.Recovering("QLowEnergyConnectionParameters::latency")

	if ptr.Pointer() != nil {
		return int(C.QLowEnergyConnectionParameters_Latency(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyConnectionParameters) SetLatency(latency int) {
	defer qt.Recovering("QLowEnergyConnectionParameters::setLatency")

	if ptr.Pointer() != nil {
		C.QLowEnergyConnectionParameters_SetLatency(ptr.Pointer(), C.int(latency))
	}
}

func (ptr *QLowEnergyConnectionParameters) SetSupervisionTimeout(timeout int) {
	defer qt.Recovering("QLowEnergyConnectionParameters::setSupervisionTimeout")

	if ptr.Pointer() != nil {
		C.QLowEnergyConnectionParameters_SetSupervisionTimeout(ptr.Pointer(), C.int(timeout))
	}
}

func (ptr *QLowEnergyConnectionParameters) SupervisionTimeout() int {
	defer qt.Recovering("QLowEnergyConnectionParameters::supervisionTimeout")

	if ptr.Pointer() != nil {
		return int(C.QLowEnergyConnectionParameters_SupervisionTimeout(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyConnectionParameters) Swap(other QLowEnergyConnectionParameters_ITF) {
	defer qt.Recovering("QLowEnergyConnectionParameters::swap")

	if ptr.Pointer() != nil {
		C.QLowEnergyConnectionParameters_Swap(ptr.Pointer(), PointerFromQLowEnergyConnectionParameters(other))
	}
}

func (ptr *QLowEnergyConnectionParameters) DestroyQLowEnergyConnectionParameters() {
	defer qt.Recovering("QLowEnergyConnectionParameters::~QLowEnergyConnectionParameters")

	if ptr.Pointer() != nil {
		C.QLowEnergyConnectionParameters_DestroyQLowEnergyConnectionParameters(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//QLowEnergyController::ControllerState
type QLowEnergyController__ControllerState int64

const (
	QLowEnergyController__UnconnectedState = QLowEnergyController__ControllerState(0)
	QLowEnergyController__ConnectingState  = QLowEnergyController__ControllerState(1)
	QLowEnergyController__ConnectedState   = QLowEnergyController__ControllerState(2)
	QLowEnergyController__DiscoveringState = QLowEnergyController__ControllerState(3)
	QLowEnergyController__DiscoveredState  = QLowEnergyController__ControllerState(4)
	QLowEnergyController__ClosingState     = QLowEnergyController__ControllerState(5)
	QLowEnergyController__AdvertisingState = QLowEnergyController__ControllerState(6)
)

//QLowEnergyController::Error
type QLowEnergyController__Error int64

const (
	QLowEnergyController__NoError                      = QLowEnergyController__Error(0)
	QLowEnergyController__UnknownError                 = QLowEnergyController__Error(1)
	QLowEnergyController__UnknownRemoteDeviceError     = QLowEnergyController__Error(2)
	QLowEnergyController__NetworkError                 = QLowEnergyController__Error(3)
	QLowEnergyController__InvalidBluetoothAdapterError = QLowEnergyController__Error(4)
	QLowEnergyController__ConnectionError              = QLowEnergyController__Error(5)
	QLowEnergyController__AdvertisingError             = QLowEnergyController__Error(6)
)

//QLowEnergyController::RemoteAddressType
type QLowEnergyController__RemoteAddressType int64

const (
	QLowEnergyController__PublicAddress = QLowEnergyController__RemoteAddressType(0)
	QLowEnergyController__RandomAddress = QLowEnergyController__RemoteAddressType(1)
)

//QLowEnergyController::Role
type QLowEnergyController__Role int64

const (
	QLowEnergyController__CentralRole    = QLowEnergyController__Role(0)
	QLowEnergyController__PeripheralRole = QLowEnergyController__Role(1)
)

type QLowEnergyController struct {
	core.QObject
}

type QLowEnergyController_ITF interface {
	core.QObject_ITF
	QLowEnergyController_PTR() *QLowEnergyController
}

func (p *QLowEnergyController) QLowEnergyController_PTR() *QLowEnergyController {
	return p
}

func (p *QLowEnergyController) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QLowEnergyController) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQLowEnergyController(ptr QLowEnergyController_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLowEnergyController_PTR().Pointer()
	}
	return nil
}

func NewQLowEnergyControllerFromPointer(ptr unsafe.Pointer) *QLowEnergyController {
	var n = new(QLowEnergyController)
	n.SetPointer(ptr)
	return n
}

func newQLowEnergyControllerFromPointer(ptr unsafe.Pointer) *QLowEnergyController {
	var n = NewQLowEnergyControllerFromPointer(ptr)
	for len(n.ObjectName()) < len("QLowEnergyController_") {
		n.SetObjectName("QLowEnergyController_" + qt.Identifier())
	}
	return n
}

//export callbackQLowEnergyController_Connected
func callbackQLowEnergyController_Connected(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QLowEnergyController::connected")

	if signal := qt.GetSignal(C.GoString(ptrName), "connected"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QLowEnergyController) ConnectConnected(f func()) {
	defer qt.Recovering("connect QLowEnergyController::connected")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_ConnectConnected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "connected", f)
	}
}

func (ptr *QLowEnergyController) DisconnectConnected() {
	defer qt.Recovering("disconnect QLowEnergyController::connected")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_DisconnectConnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "connected")
	}
}

func (ptr *QLowEnergyController) Connected() {
	defer qt.Recovering("QLowEnergyController::connected")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_Connected(ptr.Pointer())
	}
}

//export callbackQLowEnergyController_ConnectionUpdated
func callbackQLowEnergyController_ConnectionUpdated(ptr unsafe.Pointer, ptrName *C.char, newParameters unsafe.Pointer) {
	defer qt.Recovering("callback QLowEnergyController::connectionUpdated")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectionUpdated"); signal != nil {
		signal.(func(*QLowEnergyConnectionParameters))(NewQLowEnergyConnectionParametersFromPointer(newParameters))
	}

}

func (ptr *QLowEnergyController) ConnectConnectionUpdated(f func(newParameters *QLowEnergyConnectionParameters)) {
	defer qt.Recovering("connect QLowEnergyController::connectionUpdated")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_ConnectConnectionUpdated(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "connectionUpdated", f)
	}
}

func (ptr *QLowEnergyController) DisconnectConnectionUpdated() {
	defer qt.Recovering("disconnect QLowEnergyController::connectionUpdated")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_DisconnectConnectionUpdated(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "connectionUpdated")
	}
}

func (ptr *QLowEnergyController) ConnectionUpdated(newParameters QLowEnergyConnectionParameters_ITF) {
	defer qt.Recovering("QLowEnergyController::connectionUpdated")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_ConnectionUpdated(ptr.Pointer(), PointerFromQLowEnergyConnectionParameters(newParameters))
	}
}

//export callbackQLowEnergyController_Disconnected
func callbackQLowEnergyController_Disconnected(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QLowEnergyController::disconnected")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnected"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QLowEnergyController) ConnectDisconnected(f func()) {
	defer qt.Recovering("connect QLowEnergyController::disconnected")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_ConnectDisconnected(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "disconnected", f)
	}
}

func (ptr *QLowEnergyController) DisconnectDisconnected() {
	defer qt.Recovering("disconnect QLowEnergyController::disconnected")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_DisconnectDisconnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "disconnected")
	}
}

func (ptr *QLowEnergyController) Disconnected() {
	defer qt.Recovering("QLowEnergyController::disconnected")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_Disconnected(ptr.Pointer())
	}
}

//export callbackQLowEnergyController_DiscoveryFinished
func callbackQLowEnergyController_DiscoveryFinished(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QLowEnergyController::discoveryFinished")

	if signal := qt.GetSignal(C.GoString(ptrName), "discoveryFinished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QLowEnergyController) ConnectDiscoveryFinished(f func()) {
	defer qt.Recovering("connect QLowEnergyController::discoveryFinished")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_ConnectDiscoveryFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "discoveryFinished", f)
	}
}

func (ptr *QLowEnergyController) DisconnectDiscoveryFinished() {
	defer qt.Recovering("disconnect QLowEnergyController::discoveryFinished")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_DisconnectDiscoveryFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "discoveryFinished")
	}
}

func (ptr *QLowEnergyController) DiscoveryFinished() {
	defer qt.Recovering("QLowEnergyController::discoveryFinished")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_DiscoveryFinished(ptr.Pointer())
	}
}

//export callbackQLowEnergyController_Error2
func callbackQLowEnergyController_Error2(ptr unsafe.Pointer, ptrName *C.char, newError C.int) {
	defer qt.Recovering("callback QLowEnergyController::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(QLowEnergyController__Error))(QLowEnergyController__Error(newError))
	}

}

func (ptr *QLowEnergyController) ConnectError2(f func(newError QLowEnergyController__Error)) {
	defer qt.Recovering("connect QLowEnergyController::error")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error2", f)
	}
}

func (ptr *QLowEnergyController) DisconnectError2() {
	defer qt.Recovering("disconnect QLowEnergyController::error")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error2")
	}
}

func (ptr *QLowEnergyController) Error2(newError QLowEnergyController__Error) {
	defer qt.Recovering("QLowEnergyController::error")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_Error2(ptr.Pointer(), C.int(newError))
	}
}

//export callbackQLowEnergyController_ServiceDiscovered
func callbackQLowEnergyController_ServiceDiscovered(ptr unsafe.Pointer, ptrName *C.char, newService unsafe.Pointer) {
	defer qt.Recovering("callback QLowEnergyController::serviceDiscovered")

	if signal := qt.GetSignal(C.GoString(ptrName), "serviceDiscovered"); signal != nil {
		signal.(func(*QBluetoothUuid))(NewQBluetoothUuidFromPointer(newService))
	}

}

func (ptr *QLowEnergyController) ConnectServiceDiscovered(f func(newService *QBluetoothUuid)) {
	defer qt.Recovering("connect QLowEnergyController::serviceDiscovered")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_ConnectServiceDiscovered(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "serviceDiscovered", f)
	}
}

func (ptr *QLowEnergyController) DisconnectServiceDiscovered() {
	defer qt.Recovering("disconnect QLowEnergyController::serviceDiscovered")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_DisconnectServiceDiscovered(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "serviceDiscovered")
	}
}

func (ptr *QLowEnergyController) ServiceDiscovered(newService QBluetoothUuid_ITF) {
	defer qt.Recovering("QLowEnergyController::serviceDiscovered")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_ServiceDiscovered(ptr.Pointer(), PointerFromQBluetoothUuid(newService))
	}
}

//export callbackQLowEnergyController_StateChanged
func callbackQLowEnergyController_StateChanged(ptr unsafe.Pointer, ptrName *C.char, state C.int) {
	defer qt.Recovering("callback QLowEnergyController::stateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stateChanged"); signal != nil {
		signal.(func(QLowEnergyController__ControllerState))(QLowEnergyController__ControllerState(state))
	}

}

func (ptr *QLowEnergyController) ConnectStateChanged(f func(state QLowEnergyController__ControllerState)) {
	defer qt.Recovering("connect QLowEnergyController::stateChanged")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QLowEnergyController) DisconnectStateChanged() {
	defer qt.Recovering("disconnect QLowEnergyController::stateChanged")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

func (ptr *QLowEnergyController) StateChanged(state QLowEnergyController__ControllerState) {
	defer qt.Recovering("QLowEnergyController::stateChanged")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_StateChanged(ptr.Pointer(), C.int(state))
	}
}

func (ptr *QLowEnergyController) AddService(service QLowEnergyServiceData_ITF, parent core.QObject_ITF) *QLowEnergyService {
	defer qt.Recovering("QLowEnergyController::addService")

	if ptr.Pointer() != nil {
		return NewQLowEnergyServiceFromPointer(C.QLowEnergyController_AddService(ptr.Pointer(), PointerFromQLowEnergyServiceData(service), core.PointerFromQObject(parent)))
	}
	return nil
}

func (ptr *QLowEnergyController) ConnectToDevice() {
	defer qt.Recovering("QLowEnergyController::connectToDevice")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_ConnectToDevice(ptr.Pointer())
	}
}

func QLowEnergyController_CreateCentral(remoteDevice QBluetoothDeviceInfo_ITF, parent core.QObject_ITF) *QLowEnergyController {
	defer qt.Recovering("QLowEnergyController::createCentral")

	return NewQLowEnergyControllerFromPointer(C.QLowEnergyController_QLowEnergyController_CreateCentral(PointerFromQBluetoothDeviceInfo(remoteDevice), core.PointerFromQObject(parent)))
}

func (ptr *QLowEnergyController) CreateCentral(remoteDevice QBluetoothDeviceInfo_ITF, parent core.QObject_ITF) *QLowEnergyController {
	defer qt.Recovering("QLowEnergyController::createCentral")

	return NewQLowEnergyControllerFromPointer(C.QLowEnergyController_QLowEnergyController_CreateCentral(PointerFromQBluetoothDeviceInfo(remoteDevice), core.PointerFromQObject(parent)))
}

func QLowEnergyController_CreatePeripheral(parent core.QObject_ITF) *QLowEnergyController {
	defer qt.Recovering("QLowEnergyController::createPeripheral")

	return NewQLowEnergyControllerFromPointer(C.QLowEnergyController_QLowEnergyController_CreatePeripheral(core.PointerFromQObject(parent)))
}

func (ptr *QLowEnergyController) CreatePeripheral(parent core.QObject_ITF) *QLowEnergyController {
	defer qt.Recovering("QLowEnergyController::createPeripheral")

	return NewQLowEnergyControllerFromPointer(C.QLowEnergyController_QLowEnergyController_CreatePeripheral(core.PointerFromQObject(parent)))
}

func (ptr *QLowEnergyController) CreateServiceObject(serviceUuid QBluetoothUuid_ITF, parent core.QObject_ITF) *QLowEnergyService {
	defer qt.Recovering("QLowEnergyController::createServiceObject")

	if ptr.Pointer() != nil {
		return NewQLowEnergyServiceFromPointer(C.QLowEnergyController_CreateServiceObject(ptr.Pointer(), PointerFromQBluetoothUuid(serviceUuid), core.PointerFromQObject(parent)))
	}
	return nil
}

func (ptr *QLowEnergyController) DisconnectFromDevice() {
	defer qt.Recovering("QLowEnergyController::disconnectFromDevice")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_DisconnectFromDevice(ptr.Pointer())
	}
}

func (ptr *QLowEnergyController) DiscoverServices() {
	defer qt.Recovering("QLowEnergyController::discoverServices")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_DiscoverServices(ptr.Pointer())
	}
}

func (ptr *QLowEnergyController) Error() QLowEnergyController__Error {
	defer qt.Recovering("QLowEnergyController::error")

	if ptr.Pointer() != nil {
		return QLowEnergyController__Error(C.QLowEnergyController_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyController) ErrorString() string {
	defer qt.Recovering("QLowEnergyController::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QLowEnergyController_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLowEnergyController) LocalAddress() *QBluetoothAddress {
	defer qt.Recovering("QLowEnergyController::localAddress")

	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothAddressFromPointer(C.QLowEnergyController_LocalAddress(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyController) RemoteAddress() *QBluetoothAddress {
	defer qt.Recovering("QLowEnergyController::remoteAddress")

	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothAddressFromPointer(C.QLowEnergyController_RemoteAddress(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyController) RemoteAddressType() QLowEnergyController__RemoteAddressType {
	defer qt.Recovering("QLowEnergyController::remoteAddressType")

	if ptr.Pointer() != nil {
		return QLowEnergyController__RemoteAddressType(C.QLowEnergyController_RemoteAddressType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyController) RemoteName() string {
	defer qt.Recovering("QLowEnergyController::remoteName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QLowEnergyController_RemoteName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLowEnergyController) RequestConnectionUpdate(parameters QLowEnergyConnectionParameters_ITF) {
	defer qt.Recovering("QLowEnergyController::requestConnectionUpdate")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_RequestConnectionUpdate(ptr.Pointer(), PointerFromQLowEnergyConnectionParameters(parameters))
	}
}

func (ptr *QLowEnergyController) Role() QLowEnergyController__Role {
	defer qt.Recovering("QLowEnergyController::role")

	if ptr.Pointer() != nil {
		return QLowEnergyController__Role(C.QLowEnergyController_Role(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyController) SetRemoteAddressType(ty QLowEnergyController__RemoteAddressType) {
	defer qt.Recovering("QLowEnergyController::setRemoteAddressType")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_SetRemoteAddressType(ptr.Pointer(), C.int(ty))
	}
}

func (ptr *QLowEnergyController) StartAdvertising(parameters QLowEnergyAdvertisingParameters_ITF, advertisingData QLowEnergyAdvertisingData_ITF, scanResponseData QLowEnergyAdvertisingData_ITF) {
	defer qt.Recovering("QLowEnergyController::startAdvertising")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_StartAdvertising(ptr.Pointer(), PointerFromQLowEnergyAdvertisingParameters(parameters), PointerFromQLowEnergyAdvertisingData(advertisingData), PointerFromQLowEnergyAdvertisingData(scanResponseData))
	}
}

func (ptr *QLowEnergyController) State() QLowEnergyController__ControllerState {
	defer qt.Recovering("QLowEnergyController::state")

	if ptr.Pointer() != nil {
		return QLowEnergyController__ControllerState(C.QLowEnergyController_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyController) StopAdvertising() {
	defer qt.Recovering("QLowEnergyController::stopAdvertising")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_StopAdvertising(ptr.Pointer())
	}
}

func (ptr *QLowEnergyController) DestroyQLowEnergyController() {
	defer qt.Recovering("QLowEnergyController::~QLowEnergyController")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QLowEnergyController_DestroyQLowEnergyController(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQLowEnergyController_TimerEvent
func callbackQLowEnergyController_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLowEnergyController::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQLowEnergyControllerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QLowEnergyController) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QLowEnergyController::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QLowEnergyController) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QLowEnergyController::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QLowEnergyController) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QLowEnergyController::timerEvent")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QLowEnergyController) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QLowEnergyController::timerEvent")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQLowEnergyController_ChildEvent
func callbackQLowEnergyController_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLowEnergyController::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQLowEnergyControllerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QLowEnergyController) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QLowEnergyController::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QLowEnergyController) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QLowEnergyController::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QLowEnergyController) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QLowEnergyController::childEvent")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QLowEnergyController) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QLowEnergyController::childEvent")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQLowEnergyController_ConnectNotify
func callbackQLowEnergyController_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QLowEnergyController::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQLowEnergyControllerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QLowEnergyController) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QLowEnergyController::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QLowEnergyController) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QLowEnergyController::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QLowEnergyController) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QLowEnergyController::connectNotify")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QLowEnergyController) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QLowEnergyController::connectNotify")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQLowEnergyController_CustomEvent
func callbackQLowEnergyController_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLowEnergyController::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQLowEnergyControllerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QLowEnergyController) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QLowEnergyController::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QLowEnergyController) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QLowEnergyController::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QLowEnergyController) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QLowEnergyController::customEvent")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QLowEnergyController) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QLowEnergyController::customEvent")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQLowEnergyController_DeleteLater
func callbackQLowEnergyController_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QLowEnergyController::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQLowEnergyControllerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QLowEnergyController) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QLowEnergyController::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QLowEnergyController) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QLowEnergyController::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QLowEnergyController) DeleteLater() {
	defer qt.Recovering("QLowEnergyController::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QLowEnergyController_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QLowEnergyController) DeleteLaterDefault() {
	defer qt.Recovering("QLowEnergyController::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QLowEnergyController_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQLowEnergyController_DisconnectNotify
func callbackQLowEnergyController_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QLowEnergyController::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQLowEnergyControllerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QLowEnergyController) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QLowEnergyController::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QLowEnergyController) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QLowEnergyController::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QLowEnergyController) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QLowEnergyController::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QLowEnergyController) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QLowEnergyController::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QLowEnergyController_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQLowEnergyController_Event
func callbackQLowEnergyController_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QLowEnergyController::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQLowEnergyControllerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QLowEnergyController) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QLowEnergyController::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QLowEnergyController) DisconnectEvent() {
	defer qt.Recovering("disconnect QLowEnergyController::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QLowEnergyController) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QLowEnergyController::event")

	if ptr.Pointer() != nil {
		return C.QLowEnergyController_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QLowEnergyController) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QLowEnergyController::event")

	if ptr.Pointer() != nil {
		return C.QLowEnergyController_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQLowEnergyController_EventFilter
func callbackQLowEnergyController_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QLowEnergyController::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQLowEnergyControllerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QLowEnergyController) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QLowEnergyController::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QLowEnergyController) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QLowEnergyController::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QLowEnergyController) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QLowEnergyController::eventFilter")

	if ptr.Pointer() != nil {
		return C.QLowEnergyController_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QLowEnergyController) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QLowEnergyController::eventFilter")

	if ptr.Pointer() != nil {
		return C.QLowEnergyController_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQLowEnergyController_MetaObject
func callbackQLowEnergyController_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QLowEnergyController::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQLowEnergyControllerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QLowEnergyController) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QLowEnergyController::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QLowEnergyController) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QLowEnergyController::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QLowEnergyController) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QLowEnergyController::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QLowEnergyController_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLowEnergyController) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QLowEnergyController::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QLowEnergyController_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QLowEnergyDescriptor struct {
	ptr unsafe.Pointer
}

type QLowEnergyDescriptor_ITF interface {
	QLowEnergyDescriptor_PTR() *QLowEnergyDescriptor
}

func (p *QLowEnergyDescriptor) QLowEnergyDescriptor_PTR() *QLowEnergyDescriptor {
	return p
}

func (p *QLowEnergyDescriptor) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QLowEnergyDescriptor) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQLowEnergyDescriptor(ptr QLowEnergyDescriptor_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLowEnergyDescriptor_PTR().Pointer()
	}
	return nil
}

func NewQLowEnergyDescriptorFromPointer(ptr unsafe.Pointer) *QLowEnergyDescriptor {
	var n = new(QLowEnergyDescriptor)
	n.SetPointer(ptr)
	return n
}

func newQLowEnergyDescriptorFromPointer(ptr unsafe.Pointer) *QLowEnergyDescriptor {
	var n = NewQLowEnergyDescriptorFromPointer(ptr)
	return n
}

func NewQLowEnergyDescriptor() *QLowEnergyDescriptor {
	defer qt.Recovering("QLowEnergyDescriptor::QLowEnergyDescriptor")

	return newQLowEnergyDescriptorFromPointer(C.QLowEnergyDescriptor_NewQLowEnergyDescriptor())
}

func NewQLowEnergyDescriptor2(other QLowEnergyDescriptor_ITF) *QLowEnergyDescriptor {
	defer qt.Recovering("QLowEnergyDescriptor::QLowEnergyDescriptor")

	return newQLowEnergyDescriptorFromPointer(C.QLowEnergyDescriptor_NewQLowEnergyDescriptor2(PointerFromQLowEnergyDescriptor(other)))
}

func (ptr *QLowEnergyDescriptor) IsValid() bool {
	defer qt.Recovering("QLowEnergyDescriptor::isValid")

	if ptr.Pointer() != nil {
		return C.QLowEnergyDescriptor_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLowEnergyDescriptor) Name() string {
	defer qt.Recovering("QLowEnergyDescriptor::name")

	if ptr.Pointer() != nil {
		return C.GoString(C.QLowEnergyDescriptor_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLowEnergyDescriptor) Type() QBluetoothUuid__DescriptorType {
	defer qt.Recovering("QLowEnergyDescriptor::type")

	if ptr.Pointer() != nil {
		return QBluetoothUuid__DescriptorType(C.QLowEnergyDescriptor_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyDescriptor) Uuid() *QBluetoothUuid {
	defer qt.Recovering("QLowEnergyDescriptor::uuid")

	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothUuidFromPointer(C.QLowEnergyDescriptor_Uuid(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyDescriptor) Value() string {
	defer qt.Recovering("QLowEnergyDescriptor::value")

	if ptr.Pointer() != nil {
		return qt.HexDecodeToString(C.GoString(C.QLowEnergyDescriptor_Value(ptr.Pointer())))
	}
	return ""
}

func (ptr *QLowEnergyDescriptor) DestroyQLowEnergyDescriptor() {
	defer qt.Recovering("QLowEnergyDescriptor::~QLowEnergyDescriptor")

	if ptr.Pointer() != nil {
		C.QLowEnergyDescriptor_DestroyQLowEnergyDescriptor(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

type QLowEnergyDescriptorData struct {
	ptr unsafe.Pointer
}

type QLowEnergyDescriptorData_ITF interface {
	QLowEnergyDescriptorData_PTR() *QLowEnergyDescriptorData
}

func (p *QLowEnergyDescriptorData) QLowEnergyDescriptorData_PTR() *QLowEnergyDescriptorData {
	return p
}

func (p *QLowEnergyDescriptorData) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QLowEnergyDescriptorData) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQLowEnergyDescriptorData(ptr QLowEnergyDescriptorData_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLowEnergyDescriptorData_PTR().Pointer()
	}
	return nil
}

func NewQLowEnergyDescriptorDataFromPointer(ptr unsafe.Pointer) *QLowEnergyDescriptorData {
	var n = new(QLowEnergyDescriptorData)
	n.SetPointer(ptr)
	return n
}

func newQLowEnergyDescriptorDataFromPointer(ptr unsafe.Pointer) *QLowEnergyDescriptorData {
	var n = NewQLowEnergyDescriptorDataFromPointer(ptr)
	return n
}

func NewQLowEnergyDescriptorData() *QLowEnergyDescriptorData {
	defer qt.Recovering("QLowEnergyDescriptorData::QLowEnergyDescriptorData")

	return newQLowEnergyDescriptorDataFromPointer(C.QLowEnergyDescriptorData_NewQLowEnergyDescriptorData())
}

func NewQLowEnergyDescriptorData2(uuid QBluetoothUuid_ITF, value string) *QLowEnergyDescriptorData {
	defer qt.Recovering("QLowEnergyDescriptorData::QLowEnergyDescriptorData")

	var valueC = C.CString(hex.EncodeToString([]byte(value)))
	defer C.free(unsafe.Pointer(valueC))
	return newQLowEnergyDescriptorDataFromPointer(C.QLowEnergyDescriptorData_NewQLowEnergyDescriptorData2(PointerFromQBluetoothUuid(uuid), valueC))
}

func NewQLowEnergyDescriptorData3(other QLowEnergyDescriptorData_ITF) *QLowEnergyDescriptorData {
	defer qt.Recovering("QLowEnergyDescriptorData::QLowEnergyDescriptorData")

	return newQLowEnergyDescriptorDataFromPointer(C.QLowEnergyDescriptorData_NewQLowEnergyDescriptorData3(PointerFromQLowEnergyDescriptorData(other)))
}

func (ptr *QLowEnergyDescriptorData) IsReadable() bool {
	defer qt.Recovering("QLowEnergyDescriptorData::isReadable")

	if ptr.Pointer() != nil {
		return C.QLowEnergyDescriptorData_IsReadable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLowEnergyDescriptorData) IsValid() bool {
	defer qt.Recovering("QLowEnergyDescriptorData::isValid")

	if ptr.Pointer() != nil {
		return C.QLowEnergyDescriptorData_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLowEnergyDescriptorData) IsWritable() bool {
	defer qt.Recovering("QLowEnergyDescriptorData::isWritable")

	if ptr.Pointer() != nil {
		return C.QLowEnergyDescriptorData_IsWritable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLowEnergyDescriptorData) SetUuid(uuid QBluetoothUuid_ITF) {
	defer qt.Recovering("QLowEnergyDescriptorData::setUuid")

	if ptr.Pointer() != nil {
		C.QLowEnergyDescriptorData_SetUuid(ptr.Pointer(), PointerFromQBluetoothUuid(uuid))
	}
}

func (ptr *QLowEnergyDescriptorData) SetValue(value string) {
	defer qt.Recovering("QLowEnergyDescriptorData::setValue")

	if ptr.Pointer() != nil {
		var valueC = C.CString(hex.EncodeToString([]byte(value)))
		defer C.free(unsafe.Pointer(valueC))
		C.QLowEnergyDescriptorData_SetValue(ptr.Pointer(), valueC)
	}
}

func (ptr *QLowEnergyDescriptorData) Swap(other QLowEnergyDescriptorData_ITF) {
	defer qt.Recovering("QLowEnergyDescriptorData::swap")

	if ptr.Pointer() != nil {
		C.QLowEnergyDescriptorData_Swap(ptr.Pointer(), PointerFromQLowEnergyDescriptorData(other))
	}
}

func (ptr *QLowEnergyDescriptorData) Uuid() *QBluetoothUuid {
	defer qt.Recovering("QLowEnergyDescriptorData::uuid")

	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothUuidFromPointer(C.QLowEnergyDescriptorData_Uuid(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyDescriptorData) Value() string {
	defer qt.Recovering("QLowEnergyDescriptorData::value")

	if ptr.Pointer() != nil {
		return qt.HexDecodeToString(C.GoString(C.QLowEnergyDescriptorData_Value(ptr.Pointer())))
	}
	return ""
}

func (ptr *QLowEnergyDescriptorData) DestroyQLowEnergyDescriptorData() {
	defer qt.Recovering("QLowEnergyDescriptorData::~QLowEnergyDescriptorData")

	if ptr.Pointer() != nil {
		C.QLowEnergyDescriptorData_DestroyQLowEnergyDescriptorData(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//QLowEnergyService::ServiceError
type QLowEnergyService__ServiceError int64

const (
	QLowEnergyService__NoError                  = QLowEnergyService__ServiceError(0)
	QLowEnergyService__OperationError           = QLowEnergyService__ServiceError(1)
	QLowEnergyService__CharacteristicWriteError = QLowEnergyService__ServiceError(2)
	QLowEnergyService__DescriptorWriteError     = QLowEnergyService__ServiceError(3)
	QLowEnergyService__UnknownError             = QLowEnergyService__ServiceError(4)
	QLowEnergyService__CharacteristicReadError  = QLowEnergyService__ServiceError(5)
	QLowEnergyService__DescriptorReadError      = QLowEnergyService__ServiceError(6)
)

//QLowEnergyService::ServiceState
type QLowEnergyService__ServiceState int64

const (
	QLowEnergyService__InvalidService      = QLowEnergyService__ServiceState(0)
	QLowEnergyService__DiscoveryRequired   = QLowEnergyService__ServiceState(1)
	QLowEnergyService__DiscoveringServices = QLowEnergyService__ServiceState(2)
	QLowEnergyService__ServiceDiscovered   = QLowEnergyService__ServiceState(3)
	QLowEnergyService__LocalService        = QLowEnergyService__ServiceState(4)
)

//QLowEnergyService::ServiceType
type QLowEnergyService__ServiceType int64

const (
	QLowEnergyService__PrimaryService  = QLowEnergyService__ServiceType(0x0001)
	QLowEnergyService__IncludedService = QLowEnergyService__ServiceType(0x0002)
)

//QLowEnergyService::WriteMode
type QLowEnergyService__WriteMode int64

const (
	QLowEnergyService__WriteWithResponse    = QLowEnergyService__WriteMode(0)
	QLowEnergyService__WriteWithoutResponse = QLowEnergyService__WriteMode(1)
	QLowEnergyService__WriteSigned          = QLowEnergyService__WriteMode(2)
)

type QLowEnergyService struct {
	core.QObject
}

type QLowEnergyService_ITF interface {
	core.QObject_ITF
	QLowEnergyService_PTR() *QLowEnergyService
}

func (p *QLowEnergyService) QLowEnergyService_PTR() *QLowEnergyService {
	return p
}

func (p *QLowEnergyService) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QLowEnergyService) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQLowEnergyService(ptr QLowEnergyService_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLowEnergyService_PTR().Pointer()
	}
	return nil
}

func NewQLowEnergyServiceFromPointer(ptr unsafe.Pointer) *QLowEnergyService {
	var n = new(QLowEnergyService)
	n.SetPointer(ptr)
	return n
}

func newQLowEnergyServiceFromPointer(ptr unsafe.Pointer) *QLowEnergyService {
	var n = NewQLowEnergyServiceFromPointer(ptr)
	for len(n.ObjectName()) < len("QLowEnergyService_") {
		n.SetObjectName("QLowEnergyService_" + qt.Identifier())
	}
	return n
}

//export callbackQLowEnergyService_CharacteristicChanged
func callbackQLowEnergyService_CharacteristicChanged(ptr unsafe.Pointer, ptrName *C.char, characteristic unsafe.Pointer, newValue *C.char) {
	defer qt.Recovering("callback QLowEnergyService::characteristicChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "characteristicChanged"); signal != nil {
		signal.(func(*QLowEnergyCharacteristic, string))(NewQLowEnergyCharacteristicFromPointer(characteristic), qt.HexDecodeToString(C.GoString(newValue)))
	}

}

func (ptr *QLowEnergyService) ConnectCharacteristicChanged(f func(characteristic *QLowEnergyCharacteristic, newValue string)) {
	defer qt.Recovering("connect QLowEnergyService::characteristicChanged")

	if ptr.Pointer() != nil {
		C.QLowEnergyService_ConnectCharacteristicChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "characteristicChanged", f)
	}
}

func (ptr *QLowEnergyService) DisconnectCharacteristicChanged() {
	defer qt.Recovering("disconnect QLowEnergyService::characteristicChanged")

	if ptr.Pointer() != nil {
		C.QLowEnergyService_DisconnectCharacteristicChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "characteristicChanged")
	}
}

func (ptr *QLowEnergyService) CharacteristicChanged(characteristic QLowEnergyCharacteristic_ITF, newValue string) {
	defer qt.Recovering("QLowEnergyService::characteristicChanged")

	if ptr.Pointer() != nil {
		var newValueC = C.CString(hex.EncodeToString([]byte(newValue)))
		defer C.free(unsafe.Pointer(newValueC))
		C.QLowEnergyService_CharacteristicChanged(ptr.Pointer(), PointerFromQLowEnergyCharacteristic(characteristic), newValueC)
	}
}

//export callbackQLowEnergyService_CharacteristicRead
func callbackQLowEnergyService_CharacteristicRead(ptr unsafe.Pointer, ptrName *C.char, characteristic unsafe.Pointer, value *C.char) {
	defer qt.Recovering("callback QLowEnergyService::characteristicRead")

	if signal := qt.GetSignal(C.GoString(ptrName), "characteristicRead"); signal != nil {
		signal.(func(*QLowEnergyCharacteristic, string))(NewQLowEnergyCharacteristicFromPointer(characteristic), qt.HexDecodeToString(C.GoString(value)))
	}

}

func (ptr *QLowEnergyService) ConnectCharacteristicRead(f func(characteristic *QLowEnergyCharacteristic, value string)) {
	defer qt.Recovering("connect QLowEnergyService::characteristicRead")

	if ptr.Pointer() != nil {
		C.QLowEnergyService_ConnectCharacteristicRead(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "characteristicRead", f)
	}
}

func (ptr *QLowEnergyService) DisconnectCharacteristicRead() {
	defer qt.Recovering("disconnect QLowEnergyService::characteristicRead")

	if ptr.Pointer() != nil {
		C.QLowEnergyService_DisconnectCharacteristicRead(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "characteristicRead")
	}
}

func (ptr *QLowEnergyService) CharacteristicRead(characteristic QLowEnergyCharacteristic_ITF, value string) {
	defer qt.Recovering("QLowEnergyService::characteristicRead")

	if ptr.Pointer() != nil {
		var valueC = C.CString(hex.EncodeToString([]byte(value)))
		defer C.free(unsafe.Pointer(valueC))
		C.QLowEnergyService_CharacteristicRead(ptr.Pointer(), PointerFromQLowEnergyCharacteristic(characteristic), valueC)
	}
}

//export callbackQLowEnergyService_CharacteristicWritten
func callbackQLowEnergyService_CharacteristicWritten(ptr unsafe.Pointer, ptrName *C.char, characteristic unsafe.Pointer, newValue *C.char) {
	defer qt.Recovering("callback QLowEnergyService::characteristicWritten")

	if signal := qt.GetSignal(C.GoString(ptrName), "characteristicWritten"); signal != nil {
		signal.(func(*QLowEnergyCharacteristic, string))(NewQLowEnergyCharacteristicFromPointer(characteristic), qt.HexDecodeToString(C.GoString(newValue)))
	}

}

func (ptr *QLowEnergyService) ConnectCharacteristicWritten(f func(characteristic *QLowEnergyCharacteristic, newValue string)) {
	defer qt.Recovering("connect QLowEnergyService::characteristicWritten")

	if ptr.Pointer() != nil {
		C.QLowEnergyService_ConnectCharacteristicWritten(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "characteristicWritten", f)
	}
}

func (ptr *QLowEnergyService) DisconnectCharacteristicWritten() {
	defer qt.Recovering("disconnect QLowEnergyService::characteristicWritten")

	if ptr.Pointer() != nil {
		C.QLowEnergyService_DisconnectCharacteristicWritten(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "characteristicWritten")
	}
}

func (ptr *QLowEnergyService) CharacteristicWritten(characteristic QLowEnergyCharacteristic_ITF, newValue string) {
	defer qt.Recovering("QLowEnergyService::characteristicWritten")

	if ptr.Pointer() != nil {
		var newValueC = C.CString(hex.EncodeToString([]byte(newValue)))
		defer C.free(unsafe.Pointer(newValueC))
		C.QLowEnergyService_CharacteristicWritten(ptr.Pointer(), PointerFromQLowEnergyCharacteristic(characteristic), newValueC)
	}
}

//export callbackQLowEnergyService_DescriptorRead
func callbackQLowEnergyService_DescriptorRead(ptr unsafe.Pointer, ptrName *C.char, descriptor unsafe.Pointer, value *C.char) {
	defer qt.Recovering("callback QLowEnergyService::descriptorRead")

	if signal := qt.GetSignal(C.GoString(ptrName), "descriptorRead"); signal != nil {
		signal.(func(*QLowEnergyDescriptor, string))(NewQLowEnergyDescriptorFromPointer(descriptor), qt.HexDecodeToString(C.GoString(value)))
	}

}

func (ptr *QLowEnergyService) ConnectDescriptorRead(f func(descriptor *QLowEnergyDescriptor, value string)) {
	defer qt.Recovering("connect QLowEnergyService::descriptorRead")

	if ptr.Pointer() != nil {
		C.QLowEnergyService_ConnectDescriptorRead(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "descriptorRead", f)
	}
}

func (ptr *QLowEnergyService) DisconnectDescriptorRead() {
	defer qt.Recovering("disconnect QLowEnergyService::descriptorRead")

	if ptr.Pointer() != nil {
		C.QLowEnergyService_DisconnectDescriptorRead(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "descriptorRead")
	}
}

func (ptr *QLowEnergyService) DescriptorRead(descriptor QLowEnergyDescriptor_ITF, value string) {
	defer qt.Recovering("QLowEnergyService::descriptorRead")

	if ptr.Pointer() != nil {
		var valueC = C.CString(hex.EncodeToString([]byte(value)))
		defer C.free(unsafe.Pointer(valueC))
		C.QLowEnergyService_DescriptorRead(ptr.Pointer(), PointerFromQLowEnergyDescriptor(descriptor), valueC)
	}
}

//export callbackQLowEnergyService_DescriptorWritten
func callbackQLowEnergyService_DescriptorWritten(ptr unsafe.Pointer, ptrName *C.char, descriptor unsafe.Pointer, newValue *C.char) {
	defer qt.Recovering("callback QLowEnergyService::descriptorWritten")

	if signal := qt.GetSignal(C.GoString(ptrName), "descriptorWritten"); signal != nil {
		signal.(func(*QLowEnergyDescriptor, string))(NewQLowEnergyDescriptorFromPointer(descriptor), qt.HexDecodeToString(C.GoString(newValue)))
	}

}

func (ptr *QLowEnergyService) ConnectDescriptorWritten(f func(descriptor *QLowEnergyDescriptor, newValue string)) {
	defer qt.Recovering("connect QLowEnergyService::descriptorWritten")

	if ptr.Pointer() != nil {
		C.QLowEnergyService_ConnectDescriptorWritten(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "descriptorWritten", f)
	}
}

func (ptr *QLowEnergyService) DisconnectDescriptorWritten() {
	defer qt.Recovering("disconnect QLowEnergyService::descriptorWritten")

	if ptr.Pointer() != nil {
		C.QLowEnergyService_DisconnectDescriptorWritten(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "descriptorWritten")
	}
}

func (ptr *QLowEnergyService) DescriptorWritten(descriptor QLowEnergyDescriptor_ITF, newValue string) {
	defer qt.Recovering("QLowEnergyService::descriptorWritten")

	if ptr.Pointer() != nil {
		var newValueC = C.CString(hex.EncodeToString([]byte(newValue)))
		defer C.free(unsafe.Pointer(newValueC))
		C.QLowEnergyService_DescriptorWritten(ptr.Pointer(), PointerFromQLowEnergyDescriptor(descriptor), newValueC)
	}
}

//export callbackQLowEnergyService_Error2
func callbackQLowEnergyService_Error2(ptr unsafe.Pointer, ptrName *C.char, newError C.int) {
	defer qt.Recovering("callback QLowEnergyService::error")

	if signal := qt.GetSignal(C.GoString(ptrName), "error2"); signal != nil {
		signal.(func(QLowEnergyService__ServiceError))(QLowEnergyService__ServiceError(newError))
	}

}

func (ptr *QLowEnergyService) ConnectError2(f func(newError QLowEnergyService__ServiceError)) {
	defer qt.Recovering("connect QLowEnergyService::error")

	if ptr.Pointer() != nil {
		C.QLowEnergyService_ConnectError2(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "error2", f)
	}
}

func (ptr *QLowEnergyService) DisconnectError2() {
	defer qt.Recovering("disconnect QLowEnergyService::error")

	if ptr.Pointer() != nil {
		C.QLowEnergyService_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "error2")
	}
}

func (ptr *QLowEnergyService) Error2(newError QLowEnergyService__ServiceError) {
	defer qt.Recovering("QLowEnergyService::error")

	if ptr.Pointer() != nil {
		C.QLowEnergyService_Error2(ptr.Pointer(), C.int(newError))
	}
}

//export callbackQLowEnergyService_StateChanged
func callbackQLowEnergyService_StateChanged(ptr unsafe.Pointer, ptrName *C.char, newState C.int) {
	defer qt.Recovering("callback QLowEnergyService::stateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stateChanged"); signal != nil {
		signal.(func(QLowEnergyService__ServiceState))(QLowEnergyService__ServiceState(newState))
	}

}

func (ptr *QLowEnergyService) ConnectStateChanged(f func(newState QLowEnergyService__ServiceState)) {
	defer qt.Recovering("connect QLowEnergyService::stateChanged")

	if ptr.Pointer() != nil {
		C.QLowEnergyService_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QLowEnergyService) DisconnectStateChanged() {
	defer qt.Recovering("disconnect QLowEnergyService::stateChanged")

	if ptr.Pointer() != nil {
		C.QLowEnergyService_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

func (ptr *QLowEnergyService) StateChanged(newState QLowEnergyService__ServiceState) {
	defer qt.Recovering("QLowEnergyService::stateChanged")

	if ptr.Pointer() != nil {
		C.QLowEnergyService_StateChanged(ptr.Pointer(), C.int(newState))
	}
}

func (ptr *QLowEnergyService) Characteristic(uuid QBluetoothUuid_ITF) *QLowEnergyCharacteristic {
	defer qt.Recovering("QLowEnergyService::characteristic")

	if ptr.Pointer() != nil {
		var tmpValue = NewQLowEnergyCharacteristicFromPointer(C.QLowEnergyService_Characteristic(ptr.Pointer(), PointerFromQBluetoothUuid(uuid)))
		runtime.SetFinalizer(tmpValue, (*QLowEnergyCharacteristic).DestroyQLowEnergyCharacteristic)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyService) Contains(characteristic QLowEnergyCharacteristic_ITF) bool {
	defer qt.Recovering("QLowEnergyService::contains")

	if ptr.Pointer() != nil {
		return C.QLowEnergyService_Contains(ptr.Pointer(), PointerFromQLowEnergyCharacteristic(characteristic)) != 0
	}
	return false
}

func (ptr *QLowEnergyService) Contains2(descriptor QLowEnergyDescriptor_ITF) bool {
	defer qt.Recovering("QLowEnergyService::contains")

	if ptr.Pointer() != nil {
		return C.QLowEnergyService_Contains2(ptr.Pointer(), PointerFromQLowEnergyDescriptor(descriptor)) != 0
	}
	return false
}

func (ptr *QLowEnergyService) DiscoverDetails() {
	defer qt.Recovering("QLowEnergyService::discoverDetails")

	if ptr.Pointer() != nil {
		C.QLowEnergyService_DiscoverDetails(ptr.Pointer())
	}
}

func (ptr *QLowEnergyService) Error() QLowEnergyService__ServiceError {
	defer qt.Recovering("QLowEnergyService::error")

	if ptr.Pointer() != nil {
		return QLowEnergyService__ServiceError(C.QLowEnergyService_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyService) ReadCharacteristic(characteristic QLowEnergyCharacteristic_ITF) {
	defer qt.Recovering("QLowEnergyService::readCharacteristic")

	if ptr.Pointer() != nil {
		C.QLowEnergyService_ReadCharacteristic(ptr.Pointer(), PointerFromQLowEnergyCharacteristic(characteristic))
	}
}

func (ptr *QLowEnergyService) ReadDescriptor(descriptor QLowEnergyDescriptor_ITF) {
	defer qt.Recovering("QLowEnergyService::readDescriptor")

	if ptr.Pointer() != nil {
		C.QLowEnergyService_ReadDescriptor(ptr.Pointer(), PointerFromQLowEnergyDescriptor(descriptor))
	}
}

func (ptr *QLowEnergyService) ServiceName() string {
	defer qt.Recovering("QLowEnergyService::serviceName")

	if ptr.Pointer() != nil {
		return C.GoString(C.QLowEnergyService_ServiceName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLowEnergyService) ServiceUuid() *QBluetoothUuid {
	defer qt.Recovering("QLowEnergyService::serviceUuid")

	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothUuidFromPointer(C.QLowEnergyService_ServiceUuid(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyService) State() QLowEnergyService__ServiceState {
	defer qt.Recovering("QLowEnergyService::state")

	if ptr.Pointer() != nil {
		return QLowEnergyService__ServiceState(C.QLowEnergyService_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyService) Type() QLowEnergyService__ServiceType {
	defer qt.Recovering("QLowEnergyService::type")

	if ptr.Pointer() != nil {
		return QLowEnergyService__ServiceType(C.QLowEnergyService_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyService) WriteCharacteristic(characteristic QLowEnergyCharacteristic_ITF, newValue string, mode QLowEnergyService__WriteMode) {
	defer qt.Recovering("QLowEnergyService::writeCharacteristic")

	if ptr.Pointer() != nil {
		var newValueC = C.CString(hex.EncodeToString([]byte(newValue)))
		defer C.free(unsafe.Pointer(newValueC))
		C.QLowEnergyService_WriteCharacteristic(ptr.Pointer(), PointerFromQLowEnergyCharacteristic(characteristic), newValueC, C.int(mode))
	}
}

func (ptr *QLowEnergyService) WriteDescriptor(descriptor QLowEnergyDescriptor_ITF, newValue string) {
	defer qt.Recovering("QLowEnergyService::writeDescriptor")

	if ptr.Pointer() != nil {
		var newValueC = C.CString(hex.EncodeToString([]byte(newValue)))
		defer C.free(unsafe.Pointer(newValueC))
		C.QLowEnergyService_WriteDescriptor(ptr.Pointer(), PointerFromQLowEnergyDescriptor(descriptor), newValueC)
	}
}

func (ptr *QLowEnergyService) DestroyQLowEnergyService() {
	defer qt.Recovering("QLowEnergyService::~QLowEnergyService")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QLowEnergyService_DestroyQLowEnergyService(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQLowEnergyService_TimerEvent
func callbackQLowEnergyService_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLowEnergyService::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQLowEnergyServiceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QLowEnergyService) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QLowEnergyService::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QLowEnergyService) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QLowEnergyService::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QLowEnergyService) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QLowEnergyService::timerEvent")

	if ptr.Pointer() != nil {
		C.QLowEnergyService_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QLowEnergyService) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QLowEnergyService::timerEvent")

	if ptr.Pointer() != nil {
		C.QLowEnergyService_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQLowEnergyService_ChildEvent
func callbackQLowEnergyService_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLowEnergyService::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQLowEnergyServiceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QLowEnergyService) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QLowEnergyService::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QLowEnergyService) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QLowEnergyService::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QLowEnergyService) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QLowEnergyService::childEvent")

	if ptr.Pointer() != nil {
		C.QLowEnergyService_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QLowEnergyService) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QLowEnergyService::childEvent")

	if ptr.Pointer() != nil {
		C.QLowEnergyService_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQLowEnergyService_ConnectNotify
func callbackQLowEnergyService_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QLowEnergyService::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQLowEnergyServiceFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QLowEnergyService) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QLowEnergyService::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QLowEnergyService) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QLowEnergyService::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QLowEnergyService) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QLowEnergyService::connectNotify")

	if ptr.Pointer() != nil {
		C.QLowEnergyService_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QLowEnergyService) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QLowEnergyService::connectNotify")

	if ptr.Pointer() != nil {
		C.QLowEnergyService_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQLowEnergyService_CustomEvent
func callbackQLowEnergyService_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QLowEnergyService::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQLowEnergyServiceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QLowEnergyService) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QLowEnergyService::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QLowEnergyService) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QLowEnergyService::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QLowEnergyService) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QLowEnergyService::customEvent")

	if ptr.Pointer() != nil {
		C.QLowEnergyService_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QLowEnergyService) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QLowEnergyService::customEvent")

	if ptr.Pointer() != nil {
		C.QLowEnergyService_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQLowEnergyService_DeleteLater
func callbackQLowEnergyService_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QLowEnergyService::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQLowEnergyServiceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QLowEnergyService) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QLowEnergyService::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QLowEnergyService) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QLowEnergyService::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QLowEnergyService) DeleteLater() {
	defer qt.Recovering("QLowEnergyService::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QLowEnergyService_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QLowEnergyService) DeleteLaterDefault() {
	defer qt.Recovering("QLowEnergyService::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QLowEnergyService_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQLowEnergyService_DisconnectNotify
func callbackQLowEnergyService_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QLowEnergyService::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQLowEnergyServiceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QLowEnergyService) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QLowEnergyService::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QLowEnergyService) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QLowEnergyService::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QLowEnergyService) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QLowEnergyService::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QLowEnergyService_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QLowEnergyService) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QLowEnergyService::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QLowEnergyService_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQLowEnergyService_Event
func callbackQLowEnergyService_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QLowEnergyService::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQLowEnergyServiceFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QLowEnergyService) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QLowEnergyService::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QLowEnergyService) DisconnectEvent() {
	defer qt.Recovering("disconnect QLowEnergyService::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QLowEnergyService) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QLowEnergyService::event")

	if ptr.Pointer() != nil {
		return C.QLowEnergyService_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QLowEnergyService) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QLowEnergyService::event")

	if ptr.Pointer() != nil {
		return C.QLowEnergyService_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQLowEnergyService_EventFilter
func callbackQLowEnergyService_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QLowEnergyService::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQLowEnergyServiceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QLowEnergyService) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QLowEnergyService::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QLowEnergyService) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QLowEnergyService::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QLowEnergyService) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QLowEnergyService::eventFilter")

	if ptr.Pointer() != nil {
		return C.QLowEnergyService_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QLowEnergyService) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QLowEnergyService::eventFilter")

	if ptr.Pointer() != nil {
		return C.QLowEnergyService_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQLowEnergyService_MetaObject
func callbackQLowEnergyService_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QLowEnergyService::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQLowEnergyServiceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QLowEnergyService) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QLowEnergyService::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QLowEnergyService) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QLowEnergyService::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QLowEnergyService) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QLowEnergyService::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QLowEnergyService_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QLowEnergyService) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QLowEnergyService::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QLowEnergyService_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QLowEnergyServiceData::ServiceType
type QLowEnergyServiceData__ServiceType int64

const (
	QLowEnergyServiceData__ServiceTypePrimary   = QLowEnergyServiceData__ServiceType(0x2800)
	QLowEnergyServiceData__ServiceTypeSecondary = QLowEnergyServiceData__ServiceType(0x2801)
)

type QLowEnergyServiceData struct {
	ptr unsafe.Pointer
}

type QLowEnergyServiceData_ITF interface {
	QLowEnergyServiceData_PTR() *QLowEnergyServiceData
}

func (p *QLowEnergyServiceData) QLowEnergyServiceData_PTR() *QLowEnergyServiceData {
	return p
}

func (p *QLowEnergyServiceData) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QLowEnergyServiceData) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQLowEnergyServiceData(ptr QLowEnergyServiceData_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QLowEnergyServiceData_PTR().Pointer()
	}
	return nil
}

func NewQLowEnergyServiceDataFromPointer(ptr unsafe.Pointer) *QLowEnergyServiceData {
	var n = new(QLowEnergyServiceData)
	n.SetPointer(ptr)
	return n
}

func newQLowEnergyServiceDataFromPointer(ptr unsafe.Pointer) *QLowEnergyServiceData {
	var n = NewQLowEnergyServiceDataFromPointer(ptr)
	return n
}

func NewQLowEnergyServiceData() *QLowEnergyServiceData {
	defer qt.Recovering("QLowEnergyServiceData::QLowEnergyServiceData")

	return newQLowEnergyServiceDataFromPointer(C.QLowEnergyServiceData_NewQLowEnergyServiceData())
}

func NewQLowEnergyServiceData2(other QLowEnergyServiceData_ITF) *QLowEnergyServiceData {
	defer qt.Recovering("QLowEnergyServiceData::QLowEnergyServiceData")

	return newQLowEnergyServiceDataFromPointer(C.QLowEnergyServiceData_NewQLowEnergyServiceData2(PointerFromQLowEnergyServiceData(other)))
}

func (ptr *QLowEnergyServiceData) AddCharacteristic(characteristic QLowEnergyCharacteristicData_ITF) {
	defer qt.Recovering("QLowEnergyServiceData::addCharacteristic")

	if ptr.Pointer() != nil {
		C.QLowEnergyServiceData_AddCharacteristic(ptr.Pointer(), PointerFromQLowEnergyCharacteristicData(characteristic))
	}
}

func (ptr *QLowEnergyServiceData) AddIncludedService(service QLowEnergyService_ITF) {
	defer qt.Recovering("QLowEnergyServiceData::addIncludedService")

	if ptr.Pointer() != nil {
		C.QLowEnergyServiceData_AddIncludedService(ptr.Pointer(), PointerFromQLowEnergyService(service))
	}
}

func (ptr *QLowEnergyServiceData) IsValid() bool {
	defer qt.Recovering("QLowEnergyServiceData::isValid")

	if ptr.Pointer() != nil {
		return C.QLowEnergyServiceData_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLowEnergyServiceData) SetType(ty QLowEnergyServiceData__ServiceType) {
	defer qt.Recovering("QLowEnergyServiceData::setType")

	if ptr.Pointer() != nil {
		C.QLowEnergyServiceData_SetType(ptr.Pointer(), C.int(ty))
	}
}

func (ptr *QLowEnergyServiceData) SetUuid(uuid QBluetoothUuid_ITF) {
	defer qt.Recovering("QLowEnergyServiceData::setUuid")

	if ptr.Pointer() != nil {
		C.QLowEnergyServiceData_SetUuid(ptr.Pointer(), PointerFromQBluetoothUuid(uuid))
	}
}

func (ptr *QLowEnergyServiceData) Swap(other QLowEnergyServiceData_ITF) {
	defer qt.Recovering("QLowEnergyServiceData::swap")

	if ptr.Pointer() != nil {
		C.QLowEnergyServiceData_Swap(ptr.Pointer(), PointerFromQLowEnergyServiceData(other))
	}
}

func (ptr *QLowEnergyServiceData) Type() QLowEnergyServiceData__ServiceType {
	defer qt.Recovering("QLowEnergyServiceData::type")

	if ptr.Pointer() != nil {
		return QLowEnergyServiceData__ServiceType(C.QLowEnergyServiceData_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyServiceData) Uuid() *QBluetoothUuid {
	defer qt.Recovering("QLowEnergyServiceData::uuid")

	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothUuidFromPointer(C.QLowEnergyServiceData_Uuid(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyServiceData) DestroyQLowEnergyServiceData() {
	defer qt.Recovering("QLowEnergyServiceData::~QLowEnergyServiceData")

	if ptr.Pointer() != nil {
		C.QLowEnergyServiceData_DestroyQLowEnergyServiceData(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
