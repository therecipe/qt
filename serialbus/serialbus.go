// +build !minimal

package serialbus

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "serialbus.h"
import "C"
import (
	"fmt"
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"runtime"
	"unsafe"
)

func cGoUnpackString(s C.struct_QtSerialBus_PackedString) string {
	if len := int(s.len); len == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}

type Diagnostics struct {
	ptr unsafe.Pointer
}

type Diagnostics_ITF interface {
	Diagnostics_PTR() *Diagnostics
}

func (ptr *Diagnostics) Diagnostics_PTR() *Diagnostics {
	return ptr
}

func (ptr *Diagnostics) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *Diagnostics) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromDiagnostics(ptr Diagnostics_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.Diagnostics_PTR().Pointer()
	}
	return nil
}

func NewDiagnosticsFromPointer(ptr unsafe.Pointer) *Diagnostics {
	var n = new(Diagnostics)
	n.SetPointer(ptr)
	return n
}

func (ptr *Diagnostics) DestroyDiagnostics() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=Diagnostics__SubFunctionCode
//Diagnostics::SubFunctionCode
type Diagnostics__SubFunctionCode int64

const (
	Diagnostics__ReturnQueryData                    Diagnostics__SubFunctionCode = Diagnostics__SubFunctionCode(0x0000)
	Diagnostics__RestartCommunicationsOption        Diagnostics__SubFunctionCode = Diagnostics__SubFunctionCode(0x0001)
	Diagnostics__ReturnDiagnosticRegister           Diagnostics__SubFunctionCode = Diagnostics__SubFunctionCode(0x0002)
	Diagnostics__ChangeAsciiInputDelimiter          Diagnostics__SubFunctionCode = Diagnostics__SubFunctionCode(0x0003)
	Diagnostics__ForceListenOnlyMode                Diagnostics__SubFunctionCode = Diagnostics__SubFunctionCode(0x0004)
	Diagnostics__ClearCountersAndDiagnosticRegister Diagnostics__SubFunctionCode = Diagnostics__SubFunctionCode(0x000a)
	Diagnostics__ReturnBusMessageCount              Diagnostics__SubFunctionCode = Diagnostics__SubFunctionCode(0x000b)
	Diagnostics__ReturnBusCommunicationErrorCount   Diagnostics__SubFunctionCode = Diagnostics__SubFunctionCode(0x000c)
	Diagnostics__ReturnBusExceptionErrorCount       Diagnostics__SubFunctionCode = Diagnostics__SubFunctionCode(0x000d)
	Diagnostics__ReturnServerMessageCount           Diagnostics__SubFunctionCode = Diagnostics__SubFunctionCode(0x000e)
	Diagnostics__ReturnServerNoResponseCount        Diagnostics__SubFunctionCode = Diagnostics__SubFunctionCode(0x000f)
	Diagnostics__ReturnServerNAKCount               Diagnostics__SubFunctionCode = Diagnostics__SubFunctionCode(0x0010)
	Diagnostics__ReturnServerBusyCount              Diagnostics__SubFunctionCode = Diagnostics__SubFunctionCode(0x0011)
	Diagnostics__ReturnBusCharacterOverrunCount     Diagnostics__SubFunctionCode = Diagnostics__SubFunctionCode(0x0012)
	Diagnostics__ClearOverrunCounterAndFlag         Diagnostics__SubFunctionCode = Diagnostics__SubFunctionCode(0x0014)
)

type EncapsulatedInterfaceTransport struct {
	ptr unsafe.Pointer
}

type EncapsulatedInterfaceTransport_ITF interface {
	EncapsulatedInterfaceTransport_PTR() *EncapsulatedInterfaceTransport
}

func (ptr *EncapsulatedInterfaceTransport) EncapsulatedInterfaceTransport_PTR() *EncapsulatedInterfaceTransport {
	return ptr
}

func (ptr *EncapsulatedInterfaceTransport) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *EncapsulatedInterfaceTransport) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromEncapsulatedInterfaceTransport(ptr EncapsulatedInterfaceTransport_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.EncapsulatedInterfaceTransport_PTR().Pointer()
	}
	return nil
}

func NewEncapsulatedInterfaceTransportFromPointer(ptr unsafe.Pointer) *EncapsulatedInterfaceTransport {
	var n = new(EncapsulatedInterfaceTransport)
	n.SetPointer(ptr)
	return n
}

func (ptr *EncapsulatedInterfaceTransport) DestroyEncapsulatedInterfaceTransport() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=EncapsulatedInterfaceTransport__SubFunctionCode
//EncapsulatedInterfaceTransport::SubFunctionCode
type EncapsulatedInterfaceTransport__SubFunctionCode int64

const (
	EncapsulatedInterfaceTransport__CanOpenGeneralReference  EncapsulatedInterfaceTransport__SubFunctionCode = EncapsulatedInterfaceTransport__SubFunctionCode(0x0D)
	EncapsulatedInterfaceTransport__ReadDeviceIdentification EncapsulatedInterfaceTransport__SubFunctionCode = EncapsulatedInterfaceTransport__SubFunctionCode(0x0E)
)

type QCanBus struct {
	core.QObject
}

type QCanBus_ITF interface {
	core.QObject_ITF
	QCanBus_PTR() *QCanBus
}

func (ptr *QCanBus) QCanBus_PTR() *QCanBus {
	return ptr
}

func (ptr *QCanBus) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QCanBus) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
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

func (ptr *QCanBus) DestroyQCanBus() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QCanBus) CreateDevice(plugin core.QByteArray_ITF, interfaceName string) *QCanBusDevice {
	if ptr.Pointer() != nil {
		var interfaceNameC = C.CString(interfaceName)
		defer C.free(unsafe.Pointer(interfaceNameC))
		var tmpValue = NewQCanBusDeviceFromPointer(C.QCanBus_CreateDevice(ptr.Pointer(), core.PointerFromQByteArray(plugin), interfaceNameC))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func QCanBus_Instance() *QCanBus {
	var tmpValue = NewQCanBusFromPointer(C.QCanBus_QCanBus_Instance())
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QCanBus) Instance() *QCanBus {
	var tmpValue = NewQCanBusFromPointer(C.QCanBus_QCanBus_Instance())
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QCanBus) Plugins() []*core.QByteArray {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtSerialBus_PackedList) []*core.QByteArray {
			var out = make([]*core.QByteArray, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQCanBusFromPointer(l.data).plugins_atList(i)
			}
			return out
		}(C.QCanBus_Plugins(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCanBus) plugins_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QCanBus_plugins_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

//export callbackQCanBus_TimerEvent
func callbackQCanBus_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCanBus::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQCanBusFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QCanBus) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBus::timerEvent", f)
	}
}

func (ptr *QCanBus) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBus::timerEvent")
	}
}

func (ptr *QCanBus) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QCanBus_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCanBus) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QCanBus_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQCanBus_ChildEvent
func callbackQCanBus_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCanBus::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCanBusFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCanBus) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBus::childEvent", f)
	}
}

func (ptr *QCanBus) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBus::childEvent")
	}
}

func (ptr *QCanBus) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QCanBus_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCanBus) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QCanBus_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQCanBus_ConnectNotify
func callbackQCanBus_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCanBus::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQCanBusFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QCanBus) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBus::connectNotify", f)
	}
}

func (ptr *QCanBus) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBus::connectNotify")
	}
}

func (ptr *QCanBus) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QCanBus_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QCanBus) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QCanBus_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQCanBus_CustomEvent
func callbackQCanBus_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCanBus::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCanBusFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCanBus) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBus::customEvent", f)
	}
}

func (ptr *QCanBus) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBus::customEvent")
	}
}

func (ptr *QCanBus) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QCanBus_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCanBus) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QCanBus_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQCanBus_DeleteLater
func callbackQCanBus_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCanBus::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQCanBusFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QCanBus) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBus::deleteLater", f)
	}
}

func (ptr *QCanBus) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBus::deleteLater")
	}
}

func (ptr *QCanBus) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QCanBus_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QCanBus) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QCanBus_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQCanBus_DisconnectNotify
func callbackQCanBus_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCanBus::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQCanBusFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QCanBus) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBus::disconnectNotify", f)
	}
}

func (ptr *QCanBus) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBus::disconnectNotify")
	}
}

func (ptr *QCanBus) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QCanBus_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QCanBus) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QCanBus_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQCanBus_Event
func callbackQCanBus_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCanBus::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQCanBusFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QCanBus) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBus::event", f)
	}
}

func (ptr *QCanBus) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBus::event")
	}
}

func (ptr *QCanBus) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QCanBus_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QCanBus) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QCanBus_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQCanBus_EventFilter
func callbackQCanBus_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCanBus::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQCanBusFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QCanBus) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBus::eventFilter", f)
	}
}

func (ptr *QCanBus) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBus::eventFilter")
	}
}

func (ptr *QCanBus) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QCanBus_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QCanBus) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QCanBus_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQCanBus_MetaObject
func callbackQCanBus_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCanBus::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQCanBusFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QCanBus) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBus::metaObject", f)
	}
}

func (ptr *QCanBus) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBus::metaObject")
	}
}

func (ptr *QCanBus) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QCanBus_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCanBus) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QCanBus_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QCanBusDevice struct {
	core.QObject
}

type QCanBusDevice_ITF interface {
	core.QObject_ITF
	QCanBusDevice_PTR() *QCanBusDevice
}

func (ptr *QCanBusDevice) QCanBusDevice_PTR() *QCanBusDevice {
	return ptr
}

func (ptr *QCanBusDevice) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QCanBusDevice) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
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

func (ptr *QCanBusDevice) DestroyQCanBusDevice() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QCanBusDevice__CanBusDeviceState
//QCanBusDevice::CanBusDeviceState
type QCanBusDevice__CanBusDeviceState int64

const (
	QCanBusDevice__UnconnectedState QCanBusDevice__CanBusDeviceState = QCanBusDevice__CanBusDeviceState(0)
	QCanBusDevice__ConnectingState  QCanBusDevice__CanBusDeviceState = QCanBusDevice__CanBusDeviceState(1)
	QCanBusDevice__ConnectedState   QCanBusDevice__CanBusDeviceState = QCanBusDevice__CanBusDeviceState(2)
	QCanBusDevice__ClosingState     QCanBusDevice__CanBusDeviceState = QCanBusDevice__CanBusDeviceState(3)
)

//go:generate stringer -type=QCanBusDevice__CanBusError
//QCanBusDevice::CanBusError
type QCanBusDevice__CanBusError int64

const (
	QCanBusDevice__NoError            QCanBusDevice__CanBusError = QCanBusDevice__CanBusError(0)
	QCanBusDevice__ReadError          QCanBusDevice__CanBusError = QCanBusDevice__CanBusError(1)
	QCanBusDevice__WriteError         QCanBusDevice__CanBusError = QCanBusDevice__CanBusError(2)
	QCanBusDevice__ConnectionError    QCanBusDevice__CanBusError = QCanBusDevice__CanBusError(3)
	QCanBusDevice__ConfigurationError QCanBusDevice__CanBusError = QCanBusDevice__CanBusError(4)
	QCanBusDevice__UnknownError       QCanBusDevice__CanBusError = QCanBusDevice__CanBusError(5)
)

//go:generate stringer -type=QCanBusDevice__ConfigurationKey
//QCanBusDevice::ConfigurationKey
type QCanBusDevice__ConfigurationKey int64

const (
	QCanBusDevice__RawFilterKey   QCanBusDevice__ConfigurationKey = QCanBusDevice__ConfigurationKey(0)
	QCanBusDevice__ErrorFilterKey QCanBusDevice__ConfigurationKey = QCanBusDevice__ConfigurationKey(1)
	QCanBusDevice__LoopbackKey    QCanBusDevice__ConfigurationKey = QCanBusDevice__ConfigurationKey(2)
	QCanBusDevice__ReceiveOwnKey  QCanBusDevice__ConfigurationKey = QCanBusDevice__ConfigurationKey(3)
	QCanBusDevice__BitRateKey     QCanBusDevice__ConfigurationKey = QCanBusDevice__ConfigurationKey(4)
	QCanBusDevice__CanFdKey       QCanBusDevice__ConfigurationKey = QCanBusDevice__ConfigurationKey(5)
	QCanBusDevice__UserKey        QCanBusDevice__ConfigurationKey = QCanBusDevice__ConfigurationKey(30)
)

func (ptr *QCanBusDevice) FramesAvailable() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QCanBusDevice_FramesAvailable(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCanBusDevice) FramesToWrite() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QCanBusDevice_FramesToWrite(ptr.Pointer()))
	}
	return 0
}

func NewQCanBusDevice(parent core.QObject_ITF) *QCanBusDevice {
	var tmpValue = NewQCanBusDeviceFromPointer(C.QCanBusDevice_NewQCanBusDevice(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQCanBusDevice_Close
func callbackQCanBusDevice_Close(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCanBusDevice::close"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QCanBusDevice) ConnectClose(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBusDevice::close", f)
	}
}

func (ptr *QCanBusDevice) DisconnectClose() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBusDevice::close")
	}
}

func (ptr *QCanBusDevice) Close() {
	if ptr.Pointer() != nil {
		C.QCanBusDevice_Close(ptr.Pointer())
	}
}

func (ptr *QCanBusDevice) ConfigurationParameter(key int) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QCanBusDevice_ConfigurationParameter(ptr.Pointer(), C.int(int32(key))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QCanBusDevice) ConnectDevice() bool {
	if ptr.Pointer() != nil {
		return C.QCanBusDevice_ConnectDevice(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCanBusDevice) DisconnectDevice() {
	if ptr.Pointer() != nil {
		C.QCanBusDevice_DisconnectDevice(ptr.Pointer())
	}
}

func (ptr *QCanBusDevice) EnqueueOutgoingFrame(newFrame QCanBusFrame_ITF) {
	if ptr.Pointer() != nil {
		C.QCanBusDevice_EnqueueOutgoingFrame(ptr.Pointer(), PointerFromQCanBusFrame(newFrame))
	}
}

func (ptr *QCanBusDevice) Error() QCanBusDevice__CanBusError {
	if ptr.Pointer() != nil {
		return QCanBusDevice__CanBusError(C.QCanBusDevice_Error(ptr.Pointer()))
	}
	return 0
}

//export callbackQCanBusDevice_ErrorOccurred
func callbackQCanBusDevice_ErrorOccurred(ptr unsafe.Pointer, error C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCanBusDevice::errorOccurred"); signal != nil {
		signal.(func(QCanBusDevice__CanBusError))(QCanBusDevice__CanBusError(error))
	}

}

func (ptr *QCanBusDevice) ConnectErrorOccurred(f func(error QCanBusDevice__CanBusError)) {
	if ptr.Pointer() != nil {
		C.QCanBusDevice_ConnectErrorOccurred(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBusDevice::errorOccurred", f)
	}
}

func (ptr *QCanBusDevice) DisconnectErrorOccurred() {
	if ptr.Pointer() != nil {
		C.QCanBusDevice_DisconnectErrorOccurred(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBusDevice::errorOccurred")
	}
}

func (ptr *QCanBusDevice) ErrorOccurred(error QCanBusDevice__CanBusError) {
	if ptr.Pointer() != nil {
		C.QCanBusDevice_ErrorOccurred(ptr.Pointer(), C.longlong(error))
	}
}

func (ptr *QCanBusDevice) ErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QCanBusDevice_ErrorString(ptr.Pointer()))
	}
	return ""
}

//export callbackQCanBusDevice_FramesReceived
func callbackQCanBusDevice_FramesReceived(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCanBusDevice::framesReceived"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QCanBusDevice) ConnectFramesReceived(f func()) {
	if ptr.Pointer() != nil {
		C.QCanBusDevice_ConnectFramesReceived(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBusDevice::framesReceived", f)
	}
}

func (ptr *QCanBusDevice) DisconnectFramesReceived() {
	if ptr.Pointer() != nil {
		C.QCanBusDevice_DisconnectFramesReceived(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBusDevice::framesReceived")
	}
}

func (ptr *QCanBusDevice) FramesReceived() {
	if ptr.Pointer() != nil {
		C.QCanBusDevice_FramesReceived(ptr.Pointer())
	}
}

//export callbackQCanBusDevice_FramesWritten
func callbackQCanBusDevice_FramesWritten(ptr unsafe.Pointer, framesCount C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCanBusDevice::framesWritten"); signal != nil {
		signal.(func(int64))(int64(framesCount))
	}

}

func (ptr *QCanBusDevice) ConnectFramesWritten(f func(framesCount int64)) {
	if ptr.Pointer() != nil {
		C.QCanBusDevice_ConnectFramesWritten(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBusDevice::framesWritten", f)
	}
}

func (ptr *QCanBusDevice) DisconnectFramesWritten() {
	if ptr.Pointer() != nil {
		C.QCanBusDevice_DisconnectFramesWritten(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBusDevice::framesWritten")
	}
}

func (ptr *QCanBusDevice) FramesWritten(framesCount int64) {
	if ptr.Pointer() != nil {
		C.QCanBusDevice_FramesWritten(ptr.Pointer(), C.longlong(framesCount))
	}
}

func (ptr *QCanBusDevice) HasOutgoingFrames() bool {
	if ptr.Pointer() != nil {
		return C.QCanBusDevice_HasOutgoingFrames(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQCanBusDevice_InterpretErrorFrame
func callbackQCanBusDevice_InterpretErrorFrame(ptr unsafe.Pointer, frame unsafe.Pointer) *C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCanBusDevice::interpretErrorFrame"); signal != nil {
		return C.CString(signal.(func(*QCanBusFrame) string)(NewQCanBusFrameFromPointer(frame)))
	}

	return C.CString("")
}

func (ptr *QCanBusDevice) ConnectInterpretErrorFrame(f func(frame *QCanBusFrame) string) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBusDevice::interpretErrorFrame", f)
	}
}

func (ptr *QCanBusDevice) DisconnectInterpretErrorFrame() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBusDevice::interpretErrorFrame")
	}
}

func (ptr *QCanBusDevice) InterpretErrorFrame(frame QCanBusFrame_ITF) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QCanBusDevice_InterpretErrorFrame(ptr.Pointer(), PointerFromQCanBusFrame(frame)))
	}
	return ""
}

//export callbackQCanBusDevice_Open
func callbackQCanBusDevice_Open(ptr unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCanBusDevice::open"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QCanBusDevice) ConnectOpen(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBusDevice::open", f)
	}
}

func (ptr *QCanBusDevice) DisconnectOpen() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBusDevice::open")
	}
}

func (ptr *QCanBusDevice) Open() bool {
	if ptr.Pointer() != nil {
		return C.QCanBusDevice_Open(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQCanBusDevice_SetConfigurationParameter
func callbackQCanBusDevice_SetConfigurationParameter(ptr unsafe.Pointer, key C.int, value unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCanBusDevice::setConfigurationParameter"); signal != nil {
		signal.(func(int, *core.QVariant))(int(int32(key)), core.NewQVariantFromPointer(value))
	} else {
		NewQCanBusDeviceFromPointer(ptr).SetConfigurationParameterDefault(int(int32(key)), core.NewQVariantFromPointer(value))
	}
}

func (ptr *QCanBusDevice) ConnectSetConfigurationParameter(f func(key int, value *core.QVariant)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBusDevice::setConfigurationParameter", f)
	}
}

func (ptr *QCanBusDevice) DisconnectSetConfigurationParameter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBusDevice::setConfigurationParameter")
	}
}

func (ptr *QCanBusDevice) SetConfigurationParameter(key int, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QCanBusDevice_SetConfigurationParameter(ptr.Pointer(), C.int(int32(key)), core.PointerFromQVariant(value))
	}
}

func (ptr *QCanBusDevice) SetConfigurationParameterDefault(key int, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QCanBusDevice_SetConfigurationParameterDefault(ptr.Pointer(), C.int(int32(key)), core.PointerFromQVariant(value))
	}
}

func (ptr *QCanBusDevice) SetError(errorText string, errorId QCanBusDevice__CanBusError) {
	if ptr.Pointer() != nil {
		var errorTextC = C.CString(errorText)
		defer C.free(unsafe.Pointer(errorTextC))
		C.QCanBusDevice_SetError(ptr.Pointer(), errorTextC, C.longlong(errorId))
	}
}

func (ptr *QCanBusDevice) SetState(newState QCanBusDevice__CanBusDeviceState) {
	if ptr.Pointer() != nil {
		C.QCanBusDevice_SetState(ptr.Pointer(), C.longlong(newState))
	}
}

func (ptr *QCanBusDevice) State() QCanBusDevice__CanBusDeviceState {
	if ptr.Pointer() != nil {
		return QCanBusDevice__CanBusDeviceState(C.QCanBusDevice_State(ptr.Pointer()))
	}
	return 0
}

//export callbackQCanBusDevice_StateChanged
func callbackQCanBusDevice_StateChanged(ptr unsafe.Pointer, state C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCanBusDevice::stateChanged"); signal != nil {
		signal.(func(QCanBusDevice__CanBusDeviceState))(QCanBusDevice__CanBusDeviceState(state))
	}

}

func (ptr *QCanBusDevice) ConnectStateChanged(f func(state QCanBusDevice__CanBusDeviceState)) {
	if ptr.Pointer() != nil {
		C.QCanBusDevice_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBusDevice::stateChanged", f)
	}
}

func (ptr *QCanBusDevice) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QCanBusDevice_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBusDevice::stateChanged")
	}
}

func (ptr *QCanBusDevice) StateChanged(state QCanBusDevice__CanBusDeviceState) {
	if ptr.Pointer() != nil {
		C.QCanBusDevice_StateChanged(ptr.Pointer(), C.longlong(state))
	}
}

//export callbackQCanBusDevice_WriteFrame
func callbackQCanBusDevice_WriteFrame(ptr unsafe.Pointer, frame unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCanBusDevice::writeFrame"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QCanBusFrame) bool)(NewQCanBusFrameFromPointer(frame)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QCanBusDevice) ConnectWriteFrame(f func(frame *QCanBusFrame) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBusDevice::writeFrame", f)
	}
}

func (ptr *QCanBusDevice) DisconnectWriteFrame() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBusDevice::writeFrame")
	}
}

func (ptr *QCanBusDevice) WriteFrame(frame QCanBusFrame_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QCanBusDevice_WriteFrame(ptr.Pointer(), PointerFromQCanBusFrame(frame)) != 0
	}
	return false
}

//export callbackQCanBusDevice_TimerEvent
func callbackQCanBusDevice_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCanBusDevice::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQCanBusDeviceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QCanBusDevice) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBusDevice::timerEvent", f)
	}
}

func (ptr *QCanBusDevice) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBusDevice::timerEvent")
	}
}

func (ptr *QCanBusDevice) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QCanBusDevice_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QCanBusDevice) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QCanBusDevice_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQCanBusDevice_ChildEvent
func callbackQCanBusDevice_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCanBusDevice::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCanBusDeviceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCanBusDevice) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBusDevice::childEvent", f)
	}
}

func (ptr *QCanBusDevice) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBusDevice::childEvent")
	}
}

func (ptr *QCanBusDevice) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QCanBusDevice_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QCanBusDevice) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QCanBusDevice_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQCanBusDevice_ConnectNotify
func callbackQCanBusDevice_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCanBusDevice::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQCanBusDeviceFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QCanBusDevice) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBusDevice::connectNotify", f)
	}
}

func (ptr *QCanBusDevice) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBusDevice::connectNotify")
	}
}

func (ptr *QCanBusDevice) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QCanBusDevice_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QCanBusDevice) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QCanBusDevice_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQCanBusDevice_CustomEvent
func callbackQCanBusDevice_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCanBusDevice::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCanBusDeviceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCanBusDevice) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBusDevice::customEvent", f)
	}
}

func (ptr *QCanBusDevice) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBusDevice::customEvent")
	}
}

func (ptr *QCanBusDevice) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QCanBusDevice_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QCanBusDevice) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QCanBusDevice_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQCanBusDevice_DeleteLater
func callbackQCanBusDevice_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCanBusDevice::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQCanBusDeviceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QCanBusDevice) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBusDevice::deleteLater", f)
	}
}

func (ptr *QCanBusDevice) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBusDevice::deleteLater")
	}
}

func (ptr *QCanBusDevice) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QCanBusDevice_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QCanBusDevice) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QCanBusDevice_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQCanBusDevice_DisconnectNotify
func callbackQCanBusDevice_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCanBusDevice::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQCanBusDeviceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QCanBusDevice) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBusDevice::disconnectNotify", f)
	}
}

func (ptr *QCanBusDevice) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBusDevice::disconnectNotify")
	}
}

func (ptr *QCanBusDevice) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QCanBusDevice_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QCanBusDevice) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QCanBusDevice_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQCanBusDevice_Event
func callbackQCanBusDevice_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCanBusDevice::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQCanBusDeviceFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QCanBusDevice) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBusDevice::event", f)
	}
}

func (ptr *QCanBusDevice) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBusDevice::event")
	}
}

func (ptr *QCanBusDevice) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QCanBusDevice_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QCanBusDevice) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QCanBusDevice_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQCanBusDevice_EventFilter
func callbackQCanBusDevice_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCanBusDevice::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQCanBusDeviceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QCanBusDevice) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBusDevice::eventFilter", f)
	}
}

func (ptr *QCanBusDevice) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBusDevice::eventFilter")
	}
}

func (ptr *QCanBusDevice) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QCanBusDevice_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QCanBusDevice) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QCanBusDevice_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQCanBusDevice_MetaObject
func callbackQCanBusDevice_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCanBusDevice::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQCanBusDeviceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QCanBusDevice) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBusDevice::metaObject", f)
	}
}

func (ptr *QCanBusDevice) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBusDevice::metaObject")
	}
}

func (ptr *QCanBusDevice) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QCanBusDevice_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QCanBusDevice) MetaObjectDefault() *core.QMetaObject {
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

func (ptr *QCanBusFactory) QCanBusFactory_PTR() *QCanBusFactory {
	return ptr
}

func (ptr *QCanBusFactory) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QCanBusFactory) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
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

func (ptr *QCanBusFactory) DestroyQCanBusFactory() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQCanBusFactory_CreateDevice
func callbackQCanBusFactory_CreateDevice(ptr unsafe.Pointer, interfaceName C.struct_QtSerialBus_PackedString) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QCanBusFactory::createDevice"); signal != nil {
		return PointerFromQCanBusDevice(signal.(func(string) *QCanBusDevice)(cGoUnpackString(interfaceName)))
	}

	return PointerFromQCanBusDevice(NewQCanBusDevice(nil))
}

func (ptr *QCanBusFactory) ConnectCreateDevice(f func(interfaceName string) *QCanBusDevice) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBusFactory::createDevice", f)
	}
}

func (ptr *QCanBusFactory) DisconnectCreateDevice() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QCanBusFactory::createDevice")
	}
}

func (ptr *QCanBusFactory) CreateDevice(interfaceName string) *QCanBusDevice {
	if ptr.Pointer() != nil {
		var interfaceNameC = C.CString(interfaceName)
		defer C.free(unsafe.Pointer(interfaceNameC))
		var tmpValue = NewQCanBusDeviceFromPointer(C.QCanBusFactory_CreateDevice(ptr.Pointer(), interfaceNameC))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

type QCanBusFrame struct {
	ptr unsafe.Pointer
}

type QCanBusFrame_ITF interface {
	QCanBusFrame_PTR() *QCanBusFrame
}

func (ptr *QCanBusFrame) QCanBusFrame_PTR() *QCanBusFrame {
	return ptr
}

func (ptr *QCanBusFrame) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QCanBusFrame) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
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

func (ptr *QCanBusFrame) DestroyQCanBusFrame() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QCanBusFrame__FrameError
//QCanBusFrame::FrameError
type QCanBusFrame__FrameError int64

var (
	QCanBusFrame__NoError                    QCanBusFrame__FrameError = QCanBusFrame__FrameError(0)
	QCanBusFrame__TransmissionTimeoutError   QCanBusFrame__FrameError = QCanBusFrame__FrameError(C.QCanBusFrame_TransmissionTimeoutError_Type())
	QCanBusFrame__LostArbitrationError       QCanBusFrame__FrameError = QCanBusFrame__FrameError(C.QCanBusFrame_LostArbitrationError_Type())
	QCanBusFrame__ControllerError            QCanBusFrame__FrameError = QCanBusFrame__FrameError(C.QCanBusFrame_ControllerError_Type())
	QCanBusFrame__ProtocolViolationError     QCanBusFrame__FrameError = QCanBusFrame__FrameError(C.QCanBusFrame_ProtocolViolationError_Type())
	QCanBusFrame__TransceiverError           QCanBusFrame__FrameError = QCanBusFrame__FrameError(C.QCanBusFrame_TransceiverError_Type())
	QCanBusFrame__MissingAcknowledgmentError QCanBusFrame__FrameError = QCanBusFrame__FrameError(C.QCanBusFrame_MissingAcknowledgmentError_Type())
	QCanBusFrame__BusOffError                QCanBusFrame__FrameError = QCanBusFrame__FrameError(C.QCanBusFrame_BusOffError_Type())
	QCanBusFrame__BusError                   QCanBusFrame__FrameError = QCanBusFrame__FrameError(C.QCanBusFrame_BusError_Type())
	QCanBusFrame__ControllerRestartError     QCanBusFrame__FrameError = QCanBusFrame__FrameError(C.QCanBusFrame_ControllerRestartError_Type())
	QCanBusFrame__UnknownError               QCanBusFrame__FrameError = QCanBusFrame__FrameError(C.QCanBusFrame_UnknownError_Type())
	QCanBusFrame__AnyError                   QCanBusFrame__FrameError = QCanBusFrame__FrameError(C.QCanBusFrame_AnyError_Type())
)

//go:generate stringer -type=QCanBusFrame__FrameType
//QCanBusFrame::FrameType
type QCanBusFrame__FrameType int64

const (
	QCanBusFrame__UnknownFrame       QCanBusFrame__FrameType = QCanBusFrame__FrameType(0x0)
	QCanBusFrame__DataFrame          QCanBusFrame__FrameType = QCanBusFrame__FrameType(0x1)
	QCanBusFrame__ErrorFrame         QCanBusFrame__FrameType = QCanBusFrame__FrameType(0x2)
	QCanBusFrame__RemoteRequestFrame QCanBusFrame__FrameType = QCanBusFrame__FrameType(0x3)
	QCanBusFrame__InvalidFrame       QCanBusFrame__FrameType = QCanBusFrame__FrameType(0x4)
)

func NewQCanBusFrame(ty QCanBusFrame__FrameType) *QCanBusFrame {
	var tmpValue = NewQCanBusFrameFromPointer(C.QCanBusFrame_NewQCanBusFrame(C.longlong(ty)))
	runtime.SetFinalizer(tmpValue, (*QCanBusFrame).DestroyQCanBusFrame)
	return tmpValue
}

func NewQCanBusFrame2(identifier uint, data core.QByteArray_ITF) *QCanBusFrame {
	var tmpValue = NewQCanBusFrameFromPointer(C.QCanBusFrame_NewQCanBusFrame2(C.uint(uint32(identifier)), core.PointerFromQByteArray(data)))
	runtime.SetFinalizer(tmpValue, (*QCanBusFrame).DestroyQCanBusFrame)
	return tmpValue
}

func (ptr *QCanBusFrame) Error() QCanBusFrame__FrameError {
	if ptr.Pointer() != nil {
		return QCanBusFrame__FrameError(C.QCanBusFrame_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCanBusFrame) FrameId() uint {
	if ptr.Pointer() != nil {
		return uint(uint32(C.QCanBusFrame_FrameId(ptr.Pointer())))
	}
	return 0
}

func (ptr *QCanBusFrame) FrameType() QCanBusFrame__FrameType {
	if ptr.Pointer() != nil {
		return QCanBusFrame__FrameType(C.QCanBusFrame_FrameType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCanBusFrame) HasExtendedFrameFormat() bool {
	if ptr.Pointer() != nil {
		return C.QCanBusFrame_HasExtendedFrameFormat(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCanBusFrame) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QCanBusFrame_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCanBusFrame) Payload() *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QCanBusFrame_Payload(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QCanBusFrame) SetError(error QCanBusFrame__FrameError) {
	if ptr.Pointer() != nil {
		C.QCanBusFrame_SetError(ptr.Pointer(), C.longlong(error))
	}
}

func (ptr *QCanBusFrame) SetExtendedFrameFormat(isExtended bool) {
	if ptr.Pointer() != nil {
		C.QCanBusFrame_SetExtendedFrameFormat(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(isExtended))))
	}
}

func (ptr *QCanBusFrame) SetFrameId(newFrameId uint) {
	if ptr.Pointer() != nil {
		C.QCanBusFrame_SetFrameId(ptr.Pointer(), C.uint(uint32(newFrameId)))
	}
}

func (ptr *QCanBusFrame) SetFrameType(newType QCanBusFrame__FrameType) {
	if ptr.Pointer() != nil {
		C.QCanBusFrame_SetFrameType(ptr.Pointer(), C.longlong(newType))
	}
}

func (ptr *QCanBusFrame) SetPayload(data core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QCanBusFrame_SetPayload(ptr.Pointer(), core.PointerFromQByteArray(data))
	}
}

type QModbusClient struct {
	QModbusDevice
}

type QModbusClient_ITF interface {
	QModbusDevice_ITF
	QModbusClient_PTR() *QModbusClient
}

func (ptr *QModbusClient) QModbusClient_PTR() *QModbusClient {
	return ptr
}

func (ptr *QModbusClient) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QModbusDevice_PTR().Pointer()
	}
	return nil
}

func (ptr *QModbusClient) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QModbusDevice_PTR().SetPointer(p)
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

func (ptr *QModbusClient) DestroyQModbusClient() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QModbusClient) Timeout() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QModbusClient_Timeout(ptr.Pointer())))
	}
	return 0
}

func NewQModbusClient(parent core.QObject_ITF) *QModbusClient {
	var tmpValue = NewQModbusClientFromPointer(C.QModbusClient_NewQModbusClient(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QModbusClient) NumberOfRetries() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QModbusClient_NumberOfRetries(ptr.Pointer())))
	}
	return 0
}

func (ptr *QModbusClient) SendRawRequest(request QModbusRequest_ITF, serverAddress int) *QModbusReply {
	if ptr.Pointer() != nil {
		var tmpValue = NewQModbusReplyFromPointer(C.QModbusClient_SendRawRequest(ptr.Pointer(), PointerFromQModbusRequest(request), C.int(int32(serverAddress))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QModbusClient) SendReadRequest(read QModbusDataUnit_ITF, serverAddress int) *QModbusReply {
	if ptr.Pointer() != nil {
		var tmpValue = NewQModbusReplyFromPointer(C.QModbusClient_SendReadRequest(ptr.Pointer(), PointerFromQModbusDataUnit(read), C.int(int32(serverAddress))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QModbusClient) SendReadWriteRequest(read QModbusDataUnit_ITF, write QModbusDataUnit_ITF, serverAddress int) *QModbusReply {
	if ptr.Pointer() != nil {
		var tmpValue = NewQModbusReplyFromPointer(C.QModbusClient_SendReadWriteRequest(ptr.Pointer(), PointerFromQModbusDataUnit(read), PointerFromQModbusDataUnit(write), C.int(int32(serverAddress))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QModbusClient) SendWriteRequest(write QModbusDataUnit_ITF, serverAddress int) *QModbusReply {
	if ptr.Pointer() != nil {
		var tmpValue = NewQModbusReplyFromPointer(C.QModbusClient_SendWriteRequest(ptr.Pointer(), PointerFromQModbusDataUnit(write), C.int(int32(serverAddress))))
		if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QModbusClient) SetNumberOfRetries(number int) {
	if ptr.Pointer() != nil {
		C.QModbusClient_SetNumberOfRetries(ptr.Pointer(), C.int(int32(number)))
	}
}

func (ptr *QModbusClient) SetTimeout(newTimeout int) {
	if ptr.Pointer() != nil {
		C.QModbusClient_SetTimeout(ptr.Pointer(), C.int(int32(newTimeout)))
	}
}

//export callbackQModbusClient_TimeoutChanged
func callbackQModbusClient_TimeoutChanged(ptr unsafe.Pointer, newTimeout C.int) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusClient::timeoutChanged"); signal != nil {
		signal.(func(int))(int(int32(newTimeout)))
	}

}

func (ptr *QModbusClient) ConnectTimeoutChanged(f func(newTimeout int)) {
	if ptr.Pointer() != nil {
		C.QModbusClient_ConnectTimeoutChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusClient::timeoutChanged", f)
	}
}

func (ptr *QModbusClient) DisconnectTimeoutChanged() {
	if ptr.Pointer() != nil {
		C.QModbusClient_DisconnectTimeoutChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusClient::timeoutChanged")
	}
}

func (ptr *QModbusClient) TimeoutChanged(newTimeout int) {
	if ptr.Pointer() != nil {
		C.QModbusClient_TimeoutChanged(ptr.Pointer(), C.int(int32(newTimeout)))
	}
}

//export callbackQModbusClient_ProcessPrivateResponse
func callbackQModbusClient_ProcessPrivateResponse(ptr unsafe.Pointer, response unsafe.Pointer, data unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusClient::processPrivateResponse"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QModbusResponse, *QModbusDataUnit) bool)(NewQModbusResponseFromPointer(response), NewQModbusDataUnitFromPointer(data)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusClientFromPointer(ptr).ProcessPrivateResponseDefault(NewQModbusResponseFromPointer(response), NewQModbusDataUnitFromPointer(data)))))
}

func (ptr *QModbusClient) ConnectProcessPrivateResponse(f func(response *QModbusResponse, data *QModbusDataUnit) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusClient::processPrivateResponse", f)
	}
}

func (ptr *QModbusClient) DisconnectProcessPrivateResponse() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusClient::processPrivateResponse")
	}
}

func (ptr *QModbusClient) ProcessPrivateResponse(response QModbusResponse_ITF, data QModbusDataUnit_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusClient_ProcessPrivateResponse(ptr.Pointer(), PointerFromQModbusResponse(response), PointerFromQModbusDataUnit(data)) != 0
	}
	return false
}

func (ptr *QModbusClient) ProcessPrivateResponseDefault(response QModbusResponse_ITF, data QModbusDataUnit_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusClient_ProcessPrivateResponseDefault(ptr.Pointer(), PointerFromQModbusResponse(response), PointerFromQModbusDataUnit(data)) != 0
	}
	return false
}

//export callbackQModbusClient_ProcessResponse
func callbackQModbusClient_ProcessResponse(ptr unsafe.Pointer, response unsafe.Pointer, data unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusClient::processResponse"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QModbusResponse, *QModbusDataUnit) bool)(NewQModbusResponseFromPointer(response), NewQModbusDataUnitFromPointer(data)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusClientFromPointer(ptr).ProcessResponseDefault(NewQModbusResponseFromPointer(response), NewQModbusDataUnitFromPointer(data)))))
}

func (ptr *QModbusClient) ConnectProcessResponse(f func(response *QModbusResponse, data *QModbusDataUnit) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusClient::processResponse", f)
	}
}

func (ptr *QModbusClient) DisconnectProcessResponse() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusClient::processResponse")
	}
}

func (ptr *QModbusClient) ProcessResponse(response QModbusResponse_ITF, data QModbusDataUnit_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusClient_ProcessResponse(ptr.Pointer(), PointerFromQModbusResponse(response), PointerFromQModbusDataUnit(data)) != 0
	}
	return false
}

func (ptr *QModbusClient) ProcessResponseDefault(response QModbusResponse_ITF, data QModbusDataUnit_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusClient_ProcessResponseDefault(ptr.Pointer(), PointerFromQModbusResponse(response), PointerFromQModbusDataUnit(data)) != 0
	}
	return false
}

//export callbackQModbusClient_Close
func callbackQModbusClient_Close(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusClient::close"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QModbusClient) ConnectClose(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusClient::close", f)
	}
}

func (ptr *QModbusClient) DisconnectClose() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusClient::close")
	}
}

func (ptr *QModbusClient) Close() {
	if ptr.Pointer() != nil {
		C.QModbusClient_Close(ptr.Pointer())
	}
}

//export callbackQModbusClient_Open
func callbackQModbusClient_Open(ptr unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusClient::open"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QModbusClient) ConnectOpen(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusClient::open", f)
	}
}

func (ptr *QModbusClient) DisconnectOpen() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusClient::open")
	}
}

func (ptr *QModbusClient) Open() bool {
	if ptr.Pointer() != nil {
		return C.QModbusClient_Open(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQModbusClient_TimerEvent
func callbackQModbusClient_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusClient::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQModbusClientFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QModbusClient) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusClient::timerEvent", f)
	}
}

func (ptr *QModbusClient) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusClient::timerEvent")
	}
}

func (ptr *QModbusClient) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusClient_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QModbusClient) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusClient_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQModbusClient_ChildEvent
func callbackQModbusClient_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusClient::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQModbusClientFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QModbusClient) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusClient::childEvent", f)
	}
}

func (ptr *QModbusClient) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusClient::childEvent")
	}
}

func (ptr *QModbusClient) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusClient_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QModbusClient) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusClient_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQModbusClient_ConnectNotify
func callbackQModbusClient_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusClient::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQModbusClientFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QModbusClient) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusClient::connectNotify", f)
	}
}

func (ptr *QModbusClient) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusClient::connectNotify")
	}
}

func (ptr *QModbusClient) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusClient_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QModbusClient) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusClient_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQModbusClient_CustomEvent
func callbackQModbusClient_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusClient::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQModbusClientFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QModbusClient) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusClient::customEvent", f)
	}
}

func (ptr *QModbusClient) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusClient::customEvent")
	}
}

func (ptr *QModbusClient) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusClient_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QModbusClient) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusClient_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQModbusClient_DeleteLater
func callbackQModbusClient_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusClient::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQModbusClientFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QModbusClient) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusClient::deleteLater", f)
	}
}

func (ptr *QModbusClient) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusClient::deleteLater")
	}
}

func (ptr *QModbusClient) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QModbusClient_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QModbusClient) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QModbusClient_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQModbusClient_DisconnectNotify
func callbackQModbusClient_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusClient::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQModbusClientFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QModbusClient) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusClient::disconnectNotify", f)
	}
}

func (ptr *QModbusClient) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusClient::disconnectNotify")
	}
}

func (ptr *QModbusClient) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusClient_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QModbusClient) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusClient_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQModbusClient_Event
func callbackQModbusClient_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusClient::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusClientFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QModbusClient) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusClient::event", f)
	}
}

func (ptr *QModbusClient) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusClient::event")
	}
}

func (ptr *QModbusClient) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusClient_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QModbusClient) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusClient_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQModbusClient_EventFilter
func callbackQModbusClient_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusClient::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusClientFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QModbusClient) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusClient::eventFilter", f)
	}
}

func (ptr *QModbusClient) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusClient::eventFilter")
	}
}

func (ptr *QModbusClient) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusClient_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QModbusClient) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusClient_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQModbusClient_MetaObject
func callbackQModbusClient_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusClient::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQModbusClientFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QModbusClient) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusClient::metaObject", f)
	}
}

func (ptr *QModbusClient) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusClient::metaObject")
	}
}

func (ptr *QModbusClient) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QModbusClient_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QModbusClient) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QModbusClient_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QModbusDataUnit struct {
	ptr unsafe.Pointer
}

type QModbusDataUnit_ITF interface {
	QModbusDataUnit_PTR() *QModbusDataUnit
}

func (ptr *QModbusDataUnit) QModbusDataUnit_PTR() *QModbusDataUnit {
	return ptr
}

func (ptr *QModbusDataUnit) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QModbusDataUnit) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
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

func (ptr *QModbusDataUnit) DestroyQModbusDataUnit() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QModbusDataUnit__RegisterType
//QModbusDataUnit::RegisterType
type QModbusDataUnit__RegisterType int64

const (
	QModbusDataUnit__Invalid          QModbusDataUnit__RegisterType = QModbusDataUnit__RegisterType(0)
	QModbusDataUnit__DiscreteInputs   QModbusDataUnit__RegisterType = QModbusDataUnit__RegisterType(1)
	QModbusDataUnit__Coils            QModbusDataUnit__RegisterType = QModbusDataUnit__RegisterType(2)
	QModbusDataUnit__InputRegisters   QModbusDataUnit__RegisterType = QModbusDataUnit__RegisterType(3)
	QModbusDataUnit__HoldingRegisters QModbusDataUnit__RegisterType = QModbusDataUnit__RegisterType(4)
)

func NewQModbusDataUnit() *QModbusDataUnit {
	var tmpValue = NewQModbusDataUnitFromPointer(C.QModbusDataUnit_NewQModbusDataUnit())
	runtime.SetFinalizer(tmpValue, (*QModbusDataUnit).DestroyQModbusDataUnit)
	return tmpValue
}

func NewQModbusDataUnit2(ty QModbusDataUnit__RegisterType) *QModbusDataUnit {
	var tmpValue = NewQModbusDataUnitFromPointer(C.QModbusDataUnit_NewQModbusDataUnit2(C.longlong(ty)))
	runtime.SetFinalizer(tmpValue, (*QModbusDataUnit).DestroyQModbusDataUnit)
	return tmpValue
}

func NewQModbusDataUnit3(ty QModbusDataUnit__RegisterType, address int, size uint16) *QModbusDataUnit {
	var tmpValue = NewQModbusDataUnitFromPointer(C.QModbusDataUnit_NewQModbusDataUnit3(C.longlong(ty), C.int(int32(address)), C.ushort(size)))
	runtime.SetFinalizer(tmpValue, (*QModbusDataUnit).DestroyQModbusDataUnit)
	return tmpValue
}

func (ptr *QModbusDataUnit) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QModbusDataUnit_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QModbusDataUnit) RegisterType() QModbusDataUnit__RegisterType {
	if ptr.Pointer() != nil {
		return QModbusDataUnit__RegisterType(C.QModbusDataUnit_RegisterType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QModbusDataUnit) SetRegisterType(ty QModbusDataUnit__RegisterType) {
	if ptr.Pointer() != nil {
		C.QModbusDataUnit_SetRegisterType(ptr.Pointer(), C.longlong(ty))
	}
}

func (ptr *QModbusDataUnit) SetStartAddress(address int) {
	if ptr.Pointer() != nil {
		C.QModbusDataUnit_SetStartAddress(ptr.Pointer(), C.int(int32(address)))
	}
}

func (ptr *QModbusDataUnit) SetValue(index int, value uint16) {
	if ptr.Pointer() != nil {
		C.QModbusDataUnit_SetValue(ptr.Pointer(), C.int(int32(index)), C.ushort(value))
	}
}

func (ptr *QModbusDataUnit) SetValueCount(newCount uint) {
	if ptr.Pointer() != nil {
		C.QModbusDataUnit_SetValueCount(ptr.Pointer(), C.uint(uint32(newCount)))
	}
}

func (ptr *QModbusDataUnit) StartAddress() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QModbusDataUnit_StartAddress(ptr.Pointer())))
	}
	return 0
}

func (ptr *QModbusDataUnit) Value(index int) uint16 {
	if ptr.Pointer() != nil {
		return uint16(C.QModbusDataUnit_Value(ptr.Pointer(), C.int(int32(index))))
	}
	return 0
}

func (ptr *QModbusDataUnit) ValueCount() uint {
	if ptr.Pointer() != nil {
		return uint(uint32(C.QModbusDataUnit_ValueCount(ptr.Pointer())))
	}
	return 0
}

type QModbusDevice struct {
	core.QObject
}

type QModbusDevice_ITF interface {
	core.QObject_ITF
	QModbusDevice_PTR() *QModbusDevice
}

func (ptr *QModbusDevice) QModbusDevice_PTR() *QModbusDevice {
	return ptr
}

func (ptr *QModbusDevice) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QModbusDevice) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
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

//go:generate stringer -type=QModbusDevice__ConnectionParameter
//QModbusDevice::ConnectionParameter
type QModbusDevice__ConnectionParameter int64

const (
	QModbusDevice__SerialPortNameParameter QModbusDevice__ConnectionParameter = QModbusDevice__ConnectionParameter(0)
	QModbusDevice__SerialParityParameter   QModbusDevice__ConnectionParameter = QModbusDevice__ConnectionParameter(1)
	QModbusDevice__SerialBaudRateParameter QModbusDevice__ConnectionParameter = QModbusDevice__ConnectionParameter(2)
	QModbusDevice__SerialDataBitsParameter QModbusDevice__ConnectionParameter = QModbusDevice__ConnectionParameter(3)
	QModbusDevice__SerialStopBitsParameter QModbusDevice__ConnectionParameter = QModbusDevice__ConnectionParameter(4)
	QModbusDevice__NetworkPortParameter    QModbusDevice__ConnectionParameter = QModbusDevice__ConnectionParameter(5)
	QModbusDevice__NetworkAddressParameter QModbusDevice__ConnectionParameter = QModbusDevice__ConnectionParameter(6)
	QModbusDevice__UserParameter           QModbusDevice__ConnectionParameter = QModbusDevice__ConnectionParameter(0x100)
)

//go:generate stringer -type=QModbusDevice__Error
//QModbusDevice::Error
type QModbusDevice__Error int64

const (
	QModbusDevice__NoError            QModbusDevice__Error = QModbusDevice__Error(0)
	QModbusDevice__ReadError          QModbusDevice__Error = QModbusDevice__Error(1)
	QModbusDevice__WriteError         QModbusDevice__Error = QModbusDevice__Error(2)
	QModbusDevice__ConnectionError    QModbusDevice__Error = QModbusDevice__Error(3)
	QModbusDevice__ConfigurationError QModbusDevice__Error = QModbusDevice__Error(4)
	QModbusDevice__TimeoutError       QModbusDevice__Error = QModbusDevice__Error(5)
	QModbusDevice__ProtocolError      QModbusDevice__Error = QModbusDevice__Error(6)
	QModbusDevice__ReplyAbortedError  QModbusDevice__Error = QModbusDevice__Error(7)
	QModbusDevice__UnknownError       QModbusDevice__Error = QModbusDevice__Error(8)
)

//go:generate stringer -type=QModbusDevice__State
//QModbusDevice::State
type QModbusDevice__State int64

const (
	QModbusDevice__UnconnectedState QModbusDevice__State = QModbusDevice__State(0)
	QModbusDevice__ConnectingState  QModbusDevice__State = QModbusDevice__State(1)
	QModbusDevice__ConnectedState   QModbusDevice__State = QModbusDevice__State(2)
	QModbusDevice__ClosingState     QModbusDevice__State = QModbusDevice__State(3)
)

func NewQModbusDevice(parent core.QObject_ITF) *QModbusDevice {
	var tmpValue = NewQModbusDeviceFromPointer(C.QModbusDevice_NewQModbusDevice(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQModbusDevice_Close
func callbackQModbusDevice_Close(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusDevice::close"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QModbusDevice) ConnectClose(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusDevice::close", f)
	}
}

func (ptr *QModbusDevice) DisconnectClose() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusDevice::close")
	}
}

func (ptr *QModbusDevice) Close() {
	if ptr.Pointer() != nil {
		C.QModbusDevice_Close(ptr.Pointer())
	}
}

func (ptr *QModbusDevice) ConnectDevice() bool {
	if ptr.Pointer() != nil {
		return C.QModbusDevice_ConnectDevice(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QModbusDevice) ConnectionParameter(parameter int) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QModbusDevice_ConnectionParameter(ptr.Pointer(), C.int(int32(parameter))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QModbusDevice) DisconnectDevice() {
	if ptr.Pointer() != nil {
		C.QModbusDevice_DisconnectDevice(ptr.Pointer())
	}
}

func (ptr *QModbusDevice) Error() QModbusDevice__Error {
	if ptr.Pointer() != nil {
		return QModbusDevice__Error(C.QModbusDevice_Error(ptr.Pointer()))
	}
	return 0
}

//export callbackQModbusDevice_ErrorOccurred
func callbackQModbusDevice_ErrorOccurred(ptr unsafe.Pointer, error C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusDevice::errorOccurred"); signal != nil {
		signal.(func(QModbusDevice__Error))(QModbusDevice__Error(error))
	}

}

func (ptr *QModbusDevice) ConnectErrorOccurred(f func(error QModbusDevice__Error)) {
	if ptr.Pointer() != nil {
		C.QModbusDevice_ConnectErrorOccurred(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusDevice::errorOccurred", f)
	}
}

func (ptr *QModbusDevice) DisconnectErrorOccurred() {
	if ptr.Pointer() != nil {
		C.QModbusDevice_DisconnectErrorOccurred(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusDevice::errorOccurred")
	}
}

func (ptr *QModbusDevice) ErrorOccurred(error QModbusDevice__Error) {
	if ptr.Pointer() != nil {
		C.QModbusDevice_ErrorOccurred(ptr.Pointer(), C.longlong(error))
	}
}

func (ptr *QModbusDevice) ErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QModbusDevice_ErrorString(ptr.Pointer()))
	}
	return ""
}

//export callbackQModbusDevice_Open
func callbackQModbusDevice_Open(ptr unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusDevice::open"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QModbusDevice) ConnectOpen(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusDevice::open", f)
	}
}

func (ptr *QModbusDevice) DisconnectOpen() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusDevice::open")
	}
}

func (ptr *QModbusDevice) Open() bool {
	if ptr.Pointer() != nil {
		return C.QModbusDevice_Open(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QModbusDevice) SetConnectionParameter(parameter int, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusDevice_SetConnectionParameter(ptr.Pointer(), C.int(int32(parameter)), core.PointerFromQVariant(value))
	}
}

func (ptr *QModbusDevice) SetError(errorText string, error QModbusDevice__Error) {
	if ptr.Pointer() != nil {
		var errorTextC = C.CString(errorText)
		defer C.free(unsafe.Pointer(errorTextC))
		C.QModbusDevice_SetError(ptr.Pointer(), errorTextC, C.longlong(error))
	}
}

func (ptr *QModbusDevice) SetState(newState QModbusDevice__State) {
	if ptr.Pointer() != nil {
		C.QModbusDevice_SetState(ptr.Pointer(), C.longlong(newState))
	}
}

func (ptr *QModbusDevice) State() QModbusDevice__State {
	if ptr.Pointer() != nil {
		return QModbusDevice__State(C.QModbusDevice_State(ptr.Pointer()))
	}
	return 0
}

//export callbackQModbusDevice_StateChanged
func callbackQModbusDevice_StateChanged(ptr unsafe.Pointer, state C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusDevice::stateChanged"); signal != nil {
		signal.(func(QModbusDevice__State))(QModbusDevice__State(state))
	}

}

func (ptr *QModbusDevice) ConnectStateChanged(f func(state QModbusDevice__State)) {
	if ptr.Pointer() != nil {
		C.QModbusDevice_ConnectStateChanged(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusDevice::stateChanged", f)
	}
}

func (ptr *QModbusDevice) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QModbusDevice_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusDevice::stateChanged")
	}
}

func (ptr *QModbusDevice) StateChanged(state QModbusDevice__State) {
	if ptr.Pointer() != nil {
		C.QModbusDevice_StateChanged(ptr.Pointer(), C.longlong(state))
	}
}

func (ptr *QModbusDevice) DestroyQModbusDevice() {
	if ptr.Pointer() != nil {
		C.QModbusDevice_DestroyQModbusDevice(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQModbusDevice_TimerEvent
func callbackQModbusDevice_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusDevice::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQModbusDeviceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QModbusDevice) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusDevice::timerEvent", f)
	}
}

func (ptr *QModbusDevice) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusDevice::timerEvent")
	}
}

func (ptr *QModbusDevice) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusDevice_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QModbusDevice) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusDevice_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQModbusDevice_ChildEvent
func callbackQModbusDevice_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusDevice::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQModbusDeviceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QModbusDevice) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusDevice::childEvent", f)
	}
}

func (ptr *QModbusDevice) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusDevice::childEvent")
	}
}

func (ptr *QModbusDevice) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusDevice_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QModbusDevice) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusDevice_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQModbusDevice_ConnectNotify
func callbackQModbusDevice_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusDevice::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQModbusDeviceFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QModbusDevice) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusDevice::connectNotify", f)
	}
}

func (ptr *QModbusDevice) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusDevice::connectNotify")
	}
}

func (ptr *QModbusDevice) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusDevice_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QModbusDevice) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusDevice_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQModbusDevice_CustomEvent
func callbackQModbusDevice_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusDevice::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQModbusDeviceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QModbusDevice) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusDevice::customEvent", f)
	}
}

func (ptr *QModbusDevice) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusDevice::customEvent")
	}
}

func (ptr *QModbusDevice) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusDevice_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QModbusDevice) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusDevice_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQModbusDevice_DeleteLater
func callbackQModbusDevice_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusDevice::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQModbusDeviceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QModbusDevice) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusDevice::deleteLater", f)
	}
}

func (ptr *QModbusDevice) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusDevice::deleteLater")
	}
}

func (ptr *QModbusDevice) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QModbusDevice_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QModbusDevice) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QModbusDevice_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQModbusDevice_DisconnectNotify
func callbackQModbusDevice_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusDevice::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQModbusDeviceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QModbusDevice) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusDevice::disconnectNotify", f)
	}
}

func (ptr *QModbusDevice) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusDevice::disconnectNotify")
	}
}

func (ptr *QModbusDevice) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusDevice_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QModbusDevice) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusDevice_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQModbusDevice_Event
func callbackQModbusDevice_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusDevice::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusDeviceFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QModbusDevice) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusDevice::event", f)
	}
}

func (ptr *QModbusDevice) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusDevice::event")
	}
}

func (ptr *QModbusDevice) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusDevice_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QModbusDevice) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusDevice_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQModbusDevice_EventFilter
func callbackQModbusDevice_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusDevice::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusDeviceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QModbusDevice) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusDevice::eventFilter", f)
	}
}

func (ptr *QModbusDevice) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusDevice::eventFilter")
	}
}

func (ptr *QModbusDevice) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusDevice_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QModbusDevice) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusDevice_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQModbusDevice_MetaObject
func callbackQModbusDevice_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusDevice::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQModbusDeviceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QModbusDevice) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusDevice::metaObject", f)
	}
}

func (ptr *QModbusDevice) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusDevice::metaObject")
	}
}

func (ptr *QModbusDevice) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QModbusDevice_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QModbusDevice) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QModbusDevice_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QModbusDeviceIdentification struct {
	ptr unsafe.Pointer
}

type QModbusDeviceIdentification_ITF interface {
	QModbusDeviceIdentification_PTR() *QModbusDeviceIdentification
}

func (ptr *QModbusDeviceIdentification) QModbusDeviceIdentification_PTR() *QModbusDeviceIdentification {
	return ptr
}

func (ptr *QModbusDeviceIdentification) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QModbusDeviceIdentification) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
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

func (ptr *QModbusDeviceIdentification) DestroyQModbusDeviceIdentification() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QModbusDeviceIdentification__ConformityLevel
//QModbusDeviceIdentification::ConformityLevel
type QModbusDeviceIdentification__ConformityLevel int64

const (
	QModbusDeviceIdentification__BasicConformityLevel              QModbusDeviceIdentification__ConformityLevel = QModbusDeviceIdentification__ConformityLevel(0x01)
	QModbusDeviceIdentification__RegularConformityLevel            QModbusDeviceIdentification__ConformityLevel = QModbusDeviceIdentification__ConformityLevel(0x02)
	QModbusDeviceIdentification__ExtendedConformityLevel           QModbusDeviceIdentification__ConformityLevel = QModbusDeviceIdentification__ConformityLevel(0x03)
	QModbusDeviceIdentification__BasicIndividualConformityLevel    QModbusDeviceIdentification__ConformityLevel = QModbusDeviceIdentification__ConformityLevel(0x81)
	QModbusDeviceIdentification__RegularIndividualConformityLevel  QModbusDeviceIdentification__ConformityLevel = QModbusDeviceIdentification__ConformityLevel(0x82)
	QModbusDeviceIdentification__ExtendedIndividualConformityLevel QModbusDeviceIdentification__ConformityLevel = QModbusDeviceIdentification__ConformityLevel(0x83)
)

//go:generate stringer -type=QModbusDeviceIdentification__ObjectId
//QModbusDeviceIdentification::ObjectId
type QModbusDeviceIdentification__ObjectId int64

const (
	QModbusDeviceIdentification__VendorNameObjectId          QModbusDeviceIdentification__ObjectId = QModbusDeviceIdentification__ObjectId(0x00)
	QModbusDeviceIdentification__ProductCodeObjectId         QModbusDeviceIdentification__ObjectId = QModbusDeviceIdentification__ObjectId(0x01)
	QModbusDeviceIdentification__MajorMinorRevisionObjectId  QModbusDeviceIdentification__ObjectId = QModbusDeviceIdentification__ObjectId(0x02)
	QModbusDeviceIdentification__VendorUrlObjectId           QModbusDeviceIdentification__ObjectId = QModbusDeviceIdentification__ObjectId(0x03)
	QModbusDeviceIdentification__ProductNameObjectId         QModbusDeviceIdentification__ObjectId = QModbusDeviceIdentification__ObjectId(0x04)
	QModbusDeviceIdentification__ModelNameObjectId           QModbusDeviceIdentification__ObjectId = QModbusDeviceIdentification__ObjectId(0x05)
	QModbusDeviceIdentification__UserApplicationNameObjectId QModbusDeviceIdentification__ObjectId = QModbusDeviceIdentification__ObjectId(0x06)
	QModbusDeviceIdentification__ReservedObjectId            QModbusDeviceIdentification__ObjectId = QModbusDeviceIdentification__ObjectId(0x07)
	QModbusDeviceIdentification__ProductDependentObjectId    QModbusDeviceIdentification__ObjectId = QModbusDeviceIdentification__ObjectId(0x80)
	QModbusDeviceIdentification__UndefinedObjectId           QModbusDeviceIdentification__ObjectId = QModbusDeviceIdentification__ObjectId(0x100)
)

//go:generate stringer -type=QModbusDeviceIdentification__ReadDeviceIdCode
//QModbusDeviceIdentification::ReadDeviceIdCode
type QModbusDeviceIdentification__ReadDeviceIdCode int64

const (
	QModbusDeviceIdentification__BasicReadDeviceIdCode      QModbusDeviceIdentification__ReadDeviceIdCode = QModbusDeviceIdentification__ReadDeviceIdCode(0x01)
	QModbusDeviceIdentification__RegularReadDeviceIdCode    QModbusDeviceIdentification__ReadDeviceIdCode = QModbusDeviceIdentification__ReadDeviceIdCode(0x02)
	QModbusDeviceIdentification__ExtendedReadDeviceIdCode   QModbusDeviceIdentification__ReadDeviceIdCode = QModbusDeviceIdentification__ReadDeviceIdCode(0x03)
	QModbusDeviceIdentification__IndividualReadDeviceIdCode QModbusDeviceIdentification__ReadDeviceIdCode = QModbusDeviceIdentification__ReadDeviceIdCode(0x04)
)

func NewQModbusDeviceIdentification() *QModbusDeviceIdentification {
	var tmpValue = NewQModbusDeviceIdentificationFromPointer(C.QModbusDeviceIdentification_NewQModbusDeviceIdentification())
	runtime.SetFinalizer(tmpValue, (*QModbusDeviceIdentification).DestroyQModbusDeviceIdentification)
	return tmpValue
}

func (ptr *QModbusDeviceIdentification) ConformityLevel() QModbusDeviceIdentification__ConformityLevel {
	if ptr.Pointer() != nil {
		return QModbusDeviceIdentification__ConformityLevel(C.QModbusDeviceIdentification_ConformityLevel(ptr.Pointer()))
	}
	return 0
}

func (ptr *QModbusDeviceIdentification) Contains(objectId uint) bool {
	if ptr.Pointer() != nil {
		return C.QModbusDeviceIdentification_Contains(ptr.Pointer(), C.uint(uint32(objectId))) != 0
	}
	return false
}

func QModbusDeviceIdentification_FromByteArray(ba core.QByteArray_ITF) *QModbusDeviceIdentification {
	var tmpValue = NewQModbusDeviceIdentificationFromPointer(C.QModbusDeviceIdentification_QModbusDeviceIdentification_FromByteArray(core.PointerFromQByteArray(ba)))
	runtime.SetFinalizer(tmpValue, (*QModbusDeviceIdentification).DestroyQModbusDeviceIdentification)
	return tmpValue
}

func (ptr *QModbusDeviceIdentification) FromByteArray(ba core.QByteArray_ITF) *QModbusDeviceIdentification {
	var tmpValue = NewQModbusDeviceIdentificationFromPointer(C.QModbusDeviceIdentification_QModbusDeviceIdentification_FromByteArray(core.PointerFromQByteArray(ba)))
	runtime.SetFinalizer(tmpValue, (*QModbusDeviceIdentification).DestroyQModbusDeviceIdentification)
	return tmpValue
}

func (ptr *QModbusDeviceIdentification) Insert(objectId uint, value core.QByteArray_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusDeviceIdentification_Insert(ptr.Pointer(), C.uint(uint32(objectId)), core.PointerFromQByteArray(value)) != 0
	}
	return false
}

func (ptr *QModbusDeviceIdentification) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QModbusDeviceIdentification_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QModbusDeviceIdentification) Remove(objectId uint) {
	if ptr.Pointer() != nil {
		C.QModbusDeviceIdentification_Remove(ptr.Pointer(), C.uint(uint32(objectId)))
	}
}

func (ptr *QModbusDeviceIdentification) SetConformityLevel(level QModbusDeviceIdentification__ConformityLevel) {
	if ptr.Pointer() != nil {
		C.QModbusDeviceIdentification_SetConformityLevel(ptr.Pointer(), C.longlong(level))
	}
}

func (ptr *QModbusDeviceIdentification) Value(objectId uint) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QModbusDeviceIdentification_Value(ptr.Pointer(), C.uint(uint32(objectId))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

type QModbusExceptionResponse struct {
	QModbusResponse
}

type QModbusExceptionResponse_ITF interface {
	QModbusResponse_ITF
	QModbusExceptionResponse_PTR() *QModbusExceptionResponse
}

func (ptr *QModbusExceptionResponse) QModbusExceptionResponse_PTR() *QModbusExceptionResponse {
	return ptr
}

func (ptr *QModbusExceptionResponse) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QModbusResponse_PTR().Pointer()
	}
	return nil
}

func (ptr *QModbusExceptionResponse) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QModbusResponse_PTR().SetPointer(p)
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

func (ptr *QModbusExceptionResponse) DestroyQModbusExceptionResponse() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func NewQModbusExceptionResponse() *QModbusExceptionResponse {
	return NewQModbusExceptionResponseFromPointer(C.QModbusExceptionResponse_NewQModbusExceptionResponse())
}

func NewQModbusExceptionResponse3(code QModbusPdu__FunctionCode, ec QModbusPdu__ExceptionCode) *QModbusExceptionResponse {
	return NewQModbusExceptionResponseFromPointer(C.QModbusExceptionResponse_NewQModbusExceptionResponse3(C.longlong(code), C.longlong(ec)))
}

func NewQModbusExceptionResponse2(pdu QModbusPdu_ITF) *QModbusExceptionResponse {
	return NewQModbusExceptionResponseFromPointer(C.QModbusExceptionResponse_NewQModbusExceptionResponse2(PointerFromQModbusPdu(pdu)))
}

func (ptr *QModbusExceptionResponse) SetExceptionCode(ec QModbusPdu__ExceptionCode) {
	if ptr.Pointer() != nil {
		C.QModbusExceptionResponse_SetExceptionCode(ptr.Pointer(), C.longlong(ec))
	}
}

//export callbackQModbusExceptionResponse_SetFunctionCode
func callbackQModbusExceptionResponse_SetFunctionCode(ptr unsafe.Pointer, c C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusExceptionResponse::setFunctionCode"); signal != nil {
		signal.(func(QModbusPdu__FunctionCode))(QModbusPdu__FunctionCode(c))
	} else {
		NewQModbusExceptionResponseFromPointer(ptr).SetFunctionCodeDefault(QModbusPdu__FunctionCode(c))
	}
}

func (ptr *QModbusExceptionResponse) ConnectSetFunctionCode(f func(c QModbusPdu__FunctionCode)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusExceptionResponse::setFunctionCode", f)
	}
}

func (ptr *QModbusExceptionResponse) DisconnectSetFunctionCode() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusExceptionResponse::setFunctionCode")
	}
}

func (ptr *QModbusExceptionResponse) SetFunctionCode(c QModbusPdu__FunctionCode) {
	if ptr.Pointer() != nil {
		C.QModbusExceptionResponse_SetFunctionCode(ptr.Pointer(), C.longlong(c))
	}
}

func (ptr *QModbusExceptionResponse) SetFunctionCodeDefault(c QModbusPdu__FunctionCode) {
	if ptr.Pointer() != nil {
		C.QModbusExceptionResponse_SetFunctionCodeDefault(ptr.Pointer(), C.longlong(c))
	}
}

type QModbusPdu struct {
	ptr unsafe.Pointer
}

type QModbusPdu_ITF interface {
	QModbusPdu_PTR() *QModbusPdu
}

func (ptr *QModbusPdu) QModbusPdu_PTR() *QModbusPdu {
	return ptr
}

func (ptr *QModbusPdu) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QModbusPdu) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
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

//go:generate stringer -type=QModbusPdu__ExceptionCode
//QModbusPdu::ExceptionCode
type QModbusPdu__ExceptionCode int64

const (
	QModbusPdu__IllegalFunction                    QModbusPdu__ExceptionCode = QModbusPdu__ExceptionCode(0x01)
	QModbusPdu__IllegalDataAddress                 QModbusPdu__ExceptionCode = QModbusPdu__ExceptionCode(0x02)
	QModbusPdu__IllegalDataValue                   QModbusPdu__ExceptionCode = QModbusPdu__ExceptionCode(0x03)
	QModbusPdu__ServerDeviceFailure                QModbusPdu__ExceptionCode = QModbusPdu__ExceptionCode(0x04)
	QModbusPdu__Acknowledge                        QModbusPdu__ExceptionCode = QModbusPdu__ExceptionCode(0x05)
	QModbusPdu__ServerDeviceBusy                   QModbusPdu__ExceptionCode = QModbusPdu__ExceptionCode(0x06)
	QModbusPdu__NegativeAcknowledge                QModbusPdu__ExceptionCode = QModbusPdu__ExceptionCode(0x07)
	QModbusPdu__MemoryParityError                  QModbusPdu__ExceptionCode = QModbusPdu__ExceptionCode(0x08)
	QModbusPdu__GatewayPathUnavailable             QModbusPdu__ExceptionCode = QModbusPdu__ExceptionCode(0x0A)
	QModbusPdu__GatewayTargetDeviceFailedToRespond QModbusPdu__ExceptionCode = QModbusPdu__ExceptionCode(0x0B)
	QModbusPdu__ExtendedException                  QModbusPdu__ExceptionCode = QModbusPdu__ExceptionCode(0xFF)
)

//go:generate stringer -type=QModbusPdu__FunctionCode
//QModbusPdu::FunctionCode
type QModbusPdu__FunctionCode int64

const (
	QModbusPdu__Invalid                        QModbusPdu__FunctionCode = QModbusPdu__FunctionCode(0x00)
	QModbusPdu__ReadCoils                      QModbusPdu__FunctionCode = QModbusPdu__FunctionCode(0x01)
	QModbusPdu__ReadDiscreteInputs             QModbusPdu__FunctionCode = QModbusPdu__FunctionCode(0x02)
	QModbusPdu__ReadHoldingRegisters           QModbusPdu__FunctionCode = QModbusPdu__FunctionCode(0x03)
	QModbusPdu__ReadInputRegisters             QModbusPdu__FunctionCode = QModbusPdu__FunctionCode(0x04)
	QModbusPdu__WriteSingleCoil                QModbusPdu__FunctionCode = QModbusPdu__FunctionCode(0x05)
	QModbusPdu__WriteSingleRegister            QModbusPdu__FunctionCode = QModbusPdu__FunctionCode(0x06)
	QModbusPdu__ReadExceptionStatus            QModbusPdu__FunctionCode = QModbusPdu__FunctionCode(0x07)
	QModbusPdu__Diagnostics                    QModbusPdu__FunctionCode = QModbusPdu__FunctionCode(0x08)
	QModbusPdu__GetCommEventCounter            QModbusPdu__FunctionCode = QModbusPdu__FunctionCode(0x0B)
	QModbusPdu__GetCommEventLog                QModbusPdu__FunctionCode = QModbusPdu__FunctionCode(0x0C)
	QModbusPdu__WriteMultipleCoils             QModbusPdu__FunctionCode = QModbusPdu__FunctionCode(0x0F)
	QModbusPdu__WriteMultipleRegisters         QModbusPdu__FunctionCode = QModbusPdu__FunctionCode(0x10)
	QModbusPdu__ReportServerId                 QModbusPdu__FunctionCode = QModbusPdu__FunctionCode(0x11)
	QModbusPdu__ReadFileRecord                 QModbusPdu__FunctionCode = QModbusPdu__FunctionCode(0x14)
	QModbusPdu__WriteFileRecord                QModbusPdu__FunctionCode = QModbusPdu__FunctionCode(0x15)
	QModbusPdu__MaskWriteRegister              QModbusPdu__FunctionCode = QModbusPdu__FunctionCode(0x16)
	QModbusPdu__ReadWriteMultipleRegisters     QModbusPdu__FunctionCode = QModbusPdu__FunctionCode(0x17)
	QModbusPdu__ReadFifoQueue                  QModbusPdu__FunctionCode = QModbusPdu__FunctionCode(0x18)
	QModbusPdu__EncapsulatedInterfaceTransport QModbusPdu__FunctionCode = QModbusPdu__FunctionCode(0x2B)
	QModbusPdu__UndefinedFunctionCode          QModbusPdu__FunctionCode = QModbusPdu__FunctionCode(0x100)
)

func NewQModbusPdu() *QModbusPdu {
	return NewQModbusPduFromPointer(C.QModbusPdu_NewQModbusPdu())
}

func NewQModbusPdu2(code QModbusPdu__FunctionCode, data core.QByteArray_ITF) *QModbusPdu {
	return NewQModbusPduFromPointer(C.QModbusPdu_NewQModbusPdu2(C.longlong(code), core.PointerFromQByteArray(data)))
}

func NewQModbusPdu3(other QModbusPdu_ITF) *QModbusPdu {
	return NewQModbusPduFromPointer(C.QModbusPdu_NewQModbusPdu3(PointerFromQModbusPdu(other)))
}

func (ptr *QModbusPdu) DataSize() int16 {
	if ptr.Pointer() != nil {
		return int16(C.QModbusPdu_DataSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QModbusPdu) Data() *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QModbusPdu_Data(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QModbusPdu) ExceptionCode() QModbusPdu__ExceptionCode {
	if ptr.Pointer() != nil {
		return QModbusPdu__ExceptionCode(C.QModbusPdu_ExceptionCode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QModbusPdu) FunctionCode() QModbusPdu__FunctionCode {
	if ptr.Pointer() != nil {
		return QModbusPdu__FunctionCode(C.QModbusPdu_FunctionCode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QModbusPdu) IsException() bool {
	if ptr.Pointer() != nil {
		return C.QModbusPdu_IsException(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QModbusPdu) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QModbusPdu_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QModbusPdu) SetData(data core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusPdu_SetData(ptr.Pointer(), core.PointerFromQByteArray(data))
	}
}

//export callbackQModbusPdu_SetFunctionCode
func callbackQModbusPdu_SetFunctionCode(ptr unsafe.Pointer, code C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusPdu::setFunctionCode"); signal != nil {
		signal.(func(QModbusPdu__FunctionCode))(QModbusPdu__FunctionCode(code))
	} else {
		NewQModbusPduFromPointer(ptr).SetFunctionCodeDefault(QModbusPdu__FunctionCode(code))
	}
}

func (ptr *QModbusPdu) ConnectSetFunctionCode(f func(code QModbusPdu__FunctionCode)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusPdu::setFunctionCode", f)
	}
}

func (ptr *QModbusPdu) DisconnectSetFunctionCode() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusPdu::setFunctionCode")
	}
}

func (ptr *QModbusPdu) SetFunctionCode(code QModbusPdu__FunctionCode) {
	if ptr.Pointer() != nil {
		C.QModbusPdu_SetFunctionCode(ptr.Pointer(), C.longlong(code))
	}
}

func (ptr *QModbusPdu) SetFunctionCodeDefault(code QModbusPdu__FunctionCode) {
	if ptr.Pointer() != nil {
		C.QModbusPdu_SetFunctionCodeDefault(ptr.Pointer(), C.longlong(code))
	}
}

func (ptr *QModbusPdu) Size() int16 {
	if ptr.Pointer() != nil {
		return int16(C.QModbusPdu_Size(ptr.Pointer()))
	}
	return 0
}

//export callbackQModbusPdu_DestroyQModbusPdu
func callbackQModbusPdu_DestroyQModbusPdu(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusPdu::~QModbusPdu"); signal != nil {
		signal.(func())()
	} else {
		NewQModbusPduFromPointer(ptr).DestroyQModbusPduDefault()
	}
}

func (ptr *QModbusPdu) ConnectDestroyQModbusPdu(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusPdu::~QModbusPdu", f)
	}
}

func (ptr *QModbusPdu) DisconnectDestroyQModbusPdu() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusPdu::~QModbusPdu")
	}
}

func (ptr *QModbusPdu) DestroyQModbusPdu() {
	if ptr.Pointer() != nil {
		C.QModbusPdu_DestroyQModbusPdu(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QModbusPdu) DestroyQModbusPduDefault() {
	if ptr.Pointer() != nil {
		C.QModbusPdu_DestroyQModbusPduDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

type QModbusReply struct {
	core.QObject
}

type QModbusReply_ITF interface {
	core.QObject_ITF
	QModbusReply_PTR() *QModbusReply
}

func (ptr *QModbusReply) QModbusReply_PTR() *QModbusReply {
	return ptr
}

func (ptr *QModbusReply) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QModbusReply) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
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

func (ptr *QModbusReply) DestroyQModbusReply() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QModbusReply__ReplyType
//QModbusReply::ReplyType
type QModbusReply__ReplyType int64

const (
	QModbusReply__Raw    QModbusReply__ReplyType = QModbusReply__ReplyType(0)
	QModbusReply__Common QModbusReply__ReplyType = QModbusReply__ReplyType(1)
)

func NewQModbusReply(ty QModbusReply__ReplyType, serverAddress int, parent core.QObject_ITF) *QModbusReply {
	var tmpValue = NewQModbusReplyFromPointer(C.QModbusReply_NewQModbusReply(C.longlong(ty), C.int(int32(serverAddress)), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QModbusReply) Error() QModbusDevice__Error {
	if ptr.Pointer() != nil {
		return QModbusDevice__Error(C.QModbusReply_Error(ptr.Pointer()))
	}
	return 0
}

//export callbackQModbusReply_ErrorOccurred
func callbackQModbusReply_ErrorOccurred(ptr unsafe.Pointer, error C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusReply::errorOccurred"); signal != nil {
		signal.(func(QModbusDevice__Error))(QModbusDevice__Error(error))
	}

}

func (ptr *QModbusReply) ConnectErrorOccurred(f func(error QModbusDevice__Error)) {
	if ptr.Pointer() != nil {
		C.QModbusReply_ConnectErrorOccurred(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusReply::errorOccurred", f)
	}
}

func (ptr *QModbusReply) DisconnectErrorOccurred() {
	if ptr.Pointer() != nil {
		C.QModbusReply_DisconnectErrorOccurred(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusReply::errorOccurred")
	}
}

func (ptr *QModbusReply) ErrorOccurred(error QModbusDevice__Error) {
	if ptr.Pointer() != nil {
		C.QModbusReply_ErrorOccurred(ptr.Pointer(), C.longlong(error))
	}
}

func (ptr *QModbusReply) ErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QModbusReply_ErrorString(ptr.Pointer()))
	}
	return ""
}

//export callbackQModbusReply_Finished
func callbackQModbusReply_Finished(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusReply::finished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QModbusReply) ConnectFinished(f func()) {
	if ptr.Pointer() != nil {
		C.QModbusReply_ConnectFinished(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusReply::finished", f)
	}
}

func (ptr *QModbusReply) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QModbusReply_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusReply::finished")
	}
}

func (ptr *QModbusReply) Finished() {
	if ptr.Pointer() != nil {
		C.QModbusReply_Finished(ptr.Pointer())
	}
}

func (ptr *QModbusReply) IsFinished() bool {
	if ptr.Pointer() != nil {
		return C.QModbusReply_IsFinished(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QModbusReply) RawResult() *QModbusResponse {
	if ptr.Pointer() != nil {
		var tmpValue = NewQModbusResponseFromPointer(C.QModbusReply_RawResult(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QModbusResponse).DestroyQModbusResponse)
		return tmpValue
	}
	return nil
}

func (ptr *QModbusReply) Result() *QModbusDataUnit {
	if ptr.Pointer() != nil {
		var tmpValue = NewQModbusDataUnitFromPointer(C.QModbusReply_Result(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QModbusDataUnit).DestroyQModbusDataUnit)
		return tmpValue
	}
	return nil
}

func (ptr *QModbusReply) ServerAddress() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QModbusReply_ServerAddress(ptr.Pointer())))
	}
	return 0
}

func (ptr *QModbusReply) Type() QModbusReply__ReplyType {
	if ptr.Pointer() != nil {
		return QModbusReply__ReplyType(C.QModbusReply_Type(ptr.Pointer()))
	}
	return 0
}

//export callbackQModbusReply_TimerEvent
func callbackQModbusReply_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusReply::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQModbusReplyFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QModbusReply) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusReply::timerEvent", f)
	}
}

func (ptr *QModbusReply) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusReply::timerEvent")
	}
}

func (ptr *QModbusReply) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusReply_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QModbusReply) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusReply_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQModbusReply_ChildEvent
func callbackQModbusReply_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusReply::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQModbusReplyFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QModbusReply) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusReply::childEvent", f)
	}
}

func (ptr *QModbusReply) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusReply::childEvent")
	}
}

func (ptr *QModbusReply) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusReply_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QModbusReply) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusReply_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQModbusReply_ConnectNotify
func callbackQModbusReply_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusReply::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQModbusReplyFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QModbusReply) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusReply::connectNotify", f)
	}
}

func (ptr *QModbusReply) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusReply::connectNotify")
	}
}

func (ptr *QModbusReply) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusReply_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QModbusReply) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusReply_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQModbusReply_CustomEvent
func callbackQModbusReply_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusReply::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQModbusReplyFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QModbusReply) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusReply::customEvent", f)
	}
}

func (ptr *QModbusReply) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusReply::customEvent")
	}
}

func (ptr *QModbusReply) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusReply_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QModbusReply) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusReply_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQModbusReply_DeleteLater
func callbackQModbusReply_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusReply::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQModbusReplyFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QModbusReply) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusReply::deleteLater", f)
	}
}

func (ptr *QModbusReply) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusReply::deleteLater")
	}
}

func (ptr *QModbusReply) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QModbusReply_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QModbusReply) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QModbusReply_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQModbusReply_DisconnectNotify
func callbackQModbusReply_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusReply::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQModbusReplyFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QModbusReply) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusReply::disconnectNotify", f)
	}
}

func (ptr *QModbusReply) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusReply::disconnectNotify")
	}
}

func (ptr *QModbusReply) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusReply_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QModbusReply) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusReply_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQModbusReply_Event
func callbackQModbusReply_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusReply::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusReplyFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QModbusReply) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusReply::event", f)
	}
}

func (ptr *QModbusReply) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusReply::event")
	}
}

func (ptr *QModbusReply) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusReply_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QModbusReply) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusReply_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQModbusReply_EventFilter
func callbackQModbusReply_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusReply::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusReplyFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QModbusReply) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusReply::eventFilter", f)
	}
}

func (ptr *QModbusReply) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusReply::eventFilter")
	}
}

func (ptr *QModbusReply) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusReply_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QModbusReply) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusReply_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQModbusReply_MetaObject
func callbackQModbusReply_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusReply::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQModbusReplyFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QModbusReply) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusReply::metaObject", f)
	}
}

func (ptr *QModbusReply) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusReply::metaObject")
	}
}

func (ptr *QModbusReply) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QModbusReply_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QModbusReply) MetaObjectDefault() *core.QMetaObject {
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

func (ptr *QModbusRequest) QModbusRequest_PTR() *QModbusRequest {
	return ptr
}

func (ptr *QModbusRequest) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QModbusPdu_PTR().Pointer()
	}
	return nil
}

func (ptr *QModbusRequest) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QModbusPdu_PTR().SetPointer(p)
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

func (ptr *QModbusRequest) DestroyQModbusRequest() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func NewQModbusRequest() *QModbusRequest {
	var tmpValue = NewQModbusRequestFromPointer(C.QModbusRequest_NewQModbusRequest())
	runtime.SetFinalizer(tmpValue, (*QModbusRequest).DestroyQModbusRequest)
	return tmpValue
}

func NewQModbusRequest3(code QModbusPdu__FunctionCode, data core.QByteArray_ITF) *QModbusRequest {
	var tmpValue = NewQModbusRequestFromPointer(C.QModbusRequest_NewQModbusRequest3(C.longlong(code), core.PointerFromQByteArray(data)))
	runtime.SetFinalizer(tmpValue, (*QModbusRequest).DestroyQModbusRequest)
	return tmpValue
}

func NewQModbusRequest2(pdu QModbusPdu_ITF) *QModbusRequest {
	var tmpValue = NewQModbusRequestFromPointer(C.QModbusRequest_NewQModbusRequest2(PointerFromQModbusPdu(pdu)))
	runtime.SetFinalizer(tmpValue, (*QModbusRequest).DestroyQModbusRequest)
	return tmpValue
}

func QModbusRequest_CalculateDataSize(request QModbusRequest_ITF) int {
	return int(int32(C.QModbusRequest_QModbusRequest_CalculateDataSize(PointerFromQModbusRequest(request))))
}

func (ptr *QModbusRequest) CalculateDataSize(request QModbusRequest_ITF) int {
	return int(int32(C.QModbusRequest_QModbusRequest_CalculateDataSize(PointerFromQModbusRequest(request))))
}

func QModbusRequest_MinimumDataSize(request QModbusRequest_ITF) int {
	return int(int32(C.QModbusRequest_QModbusRequest_MinimumDataSize(PointerFromQModbusRequest(request))))
}

func (ptr *QModbusRequest) MinimumDataSize(request QModbusRequest_ITF) int {
	return int(int32(C.QModbusRequest_QModbusRequest_MinimumDataSize(PointerFromQModbusRequest(request))))
}

//export callbackQModbusRequest_SetFunctionCode
func callbackQModbusRequest_SetFunctionCode(ptr unsafe.Pointer, code C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusRequest::setFunctionCode"); signal != nil {
		signal.(func(QModbusPdu__FunctionCode))(QModbusPdu__FunctionCode(code))
	} else {
		NewQModbusRequestFromPointer(ptr).SetFunctionCodeDefault(QModbusPdu__FunctionCode(code))
	}
}

func (ptr *QModbusRequest) ConnectSetFunctionCode(f func(code QModbusPdu__FunctionCode)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRequest::setFunctionCode", f)
	}
}

func (ptr *QModbusRequest) DisconnectSetFunctionCode() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRequest::setFunctionCode")
	}
}

func (ptr *QModbusRequest) SetFunctionCode(code QModbusPdu__FunctionCode) {
	if ptr.Pointer() != nil {
		C.QModbusRequest_SetFunctionCode(ptr.Pointer(), C.longlong(code))
	}
}

func (ptr *QModbusRequest) SetFunctionCodeDefault(code QModbusPdu__FunctionCode) {
	if ptr.Pointer() != nil {
		C.QModbusRequest_SetFunctionCodeDefault(ptr.Pointer(), C.longlong(code))
	}
}

type QModbusResponse struct {
	QModbusPdu
}

type QModbusResponse_ITF interface {
	QModbusPdu_ITF
	QModbusResponse_PTR() *QModbusResponse
}

func (ptr *QModbusResponse) QModbusResponse_PTR() *QModbusResponse {
	return ptr
}

func (ptr *QModbusResponse) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QModbusPdu_PTR().Pointer()
	}
	return nil
}

func (ptr *QModbusResponse) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QModbusPdu_PTR().SetPointer(p)
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

func (ptr *QModbusResponse) DestroyQModbusResponse() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func NewQModbusResponse() *QModbusResponse {
	var tmpValue = NewQModbusResponseFromPointer(C.QModbusResponse_NewQModbusResponse())
	runtime.SetFinalizer(tmpValue, (*QModbusResponse).DestroyQModbusResponse)
	return tmpValue
}

func NewQModbusResponse3(code QModbusPdu__FunctionCode, data core.QByteArray_ITF) *QModbusResponse {
	var tmpValue = NewQModbusResponseFromPointer(C.QModbusResponse_NewQModbusResponse3(C.longlong(code), core.PointerFromQByteArray(data)))
	runtime.SetFinalizer(tmpValue, (*QModbusResponse).DestroyQModbusResponse)
	return tmpValue
}

func NewQModbusResponse2(pdu QModbusPdu_ITF) *QModbusResponse {
	var tmpValue = NewQModbusResponseFromPointer(C.QModbusResponse_NewQModbusResponse2(PointerFromQModbusPdu(pdu)))
	runtime.SetFinalizer(tmpValue, (*QModbusResponse).DestroyQModbusResponse)
	return tmpValue
}

func QModbusResponse_CalculateDataSize(response QModbusResponse_ITF) int {
	return int(int32(C.QModbusResponse_QModbusResponse_CalculateDataSize(PointerFromQModbusResponse(response))))
}

func (ptr *QModbusResponse) CalculateDataSize(response QModbusResponse_ITF) int {
	return int(int32(C.QModbusResponse_QModbusResponse_CalculateDataSize(PointerFromQModbusResponse(response))))
}

func QModbusResponse_MinimumDataSize(response QModbusResponse_ITF) int {
	return int(int32(C.QModbusResponse_QModbusResponse_MinimumDataSize(PointerFromQModbusResponse(response))))
}

func (ptr *QModbusResponse) MinimumDataSize(response QModbusResponse_ITF) int {
	return int(int32(C.QModbusResponse_QModbusResponse_MinimumDataSize(PointerFromQModbusResponse(response))))
}

//export callbackQModbusResponse_SetFunctionCode
func callbackQModbusResponse_SetFunctionCode(ptr unsafe.Pointer, code C.longlong) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusResponse::setFunctionCode"); signal != nil {
		signal.(func(QModbusPdu__FunctionCode))(QModbusPdu__FunctionCode(code))
	} else {
		NewQModbusResponseFromPointer(ptr).SetFunctionCodeDefault(QModbusPdu__FunctionCode(code))
	}
}

func (ptr *QModbusResponse) ConnectSetFunctionCode(f func(code QModbusPdu__FunctionCode)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusResponse::setFunctionCode", f)
	}
}

func (ptr *QModbusResponse) DisconnectSetFunctionCode() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusResponse::setFunctionCode")
	}
}

func (ptr *QModbusResponse) SetFunctionCode(code QModbusPdu__FunctionCode) {
	if ptr.Pointer() != nil {
		C.QModbusResponse_SetFunctionCode(ptr.Pointer(), C.longlong(code))
	}
}

func (ptr *QModbusResponse) SetFunctionCodeDefault(code QModbusPdu__FunctionCode) {
	if ptr.Pointer() != nil {
		C.QModbusResponse_SetFunctionCodeDefault(ptr.Pointer(), C.longlong(code))
	}
}

type QModbusRtuSerialMaster struct {
	QModbusClient
}

type QModbusRtuSerialMaster_ITF interface {
	QModbusClient_ITF
	QModbusRtuSerialMaster_PTR() *QModbusRtuSerialMaster
}

func (ptr *QModbusRtuSerialMaster) QModbusRtuSerialMaster_PTR() *QModbusRtuSerialMaster {
	return ptr
}

func (ptr *QModbusRtuSerialMaster) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QModbusClient_PTR().Pointer()
	}
	return nil
}

func (ptr *QModbusRtuSerialMaster) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QModbusClient_PTR().SetPointer(p)
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

func (ptr *QModbusRtuSerialMaster) DestroyQModbusRtuSerialMaster() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func NewQModbusRtuSerialMaster(parent core.QObject_ITF) *QModbusRtuSerialMaster {
	var tmpValue = NewQModbusRtuSerialMasterFromPointer(C.QModbusRtuSerialMaster_NewQModbusRtuSerialMaster(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QModbusRtuSerialMaster) InterFrameDelay() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QModbusRtuSerialMaster_InterFrameDelay(ptr.Pointer())))
	}
	return 0
}

func (ptr *QModbusRtuSerialMaster) SetInterFrameDelay(microseconds int) {
	if ptr.Pointer() != nil {
		C.QModbusRtuSerialMaster_SetInterFrameDelay(ptr.Pointer(), C.int(int32(microseconds)))
	}
}

//export callbackQModbusRtuSerialMaster_ProcessPrivateResponse
func callbackQModbusRtuSerialMaster_ProcessPrivateResponse(ptr unsafe.Pointer, response unsafe.Pointer, data unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusRtuSerialMaster::processPrivateResponse"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QModbusResponse, *QModbusDataUnit) bool)(NewQModbusResponseFromPointer(response), NewQModbusDataUnitFromPointer(data)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusRtuSerialMasterFromPointer(ptr).ProcessPrivateResponseDefault(NewQModbusResponseFromPointer(response), NewQModbusDataUnitFromPointer(data)))))
}

func (ptr *QModbusRtuSerialMaster) ConnectProcessPrivateResponse(f func(response *QModbusResponse, data *QModbusDataUnit) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialMaster::processPrivateResponse", f)
	}
}

func (ptr *QModbusRtuSerialMaster) DisconnectProcessPrivateResponse() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialMaster::processPrivateResponse")
	}
}

func (ptr *QModbusRtuSerialMaster) ProcessPrivateResponse(response QModbusResponse_ITF, data QModbusDataUnit_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusRtuSerialMaster_ProcessPrivateResponse(ptr.Pointer(), PointerFromQModbusResponse(response), PointerFromQModbusDataUnit(data)) != 0
	}
	return false
}

func (ptr *QModbusRtuSerialMaster) ProcessPrivateResponseDefault(response QModbusResponse_ITF, data QModbusDataUnit_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusRtuSerialMaster_ProcessPrivateResponseDefault(ptr.Pointer(), PointerFromQModbusResponse(response), PointerFromQModbusDataUnit(data)) != 0
	}
	return false
}

//export callbackQModbusRtuSerialMaster_ProcessResponse
func callbackQModbusRtuSerialMaster_ProcessResponse(ptr unsafe.Pointer, response unsafe.Pointer, data unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusRtuSerialMaster::processResponse"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QModbusResponse, *QModbusDataUnit) bool)(NewQModbusResponseFromPointer(response), NewQModbusDataUnitFromPointer(data)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusRtuSerialMasterFromPointer(ptr).ProcessResponseDefault(NewQModbusResponseFromPointer(response), NewQModbusDataUnitFromPointer(data)))))
}

func (ptr *QModbusRtuSerialMaster) ConnectProcessResponse(f func(response *QModbusResponse, data *QModbusDataUnit) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialMaster::processResponse", f)
	}
}

func (ptr *QModbusRtuSerialMaster) DisconnectProcessResponse() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialMaster::processResponse")
	}
}

func (ptr *QModbusRtuSerialMaster) ProcessResponse(response QModbusResponse_ITF, data QModbusDataUnit_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusRtuSerialMaster_ProcessResponse(ptr.Pointer(), PointerFromQModbusResponse(response), PointerFromQModbusDataUnit(data)) != 0
	}
	return false
}

func (ptr *QModbusRtuSerialMaster) ProcessResponseDefault(response QModbusResponse_ITF, data QModbusDataUnit_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusRtuSerialMaster_ProcessResponseDefault(ptr.Pointer(), PointerFromQModbusResponse(response), PointerFromQModbusDataUnit(data)) != 0
	}
	return false
}

//export callbackQModbusRtuSerialMaster_Close
func callbackQModbusRtuSerialMaster_Close(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusRtuSerialMaster::close"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QModbusRtuSerialMaster) ConnectClose(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialMaster::close", f)
	}
}

func (ptr *QModbusRtuSerialMaster) DisconnectClose() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialMaster::close")
	}
}

func (ptr *QModbusRtuSerialMaster) Close() {
	if ptr.Pointer() != nil {
		C.QModbusRtuSerialMaster_Close(ptr.Pointer())
	}
}

//export callbackQModbusRtuSerialMaster_Open
func callbackQModbusRtuSerialMaster_Open(ptr unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusRtuSerialMaster::open"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QModbusRtuSerialMaster) ConnectOpen(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialMaster::open", f)
	}
}

func (ptr *QModbusRtuSerialMaster) DisconnectOpen() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialMaster::open")
	}
}

func (ptr *QModbusRtuSerialMaster) Open() bool {
	if ptr.Pointer() != nil {
		return C.QModbusRtuSerialMaster_Open(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQModbusRtuSerialMaster_TimerEvent
func callbackQModbusRtuSerialMaster_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusRtuSerialMaster::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQModbusRtuSerialMasterFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QModbusRtuSerialMaster) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialMaster::timerEvent", f)
	}
}

func (ptr *QModbusRtuSerialMaster) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialMaster::timerEvent")
	}
}

func (ptr *QModbusRtuSerialMaster) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusRtuSerialMaster_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QModbusRtuSerialMaster) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusRtuSerialMaster_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQModbusRtuSerialMaster_ChildEvent
func callbackQModbusRtuSerialMaster_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusRtuSerialMaster::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQModbusRtuSerialMasterFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QModbusRtuSerialMaster) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialMaster::childEvent", f)
	}
}

func (ptr *QModbusRtuSerialMaster) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialMaster::childEvent")
	}
}

func (ptr *QModbusRtuSerialMaster) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusRtuSerialMaster_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QModbusRtuSerialMaster) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusRtuSerialMaster_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQModbusRtuSerialMaster_ConnectNotify
func callbackQModbusRtuSerialMaster_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusRtuSerialMaster::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQModbusRtuSerialMasterFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QModbusRtuSerialMaster) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialMaster::connectNotify", f)
	}
}

func (ptr *QModbusRtuSerialMaster) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialMaster::connectNotify")
	}
}

func (ptr *QModbusRtuSerialMaster) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusRtuSerialMaster_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QModbusRtuSerialMaster) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusRtuSerialMaster_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQModbusRtuSerialMaster_CustomEvent
func callbackQModbusRtuSerialMaster_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusRtuSerialMaster::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQModbusRtuSerialMasterFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QModbusRtuSerialMaster) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialMaster::customEvent", f)
	}
}

func (ptr *QModbusRtuSerialMaster) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialMaster::customEvent")
	}
}

func (ptr *QModbusRtuSerialMaster) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusRtuSerialMaster_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QModbusRtuSerialMaster) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusRtuSerialMaster_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQModbusRtuSerialMaster_DeleteLater
func callbackQModbusRtuSerialMaster_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusRtuSerialMaster::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQModbusRtuSerialMasterFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QModbusRtuSerialMaster) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialMaster::deleteLater", f)
	}
}

func (ptr *QModbusRtuSerialMaster) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialMaster::deleteLater")
	}
}

func (ptr *QModbusRtuSerialMaster) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QModbusRtuSerialMaster_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QModbusRtuSerialMaster) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QModbusRtuSerialMaster_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQModbusRtuSerialMaster_DisconnectNotify
func callbackQModbusRtuSerialMaster_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusRtuSerialMaster::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQModbusRtuSerialMasterFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QModbusRtuSerialMaster) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialMaster::disconnectNotify", f)
	}
}

func (ptr *QModbusRtuSerialMaster) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialMaster::disconnectNotify")
	}
}

func (ptr *QModbusRtuSerialMaster) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusRtuSerialMaster_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QModbusRtuSerialMaster) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusRtuSerialMaster_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQModbusRtuSerialMaster_Event
func callbackQModbusRtuSerialMaster_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusRtuSerialMaster::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusRtuSerialMasterFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QModbusRtuSerialMaster) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialMaster::event", f)
	}
}

func (ptr *QModbusRtuSerialMaster) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialMaster::event")
	}
}

func (ptr *QModbusRtuSerialMaster) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusRtuSerialMaster_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QModbusRtuSerialMaster) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusRtuSerialMaster_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQModbusRtuSerialMaster_EventFilter
func callbackQModbusRtuSerialMaster_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusRtuSerialMaster::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusRtuSerialMasterFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QModbusRtuSerialMaster) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialMaster::eventFilter", f)
	}
}

func (ptr *QModbusRtuSerialMaster) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialMaster::eventFilter")
	}
}

func (ptr *QModbusRtuSerialMaster) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusRtuSerialMaster_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QModbusRtuSerialMaster) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusRtuSerialMaster_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQModbusRtuSerialMaster_MetaObject
func callbackQModbusRtuSerialMaster_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusRtuSerialMaster::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQModbusRtuSerialMasterFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QModbusRtuSerialMaster) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialMaster::metaObject", f)
	}
}

func (ptr *QModbusRtuSerialMaster) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialMaster::metaObject")
	}
}

func (ptr *QModbusRtuSerialMaster) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QModbusRtuSerialMaster_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QModbusRtuSerialMaster) MetaObjectDefault() *core.QMetaObject {
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

func (ptr *QModbusRtuSerialSlave) QModbusRtuSerialSlave_PTR() *QModbusRtuSerialSlave {
	return ptr
}

func (ptr *QModbusRtuSerialSlave) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QModbusServer_PTR().Pointer()
	}
	return nil
}

func (ptr *QModbusRtuSerialSlave) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QModbusServer_PTR().SetPointer(p)
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
func NewQModbusRtuSerialSlave(parent core.QObject_ITF) *QModbusRtuSerialSlave {
	var tmpValue = NewQModbusRtuSerialSlaveFromPointer(C.QModbusRtuSerialSlave_NewQModbusRtuSerialSlave(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QModbusRtuSerialSlave) DestroyQModbusRtuSerialSlave() {
	if ptr.Pointer() != nil {
		C.QModbusRtuSerialSlave_DestroyQModbusRtuSerialSlave(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQModbusRtuSerialSlave_ProcessPrivateRequest
func callbackQModbusRtuSerialSlave_ProcessPrivateRequest(ptr unsafe.Pointer, request unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusRtuSerialSlave::processPrivateRequest"); signal != nil {
		return PointerFromQModbusResponse(signal.(func(*QModbusPdu) *QModbusResponse)(NewQModbusPduFromPointer(request)))
	}

	return PointerFromQModbusResponse(NewQModbusRtuSerialSlaveFromPointer(ptr).ProcessPrivateRequestDefault(NewQModbusPduFromPointer(request)))
}

func (ptr *QModbusRtuSerialSlave) ConnectProcessPrivateRequest(f func(request *QModbusPdu) *QModbusResponse) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialSlave::processPrivateRequest", f)
	}
}

func (ptr *QModbusRtuSerialSlave) DisconnectProcessPrivateRequest() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialSlave::processPrivateRequest")
	}
}

func (ptr *QModbusRtuSerialSlave) ProcessPrivateRequest(request QModbusPdu_ITF) *QModbusResponse {
	if ptr.Pointer() != nil {
		var tmpValue = NewQModbusResponseFromPointer(C.QModbusRtuSerialSlave_ProcessPrivateRequest(ptr.Pointer(), PointerFromQModbusPdu(request)))
		runtime.SetFinalizer(tmpValue, (*QModbusResponse).DestroyQModbusResponse)
		return tmpValue
	}
	return nil
}

func (ptr *QModbusRtuSerialSlave) ProcessPrivateRequestDefault(request QModbusPdu_ITF) *QModbusResponse {
	if ptr.Pointer() != nil {
		var tmpValue = NewQModbusResponseFromPointer(C.QModbusRtuSerialSlave_ProcessPrivateRequestDefault(ptr.Pointer(), PointerFromQModbusPdu(request)))
		runtime.SetFinalizer(tmpValue, (*QModbusResponse).DestroyQModbusResponse)
		return tmpValue
	}
	return nil
}

//export callbackQModbusRtuSerialSlave_ProcessRequest
func callbackQModbusRtuSerialSlave_ProcessRequest(ptr unsafe.Pointer, request unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusRtuSerialSlave::processRequest"); signal != nil {
		return PointerFromQModbusResponse(signal.(func(*QModbusPdu) *QModbusResponse)(NewQModbusPduFromPointer(request)))
	}

	return PointerFromQModbusResponse(NewQModbusRtuSerialSlaveFromPointer(ptr).ProcessRequestDefault(NewQModbusPduFromPointer(request)))
}

func (ptr *QModbusRtuSerialSlave) ConnectProcessRequest(f func(request *QModbusPdu) *QModbusResponse) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialSlave::processRequest", f)
	}
}

func (ptr *QModbusRtuSerialSlave) DisconnectProcessRequest() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialSlave::processRequest")
	}
}

func (ptr *QModbusRtuSerialSlave) ProcessRequest(request QModbusPdu_ITF) *QModbusResponse {
	if ptr.Pointer() != nil {
		var tmpValue = NewQModbusResponseFromPointer(C.QModbusRtuSerialSlave_ProcessRequest(ptr.Pointer(), PointerFromQModbusPdu(request)))
		runtime.SetFinalizer(tmpValue, (*QModbusResponse).DestroyQModbusResponse)
		return tmpValue
	}
	return nil
}

func (ptr *QModbusRtuSerialSlave) ProcessRequestDefault(request QModbusPdu_ITF) *QModbusResponse {
	if ptr.Pointer() != nil {
		var tmpValue = NewQModbusResponseFromPointer(C.QModbusRtuSerialSlave_ProcessRequestDefault(ptr.Pointer(), PointerFromQModbusPdu(request)))
		runtime.SetFinalizer(tmpValue, (*QModbusResponse).DestroyQModbusResponse)
		return tmpValue
	}
	return nil
}

//export callbackQModbusRtuSerialSlave_ProcessesBroadcast
func callbackQModbusRtuSerialSlave_ProcessesBroadcast(ptr unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusRtuSerialSlave::processesBroadcast"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusRtuSerialSlaveFromPointer(ptr).ProcessesBroadcastDefault())))
}

func (ptr *QModbusRtuSerialSlave) ConnectProcessesBroadcast(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialSlave::processesBroadcast", f)
	}
}

func (ptr *QModbusRtuSerialSlave) DisconnectProcessesBroadcast() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialSlave::processesBroadcast")
	}
}

func (ptr *QModbusRtuSerialSlave) ProcessesBroadcast() bool {
	if ptr.Pointer() != nil {
		return C.QModbusRtuSerialSlave_ProcessesBroadcast(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QModbusRtuSerialSlave) ProcessesBroadcastDefault() bool {
	if ptr.Pointer() != nil {
		return C.QModbusRtuSerialSlave_ProcessesBroadcastDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQModbusRtuSerialSlave_ReadData
func callbackQModbusRtuSerialSlave_ReadData(ptr unsafe.Pointer, newData unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusRtuSerialSlave::readData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QModbusDataUnit) bool)(NewQModbusDataUnitFromPointer(newData)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusRtuSerialSlaveFromPointer(ptr).ReadDataDefault(NewQModbusDataUnitFromPointer(newData)))))
}

func (ptr *QModbusRtuSerialSlave) ConnectReadData(f func(newData *QModbusDataUnit) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialSlave::readData", f)
	}
}

func (ptr *QModbusRtuSerialSlave) DisconnectReadData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialSlave::readData")
	}
}

func (ptr *QModbusRtuSerialSlave) ReadData(newData QModbusDataUnit_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusRtuSerialSlave_ReadData(ptr.Pointer(), PointerFromQModbusDataUnit(newData)) != 0
	}
	return false
}

func (ptr *QModbusRtuSerialSlave) ReadDataDefault(newData QModbusDataUnit_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusRtuSerialSlave_ReadDataDefault(ptr.Pointer(), PointerFromQModbusDataUnit(newData)) != 0
	}
	return false
}

//export callbackQModbusRtuSerialSlave_SetValue
func callbackQModbusRtuSerialSlave_SetValue(ptr unsafe.Pointer, option C.int, newValue unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusRtuSerialSlave::setValue"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, *core.QVariant) bool)(int(int32(option)), core.NewQVariantFromPointer(newValue)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusRtuSerialSlaveFromPointer(ptr).SetValueDefault(int(int32(option)), core.NewQVariantFromPointer(newValue)))))
}

func (ptr *QModbusRtuSerialSlave) ConnectSetValue(f func(option int, newValue *core.QVariant) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialSlave::setValue", f)
	}
}

func (ptr *QModbusRtuSerialSlave) DisconnectSetValue() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialSlave::setValue")
	}
}

func (ptr *QModbusRtuSerialSlave) SetValue(option int, newValue core.QVariant_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusRtuSerialSlave_SetValue(ptr.Pointer(), C.int(int32(option)), core.PointerFromQVariant(newValue)) != 0
	}
	return false
}

func (ptr *QModbusRtuSerialSlave) SetValueDefault(option int, newValue core.QVariant_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusRtuSerialSlave_SetValueDefault(ptr.Pointer(), C.int(int32(option)), core.PointerFromQVariant(newValue)) != 0
	}
	return false
}

//export callbackQModbusRtuSerialSlave_Value
func callbackQModbusRtuSerialSlave_Value(ptr unsafe.Pointer, option C.int) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusRtuSerialSlave::value"); signal != nil {
		return core.PointerFromQVariant(signal.(func(int) *core.QVariant)(int(int32(option))))
	}

	return core.PointerFromQVariant(NewQModbusRtuSerialSlaveFromPointer(ptr).ValueDefault(int(int32(option))))
}

func (ptr *QModbusRtuSerialSlave) ConnectValue(f func(option int) *core.QVariant) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialSlave::value", f)
	}
}

func (ptr *QModbusRtuSerialSlave) DisconnectValue() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialSlave::value")
	}
}

func (ptr *QModbusRtuSerialSlave) Value(option int) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QModbusRtuSerialSlave_Value(ptr.Pointer(), C.int(int32(option))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QModbusRtuSerialSlave) ValueDefault(option int) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QModbusRtuSerialSlave_ValueDefault(ptr.Pointer(), C.int(int32(option))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQModbusRtuSerialSlave_WriteData
func callbackQModbusRtuSerialSlave_WriteData(ptr unsafe.Pointer, newData unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusRtuSerialSlave::writeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QModbusDataUnit) bool)(NewQModbusDataUnitFromPointer(newData)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusRtuSerialSlaveFromPointer(ptr).WriteDataDefault(NewQModbusDataUnitFromPointer(newData)))))
}

func (ptr *QModbusRtuSerialSlave) ConnectWriteData(f func(newData *QModbusDataUnit) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialSlave::writeData", f)
	}
}

func (ptr *QModbusRtuSerialSlave) DisconnectWriteData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialSlave::writeData")
	}
}

func (ptr *QModbusRtuSerialSlave) WriteData(newData QModbusDataUnit_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusRtuSerialSlave_WriteData(ptr.Pointer(), PointerFromQModbusDataUnit(newData)) != 0
	}
	return false
}

func (ptr *QModbusRtuSerialSlave) WriteDataDefault(newData QModbusDataUnit_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusRtuSerialSlave_WriteDataDefault(ptr.Pointer(), PointerFromQModbusDataUnit(newData)) != 0
	}
	return false
}

//export callbackQModbusRtuSerialSlave_Close
func callbackQModbusRtuSerialSlave_Close(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusRtuSerialSlave::close"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QModbusRtuSerialSlave) ConnectClose(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialSlave::close", f)
	}
}

func (ptr *QModbusRtuSerialSlave) DisconnectClose() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialSlave::close")
	}
}

func (ptr *QModbusRtuSerialSlave) Close() {
	if ptr.Pointer() != nil {
		C.QModbusRtuSerialSlave_Close(ptr.Pointer())
	}
}

//export callbackQModbusRtuSerialSlave_Open
func callbackQModbusRtuSerialSlave_Open(ptr unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusRtuSerialSlave::open"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QModbusRtuSerialSlave) ConnectOpen(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialSlave::open", f)
	}
}

func (ptr *QModbusRtuSerialSlave) DisconnectOpen() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialSlave::open")
	}
}

func (ptr *QModbusRtuSerialSlave) Open() bool {
	if ptr.Pointer() != nil {
		return C.QModbusRtuSerialSlave_Open(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQModbusRtuSerialSlave_TimerEvent
func callbackQModbusRtuSerialSlave_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusRtuSerialSlave::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQModbusRtuSerialSlaveFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QModbusRtuSerialSlave) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialSlave::timerEvent", f)
	}
}

func (ptr *QModbusRtuSerialSlave) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialSlave::timerEvent")
	}
}

func (ptr *QModbusRtuSerialSlave) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusRtuSerialSlave_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QModbusRtuSerialSlave) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusRtuSerialSlave_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQModbusRtuSerialSlave_ChildEvent
func callbackQModbusRtuSerialSlave_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusRtuSerialSlave::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQModbusRtuSerialSlaveFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QModbusRtuSerialSlave) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialSlave::childEvent", f)
	}
}

func (ptr *QModbusRtuSerialSlave) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialSlave::childEvent")
	}
}

func (ptr *QModbusRtuSerialSlave) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusRtuSerialSlave_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QModbusRtuSerialSlave) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusRtuSerialSlave_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQModbusRtuSerialSlave_ConnectNotify
func callbackQModbusRtuSerialSlave_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusRtuSerialSlave::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQModbusRtuSerialSlaveFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QModbusRtuSerialSlave) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialSlave::connectNotify", f)
	}
}

func (ptr *QModbusRtuSerialSlave) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialSlave::connectNotify")
	}
}

func (ptr *QModbusRtuSerialSlave) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusRtuSerialSlave_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QModbusRtuSerialSlave) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusRtuSerialSlave_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQModbusRtuSerialSlave_CustomEvent
func callbackQModbusRtuSerialSlave_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusRtuSerialSlave::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQModbusRtuSerialSlaveFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QModbusRtuSerialSlave) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialSlave::customEvent", f)
	}
}

func (ptr *QModbusRtuSerialSlave) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialSlave::customEvent")
	}
}

func (ptr *QModbusRtuSerialSlave) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusRtuSerialSlave_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QModbusRtuSerialSlave) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusRtuSerialSlave_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQModbusRtuSerialSlave_DeleteLater
func callbackQModbusRtuSerialSlave_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusRtuSerialSlave::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQModbusRtuSerialSlaveFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QModbusRtuSerialSlave) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialSlave::deleteLater", f)
	}
}

func (ptr *QModbusRtuSerialSlave) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialSlave::deleteLater")
	}
}

func (ptr *QModbusRtuSerialSlave) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QModbusRtuSerialSlave_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QModbusRtuSerialSlave) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QModbusRtuSerialSlave_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQModbusRtuSerialSlave_DisconnectNotify
func callbackQModbusRtuSerialSlave_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusRtuSerialSlave::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQModbusRtuSerialSlaveFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QModbusRtuSerialSlave) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialSlave::disconnectNotify", f)
	}
}

func (ptr *QModbusRtuSerialSlave) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialSlave::disconnectNotify")
	}
}

func (ptr *QModbusRtuSerialSlave) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusRtuSerialSlave_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QModbusRtuSerialSlave) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusRtuSerialSlave_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQModbusRtuSerialSlave_Event
func callbackQModbusRtuSerialSlave_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusRtuSerialSlave::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusRtuSerialSlaveFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QModbusRtuSerialSlave) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialSlave::event", f)
	}
}

func (ptr *QModbusRtuSerialSlave) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialSlave::event")
	}
}

func (ptr *QModbusRtuSerialSlave) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusRtuSerialSlave_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QModbusRtuSerialSlave) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusRtuSerialSlave_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQModbusRtuSerialSlave_EventFilter
func callbackQModbusRtuSerialSlave_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusRtuSerialSlave::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusRtuSerialSlaveFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QModbusRtuSerialSlave) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialSlave::eventFilter", f)
	}
}

func (ptr *QModbusRtuSerialSlave) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialSlave::eventFilter")
	}
}

func (ptr *QModbusRtuSerialSlave) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusRtuSerialSlave_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QModbusRtuSerialSlave) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusRtuSerialSlave_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQModbusRtuSerialSlave_MetaObject
func callbackQModbusRtuSerialSlave_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusRtuSerialSlave::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQModbusRtuSerialSlaveFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QModbusRtuSerialSlave) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialSlave::metaObject", f)
	}
}

func (ptr *QModbusRtuSerialSlave) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusRtuSerialSlave::metaObject")
	}
}

func (ptr *QModbusRtuSerialSlave) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QModbusRtuSerialSlave_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QModbusRtuSerialSlave) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QModbusRtuSerialSlave_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QModbusServer struct {
	QModbusDevice
}

type QModbusServer_ITF interface {
	QModbusDevice_ITF
	QModbusServer_PTR() *QModbusServer
}

func (ptr *QModbusServer) QModbusServer_PTR() *QModbusServer {
	return ptr
}

func (ptr *QModbusServer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QModbusDevice_PTR().Pointer()
	}
	return nil
}

func (ptr *QModbusServer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QModbusDevice_PTR().SetPointer(p)
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

func (ptr *QModbusServer) DestroyQModbusServer() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//go:generate stringer -type=QModbusServer__Option
//QModbusServer::Option
type QModbusServer__Option int64

const (
	QModbusServer__DiagnosticRegister    QModbusServer__Option = QModbusServer__Option(0)
	QModbusServer__ExceptionStatusOffset QModbusServer__Option = QModbusServer__Option(1)
	QModbusServer__DeviceBusy            QModbusServer__Option = QModbusServer__Option(2)
	QModbusServer__AsciiInputDelimiter   QModbusServer__Option = QModbusServer__Option(3)
	QModbusServer__ListenOnlyMode        QModbusServer__Option = QModbusServer__Option(4)
	QModbusServer__ServerIdentifier      QModbusServer__Option = QModbusServer__Option(5)
	QModbusServer__RunIndicatorStatus    QModbusServer__Option = QModbusServer__Option(6)
	QModbusServer__AdditionalData        QModbusServer__Option = QModbusServer__Option(7)
	QModbusServer__DeviceIdentification  QModbusServer__Option = QModbusServer__Option(8)
	QModbusServer__UserOption            QModbusServer__Option = QModbusServer__Option(0x100)
)

func NewQModbusServer(parent core.QObject_ITF) *QModbusServer {
	var tmpValue = NewQModbusServerFromPointer(C.QModbusServer_NewQModbusServer(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QModbusServer) Data(newData QModbusDataUnit_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusServer_Data(ptr.Pointer(), PointerFromQModbusDataUnit(newData)) != 0
	}
	return false
}

func (ptr *QModbusServer) Data2(table QModbusDataUnit__RegisterType, address uint16, data uint16) bool {
	if ptr.Pointer() != nil {
		return C.QModbusServer_Data2(ptr.Pointer(), C.longlong(table), C.ushort(address), C.ushort(data)) != 0
	}
	return false
}

//export callbackQModbusServer_DataWritten
func callbackQModbusServer_DataWritten(ptr unsafe.Pointer, regist C.longlong, address C.int, size C.int) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusServer::dataWritten"); signal != nil {
		signal.(func(QModbusDataUnit__RegisterType, int, int))(QModbusDataUnit__RegisterType(regist), int(int32(address)), int(int32(size)))
	}

}

func (ptr *QModbusServer) ConnectDataWritten(f func(regist QModbusDataUnit__RegisterType, address int, size int)) {
	if ptr.Pointer() != nil {
		C.QModbusServer_ConnectDataWritten(ptr.Pointer())
		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusServer::dataWritten", f)
	}
}

func (ptr *QModbusServer) DisconnectDataWritten() {
	if ptr.Pointer() != nil {
		C.QModbusServer_DisconnectDataWritten(ptr.Pointer())
		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusServer::dataWritten")
	}
}

func (ptr *QModbusServer) DataWritten(regist QModbusDataUnit__RegisterType, address int, size int) {
	if ptr.Pointer() != nil {
		C.QModbusServer_DataWritten(ptr.Pointer(), C.longlong(regist), C.int(int32(address)), C.int(int32(size)))
	}
}

//export callbackQModbusServer_ProcessPrivateRequest
func callbackQModbusServer_ProcessPrivateRequest(ptr unsafe.Pointer, request unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusServer::processPrivateRequest"); signal != nil {
		return PointerFromQModbusResponse(signal.(func(*QModbusPdu) *QModbusResponse)(NewQModbusPduFromPointer(request)))
	}

	return PointerFromQModbusResponse(NewQModbusServerFromPointer(ptr).ProcessPrivateRequestDefault(NewQModbusPduFromPointer(request)))
}

func (ptr *QModbusServer) ConnectProcessPrivateRequest(f func(request *QModbusPdu) *QModbusResponse) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusServer::processPrivateRequest", f)
	}
}

func (ptr *QModbusServer) DisconnectProcessPrivateRequest() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusServer::processPrivateRequest")
	}
}

func (ptr *QModbusServer) ProcessPrivateRequest(request QModbusPdu_ITF) *QModbusResponse {
	if ptr.Pointer() != nil {
		var tmpValue = NewQModbusResponseFromPointer(C.QModbusServer_ProcessPrivateRequest(ptr.Pointer(), PointerFromQModbusPdu(request)))
		runtime.SetFinalizer(tmpValue, (*QModbusResponse).DestroyQModbusResponse)
		return tmpValue
	}
	return nil
}

func (ptr *QModbusServer) ProcessPrivateRequestDefault(request QModbusPdu_ITF) *QModbusResponse {
	if ptr.Pointer() != nil {
		var tmpValue = NewQModbusResponseFromPointer(C.QModbusServer_ProcessPrivateRequestDefault(ptr.Pointer(), PointerFromQModbusPdu(request)))
		runtime.SetFinalizer(tmpValue, (*QModbusResponse).DestroyQModbusResponse)
		return tmpValue
	}
	return nil
}

//export callbackQModbusServer_ProcessRequest
func callbackQModbusServer_ProcessRequest(ptr unsafe.Pointer, request unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusServer::processRequest"); signal != nil {
		return PointerFromQModbusResponse(signal.(func(*QModbusPdu) *QModbusResponse)(NewQModbusPduFromPointer(request)))
	}

	return PointerFromQModbusResponse(NewQModbusServerFromPointer(ptr).ProcessRequestDefault(NewQModbusPduFromPointer(request)))
}

func (ptr *QModbusServer) ConnectProcessRequest(f func(request *QModbusPdu) *QModbusResponse) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusServer::processRequest", f)
	}
}

func (ptr *QModbusServer) DisconnectProcessRequest() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusServer::processRequest")
	}
}

func (ptr *QModbusServer) ProcessRequest(request QModbusPdu_ITF) *QModbusResponse {
	if ptr.Pointer() != nil {
		var tmpValue = NewQModbusResponseFromPointer(C.QModbusServer_ProcessRequest(ptr.Pointer(), PointerFromQModbusPdu(request)))
		runtime.SetFinalizer(tmpValue, (*QModbusResponse).DestroyQModbusResponse)
		return tmpValue
	}
	return nil
}

func (ptr *QModbusServer) ProcessRequestDefault(request QModbusPdu_ITF) *QModbusResponse {
	if ptr.Pointer() != nil {
		var tmpValue = NewQModbusResponseFromPointer(C.QModbusServer_ProcessRequestDefault(ptr.Pointer(), PointerFromQModbusPdu(request)))
		runtime.SetFinalizer(tmpValue, (*QModbusResponse).DestroyQModbusResponse)
		return tmpValue
	}
	return nil
}

//export callbackQModbusServer_ProcessesBroadcast
func callbackQModbusServer_ProcessesBroadcast(ptr unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusServer::processesBroadcast"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusServerFromPointer(ptr).ProcessesBroadcastDefault())))
}

func (ptr *QModbusServer) ConnectProcessesBroadcast(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusServer::processesBroadcast", f)
	}
}

func (ptr *QModbusServer) DisconnectProcessesBroadcast() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusServer::processesBroadcast")
	}
}

func (ptr *QModbusServer) ProcessesBroadcast() bool {
	if ptr.Pointer() != nil {
		return C.QModbusServer_ProcessesBroadcast(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QModbusServer) ProcessesBroadcastDefault() bool {
	if ptr.Pointer() != nil {
		return C.QModbusServer_ProcessesBroadcastDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQModbusServer_ReadData
func callbackQModbusServer_ReadData(ptr unsafe.Pointer, newData unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusServer::readData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QModbusDataUnit) bool)(NewQModbusDataUnitFromPointer(newData)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusServerFromPointer(ptr).ReadDataDefault(NewQModbusDataUnitFromPointer(newData)))))
}

func (ptr *QModbusServer) ConnectReadData(f func(newData *QModbusDataUnit) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusServer::readData", f)
	}
}

func (ptr *QModbusServer) DisconnectReadData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusServer::readData")
	}
}

func (ptr *QModbusServer) ReadData(newData QModbusDataUnit_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusServer_ReadData(ptr.Pointer(), PointerFromQModbusDataUnit(newData)) != 0
	}
	return false
}

func (ptr *QModbusServer) ReadDataDefault(newData QModbusDataUnit_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusServer_ReadDataDefault(ptr.Pointer(), PointerFromQModbusDataUnit(newData)) != 0
	}
	return false
}

func (ptr *QModbusServer) ServerAddress() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QModbusServer_ServerAddress(ptr.Pointer())))
	}
	return 0
}

func (ptr *QModbusServer) SetData2(table QModbusDataUnit__RegisterType, address uint16, data uint16) bool {
	if ptr.Pointer() != nil {
		return C.QModbusServer_SetData2(ptr.Pointer(), C.longlong(table), C.ushort(address), C.ushort(data)) != 0
	}
	return false
}

func (ptr *QModbusServer) SetData(newData QModbusDataUnit_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusServer_SetData(ptr.Pointer(), PointerFromQModbusDataUnit(newData)) != 0
	}
	return false
}

func (ptr *QModbusServer) SetServerAddress(serverAddress int) {
	if ptr.Pointer() != nil {
		C.QModbusServer_SetServerAddress(ptr.Pointer(), C.int(int32(serverAddress)))
	}
}

//export callbackQModbusServer_SetValue
func callbackQModbusServer_SetValue(ptr unsafe.Pointer, option C.int, newValue unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusServer::setValue"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, *core.QVariant) bool)(int(int32(option)), core.NewQVariantFromPointer(newValue)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusServerFromPointer(ptr).SetValueDefault(int(int32(option)), core.NewQVariantFromPointer(newValue)))))
}

func (ptr *QModbusServer) ConnectSetValue(f func(option int, newValue *core.QVariant) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusServer::setValue", f)
	}
}

func (ptr *QModbusServer) DisconnectSetValue() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusServer::setValue")
	}
}

func (ptr *QModbusServer) SetValue(option int, newValue core.QVariant_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusServer_SetValue(ptr.Pointer(), C.int(int32(option)), core.PointerFromQVariant(newValue)) != 0
	}
	return false
}

func (ptr *QModbusServer) SetValueDefault(option int, newValue core.QVariant_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusServer_SetValueDefault(ptr.Pointer(), C.int(int32(option)), core.PointerFromQVariant(newValue)) != 0
	}
	return false
}

//export callbackQModbusServer_Value
func callbackQModbusServer_Value(ptr unsafe.Pointer, option C.int) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusServer::value"); signal != nil {
		return core.PointerFromQVariant(signal.(func(int) *core.QVariant)(int(int32(option))))
	}

	return core.PointerFromQVariant(NewQModbusServerFromPointer(ptr).ValueDefault(int(int32(option))))
}

func (ptr *QModbusServer) ConnectValue(f func(option int) *core.QVariant) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusServer::value", f)
	}
}

func (ptr *QModbusServer) DisconnectValue() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusServer::value")
	}
}

func (ptr *QModbusServer) Value(option int) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QModbusServer_Value(ptr.Pointer(), C.int(int32(option))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QModbusServer) ValueDefault(option int) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QModbusServer_ValueDefault(ptr.Pointer(), C.int(int32(option))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQModbusServer_WriteData
func callbackQModbusServer_WriteData(ptr unsafe.Pointer, newData unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusServer::writeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QModbusDataUnit) bool)(NewQModbusDataUnitFromPointer(newData)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusServerFromPointer(ptr).WriteDataDefault(NewQModbusDataUnitFromPointer(newData)))))
}

func (ptr *QModbusServer) ConnectWriteData(f func(newData *QModbusDataUnit) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusServer::writeData", f)
	}
}

func (ptr *QModbusServer) DisconnectWriteData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusServer::writeData")
	}
}

func (ptr *QModbusServer) WriteData(newData QModbusDataUnit_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusServer_WriteData(ptr.Pointer(), PointerFromQModbusDataUnit(newData)) != 0
	}
	return false
}

func (ptr *QModbusServer) WriteDataDefault(newData QModbusDataUnit_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusServer_WriteDataDefault(ptr.Pointer(), PointerFromQModbusDataUnit(newData)) != 0
	}
	return false
}

//export callbackQModbusServer_Close
func callbackQModbusServer_Close(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusServer::close"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QModbusServer) ConnectClose(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusServer::close", f)
	}
}

func (ptr *QModbusServer) DisconnectClose() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusServer::close")
	}
}

func (ptr *QModbusServer) Close() {
	if ptr.Pointer() != nil {
		C.QModbusServer_Close(ptr.Pointer())
	}
}

//export callbackQModbusServer_Open
func callbackQModbusServer_Open(ptr unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusServer::open"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QModbusServer) ConnectOpen(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusServer::open", f)
	}
}

func (ptr *QModbusServer) DisconnectOpen() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusServer::open")
	}
}

func (ptr *QModbusServer) Open() bool {
	if ptr.Pointer() != nil {
		return C.QModbusServer_Open(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQModbusServer_TimerEvent
func callbackQModbusServer_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusServer::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQModbusServerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QModbusServer) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusServer::timerEvent", f)
	}
}

func (ptr *QModbusServer) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusServer::timerEvent")
	}
}

func (ptr *QModbusServer) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusServer_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QModbusServer) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusServer_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQModbusServer_ChildEvent
func callbackQModbusServer_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusServer::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQModbusServerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QModbusServer) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusServer::childEvent", f)
	}
}

func (ptr *QModbusServer) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusServer::childEvent")
	}
}

func (ptr *QModbusServer) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusServer_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QModbusServer) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusServer_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQModbusServer_ConnectNotify
func callbackQModbusServer_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusServer::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQModbusServerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QModbusServer) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusServer::connectNotify", f)
	}
}

func (ptr *QModbusServer) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusServer::connectNotify")
	}
}

func (ptr *QModbusServer) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusServer_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QModbusServer) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusServer_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQModbusServer_CustomEvent
func callbackQModbusServer_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusServer::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQModbusServerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QModbusServer) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusServer::customEvent", f)
	}
}

func (ptr *QModbusServer) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusServer::customEvent")
	}
}

func (ptr *QModbusServer) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusServer_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QModbusServer) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusServer_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQModbusServer_DeleteLater
func callbackQModbusServer_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusServer::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQModbusServerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QModbusServer) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusServer::deleteLater", f)
	}
}

func (ptr *QModbusServer) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusServer::deleteLater")
	}
}

func (ptr *QModbusServer) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QModbusServer_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QModbusServer) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QModbusServer_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQModbusServer_DisconnectNotify
func callbackQModbusServer_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusServer::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQModbusServerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QModbusServer) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusServer::disconnectNotify", f)
	}
}

func (ptr *QModbusServer) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusServer::disconnectNotify")
	}
}

func (ptr *QModbusServer) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusServer_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QModbusServer) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusServer_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQModbusServer_Event
func callbackQModbusServer_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusServer::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusServerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QModbusServer) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusServer::event", f)
	}
}

func (ptr *QModbusServer) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusServer::event")
	}
}

func (ptr *QModbusServer) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusServer_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QModbusServer) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusServer_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQModbusServer_EventFilter
func callbackQModbusServer_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusServer::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusServerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QModbusServer) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusServer::eventFilter", f)
	}
}

func (ptr *QModbusServer) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusServer::eventFilter")
	}
}

func (ptr *QModbusServer) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusServer_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QModbusServer) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusServer_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQModbusServer_MetaObject
func callbackQModbusServer_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusServer::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQModbusServerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QModbusServer) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusServer::metaObject", f)
	}
}

func (ptr *QModbusServer) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusServer::metaObject")
	}
}

func (ptr *QModbusServer) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QModbusServer_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QModbusServer) MetaObjectDefault() *core.QMetaObject {
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

func (ptr *QModbusTcpClient) QModbusTcpClient_PTR() *QModbusTcpClient {
	return ptr
}

func (ptr *QModbusTcpClient) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QModbusClient_PTR().Pointer()
	}
	return nil
}

func (ptr *QModbusTcpClient) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QModbusClient_PTR().SetPointer(p)
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
func NewQModbusTcpClient(parent core.QObject_ITF) *QModbusTcpClient {
	var tmpValue = NewQModbusTcpClientFromPointer(C.QModbusTcpClient_NewQModbusTcpClient(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QModbusTcpClient) DestroyQModbusTcpClient() {
	if ptr.Pointer() != nil {
		C.QModbusTcpClient_DestroyQModbusTcpClient(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQModbusTcpClient_ProcessPrivateResponse
func callbackQModbusTcpClient_ProcessPrivateResponse(ptr unsafe.Pointer, response unsafe.Pointer, data unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusTcpClient::processPrivateResponse"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QModbusResponse, *QModbusDataUnit) bool)(NewQModbusResponseFromPointer(response), NewQModbusDataUnitFromPointer(data)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusTcpClientFromPointer(ptr).ProcessPrivateResponseDefault(NewQModbusResponseFromPointer(response), NewQModbusDataUnitFromPointer(data)))))
}

func (ptr *QModbusTcpClient) ConnectProcessPrivateResponse(f func(response *QModbusResponse, data *QModbusDataUnit) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpClient::processPrivateResponse", f)
	}
}

func (ptr *QModbusTcpClient) DisconnectProcessPrivateResponse() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpClient::processPrivateResponse")
	}
}

func (ptr *QModbusTcpClient) ProcessPrivateResponse(response QModbusResponse_ITF, data QModbusDataUnit_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusTcpClient_ProcessPrivateResponse(ptr.Pointer(), PointerFromQModbusResponse(response), PointerFromQModbusDataUnit(data)) != 0
	}
	return false
}

func (ptr *QModbusTcpClient) ProcessPrivateResponseDefault(response QModbusResponse_ITF, data QModbusDataUnit_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusTcpClient_ProcessPrivateResponseDefault(ptr.Pointer(), PointerFromQModbusResponse(response), PointerFromQModbusDataUnit(data)) != 0
	}
	return false
}

//export callbackQModbusTcpClient_ProcessResponse
func callbackQModbusTcpClient_ProcessResponse(ptr unsafe.Pointer, response unsafe.Pointer, data unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusTcpClient::processResponse"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QModbusResponse, *QModbusDataUnit) bool)(NewQModbusResponseFromPointer(response), NewQModbusDataUnitFromPointer(data)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusTcpClientFromPointer(ptr).ProcessResponseDefault(NewQModbusResponseFromPointer(response), NewQModbusDataUnitFromPointer(data)))))
}

func (ptr *QModbusTcpClient) ConnectProcessResponse(f func(response *QModbusResponse, data *QModbusDataUnit) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpClient::processResponse", f)
	}
}

func (ptr *QModbusTcpClient) DisconnectProcessResponse() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpClient::processResponse")
	}
}

func (ptr *QModbusTcpClient) ProcessResponse(response QModbusResponse_ITF, data QModbusDataUnit_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusTcpClient_ProcessResponse(ptr.Pointer(), PointerFromQModbusResponse(response), PointerFromQModbusDataUnit(data)) != 0
	}
	return false
}

func (ptr *QModbusTcpClient) ProcessResponseDefault(response QModbusResponse_ITF, data QModbusDataUnit_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusTcpClient_ProcessResponseDefault(ptr.Pointer(), PointerFromQModbusResponse(response), PointerFromQModbusDataUnit(data)) != 0
	}
	return false
}

//export callbackQModbusTcpClient_Close
func callbackQModbusTcpClient_Close(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusTcpClient::close"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QModbusTcpClient) ConnectClose(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpClient::close", f)
	}
}

func (ptr *QModbusTcpClient) DisconnectClose() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpClient::close")
	}
}

func (ptr *QModbusTcpClient) Close() {
	if ptr.Pointer() != nil {
		C.QModbusTcpClient_Close(ptr.Pointer())
	}
}

//export callbackQModbusTcpClient_Open
func callbackQModbusTcpClient_Open(ptr unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusTcpClient::open"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QModbusTcpClient) ConnectOpen(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpClient::open", f)
	}
}

func (ptr *QModbusTcpClient) DisconnectOpen() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpClient::open")
	}
}

func (ptr *QModbusTcpClient) Open() bool {
	if ptr.Pointer() != nil {
		return C.QModbusTcpClient_Open(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQModbusTcpClient_TimerEvent
func callbackQModbusTcpClient_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusTcpClient::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQModbusTcpClientFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QModbusTcpClient) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpClient::timerEvent", f)
	}
}

func (ptr *QModbusTcpClient) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpClient::timerEvent")
	}
}

func (ptr *QModbusTcpClient) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusTcpClient_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QModbusTcpClient) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusTcpClient_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQModbusTcpClient_ChildEvent
func callbackQModbusTcpClient_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusTcpClient::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQModbusTcpClientFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QModbusTcpClient) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpClient::childEvent", f)
	}
}

func (ptr *QModbusTcpClient) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpClient::childEvent")
	}
}

func (ptr *QModbusTcpClient) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusTcpClient_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QModbusTcpClient) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusTcpClient_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQModbusTcpClient_ConnectNotify
func callbackQModbusTcpClient_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusTcpClient::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQModbusTcpClientFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QModbusTcpClient) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpClient::connectNotify", f)
	}
}

func (ptr *QModbusTcpClient) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpClient::connectNotify")
	}
}

func (ptr *QModbusTcpClient) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusTcpClient_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QModbusTcpClient) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusTcpClient_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQModbusTcpClient_CustomEvent
func callbackQModbusTcpClient_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusTcpClient::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQModbusTcpClientFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QModbusTcpClient) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpClient::customEvent", f)
	}
}

func (ptr *QModbusTcpClient) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpClient::customEvent")
	}
}

func (ptr *QModbusTcpClient) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusTcpClient_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QModbusTcpClient) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusTcpClient_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQModbusTcpClient_DeleteLater
func callbackQModbusTcpClient_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusTcpClient::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQModbusTcpClientFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QModbusTcpClient) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpClient::deleteLater", f)
	}
}

func (ptr *QModbusTcpClient) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpClient::deleteLater")
	}
}

func (ptr *QModbusTcpClient) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QModbusTcpClient_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QModbusTcpClient) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QModbusTcpClient_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQModbusTcpClient_DisconnectNotify
func callbackQModbusTcpClient_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusTcpClient::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQModbusTcpClientFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QModbusTcpClient) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpClient::disconnectNotify", f)
	}
}

func (ptr *QModbusTcpClient) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpClient::disconnectNotify")
	}
}

func (ptr *QModbusTcpClient) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusTcpClient_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QModbusTcpClient) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusTcpClient_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQModbusTcpClient_Event
func callbackQModbusTcpClient_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusTcpClient::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusTcpClientFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QModbusTcpClient) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpClient::event", f)
	}
}

func (ptr *QModbusTcpClient) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpClient::event")
	}
}

func (ptr *QModbusTcpClient) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusTcpClient_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QModbusTcpClient) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusTcpClient_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQModbusTcpClient_EventFilter
func callbackQModbusTcpClient_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusTcpClient::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusTcpClientFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QModbusTcpClient) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpClient::eventFilter", f)
	}
}

func (ptr *QModbusTcpClient) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpClient::eventFilter")
	}
}

func (ptr *QModbusTcpClient) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusTcpClient_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QModbusTcpClient) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusTcpClient_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQModbusTcpClient_MetaObject
func callbackQModbusTcpClient_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusTcpClient::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQModbusTcpClientFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QModbusTcpClient) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpClient::metaObject", f)
	}
}

func (ptr *QModbusTcpClient) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpClient::metaObject")
	}
}

func (ptr *QModbusTcpClient) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QModbusTcpClient_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QModbusTcpClient) MetaObjectDefault() *core.QMetaObject {
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

func (ptr *QModbusTcpServer) QModbusTcpServer_PTR() *QModbusTcpServer {
	return ptr
}

func (ptr *QModbusTcpServer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QModbusServer_PTR().Pointer()
	}
	return nil
}

func (ptr *QModbusTcpServer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QModbusServer_PTR().SetPointer(p)
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
func NewQModbusTcpServer(parent core.QObject_ITF) *QModbusTcpServer {
	var tmpValue = NewQModbusTcpServerFromPointer(C.QModbusTcpServer_NewQModbusTcpServer(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(fmt.Sprint(tmpValue.Pointer()), "QObject::destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QModbusTcpServer) DestroyQModbusTcpServer() {
	if ptr.Pointer() != nil {
		C.QModbusTcpServer_DestroyQModbusTcpServer(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQModbusTcpServer_ProcessPrivateRequest
func callbackQModbusTcpServer_ProcessPrivateRequest(ptr unsafe.Pointer, request unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusTcpServer::processPrivateRequest"); signal != nil {
		return PointerFromQModbusResponse(signal.(func(*QModbusPdu) *QModbusResponse)(NewQModbusPduFromPointer(request)))
	}

	return PointerFromQModbusResponse(NewQModbusTcpServerFromPointer(ptr).ProcessPrivateRequestDefault(NewQModbusPduFromPointer(request)))
}

func (ptr *QModbusTcpServer) ConnectProcessPrivateRequest(f func(request *QModbusPdu) *QModbusResponse) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpServer::processPrivateRequest", f)
	}
}

func (ptr *QModbusTcpServer) DisconnectProcessPrivateRequest() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpServer::processPrivateRequest")
	}
}

func (ptr *QModbusTcpServer) ProcessPrivateRequest(request QModbusPdu_ITF) *QModbusResponse {
	if ptr.Pointer() != nil {
		var tmpValue = NewQModbusResponseFromPointer(C.QModbusTcpServer_ProcessPrivateRequest(ptr.Pointer(), PointerFromQModbusPdu(request)))
		runtime.SetFinalizer(tmpValue, (*QModbusResponse).DestroyQModbusResponse)
		return tmpValue
	}
	return nil
}

func (ptr *QModbusTcpServer) ProcessPrivateRequestDefault(request QModbusPdu_ITF) *QModbusResponse {
	if ptr.Pointer() != nil {
		var tmpValue = NewQModbusResponseFromPointer(C.QModbusTcpServer_ProcessPrivateRequestDefault(ptr.Pointer(), PointerFromQModbusPdu(request)))
		runtime.SetFinalizer(tmpValue, (*QModbusResponse).DestroyQModbusResponse)
		return tmpValue
	}
	return nil
}

//export callbackQModbusTcpServer_ProcessRequest
func callbackQModbusTcpServer_ProcessRequest(ptr unsafe.Pointer, request unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusTcpServer::processRequest"); signal != nil {
		return PointerFromQModbusResponse(signal.(func(*QModbusPdu) *QModbusResponse)(NewQModbusPduFromPointer(request)))
	}

	return PointerFromQModbusResponse(NewQModbusTcpServerFromPointer(ptr).ProcessRequestDefault(NewQModbusPduFromPointer(request)))
}

func (ptr *QModbusTcpServer) ConnectProcessRequest(f func(request *QModbusPdu) *QModbusResponse) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpServer::processRequest", f)
	}
}

func (ptr *QModbusTcpServer) DisconnectProcessRequest() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpServer::processRequest")
	}
}

func (ptr *QModbusTcpServer) ProcessRequest(request QModbusPdu_ITF) *QModbusResponse {
	if ptr.Pointer() != nil {
		var tmpValue = NewQModbusResponseFromPointer(C.QModbusTcpServer_ProcessRequest(ptr.Pointer(), PointerFromQModbusPdu(request)))
		runtime.SetFinalizer(tmpValue, (*QModbusResponse).DestroyQModbusResponse)
		return tmpValue
	}
	return nil
}

func (ptr *QModbusTcpServer) ProcessRequestDefault(request QModbusPdu_ITF) *QModbusResponse {
	if ptr.Pointer() != nil {
		var tmpValue = NewQModbusResponseFromPointer(C.QModbusTcpServer_ProcessRequestDefault(ptr.Pointer(), PointerFromQModbusPdu(request)))
		runtime.SetFinalizer(tmpValue, (*QModbusResponse).DestroyQModbusResponse)
		return tmpValue
	}
	return nil
}

//export callbackQModbusTcpServer_ProcessesBroadcast
func callbackQModbusTcpServer_ProcessesBroadcast(ptr unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusTcpServer::processesBroadcast"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusTcpServerFromPointer(ptr).ProcessesBroadcastDefault())))
}

func (ptr *QModbusTcpServer) ConnectProcessesBroadcast(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpServer::processesBroadcast", f)
	}
}

func (ptr *QModbusTcpServer) DisconnectProcessesBroadcast() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpServer::processesBroadcast")
	}
}

func (ptr *QModbusTcpServer) ProcessesBroadcast() bool {
	if ptr.Pointer() != nil {
		return C.QModbusTcpServer_ProcessesBroadcast(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QModbusTcpServer) ProcessesBroadcastDefault() bool {
	if ptr.Pointer() != nil {
		return C.QModbusTcpServer_ProcessesBroadcastDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQModbusTcpServer_ReadData
func callbackQModbusTcpServer_ReadData(ptr unsafe.Pointer, newData unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusTcpServer::readData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QModbusDataUnit) bool)(NewQModbusDataUnitFromPointer(newData)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusTcpServerFromPointer(ptr).ReadDataDefault(NewQModbusDataUnitFromPointer(newData)))))
}

func (ptr *QModbusTcpServer) ConnectReadData(f func(newData *QModbusDataUnit) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpServer::readData", f)
	}
}

func (ptr *QModbusTcpServer) DisconnectReadData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpServer::readData")
	}
}

func (ptr *QModbusTcpServer) ReadData(newData QModbusDataUnit_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusTcpServer_ReadData(ptr.Pointer(), PointerFromQModbusDataUnit(newData)) != 0
	}
	return false
}

func (ptr *QModbusTcpServer) ReadDataDefault(newData QModbusDataUnit_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusTcpServer_ReadDataDefault(ptr.Pointer(), PointerFromQModbusDataUnit(newData)) != 0
	}
	return false
}

//export callbackQModbusTcpServer_SetValue
func callbackQModbusTcpServer_SetValue(ptr unsafe.Pointer, option C.int, newValue unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusTcpServer::setValue"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, *core.QVariant) bool)(int(int32(option)), core.NewQVariantFromPointer(newValue)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusTcpServerFromPointer(ptr).SetValueDefault(int(int32(option)), core.NewQVariantFromPointer(newValue)))))
}

func (ptr *QModbusTcpServer) ConnectSetValue(f func(option int, newValue *core.QVariant) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpServer::setValue", f)
	}
}

func (ptr *QModbusTcpServer) DisconnectSetValue() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpServer::setValue")
	}
}

func (ptr *QModbusTcpServer) SetValue(option int, newValue core.QVariant_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusTcpServer_SetValue(ptr.Pointer(), C.int(int32(option)), core.PointerFromQVariant(newValue)) != 0
	}
	return false
}

func (ptr *QModbusTcpServer) SetValueDefault(option int, newValue core.QVariant_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusTcpServer_SetValueDefault(ptr.Pointer(), C.int(int32(option)), core.PointerFromQVariant(newValue)) != 0
	}
	return false
}

//export callbackQModbusTcpServer_Value
func callbackQModbusTcpServer_Value(ptr unsafe.Pointer, option C.int) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusTcpServer::value"); signal != nil {
		return core.PointerFromQVariant(signal.(func(int) *core.QVariant)(int(int32(option))))
	}

	return core.PointerFromQVariant(NewQModbusTcpServerFromPointer(ptr).ValueDefault(int(int32(option))))
}

func (ptr *QModbusTcpServer) ConnectValue(f func(option int) *core.QVariant) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpServer::value", f)
	}
}

func (ptr *QModbusTcpServer) DisconnectValue() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpServer::value")
	}
}

func (ptr *QModbusTcpServer) Value(option int) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QModbusTcpServer_Value(ptr.Pointer(), C.int(int32(option))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QModbusTcpServer) ValueDefault(option int) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QModbusTcpServer_ValueDefault(ptr.Pointer(), C.int(int32(option))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

//export callbackQModbusTcpServer_WriteData
func callbackQModbusTcpServer_WriteData(ptr unsafe.Pointer, newData unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusTcpServer::writeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QModbusDataUnit) bool)(NewQModbusDataUnitFromPointer(newData)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusTcpServerFromPointer(ptr).WriteDataDefault(NewQModbusDataUnitFromPointer(newData)))))
}

func (ptr *QModbusTcpServer) ConnectWriteData(f func(newData *QModbusDataUnit) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpServer::writeData", f)
	}
}

func (ptr *QModbusTcpServer) DisconnectWriteData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpServer::writeData")
	}
}

func (ptr *QModbusTcpServer) WriteData(newData QModbusDataUnit_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusTcpServer_WriteData(ptr.Pointer(), PointerFromQModbusDataUnit(newData)) != 0
	}
	return false
}

func (ptr *QModbusTcpServer) WriteDataDefault(newData QModbusDataUnit_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusTcpServer_WriteDataDefault(ptr.Pointer(), PointerFromQModbusDataUnit(newData)) != 0
	}
	return false
}

//export callbackQModbusTcpServer_Close
func callbackQModbusTcpServer_Close(ptr unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusTcpServer::close"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QModbusTcpServer) ConnectClose(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpServer::close", f)
	}
}

func (ptr *QModbusTcpServer) DisconnectClose() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpServer::close")
	}
}

func (ptr *QModbusTcpServer) Close() {
	if ptr.Pointer() != nil {
		C.QModbusTcpServer_Close(ptr.Pointer())
	}
}

//export callbackQModbusTcpServer_Open
func callbackQModbusTcpServer_Open(ptr unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusTcpServer::open"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QModbusTcpServer) ConnectOpen(f func() bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpServer::open", f)
	}
}

func (ptr *QModbusTcpServer) DisconnectOpen() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpServer::open")
	}
}

func (ptr *QModbusTcpServer) Open() bool {
	if ptr.Pointer() != nil {
		return C.QModbusTcpServer_Open(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQModbusTcpServer_TimerEvent
func callbackQModbusTcpServer_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusTcpServer::timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQModbusTcpServerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QModbusTcpServer) ConnectTimerEvent(f func(event *core.QTimerEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpServer::timerEvent", f)
	}
}

func (ptr *QModbusTcpServer) DisconnectTimerEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpServer::timerEvent")
	}
}

func (ptr *QModbusTcpServer) TimerEvent(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusTcpServer_TimerEvent(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

func (ptr *QModbusTcpServer) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusTcpServer_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQModbusTcpServer_ChildEvent
func callbackQModbusTcpServer_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusTcpServer::childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQModbusTcpServerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QModbusTcpServer) ConnectChildEvent(f func(event *core.QChildEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpServer::childEvent", f)
	}
}

func (ptr *QModbusTcpServer) DisconnectChildEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpServer::childEvent")
	}
}

func (ptr *QModbusTcpServer) ChildEvent(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusTcpServer_ChildEvent(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

func (ptr *QModbusTcpServer) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusTcpServer_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQModbusTcpServer_ConnectNotify
func callbackQModbusTcpServer_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusTcpServer::connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQModbusTcpServerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QModbusTcpServer) ConnectConnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpServer::connectNotify", f)
	}
}

func (ptr *QModbusTcpServer) DisconnectConnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpServer::connectNotify")
	}
}

func (ptr *QModbusTcpServer) ConnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusTcpServer_ConnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QModbusTcpServer) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusTcpServer_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQModbusTcpServer_CustomEvent
func callbackQModbusTcpServer_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusTcpServer::customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQModbusTcpServerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QModbusTcpServer) ConnectCustomEvent(f func(event *core.QEvent)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpServer::customEvent", f)
	}
}

func (ptr *QModbusTcpServer) DisconnectCustomEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpServer::customEvent")
	}
}

func (ptr *QModbusTcpServer) CustomEvent(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusTcpServer_CustomEvent(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

func (ptr *QModbusTcpServer) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusTcpServer_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQModbusTcpServer_DeleteLater
func callbackQModbusTcpServer_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusTcpServer::deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQModbusTcpServerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QModbusTcpServer) ConnectDeleteLater(f func()) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpServer::deleteLater", f)
	}
}

func (ptr *QModbusTcpServer) DisconnectDeleteLater() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpServer::deleteLater")
	}
}

func (ptr *QModbusTcpServer) DeleteLater() {
	if ptr.Pointer() != nil {
		C.QModbusTcpServer_DeleteLater(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

func (ptr *QModbusTcpServer) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QModbusTcpServer_DeleteLaterDefault(ptr.Pointer())
		qt.DisconnectAllSignals(fmt.Sprint(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}

//export callbackQModbusTcpServer_DisconnectNotify
func callbackQModbusTcpServer_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusTcpServer::disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQModbusTcpServerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QModbusTcpServer) ConnectDisconnectNotify(f func(sign *core.QMetaMethod)) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpServer::disconnectNotify", f)
	}
}

func (ptr *QModbusTcpServer) DisconnectDisconnectNotify() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpServer::disconnectNotify")
	}
}

func (ptr *QModbusTcpServer) DisconnectNotify(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusTcpServer_DisconnectNotify(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

func (ptr *QModbusTcpServer) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusTcpServer_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQModbusTcpServer_Event
func callbackQModbusTcpServer_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusTcpServer::event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusTcpServerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QModbusTcpServer) ConnectEvent(f func(e *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpServer::event", f)
	}
}

func (ptr *QModbusTcpServer) DisconnectEvent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpServer::event")
	}
}

func (ptr *QModbusTcpServer) Event(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusTcpServer_Event(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

func (ptr *QModbusTcpServer) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusTcpServer_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQModbusTcpServer_EventFilter
func callbackQModbusTcpServer_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusTcpServer::eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusTcpServerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QModbusTcpServer) ConnectEventFilter(f func(watched *core.QObject, event *core.QEvent) bool) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpServer::eventFilter", f)
	}
}

func (ptr *QModbusTcpServer) DisconnectEventFilter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpServer::eventFilter")
	}
}

func (ptr *QModbusTcpServer) EventFilter(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusTcpServer_EventFilter(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

func (ptr *QModbusTcpServer) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusTcpServer_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQModbusTcpServer_MetaObject
func callbackQModbusTcpServer_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {

	if signal := qt.GetSignal(fmt.Sprint(ptr), "QModbusTcpServer::metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQModbusTcpServerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QModbusTcpServer) ConnectMetaObject(f func() *core.QMetaObject) {
	if ptr.Pointer() != nil {

		qt.ConnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpServer::metaObject", f)
	}
}

func (ptr *QModbusTcpServer) DisconnectMetaObject() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(fmt.Sprint(ptr.Pointer()), "QModbusTcpServer::metaObject")
	}
}

func (ptr *QModbusTcpServer) MetaObject() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QModbusTcpServer_MetaObject(ptr.Pointer()))
	}
	return nil
}

func (ptr *QModbusTcpServer) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QModbusTcpServer_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}
