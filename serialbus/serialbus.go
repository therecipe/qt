// +build !minimal

package serialbus

//#include <stdlib.h>
//#include "serialbus.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"runtime"
	"unsafe"
)

type QCanBus struct {
	core.QObject
}

type QCanBus_ITF interface {
	core.QObject_ITF
	QCanBus_PTR() *QCanBus
}

func (p *QCanBus) QCanBus_PTR() *QCanBus {
	return p
}

func (p *QCanBus) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QCanBus) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQCanBus(ptr QCanBus_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCanBus_PTR().Pointer()
	}
	return nil
}

func NewQCanBusFromPointer(ptr unsafe.Pointer) *QCanBus {
	var n = new(QCanBus)
	n.SetPointer(ptr)
	return n
}

func newQCanBusFromPointer(ptr unsafe.Pointer) *QCanBus {
	var n = NewQCanBusFromPointer(ptr)
	for len(n.ObjectName()) < len("QCanBus_") {
		n.SetObjectName("QCanBus_" + qt.Identifier())
	}
	return n
}

func (ptr *QCanBus) DestroyQCanBus() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

func (ptr *QCanBus) CreateDevice(plugin string, interfaceName string) *QCanBusDevice {
	defer qt.Recovering("QCanBus::createDevice")

	if ptr.Pointer() != nil {
		var pluginC = C.CString(plugin)
		defer C.free(unsafe.Pointer(pluginC))
		var interfaceNameC = C.CString(interfaceName)
		defer C.free(unsafe.Pointer(interfaceNameC))
		return NewQCanBusDeviceFromPointer(C.QCanBus_CreateDevice(ptr.Pointer(), pluginC, interfaceNameC))
	}
	return nil
}

func QCanBus_Instance() *QCanBus {
	defer qt.Recovering("QCanBus::instance")

	return NewQCanBusFromPointer(C.QCanBus_QCanBus_Instance())
}

func (ptr *QCanBus) Instance() *QCanBus {
	defer qt.Recovering("QCanBus::instance")

	return NewQCanBusFromPointer(C.QCanBus_QCanBus_Instance())
}

//export callbackQCanBus_TimerEvent
func callbackQCanBus_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCanBus::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQCanBusFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QCanBus) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QCanBus::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCanBus) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCanBus::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QCanBus) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCanBus::timerEvent")

	if ptr.Pointer() != nil {
		C.QCanBus_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCanBus) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCanBus::timerEvent")

	if ptr.Pointer() != nil {
		C.QCanBus_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQCanBus_ChildEvent
func callbackQCanBus_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCanBus::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCanBusFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCanBus) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCanBus::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCanBus) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCanBus::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QCanBus) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCanBus::childEvent")

	if ptr.Pointer() != nil {
		C.QCanBus_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCanBus) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCanBus::childEvent")

	if ptr.Pointer() != nil {
		C.QCanBus_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQCanBus_ConnectNotify
func callbackQCanBus_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QCanBus::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQCanBusFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QCanBus) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QCanBus::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QCanBus) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QCanBus::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QCanBus) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QCanBus::connectNotify")

	if ptr.Pointer() != nil {
		C.QCanBus_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QCanBus) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QCanBus::connectNotify")

	if ptr.Pointer() != nil {
		C.QCanBus_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQCanBus_CustomEvent
func callbackQCanBus_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCanBus::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCanBusFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCanBus) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCanBus::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCanBus) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCanBus::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QCanBus) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCanBus::customEvent")

	if ptr.Pointer() != nil {
		C.QCanBus_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCanBus) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCanBus::customEvent")

	if ptr.Pointer() != nil {
		C.QCanBus_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQCanBus_DeleteLater
func callbackQCanBus_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QCanBus::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQCanBusFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QCanBus) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QCanBus::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QCanBus) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QCanBus::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QCanBus) DeleteLater() {
	defer qt.Recovering("QCanBus::deleteLater")

	if ptr.Pointer() != nil {
		C.QCanBus_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCanBus) DeleteLaterDefault() {
	defer qt.Recovering("QCanBus::deleteLater")

	if ptr.Pointer() != nil {
		C.QCanBus_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQCanBus_DisconnectNotify
func callbackQCanBus_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QCanBus::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQCanBusFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QCanBus) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QCanBus::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QCanBus) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QCanBus::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QCanBus) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QCanBus::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QCanBus_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QCanBus) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QCanBus::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QCanBus_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQCanBus_Event
func callbackQCanBus_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QCanBus::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQCanBusFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QCanBus) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QCanBus::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QCanBus) DisconnectEvent() {
	defer qt.Recovering("disconnect QCanBus::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QCanBus) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QCanBus::event")

	if ptr.Pointer() != nil {
		return C.QCanBus_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QCanBus) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QCanBus::event")

	if ptr.Pointer() != nil {
		return C.QCanBus_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQCanBus_EventFilter
func callbackQCanBus_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QCanBus::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQCanBusFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QCanBus) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QCanBus::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QCanBus) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QCanBus::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QCanBus) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QCanBus::eventFilter")

	if ptr.Pointer() != nil {
		return C.QCanBus_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QCanBus) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QCanBus::eventFilter")

	if ptr.Pointer() != nil {
		return C.QCanBus_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQCanBus_MetaObject
func callbackQCanBus_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QCanBus::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQCanBusFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QCanBus) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QCanBus::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QCanBus) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QCanBus::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QCanBus) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QCanBus::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QCanBus_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCanBus) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QCanBus::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QCanBus_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QCanBusDevice::CanBusDeviceState
type QCanBusDevice__CanBusDeviceState int64

const (
	QCanBusDevice__UnconnectedState = QCanBusDevice__CanBusDeviceState(0)
	QCanBusDevice__ConnectingState  = QCanBusDevice__CanBusDeviceState(1)
	QCanBusDevice__ConnectedState   = QCanBusDevice__CanBusDeviceState(2)
	QCanBusDevice__ClosingState     = QCanBusDevice__CanBusDeviceState(3)
)

//QCanBusDevice::CanBusError
type QCanBusDevice__CanBusError int64

const (
	QCanBusDevice__NoError            = QCanBusDevice__CanBusError(0)
	QCanBusDevice__ReadError          = QCanBusDevice__CanBusError(1)
	QCanBusDevice__WriteError         = QCanBusDevice__CanBusError(2)
	QCanBusDevice__ConnectionError    = QCanBusDevice__CanBusError(3)
	QCanBusDevice__ConfigurationError = QCanBusDevice__CanBusError(4)
	QCanBusDevice__UnknownError       = QCanBusDevice__CanBusError(5)
)

//QCanBusDevice::ConfigurationKey
type QCanBusDevice__ConfigurationKey int64

const (
	QCanBusDevice__RawFilterKey   = QCanBusDevice__ConfigurationKey(0)
	QCanBusDevice__ErrorFilterKey = QCanBusDevice__ConfigurationKey(1)
	QCanBusDevice__LoopbackKey    = QCanBusDevice__ConfigurationKey(2)
	QCanBusDevice__ReceiveOwnKey  = QCanBusDevice__ConfigurationKey(3)
	QCanBusDevice__BitRateKey     = QCanBusDevice__ConfigurationKey(4)
	QCanBusDevice__CanFdKey       = QCanBusDevice__ConfigurationKey(5)
	QCanBusDevice__UserKey        = QCanBusDevice__ConfigurationKey(30)
)

type QCanBusDevice struct {
	core.QObject
}

type QCanBusDevice_ITF interface {
	core.QObject_ITF
	QCanBusDevice_PTR() *QCanBusDevice
}

func (p *QCanBusDevice) QCanBusDevice_PTR() *QCanBusDevice {
	return p
}

func (p *QCanBusDevice) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QCanBusDevice) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQCanBusDevice(ptr QCanBusDevice_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCanBusDevice_PTR().Pointer()
	}
	return nil
}

func NewQCanBusDeviceFromPointer(ptr unsafe.Pointer) *QCanBusDevice {
	var n = new(QCanBusDevice)
	n.SetPointer(ptr)
	return n
}

func newQCanBusDeviceFromPointer(ptr unsafe.Pointer) *QCanBusDevice {
	var n = NewQCanBusDeviceFromPointer(ptr)
	for len(n.ObjectName()) < len("QCanBusDevice_") {
		n.SetObjectName("QCanBusDevice_" + qt.Identifier())
	}
	return n
}

func (ptr *QCanBusDevice) DestroyQCanBusDevice() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

func (ptr *QCanBusDevice) FramesAvailable() int64 {
	defer qt.Recovering("QCanBusDevice::framesAvailable")

	if ptr.Pointer() != nil {
		return int64(C.QCanBusDevice_FramesAvailable(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCanBusDevice) FramesToWrite() int64 {
	defer qt.Recovering("QCanBusDevice::framesToWrite")

	if ptr.Pointer() != nil {
		return int64(C.QCanBusDevice_FramesToWrite(ptr.Pointer()))
	}
	return 0
}

func NewQCanBusDevice(parent core.QObject_ITF) *QCanBusDevice {
	defer qt.Recovering("QCanBusDevice::QCanBusDevice")

	return newQCanBusDeviceFromPointer(C.QCanBusDevice_NewQCanBusDevice(core.PointerFromQObject(parent)))
}

//export callbackQCanBusDevice_Close
func callbackQCanBusDevice_Close(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QCanBusDevice::close")

	if signal := qt.GetSignal(C.GoString(ptrName), "close"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QCanBusDevice) ConnectClose(f func()) {
	defer qt.Recovering("connect QCanBusDevice::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "close", f)
	}
}

func (ptr *QCanBusDevice) DisconnectClose() {
	defer qt.Recovering("disconnect QCanBusDevice::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "close")
	}
}

func (ptr *QCanBusDevice) Close() {
	defer qt.Recovering("QCanBusDevice::close")

	if ptr.Pointer() != nil {
		C.QCanBusDevice_Close(ptr.Pointer())
	}
}

func (ptr *QCanBusDevice) ConfigurationParameter(key int) *core.QVariant {
	defer qt.Recovering("QCanBusDevice::configurationParameter")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QCanBusDevice_ConfigurationParameter(ptr.Pointer(), C.int(key)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QCanBusDevice) ConnectDevice() bool {
	defer qt.Recovering("QCanBusDevice::connectDevice")

	if ptr.Pointer() != nil {
		return C.QCanBusDevice_ConnectDevice(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCanBusDevice) DisconnectDevice() {
	defer qt.Recovering("QCanBusDevice::disconnectDevice")

	if ptr.Pointer() != nil {
		C.QCanBusDevice_DisconnectDevice(ptr.Pointer())
	}
}

func (ptr *QCanBusDevice) EnqueueOutgoingFrame(newFrame QCanBusFrame_ITF) {
	defer qt.Recovering("QCanBusDevice::enqueueOutgoingFrame")

	if ptr.Pointer() != nil {
		C.QCanBusDevice_EnqueueOutgoingFrame(ptr.Pointer(), PointerFromQCanBusFrame(newFrame))
	}
}

func (ptr *QCanBusDevice) Error() QCanBusDevice__CanBusError {
	defer qt.Recovering("QCanBusDevice::error")

	if ptr.Pointer() != nil {
		return QCanBusDevice__CanBusError(C.QCanBusDevice_Error(ptr.Pointer()))
	}
	return 0
}

//export callbackQCanBusDevice_ErrorOccurred
func callbackQCanBusDevice_ErrorOccurred(ptr unsafe.Pointer, ptrName *C.char, error C.int) {
	defer qt.Recovering("callback QCanBusDevice::errorOccurred")

	if signal := qt.GetSignal(C.GoString(ptrName), "errorOccurred"); signal != nil {
		signal.(func(QCanBusDevice__CanBusError))(QCanBusDevice__CanBusError(error))
	}

}

func (ptr *QCanBusDevice) ConnectErrorOccurred(f func(error QCanBusDevice__CanBusError)) {
	defer qt.Recovering("connect QCanBusDevice::errorOccurred")

	if ptr.Pointer() != nil {
		C.QCanBusDevice_ConnectErrorOccurred(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "errorOccurred", f)
	}
}

func (ptr *QCanBusDevice) DisconnectErrorOccurred() {
	defer qt.Recovering("disconnect QCanBusDevice::errorOccurred")

	if ptr.Pointer() != nil {
		C.QCanBusDevice_DisconnectErrorOccurred(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "errorOccurred")
	}
}

func (ptr *QCanBusDevice) ErrorOccurred(error QCanBusDevice__CanBusError) {
	defer qt.Recovering("QCanBusDevice::errorOccurred")

	if ptr.Pointer() != nil {
		C.QCanBusDevice_ErrorOccurred(ptr.Pointer(), C.int(error))
	}
}

func (ptr *QCanBusDevice) ErrorString() string {
	defer qt.Recovering("QCanBusDevice::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QCanBusDevice_ErrorString(ptr.Pointer()))
	}
	return ""
}

//export callbackQCanBusDevice_FramesReceived
func callbackQCanBusDevice_FramesReceived(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QCanBusDevice::framesReceived")

	if signal := qt.GetSignal(C.GoString(ptrName), "framesReceived"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QCanBusDevice) ConnectFramesReceived(f func()) {
	defer qt.Recovering("connect QCanBusDevice::framesReceived")

	if ptr.Pointer() != nil {
		C.QCanBusDevice_ConnectFramesReceived(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "framesReceived", f)
	}
}

func (ptr *QCanBusDevice) DisconnectFramesReceived() {
	defer qt.Recovering("disconnect QCanBusDevice::framesReceived")

	if ptr.Pointer() != nil {
		C.QCanBusDevice_DisconnectFramesReceived(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "framesReceived")
	}
}

func (ptr *QCanBusDevice) FramesReceived() {
	defer qt.Recovering("QCanBusDevice::framesReceived")

	if ptr.Pointer() != nil {
		C.QCanBusDevice_FramesReceived(ptr.Pointer())
	}
}

//export callbackQCanBusDevice_FramesWritten
func callbackQCanBusDevice_FramesWritten(ptr unsafe.Pointer, ptrName *C.char, framesCount C.longlong) {
	defer qt.Recovering("callback QCanBusDevice::framesWritten")

	if signal := qt.GetSignal(C.GoString(ptrName), "framesWritten"); signal != nil {
		signal.(func(int64))(int64(framesCount))
	}

}

func (ptr *QCanBusDevice) ConnectFramesWritten(f func(framesCount int64)) {
	defer qt.Recovering("connect QCanBusDevice::framesWritten")

	if ptr.Pointer() != nil {
		C.QCanBusDevice_ConnectFramesWritten(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "framesWritten", f)
	}
}

func (ptr *QCanBusDevice) DisconnectFramesWritten() {
	defer qt.Recovering("disconnect QCanBusDevice::framesWritten")

	if ptr.Pointer() != nil {
		C.QCanBusDevice_DisconnectFramesWritten(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "framesWritten")
	}
}

func (ptr *QCanBusDevice) FramesWritten(framesCount int64) {
	defer qt.Recovering("QCanBusDevice::framesWritten")

	if ptr.Pointer() != nil {
		C.QCanBusDevice_FramesWritten(ptr.Pointer(), C.longlong(framesCount))
	}
}

func (ptr *QCanBusDevice) HasOutgoingFrames() bool {
	defer qt.Recovering("QCanBusDevice::hasOutgoingFrames")

	if ptr.Pointer() != nil {
		return C.QCanBusDevice_HasOutgoingFrames(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQCanBusDevice_InterpretErrorFrame
func callbackQCanBusDevice_InterpretErrorFrame(ptr unsafe.Pointer, ptrName *C.char, frame unsafe.Pointer) *C.char {
	defer qt.Recovering("callback QCanBusDevice::interpretErrorFrame")

	if signal := qt.GetSignal(C.GoString(ptrName), "interpretErrorFrame"); signal != nil {
		return C.CString(signal.(func(*QCanBusFrame) string)(NewQCanBusFrameFromPointer(frame)))
	}

	return C.CString("")
}

func (ptr *QCanBusDevice) ConnectInterpretErrorFrame(f func(frame *QCanBusFrame) string) {
	defer qt.Recovering("connect QCanBusDevice::interpretErrorFrame")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "interpretErrorFrame", f)
	}
}

func (ptr *QCanBusDevice) DisconnectInterpretErrorFrame(frame QCanBusFrame_ITF) {
	defer qt.Recovering("disconnect QCanBusDevice::interpretErrorFrame")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "interpretErrorFrame")
	}
}

func (ptr *QCanBusDevice) InterpretErrorFrame(frame QCanBusFrame_ITF) string {
	defer qt.Recovering("QCanBusDevice::interpretErrorFrame")

	if ptr.Pointer() != nil {
		return C.GoString(C.QCanBusDevice_InterpretErrorFrame(ptr.Pointer(), PointerFromQCanBusFrame(frame)))
	}
	return ""
}

//export callbackQCanBusDevice_Open
func callbackQCanBusDevice_Open(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QCanBusDevice::open")

	if signal := qt.GetSignal(C.GoString(ptrName), "open"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(false))
}

func (ptr *QCanBusDevice) ConnectOpen(f func() bool) {
	defer qt.Recovering("connect QCanBusDevice::open")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "open", f)
	}
}

func (ptr *QCanBusDevice) DisconnectOpen() {
	defer qt.Recovering("disconnect QCanBusDevice::open")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "open")
	}
}

func (ptr *QCanBusDevice) Open() bool {
	defer qt.Recovering("QCanBusDevice::open")

	if ptr.Pointer() != nil {
		return C.QCanBusDevice_Open(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQCanBusDevice_SetConfigurationParameter
func callbackQCanBusDevice_SetConfigurationParameter(ptr unsafe.Pointer, ptrName *C.char, key C.int, value unsafe.Pointer) {
	defer qt.Recovering("callback QCanBusDevice::setConfigurationParameter")

	if signal := qt.GetSignal(C.GoString(ptrName), "setConfigurationParameter"); signal != nil {
		signal.(func(int, *core.QVariant))(int(key), core.NewQVariantFromPointer(value))
	} else {
		NewQCanBusDeviceFromPointer(ptr).SetConfigurationParameterDefault(int(key), core.NewQVariantFromPointer(value))
	}
}

func (ptr *QCanBusDevice) ConnectSetConfigurationParameter(f func(key int, value *core.QVariant)) {
	defer qt.Recovering("connect QCanBusDevice::setConfigurationParameter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setConfigurationParameter", f)
	}
}

func (ptr *QCanBusDevice) DisconnectSetConfigurationParameter() {
	defer qt.Recovering("disconnect QCanBusDevice::setConfigurationParameter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setConfigurationParameter")
	}
}

func (ptr *QCanBusDevice) SetConfigurationParameter(key int, value core.QVariant_ITF) {
	defer qt.Recovering("QCanBusDevice::setConfigurationParameter")

	if ptr.Pointer() != nil {
		C.QCanBusDevice_SetConfigurationParameter(ptr.Pointer(), C.int(key), core.PointerFromQVariant(value))
	}
}

func (ptr *QCanBusDevice) SetConfigurationParameterDefault(key int, value core.QVariant_ITF) {
	defer qt.Recovering("QCanBusDevice::setConfigurationParameter")

	if ptr.Pointer() != nil {
		C.QCanBusDevice_SetConfigurationParameterDefault(ptr.Pointer(), C.int(key), core.PointerFromQVariant(value))
	}
}

func (ptr *QCanBusDevice) SetError(errorText string, errorId QCanBusDevice__CanBusError) {
	defer qt.Recovering("QCanBusDevice::setError")

	if ptr.Pointer() != nil {
		var errorTextC = C.CString(errorText)
		defer C.free(unsafe.Pointer(errorTextC))
		C.QCanBusDevice_SetError(ptr.Pointer(), errorTextC, C.int(errorId))
	}
}

func (ptr *QCanBusDevice) SetState(newState QCanBusDevice__CanBusDeviceState) {
	defer qt.Recovering("QCanBusDevice::setState")

	if ptr.Pointer() != nil {
		C.QCanBusDevice_SetState(ptr.Pointer(), C.int(newState))
	}
}

func (ptr *QCanBusDevice) State() QCanBusDevice__CanBusDeviceState {
	defer qt.Recovering("QCanBusDevice::state")

	if ptr.Pointer() != nil {
		return QCanBusDevice__CanBusDeviceState(C.QCanBusDevice_State(ptr.Pointer()))
	}
	return 0
}

//export callbackQCanBusDevice_StateChanged
func callbackQCanBusDevice_StateChanged(ptr unsafe.Pointer, ptrName *C.char, state C.int) {
	defer qt.Recovering("callback QCanBusDevice::stateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stateChanged"); signal != nil {
		signal.(func(QCanBusDevice__CanBusDeviceState))(QCanBusDevice__CanBusDeviceState(state))
	}

}

func (ptr *QCanBusDevice) ConnectStateChanged(f func(state QCanBusDevice__CanBusDeviceState)) {
	defer qt.Recovering("connect QCanBusDevice::stateChanged")

	if ptr.Pointer() != nil {
		C.QCanBusDevice_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QCanBusDevice) DisconnectStateChanged() {
	defer qt.Recovering("disconnect QCanBusDevice::stateChanged")

	if ptr.Pointer() != nil {
		C.QCanBusDevice_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

func (ptr *QCanBusDevice) StateChanged(state QCanBusDevice__CanBusDeviceState) {
	defer qt.Recovering("QCanBusDevice::stateChanged")

	if ptr.Pointer() != nil {
		C.QCanBusDevice_StateChanged(ptr.Pointer(), C.int(state))
	}
}

//export callbackQCanBusDevice_WriteFrame
func callbackQCanBusDevice_WriteFrame(ptr unsafe.Pointer, ptrName *C.char, frame unsafe.Pointer) C.int {
	defer qt.Recovering("callback QCanBusDevice::writeFrame")

	if signal := qt.GetSignal(C.GoString(ptrName), "writeFrame"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*QCanBusFrame) bool)(NewQCanBusFrameFromPointer(frame))))
	}

	return C.int(qt.GoBoolToInt(false))
}

func (ptr *QCanBusDevice) ConnectWriteFrame(f func(frame *QCanBusFrame) bool) {
	defer qt.Recovering("connect QCanBusDevice::writeFrame")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "writeFrame", f)
	}
}

func (ptr *QCanBusDevice) DisconnectWriteFrame(frame QCanBusFrame_ITF) {
	defer qt.Recovering("disconnect QCanBusDevice::writeFrame")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "writeFrame")
	}
}

func (ptr *QCanBusDevice) WriteFrame(frame QCanBusFrame_ITF) bool {
	defer qt.Recovering("QCanBusDevice::writeFrame")

	if ptr.Pointer() != nil {
		return C.QCanBusDevice_WriteFrame(ptr.Pointer(), PointerFromQCanBusFrame(frame)) != 0
	}
	return false
}

//export callbackQCanBusDevice_TimerEvent
func callbackQCanBusDevice_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCanBusDevice::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQCanBusDeviceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QCanBusDevice) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QCanBusDevice::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QCanBusDevice) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QCanBusDevice::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QCanBusDevice) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCanBusDevice::timerEvent")

	if ptr.Pointer() != nil {
		C.QCanBusDevice_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCanBusDevice) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QCanBusDevice::timerEvent")

	if ptr.Pointer() != nil {
		C.QCanBusDevice_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQCanBusDevice_ChildEvent
func callbackQCanBusDevice_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCanBusDevice::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCanBusDeviceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCanBusDevice) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QCanBusDevice::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QCanBusDevice) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QCanBusDevice::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QCanBusDevice) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCanBusDevice::childEvent")

	if ptr.Pointer() != nil {
		C.QCanBusDevice_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCanBusDevice) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QCanBusDevice::childEvent")

	if ptr.Pointer() != nil {
		C.QCanBusDevice_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQCanBusDevice_ConnectNotify
func callbackQCanBusDevice_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QCanBusDevice::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQCanBusDeviceFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QCanBusDevice) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QCanBusDevice::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QCanBusDevice) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QCanBusDevice::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QCanBusDevice) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QCanBusDevice::connectNotify")

	if ptr.Pointer() != nil {
		C.QCanBusDevice_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QCanBusDevice) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QCanBusDevice::connectNotify")

	if ptr.Pointer() != nil {
		C.QCanBusDevice_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQCanBusDevice_CustomEvent
func callbackQCanBusDevice_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QCanBusDevice::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCanBusDeviceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCanBusDevice) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QCanBusDevice::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QCanBusDevice) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QCanBusDevice::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QCanBusDevice) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QCanBusDevice::customEvent")

	if ptr.Pointer() != nil {
		C.QCanBusDevice_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCanBusDevice) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QCanBusDevice::customEvent")

	if ptr.Pointer() != nil {
		C.QCanBusDevice_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQCanBusDevice_DeleteLater
func callbackQCanBusDevice_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QCanBusDevice::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQCanBusDeviceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QCanBusDevice) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QCanBusDevice::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QCanBusDevice) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QCanBusDevice::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QCanBusDevice) DeleteLater() {
	defer qt.Recovering("QCanBusDevice::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QCanBusDevice_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCanBusDevice) DeleteLaterDefault() {
	defer qt.Recovering("QCanBusDevice::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QCanBusDevice_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQCanBusDevice_DisconnectNotify
func callbackQCanBusDevice_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QCanBusDevice::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQCanBusDeviceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QCanBusDevice) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QCanBusDevice::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QCanBusDevice) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QCanBusDevice::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QCanBusDevice) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QCanBusDevice::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QCanBusDevice_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QCanBusDevice) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QCanBusDevice::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QCanBusDevice_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQCanBusDevice_Event
func callbackQCanBusDevice_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QCanBusDevice::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQCanBusDeviceFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QCanBusDevice) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QCanBusDevice::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QCanBusDevice) DisconnectEvent() {
	defer qt.Recovering("disconnect QCanBusDevice::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QCanBusDevice) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QCanBusDevice::event")

	if ptr.Pointer() != nil {
		return C.QCanBusDevice_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QCanBusDevice) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QCanBusDevice::event")

	if ptr.Pointer() != nil {
		return C.QCanBusDevice_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQCanBusDevice_EventFilter
func callbackQCanBusDevice_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QCanBusDevice::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQCanBusDeviceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QCanBusDevice) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QCanBusDevice::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QCanBusDevice) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QCanBusDevice::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QCanBusDevice) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QCanBusDevice::eventFilter")

	if ptr.Pointer() != nil {
		return C.QCanBusDevice_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QCanBusDevice) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QCanBusDevice::eventFilter")

	if ptr.Pointer() != nil {
		return C.QCanBusDevice_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQCanBusDevice_MetaObject
func callbackQCanBusDevice_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QCanBusDevice::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQCanBusDeviceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QCanBusDevice) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QCanBusDevice::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QCanBusDevice) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QCanBusDevice::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QCanBusDevice) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QCanBusDevice::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QCanBusDevice_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCanBusDevice) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QCanBusDevice::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QCanBusDevice_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QCanBusFactory struct {
	ptr unsafe.Pointer
}

type QCanBusFactory_ITF interface {
	QCanBusFactory_PTR() *QCanBusFactory
}

func (p *QCanBusFactory) QCanBusFactory_PTR() *QCanBusFactory {
	return p
}

func (p *QCanBusFactory) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QCanBusFactory) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQCanBusFactory(ptr QCanBusFactory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCanBusFactory_PTR().Pointer()
	}
	return nil
}

func NewQCanBusFactoryFromPointer(ptr unsafe.Pointer) *QCanBusFactory {
	var n = new(QCanBusFactory)
	n.SetPointer(ptr)
	return n
}

func newQCanBusFactoryFromPointer(ptr unsafe.Pointer) *QCanBusFactory {
	var n = NewQCanBusFactoryFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QCanBusFactory_") {
		n.SetObjectNameAbs("QCanBusFactory_" + qt.Identifier())
	}
	return n
}

func (ptr *QCanBusFactory) DestroyQCanBusFactory() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

//export callbackQCanBusFactory_CreateDevice
func callbackQCanBusFactory_CreateDevice(ptr unsafe.Pointer, ptrName *C.char, interfaceName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QCanBusFactory::createDevice")

	if signal := qt.GetSignal(C.GoString(ptrName), "createDevice"); signal != nil {
		return PointerFromQCanBusDevice(signal.(func(string) *QCanBusDevice)(C.GoString(interfaceName)))
	}

	return PointerFromQCanBusDevice(nil)
}

func (ptr *QCanBusFactory) ConnectCreateDevice(f func(interfaceName string) *QCanBusDevice) {
	defer qt.Recovering("connect QCanBusFactory::createDevice")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "createDevice", f)
	}
}

func (ptr *QCanBusFactory) DisconnectCreateDevice(interfaceName string) {
	defer qt.Recovering("disconnect QCanBusFactory::createDevice")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "createDevice")
	}
}

func (ptr *QCanBusFactory) CreateDevice(interfaceName string) *QCanBusDevice {
	defer qt.Recovering("QCanBusFactory::createDevice")

	if ptr.Pointer() != nil {
		var interfaceNameC = C.CString(interfaceName)
		defer C.free(unsafe.Pointer(interfaceNameC))
		return NewQCanBusDeviceFromPointer(C.QCanBusFactory_CreateDevice(ptr.Pointer(), interfaceNameC))
	}
	return nil
}

func (ptr *QCanBusFactory) ObjectNameAbs() string {
	defer qt.Recovering("QCanBusFactory::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QCanBusFactory_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QCanBusFactory) SetObjectNameAbs(name string) {
	defer qt.Recovering("QCanBusFactory::setObjectNameAbs")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QCanBusFactory_SetObjectNameAbs(ptr.Pointer(), nameC)
	}
}

//QCanBusFrame::FrameError
type QCanBusFrame__FrameError int64

var (
	QCanBusFrame__NoError                    = QCanBusFrame__FrameError(0)
	QCanBusFrame__TransmissionTimeoutError   = QCanBusFrame__FrameError(C.QCanBusFrame_TransmissionTimeoutError_Type())
	QCanBusFrame__LostArbitrationError       = QCanBusFrame__FrameError(C.QCanBusFrame_LostArbitrationError_Type())
	QCanBusFrame__ControllerError            = QCanBusFrame__FrameError(C.QCanBusFrame_ControllerError_Type())
	QCanBusFrame__ProtocolViolationError     = QCanBusFrame__FrameError(C.QCanBusFrame_ProtocolViolationError_Type())
	QCanBusFrame__TransceiverError           = QCanBusFrame__FrameError(C.QCanBusFrame_TransceiverError_Type())
	QCanBusFrame__MissingAcknowledgmentError = QCanBusFrame__FrameError(C.QCanBusFrame_MissingAcknowledgmentError_Type())
	QCanBusFrame__BusOffError                = QCanBusFrame__FrameError(C.QCanBusFrame_BusOffError_Type())
	QCanBusFrame__BusError                   = QCanBusFrame__FrameError(C.QCanBusFrame_BusError_Type())
	QCanBusFrame__ControllerRestartError     = QCanBusFrame__FrameError(C.QCanBusFrame_ControllerRestartError_Type())
	QCanBusFrame__UnknownError               = QCanBusFrame__FrameError(C.QCanBusFrame_UnknownError_Type())
	QCanBusFrame__AnyError                   = QCanBusFrame__FrameError(C.QCanBusFrame_AnyError_Type())
)

//QCanBusFrame::FrameType
type QCanBusFrame__FrameType int64

const (
	QCanBusFrame__UnknownFrame       = QCanBusFrame__FrameType(0x0)
	QCanBusFrame__DataFrame          = QCanBusFrame__FrameType(0x1)
	QCanBusFrame__ErrorFrame         = QCanBusFrame__FrameType(0x2)
	QCanBusFrame__RemoteRequestFrame = QCanBusFrame__FrameType(0x3)
	QCanBusFrame__InvalidFrame       = QCanBusFrame__FrameType(0x4)
)

type QCanBusFrame struct {
	ptr unsafe.Pointer
}

type QCanBusFrame_ITF interface {
	QCanBusFrame_PTR() *QCanBusFrame
}

func (p *QCanBusFrame) QCanBusFrame_PTR() *QCanBusFrame {
	return p
}

func (p *QCanBusFrame) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QCanBusFrame) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQCanBusFrame(ptr QCanBusFrame_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCanBusFrame_PTR().Pointer()
	}
	return nil
}

func NewQCanBusFrameFromPointer(ptr unsafe.Pointer) *QCanBusFrame {
	var n = new(QCanBusFrame)
	n.SetPointer(ptr)
	return n
}

func newQCanBusFrameFromPointer(ptr unsafe.Pointer) *QCanBusFrame {
	var n = NewQCanBusFrameFromPointer(ptr)
	return n
}

func (ptr *QCanBusFrame) DestroyQCanBusFrame() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

func NewQCanBusFrame(ty QCanBusFrame__FrameType) *QCanBusFrame {
	defer qt.Recovering("QCanBusFrame::QCanBusFrame")

	return newQCanBusFrameFromPointer(C.QCanBusFrame_NewQCanBusFrame(C.int(ty)))
}

func (ptr *QCanBusFrame) Error() QCanBusFrame__FrameError {
	defer qt.Recovering("QCanBusFrame::error")

	if ptr.Pointer() != nil {
		return QCanBusFrame__FrameError(C.QCanBusFrame_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCanBusFrame) FrameType() QCanBusFrame__FrameType {
	defer qt.Recovering("QCanBusFrame::frameType")

	if ptr.Pointer() != nil {
		return QCanBusFrame__FrameType(C.QCanBusFrame_FrameType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCanBusFrame) HasExtendedFrameFormat() bool {
	defer qt.Recovering("QCanBusFrame::hasExtendedFrameFormat")

	if ptr.Pointer() != nil {
		return C.QCanBusFrame_HasExtendedFrameFormat(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCanBusFrame) IsValid() bool {
	defer qt.Recovering("QCanBusFrame::isValid")

	if ptr.Pointer() != nil {
		return C.QCanBusFrame_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCanBusFrame) Payload() string {
	defer qt.Recovering("QCanBusFrame::payload")

	if ptr.Pointer() != nil {
		return C.GoString(C.QCanBusFrame_Payload(ptr.Pointer()))
	}
	return ""
}

func (ptr *QCanBusFrame) SetError(error QCanBusFrame__FrameError) {
	defer qt.Recovering("QCanBusFrame::setError")

	if ptr.Pointer() != nil {
		C.QCanBusFrame_SetError(ptr.Pointer(), C.int(error))
	}
}

func (ptr *QCanBusFrame) SetExtendedFrameFormat(isExtended bool) {
	defer qt.Recovering("QCanBusFrame::setExtendedFrameFormat")

	if ptr.Pointer() != nil {
		C.QCanBusFrame_SetExtendedFrameFormat(ptr.Pointer(), C.int(qt.GoBoolToInt(isExtended)))
	}
}

func (ptr *QCanBusFrame) SetFrameType(newType QCanBusFrame__FrameType) {
	defer qt.Recovering("QCanBusFrame::setFrameType")

	if ptr.Pointer() != nil {
		C.QCanBusFrame_SetFrameType(ptr.Pointer(), C.int(newType))
	}
}

func (ptr *QCanBusFrame) SetPayload(data string) {
	defer qt.Recovering("QCanBusFrame::setPayload")

	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		C.QCanBusFrame_SetPayload(ptr.Pointer(), dataC)
	}
}

type QModbusClient struct {
	QModbusDevice
}

type QModbusClient_ITF interface {
	QModbusDevice_ITF
	QModbusClient_PTR() *QModbusClient
}

func (p *QModbusClient) QModbusClient_PTR() *QModbusClient {
	return p
}

func (p *QModbusClient) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QModbusDevice_PTR().Pointer()
	}
	return nil
}

func (p *QModbusClient) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QModbusDevice_PTR().SetPointer(ptr)
	}
}

func PointerFromQModbusClient(ptr QModbusClient_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QModbusClient_PTR().Pointer()
	}
	return nil
}

func NewQModbusClientFromPointer(ptr unsafe.Pointer) *QModbusClient {
	var n = new(QModbusClient)
	n.SetPointer(ptr)
	return n
}

func newQModbusClientFromPointer(ptr unsafe.Pointer) *QModbusClient {
	var n = NewQModbusClientFromPointer(ptr)
	for len(n.ObjectName()) < len("QModbusClient_") {
		n.SetObjectName("QModbusClient_" + qt.Identifier())
	}
	return n
}

func (ptr *QModbusClient) DestroyQModbusClient() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

func (ptr *QModbusClient) Timeout() int {
	defer qt.Recovering("QModbusClient::timeout")

	if ptr.Pointer() != nil {
		return int(C.QModbusClient_Timeout(ptr.Pointer()))
	}
	return 0
}

func NewQModbusClient(parent core.QObject_ITF) *QModbusClient {
	defer qt.Recovering("QModbusClient::QModbusClient")

	return newQModbusClientFromPointer(C.QModbusClient_NewQModbusClient(core.PointerFromQObject(parent)))
}

func (ptr *QModbusClient) NumberOfRetries() int {
	defer qt.Recovering("QModbusClient::numberOfRetries")

	if ptr.Pointer() != nil {
		return int(C.QModbusClient_NumberOfRetries(ptr.Pointer()))
	}
	return 0
}

func (ptr *QModbusClient) SendRawRequest(request QModbusRequest_ITF, serverAddress int) *QModbusReply {
	defer qt.Recovering("QModbusClient::sendRawRequest")

	if ptr.Pointer() != nil {
		return NewQModbusReplyFromPointer(C.QModbusClient_SendRawRequest(ptr.Pointer(), PointerFromQModbusRequest(request), C.int(serverAddress)))
	}
	return nil
}

func (ptr *QModbusClient) SendReadRequest(read QModbusDataUnit_ITF, serverAddress int) *QModbusReply {
	defer qt.Recovering("QModbusClient::sendReadRequest")

	if ptr.Pointer() != nil {
		return NewQModbusReplyFromPointer(C.QModbusClient_SendReadRequest(ptr.Pointer(), PointerFromQModbusDataUnit(read), C.int(serverAddress)))
	}
	return nil
}

func (ptr *QModbusClient) SendReadWriteRequest(read QModbusDataUnit_ITF, write QModbusDataUnit_ITF, serverAddress int) *QModbusReply {
	defer qt.Recovering("QModbusClient::sendReadWriteRequest")

	if ptr.Pointer() != nil {
		return NewQModbusReplyFromPointer(C.QModbusClient_SendReadWriteRequest(ptr.Pointer(), PointerFromQModbusDataUnit(read), PointerFromQModbusDataUnit(write), C.int(serverAddress)))
	}
	return nil
}

func (ptr *QModbusClient) SendWriteRequest(write QModbusDataUnit_ITF, serverAddress int) *QModbusReply {
	defer qt.Recovering("QModbusClient::sendWriteRequest")

	if ptr.Pointer() != nil {
		return NewQModbusReplyFromPointer(C.QModbusClient_SendWriteRequest(ptr.Pointer(), PointerFromQModbusDataUnit(write), C.int(serverAddress)))
	}
	return nil
}

func (ptr *QModbusClient) SetNumberOfRetries(number int) {
	defer qt.Recovering("QModbusClient::setNumberOfRetries")

	if ptr.Pointer() != nil {
		C.QModbusClient_SetNumberOfRetries(ptr.Pointer(), C.int(number))
	}
}

func (ptr *QModbusClient) SetTimeout(newTimeout int) {
	defer qt.Recovering("QModbusClient::setTimeout")

	if ptr.Pointer() != nil {
		C.QModbusClient_SetTimeout(ptr.Pointer(), C.int(newTimeout))
	}
}

//export callbackQModbusClient_TimeoutChanged
func callbackQModbusClient_TimeoutChanged(ptr unsafe.Pointer, ptrName *C.char, newTimeout C.int) {
	defer qt.Recovering("callback QModbusClient::timeoutChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "timeoutChanged"); signal != nil {
		signal.(func(int))(int(newTimeout))
	}

}

func (ptr *QModbusClient) ConnectTimeoutChanged(f func(newTimeout int)) {
	defer qt.Recovering("connect QModbusClient::timeoutChanged")

	if ptr.Pointer() != nil {
		C.QModbusClient_ConnectTimeoutChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "timeoutChanged", f)
	}
}

func (ptr *QModbusClient) DisconnectTimeoutChanged() {
	defer qt.Recovering("disconnect QModbusClient::timeoutChanged")

	if ptr.Pointer() != nil {
		C.QModbusClient_DisconnectTimeoutChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "timeoutChanged")
	}
}

func (ptr *QModbusClient) TimeoutChanged(newTimeout int) {
	defer qt.Recovering("QModbusClient::timeoutChanged")

	if ptr.Pointer() != nil {
		C.QModbusClient_TimeoutChanged(ptr.Pointer(), C.int(newTimeout))
	}
}

//export callbackQModbusClient_ProcessPrivateResponse
func callbackQModbusClient_ProcessPrivateResponse(ptr unsafe.Pointer, ptrName *C.char, response unsafe.Pointer, data unsafe.Pointer) C.int {
	defer qt.Recovering("callback QModbusClient::processPrivateResponse")

	if signal := qt.GetSignal(C.GoString(ptrName), "processPrivateResponse"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*QModbusResponse, *QModbusDataUnit) bool)(NewQModbusResponseFromPointer(response), NewQModbusDataUnitFromPointer(data))))
	}

	return C.int(qt.GoBoolToInt(NewQModbusClientFromPointer(ptr).ProcessPrivateResponseDefault(NewQModbusResponseFromPointer(response), NewQModbusDataUnitFromPointer(data))))
}

func (ptr *QModbusClient) ConnectProcessPrivateResponse(f func(response *QModbusResponse, data *QModbusDataUnit) bool) {
	defer qt.Recovering("connect QModbusClient::processPrivateResponse")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "processPrivateResponse", f)
	}
}

func (ptr *QModbusClient) DisconnectProcessPrivateResponse() {
	defer qt.Recovering("disconnect QModbusClient::processPrivateResponse")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "processPrivateResponse")
	}
}

func (ptr *QModbusClient) ProcessPrivateResponse(response QModbusResponse_ITF, data QModbusDataUnit_ITF) bool {
	defer qt.Recovering("QModbusClient::processPrivateResponse")

	if ptr.Pointer() != nil {
		return C.QModbusClient_ProcessPrivateResponse(ptr.Pointer(), PointerFromQModbusResponse(response), PointerFromQModbusDataUnit(data)) != 0
	}
	return false
}

func (ptr *QModbusClient) ProcessPrivateResponseDefault(response QModbusResponse_ITF, data QModbusDataUnit_ITF) bool {
	defer qt.Recovering("QModbusClient::processPrivateResponse")

	if ptr.Pointer() != nil {
		return C.QModbusClient_ProcessPrivateResponseDefault(ptr.Pointer(), PointerFromQModbusResponse(response), PointerFromQModbusDataUnit(data)) != 0
	}
	return false
}

//export callbackQModbusClient_ProcessResponse
func callbackQModbusClient_ProcessResponse(ptr unsafe.Pointer, ptrName *C.char, response unsafe.Pointer, data unsafe.Pointer) C.int {
	defer qt.Recovering("callback QModbusClient::processResponse")

	if signal := qt.GetSignal(C.GoString(ptrName), "processResponse"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*QModbusResponse, *QModbusDataUnit) bool)(NewQModbusResponseFromPointer(response), NewQModbusDataUnitFromPointer(data))))
	}

	return C.int(qt.GoBoolToInt(NewQModbusClientFromPointer(ptr).ProcessResponseDefault(NewQModbusResponseFromPointer(response), NewQModbusDataUnitFromPointer(data))))
}

func (ptr *QModbusClient) ConnectProcessResponse(f func(response *QModbusResponse, data *QModbusDataUnit) bool) {
	defer qt.Recovering("connect QModbusClient::processResponse")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "processResponse", f)
	}
}

func (ptr *QModbusClient) DisconnectProcessResponse() {
	defer qt.Recovering("disconnect QModbusClient::processResponse")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "processResponse")
	}
}

func (ptr *QModbusClient) ProcessResponse(response QModbusResponse_ITF, data QModbusDataUnit_ITF) bool {
	defer qt.Recovering("QModbusClient::processResponse")

	if ptr.Pointer() != nil {
		return C.QModbusClient_ProcessResponse(ptr.Pointer(), PointerFromQModbusResponse(response), PointerFromQModbusDataUnit(data)) != 0
	}
	return false
}

func (ptr *QModbusClient) ProcessResponseDefault(response QModbusResponse_ITF, data QModbusDataUnit_ITF) bool {
	defer qt.Recovering("QModbusClient::processResponse")

	if ptr.Pointer() != nil {
		return C.QModbusClient_ProcessResponseDefault(ptr.Pointer(), PointerFromQModbusResponse(response), PointerFromQModbusDataUnit(data)) != 0
	}
	return false
}

//export callbackQModbusClient_Close
func callbackQModbusClient_Close(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QModbusClient::close")

	if signal := qt.GetSignal(C.GoString(ptrName), "close"); signal != nil {
		signal.(func())()
	} else {

	}
}

func (ptr *QModbusClient) ConnectClose(f func()) {
	defer qt.Recovering("connect QModbusClient::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "close", f)
	}
}

func (ptr *QModbusClient) DisconnectClose() {
	defer qt.Recovering("disconnect QModbusClient::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "close")
	}
}

func (ptr *QModbusClient) Close() {
	defer qt.Recovering("QModbusClient::close")

	if ptr.Pointer() != nil {
		C.QModbusClient_Close(ptr.Pointer())
	}
}

//export callbackQModbusClient_Open
func callbackQModbusClient_Open(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QModbusClient::open")

	if signal := qt.GetSignal(C.GoString(ptrName), "open"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(false))
}

func (ptr *QModbusClient) ConnectOpen(f func() bool) {
	defer qt.Recovering("connect QModbusClient::open")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "open", f)
	}
}

func (ptr *QModbusClient) DisconnectOpen() {
	defer qt.Recovering("disconnect QModbusClient::open")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "open")
	}
}

func (ptr *QModbusClient) Open() bool {
	defer qt.Recovering("QModbusClient::open")

	if ptr.Pointer() != nil {
		return C.QModbusClient_Open(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQModbusClient_TimerEvent
func callbackQModbusClient_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QModbusClient::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQModbusClientFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QModbusClient) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QModbusClient::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QModbusClient) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QModbusClient::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QModbusClient) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QModbusClient::timerEvent")

	if ptr.Pointer() != nil {
		C.QModbusClient_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QModbusClient) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QModbusClient::timerEvent")

	if ptr.Pointer() != nil {
		C.QModbusClient_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQModbusClient_ChildEvent
func callbackQModbusClient_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QModbusClient::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQModbusClientFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QModbusClient) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QModbusClient::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QModbusClient) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QModbusClient::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QModbusClient) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QModbusClient::childEvent")

	if ptr.Pointer() != nil {
		C.QModbusClient_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QModbusClient) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QModbusClient::childEvent")

	if ptr.Pointer() != nil {
		C.QModbusClient_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQModbusClient_ConnectNotify
func callbackQModbusClient_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QModbusClient::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQModbusClientFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QModbusClient) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QModbusClient::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QModbusClient) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QModbusClient::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QModbusClient) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QModbusClient::connectNotify")

	if ptr.Pointer() != nil {
		C.QModbusClient_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QModbusClient) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QModbusClient::connectNotify")

	if ptr.Pointer() != nil {
		C.QModbusClient_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQModbusClient_CustomEvent
func callbackQModbusClient_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QModbusClient::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQModbusClientFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QModbusClient) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QModbusClient::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QModbusClient) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QModbusClient::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QModbusClient) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QModbusClient::customEvent")

	if ptr.Pointer() != nil {
		C.QModbusClient_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QModbusClient) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QModbusClient::customEvent")

	if ptr.Pointer() != nil {
		C.QModbusClient_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQModbusClient_DeleteLater
func callbackQModbusClient_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QModbusClient::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQModbusClientFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QModbusClient) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QModbusClient::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QModbusClient) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QModbusClient::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QModbusClient) DeleteLater() {
	defer qt.Recovering("QModbusClient::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QModbusClient_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QModbusClient) DeleteLaterDefault() {
	defer qt.Recovering("QModbusClient::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QModbusClient_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQModbusClient_DisconnectNotify
func callbackQModbusClient_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QModbusClient::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQModbusClientFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QModbusClient) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QModbusClient::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QModbusClient) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QModbusClient::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QModbusClient) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QModbusClient::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QModbusClient_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QModbusClient) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QModbusClient::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QModbusClient_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQModbusClient_Event
func callbackQModbusClient_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QModbusClient::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQModbusClientFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QModbusClient) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QModbusClient::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QModbusClient) DisconnectEvent() {
	defer qt.Recovering("disconnect QModbusClient::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QModbusClient) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QModbusClient::event")

	if ptr.Pointer() != nil {
		return C.QModbusClient_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QModbusClient) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QModbusClient::event")

	if ptr.Pointer() != nil {
		return C.QModbusClient_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQModbusClient_EventFilter
func callbackQModbusClient_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QModbusClient::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQModbusClientFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QModbusClient) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QModbusClient::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QModbusClient) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QModbusClient::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QModbusClient) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QModbusClient::eventFilter")

	if ptr.Pointer() != nil {
		return C.QModbusClient_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QModbusClient) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QModbusClient::eventFilter")

	if ptr.Pointer() != nil {
		return C.QModbusClient_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQModbusClient_MetaObject
func callbackQModbusClient_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QModbusClient::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQModbusClientFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QModbusClient) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QModbusClient::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QModbusClient) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QModbusClient::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QModbusClient) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QModbusClient::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QModbusClient_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QModbusClient) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QModbusClient::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QModbusClient_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QModbusDataUnit::RegisterType
type QModbusDataUnit__RegisterType int64

const (
	QModbusDataUnit__Invalid          = QModbusDataUnit__RegisterType(0)
	QModbusDataUnit__DiscreteInputs   = QModbusDataUnit__RegisterType(1)
	QModbusDataUnit__Coils            = QModbusDataUnit__RegisterType(2)
	QModbusDataUnit__InputRegisters   = QModbusDataUnit__RegisterType(3)
	QModbusDataUnit__HoldingRegisters = QModbusDataUnit__RegisterType(4)
)

type QModbusDataUnit struct {
	ptr unsafe.Pointer
}

type QModbusDataUnit_ITF interface {
	QModbusDataUnit_PTR() *QModbusDataUnit
}

func (p *QModbusDataUnit) QModbusDataUnit_PTR() *QModbusDataUnit {
	return p
}

func (p *QModbusDataUnit) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QModbusDataUnit) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQModbusDataUnit(ptr QModbusDataUnit_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QModbusDataUnit_PTR().Pointer()
	}
	return nil
}

func NewQModbusDataUnitFromPointer(ptr unsafe.Pointer) *QModbusDataUnit {
	var n = new(QModbusDataUnit)
	n.SetPointer(ptr)
	return n
}

func newQModbusDataUnitFromPointer(ptr unsafe.Pointer) *QModbusDataUnit {
	var n = NewQModbusDataUnitFromPointer(ptr)
	return n
}

func (ptr *QModbusDataUnit) DestroyQModbusDataUnit() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

func NewQModbusDataUnit() *QModbusDataUnit {
	defer qt.Recovering("QModbusDataUnit::QModbusDataUnit")

	return newQModbusDataUnitFromPointer(C.QModbusDataUnit_NewQModbusDataUnit())
}

func NewQModbusDataUnit2(ty QModbusDataUnit__RegisterType) *QModbusDataUnit {
	defer qt.Recovering("QModbusDataUnit::QModbusDataUnit")

	return newQModbusDataUnitFromPointer(C.QModbusDataUnit_NewQModbusDataUnit2(C.int(ty)))
}

func (ptr *QModbusDataUnit) IsValid() bool {
	defer qt.Recovering("QModbusDataUnit::isValid")

	if ptr.Pointer() != nil {
		return C.QModbusDataUnit_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QModbusDataUnit) RegisterType() QModbusDataUnit__RegisterType {
	defer qt.Recovering("QModbusDataUnit::registerType")

	if ptr.Pointer() != nil {
		return QModbusDataUnit__RegisterType(C.QModbusDataUnit_RegisterType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QModbusDataUnit) SetRegisterType(ty QModbusDataUnit__RegisterType) {
	defer qt.Recovering("QModbusDataUnit::setRegisterType")

	if ptr.Pointer() != nil {
		C.QModbusDataUnit_SetRegisterType(ptr.Pointer(), C.int(ty))
	}
}

func (ptr *QModbusDataUnit) SetStartAddress(address int) {
	defer qt.Recovering("QModbusDataUnit::setStartAddress")

	if ptr.Pointer() != nil {
		C.QModbusDataUnit_SetStartAddress(ptr.Pointer(), C.int(address))
	}
}

func (ptr *QModbusDataUnit) StartAddress() int {
	defer qt.Recovering("QModbusDataUnit::startAddress")

	if ptr.Pointer() != nil {
		return int(C.QModbusDataUnit_StartAddress(ptr.Pointer()))
	}
	return 0
}

//QModbusDevice::ConnectionParameter
type QModbusDevice__ConnectionParameter int64

const (
	QModbusDevice__SerialPortNameParameter = QModbusDevice__ConnectionParameter(0)
	QModbusDevice__SerialParityParameter   = QModbusDevice__ConnectionParameter(1)
	QModbusDevice__SerialBaudRateParameter = QModbusDevice__ConnectionParameter(2)
	QModbusDevice__SerialDataBitsParameter = QModbusDevice__ConnectionParameter(3)
	QModbusDevice__SerialStopBitsParameter = QModbusDevice__ConnectionParameter(4)
	QModbusDevice__NetworkPortParameter    = QModbusDevice__ConnectionParameter(5)
	QModbusDevice__NetworkAddressParameter = QModbusDevice__ConnectionParameter(6)
	QModbusDevice__UserParameter           = QModbusDevice__ConnectionParameter(0x100)
)

//QModbusDevice::Error
type QModbusDevice__Error int64

const (
	QModbusDevice__NoError            = QModbusDevice__Error(0)
	QModbusDevice__ReadError          = QModbusDevice__Error(1)
	QModbusDevice__WriteError         = QModbusDevice__Error(2)
	QModbusDevice__ConnectionError    = QModbusDevice__Error(3)
	QModbusDevice__ConfigurationError = QModbusDevice__Error(4)
	QModbusDevice__TimeoutError       = QModbusDevice__Error(5)
	QModbusDevice__ProtocolError      = QModbusDevice__Error(6)
	QModbusDevice__ReplyAbortedError  = QModbusDevice__Error(7)
	QModbusDevice__UnknownError       = QModbusDevice__Error(8)
)

//QModbusDevice::State
type QModbusDevice__State int64

const (
	QModbusDevice__UnconnectedState = QModbusDevice__State(0)
	QModbusDevice__ConnectingState  = QModbusDevice__State(1)
	QModbusDevice__ConnectedState   = QModbusDevice__State(2)
	QModbusDevice__ClosingState     = QModbusDevice__State(3)
)

type QModbusDevice struct {
	core.QObject
}

type QModbusDevice_ITF interface {
	core.QObject_ITF
	QModbusDevice_PTR() *QModbusDevice
}

func (p *QModbusDevice) QModbusDevice_PTR() *QModbusDevice {
	return p
}

func (p *QModbusDevice) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QModbusDevice) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQModbusDevice(ptr QModbusDevice_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QModbusDevice_PTR().Pointer()
	}
	return nil
}

func NewQModbusDeviceFromPointer(ptr unsafe.Pointer) *QModbusDevice {
	var n = new(QModbusDevice)
	n.SetPointer(ptr)
	return n
}

func newQModbusDeviceFromPointer(ptr unsafe.Pointer) *QModbusDevice {
	var n = NewQModbusDeviceFromPointer(ptr)
	for len(n.ObjectName()) < len("QModbusDevice_") {
		n.SetObjectName("QModbusDevice_" + qt.Identifier())
	}
	return n
}

func NewQModbusDevice(parent core.QObject_ITF) *QModbusDevice {
	defer qt.Recovering("QModbusDevice::QModbusDevice")

	return newQModbusDeviceFromPointer(C.QModbusDevice_NewQModbusDevice(core.PointerFromQObject(parent)))
}

//export callbackQModbusDevice_Close
func callbackQModbusDevice_Close(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QModbusDevice::close")

	if signal := qt.GetSignal(C.GoString(ptrName), "close"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QModbusDevice) ConnectClose(f func()) {
	defer qt.Recovering("connect QModbusDevice::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "close", f)
	}
}

func (ptr *QModbusDevice) DisconnectClose() {
	defer qt.Recovering("disconnect QModbusDevice::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "close")
	}
}

func (ptr *QModbusDevice) Close() {
	defer qt.Recovering("QModbusDevice::close")

	if ptr.Pointer() != nil {
		C.QModbusDevice_Close(ptr.Pointer())
	}
}

func (ptr *QModbusDevice) ConnectDevice() bool {
	defer qt.Recovering("QModbusDevice::connectDevice")

	if ptr.Pointer() != nil {
		return C.QModbusDevice_ConnectDevice(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QModbusDevice) ConnectionParameter(parameter int) *core.QVariant {
	defer qt.Recovering("QModbusDevice::connectionParameter")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QModbusDevice_ConnectionParameter(ptr.Pointer(), C.int(parameter)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QModbusDevice) DisconnectDevice() {
	defer qt.Recovering("QModbusDevice::disconnectDevice")

	if ptr.Pointer() != nil {
		C.QModbusDevice_DisconnectDevice(ptr.Pointer())
	}
}

func (ptr *QModbusDevice) Error() QModbusDevice__Error {
	defer qt.Recovering("QModbusDevice::error")

	if ptr.Pointer() != nil {
		return QModbusDevice__Error(C.QModbusDevice_Error(ptr.Pointer()))
	}
	return 0
}

//export callbackQModbusDevice_ErrorOccurred
func callbackQModbusDevice_ErrorOccurred(ptr unsafe.Pointer, ptrName *C.char, error C.int) {
	defer qt.Recovering("callback QModbusDevice::errorOccurred")

	if signal := qt.GetSignal(C.GoString(ptrName), "errorOccurred"); signal != nil {
		signal.(func(QModbusDevice__Error))(QModbusDevice__Error(error))
	}

}

func (ptr *QModbusDevice) ConnectErrorOccurred(f func(error QModbusDevice__Error)) {
	defer qt.Recovering("connect QModbusDevice::errorOccurred")

	if ptr.Pointer() != nil {
		C.QModbusDevice_ConnectErrorOccurred(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "errorOccurred", f)
	}
}

func (ptr *QModbusDevice) DisconnectErrorOccurred() {
	defer qt.Recovering("disconnect QModbusDevice::errorOccurred")

	if ptr.Pointer() != nil {
		C.QModbusDevice_DisconnectErrorOccurred(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "errorOccurred")
	}
}

func (ptr *QModbusDevice) ErrorOccurred(error QModbusDevice__Error) {
	defer qt.Recovering("QModbusDevice::errorOccurred")

	if ptr.Pointer() != nil {
		C.QModbusDevice_ErrorOccurred(ptr.Pointer(), C.int(error))
	}
}

func (ptr *QModbusDevice) ErrorString() string {
	defer qt.Recovering("QModbusDevice::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QModbusDevice_ErrorString(ptr.Pointer()))
	}
	return ""
}

//export callbackQModbusDevice_Open
func callbackQModbusDevice_Open(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QModbusDevice::open")

	if signal := qt.GetSignal(C.GoString(ptrName), "open"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(false))
}

func (ptr *QModbusDevice) ConnectOpen(f func() bool) {
	defer qt.Recovering("connect QModbusDevice::open")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "open", f)
	}
}

func (ptr *QModbusDevice) DisconnectOpen() {
	defer qt.Recovering("disconnect QModbusDevice::open")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "open")
	}
}

func (ptr *QModbusDevice) Open() bool {
	defer qt.Recovering("QModbusDevice::open")

	if ptr.Pointer() != nil {
		return C.QModbusDevice_Open(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QModbusDevice) SetConnectionParameter(parameter int, value core.QVariant_ITF) {
	defer qt.Recovering("QModbusDevice::setConnectionParameter")

	if ptr.Pointer() != nil {
		C.QModbusDevice_SetConnectionParameter(ptr.Pointer(), C.int(parameter), core.PointerFromQVariant(value))
	}
}

func (ptr *QModbusDevice) SetError(errorText string, error QModbusDevice__Error) {
	defer qt.Recovering("QModbusDevice::setError")

	if ptr.Pointer() != nil {
		var errorTextC = C.CString(errorText)
		defer C.free(unsafe.Pointer(errorTextC))
		C.QModbusDevice_SetError(ptr.Pointer(), errorTextC, C.int(error))
	}
}

func (ptr *QModbusDevice) SetState(newState QModbusDevice__State) {
	defer qt.Recovering("QModbusDevice::setState")

	if ptr.Pointer() != nil {
		C.QModbusDevice_SetState(ptr.Pointer(), C.int(newState))
	}
}

func (ptr *QModbusDevice) State() QModbusDevice__State {
	defer qt.Recovering("QModbusDevice::state")

	if ptr.Pointer() != nil {
		return QModbusDevice__State(C.QModbusDevice_State(ptr.Pointer()))
	}
	return 0
}

//export callbackQModbusDevice_StateChanged
func callbackQModbusDevice_StateChanged(ptr unsafe.Pointer, ptrName *C.char, state C.int) {
	defer qt.Recovering("callback QModbusDevice::stateChanged")

	if signal := qt.GetSignal(C.GoString(ptrName), "stateChanged"); signal != nil {
		signal.(func(QModbusDevice__State))(QModbusDevice__State(state))
	}

}

func (ptr *QModbusDevice) ConnectStateChanged(f func(state QModbusDevice__State)) {
	defer qt.Recovering("connect QModbusDevice::stateChanged")

	if ptr.Pointer() != nil {
		C.QModbusDevice_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "stateChanged", f)
	}
}

func (ptr *QModbusDevice) DisconnectStateChanged() {
	defer qt.Recovering("disconnect QModbusDevice::stateChanged")

	if ptr.Pointer() != nil {
		C.QModbusDevice_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "stateChanged")
	}
}

func (ptr *QModbusDevice) StateChanged(state QModbusDevice__State) {
	defer qt.Recovering("QModbusDevice::stateChanged")

	if ptr.Pointer() != nil {
		C.QModbusDevice_StateChanged(ptr.Pointer(), C.int(state))
	}
}

func (ptr *QModbusDevice) DestroyQModbusDevice() {
	defer qt.Recovering("QModbusDevice::~QModbusDevice")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QModbusDevice_DestroyQModbusDevice(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQModbusDevice_TimerEvent
func callbackQModbusDevice_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QModbusDevice::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQModbusDeviceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QModbusDevice) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QModbusDevice::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QModbusDevice) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QModbusDevice::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QModbusDevice) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QModbusDevice::timerEvent")

	if ptr.Pointer() != nil {
		C.QModbusDevice_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QModbusDevice) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QModbusDevice::timerEvent")

	if ptr.Pointer() != nil {
		C.QModbusDevice_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQModbusDevice_ChildEvent
func callbackQModbusDevice_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QModbusDevice::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQModbusDeviceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QModbusDevice) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QModbusDevice::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QModbusDevice) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QModbusDevice::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QModbusDevice) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QModbusDevice::childEvent")

	if ptr.Pointer() != nil {
		C.QModbusDevice_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QModbusDevice) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QModbusDevice::childEvent")

	if ptr.Pointer() != nil {
		C.QModbusDevice_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQModbusDevice_ConnectNotify
func callbackQModbusDevice_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QModbusDevice::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQModbusDeviceFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QModbusDevice) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QModbusDevice::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QModbusDevice) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QModbusDevice::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QModbusDevice) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QModbusDevice::connectNotify")

	if ptr.Pointer() != nil {
		C.QModbusDevice_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QModbusDevice) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QModbusDevice::connectNotify")

	if ptr.Pointer() != nil {
		C.QModbusDevice_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQModbusDevice_CustomEvent
func callbackQModbusDevice_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QModbusDevice::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQModbusDeviceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QModbusDevice) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QModbusDevice::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QModbusDevice) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QModbusDevice::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QModbusDevice) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QModbusDevice::customEvent")

	if ptr.Pointer() != nil {
		C.QModbusDevice_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QModbusDevice) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QModbusDevice::customEvent")

	if ptr.Pointer() != nil {
		C.QModbusDevice_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQModbusDevice_DeleteLater
func callbackQModbusDevice_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QModbusDevice::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQModbusDeviceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QModbusDevice) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QModbusDevice::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QModbusDevice) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QModbusDevice::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QModbusDevice) DeleteLater() {
	defer qt.Recovering("QModbusDevice::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QModbusDevice_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QModbusDevice) DeleteLaterDefault() {
	defer qt.Recovering("QModbusDevice::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QModbusDevice_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQModbusDevice_DisconnectNotify
func callbackQModbusDevice_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QModbusDevice::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQModbusDeviceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QModbusDevice) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QModbusDevice::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QModbusDevice) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QModbusDevice::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QModbusDevice) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QModbusDevice::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QModbusDevice_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QModbusDevice) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QModbusDevice::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QModbusDevice_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQModbusDevice_Event
func callbackQModbusDevice_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QModbusDevice::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQModbusDeviceFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QModbusDevice) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QModbusDevice::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QModbusDevice) DisconnectEvent() {
	defer qt.Recovering("disconnect QModbusDevice::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QModbusDevice) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QModbusDevice::event")

	if ptr.Pointer() != nil {
		return C.QModbusDevice_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QModbusDevice) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QModbusDevice::event")

	if ptr.Pointer() != nil {
		return C.QModbusDevice_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQModbusDevice_EventFilter
func callbackQModbusDevice_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QModbusDevice::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQModbusDeviceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QModbusDevice) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QModbusDevice::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QModbusDevice) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QModbusDevice::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QModbusDevice) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QModbusDevice::eventFilter")

	if ptr.Pointer() != nil {
		return C.QModbusDevice_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QModbusDevice) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QModbusDevice::eventFilter")

	if ptr.Pointer() != nil {
		return C.QModbusDevice_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQModbusDevice_MetaObject
func callbackQModbusDevice_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QModbusDevice::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQModbusDeviceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QModbusDevice) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QModbusDevice::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QModbusDevice) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QModbusDevice::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QModbusDevice) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QModbusDevice::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QModbusDevice_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QModbusDevice) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QModbusDevice::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QModbusDevice_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QModbusDeviceIdentification::ConformityLevel
type QModbusDeviceIdentification__ConformityLevel int64

const (
	QModbusDeviceIdentification__BasicConformityLevel              = QModbusDeviceIdentification__ConformityLevel(0x01)
	QModbusDeviceIdentification__RegularConformityLevel            = QModbusDeviceIdentification__ConformityLevel(0x02)
	QModbusDeviceIdentification__ExtendedConformityLevel           = QModbusDeviceIdentification__ConformityLevel(0x03)
	QModbusDeviceIdentification__BasicIndividualConformityLevel    = QModbusDeviceIdentification__ConformityLevel(0x81)
	QModbusDeviceIdentification__RegularIndividualConformityLevel  = QModbusDeviceIdentification__ConformityLevel(0x82)
	QModbusDeviceIdentification__ExtendedIndividualConformityLevel = QModbusDeviceIdentification__ConformityLevel(0x83)
)

//QModbusDeviceIdentification::ObjectId
type QModbusDeviceIdentification__ObjectId int64

const (
	QModbusDeviceIdentification__VendorNameObjectId          = QModbusDeviceIdentification__ObjectId(0x00)
	QModbusDeviceIdentification__ProductCodeObjectId         = QModbusDeviceIdentification__ObjectId(0x01)
	QModbusDeviceIdentification__MajorMinorRevisionObjectId  = QModbusDeviceIdentification__ObjectId(0x02)
	QModbusDeviceIdentification__VendorUrlObjectId           = QModbusDeviceIdentification__ObjectId(0x03)
	QModbusDeviceIdentification__ProductNameObjectId         = QModbusDeviceIdentification__ObjectId(0x04)
	QModbusDeviceIdentification__ModelNameObjectId           = QModbusDeviceIdentification__ObjectId(0x05)
	QModbusDeviceIdentification__UserApplicationNameObjectId = QModbusDeviceIdentification__ObjectId(0x06)
	QModbusDeviceIdentification__ReservedObjectId            = QModbusDeviceIdentification__ObjectId(0x07)
	QModbusDeviceIdentification__ProductDependentObjectId    = QModbusDeviceIdentification__ObjectId(0x80)
	QModbusDeviceIdentification__UndefinedObjectId           = QModbusDeviceIdentification__ObjectId(0x100)
)

//QModbusDeviceIdentification::ReadDeviceIdCode
type QModbusDeviceIdentification__ReadDeviceIdCode int64

const (
	QModbusDeviceIdentification__BasicReadDeviceIdCode      = QModbusDeviceIdentification__ReadDeviceIdCode(0x01)
	QModbusDeviceIdentification__RegularReadDeviceIdCode    = QModbusDeviceIdentification__ReadDeviceIdCode(0x02)
	QModbusDeviceIdentification__ExtendedReadDeviceIdCode   = QModbusDeviceIdentification__ReadDeviceIdCode(0x03)
	QModbusDeviceIdentification__IndividualReadDeviceIdCode = QModbusDeviceIdentification__ReadDeviceIdCode(0x04)
)

type QModbusDeviceIdentification struct {
	ptr unsafe.Pointer
}

type QModbusDeviceIdentification_ITF interface {
	QModbusDeviceIdentification_PTR() *QModbusDeviceIdentification
}

func (p *QModbusDeviceIdentification) QModbusDeviceIdentification_PTR() *QModbusDeviceIdentification {
	return p
}

func (p *QModbusDeviceIdentification) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QModbusDeviceIdentification) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQModbusDeviceIdentification(ptr QModbusDeviceIdentification_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QModbusDeviceIdentification_PTR().Pointer()
	}
	return nil
}

func NewQModbusDeviceIdentificationFromPointer(ptr unsafe.Pointer) *QModbusDeviceIdentification {
	var n = new(QModbusDeviceIdentification)
	n.SetPointer(ptr)
	return n
}

func newQModbusDeviceIdentificationFromPointer(ptr unsafe.Pointer) *QModbusDeviceIdentification {
	var n = NewQModbusDeviceIdentificationFromPointer(ptr)
	return n
}

func (ptr *QModbusDeviceIdentification) DestroyQModbusDeviceIdentification() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

func NewQModbusDeviceIdentification() *QModbusDeviceIdentification {
	defer qt.Recovering("QModbusDeviceIdentification::QModbusDeviceIdentification")

	return newQModbusDeviceIdentificationFromPointer(C.QModbusDeviceIdentification_NewQModbusDeviceIdentification())
}

func (ptr *QModbusDeviceIdentification) ConformityLevel() QModbusDeviceIdentification__ConformityLevel {
	defer qt.Recovering("QModbusDeviceIdentification::conformityLevel")

	if ptr.Pointer() != nil {
		return QModbusDeviceIdentification__ConformityLevel(C.QModbusDeviceIdentification_ConformityLevel(ptr.Pointer()))
	}
	return 0
}

func (ptr *QModbusDeviceIdentification) IsValid() bool {
	defer qt.Recovering("QModbusDeviceIdentification::isValid")

	if ptr.Pointer() != nil {
		return C.QModbusDeviceIdentification_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QModbusDeviceIdentification) SetConformityLevel(level QModbusDeviceIdentification__ConformityLevel) {
	defer qt.Recovering("QModbusDeviceIdentification::setConformityLevel")

	if ptr.Pointer() != nil {
		C.QModbusDeviceIdentification_SetConformityLevel(ptr.Pointer(), C.int(level))
	}
}

type QModbusExceptionResponse struct {
	QModbusResponse
}

type QModbusExceptionResponse_ITF interface {
	QModbusResponse_ITF
	QModbusExceptionResponse_PTR() *QModbusExceptionResponse
}

func (p *QModbusExceptionResponse) QModbusExceptionResponse_PTR() *QModbusExceptionResponse {
	return p
}

func (p *QModbusExceptionResponse) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QModbusResponse_PTR().Pointer()
	}
	return nil
}

func (p *QModbusExceptionResponse) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QModbusResponse_PTR().SetPointer(ptr)
	}
}

func PointerFromQModbusExceptionResponse(ptr QModbusExceptionResponse_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QModbusExceptionResponse_PTR().Pointer()
	}
	return nil
}

func NewQModbusExceptionResponseFromPointer(ptr unsafe.Pointer) *QModbusExceptionResponse {
	var n = new(QModbusExceptionResponse)
	n.SetPointer(ptr)
	return n
}

func newQModbusExceptionResponseFromPointer(ptr unsafe.Pointer) *QModbusExceptionResponse {
	var n = NewQModbusExceptionResponseFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QModbusExceptionResponse_") {
		n.SetObjectNameAbs("QModbusExceptionResponse_" + qt.Identifier())
	}
	return n
}

func (ptr *QModbusExceptionResponse) DestroyQModbusExceptionResponse() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

func NewQModbusExceptionResponse() *QModbusExceptionResponse {
	defer qt.Recovering("QModbusExceptionResponse::QModbusExceptionResponse")

	return newQModbusExceptionResponseFromPointer(C.QModbusExceptionResponse_NewQModbusExceptionResponse())
}

func NewQModbusExceptionResponse3(code QModbusPdu__FunctionCode, ec QModbusPdu__ExceptionCode) *QModbusExceptionResponse {
	defer qt.Recovering("QModbusExceptionResponse::QModbusExceptionResponse")

	return newQModbusExceptionResponseFromPointer(C.QModbusExceptionResponse_NewQModbusExceptionResponse3(C.int(code), C.int(ec)))
}

func NewQModbusExceptionResponse2(pdu QModbusPdu_ITF) *QModbusExceptionResponse {
	defer qt.Recovering("QModbusExceptionResponse::QModbusExceptionResponse")

	return newQModbusExceptionResponseFromPointer(C.QModbusExceptionResponse_NewQModbusExceptionResponse2(PointerFromQModbusPdu(pdu)))
}

func (ptr *QModbusExceptionResponse) SetExceptionCode(ec QModbusPdu__ExceptionCode) {
	defer qt.Recovering("QModbusExceptionResponse::setExceptionCode")

	if ptr.Pointer() != nil {
		C.QModbusExceptionResponse_SetExceptionCode(ptr.Pointer(), C.int(ec))
	}
}

//export callbackQModbusExceptionResponse_SetFunctionCode
func callbackQModbusExceptionResponse_SetFunctionCode(ptr unsafe.Pointer, ptrName *C.char, c C.int) {
	defer qt.Recovering("callback QModbusExceptionResponse::setFunctionCode")

	if signal := qt.GetSignal(C.GoString(ptrName), "setFunctionCode"); signal != nil {
		signal.(func(QModbusPdu__FunctionCode))(QModbusPdu__FunctionCode(c))
	} else {
		NewQModbusExceptionResponseFromPointer(ptr).SetFunctionCodeDefault(QModbusPdu__FunctionCode(c))
	}
}

func (ptr *QModbusExceptionResponse) ConnectSetFunctionCode(f func(c QModbusPdu__FunctionCode)) {
	defer qt.Recovering("connect QModbusExceptionResponse::setFunctionCode")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setFunctionCode", f)
	}
}

func (ptr *QModbusExceptionResponse) DisconnectSetFunctionCode() {
	defer qt.Recovering("disconnect QModbusExceptionResponse::setFunctionCode")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setFunctionCode")
	}
}

func (ptr *QModbusExceptionResponse) SetFunctionCode(c QModbusPdu__FunctionCode) {
	defer qt.Recovering("QModbusExceptionResponse::setFunctionCode")

	if ptr.Pointer() != nil {
		C.QModbusExceptionResponse_SetFunctionCode(ptr.Pointer(), C.int(c))
	}
}

func (ptr *QModbusExceptionResponse) SetFunctionCodeDefault(c QModbusPdu__FunctionCode) {
	defer qt.Recovering("QModbusExceptionResponse::setFunctionCode")

	if ptr.Pointer() != nil {
		C.QModbusExceptionResponse_SetFunctionCodeDefault(ptr.Pointer(), C.int(c))
	}
}

func (ptr *QModbusExceptionResponse) ObjectNameAbs() string {
	defer qt.Recovering("QModbusExceptionResponse::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QModbusExceptionResponse_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QModbusExceptionResponse) SetObjectNameAbs(name string) {
	defer qt.Recovering("QModbusExceptionResponse::setObjectNameAbs")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QModbusExceptionResponse_SetObjectNameAbs(ptr.Pointer(), nameC)
	}
}

//QModbusPdu::ExceptionCode
type QModbusPdu__ExceptionCode int64

const (
	QModbusPdu__IllegalFunction                    = QModbusPdu__ExceptionCode(0x01)
	QModbusPdu__IllegalDataAddress                 = QModbusPdu__ExceptionCode(0x02)
	QModbusPdu__IllegalDataValue                   = QModbusPdu__ExceptionCode(0x03)
	QModbusPdu__ServerDeviceFailure                = QModbusPdu__ExceptionCode(0x04)
	QModbusPdu__Acknowledge                        = QModbusPdu__ExceptionCode(0x05)
	QModbusPdu__ServerDeviceBusy                   = QModbusPdu__ExceptionCode(0x06)
	QModbusPdu__NegativeAcknowledge                = QModbusPdu__ExceptionCode(0x07)
	QModbusPdu__MemoryParityError                  = QModbusPdu__ExceptionCode(0x08)
	QModbusPdu__GatewayPathUnavailable             = QModbusPdu__ExceptionCode(0x0A)
	QModbusPdu__GatewayTargetDeviceFailedToRespond = QModbusPdu__ExceptionCode(0x0B)
	QModbusPdu__ExtendedException                  = QModbusPdu__ExceptionCode(0xFF)
)

//QModbusPdu::FunctionCode
type QModbusPdu__FunctionCode int64

const (
	QModbusPdu__Invalid                        = QModbusPdu__FunctionCode(0x00)
	QModbusPdu__ReadCoils                      = QModbusPdu__FunctionCode(0x01)
	QModbusPdu__ReadDiscreteInputs             = QModbusPdu__FunctionCode(0x02)
	QModbusPdu__ReadHoldingRegisters           = QModbusPdu__FunctionCode(0x03)
	QModbusPdu__ReadInputRegisters             = QModbusPdu__FunctionCode(0x04)
	QModbusPdu__WriteSingleCoil                = QModbusPdu__FunctionCode(0x05)
	QModbusPdu__WriteSingleRegister            = QModbusPdu__FunctionCode(0x06)
	QModbusPdu__ReadExceptionStatus            = QModbusPdu__FunctionCode(0x07)
	QModbusPdu__Diagnostics                    = QModbusPdu__FunctionCode(0x08)
	QModbusPdu__GetCommEventCounter            = QModbusPdu__FunctionCode(0x0B)
	QModbusPdu__GetCommEventLog                = QModbusPdu__FunctionCode(0x0C)
	QModbusPdu__WriteMultipleCoils             = QModbusPdu__FunctionCode(0x0F)
	QModbusPdu__WriteMultipleRegisters         = QModbusPdu__FunctionCode(0x10)
	QModbusPdu__ReportServerId                 = QModbusPdu__FunctionCode(0x11)
	QModbusPdu__ReadFileRecord                 = QModbusPdu__FunctionCode(0x14)
	QModbusPdu__WriteFileRecord                = QModbusPdu__FunctionCode(0x15)
	QModbusPdu__MaskWriteRegister              = QModbusPdu__FunctionCode(0x16)
	QModbusPdu__ReadWriteMultipleRegisters     = QModbusPdu__FunctionCode(0x17)
	QModbusPdu__ReadFifoQueue                  = QModbusPdu__FunctionCode(0x18)
	QModbusPdu__EncapsulatedInterfaceTransport = QModbusPdu__FunctionCode(0x2B)
	QModbusPdu__UndefinedFunctionCode          = QModbusPdu__FunctionCode(0x100)
)

type QModbusPdu struct {
	ptr unsafe.Pointer
}

type QModbusPdu_ITF interface {
	QModbusPdu_PTR() *QModbusPdu
}

func (p *QModbusPdu) QModbusPdu_PTR() *QModbusPdu {
	return p
}

func (p *QModbusPdu) Pointer() unsafe.Pointer {
	if p != nil {
		return p.ptr
	}
	return nil
}

func (p *QModbusPdu) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.ptr = ptr
	}
}

func PointerFromQModbusPdu(ptr QModbusPdu_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QModbusPdu_PTR().Pointer()
	}
	return nil
}

func NewQModbusPduFromPointer(ptr unsafe.Pointer) *QModbusPdu {
	var n = new(QModbusPdu)
	n.SetPointer(ptr)
	return n
}

func newQModbusPduFromPointer(ptr unsafe.Pointer) *QModbusPdu {
	var n = NewQModbusPduFromPointer(ptr)
	for len(n.ObjectNameAbs()) < len("QModbusPdu_") {
		n.SetObjectNameAbs("QModbusPdu_" + qt.Identifier())
	}
	return n
}

func NewQModbusPdu() *QModbusPdu {
	defer qt.Recovering("QModbusPdu::QModbusPdu")

	return newQModbusPduFromPointer(C.QModbusPdu_NewQModbusPdu())
}

func NewQModbusPdu2(code QModbusPdu__FunctionCode, data string) *QModbusPdu {
	defer qt.Recovering("QModbusPdu::QModbusPdu")

	var dataC = C.CString(data)
	defer C.free(unsafe.Pointer(dataC))
	return newQModbusPduFromPointer(C.QModbusPdu_NewQModbusPdu2(C.int(code), dataC))
}

func NewQModbusPdu3(other QModbusPdu_ITF) *QModbusPdu {
	defer qt.Recovering("QModbusPdu::QModbusPdu")

	return newQModbusPduFromPointer(C.QModbusPdu_NewQModbusPdu3(PointerFromQModbusPdu(other)))
}

func (ptr *QModbusPdu) Data() string {
	defer qt.Recovering("QModbusPdu::data")

	if ptr.Pointer() != nil {
		return C.GoString(C.QModbusPdu_Data(ptr.Pointer()))
	}
	return ""
}

func (ptr *QModbusPdu) ExceptionCode() QModbusPdu__ExceptionCode {
	defer qt.Recovering("QModbusPdu::exceptionCode")

	if ptr.Pointer() != nil {
		return QModbusPdu__ExceptionCode(C.QModbusPdu_ExceptionCode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QModbusPdu) FunctionCode() QModbusPdu__FunctionCode {
	defer qt.Recovering("QModbusPdu::functionCode")

	if ptr.Pointer() != nil {
		return QModbusPdu__FunctionCode(C.QModbusPdu_FunctionCode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QModbusPdu) IsException() bool {
	defer qt.Recovering("QModbusPdu::isException")

	if ptr.Pointer() != nil {
		return C.QModbusPdu_IsException(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QModbusPdu) IsValid() bool {
	defer qt.Recovering("QModbusPdu::isValid")

	if ptr.Pointer() != nil {
		return C.QModbusPdu_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QModbusPdu) SetData(data string) {
	defer qt.Recovering("QModbusPdu::setData")

	if ptr.Pointer() != nil {
		var dataC = C.CString(data)
		defer C.free(unsafe.Pointer(dataC))
		C.QModbusPdu_SetData(ptr.Pointer(), dataC)
	}
}

//export callbackQModbusPdu_SetFunctionCode
func callbackQModbusPdu_SetFunctionCode(ptr unsafe.Pointer, ptrName *C.char, code C.int) {
	defer qt.Recovering("callback QModbusPdu::setFunctionCode")

	if signal := qt.GetSignal(C.GoString(ptrName), "setFunctionCode"); signal != nil {
		signal.(func(QModbusPdu__FunctionCode))(QModbusPdu__FunctionCode(code))
	} else {
		NewQModbusPduFromPointer(ptr).SetFunctionCodeDefault(QModbusPdu__FunctionCode(code))
	}
}

func (ptr *QModbusPdu) ConnectSetFunctionCode(f func(code QModbusPdu__FunctionCode)) {
	defer qt.Recovering("connect QModbusPdu::setFunctionCode")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setFunctionCode", f)
	}
}

func (ptr *QModbusPdu) DisconnectSetFunctionCode() {
	defer qt.Recovering("disconnect QModbusPdu::setFunctionCode")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setFunctionCode")
	}
}

func (ptr *QModbusPdu) SetFunctionCode(code QModbusPdu__FunctionCode) {
	defer qt.Recovering("QModbusPdu::setFunctionCode")

	if ptr.Pointer() != nil {
		C.QModbusPdu_SetFunctionCode(ptr.Pointer(), C.int(code))
	}
}

func (ptr *QModbusPdu) SetFunctionCodeDefault(code QModbusPdu__FunctionCode) {
	defer qt.Recovering("QModbusPdu::setFunctionCode")

	if ptr.Pointer() != nil {
		C.QModbusPdu_SetFunctionCodeDefault(ptr.Pointer(), C.int(code))
	}
}

func (ptr *QModbusPdu) DestroyQModbusPdu() {
	defer qt.Recovering("QModbusPdu::~QModbusPdu")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectNameAbs())
		C.QModbusPdu_DestroyQModbusPdu(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QModbusPdu) ObjectNameAbs() string {
	defer qt.Recovering("QModbusPdu::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QModbusPdu_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QModbusPdu) SetObjectNameAbs(name string) {
	defer qt.Recovering("QModbusPdu::setObjectNameAbs")

	if ptr.Pointer() != nil {
		var nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
		C.QModbusPdu_SetObjectNameAbs(ptr.Pointer(), nameC)
	}
}

//QModbusReply::ReplyType
type QModbusReply__ReplyType int64

const (
	QModbusReply__Raw    = QModbusReply__ReplyType(0)
	QModbusReply__Common = QModbusReply__ReplyType(1)
)

type QModbusReply struct {
	core.QObject
}

type QModbusReply_ITF interface {
	core.QObject_ITF
	QModbusReply_PTR() *QModbusReply
}

func (p *QModbusReply) QModbusReply_PTR() *QModbusReply {
	return p
}

func (p *QModbusReply) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QObject_PTR().Pointer()
	}
	return nil
}

func (p *QModbusReply) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QObject_PTR().SetPointer(ptr)
	}
}

func PointerFromQModbusReply(ptr QModbusReply_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QModbusReply_PTR().Pointer()
	}
	return nil
}

func NewQModbusReplyFromPointer(ptr unsafe.Pointer) *QModbusReply {
	var n = new(QModbusReply)
	n.SetPointer(ptr)
	return n
}

func newQModbusReplyFromPointer(ptr unsafe.Pointer) *QModbusReply {
	var n = NewQModbusReplyFromPointer(ptr)
	for len(n.ObjectName()) < len("QModbusReply_") {
		n.SetObjectName("QModbusReply_" + qt.Identifier())
	}
	return n
}

func (ptr *QModbusReply) DestroyQModbusReply() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

func NewQModbusReply(ty QModbusReply__ReplyType, serverAddress int, parent core.QObject_ITF) *QModbusReply {
	defer qt.Recovering("QModbusReply::QModbusReply")

	return newQModbusReplyFromPointer(C.QModbusReply_NewQModbusReply(C.int(ty), C.int(serverAddress), core.PointerFromQObject(parent)))
}

func (ptr *QModbusReply) Error() QModbusDevice__Error {
	defer qt.Recovering("QModbusReply::error")

	if ptr.Pointer() != nil {
		return QModbusDevice__Error(C.QModbusReply_Error(ptr.Pointer()))
	}
	return 0
}

//export callbackQModbusReply_ErrorOccurred
func callbackQModbusReply_ErrorOccurred(ptr unsafe.Pointer, ptrName *C.char, error C.int) {
	defer qt.Recovering("callback QModbusReply::errorOccurred")

	if signal := qt.GetSignal(C.GoString(ptrName), "errorOccurred"); signal != nil {
		signal.(func(QModbusDevice__Error))(QModbusDevice__Error(error))
	}

}

func (ptr *QModbusReply) ConnectErrorOccurred(f func(error QModbusDevice__Error)) {
	defer qt.Recovering("connect QModbusReply::errorOccurred")

	if ptr.Pointer() != nil {
		C.QModbusReply_ConnectErrorOccurred(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "errorOccurred", f)
	}
}

func (ptr *QModbusReply) DisconnectErrorOccurred() {
	defer qt.Recovering("disconnect QModbusReply::errorOccurred")

	if ptr.Pointer() != nil {
		C.QModbusReply_DisconnectErrorOccurred(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "errorOccurred")
	}
}

func (ptr *QModbusReply) ErrorOccurred(error QModbusDevice__Error) {
	defer qt.Recovering("QModbusReply::errorOccurred")

	if ptr.Pointer() != nil {
		C.QModbusReply_ErrorOccurred(ptr.Pointer(), C.int(error))
	}
}

func (ptr *QModbusReply) ErrorString() string {
	defer qt.Recovering("QModbusReply::errorString")

	if ptr.Pointer() != nil {
		return C.GoString(C.QModbusReply_ErrorString(ptr.Pointer()))
	}
	return ""
}

//export callbackQModbusReply_Finished
func callbackQModbusReply_Finished(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QModbusReply::finished")

	if signal := qt.GetSignal(C.GoString(ptrName), "finished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QModbusReply) ConnectFinished(f func()) {
	defer qt.Recovering("connect QModbusReply::finished")

	if ptr.Pointer() != nil {
		C.QModbusReply_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "finished", f)
	}
}

func (ptr *QModbusReply) DisconnectFinished() {
	defer qt.Recovering("disconnect QModbusReply::finished")

	if ptr.Pointer() != nil {
		C.QModbusReply_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "finished")
	}
}

func (ptr *QModbusReply) Finished() {
	defer qt.Recovering("QModbusReply::finished")

	if ptr.Pointer() != nil {
		C.QModbusReply_Finished(ptr.Pointer())
	}
}

func (ptr *QModbusReply) IsFinished() bool {
	defer qt.Recovering("QModbusReply::isFinished")

	if ptr.Pointer() != nil {
		return C.QModbusReply_IsFinished(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QModbusReply) ServerAddress() int {
	defer qt.Recovering("QModbusReply::serverAddress")

	if ptr.Pointer() != nil {
		return int(C.QModbusReply_ServerAddress(ptr.Pointer()))
	}
	return 0
}

func (ptr *QModbusReply) Type() QModbusReply__ReplyType {
	defer qt.Recovering("QModbusReply::type")

	if ptr.Pointer() != nil {
		return QModbusReply__ReplyType(C.QModbusReply_Type(ptr.Pointer()))
	}
	return 0
}

//export callbackQModbusReply_TimerEvent
func callbackQModbusReply_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QModbusReply::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQModbusReplyFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QModbusReply) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QModbusReply::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QModbusReply) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QModbusReply::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QModbusReply) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QModbusReply::timerEvent")

	if ptr.Pointer() != nil {
		C.QModbusReply_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QModbusReply) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QModbusReply::timerEvent")

	if ptr.Pointer() != nil {
		C.QModbusReply_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQModbusReply_ChildEvent
func callbackQModbusReply_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QModbusReply::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQModbusReplyFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QModbusReply) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QModbusReply::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QModbusReply) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QModbusReply::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QModbusReply) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QModbusReply::childEvent")

	if ptr.Pointer() != nil {
		C.QModbusReply_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QModbusReply) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QModbusReply::childEvent")

	if ptr.Pointer() != nil {
		C.QModbusReply_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQModbusReply_ConnectNotify
func callbackQModbusReply_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QModbusReply::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQModbusReplyFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QModbusReply) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QModbusReply::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QModbusReply) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QModbusReply::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QModbusReply) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QModbusReply::connectNotify")

	if ptr.Pointer() != nil {
		C.QModbusReply_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QModbusReply) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QModbusReply::connectNotify")

	if ptr.Pointer() != nil {
		C.QModbusReply_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQModbusReply_CustomEvent
func callbackQModbusReply_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QModbusReply::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQModbusReplyFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QModbusReply) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QModbusReply::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QModbusReply) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QModbusReply::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QModbusReply) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QModbusReply::customEvent")

	if ptr.Pointer() != nil {
		C.QModbusReply_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QModbusReply) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QModbusReply::customEvent")

	if ptr.Pointer() != nil {
		C.QModbusReply_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQModbusReply_DeleteLater
func callbackQModbusReply_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QModbusReply::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQModbusReplyFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QModbusReply) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QModbusReply::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QModbusReply) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QModbusReply::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QModbusReply) DeleteLater() {
	defer qt.Recovering("QModbusReply::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QModbusReply_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QModbusReply) DeleteLaterDefault() {
	defer qt.Recovering("QModbusReply::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QModbusReply_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQModbusReply_DisconnectNotify
func callbackQModbusReply_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QModbusReply::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQModbusReplyFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QModbusReply) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QModbusReply::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QModbusReply) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QModbusReply::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QModbusReply) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QModbusReply::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QModbusReply_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QModbusReply) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QModbusReply::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QModbusReply_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQModbusReply_Event
func callbackQModbusReply_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QModbusReply::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQModbusReplyFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QModbusReply) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QModbusReply::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QModbusReply) DisconnectEvent() {
	defer qt.Recovering("disconnect QModbusReply::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QModbusReply) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QModbusReply::event")

	if ptr.Pointer() != nil {
		return C.QModbusReply_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QModbusReply) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QModbusReply::event")

	if ptr.Pointer() != nil {
		return C.QModbusReply_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQModbusReply_EventFilter
func callbackQModbusReply_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QModbusReply::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQModbusReplyFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QModbusReply) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QModbusReply::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QModbusReply) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QModbusReply::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QModbusReply) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QModbusReply::eventFilter")

	if ptr.Pointer() != nil {
		return C.QModbusReply_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QModbusReply) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QModbusReply::eventFilter")

	if ptr.Pointer() != nil {
		return C.QModbusReply_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQModbusReply_MetaObject
func callbackQModbusReply_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QModbusReply::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQModbusReplyFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QModbusReply) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QModbusReply::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QModbusReply) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QModbusReply::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QModbusReply) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QModbusReply::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QModbusReply_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QModbusReply) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QModbusReply::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QModbusReply_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QModbusRequest struct {
	QModbusPdu
}

type QModbusRequest_ITF interface {
	QModbusPdu_ITF
	QModbusRequest_PTR() *QModbusRequest
}

func (p *QModbusRequest) QModbusRequest_PTR() *QModbusRequest {
	return p
}

func (p *QModbusRequest) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QModbusPdu_PTR().Pointer()
	}
	return nil
}

func (p *QModbusRequest) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QModbusPdu_PTR().SetPointer(ptr)
	}
}

func PointerFromQModbusRequest(ptr QModbusRequest_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QModbusRequest_PTR().Pointer()
	}
	return nil
}

func NewQModbusRequestFromPointer(ptr unsafe.Pointer) *QModbusRequest {
	var n = new(QModbusRequest)
	n.SetPointer(ptr)
	return n
}

func newQModbusRequestFromPointer(ptr unsafe.Pointer) *QModbusRequest {
	var n = NewQModbusRequestFromPointer(ptr)
	return n
}

func (ptr *QModbusRequest) DestroyQModbusRequest() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

func NewQModbusRequest() *QModbusRequest {
	defer qt.Recovering("QModbusRequest::QModbusRequest")

	return newQModbusRequestFromPointer(C.QModbusRequest_NewQModbusRequest())
}

func NewQModbusRequest3(code QModbusPdu__FunctionCode, data string) *QModbusRequest {
	defer qt.Recovering("QModbusRequest::QModbusRequest")

	var dataC = C.CString(data)
	defer C.free(unsafe.Pointer(dataC))
	return newQModbusRequestFromPointer(C.QModbusRequest_NewQModbusRequest3(C.int(code), dataC))
}

func NewQModbusRequest2(pdu QModbusPdu_ITF) *QModbusRequest {
	defer qt.Recovering("QModbusRequest::QModbusRequest")

	return newQModbusRequestFromPointer(C.QModbusRequest_NewQModbusRequest2(PointerFromQModbusPdu(pdu)))
}

func QModbusRequest_CalculateDataSize(request QModbusRequest_ITF) int {
	defer qt.Recovering("QModbusRequest::calculateDataSize")

	return int(C.QModbusRequest_QModbusRequest_CalculateDataSize(PointerFromQModbusRequest(request)))
}

func (ptr *QModbusRequest) CalculateDataSize(request QModbusRequest_ITF) int {
	defer qt.Recovering("QModbusRequest::calculateDataSize")

	return int(C.QModbusRequest_QModbusRequest_CalculateDataSize(PointerFromQModbusRequest(request)))
}

func QModbusRequest_MinimumDataSize(request QModbusRequest_ITF) int {
	defer qt.Recovering("QModbusRequest::minimumDataSize")

	return int(C.QModbusRequest_QModbusRequest_MinimumDataSize(PointerFromQModbusRequest(request)))
}

func (ptr *QModbusRequest) MinimumDataSize(request QModbusRequest_ITF) int {
	defer qt.Recovering("QModbusRequest::minimumDataSize")

	return int(C.QModbusRequest_QModbusRequest_MinimumDataSize(PointerFromQModbusRequest(request)))
}

//export callbackQModbusRequest_SetFunctionCode
func callbackQModbusRequest_SetFunctionCode(ptr unsafe.Pointer, ptrName *C.char, code C.int) {
	defer qt.Recovering("callback QModbusRequest::setFunctionCode")

	if signal := qt.GetSignal(C.GoString(ptrName), "setFunctionCode"); signal != nil {
		signal.(func(QModbusPdu__FunctionCode))(QModbusPdu__FunctionCode(code))
	} else {
		NewQModbusRequestFromPointer(ptr).SetFunctionCodeDefault(QModbusPdu__FunctionCode(code))
	}
}

func (ptr *QModbusRequest) ConnectSetFunctionCode(f func(code QModbusPdu__FunctionCode)) {
	defer qt.Recovering("connect QModbusRequest::setFunctionCode")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setFunctionCode", f)
	}
}

func (ptr *QModbusRequest) DisconnectSetFunctionCode() {
	defer qt.Recovering("disconnect QModbusRequest::setFunctionCode")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setFunctionCode")
	}
}

func (ptr *QModbusRequest) SetFunctionCode(code QModbusPdu__FunctionCode) {
	defer qt.Recovering("QModbusRequest::setFunctionCode")

	if ptr.Pointer() != nil {
		C.QModbusRequest_SetFunctionCode(ptr.Pointer(), C.int(code))
	}
}

func (ptr *QModbusRequest) SetFunctionCodeDefault(code QModbusPdu__FunctionCode) {
	defer qt.Recovering("QModbusRequest::setFunctionCode")

	if ptr.Pointer() != nil {
		C.QModbusRequest_SetFunctionCodeDefault(ptr.Pointer(), C.int(code))
	}
}

type QModbusResponse struct {
	QModbusPdu
}

type QModbusResponse_ITF interface {
	QModbusPdu_ITF
	QModbusResponse_PTR() *QModbusResponse
}

func (p *QModbusResponse) QModbusResponse_PTR() *QModbusResponse {
	return p
}

func (p *QModbusResponse) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QModbusPdu_PTR().Pointer()
	}
	return nil
}

func (p *QModbusResponse) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QModbusPdu_PTR().SetPointer(ptr)
	}
}

func PointerFromQModbusResponse(ptr QModbusResponse_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QModbusResponse_PTR().Pointer()
	}
	return nil
}

func NewQModbusResponseFromPointer(ptr unsafe.Pointer) *QModbusResponse {
	var n = new(QModbusResponse)
	n.SetPointer(ptr)
	return n
}

func newQModbusResponseFromPointer(ptr unsafe.Pointer) *QModbusResponse {
	var n = NewQModbusResponseFromPointer(ptr)
	return n
}

func (ptr *QModbusResponse) DestroyQModbusResponse() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

func NewQModbusResponse() *QModbusResponse {
	defer qt.Recovering("QModbusResponse::QModbusResponse")

	return newQModbusResponseFromPointer(C.QModbusResponse_NewQModbusResponse())
}

func NewQModbusResponse3(code QModbusPdu__FunctionCode, data string) *QModbusResponse {
	defer qt.Recovering("QModbusResponse::QModbusResponse")

	var dataC = C.CString(data)
	defer C.free(unsafe.Pointer(dataC))
	return newQModbusResponseFromPointer(C.QModbusResponse_NewQModbusResponse3(C.int(code), dataC))
}

func NewQModbusResponse2(pdu QModbusPdu_ITF) *QModbusResponse {
	defer qt.Recovering("QModbusResponse::QModbusResponse")

	return newQModbusResponseFromPointer(C.QModbusResponse_NewQModbusResponse2(PointerFromQModbusPdu(pdu)))
}

func QModbusResponse_CalculateDataSize(response QModbusResponse_ITF) int {
	defer qt.Recovering("QModbusResponse::calculateDataSize")

	return int(C.QModbusResponse_QModbusResponse_CalculateDataSize(PointerFromQModbusResponse(response)))
}

func (ptr *QModbusResponse) CalculateDataSize(response QModbusResponse_ITF) int {
	defer qt.Recovering("QModbusResponse::calculateDataSize")

	return int(C.QModbusResponse_QModbusResponse_CalculateDataSize(PointerFromQModbusResponse(response)))
}

func QModbusResponse_MinimumDataSize(response QModbusResponse_ITF) int {
	defer qt.Recovering("QModbusResponse::minimumDataSize")

	return int(C.QModbusResponse_QModbusResponse_MinimumDataSize(PointerFromQModbusResponse(response)))
}

func (ptr *QModbusResponse) MinimumDataSize(response QModbusResponse_ITF) int {
	defer qt.Recovering("QModbusResponse::minimumDataSize")

	return int(C.QModbusResponse_QModbusResponse_MinimumDataSize(PointerFromQModbusResponse(response)))
}

//export callbackQModbusResponse_SetFunctionCode
func callbackQModbusResponse_SetFunctionCode(ptr unsafe.Pointer, ptrName *C.char, code C.int) {
	defer qt.Recovering("callback QModbusResponse::setFunctionCode")

	if signal := qt.GetSignal(C.GoString(ptrName), "setFunctionCode"); signal != nil {
		signal.(func(QModbusPdu__FunctionCode))(QModbusPdu__FunctionCode(code))
	} else {
		NewQModbusResponseFromPointer(ptr).SetFunctionCodeDefault(QModbusPdu__FunctionCode(code))
	}
}

func (ptr *QModbusResponse) ConnectSetFunctionCode(f func(code QModbusPdu__FunctionCode)) {
	defer qt.Recovering("connect QModbusResponse::setFunctionCode")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectNameAbs(), "setFunctionCode", f)
	}
}

func (ptr *QModbusResponse) DisconnectSetFunctionCode() {
	defer qt.Recovering("disconnect QModbusResponse::setFunctionCode")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectNameAbs(), "setFunctionCode")
	}
}

func (ptr *QModbusResponse) SetFunctionCode(code QModbusPdu__FunctionCode) {
	defer qt.Recovering("QModbusResponse::setFunctionCode")

	if ptr.Pointer() != nil {
		C.QModbusResponse_SetFunctionCode(ptr.Pointer(), C.int(code))
	}
}

func (ptr *QModbusResponse) SetFunctionCodeDefault(code QModbusPdu__FunctionCode) {
	defer qt.Recovering("QModbusResponse::setFunctionCode")

	if ptr.Pointer() != nil {
		C.QModbusResponse_SetFunctionCodeDefault(ptr.Pointer(), C.int(code))
	}
}

type QModbusRtuSerialMaster struct {
	QModbusClient
}

type QModbusRtuSerialMaster_ITF interface {
	QModbusClient_ITF
	QModbusRtuSerialMaster_PTR() *QModbusRtuSerialMaster
}

func (p *QModbusRtuSerialMaster) QModbusRtuSerialMaster_PTR() *QModbusRtuSerialMaster {
	return p
}

func (p *QModbusRtuSerialMaster) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QModbusClient_PTR().Pointer()
	}
	return nil
}

func (p *QModbusRtuSerialMaster) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QModbusClient_PTR().SetPointer(ptr)
	}
}

func PointerFromQModbusRtuSerialMaster(ptr QModbusRtuSerialMaster_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QModbusRtuSerialMaster_PTR().Pointer()
	}
	return nil
}

func NewQModbusRtuSerialMasterFromPointer(ptr unsafe.Pointer) *QModbusRtuSerialMaster {
	var n = new(QModbusRtuSerialMaster)
	n.SetPointer(ptr)
	return n
}

func newQModbusRtuSerialMasterFromPointer(ptr unsafe.Pointer) *QModbusRtuSerialMaster {
	var n = NewQModbusRtuSerialMasterFromPointer(ptr)
	for len(n.ObjectName()) < len("QModbusRtuSerialMaster_") {
		n.SetObjectName("QModbusRtuSerialMaster_" + qt.Identifier())
	}
	return n
}

func (ptr *QModbusRtuSerialMaster) DestroyQModbusRtuSerialMaster() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

func NewQModbusRtuSerialMaster(parent core.QObject_ITF) *QModbusRtuSerialMaster {
	defer qt.Recovering("QModbusRtuSerialMaster::QModbusRtuSerialMaster")

	return newQModbusRtuSerialMasterFromPointer(C.QModbusRtuSerialMaster_NewQModbusRtuSerialMaster(core.PointerFromQObject(parent)))
}

func (ptr *QModbusRtuSerialMaster) InterFrameDelay() int {
	defer qt.Recovering("QModbusRtuSerialMaster::interFrameDelay")

	if ptr.Pointer() != nil {
		return int(C.QModbusRtuSerialMaster_InterFrameDelay(ptr.Pointer()))
	}
	return 0
}

func (ptr *QModbusRtuSerialMaster) SetInterFrameDelay(microseconds int) {
	defer qt.Recovering("QModbusRtuSerialMaster::setInterFrameDelay")

	if ptr.Pointer() != nil {
		C.QModbusRtuSerialMaster_SetInterFrameDelay(ptr.Pointer(), C.int(microseconds))
	}
}

//export callbackQModbusRtuSerialMaster_ProcessPrivateResponse
func callbackQModbusRtuSerialMaster_ProcessPrivateResponse(ptr unsafe.Pointer, ptrName *C.char, response unsafe.Pointer, data unsafe.Pointer) C.int {
	defer qt.Recovering("callback QModbusRtuSerialMaster::processPrivateResponse")

	if signal := qt.GetSignal(C.GoString(ptrName), "processPrivateResponse"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*QModbusResponse, *QModbusDataUnit) bool)(NewQModbusResponseFromPointer(response), NewQModbusDataUnitFromPointer(data))))
	}

	return C.int(qt.GoBoolToInt(NewQModbusRtuSerialMasterFromPointer(ptr).ProcessPrivateResponseDefault(NewQModbusResponseFromPointer(response), NewQModbusDataUnitFromPointer(data))))
}

func (ptr *QModbusRtuSerialMaster) ConnectProcessPrivateResponse(f func(response *QModbusResponse, data *QModbusDataUnit) bool) {
	defer qt.Recovering("connect QModbusRtuSerialMaster::processPrivateResponse")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "processPrivateResponse", f)
	}
}

func (ptr *QModbusRtuSerialMaster) DisconnectProcessPrivateResponse() {
	defer qt.Recovering("disconnect QModbusRtuSerialMaster::processPrivateResponse")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "processPrivateResponse")
	}
}

func (ptr *QModbusRtuSerialMaster) ProcessPrivateResponse(response QModbusResponse_ITF, data QModbusDataUnit_ITF) bool {
	defer qt.Recovering("QModbusRtuSerialMaster::processPrivateResponse")

	if ptr.Pointer() != nil {
		return C.QModbusRtuSerialMaster_ProcessPrivateResponse(ptr.Pointer(), PointerFromQModbusResponse(response), PointerFromQModbusDataUnit(data)) != 0
	}
	return false
}

func (ptr *QModbusRtuSerialMaster) ProcessPrivateResponseDefault(response QModbusResponse_ITF, data QModbusDataUnit_ITF) bool {
	defer qt.Recovering("QModbusRtuSerialMaster::processPrivateResponse")

	if ptr.Pointer() != nil {
		return C.QModbusRtuSerialMaster_ProcessPrivateResponseDefault(ptr.Pointer(), PointerFromQModbusResponse(response), PointerFromQModbusDataUnit(data)) != 0
	}
	return false
}

//export callbackQModbusRtuSerialMaster_ProcessResponse
func callbackQModbusRtuSerialMaster_ProcessResponse(ptr unsafe.Pointer, ptrName *C.char, response unsafe.Pointer, data unsafe.Pointer) C.int {
	defer qt.Recovering("callback QModbusRtuSerialMaster::processResponse")

	if signal := qt.GetSignal(C.GoString(ptrName), "processResponse"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*QModbusResponse, *QModbusDataUnit) bool)(NewQModbusResponseFromPointer(response), NewQModbusDataUnitFromPointer(data))))
	}

	return C.int(qt.GoBoolToInt(NewQModbusRtuSerialMasterFromPointer(ptr).ProcessResponseDefault(NewQModbusResponseFromPointer(response), NewQModbusDataUnitFromPointer(data))))
}

func (ptr *QModbusRtuSerialMaster) ConnectProcessResponse(f func(response *QModbusResponse, data *QModbusDataUnit) bool) {
	defer qt.Recovering("connect QModbusRtuSerialMaster::processResponse")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "processResponse", f)
	}
}

func (ptr *QModbusRtuSerialMaster) DisconnectProcessResponse() {
	defer qt.Recovering("disconnect QModbusRtuSerialMaster::processResponse")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "processResponse")
	}
}

func (ptr *QModbusRtuSerialMaster) ProcessResponse(response QModbusResponse_ITF, data QModbusDataUnit_ITF) bool {
	defer qt.Recovering("QModbusRtuSerialMaster::processResponse")

	if ptr.Pointer() != nil {
		return C.QModbusRtuSerialMaster_ProcessResponse(ptr.Pointer(), PointerFromQModbusResponse(response), PointerFromQModbusDataUnit(data)) != 0
	}
	return false
}

func (ptr *QModbusRtuSerialMaster) ProcessResponseDefault(response QModbusResponse_ITF, data QModbusDataUnit_ITF) bool {
	defer qt.Recovering("QModbusRtuSerialMaster::processResponse")

	if ptr.Pointer() != nil {
		return C.QModbusRtuSerialMaster_ProcessResponseDefault(ptr.Pointer(), PointerFromQModbusResponse(response), PointerFromQModbusDataUnit(data)) != 0
	}
	return false
}

//export callbackQModbusRtuSerialMaster_Close
func callbackQModbusRtuSerialMaster_Close(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QModbusRtuSerialMaster::close")

	if signal := qt.GetSignal(C.GoString(ptrName), "close"); signal != nil {
		signal.(func())()
	} else {
		NewQModbusRtuSerialMasterFromPointer(ptr).CloseDefault()
	}
}

func (ptr *QModbusRtuSerialMaster) ConnectClose(f func()) {
	defer qt.Recovering("connect QModbusRtuSerialMaster::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "close", f)
	}
}

func (ptr *QModbusRtuSerialMaster) DisconnectClose() {
	defer qt.Recovering("disconnect QModbusRtuSerialMaster::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "close")
	}
}

func (ptr *QModbusRtuSerialMaster) Close() {
	defer qt.Recovering("QModbusRtuSerialMaster::close")

	if ptr.Pointer() != nil {
		C.QModbusRtuSerialMaster_Close(ptr.Pointer())
	}
}

func (ptr *QModbusRtuSerialMaster) CloseDefault() {
	defer qt.Recovering("QModbusRtuSerialMaster::close")

	if ptr.Pointer() != nil {
		C.QModbusRtuSerialMaster_CloseDefault(ptr.Pointer())
	}
}

//export callbackQModbusRtuSerialMaster_Open
func callbackQModbusRtuSerialMaster_Open(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QModbusRtuSerialMaster::open")

	if signal := qt.GetSignal(C.GoString(ptrName), "open"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewQModbusRtuSerialMasterFromPointer(ptr).OpenDefault()))
}

func (ptr *QModbusRtuSerialMaster) ConnectOpen(f func() bool) {
	defer qt.Recovering("connect QModbusRtuSerialMaster::open")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "open", f)
	}
}

func (ptr *QModbusRtuSerialMaster) DisconnectOpen() {
	defer qt.Recovering("disconnect QModbusRtuSerialMaster::open")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "open")
	}
}

func (ptr *QModbusRtuSerialMaster) Open() bool {
	defer qt.Recovering("QModbusRtuSerialMaster::open")

	if ptr.Pointer() != nil {
		return C.QModbusRtuSerialMaster_Open(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QModbusRtuSerialMaster) OpenDefault() bool {
	defer qt.Recovering("QModbusRtuSerialMaster::open")

	if ptr.Pointer() != nil {
		return C.QModbusRtuSerialMaster_OpenDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQModbusRtuSerialMaster_TimerEvent
func callbackQModbusRtuSerialMaster_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QModbusRtuSerialMaster::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQModbusRtuSerialMasterFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QModbusRtuSerialMaster) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QModbusRtuSerialMaster::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QModbusRtuSerialMaster) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QModbusRtuSerialMaster::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QModbusRtuSerialMaster) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QModbusRtuSerialMaster::timerEvent")

	if ptr.Pointer() != nil {
		C.QModbusRtuSerialMaster_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QModbusRtuSerialMaster) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QModbusRtuSerialMaster::timerEvent")

	if ptr.Pointer() != nil {
		C.QModbusRtuSerialMaster_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQModbusRtuSerialMaster_ChildEvent
func callbackQModbusRtuSerialMaster_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QModbusRtuSerialMaster::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQModbusRtuSerialMasterFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QModbusRtuSerialMaster) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QModbusRtuSerialMaster::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QModbusRtuSerialMaster) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QModbusRtuSerialMaster::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QModbusRtuSerialMaster) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QModbusRtuSerialMaster::childEvent")

	if ptr.Pointer() != nil {
		C.QModbusRtuSerialMaster_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QModbusRtuSerialMaster) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QModbusRtuSerialMaster::childEvent")

	if ptr.Pointer() != nil {
		C.QModbusRtuSerialMaster_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQModbusRtuSerialMaster_ConnectNotify
func callbackQModbusRtuSerialMaster_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QModbusRtuSerialMaster::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQModbusRtuSerialMasterFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QModbusRtuSerialMaster) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QModbusRtuSerialMaster::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QModbusRtuSerialMaster) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QModbusRtuSerialMaster::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QModbusRtuSerialMaster) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QModbusRtuSerialMaster::connectNotify")

	if ptr.Pointer() != nil {
		C.QModbusRtuSerialMaster_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QModbusRtuSerialMaster) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QModbusRtuSerialMaster::connectNotify")

	if ptr.Pointer() != nil {
		C.QModbusRtuSerialMaster_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQModbusRtuSerialMaster_CustomEvent
func callbackQModbusRtuSerialMaster_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QModbusRtuSerialMaster::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQModbusRtuSerialMasterFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QModbusRtuSerialMaster) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QModbusRtuSerialMaster::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QModbusRtuSerialMaster) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QModbusRtuSerialMaster::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QModbusRtuSerialMaster) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QModbusRtuSerialMaster::customEvent")

	if ptr.Pointer() != nil {
		C.QModbusRtuSerialMaster_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QModbusRtuSerialMaster) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QModbusRtuSerialMaster::customEvent")

	if ptr.Pointer() != nil {
		C.QModbusRtuSerialMaster_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQModbusRtuSerialMaster_DeleteLater
func callbackQModbusRtuSerialMaster_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QModbusRtuSerialMaster::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQModbusRtuSerialMasterFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QModbusRtuSerialMaster) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QModbusRtuSerialMaster::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QModbusRtuSerialMaster) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QModbusRtuSerialMaster::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QModbusRtuSerialMaster) DeleteLater() {
	defer qt.Recovering("QModbusRtuSerialMaster::deleteLater")

	if ptr.Pointer() != nil {
		C.QModbusRtuSerialMaster_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QModbusRtuSerialMaster) DeleteLaterDefault() {
	defer qt.Recovering("QModbusRtuSerialMaster::deleteLater")

	if ptr.Pointer() != nil {
		C.QModbusRtuSerialMaster_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQModbusRtuSerialMaster_DisconnectNotify
func callbackQModbusRtuSerialMaster_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QModbusRtuSerialMaster::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQModbusRtuSerialMasterFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QModbusRtuSerialMaster) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QModbusRtuSerialMaster::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QModbusRtuSerialMaster) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QModbusRtuSerialMaster::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QModbusRtuSerialMaster) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QModbusRtuSerialMaster::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QModbusRtuSerialMaster_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QModbusRtuSerialMaster) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QModbusRtuSerialMaster::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QModbusRtuSerialMaster_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQModbusRtuSerialMaster_Event
func callbackQModbusRtuSerialMaster_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QModbusRtuSerialMaster::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQModbusRtuSerialMasterFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QModbusRtuSerialMaster) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QModbusRtuSerialMaster::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QModbusRtuSerialMaster) DisconnectEvent() {
	defer qt.Recovering("disconnect QModbusRtuSerialMaster::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QModbusRtuSerialMaster) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QModbusRtuSerialMaster::event")

	if ptr.Pointer() != nil {
		return C.QModbusRtuSerialMaster_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QModbusRtuSerialMaster) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QModbusRtuSerialMaster::event")

	if ptr.Pointer() != nil {
		return C.QModbusRtuSerialMaster_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQModbusRtuSerialMaster_EventFilter
func callbackQModbusRtuSerialMaster_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QModbusRtuSerialMaster::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQModbusRtuSerialMasterFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QModbusRtuSerialMaster) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QModbusRtuSerialMaster::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QModbusRtuSerialMaster) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QModbusRtuSerialMaster::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QModbusRtuSerialMaster) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QModbusRtuSerialMaster::eventFilter")

	if ptr.Pointer() != nil {
		return C.QModbusRtuSerialMaster_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QModbusRtuSerialMaster) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QModbusRtuSerialMaster::eventFilter")

	if ptr.Pointer() != nil {
		return C.QModbusRtuSerialMaster_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQModbusRtuSerialMaster_MetaObject
func callbackQModbusRtuSerialMaster_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QModbusRtuSerialMaster::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQModbusRtuSerialMasterFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QModbusRtuSerialMaster) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QModbusRtuSerialMaster::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QModbusRtuSerialMaster) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QModbusRtuSerialMaster::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QModbusRtuSerialMaster) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QModbusRtuSerialMaster::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QModbusRtuSerialMaster_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QModbusRtuSerialMaster) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QModbusRtuSerialMaster::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QModbusRtuSerialMaster_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QModbusRtuSerialSlave struct {
	QModbusServer
}

type QModbusRtuSerialSlave_ITF interface {
	QModbusServer_ITF
	QModbusRtuSerialSlave_PTR() *QModbusRtuSerialSlave
}

func (p *QModbusRtuSerialSlave) QModbusRtuSerialSlave_PTR() *QModbusRtuSerialSlave {
	return p
}

func (p *QModbusRtuSerialSlave) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QModbusServer_PTR().Pointer()
	}
	return nil
}

func (p *QModbusRtuSerialSlave) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QModbusServer_PTR().SetPointer(ptr)
	}
}

func PointerFromQModbusRtuSerialSlave(ptr QModbusRtuSerialSlave_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QModbusRtuSerialSlave_PTR().Pointer()
	}
	return nil
}

func NewQModbusRtuSerialSlaveFromPointer(ptr unsafe.Pointer) *QModbusRtuSerialSlave {
	var n = new(QModbusRtuSerialSlave)
	n.SetPointer(ptr)
	return n
}

func newQModbusRtuSerialSlaveFromPointer(ptr unsafe.Pointer) *QModbusRtuSerialSlave {
	var n = NewQModbusRtuSerialSlaveFromPointer(ptr)
	for len(n.ObjectName()) < len("QModbusRtuSerialSlave_") {
		n.SetObjectName("QModbusRtuSerialSlave_" + qt.Identifier())
	}
	return n
}

func NewQModbusRtuSerialSlave(parent core.QObject_ITF) *QModbusRtuSerialSlave {
	defer qt.Recovering("QModbusRtuSerialSlave::QModbusRtuSerialSlave")

	return newQModbusRtuSerialSlaveFromPointer(C.QModbusRtuSerialSlave_NewQModbusRtuSerialSlave(core.PointerFromQObject(parent)))
}

func (ptr *QModbusRtuSerialSlave) DestroyQModbusRtuSerialSlave() {
	defer qt.Recovering("QModbusRtuSerialSlave::~QModbusRtuSerialSlave")

	if ptr.Pointer() != nil {
		C.QModbusRtuSerialSlave_DestroyQModbusRtuSerialSlave(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQModbusRtuSerialSlave_ProcessPrivateRequest
func callbackQModbusRtuSerialSlave_ProcessPrivateRequest(ptr unsafe.Pointer, ptrName *C.char, request unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QModbusRtuSerialSlave::processPrivateRequest")

	if signal := qt.GetSignal(C.GoString(ptrName), "processPrivateRequest"); signal != nil {
		return PointerFromQModbusResponse(signal.(func(*QModbusPdu) *QModbusResponse)(NewQModbusPduFromPointer(request)))
	}

	return PointerFromQModbusResponse(NewQModbusRtuSerialSlaveFromPointer(ptr).ProcessPrivateRequestDefault(NewQModbusPduFromPointer(request)))
}

func (ptr *QModbusRtuSerialSlave) ConnectProcessPrivateRequest(f func(request *QModbusPdu) *QModbusResponse) {
	defer qt.Recovering("connect QModbusRtuSerialSlave::processPrivateRequest")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "processPrivateRequest", f)
	}
}

func (ptr *QModbusRtuSerialSlave) DisconnectProcessPrivateRequest() {
	defer qt.Recovering("disconnect QModbusRtuSerialSlave::processPrivateRequest")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "processPrivateRequest")
	}
}

func (ptr *QModbusRtuSerialSlave) ProcessPrivateRequest(request QModbusPdu_ITF) *QModbusResponse {
	defer qt.Recovering("QModbusRtuSerialSlave::processPrivateRequest")

	if ptr.Pointer() != nil {

	}
	return nil
}

func (ptr *QModbusRtuSerialSlave) ProcessPrivateRequestDefault(request QModbusPdu_ITF) *QModbusResponse {
	defer qt.Recovering("QModbusRtuSerialSlave::processPrivateRequest")

	if ptr.Pointer() != nil {

	}
	return nil
}

//export callbackQModbusRtuSerialSlave_ProcessRequest
func callbackQModbusRtuSerialSlave_ProcessRequest(ptr unsafe.Pointer, ptrName *C.char, request unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QModbusRtuSerialSlave::processRequest")

	if signal := qt.GetSignal(C.GoString(ptrName), "processRequest"); signal != nil {
		return PointerFromQModbusResponse(signal.(func(*QModbusPdu) *QModbusResponse)(NewQModbusPduFromPointer(request)))
	}

	return PointerFromQModbusResponse(NewQModbusRtuSerialSlaveFromPointer(ptr).ProcessRequestDefault(NewQModbusPduFromPointer(request)))
}

func (ptr *QModbusRtuSerialSlave) ConnectProcessRequest(f func(request *QModbusPdu) *QModbusResponse) {
	defer qt.Recovering("connect QModbusRtuSerialSlave::processRequest")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "processRequest", f)
	}
}

func (ptr *QModbusRtuSerialSlave) DisconnectProcessRequest() {
	defer qt.Recovering("disconnect QModbusRtuSerialSlave::processRequest")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "processRequest")
	}
}

func (ptr *QModbusRtuSerialSlave) ProcessRequest(request QModbusPdu_ITF) *QModbusResponse {
	defer qt.Recovering("QModbusRtuSerialSlave::processRequest")

	if ptr.Pointer() != nil {

	}
	return nil
}

func (ptr *QModbusRtuSerialSlave) ProcessRequestDefault(request QModbusPdu_ITF) *QModbusResponse {
	defer qt.Recovering("QModbusRtuSerialSlave::processRequest")

	if ptr.Pointer() != nil {

	}
	return nil
}

//export callbackQModbusRtuSerialSlave_ProcessesBroadcast
func callbackQModbusRtuSerialSlave_ProcessesBroadcast(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QModbusRtuSerialSlave::processesBroadcast")

	if signal := qt.GetSignal(C.GoString(ptrName), "processesBroadcast"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewQModbusRtuSerialSlaveFromPointer(ptr).ProcessesBroadcastDefault()))
}

func (ptr *QModbusRtuSerialSlave) ConnectProcessesBroadcast(f func() bool) {
	defer qt.Recovering("connect QModbusRtuSerialSlave::processesBroadcast")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "processesBroadcast", f)
	}
}

func (ptr *QModbusRtuSerialSlave) DisconnectProcessesBroadcast() {
	defer qt.Recovering("disconnect QModbusRtuSerialSlave::processesBroadcast")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "processesBroadcast")
	}
}

func (ptr *QModbusRtuSerialSlave) ProcessesBroadcast() bool {
	defer qt.Recovering("QModbusRtuSerialSlave::processesBroadcast")

	if ptr.Pointer() != nil {
		return C.QModbusRtuSerialSlave_ProcessesBroadcast(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QModbusRtuSerialSlave) ProcessesBroadcastDefault() bool {
	defer qt.Recovering("QModbusRtuSerialSlave::processesBroadcast")

	if ptr.Pointer() != nil {
		return C.QModbusRtuSerialSlave_ProcessesBroadcastDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQModbusRtuSerialSlave_SetValue
func callbackQModbusRtuSerialSlave_SetValue(ptr unsafe.Pointer, ptrName *C.char, option C.int, newValue unsafe.Pointer) C.int {
	defer qt.Recovering("callback QModbusRtuSerialSlave::setValue")

	if signal := qt.GetSignal(C.GoString(ptrName), "setValue"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(int, *core.QVariant) bool)(int(option), core.NewQVariantFromPointer(newValue))))
	}

	return C.int(qt.GoBoolToInt(NewQModbusRtuSerialSlaveFromPointer(ptr).SetValueDefault(int(option), core.NewQVariantFromPointer(newValue))))
}

func (ptr *QModbusRtuSerialSlave) ConnectSetValue(f func(option int, newValue *core.QVariant) bool) {
	defer qt.Recovering("connect QModbusRtuSerialSlave::setValue")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setValue", f)
	}
}

func (ptr *QModbusRtuSerialSlave) DisconnectSetValue() {
	defer qt.Recovering("disconnect QModbusRtuSerialSlave::setValue")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setValue")
	}
}

func (ptr *QModbusRtuSerialSlave) SetValue(option int, newValue core.QVariant_ITF) bool {
	defer qt.Recovering("QModbusRtuSerialSlave::setValue")

	if ptr.Pointer() != nil {
		return C.QModbusRtuSerialSlave_SetValue(ptr.Pointer(), C.int(option), core.PointerFromQVariant(newValue)) != 0
	}
	return false
}

func (ptr *QModbusRtuSerialSlave) SetValueDefault(option int, newValue core.QVariant_ITF) bool {
	defer qt.Recovering("QModbusRtuSerialSlave::setValue")

	if ptr.Pointer() != nil {
		return C.QModbusRtuSerialSlave_SetValueDefault(ptr.Pointer(), C.int(option), core.PointerFromQVariant(newValue)) != 0
	}
	return false
}

//export callbackQModbusRtuSerialSlave_Value
func callbackQModbusRtuSerialSlave_Value(ptr unsafe.Pointer, ptrName *C.char, option C.int) unsafe.Pointer {
	defer qt.Recovering("callback QModbusRtuSerialSlave::value")

	if signal := qt.GetSignal(C.GoString(ptrName), "value"); signal != nil {
		return core.PointerFromQVariant(signal.(func(int) *core.QVariant)(int(option)))
	}

	return core.PointerFromQVariant(NewQModbusRtuSerialSlaveFromPointer(ptr).ValueDefault(int(option)))
}

func (ptr *QModbusRtuSerialSlave) ConnectValue(f func(option int) *core.QVariant) {
	defer qt.Recovering("connect QModbusRtuSerialSlave::value")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "value", f)
	}
}

func (ptr *QModbusRtuSerialSlave) DisconnectValue() {
	defer qt.Recovering("disconnect QModbusRtuSerialSlave::value")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "value")
	}
}

func (ptr *QModbusRtuSerialSlave) Value(option int) *core.QVariant {
	defer qt.Recovering("QModbusRtuSerialSlave::value")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QModbusRtuSerialSlave_Value(ptr.Pointer(), C.int(option)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QModbusRtuSerialSlave) ValueDefault(option int) *core.QVariant {
	defer qt.Recovering("QModbusRtuSerialSlave::value")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QModbusRtuSerialSlave_ValueDefault(ptr.Pointer(), C.int(option)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQModbusRtuSerialSlave_WriteData
func callbackQModbusRtuSerialSlave_WriteData(ptr unsafe.Pointer, ptrName *C.char, newData unsafe.Pointer) C.int {
	defer qt.Recovering("callback QModbusRtuSerialSlave::writeData")

	if signal := qt.GetSignal(C.GoString(ptrName), "writeData"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*QModbusDataUnit) bool)(NewQModbusDataUnitFromPointer(newData))))
	}

	return C.int(qt.GoBoolToInt(NewQModbusRtuSerialSlaveFromPointer(ptr).WriteDataDefault(NewQModbusDataUnitFromPointer(newData))))
}

func (ptr *QModbusRtuSerialSlave) ConnectWriteData(f func(newData *QModbusDataUnit) bool) {
	defer qt.Recovering("connect QModbusRtuSerialSlave::writeData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "writeData", f)
	}
}

func (ptr *QModbusRtuSerialSlave) DisconnectWriteData() {
	defer qt.Recovering("disconnect QModbusRtuSerialSlave::writeData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "writeData")
	}
}

func (ptr *QModbusRtuSerialSlave) WriteData(newData QModbusDataUnit_ITF) bool {
	defer qt.Recovering("QModbusRtuSerialSlave::writeData")

	if ptr.Pointer() != nil {
		return C.QModbusRtuSerialSlave_WriteData(ptr.Pointer(), PointerFromQModbusDataUnit(newData)) != 0
	}
	return false
}

func (ptr *QModbusRtuSerialSlave) WriteDataDefault(newData QModbusDataUnit_ITF) bool {
	defer qt.Recovering("QModbusRtuSerialSlave::writeData")

	if ptr.Pointer() != nil {
		return C.QModbusRtuSerialSlave_WriteDataDefault(ptr.Pointer(), PointerFromQModbusDataUnit(newData)) != 0
	}
	return false
}

//export callbackQModbusRtuSerialSlave_Close
func callbackQModbusRtuSerialSlave_Close(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QModbusRtuSerialSlave::close")

	if signal := qt.GetSignal(C.GoString(ptrName), "close"); signal != nil {
		signal.(func())()
	} else {
		NewQModbusRtuSerialSlaveFromPointer(ptr).CloseDefault()
	}
}

func (ptr *QModbusRtuSerialSlave) ConnectClose(f func()) {
	defer qt.Recovering("connect QModbusRtuSerialSlave::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "close", f)
	}
}

func (ptr *QModbusRtuSerialSlave) DisconnectClose() {
	defer qt.Recovering("disconnect QModbusRtuSerialSlave::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "close")
	}
}

func (ptr *QModbusRtuSerialSlave) Close() {
	defer qt.Recovering("QModbusRtuSerialSlave::close")

	if ptr.Pointer() != nil {
		C.QModbusRtuSerialSlave_Close(ptr.Pointer())
	}
}

func (ptr *QModbusRtuSerialSlave) CloseDefault() {
	defer qt.Recovering("QModbusRtuSerialSlave::close")

	if ptr.Pointer() != nil {
		C.QModbusRtuSerialSlave_CloseDefault(ptr.Pointer())
	}
}

//export callbackQModbusRtuSerialSlave_Open
func callbackQModbusRtuSerialSlave_Open(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QModbusRtuSerialSlave::open")

	if signal := qt.GetSignal(C.GoString(ptrName), "open"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewQModbusRtuSerialSlaveFromPointer(ptr).OpenDefault()))
}

func (ptr *QModbusRtuSerialSlave) ConnectOpen(f func() bool) {
	defer qt.Recovering("connect QModbusRtuSerialSlave::open")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "open", f)
	}
}

func (ptr *QModbusRtuSerialSlave) DisconnectOpen() {
	defer qt.Recovering("disconnect QModbusRtuSerialSlave::open")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "open")
	}
}

func (ptr *QModbusRtuSerialSlave) Open() bool {
	defer qt.Recovering("QModbusRtuSerialSlave::open")

	if ptr.Pointer() != nil {
		return C.QModbusRtuSerialSlave_Open(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QModbusRtuSerialSlave) OpenDefault() bool {
	defer qt.Recovering("QModbusRtuSerialSlave::open")

	if ptr.Pointer() != nil {
		return C.QModbusRtuSerialSlave_OpenDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQModbusRtuSerialSlave_TimerEvent
func callbackQModbusRtuSerialSlave_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QModbusRtuSerialSlave::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQModbusRtuSerialSlaveFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QModbusRtuSerialSlave) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QModbusRtuSerialSlave::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QModbusRtuSerialSlave) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QModbusRtuSerialSlave::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QModbusRtuSerialSlave) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QModbusRtuSerialSlave::timerEvent")

	if ptr.Pointer() != nil {
		C.QModbusRtuSerialSlave_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QModbusRtuSerialSlave) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QModbusRtuSerialSlave::timerEvent")

	if ptr.Pointer() != nil {
		C.QModbusRtuSerialSlave_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQModbusRtuSerialSlave_ChildEvent
func callbackQModbusRtuSerialSlave_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QModbusRtuSerialSlave::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQModbusRtuSerialSlaveFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QModbusRtuSerialSlave) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QModbusRtuSerialSlave::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QModbusRtuSerialSlave) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QModbusRtuSerialSlave::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QModbusRtuSerialSlave) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QModbusRtuSerialSlave::childEvent")

	if ptr.Pointer() != nil {
		C.QModbusRtuSerialSlave_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QModbusRtuSerialSlave) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QModbusRtuSerialSlave::childEvent")

	if ptr.Pointer() != nil {
		C.QModbusRtuSerialSlave_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQModbusRtuSerialSlave_ConnectNotify
func callbackQModbusRtuSerialSlave_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QModbusRtuSerialSlave::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQModbusRtuSerialSlaveFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QModbusRtuSerialSlave) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QModbusRtuSerialSlave::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QModbusRtuSerialSlave) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QModbusRtuSerialSlave::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QModbusRtuSerialSlave) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QModbusRtuSerialSlave::connectNotify")

	if ptr.Pointer() != nil {
		C.QModbusRtuSerialSlave_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QModbusRtuSerialSlave) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QModbusRtuSerialSlave::connectNotify")

	if ptr.Pointer() != nil {
		C.QModbusRtuSerialSlave_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQModbusRtuSerialSlave_CustomEvent
func callbackQModbusRtuSerialSlave_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QModbusRtuSerialSlave::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQModbusRtuSerialSlaveFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QModbusRtuSerialSlave) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QModbusRtuSerialSlave::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QModbusRtuSerialSlave) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QModbusRtuSerialSlave::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QModbusRtuSerialSlave) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QModbusRtuSerialSlave::customEvent")

	if ptr.Pointer() != nil {
		C.QModbusRtuSerialSlave_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QModbusRtuSerialSlave) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QModbusRtuSerialSlave::customEvent")

	if ptr.Pointer() != nil {
		C.QModbusRtuSerialSlave_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQModbusRtuSerialSlave_DeleteLater
func callbackQModbusRtuSerialSlave_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QModbusRtuSerialSlave::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQModbusRtuSerialSlaveFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QModbusRtuSerialSlave) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QModbusRtuSerialSlave::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QModbusRtuSerialSlave) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QModbusRtuSerialSlave::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QModbusRtuSerialSlave) DeleteLater() {
	defer qt.Recovering("QModbusRtuSerialSlave::deleteLater")

	if ptr.Pointer() != nil {
		C.QModbusRtuSerialSlave_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QModbusRtuSerialSlave) DeleteLaterDefault() {
	defer qt.Recovering("QModbusRtuSerialSlave::deleteLater")

	if ptr.Pointer() != nil {
		C.QModbusRtuSerialSlave_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQModbusRtuSerialSlave_DisconnectNotify
func callbackQModbusRtuSerialSlave_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QModbusRtuSerialSlave::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQModbusRtuSerialSlaveFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QModbusRtuSerialSlave) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QModbusRtuSerialSlave::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QModbusRtuSerialSlave) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QModbusRtuSerialSlave::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QModbusRtuSerialSlave) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QModbusRtuSerialSlave::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QModbusRtuSerialSlave_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QModbusRtuSerialSlave) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QModbusRtuSerialSlave::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QModbusRtuSerialSlave_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQModbusRtuSerialSlave_Event
func callbackQModbusRtuSerialSlave_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QModbusRtuSerialSlave::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQModbusRtuSerialSlaveFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QModbusRtuSerialSlave) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QModbusRtuSerialSlave::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QModbusRtuSerialSlave) DisconnectEvent() {
	defer qt.Recovering("disconnect QModbusRtuSerialSlave::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QModbusRtuSerialSlave) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QModbusRtuSerialSlave::event")

	if ptr.Pointer() != nil {
		return C.QModbusRtuSerialSlave_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QModbusRtuSerialSlave) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QModbusRtuSerialSlave::event")

	if ptr.Pointer() != nil {
		return C.QModbusRtuSerialSlave_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQModbusRtuSerialSlave_EventFilter
func callbackQModbusRtuSerialSlave_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QModbusRtuSerialSlave::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQModbusRtuSerialSlaveFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QModbusRtuSerialSlave) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QModbusRtuSerialSlave::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QModbusRtuSerialSlave) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QModbusRtuSerialSlave::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QModbusRtuSerialSlave) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QModbusRtuSerialSlave::eventFilter")

	if ptr.Pointer() != nil {
		return C.QModbusRtuSerialSlave_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QModbusRtuSerialSlave) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QModbusRtuSerialSlave::eventFilter")

	if ptr.Pointer() != nil {
		return C.QModbusRtuSerialSlave_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQModbusRtuSerialSlave_MetaObject
func callbackQModbusRtuSerialSlave_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QModbusRtuSerialSlave::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQModbusRtuSerialSlaveFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QModbusRtuSerialSlave) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QModbusRtuSerialSlave::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QModbusRtuSerialSlave) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QModbusRtuSerialSlave::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QModbusRtuSerialSlave) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QModbusRtuSerialSlave::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QModbusRtuSerialSlave_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QModbusRtuSerialSlave) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QModbusRtuSerialSlave::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QModbusRtuSerialSlave_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//QModbusServer::Option
type QModbusServer__Option int64

const (
	QModbusServer__DiagnosticRegister    = QModbusServer__Option(0)
	QModbusServer__ExceptionStatusOffset = QModbusServer__Option(1)
	QModbusServer__DeviceBusy            = QModbusServer__Option(2)
	QModbusServer__AsciiInputDelimiter   = QModbusServer__Option(3)
	QModbusServer__ListenOnlyMode        = QModbusServer__Option(4)
	QModbusServer__ServerIdentifier      = QModbusServer__Option(5)
	QModbusServer__RunIndicatorStatus    = QModbusServer__Option(6)
	QModbusServer__AdditionalData        = QModbusServer__Option(7)
	QModbusServer__DeviceIdentification  = QModbusServer__Option(8)
	QModbusServer__UserOption            = QModbusServer__Option(0x100)
)

type QModbusServer struct {
	QModbusDevice
}

type QModbusServer_ITF interface {
	QModbusDevice_ITF
	QModbusServer_PTR() *QModbusServer
}

func (p *QModbusServer) QModbusServer_PTR() *QModbusServer {
	return p
}

func (p *QModbusServer) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QModbusDevice_PTR().Pointer()
	}
	return nil
}

func (p *QModbusServer) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QModbusDevice_PTR().SetPointer(ptr)
	}
}

func PointerFromQModbusServer(ptr QModbusServer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QModbusServer_PTR().Pointer()
	}
	return nil
}

func NewQModbusServerFromPointer(ptr unsafe.Pointer) *QModbusServer {
	var n = new(QModbusServer)
	n.SetPointer(ptr)
	return n
}

func newQModbusServerFromPointer(ptr unsafe.Pointer) *QModbusServer {
	var n = NewQModbusServerFromPointer(ptr)
	for len(n.ObjectName()) < len("QModbusServer_") {
		n.SetObjectName("QModbusServer_" + qt.Identifier())
	}
	return n
}

func (ptr *QModbusServer) DestroyQModbusServer() {
	C.free(ptr.Pointer())
	ptr.SetPointer(nil)
}

func NewQModbusServer(parent core.QObject_ITF) *QModbusServer {
	defer qt.Recovering("QModbusServer::QModbusServer")

	return newQModbusServerFromPointer(C.QModbusServer_NewQModbusServer(core.PointerFromQObject(parent)))
}

func (ptr *QModbusServer) Data(newData QModbusDataUnit_ITF) bool {
	defer qt.Recovering("QModbusServer::data")

	if ptr.Pointer() != nil {
		return C.QModbusServer_Data(ptr.Pointer(), PointerFromQModbusDataUnit(newData)) != 0
	}
	return false
}

//export callbackQModbusServer_DataWritten
func callbackQModbusServer_DataWritten(ptr unsafe.Pointer, ptrName *C.char, regist C.int, address C.int, size C.int) {
	defer qt.Recovering("callback QModbusServer::dataWritten")

	if signal := qt.GetSignal(C.GoString(ptrName), "dataWritten"); signal != nil {
		signal.(func(QModbusDataUnit__RegisterType, int, int))(QModbusDataUnit__RegisterType(regist), int(address), int(size))
	}

}

func (ptr *QModbusServer) ConnectDataWritten(f func(regist QModbusDataUnit__RegisterType, address int, size int)) {
	defer qt.Recovering("connect QModbusServer::dataWritten")

	if ptr.Pointer() != nil {
		C.QModbusServer_ConnectDataWritten(ptr.Pointer())
		qt.ConnectSignal(ptr.ObjectName(), "dataWritten", f)
	}
}

func (ptr *QModbusServer) DisconnectDataWritten() {
	defer qt.Recovering("disconnect QModbusServer::dataWritten")

	if ptr.Pointer() != nil {
		C.QModbusServer_DisconnectDataWritten(ptr.Pointer())
		qt.DisconnectSignal(ptr.ObjectName(), "dataWritten")
	}
}

func (ptr *QModbusServer) DataWritten(regist QModbusDataUnit__RegisterType, address int, size int) {
	defer qt.Recovering("QModbusServer::dataWritten")

	if ptr.Pointer() != nil {
		C.QModbusServer_DataWritten(ptr.Pointer(), C.int(regist), C.int(address), C.int(size))
	}
}

//export callbackQModbusServer_ProcessPrivateRequest
func callbackQModbusServer_ProcessPrivateRequest(ptr unsafe.Pointer, ptrName *C.char, request unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QModbusServer::processPrivateRequest")

	if signal := qt.GetSignal(C.GoString(ptrName), "processPrivateRequest"); signal != nil {
		return PointerFromQModbusResponse(signal.(func(*QModbusPdu) *QModbusResponse)(NewQModbusPduFromPointer(request)))
	}

	return PointerFromQModbusResponse(NewQModbusServerFromPointer(ptr).ProcessPrivateRequestDefault(NewQModbusPduFromPointer(request)))
}

func (ptr *QModbusServer) ConnectProcessPrivateRequest(f func(request *QModbusPdu) *QModbusResponse) {
	defer qt.Recovering("connect QModbusServer::processPrivateRequest")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "processPrivateRequest", f)
	}
}

func (ptr *QModbusServer) DisconnectProcessPrivateRequest() {
	defer qt.Recovering("disconnect QModbusServer::processPrivateRequest")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "processPrivateRequest")
	}
}

func (ptr *QModbusServer) ProcessPrivateRequest(request QModbusPdu_ITF) *QModbusResponse {
	defer qt.Recovering("QModbusServer::processPrivateRequest")

	if ptr.Pointer() != nil {

	}
	return nil
}

func (ptr *QModbusServer) ProcessPrivateRequestDefault(request QModbusPdu_ITF) *QModbusResponse {
	defer qt.Recovering("QModbusServer::processPrivateRequest")

	if ptr.Pointer() != nil {

	}
	return nil
}

//export callbackQModbusServer_ProcessRequest
func callbackQModbusServer_ProcessRequest(ptr unsafe.Pointer, ptrName *C.char, request unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QModbusServer::processRequest")

	if signal := qt.GetSignal(C.GoString(ptrName), "processRequest"); signal != nil {
		return PointerFromQModbusResponse(signal.(func(*QModbusPdu) *QModbusResponse)(NewQModbusPduFromPointer(request)))
	}

	return PointerFromQModbusResponse(NewQModbusServerFromPointer(ptr).ProcessRequestDefault(NewQModbusPduFromPointer(request)))
}

func (ptr *QModbusServer) ConnectProcessRequest(f func(request *QModbusPdu) *QModbusResponse) {
	defer qt.Recovering("connect QModbusServer::processRequest")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "processRequest", f)
	}
}

func (ptr *QModbusServer) DisconnectProcessRequest() {
	defer qt.Recovering("disconnect QModbusServer::processRequest")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "processRequest")
	}
}

func (ptr *QModbusServer) ProcessRequest(request QModbusPdu_ITF) *QModbusResponse {
	defer qt.Recovering("QModbusServer::processRequest")

	if ptr.Pointer() != nil {

	}
	return nil
}

func (ptr *QModbusServer) ProcessRequestDefault(request QModbusPdu_ITF) *QModbusResponse {
	defer qt.Recovering("QModbusServer::processRequest")

	if ptr.Pointer() != nil {

	}
	return nil
}

//export callbackQModbusServer_ProcessesBroadcast
func callbackQModbusServer_ProcessesBroadcast(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QModbusServer::processesBroadcast")

	if signal := qt.GetSignal(C.GoString(ptrName), "processesBroadcast"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewQModbusServerFromPointer(ptr).ProcessesBroadcastDefault()))
}

func (ptr *QModbusServer) ConnectProcessesBroadcast(f func() bool) {
	defer qt.Recovering("connect QModbusServer::processesBroadcast")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "processesBroadcast", f)
	}
}

func (ptr *QModbusServer) DisconnectProcessesBroadcast() {
	defer qt.Recovering("disconnect QModbusServer::processesBroadcast")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "processesBroadcast")
	}
}

func (ptr *QModbusServer) ProcessesBroadcast() bool {
	defer qt.Recovering("QModbusServer::processesBroadcast")

	if ptr.Pointer() != nil {
		return C.QModbusServer_ProcessesBroadcast(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QModbusServer) ProcessesBroadcastDefault() bool {
	defer qt.Recovering("QModbusServer::processesBroadcast")

	if ptr.Pointer() != nil {
		return C.QModbusServer_ProcessesBroadcastDefault(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QModbusServer) ServerAddress() int {
	defer qt.Recovering("QModbusServer::serverAddress")

	if ptr.Pointer() != nil {
		return int(C.QModbusServer_ServerAddress(ptr.Pointer()))
	}
	return 0
}

func (ptr *QModbusServer) SetData(newData QModbusDataUnit_ITF) bool {
	defer qt.Recovering("QModbusServer::setData")

	if ptr.Pointer() != nil {
		return C.QModbusServer_SetData(ptr.Pointer(), PointerFromQModbusDataUnit(newData)) != 0
	}
	return false
}

func (ptr *QModbusServer) SetServerAddress(serverAddress int) {
	defer qt.Recovering("QModbusServer::setServerAddress")

	if ptr.Pointer() != nil {
		C.QModbusServer_SetServerAddress(ptr.Pointer(), C.int(serverAddress))
	}
}

//export callbackQModbusServer_SetValue
func callbackQModbusServer_SetValue(ptr unsafe.Pointer, ptrName *C.char, option C.int, newValue unsafe.Pointer) C.int {
	defer qt.Recovering("callback QModbusServer::setValue")

	if signal := qt.GetSignal(C.GoString(ptrName), "setValue"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(int, *core.QVariant) bool)(int(option), core.NewQVariantFromPointer(newValue))))
	}

	return C.int(qt.GoBoolToInt(NewQModbusServerFromPointer(ptr).SetValueDefault(int(option), core.NewQVariantFromPointer(newValue))))
}

func (ptr *QModbusServer) ConnectSetValue(f func(option int, newValue *core.QVariant) bool) {
	defer qt.Recovering("connect QModbusServer::setValue")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setValue", f)
	}
}

func (ptr *QModbusServer) DisconnectSetValue() {
	defer qt.Recovering("disconnect QModbusServer::setValue")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setValue")
	}
}

func (ptr *QModbusServer) SetValue(option int, newValue core.QVariant_ITF) bool {
	defer qt.Recovering("QModbusServer::setValue")

	if ptr.Pointer() != nil {
		return C.QModbusServer_SetValue(ptr.Pointer(), C.int(option), core.PointerFromQVariant(newValue)) != 0
	}
	return false
}

func (ptr *QModbusServer) SetValueDefault(option int, newValue core.QVariant_ITF) bool {
	defer qt.Recovering("QModbusServer::setValue")

	if ptr.Pointer() != nil {
		return C.QModbusServer_SetValueDefault(ptr.Pointer(), C.int(option), core.PointerFromQVariant(newValue)) != 0
	}
	return false
}

//export callbackQModbusServer_Value
func callbackQModbusServer_Value(ptr unsafe.Pointer, ptrName *C.char, option C.int) unsafe.Pointer {
	defer qt.Recovering("callback QModbusServer::value")

	if signal := qt.GetSignal(C.GoString(ptrName), "value"); signal != nil {
		return core.PointerFromQVariant(signal.(func(int) *core.QVariant)(int(option)))
	}

	return core.PointerFromQVariant(NewQModbusServerFromPointer(ptr).ValueDefault(int(option)))
}

func (ptr *QModbusServer) ConnectValue(f func(option int) *core.QVariant) {
	defer qt.Recovering("connect QModbusServer::value")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "value", f)
	}
}

func (ptr *QModbusServer) DisconnectValue() {
	defer qt.Recovering("disconnect QModbusServer::value")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "value")
	}
}

func (ptr *QModbusServer) Value(option int) *core.QVariant {
	defer qt.Recovering("QModbusServer::value")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QModbusServer_Value(ptr.Pointer(), C.int(option)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QModbusServer) ValueDefault(option int) *core.QVariant {
	defer qt.Recovering("QModbusServer::value")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QModbusServer_ValueDefault(ptr.Pointer(), C.int(option)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQModbusServer_WriteData
func callbackQModbusServer_WriteData(ptr unsafe.Pointer, ptrName *C.char, newData unsafe.Pointer) C.int {
	defer qt.Recovering("callback QModbusServer::writeData")

	if signal := qt.GetSignal(C.GoString(ptrName), "writeData"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*QModbusDataUnit) bool)(NewQModbusDataUnitFromPointer(newData))))
	}

	return C.int(qt.GoBoolToInt(NewQModbusServerFromPointer(ptr).WriteDataDefault(NewQModbusDataUnitFromPointer(newData))))
}

func (ptr *QModbusServer) ConnectWriteData(f func(newData *QModbusDataUnit) bool) {
	defer qt.Recovering("connect QModbusServer::writeData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "writeData", f)
	}
}

func (ptr *QModbusServer) DisconnectWriteData() {
	defer qt.Recovering("disconnect QModbusServer::writeData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "writeData")
	}
}

func (ptr *QModbusServer) WriteData(newData QModbusDataUnit_ITF) bool {
	defer qt.Recovering("QModbusServer::writeData")

	if ptr.Pointer() != nil {
		return C.QModbusServer_WriteData(ptr.Pointer(), PointerFromQModbusDataUnit(newData)) != 0
	}
	return false
}

func (ptr *QModbusServer) WriteDataDefault(newData QModbusDataUnit_ITF) bool {
	defer qt.Recovering("QModbusServer::writeData")

	if ptr.Pointer() != nil {
		return C.QModbusServer_WriteDataDefault(ptr.Pointer(), PointerFromQModbusDataUnit(newData)) != 0
	}
	return false
}

//export callbackQModbusServer_Close
func callbackQModbusServer_Close(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QModbusServer::close")

	if signal := qt.GetSignal(C.GoString(ptrName), "close"); signal != nil {
		signal.(func())()
	} else {

	}
}

func (ptr *QModbusServer) ConnectClose(f func()) {
	defer qt.Recovering("connect QModbusServer::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "close", f)
	}
}

func (ptr *QModbusServer) DisconnectClose() {
	defer qt.Recovering("disconnect QModbusServer::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "close")
	}
}

func (ptr *QModbusServer) Close() {
	defer qt.Recovering("QModbusServer::close")

	if ptr.Pointer() != nil {
		C.QModbusServer_Close(ptr.Pointer())
	}
}

//export callbackQModbusServer_Open
func callbackQModbusServer_Open(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QModbusServer::open")

	if signal := qt.GetSignal(C.GoString(ptrName), "open"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(false))
}

func (ptr *QModbusServer) ConnectOpen(f func() bool) {
	defer qt.Recovering("connect QModbusServer::open")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "open", f)
	}
}

func (ptr *QModbusServer) DisconnectOpen() {
	defer qt.Recovering("disconnect QModbusServer::open")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "open")
	}
}

func (ptr *QModbusServer) Open() bool {
	defer qt.Recovering("QModbusServer::open")

	if ptr.Pointer() != nil {
		return C.QModbusServer_Open(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQModbusServer_TimerEvent
func callbackQModbusServer_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QModbusServer::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQModbusServerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QModbusServer) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QModbusServer::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QModbusServer) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QModbusServer::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QModbusServer) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QModbusServer::timerEvent")

	if ptr.Pointer() != nil {
		C.QModbusServer_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QModbusServer) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QModbusServer::timerEvent")

	if ptr.Pointer() != nil {
		C.QModbusServer_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQModbusServer_ChildEvent
func callbackQModbusServer_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QModbusServer::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQModbusServerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QModbusServer) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QModbusServer::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QModbusServer) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QModbusServer::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QModbusServer) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QModbusServer::childEvent")

	if ptr.Pointer() != nil {
		C.QModbusServer_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QModbusServer) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QModbusServer::childEvent")

	if ptr.Pointer() != nil {
		C.QModbusServer_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQModbusServer_ConnectNotify
func callbackQModbusServer_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QModbusServer::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQModbusServerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QModbusServer) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QModbusServer::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QModbusServer) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QModbusServer::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QModbusServer) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QModbusServer::connectNotify")

	if ptr.Pointer() != nil {
		C.QModbusServer_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QModbusServer) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QModbusServer::connectNotify")

	if ptr.Pointer() != nil {
		C.QModbusServer_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQModbusServer_CustomEvent
func callbackQModbusServer_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QModbusServer::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQModbusServerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QModbusServer) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QModbusServer::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QModbusServer) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QModbusServer::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QModbusServer) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QModbusServer::customEvent")

	if ptr.Pointer() != nil {
		C.QModbusServer_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QModbusServer) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QModbusServer::customEvent")

	if ptr.Pointer() != nil {
		C.QModbusServer_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQModbusServer_DeleteLater
func callbackQModbusServer_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QModbusServer::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQModbusServerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QModbusServer) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QModbusServer::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QModbusServer) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QModbusServer::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QModbusServer) DeleteLater() {
	defer qt.Recovering("QModbusServer::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QModbusServer_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QModbusServer) DeleteLaterDefault() {
	defer qt.Recovering("QModbusServer::deleteLater")

	if ptr.Pointer() != nil {
		qt.DisconnectAllSignals(ptr.ObjectName())
		C.QModbusServer_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQModbusServer_DisconnectNotify
func callbackQModbusServer_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QModbusServer::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQModbusServerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QModbusServer) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QModbusServer::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QModbusServer) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QModbusServer::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QModbusServer) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QModbusServer::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QModbusServer_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QModbusServer) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QModbusServer::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QModbusServer_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQModbusServer_Event
func callbackQModbusServer_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QModbusServer::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQModbusServerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QModbusServer) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QModbusServer::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QModbusServer) DisconnectEvent() {
	defer qt.Recovering("disconnect QModbusServer::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QModbusServer) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QModbusServer::event")

	if ptr.Pointer() != nil {
		return C.QModbusServer_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QModbusServer) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QModbusServer::event")

	if ptr.Pointer() != nil {
		return C.QModbusServer_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQModbusServer_EventFilter
func callbackQModbusServer_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QModbusServer::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQModbusServerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QModbusServer) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QModbusServer::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QModbusServer) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QModbusServer::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QModbusServer) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QModbusServer::eventFilter")

	if ptr.Pointer() != nil {
		return C.QModbusServer_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QModbusServer) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QModbusServer::eventFilter")

	if ptr.Pointer() != nil {
		return C.QModbusServer_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQModbusServer_MetaObject
func callbackQModbusServer_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QModbusServer::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQModbusServerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QModbusServer) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QModbusServer::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QModbusServer) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QModbusServer::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QModbusServer) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QModbusServer::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QModbusServer_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QModbusServer) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QModbusServer::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QModbusServer_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QModbusTcpClient struct {
	QModbusClient
}

type QModbusTcpClient_ITF interface {
	QModbusClient_ITF
	QModbusTcpClient_PTR() *QModbusTcpClient
}

func (p *QModbusTcpClient) QModbusTcpClient_PTR() *QModbusTcpClient {
	return p
}

func (p *QModbusTcpClient) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QModbusClient_PTR().Pointer()
	}
	return nil
}

func (p *QModbusTcpClient) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QModbusClient_PTR().SetPointer(ptr)
	}
}

func PointerFromQModbusTcpClient(ptr QModbusTcpClient_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QModbusTcpClient_PTR().Pointer()
	}
	return nil
}

func NewQModbusTcpClientFromPointer(ptr unsafe.Pointer) *QModbusTcpClient {
	var n = new(QModbusTcpClient)
	n.SetPointer(ptr)
	return n
}

func newQModbusTcpClientFromPointer(ptr unsafe.Pointer) *QModbusTcpClient {
	var n = NewQModbusTcpClientFromPointer(ptr)
	for len(n.ObjectName()) < len("QModbusTcpClient_") {
		n.SetObjectName("QModbusTcpClient_" + qt.Identifier())
	}
	return n
}

func NewQModbusTcpClient(parent core.QObject_ITF) *QModbusTcpClient {
	defer qt.Recovering("QModbusTcpClient::QModbusTcpClient")

	return newQModbusTcpClientFromPointer(C.QModbusTcpClient_NewQModbusTcpClient(core.PointerFromQObject(parent)))
}

func (ptr *QModbusTcpClient) DestroyQModbusTcpClient() {
	defer qt.Recovering("QModbusTcpClient::~QModbusTcpClient")

	if ptr.Pointer() != nil {
		C.QModbusTcpClient_DestroyQModbusTcpClient(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQModbusTcpClient_ProcessPrivateResponse
func callbackQModbusTcpClient_ProcessPrivateResponse(ptr unsafe.Pointer, ptrName *C.char, response unsafe.Pointer, data unsafe.Pointer) C.int {
	defer qt.Recovering("callback QModbusTcpClient::processPrivateResponse")

	if signal := qt.GetSignal(C.GoString(ptrName), "processPrivateResponse"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*QModbusResponse, *QModbusDataUnit) bool)(NewQModbusResponseFromPointer(response), NewQModbusDataUnitFromPointer(data))))
	}

	return C.int(qt.GoBoolToInt(NewQModbusTcpClientFromPointer(ptr).ProcessPrivateResponseDefault(NewQModbusResponseFromPointer(response), NewQModbusDataUnitFromPointer(data))))
}

func (ptr *QModbusTcpClient) ConnectProcessPrivateResponse(f func(response *QModbusResponse, data *QModbusDataUnit) bool) {
	defer qt.Recovering("connect QModbusTcpClient::processPrivateResponse")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "processPrivateResponse", f)
	}
}

func (ptr *QModbusTcpClient) DisconnectProcessPrivateResponse() {
	defer qt.Recovering("disconnect QModbusTcpClient::processPrivateResponse")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "processPrivateResponse")
	}
}

func (ptr *QModbusTcpClient) ProcessPrivateResponse(response QModbusResponse_ITF, data QModbusDataUnit_ITF) bool {
	defer qt.Recovering("QModbusTcpClient::processPrivateResponse")

	if ptr.Pointer() != nil {
		return C.QModbusTcpClient_ProcessPrivateResponse(ptr.Pointer(), PointerFromQModbusResponse(response), PointerFromQModbusDataUnit(data)) != 0
	}
	return false
}

func (ptr *QModbusTcpClient) ProcessPrivateResponseDefault(response QModbusResponse_ITF, data QModbusDataUnit_ITF) bool {
	defer qt.Recovering("QModbusTcpClient::processPrivateResponse")

	if ptr.Pointer() != nil {
		return C.QModbusTcpClient_ProcessPrivateResponseDefault(ptr.Pointer(), PointerFromQModbusResponse(response), PointerFromQModbusDataUnit(data)) != 0
	}
	return false
}

//export callbackQModbusTcpClient_ProcessResponse
func callbackQModbusTcpClient_ProcessResponse(ptr unsafe.Pointer, ptrName *C.char, response unsafe.Pointer, data unsafe.Pointer) C.int {
	defer qt.Recovering("callback QModbusTcpClient::processResponse")

	if signal := qt.GetSignal(C.GoString(ptrName), "processResponse"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*QModbusResponse, *QModbusDataUnit) bool)(NewQModbusResponseFromPointer(response), NewQModbusDataUnitFromPointer(data))))
	}

	return C.int(qt.GoBoolToInt(NewQModbusTcpClientFromPointer(ptr).ProcessResponseDefault(NewQModbusResponseFromPointer(response), NewQModbusDataUnitFromPointer(data))))
}

func (ptr *QModbusTcpClient) ConnectProcessResponse(f func(response *QModbusResponse, data *QModbusDataUnit) bool) {
	defer qt.Recovering("connect QModbusTcpClient::processResponse")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "processResponse", f)
	}
}

func (ptr *QModbusTcpClient) DisconnectProcessResponse() {
	defer qt.Recovering("disconnect QModbusTcpClient::processResponse")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "processResponse")
	}
}

func (ptr *QModbusTcpClient) ProcessResponse(response QModbusResponse_ITF, data QModbusDataUnit_ITF) bool {
	defer qt.Recovering("QModbusTcpClient::processResponse")

	if ptr.Pointer() != nil {
		return C.QModbusTcpClient_ProcessResponse(ptr.Pointer(), PointerFromQModbusResponse(response), PointerFromQModbusDataUnit(data)) != 0
	}
	return false
}

func (ptr *QModbusTcpClient) ProcessResponseDefault(response QModbusResponse_ITF, data QModbusDataUnit_ITF) bool {
	defer qt.Recovering("QModbusTcpClient::processResponse")

	if ptr.Pointer() != nil {
		return C.QModbusTcpClient_ProcessResponseDefault(ptr.Pointer(), PointerFromQModbusResponse(response), PointerFromQModbusDataUnit(data)) != 0
	}
	return false
}

//export callbackQModbusTcpClient_Close
func callbackQModbusTcpClient_Close(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QModbusTcpClient::close")

	if signal := qt.GetSignal(C.GoString(ptrName), "close"); signal != nil {
		signal.(func())()
	} else {
		NewQModbusTcpClientFromPointer(ptr).CloseDefault()
	}
}

func (ptr *QModbusTcpClient) ConnectClose(f func()) {
	defer qt.Recovering("connect QModbusTcpClient::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "close", f)
	}
}

func (ptr *QModbusTcpClient) DisconnectClose() {
	defer qt.Recovering("disconnect QModbusTcpClient::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "close")
	}
}

func (ptr *QModbusTcpClient) Close() {
	defer qt.Recovering("QModbusTcpClient::close")

	if ptr.Pointer() != nil {
		C.QModbusTcpClient_Close(ptr.Pointer())
	}
}

func (ptr *QModbusTcpClient) CloseDefault() {
	defer qt.Recovering("QModbusTcpClient::close")

	if ptr.Pointer() != nil {
		C.QModbusTcpClient_CloseDefault(ptr.Pointer())
	}
}

//export callbackQModbusTcpClient_Open
func callbackQModbusTcpClient_Open(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QModbusTcpClient::open")

	if signal := qt.GetSignal(C.GoString(ptrName), "open"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewQModbusTcpClientFromPointer(ptr).OpenDefault()))
}

func (ptr *QModbusTcpClient) ConnectOpen(f func() bool) {
	defer qt.Recovering("connect QModbusTcpClient::open")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "open", f)
	}
}

func (ptr *QModbusTcpClient) DisconnectOpen() {
	defer qt.Recovering("disconnect QModbusTcpClient::open")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "open")
	}
}

func (ptr *QModbusTcpClient) Open() bool {
	defer qt.Recovering("QModbusTcpClient::open")

	if ptr.Pointer() != nil {
		return C.QModbusTcpClient_Open(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QModbusTcpClient) OpenDefault() bool {
	defer qt.Recovering("QModbusTcpClient::open")

	if ptr.Pointer() != nil {
		return C.QModbusTcpClient_OpenDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQModbusTcpClient_TimerEvent
func callbackQModbusTcpClient_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QModbusTcpClient::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQModbusTcpClientFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QModbusTcpClient) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QModbusTcpClient::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QModbusTcpClient) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QModbusTcpClient::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QModbusTcpClient) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QModbusTcpClient::timerEvent")

	if ptr.Pointer() != nil {
		C.QModbusTcpClient_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QModbusTcpClient) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QModbusTcpClient::timerEvent")

	if ptr.Pointer() != nil {
		C.QModbusTcpClient_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQModbusTcpClient_ChildEvent
func callbackQModbusTcpClient_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QModbusTcpClient::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQModbusTcpClientFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QModbusTcpClient) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QModbusTcpClient::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QModbusTcpClient) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QModbusTcpClient::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QModbusTcpClient) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QModbusTcpClient::childEvent")

	if ptr.Pointer() != nil {
		C.QModbusTcpClient_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QModbusTcpClient) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QModbusTcpClient::childEvent")

	if ptr.Pointer() != nil {
		C.QModbusTcpClient_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQModbusTcpClient_ConnectNotify
func callbackQModbusTcpClient_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QModbusTcpClient::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQModbusTcpClientFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QModbusTcpClient) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QModbusTcpClient::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QModbusTcpClient) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QModbusTcpClient::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QModbusTcpClient) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QModbusTcpClient::connectNotify")

	if ptr.Pointer() != nil {
		C.QModbusTcpClient_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QModbusTcpClient) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QModbusTcpClient::connectNotify")

	if ptr.Pointer() != nil {
		C.QModbusTcpClient_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQModbusTcpClient_CustomEvent
func callbackQModbusTcpClient_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QModbusTcpClient::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQModbusTcpClientFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QModbusTcpClient) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QModbusTcpClient::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QModbusTcpClient) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QModbusTcpClient::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QModbusTcpClient) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QModbusTcpClient::customEvent")

	if ptr.Pointer() != nil {
		C.QModbusTcpClient_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QModbusTcpClient) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QModbusTcpClient::customEvent")

	if ptr.Pointer() != nil {
		C.QModbusTcpClient_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQModbusTcpClient_DeleteLater
func callbackQModbusTcpClient_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QModbusTcpClient::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQModbusTcpClientFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QModbusTcpClient) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QModbusTcpClient::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QModbusTcpClient) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QModbusTcpClient::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QModbusTcpClient) DeleteLater() {
	defer qt.Recovering("QModbusTcpClient::deleteLater")

	if ptr.Pointer() != nil {
		C.QModbusTcpClient_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QModbusTcpClient) DeleteLaterDefault() {
	defer qt.Recovering("QModbusTcpClient::deleteLater")

	if ptr.Pointer() != nil {
		C.QModbusTcpClient_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQModbusTcpClient_DisconnectNotify
func callbackQModbusTcpClient_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QModbusTcpClient::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQModbusTcpClientFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QModbusTcpClient) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QModbusTcpClient::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QModbusTcpClient) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QModbusTcpClient::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QModbusTcpClient) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QModbusTcpClient::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QModbusTcpClient_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QModbusTcpClient) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QModbusTcpClient::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QModbusTcpClient_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQModbusTcpClient_Event
func callbackQModbusTcpClient_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QModbusTcpClient::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQModbusTcpClientFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QModbusTcpClient) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QModbusTcpClient::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QModbusTcpClient) DisconnectEvent() {
	defer qt.Recovering("disconnect QModbusTcpClient::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QModbusTcpClient) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QModbusTcpClient::event")

	if ptr.Pointer() != nil {
		return C.QModbusTcpClient_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QModbusTcpClient) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QModbusTcpClient::event")

	if ptr.Pointer() != nil {
		return C.QModbusTcpClient_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQModbusTcpClient_EventFilter
func callbackQModbusTcpClient_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QModbusTcpClient::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQModbusTcpClientFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QModbusTcpClient) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QModbusTcpClient::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QModbusTcpClient) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QModbusTcpClient::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QModbusTcpClient) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QModbusTcpClient::eventFilter")

	if ptr.Pointer() != nil {
		return C.QModbusTcpClient_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QModbusTcpClient) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QModbusTcpClient::eventFilter")

	if ptr.Pointer() != nil {
		return C.QModbusTcpClient_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQModbusTcpClient_MetaObject
func callbackQModbusTcpClient_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QModbusTcpClient::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQModbusTcpClientFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QModbusTcpClient) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QModbusTcpClient::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QModbusTcpClient) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QModbusTcpClient::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QModbusTcpClient) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QModbusTcpClient::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QModbusTcpClient_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QModbusTcpClient) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QModbusTcpClient::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QModbusTcpClient_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QModbusTcpServer struct {
	QModbusServer
}

type QModbusTcpServer_ITF interface {
	QModbusServer_ITF
	QModbusTcpServer_PTR() *QModbusTcpServer
}

func (p *QModbusTcpServer) QModbusTcpServer_PTR() *QModbusTcpServer {
	return p
}

func (p *QModbusTcpServer) Pointer() unsafe.Pointer {
	if p != nil {
		return p.QModbusServer_PTR().Pointer()
	}
	return nil
}

func (p *QModbusTcpServer) SetPointer(ptr unsafe.Pointer) {
	if p != nil {
		p.QModbusServer_PTR().SetPointer(ptr)
	}
}

func PointerFromQModbusTcpServer(ptr QModbusTcpServer_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QModbusTcpServer_PTR().Pointer()
	}
	return nil
}

func NewQModbusTcpServerFromPointer(ptr unsafe.Pointer) *QModbusTcpServer {
	var n = new(QModbusTcpServer)
	n.SetPointer(ptr)
	return n
}

func newQModbusTcpServerFromPointer(ptr unsafe.Pointer) *QModbusTcpServer {
	var n = NewQModbusTcpServerFromPointer(ptr)
	for len(n.ObjectName()) < len("QModbusTcpServer_") {
		n.SetObjectName("QModbusTcpServer_" + qt.Identifier())
	}
	return n
}

func NewQModbusTcpServer(parent core.QObject_ITF) *QModbusTcpServer {
	defer qt.Recovering("QModbusTcpServer::QModbusTcpServer")

	return newQModbusTcpServerFromPointer(C.QModbusTcpServer_NewQModbusTcpServer(core.PointerFromQObject(parent)))
}

func (ptr *QModbusTcpServer) DestroyQModbusTcpServer() {
	defer qt.Recovering("QModbusTcpServer::~QModbusTcpServer")

	if ptr.Pointer() != nil {
		C.QModbusTcpServer_DestroyQModbusTcpServer(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQModbusTcpServer_ProcessPrivateRequest
func callbackQModbusTcpServer_ProcessPrivateRequest(ptr unsafe.Pointer, ptrName *C.char, request unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QModbusTcpServer::processPrivateRequest")

	if signal := qt.GetSignal(C.GoString(ptrName), "processPrivateRequest"); signal != nil {
		return PointerFromQModbusResponse(signal.(func(*QModbusPdu) *QModbusResponse)(NewQModbusPduFromPointer(request)))
	}

	return PointerFromQModbusResponse(NewQModbusTcpServerFromPointer(ptr).ProcessPrivateRequestDefault(NewQModbusPduFromPointer(request)))
}

func (ptr *QModbusTcpServer) ConnectProcessPrivateRequest(f func(request *QModbusPdu) *QModbusResponse) {
	defer qt.Recovering("connect QModbusTcpServer::processPrivateRequest")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "processPrivateRequest", f)
	}
}

func (ptr *QModbusTcpServer) DisconnectProcessPrivateRequest() {
	defer qt.Recovering("disconnect QModbusTcpServer::processPrivateRequest")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "processPrivateRequest")
	}
}

func (ptr *QModbusTcpServer) ProcessPrivateRequest(request QModbusPdu_ITF) *QModbusResponse {
	defer qt.Recovering("QModbusTcpServer::processPrivateRequest")

	if ptr.Pointer() != nil {

	}
	return nil
}

func (ptr *QModbusTcpServer) ProcessPrivateRequestDefault(request QModbusPdu_ITF) *QModbusResponse {
	defer qt.Recovering("QModbusTcpServer::processPrivateRequest")

	if ptr.Pointer() != nil {

	}
	return nil
}

//export callbackQModbusTcpServer_ProcessRequest
func callbackQModbusTcpServer_ProcessRequest(ptr unsafe.Pointer, ptrName *C.char, request unsafe.Pointer) unsafe.Pointer {
	defer qt.Recovering("callback QModbusTcpServer::processRequest")

	if signal := qt.GetSignal(C.GoString(ptrName), "processRequest"); signal != nil {
		return PointerFromQModbusResponse(signal.(func(*QModbusPdu) *QModbusResponse)(NewQModbusPduFromPointer(request)))
	}

	return PointerFromQModbusResponse(NewQModbusTcpServerFromPointer(ptr).ProcessRequestDefault(NewQModbusPduFromPointer(request)))
}

func (ptr *QModbusTcpServer) ConnectProcessRequest(f func(request *QModbusPdu) *QModbusResponse) {
	defer qt.Recovering("connect QModbusTcpServer::processRequest")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "processRequest", f)
	}
}

func (ptr *QModbusTcpServer) DisconnectProcessRequest() {
	defer qt.Recovering("disconnect QModbusTcpServer::processRequest")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "processRequest")
	}
}

func (ptr *QModbusTcpServer) ProcessRequest(request QModbusPdu_ITF) *QModbusResponse {
	defer qt.Recovering("QModbusTcpServer::processRequest")

	if ptr.Pointer() != nil {

	}
	return nil
}

func (ptr *QModbusTcpServer) ProcessRequestDefault(request QModbusPdu_ITF) *QModbusResponse {
	defer qt.Recovering("QModbusTcpServer::processRequest")

	if ptr.Pointer() != nil {

	}
	return nil
}

//export callbackQModbusTcpServer_ProcessesBroadcast
func callbackQModbusTcpServer_ProcessesBroadcast(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QModbusTcpServer::processesBroadcast")

	if signal := qt.GetSignal(C.GoString(ptrName), "processesBroadcast"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewQModbusTcpServerFromPointer(ptr).ProcessesBroadcastDefault()))
}

func (ptr *QModbusTcpServer) ConnectProcessesBroadcast(f func() bool) {
	defer qt.Recovering("connect QModbusTcpServer::processesBroadcast")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "processesBroadcast", f)
	}
}

func (ptr *QModbusTcpServer) DisconnectProcessesBroadcast() {
	defer qt.Recovering("disconnect QModbusTcpServer::processesBroadcast")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "processesBroadcast")
	}
}

func (ptr *QModbusTcpServer) ProcessesBroadcast() bool {
	defer qt.Recovering("QModbusTcpServer::processesBroadcast")

	if ptr.Pointer() != nil {
		return C.QModbusTcpServer_ProcessesBroadcast(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QModbusTcpServer) ProcessesBroadcastDefault() bool {
	defer qt.Recovering("QModbusTcpServer::processesBroadcast")

	if ptr.Pointer() != nil {
		return C.QModbusTcpServer_ProcessesBroadcastDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQModbusTcpServer_SetValue
func callbackQModbusTcpServer_SetValue(ptr unsafe.Pointer, ptrName *C.char, option C.int, newValue unsafe.Pointer) C.int {
	defer qt.Recovering("callback QModbusTcpServer::setValue")

	if signal := qt.GetSignal(C.GoString(ptrName), "setValue"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(int, *core.QVariant) bool)(int(option), core.NewQVariantFromPointer(newValue))))
	}

	return C.int(qt.GoBoolToInt(NewQModbusTcpServerFromPointer(ptr).SetValueDefault(int(option), core.NewQVariantFromPointer(newValue))))
}

func (ptr *QModbusTcpServer) ConnectSetValue(f func(option int, newValue *core.QVariant) bool) {
	defer qt.Recovering("connect QModbusTcpServer::setValue")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "setValue", f)
	}
}

func (ptr *QModbusTcpServer) DisconnectSetValue() {
	defer qt.Recovering("disconnect QModbusTcpServer::setValue")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "setValue")
	}
}

func (ptr *QModbusTcpServer) SetValue(option int, newValue core.QVariant_ITF) bool {
	defer qt.Recovering("QModbusTcpServer::setValue")

	if ptr.Pointer() != nil {
		return C.QModbusTcpServer_SetValue(ptr.Pointer(), C.int(option), core.PointerFromQVariant(newValue)) != 0
	}
	return false
}

func (ptr *QModbusTcpServer) SetValueDefault(option int, newValue core.QVariant_ITF) bool {
	defer qt.Recovering("QModbusTcpServer::setValue")

	if ptr.Pointer() != nil {
		return C.QModbusTcpServer_SetValueDefault(ptr.Pointer(), C.int(option), core.PointerFromQVariant(newValue)) != 0
	}
	return false
}

//export callbackQModbusTcpServer_Value
func callbackQModbusTcpServer_Value(ptr unsafe.Pointer, ptrName *C.char, option C.int) unsafe.Pointer {
	defer qt.Recovering("callback QModbusTcpServer::value")

	if signal := qt.GetSignal(C.GoString(ptrName), "value"); signal != nil {
		return core.PointerFromQVariant(signal.(func(int) *core.QVariant)(int(option)))
	}

	return core.PointerFromQVariant(NewQModbusTcpServerFromPointer(ptr).ValueDefault(int(option)))
}

func (ptr *QModbusTcpServer) ConnectValue(f func(option int) *core.QVariant) {
	defer qt.Recovering("connect QModbusTcpServer::value")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "value", f)
	}
}

func (ptr *QModbusTcpServer) DisconnectValue() {
	defer qt.Recovering("disconnect QModbusTcpServer::value")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "value")
	}
}

func (ptr *QModbusTcpServer) Value(option int) *core.QVariant {
	defer qt.Recovering("QModbusTcpServer::value")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QModbusTcpServer_Value(ptr.Pointer(), C.int(option)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QModbusTcpServer) ValueDefault(option int) *core.QVariant {
	defer qt.Recovering("QModbusTcpServer::value")

	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QModbusTcpServer_ValueDefault(ptr.Pointer(), C.int(option)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQModbusTcpServer_WriteData
func callbackQModbusTcpServer_WriteData(ptr unsafe.Pointer, ptrName *C.char, newData unsafe.Pointer) C.int {
	defer qt.Recovering("callback QModbusTcpServer::writeData")

	if signal := qt.GetSignal(C.GoString(ptrName), "writeData"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*QModbusDataUnit) bool)(NewQModbusDataUnitFromPointer(newData))))
	}

	return C.int(qt.GoBoolToInt(NewQModbusTcpServerFromPointer(ptr).WriteDataDefault(NewQModbusDataUnitFromPointer(newData))))
}

func (ptr *QModbusTcpServer) ConnectWriteData(f func(newData *QModbusDataUnit) bool) {
	defer qt.Recovering("connect QModbusTcpServer::writeData")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "writeData", f)
	}
}

func (ptr *QModbusTcpServer) DisconnectWriteData() {
	defer qt.Recovering("disconnect QModbusTcpServer::writeData")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "writeData")
	}
}

func (ptr *QModbusTcpServer) WriteData(newData QModbusDataUnit_ITF) bool {
	defer qt.Recovering("QModbusTcpServer::writeData")

	if ptr.Pointer() != nil {
		return C.QModbusTcpServer_WriteData(ptr.Pointer(), PointerFromQModbusDataUnit(newData)) != 0
	}
	return false
}

func (ptr *QModbusTcpServer) WriteDataDefault(newData QModbusDataUnit_ITF) bool {
	defer qt.Recovering("QModbusTcpServer::writeData")

	if ptr.Pointer() != nil {
		return C.QModbusTcpServer_WriteDataDefault(ptr.Pointer(), PointerFromQModbusDataUnit(newData)) != 0
	}
	return false
}

//export callbackQModbusTcpServer_Close
func callbackQModbusTcpServer_Close(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QModbusTcpServer::close")

	if signal := qt.GetSignal(C.GoString(ptrName), "close"); signal != nil {
		signal.(func())()
	} else {
		NewQModbusTcpServerFromPointer(ptr).CloseDefault()
	}
}

func (ptr *QModbusTcpServer) ConnectClose(f func()) {
	defer qt.Recovering("connect QModbusTcpServer::close")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "close", f)
	}
}

func (ptr *QModbusTcpServer) DisconnectClose() {
	defer qt.Recovering("disconnect QModbusTcpServer::close")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "close")
	}
}

func (ptr *QModbusTcpServer) Close() {
	defer qt.Recovering("QModbusTcpServer::close")

	if ptr.Pointer() != nil {
		C.QModbusTcpServer_Close(ptr.Pointer())
	}
}

func (ptr *QModbusTcpServer) CloseDefault() {
	defer qt.Recovering("QModbusTcpServer::close")

	if ptr.Pointer() != nil {
		C.QModbusTcpServer_CloseDefault(ptr.Pointer())
	}
}

//export callbackQModbusTcpServer_Open
func callbackQModbusTcpServer_Open(ptr unsafe.Pointer, ptrName *C.char) C.int {
	defer qt.Recovering("callback QModbusTcpServer::open")

	if signal := qt.GetSignal(C.GoString(ptrName), "open"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func() bool)()))
	}

	return C.int(qt.GoBoolToInt(NewQModbusTcpServerFromPointer(ptr).OpenDefault()))
}

func (ptr *QModbusTcpServer) ConnectOpen(f func() bool) {
	defer qt.Recovering("connect QModbusTcpServer::open")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "open", f)
	}
}

func (ptr *QModbusTcpServer) DisconnectOpen() {
	defer qt.Recovering("disconnect QModbusTcpServer::open")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "open")
	}
}

func (ptr *QModbusTcpServer) Open() bool {
	defer qt.Recovering("QModbusTcpServer::open")

	if ptr.Pointer() != nil {
		return C.QModbusTcpServer_Open(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QModbusTcpServer) OpenDefault() bool {
	defer qt.Recovering("QModbusTcpServer::open")

	if ptr.Pointer() != nil {
		return C.QModbusTcpServer_OpenDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQModbusTcpServer_TimerEvent
func callbackQModbusTcpServer_TimerEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QModbusTcpServer::timerEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQModbusTcpServerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QModbusTcpServer) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	defer qt.Recovering("connect QModbusTcpServer::timerEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "timerEvent", f)
	}
}

func (ptr *QModbusTcpServer) DisconnectTimerEvent() {
	defer qt.Recovering("disconnect QModbusTcpServer::timerEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "timerEvent")
	}
}

func (ptr *QModbusTcpServer) TimerEvent(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QModbusTcpServer::timerEvent")

	if ptr.Pointer() != nil {
		C.QModbusTcpServer_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QModbusTcpServer) TimerEventDefault(event core.QTimerEvent_ITF) {
	defer qt.Recovering("QModbusTcpServer::timerEvent")

	if ptr.Pointer() != nil {
		C.QModbusTcpServer_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQModbusTcpServer_ChildEvent
func callbackQModbusTcpServer_ChildEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QModbusTcpServer::childEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQModbusTcpServerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QModbusTcpServer) ConnectChildEvent(f func(event *core.QChildEvent)) {
	defer qt.Recovering("connect QModbusTcpServer::childEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "childEvent", f)
	}
}

func (ptr *QModbusTcpServer) DisconnectChildEvent() {
	defer qt.Recovering("disconnect QModbusTcpServer::childEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "childEvent")
	}
}

func (ptr *QModbusTcpServer) ChildEvent(event core.QChildEvent_ITF) {
	defer qt.Recovering("QModbusTcpServer::childEvent")

	if ptr.Pointer() != nil {
		C.QModbusTcpServer_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QModbusTcpServer) ChildEventDefault(event core.QChildEvent_ITF) {
	defer qt.Recovering("QModbusTcpServer::childEvent")

	if ptr.Pointer() != nil {
		C.QModbusTcpServer_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQModbusTcpServer_ConnectNotify
func callbackQModbusTcpServer_ConnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QModbusTcpServer::connectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQModbusTcpServerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QModbusTcpServer) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QModbusTcpServer::connectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "connectNotify", f)
	}
}

func (ptr *QModbusTcpServer) DisconnectConnectNotify() {
	defer qt.Recovering("disconnect QModbusTcpServer::connectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "connectNotify")
	}
}

func (ptr *QModbusTcpServer) ConnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QModbusTcpServer::connectNotify")

	if ptr.Pointer() != nil {
		C.QModbusTcpServer_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QModbusTcpServer) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QModbusTcpServer::connectNotify")

	if ptr.Pointer() != nil {
		C.QModbusTcpServer_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQModbusTcpServer_CustomEvent
func callbackQModbusTcpServer_CustomEvent(ptr unsafe.Pointer, ptrName *C.char, event unsafe.Pointer) {
	defer qt.Recovering("callback QModbusTcpServer::customEvent")

	if signal := qt.GetSignal(C.GoString(ptrName), "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQModbusTcpServerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QModbusTcpServer) ConnectCustomEvent(f func(event *core.QEvent)) {
	defer qt.Recovering("connect QModbusTcpServer::customEvent")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "customEvent", f)
	}
}

func (ptr *QModbusTcpServer) DisconnectCustomEvent() {
	defer qt.Recovering("disconnect QModbusTcpServer::customEvent")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "customEvent")
	}
}

func (ptr *QModbusTcpServer) CustomEvent(event core.QEvent_ITF) {
	defer qt.Recovering("QModbusTcpServer::customEvent")

	if ptr.Pointer() != nil {
		C.QModbusTcpServer_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QModbusTcpServer) CustomEventDefault(event core.QEvent_ITF) {
	defer qt.Recovering("QModbusTcpServer::customEvent")

	if ptr.Pointer() != nil {
		C.QModbusTcpServer_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQModbusTcpServer_DeleteLater
func callbackQModbusTcpServer_DeleteLater(ptr unsafe.Pointer, ptrName *C.char) {
	defer qt.Recovering("callback QModbusTcpServer::deleteLater")

	if signal := qt.GetSignal(C.GoString(ptrName), "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQModbusTcpServerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QModbusTcpServer) ConnectDeleteLater(f func()) {
	defer qt.Recovering("connect QModbusTcpServer::deleteLater")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "deleteLater", f)
	}
}

func (ptr *QModbusTcpServer) DisconnectDeleteLater() {
	defer qt.Recovering("disconnect QModbusTcpServer::deleteLater")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "deleteLater")
	}
}

func (ptr *QModbusTcpServer) DeleteLater() {
	defer qt.Recovering("QModbusTcpServer::deleteLater")

	if ptr.Pointer() != nil {
		C.QModbusTcpServer_DeleteLater(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QModbusTcpServer) DeleteLaterDefault() {
	defer qt.Recovering("QModbusTcpServer::deleteLater")

	if ptr.Pointer() != nil {
		C.QModbusTcpServer_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//export callbackQModbusTcpServer_DisconnectNotify
func callbackQModbusTcpServer_DisconnectNotify(ptr unsafe.Pointer, ptrName *C.char, sign unsafe.Pointer) {
	defer qt.Recovering("callback QModbusTcpServer::disconnectNotify")

	if signal := qt.GetSignal(C.GoString(ptrName), "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQModbusTcpServerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QModbusTcpServer) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	defer qt.Recovering("connect QModbusTcpServer::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "disconnectNotify", f)
	}
}

func (ptr *QModbusTcpServer) DisconnectDisconnectNotify() {
	defer qt.Recovering("disconnect QModbusTcpServer::disconnectNotify")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "disconnectNotify")
	}
}

func (ptr *QModbusTcpServer) DisconnectNotify(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QModbusTcpServer::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QModbusTcpServer_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QModbusTcpServer) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	defer qt.Recovering("QModbusTcpServer::disconnectNotify")

	if ptr.Pointer() != nil {
		C.QModbusTcpServer_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQModbusTcpServer_Event
func callbackQModbusTcpServer_Event(ptr unsafe.Pointer, ptrName *C.char, e unsafe.Pointer) C.int {
	defer qt.Recovering("callback QModbusTcpServer::event")

	if signal := qt.GetSignal(C.GoString(ptrName), "event"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e))))
	}

	return C.int(qt.GoBoolToInt(NewQModbusTcpServerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e))))
}

func (ptr *QModbusTcpServer) ConnectEvent(f func(e *core.QEvent) bool) {
	defer qt.Recovering("connect QModbusTcpServer::event")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "event", f)
	}
}

func (ptr *QModbusTcpServer) DisconnectEvent() {
	defer qt.Recovering("disconnect QModbusTcpServer::event")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "event")
	}
}

func (ptr *QModbusTcpServer) Event(e core.QEvent_ITF) bool {
	defer qt.Recovering("QModbusTcpServer::event")

	if ptr.Pointer() != nil {
		return C.QModbusTcpServer_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QModbusTcpServer) EventDefault(e core.QEvent_ITF) bool {
	defer qt.Recovering("QModbusTcpServer::event")

	if ptr.Pointer() != nil {
		return C.QModbusTcpServer_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQModbusTcpServer_EventFilter
func callbackQModbusTcpServer_EventFilter(ptr unsafe.Pointer, ptrName *C.char, watched unsafe.Pointer, event unsafe.Pointer) C.int {
	defer qt.Recovering("callback QModbusTcpServer::eventFilter")

	if signal := qt.GetSignal(C.GoString(ptrName), "eventFilter"); signal != nil {
		return C.int(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
	}

	return C.int(qt.GoBoolToInt(NewQModbusTcpServerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event))))
}

func (ptr *QModbusTcpServer) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	defer qt.Recovering("connect QModbusTcpServer::eventFilter")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "eventFilter", f)
	}
}

func (ptr *QModbusTcpServer) DisconnectEventFilter() {
	defer qt.Recovering("disconnect QModbusTcpServer::eventFilter")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "eventFilter")
	}
}

func (ptr *QModbusTcpServer) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QModbusTcpServer::eventFilter")

	if ptr.Pointer() != nil {
		return C.QModbusTcpServer_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QModbusTcpServer) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	defer qt.Recovering("QModbusTcpServer::eventFilter")

	if ptr.Pointer() != nil {
		return C.QModbusTcpServer_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQModbusTcpServer_MetaObject
func callbackQModbusTcpServer_MetaObject(ptr unsafe.Pointer, ptrName *C.char) unsafe.Pointer {
	defer qt.Recovering("callback QModbusTcpServer::metaObject")

	if signal := qt.GetSignal(C.GoString(ptrName), "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQModbusTcpServerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QModbusTcpServer) ConnectMetaObject(f func() *core.QMetaObject) {
	defer qt.Recovering("connect QModbusTcpServer::metaObject")

	if ptr.Pointer() != nil {

		qt.ConnectSignal(ptr.ObjectName(), "metaObject", f)
	}
}

func (ptr *QModbusTcpServer) DisconnectMetaObject() {
	defer qt.Recovering("disconnect QModbusTcpServer::metaObject")

	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.ObjectName(), "metaObject")
	}
}

func (ptr *QModbusTcpServer) MetaObject() *core.QMetaObject {
	defer qt.Recovering("QModbusTcpServer::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QModbusTcpServer_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QModbusTcpServer) MetaObjectDefault() *core.QMetaObject {
	defer qt.Recovering("QModbusTcpServer::metaObject")

	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QModbusTcpServer_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}
