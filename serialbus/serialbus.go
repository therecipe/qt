// +build !minimal

package serialbus

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "serialbus.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"runtime"
	"strings"
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
		runtime.SetFinalizer(ptr, nil)
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
		runtime.SetFinalizer(ptr, nil)
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
func QCanBus_Instance() *QCanBus {
	var tmpValue = NewQCanBusFromPointer(C.QCanBus_QCanBus_Instance())
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QCanBus) Instance() *QCanBus {
	var tmpValue = NewQCanBusFromPointer(C.QCanBus_QCanBus_Instance())
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QCanBus) CreateDevice(plugin string, interfaceName string, errorMessage string) *QCanBusDevice {
	if ptr.Pointer() != nil {
		var pluginC *C.char
		if plugin != "" {
			pluginC = C.CString(plugin)
			defer C.free(unsafe.Pointer(pluginC))
		}
		var interfaceNameC *C.char
		if interfaceName != "" {
			interfaceNameC = C.CString(interfaceName)
			defer C.free(unsafe.Pointer(interfaceNameC))
		}
		var errorMessageC *C.char
		if errorMessage != "" {
			errorMessageC = C.CString(errorMessage)
			defer C.free(unsafe.Pointer(errorMessageC))
		}
		var tmpValue = NewQCanBusDeviceFromPointer(C.QCanBus_CreateDevice(ptr.Pointer(), C.struct_QtSerialBus_PackedString{data: pluginC, len: C.longlong(len(plugin))}, C.struct_QtSerialBus_PackedString{data: interfaceNameC, len: C.longlong(len(interfaceName))}, C.struct_QtSerialBus_PackedString{data: errorMessageC, len: C.longlong(len(errorMessage))}))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QCanBus) AvailableDevices(plugin string, errorMessage string) []*QCanBusDeviceInfo {
	if ptr.Pointer() != nil {
		var pluginC *C.char
		if plugin != "" {
			pluginC = C.CString(plugin)
			defer C.free(unsafe.Pointer(pluginC))
		}
		var errorMessageC *C.char
		if errorMessage != "" {
			errorMessageC = C.CString(errorMessage)
			defer C.free(unsafe.Pointer(errorMessageC))
		}
		return func(l C.struct_QtSerialBus_PackedList) []*QCanBusDeviceInfo {
			var out = make([]*QCanBusDeviceInfo, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQCanBusFromPointer(l.data).__availableDevices_atList(i)
			}
			return out
		}(C.QCanBus_AvailableDevices(ptr.Pointer(), C.struct_QtSerialBus_PackedString{data: pluginC, len: C.longlong(len(plugin))}, C.struct_QtSerialBus_PackedString{data: errorMessageC, len: C.longlong(len(errorMessage))}))
	}
	return make([]*QCanBusDeviceInfo, 0)
}

func (ptr *QCanBus) Plugins() []string {
	if ptr.Pointer() != nil {
		return strings.Split(cGoUnpackString(C.QCanBus_Plugins(ptr.Pointer())), "|")
	}
	return make([]string, 0)
}

func (ptr *QCanBus) __availableDevices_atList(i int) *QCanBusDeviceInfo {
	if ptr.Pointer() != nil {
		var tmpValue = NewQCanBusDeviceInfoFromPointer(C.QCanBus___availableDevices_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QCanBusDeviceInfo).DestroyQCanBusDeviceInfo)
		return tmpValue
	}
	return nil
}

func (ptr *QCanBus) __availableDevices_setList(i QCanBusDeviceInfo_ITF) {
	if ptr.Pointer() != nil {
		C.QCanBus___availableDevices_setList(ptr.Pointer(), PointerFromQCanBusDeviceInfo(i))
	}
}

func (ptr *QCanBus) __availableDevices_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QCanBus___availableDevices_newList(ptr.Pointer()))
}

func (ptr *QCanBus) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QCanBus___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QCanBus) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QCanBus___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QCanBus) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QCanBus___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QCanBus) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QCanBus___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QCanBus) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QCanBus___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QCanBus) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QCanBus___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QCanBus) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QCanBus___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QCanBus) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QCanBus___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QCanBus) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.QCanBus___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QCanBus) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QCanBus___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QCanBus) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QCanBus___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QCanBus) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QCanBus___findChildren_newList(ptr.Pointer()))
}

func (ptr *QCanBus) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QCanBus___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QCanBus) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QCanBus___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QCanBus) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QCanBus___children_newList(ptr.Pointer()))
}

//export callbackQCanBus_Event
func callbackQCanBus_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQCanBusFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QCanBus) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QCanBus_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQCanBus_EventFilter
func callbackQCanBus_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQCanBusFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QCanBus) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QCanBus_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQCanBus_ChildEvent
func callbackQCanBus_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCanBusFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCanBus) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QCanBus_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQCanBus_ConnectNotify
func callbackQCanBus_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQCanBusFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QCanBus) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QCanBus_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQCanBus_CustomEvent
func callbackQCanBus_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCanBusFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCanBus) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QCanBus_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQCanBus_DeleteLater
func callbackQCanBus_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQCanBusFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QCanBus) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QCanBus_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQCanBus_Destroyed
func callbackQCanBus_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQCanBus_DisconnectNotify
func callbackQCanBus_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQCanBusFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QCanBus) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QCanBus_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQCanBus_ObjectNameChanged
func callbackQCanBus_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtSerialBus_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQCanBus_TimerEvent
func callbackQCanBus_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQCanBusFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QCanBus) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QCanBus_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQCanBus_MetaObject
func callbackQCanBus_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQCanBusFromPointer(ptr).MetaObjectDefault())
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
	QCanBusDevice__DataBitRateKey QCanBusDevice__ConfigurationKey = QCanBusDevice__ConfigurationKey(6)
	QCanBusDevice__UserKey        QCanBusDevice__ConfigurationKey = QCanBusDevice__ConfigurationKey(30)
)

func NewQCanBusDevice(parent core.QObject_ITF) *QCanBusDevice {
	var tmpValue = NewQCanBusDeviceFromPointer(C.QCanBusDevice_NewQCanBusDevice(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQCanBusDevice_InterpretErrorFrame
func callbackQCanBusDevice_InterpretErrorFrame(ptr unsafe.Pointer, frame unsafe.Pointer) C.struct_QtSerialBus_PackedString {
	if signal := qt.GetSignal(ptr, "interpretErrorFrame"); signal != nil {
		tempVal := signal.(func(*QCanBusFrame) string)(NewQCanBusFrameFromPointer(frame))
		return C.struct_QtSerialBus_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_QtSerialBus_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QCanBusDevice) ConnectInterpretErrorFrame(f func(frame *QCanBusFrame) string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "interpretErrorFrame"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "interpretErrorFrame", func(frame *QCanBusFrame) string {
				signal.(func(*QCanBusFrame) string)(frame)
				return f(frame)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "interpretErrorFrame", f)
		}
	}
}

func (ptr *QCanBusDevice) DisconnectInterpretErrorFrame() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "interpretErrorFrame")
	}
}

func (ptr *QCanBusDevice) InterpretErrorFrame(frame QCanBusFrame_ITF) string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QCanBusDevice_InterpretErrorFrame(ptr.Pointer(), PointerFromQCanBusFrame(frame)))
	}
	return ""
}

func (ptr *QCanBusDevice) ConnectDevice() bool {
	if ptr.Pointer() != nil {
		return C.QCanBusDevice_ConnectDevice(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQCanBusDevice_Open
func callbackQCanBusDevice_Open(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "open"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QCanBusDevice) ConnectOpen(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "open"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "open", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "open", f)
		}
	}
}

func (ptr *QCanBusDevice) DisconnectOpen() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "open")
	}
}

func (ptr *QCanBusDevice) Open() bool {
	if ptr.Pointer() != nil {
		return C.QCanBusDevice_Open(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQCanBusDevice_WaitForFramesReceived
func callbackQCanBusDevice_WaitForFramesReceived(ptr unsafe.Pointer, msecs C.int) C.char {
	if signal := qt.GetSignal(ptr, "waitForFramesReceived"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQCanBusDeviceFromPointer(ptr).WaitForFramesReceivedDefault(int(int32(msecs))))))
}

func (ptr *QCanBusDevice) ConnectWaitForFramesReceived(f func(msecs int) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "waitForFramesReceived"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "waitForFramesReceived", func(msecs int) bool {
				signal.(func(int) bool)(msecs)
				return f(msecs)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "waitForFramesReceived", f)
		}
	}
}

func (ptr *QCanBusDevice) DisconnectWaitForFramesReceived() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "waitForFramesReceived")
	}
}

func (ptr *QCanBusDevice) WaitForFramesReceived(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QCanBusDevice_WaitForFramesReceived(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

func (ptr *QCanBusDevice) WaitForFramesReceivedDefault(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QCanBusDevice_WaitForFramesReceivedDefault(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

//export callbackQCanBusDevice_WaitForFramesWritten
func callbackQCanBusDevice_WaitForFramesWritten(ptr unsafe.Pointer, msecs C.int) C.char {
	if signal := qt.GetSignal(ptr, "waitForFramesWritten"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQCanBusDeviceFromPointer(ptr).WaitForFramesWrittenDefault(int(int32(msecs))))))
}

func (ptr *QCanBusDevice) ConnectWaitForFramesWritten(f func(msecs int) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "waitForFramesWritten"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "waitForFramesWritten", func(msecs int) bool {
				signal.(func(int) bool)(msecs)
				return f(msecs)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "waitForFramesWritten", f)
		}
	}
}

func (ptr *QCanBusDevice) DisconnectWaitForFramesWritten() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "waitForFramesWritten")
	}
}

func (ptr *QCanBusDevice) WaitForFramesWritten(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QCanBusDevice_WaitForFramesWritten(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

func (ptr *QCanBusDevice) WaitForFramesWrittenDefault(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QCanBusDevice_WaitForFramesWrittenDefault(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

//export callbackQCanBusDevice_WriteFrame
func callbackQCanBusDevice_WriteFrame(ptr unsafe.Pointer, frame unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "writeFrame"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QCanBusFrame) bool)(NewQCanBusFrameFromPointer(frame)))))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QCanBusDevice) ConnectWriteFrame(f func(frame *QCanBusFrame) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "writeFrame"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "writeFrame", func(frame *QCanBusFrame) bool {
				signal.(func(*QCanBusFrame) bool)(frame)
				return f(frame)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "writeFrame", f)
		}
	}
}

func (ptr *QCanBusDevice) DisconnectWriteFrame() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "writeFrame")
	}
}

func (ptr *QCanBusDevice) WriteFrame(frame QCanBusFrame_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QCanBusDevice_WriteFrame(ptr.Pointer(), PointerFromQCanBusFrame(frame)) != 0
	}
	return false
}

//export callbackQCanBusDevice_Close
func callbackQCanBusDevice_Close(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QCanBusDevice) ConnectClose(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "close"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "close", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "close", f)
		}
	}
}

func (ptr *QCanBusDevice) DisconnectClose() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "close")
	}
}

func (ptr *QCanBusDevice) Close() {
	if ptr.Pointer() != nil {
		C.QCanBusDevice_Close(ptr.Pointer())
	}
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

func (ptr *QCanBusDevice) EnqueueReceivedFrames(newFrames []*QCanBusFrame) {
	if ptr.Pointer() != nil {
		C.QCanBusDevice_EnqueueReceivedFrames(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewQCanBusDeviceFromPointer(NewQCanBusDeviceFromPointer(nil).__enqueueReceivedFrames_newFrames_newList())
			for _, v := range newFrames {
				tmpList.__enqueueReceivedFrames_newFrames_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackQCanBusDevice_ErrorOccurred
func callbackQCanBusDevice_ErrorOccurred(ptr unsafe.Pointer, error C.longlong) {
	if signal := qt.GetSignal(ptr, "errorOccurred"); signal != nil {
		signal.(func(QCanBusDevice__CanBusError))(QCanBusDevice__CanBusError(error))
	}

}

func (ptr *QCanBusDevice) ConnectErrorOccurred(f func(error QCanBusDevice__CanBusError)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "errorOccurred") {
			C.QCanBusDevice_ConnectErrorOccurred(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "errorOccurred"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "errorOccurred", func(error QCanBusDevice__CanBusError) {
				signal.(func(QCanBusDevice__CanBusError))(error)
				f(error)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "errorOccurred", f)
		}
	}
}

func (ptr *QCanBusDevice) DisconnectErrorOccurred() {
	if ptr.Pointer() != nil {
		C.QCanBusDevice_DisconnectErrorOccurred(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "errorOccurred")
	}
}

func (ptr *QCanBusDevice) ErrorOccurred(error QCanBusDevice__CanBusError) {
	if ptr.Pointer() != nil {
		C.QCanBusDevice_ErrorOccurred(ptr.Pointer(), C.longlong(error))
	}
}

//export callbackQCanBusDevice_FramesReceived
func callbackQCanBusDevice_FramesReceived(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "framesReceived"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QCanBusDevice) ConnectFramesReceived(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "framesReceived") {
			C.QCanBusDevice_ConnectFramesReceived(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "framesReceived"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "framesReceived", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "framesReceived", f)
		}
	}
}

func (ptr *QCanBusDevice) DisconnectFramesReceived() {
	if ptr.Pointer() != nil {
		C.QCanBusDevice_DisconnectFramesReceived(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "framesReceived")
	}
}

func (ptr *QCanBusDevice) FramesReceived() {
	if ptr.Pointer() != nil {
		C.QCanBusDevice_FramesReceived(ptr.Pointer())
	}
}

//export callbackQCanBusDevice_FramesWritten
func callbackQCanBusDevice_FramesWritten(ptr unsafe.Pointer, framesCount C.longlong) {
	if signal := qt.GetSignal(ptr, "framesWritten"); signal != nil {
		signal.(func(int64))(int64(framesCount))
	}

}

func (ptr *QCanBusDevice) ConnectFramesWritten(f func(framesCount int64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "framesWritten") {
			C.QCanBusDevice_ConnectFramesWritten(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "framesWritten"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "framesWritten", func(framesCount int64) {
				signal.(func(int64))(framesCount)
				f(framesCount)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "framesWritten", f)
		}
	}
}

func (ptr *QCanBusDevice) DisconnectFramesWritten() {
	if ptr.Pointer() != nil {
		C.QCanBusDevice_DisconnectFramesWritten(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "framesWritten")
	}
}

func (ptr *QCanBusDevice) FramesWritten(framesCount int64) {
	if ptr.Pointer() != nil {
		C.QCanBusDevice_FramesWritten(ptr.Pointer(), C.longlong(framesCount))
	}
}

//export callbackQCanBusDevice_SetConfigurationParameter
func callbackQCanBusDevice_SetConfigurationParameter(ptr unsafe.Pointer, key C.int, value unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "setConfigurationParameter"); signal != nil {
		signal.(func(int, *core.QVariant))(int(int32(key)), core.NewQVariantFromPointer(value))
	} else {
		NewQCanBusDeviceFromPointer(ptr).SetConfigurationParameterDefault(int(int32(key)), core.NewQVariantFromPointer(value))
	}
}

func (ptr *QCanBusDevice) ConnectSetConfigurationParameter(f func(key int, value *core.QVariant)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setConfigurationParameter"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setConfigurationParameter", func(key int, value *core.QVariant) {
				signal.(func(int, *core.QVariant))(key, value)
				f(key, value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setConfigurationParameter", f)
		}
	}
}

func (ptr *QCanBusDevice) DisconnectSetConfigurationParameter() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setConfigurationParameter")
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
		var errorTextC *C.char
		if errorText != "" {
			errorTextC = C.CString(errorText)
			defer C.free(unsafe.Pointer(errorTextC))
		}
		C.QCanBusDevice_SetError(ptr.Pointer(), C.struct_QtSerialBus_PackedString{data: errorTextC, len: C.longlong(len(errorText))}, C.longlong(errorId))
	}
}

func (ptr *QCanBusDevice) SetState(newState QCanBusDevice__CanBusDeviceState) {
	if ptr.Pointer() != nil {
		C.QCanBusDevice_SetState(ptr.Pointer(), C.longlong(newState))
	}
}

//export callbackQCanBusDevice_StateChanged
func callbackQCanBusDevice_StateChanged(ptr unsafe.Pointer, state C.longlong) {
	if signal := qt.GetSignal(ptr, "stateChanged"); signal != nil {
		signal.(func(QCanBusDevice__CanBusDeviceState))(QCanBusDevice__CanBusDeviceState(state))
	}

}

func (ptr *QCanBusDevice) ConnectStateChanged(f func(state QCanBusDevice__CanBusDeviceState)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "stateChanged") {
			C.QCanBusDevice_ConnectStateChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "stateChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "stateChanged", func(state QCanBusDevice__CanBusDeviceState) {
				signal.(func(QCanBusDevice__CanBusDeviceState))(state)
				f(state)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stateChanged", f)
		}
	}
}

func (ptr *QCanBusDevice) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QCanBusDevice_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "stateChanged")
	}
}

func (ptr *QCanBusDevice) StateChanged(state QCanBusDevice__CanBusDeviceState) {
	if ptr.Pointer() != nil {
		C.QCanBusDevice_StateChanged(ptr.Pointer(), C.longlong(state))
	}
}

func (ptr *QCanBusDevice) State() QCanBusDevice__CanBusDeviceState {
	if ptr.Pointer() != nil {
		return QCanBusDevice__CanBusDeviceState(C.QCanBusDevice_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCanBusDevice) Error() QCanBusDevice__CanBusError {
	if ptr.Pointer() != nil {
		return QCanBusDevice__CanBusError(C.QCanBusDevice_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCanBusDevice) ErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QCanBusDevice_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QCanBusDevice) ConfigurationParameter(key int) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QCanBusDevice_ConfigurationParameter(ptr.Pointer(), C.int(int32(key))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QCanBusDevice) ConfigurationKeys() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtSerialBus_PackedList) []int {
			var out = make([]int, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQCanBusDeviceFromPointer(l.data).__configurationKeys_atList(i)
			}
			return out
		}(C.QCanBusDevice_ConfigurationKeys(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *QCanBusDevice) HasOutgoingFrames() bool {
	if ptr.Pointer() != nil {
		return C.QCanBusDevice_HasOutgoingFrames(ptr.Pointer()) != 0
	}
	return false
}

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

func (ptr *QCanBusDevice) __enqueueReceivedFrames_newFrames_setList(i QCanBusFrame_ITF) {
	if ptr.Pointer() != nil {
		C.QCanBusDevice___enqueueReceivedFrames_newFrames_setList(ptr.Pointer(), PointerFromQCanBusFrame(i))
	}
}

func (ptr *QCanBusDevice) __enqueueReceivedFrames_newFrames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QCanBusDevice___enqueueReceivedFrames_newFrames_newList(ptr.Pointer()))
}

func (ptr *QCanBusDevice) __configurationKeys_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QCanBusDevice___configurationKeys_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QCanBusDevice) __configurationKeys_setList(i int) {
	if ptr.Pointer() != nil {
		C.QCanBusDevice___configurationKeys_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QCanBusDevice) __configurationKeys_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QCanBusDevice___configurationKeys_newList(ptr.Pointer()))
}

func (ptr *QCanBusDevice) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QCanBusDevice___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QCanBusDevice) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QCanBusDevice___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QCanBusDevice) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QCanBusDevice___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QCanBusDevice) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QCanBusDevice___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QCanBusDevice) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QCanBusDevice___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QCanBusDevice) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QCanBusDevice___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QCanBusDevice) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QCanBusDevice___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QCanBusDevice) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QCanBusDevice___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QCanBusDevice) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.QCanBusDevice___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QCanBusDevice) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QCanBusDevice___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QCanBusDevice) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QCanBusDevice___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QCanBusDevice) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QCanBusDevice___findChildren_newList(ptr.Pointer()))
}

func (ptr *QCanBusDevice) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QCanBusDevice___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QCanBusDevice) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QCanBusDevice___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QCanBusDevice) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QCanBusDevice___children_newList(ptr.Pointer()))
}

//export callbackQCanBusDevice_Event
func callbackQCanBusDevice_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQCanBusDeviceFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QCanBusDevice) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QCanBusDevice_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQCanBusDevice_EventFilter
func callbackQCanBusDevice_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQCanBusDeviceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QCanBusDevice) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QCanBusDevice_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQCanBusDevice_ChildEvent
func callbackQCanBusDevice_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQCanBusDeviceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QCanBusDevice) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QCanBusDevice_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQCanBusDevice_ConnectNotify
func callbackQCanBusDevice_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQCanBusDeviceFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QCanBusDevice) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QCanBusDevice_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQCanBusDevice_CustomEvent
func callbackQCanBusDevice_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQCanBusDeviceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QCanBusDevice) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QCanBusDevice_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQCanBusDevice_DeleteLater
func callbackQCanBusDevice_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQCanBusDeviceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QCanBusDevice) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QCanBusDevice_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQCanBusDevice_Destroyed
func callbackQCanBusDevice_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQCanBusDevice_DisconnectNotify
func callbackQCanBusDevice_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQCanBusDeviceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QCanBusDevice) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QCanBusDevice_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQCanBusDevice_ObjectNameChanged
func callbackQCanBusDevice_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtSerialBus_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQCanBusDevice_TimerEvent
func callbackQCanBusDevice_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQCanBusDeviceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QCanBusDevice) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QCanBusDevice_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQCanBusDevice_MetaObject
func callbackQCanBusDevice_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQCanBusDeviceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QCanBusDevice) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QCanBusDevice_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QCanBusDeviceInfo struct {
	ptr unsafe.Pointer
}

type QCanBusDeviceInfo_ITF interface {
	QCanBusDeviceInfo_PTR() *QCanBusDeviceInfo
}

func (ptr *QCanBusDeviceInfo) QCanBusDeviceInfo_PTR() *QCanBusDeviceInfo {
	return ptr
}

func (ptr *QCanBusDeviceInfo) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QCanBusDeviceInfo) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQCanBusDeviceInfo(ptr QCanBusDeviceInfo_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCanBusDeviceInfo_PTR().Pointer()
	}
	return nil
}

func NewQCanBusDeviceInfoFromPointer(ptr unsafe.Pointer) *QCanBusDeviceInfo {
	var n = new(QCanBusDeviceInfo)
	n.SetPointer(ptr)
	return n
}
func (ptr *QCanBusDeviceInfo) Swap(other QCanBusDeviceInfo_ITF) {
	if ptr.Pointer() != nil {
		C.QCanBusDeviceInfo_Swap(ptr.Pointer(), PointerFromQCanBusDeviceInfo(other))
	}
}

func (ptr *QCanBusDeviceInfo) DestroyQCanBusDeviceInfo() {
	if ptr.Pointer() != nil {
		C.QCanBusDeviceInfo_DestroyQCanBusDeviceInfo(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QCanBusDeviceInfo) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QCanBusDeviceInfo_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QCanBusDeviceInfo) HasFlexibleDataRate() bool {
	if ptr.Pointer() != nil {
		return C.QCanBusDeviceInfo_HasFlexibleDataRate(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCanBusDeviceInfo) IsVirtual() bool {
	if ptr.Pointer() != nil {
		return C.QCanBusDeviceInfo_IsVirtual(ptr.Pointer()) != 0
	}
	return false
}

type QCanBusFactoryV2 struct {
	ptr unsafe.Pointer
}

type QCanBusFactoryV2_ITF interface {
	QCanBusFactoryV2_PTR() *QCanBusFactoryV2
}

func (ptr *QCanBusFactoryV2) QCanBusFactoryV2_PTR() *QCanBusFactoryV2 {
	return ptr
}

func (ptr *QCanBusFactoryV2) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QCanBusFactoryV2) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQCanBusFactoryV2(ptr QCanBusFactoryV2_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QCanBusFactoryV2_PTR().Pointer()
	}
	return nil
}

func NewQCanBusFactoryV2FromPointer(ptr unsafe.Pointer) *QCanBusFactoryV2 {
	var n = new(QCanBusFactoryV2)
	n.SetPointer(ptr)
	return n
}

func (ptr *QCanBusFactoryV2) DestroyQCanBusFactoryV2() {
	if ptr != nil {
		C.free(ptr.Pointer())
		qt.DisconnectAllSignals(ptr.Pointer(), "")
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQCanBusFactoryV2_CreateDevice
func callbackQCanBusFactoryV2_CreateDevice(ptr unsafe.Pointer, interfaceName C.struct_QtSerialBus_PackedString, errorMessage C.struct_QtSerialBus_PackedString) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "createDevice"); signal != nil {
		return PointerFromQCanBusDevice(signal.(func(string, string) *QCanBusDevice)(cGoUnpackString(interfaceName), cGoUnpackString(errorMessage)))
	}

	return PointerFromQCanBusDevice(NewQCanBusDevice(nil))
}

func (ptr *QCanBusFactoryV2) ConnectCreateDevice(f func(interfaceName string, errorMessage string) *QCanBusDevice) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "createDevice"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "createDevice", func(interfaceName string, errorMessage string) *QCanBusDevice {
				signal.(func(string, string) *QCanBusDevice)(interfaceName, errorMessage)
				return f(interfaceName, errorMessage)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "createDevice", f)
		}
	}
}

func (ptr *QCanBusFactoryV2) DisconnectCreateDevice() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "createDevice")
	}
}

func (ptr *QCanBusFactoryV2) CreateDevice(interfaceName string, errorMessage string) *QCanBusDevice {
	if ptr.Pointer() != nil {
		var interfaceNameC *C.char
		if interfaceName != "" {
			interfaceNameC = C.CString(interfaceName)
			defer C.free(unsafe.Pointer(interfaceNameC))
		}
		var errorMessageC *C.char
		if errorMessage != "" {
			errorMessageC = C.CString(errorMessage)
			defer C.free(unsafe.Pointer(errorMessageC))
		}
		var tmpValue = NewQCanBusDeviceFromPointer(C.QCanBusFactoryV2_CreateDevice(ptr.Pointer(), C.struct_QtSerialBus_PackedString{data: interfaceNameC, len: C.longlong(len(interfaceName))}, C.struct_QtSerialBus_PackedString{data: errorMessageC, len: C.longlong(len(errorMessage))}))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQCanBusFactoryV2_AvailableDevices
func callbackQCanBusFactoryV2_AvailableDevices(ptr unsafe.Pointer, errorMessage C.struct_QtSerialBus_PackedString) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "availableDevices"); signal != nil {
		return func() unsafe.Pointer {
			var tmpList = NewQCanBusFactoryV2FromPointer(NewQCanBusFactoryV2FromPointer(nil).__availableDevices_newList())
			for _, v := range signal.(func(string) []*QCanBusDeviceInfo)(cGoUnpackString(errorMessage)) {
				tmpList.__availableDevices_setList(v)
			}
			return tmpList.Pointer()
		}()
	}

	return func() unsafe.Pointer {
		var tmpList = NewQCanBusFactoryV2FromPointer(NewQCanBusFactoryV2FromPointer(nil).__availableDevices_newList())
		for _, v := range make([]*QCanBusDeviceInfo, 0) {
			tmpList.__availableDevices_setList(v)
		}
		return tmpList.Pointer()
	}()
}

func (ptr *QCanBusFactoryV2) ConnectAvailableDevices(f func(errorMessage string) []*QCanBusDeviceInfo) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "availableDevices"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "availableDevices", func(errorMessage string) []*QCanBusDeviceInfo {
				signal.(func(string) []*QCanBusDeviceInfo)(errorMessage)
				return f(errorMessage)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "availableDevices", f)
		}
	}
}

func (ptr *QCanBusFactoryV2) DisconnectAvailableDevices() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "availableDevices")
	}
}

func (ptr *QCanBusFactoryV2) AvailableDevices(errorMessage string) []*QCanBusDeviceInfo {
	if ptr.Pointer() != nil {
		var errorMessageC *C.char
		if errorMessage != "" {
			errorMessageC = C.CString(errorMessage)
			defer C.free(unsafe.Pointer(errorMessageC))
		}
		return func(l C.struct_QtSerialBus_PackedList) []*QCanBusDeviceInfo {
			var out = make([]*QCanBusDeviceInfo, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQCanBusFactoryV2FromPointer(l.data).__availableDevices_atList(i)
			}
			return out
		}(C.QCanBusFactoryV2_AvailableDevices(ptr.Pointer(), C.struct_QtSerialBus_PackedString{data: errorMessageC, len: C.longlong(len(errorMessage))}))
	}
	return make([]*QCanBusDeviceInfo, 0)
}

func (ptr *QCanBusFactoryV2) __availableDevices_atList(i int) *QCanBusDeviceInfo {
	if ptr.Pointer() != nil {
		var tmpValue = NewQCanBusDeviceInfoFromPointer(C.QCanBusFactoryV2___availableDevices_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QCanBusDeviceInfo).DestroyQCanBusDeviceInfo)
		return tmpValue
	}
	return nil
}

func (ptr *QCanBusFactoryV2) __availableDevices_setList(i QCanBusDeviceInfo_ITF) {
	if ptr.Pointer() != nil {
		C.QCanBusFactoryV2___availableDevices_setList(ptr.Pointer(), PointerFromQCanBusDeviceInfo(i))
	}
}

func (ptr *QCanBusFactoryV2) __availableDevices_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QCanBusFactoryV2___availableDevices_newList(ptr.Pointer()))
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
		runtime.SetFinalizer(ptr, nil)
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

func (ptr *QCanBusFrame) SetBitrateSwitch(bitrateSwitch bool) {
	if ptr.Pointer() != nil {
		C.QCanBusFrame_SetBitrateSwitch(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(bitrateSwitch))))
	}
}

func (ptr *QCanBusFrame) SetError(error QCanBusFrame__FrameError) {
	if ptr.Pointer() != nil {
		C.QCanBusFrame_SetError(ptr.Pointer(), C.longlong(error))
	}
}

func (ptr *QCanBusFrame) SetErrorStateIndicator(errorStateIndicator bool) {
	if ptr.Pointer() != nil {
		C.QCanBusFrame_SetErrorStateIndicator(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(errorStateIndicator))))
	}
}

func (ptr *QCanBusFrame) SetExtendedFrameFormat(isExtended bool) {
	if ptr.Pointer() != nil {
		C.QCanBusFrame_SetExtendedFrameFormat(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(isExtended))))
	}
}

func (ptr *QCanBusFrame) SetFlexibleDataRateFormat(isFlexibleData bool) {
	if ptr.Pointer() != nil {
		C.QCanBusFrame_SetFlexibleDataRateFormat(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(isFlexibleData))))
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

func (ptr *QCanBusFrame) Error() QCanBusFrame__FrameError {
	if ptr.Pointer() != nil {
		return QCanBusFrame__FrameError(C.QCanBusFrame_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCanBusFrame) FrameType() QCanBusFrame__FrameType {
	if ptr.Pointer() != nil {
		return QCanBusFrame__FrameType(C.QCanBusFrame_FrameType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QCanBusFrame) Payload() *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QCanBusFrame_Payload(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QCanBusFrame) ToString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QCanBusFrame_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QCanBusFrame) HasBitrateSwitch() bool {
	if ptr.Pointer() != nil {
		return C.QCanBusFrame_HasBitrateSwitch(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCanBusFrame) HasErrorStateIndicator() bool {
	if ptr.Pointer() != nil {
		return C.QCanBusFrame_HasErrorStateIndicator(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCanBusFrame) HasExtendedFrameFormat() bool {
	if ptr.Pointer() != nil {
		return C.QCanBusFrame_HasExtendedFrameFormat(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCanBusFrame) HasFlexibleDataRateFormat() bool {
	if ptr.Pointer() != nil {
		return C.QCanBusFrame_HasFlexibleDataRateFormat(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCanBusFrame) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QCanBusFrame_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QCanBusFrame) FrameId() uint {
	if ptr.Pointer() != nil {
		return uint(uint32(C.QCanBusFrame_FrameId(ptr.Pointer())))
	}
	return 0
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
func NewQModbusClient(parent core.QObject_ITF) *QModbusClient {
	var tmpValue = NewQModbusClientFromPointer(C.QModbusClient_NewQModbusClient(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QModbusClient) SendRawRequest(request QModbusRequest_ITF, serverAddress int) *QModbusReply {
	if ptr.Pointer() != nil {
		var tmpValue = NewQModbusReplyFromPointer(C.QModbusClient_SendRawRequest(ptr.Pointer(), PointerFromQModbusRequest(request), C.int(int32(serverAddress))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QModbusClient) SendReadRequest(read QModbusDataUnit_ITF, serverAddress int) *QModbusReply {
	if ptr.Pointer() != nil {
		var tmpValue = NewQModbusReplyFromPointer(C.QModbusClient_SendReadRequest(ptr.Pointer(), PointerFromQModbusDataUnit(read), C.int(int32(serverAddress))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QModbusClient) SendReadWriteRequest(read QModbusDataUnit_ITF, write QModbusDataUnit_ITF, serverAddress int) *QModbusReply {
	if ptr.Pointer() != nil {
		var tmpValue = NewQModbusReplyFromPointer(C.QModbusClient_SendReadWriteRequest(ptr.Pointer(), PointerFromQModbusDataUnit(read), PointerFromQModbusDataUnit(write), C.int(int32(serverAddress))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QModbusClient) SendWriteRequest(write QModbusDataUnit_ITF, serverAddress int) *QModbusReply {
	if ptr.Pointer() != nil {
		var tmpValue = NewQModbusReplyFromPointer(C.QModbusClient_SendWriteRequest(ptr.Pointer(), PointerFromQModbusDataUnit(write), C.int(int32(serverAddress))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQModbusClient_ProcessPrivateResponse
func callbackQModbusClient_ProcessPrivateResponse(ptr unsafe.Pointer, response unsafe.Pointer, data unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "processPrivateResponse"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QModbusResponse, *QModbusDataUnit) bool)(NewQModbusResponseFromPointer(response), NewQModbusDataUnitFromPointer(data)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusClientFromPointer(ptr).ProcessPrivateResponseDefault(NewQModbusResponseFromPointer(response), NewQModbusDataUnitFromPointer(data)))))
}

func (ptr *QModbusClient) ConnectProcessPrivateResponse(f func(response *QModbusResponse, data *QModbusDataUnit) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "processPrivateResponse"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "processPrivateResponse", func(response *QModbusResponse, data *QModbusDataUnit) bool {
				signal.(func(*QModbusResponse, *QModbusDataUnit) bool)(response, data)
				return f(response, data)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "processPrivateResponse", f)
		}
	}
}

func (ptr *QModbusClient) DisconnectProcessPrivateResponse() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "processPrivateResponse")
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
	if signal := qt.GetSignal(ptr, "processResponse"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QModbusResponse, *QModbusDataUnit) bool)(NewQModbusResponseFromPointer(response), NewQModbusDataUnitFromPointer(data)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusClientFromPointer(ptr).ProcessResponseDefault(NewQModbusResponseFromPointer(response), NewQModbusDataUnitFromPointer(data)))))
}

func (ptr *QModbusClient) ConnectProcessResponse(f func(response *QModbusResponse, data *QModbusDataUnit) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "processResponse"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "processResponse", func(response *QModbusResponse, data *QModbusDataUnit) bool {
				signal.(func(*QModbusResponse, *QModbusDataUnit) bool)(response, data)
				return f(response, data)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "processResponse", f)
		}
	}
}

func (ptr *QModbusClient) DisconnectProcessResponse() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "processResponse")
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
	if signal := qt.GetSignal(ptr, "timeoutChanged"); signal != nil {
		signal.(func(int))(int(int32(newTimeout)))
	}

}

func (ptr *QModbusClient) ConnectTimeoutChanged(f func(newTimeout int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "timeoutChanged") {
			C.QModbusClient_ConnectTimeoutChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "timeoutChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "timeoutChanged", func(newTimeout int) {
				signal.(func(int))(newTimeout)
				f(newTimeout)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "timeoutChanged", f)
		}
	}
}

func (ptr *QModbusClient) DisconnectTimeoutChanged() {
	if ptr.Pointer() != nil {
		C.QModbusClient_DisconnectTimeoutChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "timeoutChanged")
	}
}

func (ptr *QModbusClient) TimeoutChanged(newTimeout int) {
	if ptr.Pointer() != nil {
		C.QModbusClient_TimeoutChanged(ptr.Pointer(), C.int(int32(newTimeout)))
	}
}

func (ptr *QModbusClient) NumberOfRetries() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QModbusClient_NumberOfRetries(ptr.Pointer())))
	}
	return 0
}

func (ptr *QModbusClient) Timeout() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QModbusClient_Timeout(ptr.Pointer())))
	}
	return 0
}

//export callbackQModbusClient_Open
func callbackQModbusClient_Open(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "open"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusClientFromPointer(ptr).OpenDefault())))
}

func (ptr *QModbusClient) Open() bool {
	if ptr.Pointer() != nil {
		return C.QModbusClient_Open(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QModbusClient) OpenDefault() bool {
	if ptr.Pointer() != nil {
		return C.QModbusClient_OpenDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQModbusClient_Close
func callbackQModbusClient_Close(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		signal.(func())()
	} else {
		NewQModbusClientFromPointer(ptr).CloseDefault()
	}
}

func (ptr *QModbusClient) Close() {
	if ptr.Pointer() != nil {
		C.QModbusClient_Close(ptr.Pointer())
	}
}

func (ptr *QModbusClient) CloseDefault() {
	if ptr.Pointer() != nil {
		C.QModbusClient_CloseDefault(ptr.Pointer())
	}
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
		runtime.SetFinalizer(ptr, nil)
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

func NewQModbusDataUnit4(ty QModbusDataUnit__RegisterType, address int, data []uint16) *QModbusDataUnit {
	var tmpValue = NewQModbusDataUnitFromPointer(C.QModbusDataUnit_NewQModbusDataUnit4(C.longlong(ty), C.int(int32(address)), func() unsafe.Pointer {
		var tmpList = NewQModbusDataUnitFromPointer(NewQModbusDataUnitFromPointer(nil).__QModbusDataUnit_data_newList4())
		for _, v := range data {
			tmpList.__QModbusDataUnit_data_setList4(v)
		}
		return tmpList.Pointer()
	}()))
	runtime.SetFinalizer(tmpValue, (*QModbusDataUnit).DestroyQModbusDataUnit)
	return tmpValue
}

func NewQModbusDataUnit3(ty QModbusDataUnit__RegisterType, address int, size uint16) *QModbusDataUnit {
	var tmpValue = NewQModbusDataUnitFromPointer(C.QModbusDataUnit_NewQModbusDataUnit3(C.longlong(ty), C.int(int32(address)), C.ushort(size)))
	runtime.SetFinalizer(tmpValue, (*QModbusDataUnit).DestroyQModbusDataUnit)
	return tmpValue
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

func (ptr *QModbusDataUnit) SetValues(values []uint16) {
	if ptr.Pointer() != nil {
		C.QModbusDataUnit_SetValues(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewQModbusDataUnitFromPointer(NewQModbusDataUnitFromPointer(nil).__setValues_values_newList())
			for _, v := range values {
				tmpList.__setValues_values_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QModbusDataUnit) Values() []uint16 {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtSerialBus_PackedList) []uint16 {
			var out = make([]uint16, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQModbusDataUnitFromPointer(l.data).__values_atList(i)
			}
			return out
		}(C.QModbusDataUnit_Values(ptr.Pointer()))
	}
	return make([]uint16, 0)
}

func (ptr *QModbusDataUnit) RegisterType() QModbusDataUnit__RegisterType {
	if ptr.Pointer() != nil {
		return QModbusDataUnit__RegisterType(C.QModbusDataUnit_RegisterType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QModbusDataUnit) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QModbusDataUnit_IsValid(ptr.Pointer()) != 0
	}
	return false
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

func (ptr *QModbusDataUnit) __QModbusDataUnit_data_atList4(i int) uint16 {
	if ptr.Pointer() != nil {
		return uint16(C.QModbusDataUnit___QModbusDataUnit_data_atList4(ptr.Pointer(), C.int(int32(i))))
	}
	return 0
}

func (ptr *QModbusDataUnit) __QModbusDataUnit_data_setList4(i uint16) {
	if ptr.Pointer() != nil {
		C.QModbusDataUnit___QModbusDataUnit_data_setList4(ptr.Pointer(), C.ushort(i))
	}
}

func (ptr *QModbusDataUnit) __QModbusDataUnit_data_newList4() unsafe.Pointer {
	return unsafe.Pointer(C.QModbusDataUnit___QModbusDataUnit_data_newList4(ptr.Pointer()))
}

func (ptr *QModbusDataUnit) __setValues_values_atList(i int) uint16 {
	if ptr.Pointer() != nil {
		return uint16(C.QModbusDataUnit___setValues_values_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return 0
}

func (ptr *QModbusDataUnit) __setValues_values_setList(i uint16) {
	if ptr.Pointer() != nil {
		C.QModbusDataUnit___setValues_values_setList(ptr.Pointer(), C.ushort(i))
	}
}

func (ptr *QModbusDataUnit) __setValues_values_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QModbusDataUnit___setValues_values_newList(ptr.Pointer()))
}

func (ptr *QModbusDataUnit) __values_atList(i int) uint16 {
	if ptr.Pointer() != nil {
		return uint16(C.QModbusDataUnit___values_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return 0
}

func (ptr *QModbusDataUnit) __values_setList(i uint16) {
	if ptr.Pointer() != nil {
		C.QModbusDataUnit___values_setList(ptr.Pointer(), C.ushort(i))
	}
}

func (ptr *QModbusDataUnit) __values_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QModbusDataUnit___values_newList(ptr.Pointer()))
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
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QModbusDevice) ConnectDevice() bool {
	if ptr.Pointer() != nil {
		return C.QModbusDevice_ConnectDevice(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQModbusDevice_Open
func callbackQModbusDevice_Open(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "open"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QModbusDevice) ConnectOpen(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "open"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "open", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "open", f)
		}
	}
}

func (ptr *QModbusDevice) DisconnectOpen() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "open")
	}
}

func (ptr *QModbusDevice) Open() bool {
	if ptr.Pointer() != nil {
		return C.QModbusDevice_Open(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQModbusDevice_Close
func callbackQModbusDevice_Close(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QModbusDevice) ConnectClose(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "close"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "close", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "close", f)
		}
	}
}

func (ptr *QModbusDevice) DisconnectClose() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "close")
	}
}

func (ptr *QModbusDevice) Close() {
	if ptr.Pointer() != nil {
		C.QModbusDevice_Close(ptr.Pointer())
	}
}

func (ptr *QModbusDevice) DisconnectDevice() {
	if ptr.Pointer() != nil {
		C.QModbusDevice_DisconnectDevice(ptr.Pointer())
	}
}

//export callbackQModbusDevice_ErrorOccurred
func callbackQModbusDevice_ErrorOccurred(ptr unsafe.Pointer, error C.longlong) {
	if signal := qt.GetSignal(ptr, "errorOccurred"); signal != nil {
		signal.(func(QModbusDevice__Error))(QModbusDevice__Error(error))
	}

}

func (ptr *QModbusDevice) ConnectErrorOccurred(f func(error QModbusDevice__Error)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "errorOccurred") {
			C.QModbusDevice_ConnectErrorOccurred(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "errorOccurred"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "errorOccurred", func(error QModbusDevice__Error) {
				signal.(func(QModbusDevice__Error))(error)
				f(error)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "errorOccurred", f)
		}
	}
}

func (ptr *QModbusDevice) DisconnectErrorOccurred() {
	if ptr.Pointer() != nil {
		C.QModbusDevice_DisconnectErrorOccurred(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "errorOccurred")
	}
}

func (ptr *QModbusDevice) ErrorOccurred(error QModbusDevice__Error) {
	if ptr.Pointer() != nil {
		C.QModbusDevice_ErrorOccurred(ptr.Pointer(), C.longlong(error))
	}
}

func (ptr *QModbusDevice) SetConnectionParameter(parameter int, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusDevice_SetConnectionParameter(ptr.Pointer(), C.int(int32(parameter)), core.PointerFromQVariant(value))
	}
}

func (ptr *QModbusDevice) SetError(errorText string, error QModbusDevice__Error) {
	if ptr.Pointer() != nil {
		var errorTextC *C.char
		if errorText != "" {
			errorTextC = C.CString(errorText)
			defer C.free(unsafe.Pointer(errorTextC))
		}
		C.QModbusDevice_SetError(ptr.Pointer(), C.struct_QtSerialBus_PackedString{data: errorTextC, len: C.longlong(len(errorText))}, C.longlong(error))
	}
}

func (ptr *QModbusDevice) SetState(newState QModbusDevice__State) {
	if ptr.Pointer() != nil {
		C.QModbusDevice_SetState(ptr.Pointer(), C.longlong(newState))
	}
}

//export callbackQModbusDevice_StateChanged
func callbackQModbusDevice_StateChanged(ptr unsafe.Pointer, state C.longlong) {
	if signal := qt.GetSignal(ptr, "stateChanged"); signal != nil {
		signal.(func(QModbusDevice__State))(QModbusDevice__State(state))
	}

}

func (ptr *QModbusDevice) ConnectStateChanged(f func(state QModbusDevice__State)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "stateChanged") {
			C.QModbusDevice_ConnectStateChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "stateChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "stateChanged", func(state QModbusDevice__State) {
				signal.(func(QModbusDevice__State))(state)
				f(state)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stateChanged", f)
		}
	}
}

func (ptr *QModbusDevice) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QModbusDevice_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "stateChanged")
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
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QModbusDevice) Error() QModbusDevice__Error {
	if ptr.Pointer() != nil {
		return QModbusDevice__Error(C.QModbusDevice_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QModbusDevice) ErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QModbusDevice_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QModbusDevice) ConnectionParameter(parameter int) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QModbusDevice_ConnectionParameter(ptr.Pointer(), C.int(int32(parameter))))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QModbusDevice) State() QModbusDevice__State {
	if ptr.Pointer() != nil {
		return QModbusDevice__State(C.QModbusDevice_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QModbusDevice) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QModbusDevice___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QModbusDevice) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusDevice___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QModbusDevice) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QModbusDevice___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QModbusDevice) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QModbusDevice___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QModbusDevice) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusDevice___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QModbusDevice) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QModbusDevice___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QModbusDevice) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QModbusDevice___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QModbusDevice) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusDevice___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QModbusDevice) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.QModbusDevice___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QModbusDevice) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QModbusDevice___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QModbusDevice) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusDevice___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QModbusDevice) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QModbusDevice___findChildren_newList(ptr.Pointer()))
}

func (ptr *QModbusDevice) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QModbusDevice___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QModbusDevice) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusDevice___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QModbusDevice) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QModbusDevice___children_newList(ptr.Pointer()))
}

//export callbackQModbusDevice_Event
func callbackQModbusDevice_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusDeviceFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QModbusDevice) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusDevice_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQModbusDevice_EventFilter
func callbackQModbusDevice_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusDeviceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QModbusDevice) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusDevice_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQModbusDevice_ChildEvent
func callbackQModbusDevice_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQModbusDeviceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QModbusDevice) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusDevice_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQModbusDevice_ConnectNotify
func callbackQModbusDevice_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQModbusDeviceFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QModbusDevice) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusDevice_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQModbusDevice_CustomEvent
func callbackQModbusDevice_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQModbusDeviceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QModbusDevice) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusDevice_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQModbusDevice_DeleteLater
func callbackQModbusDevice_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQModbusDeviceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QModbusDevice) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QModbusDevice_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQModbusDevice_Destroyed
func callbackQModbusDevice_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQModbusDevice_DisconnectNotify
func callbackQModbusDevice_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQModbusDeviceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QModbusDevice) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusDevice_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQModbusDevice_ObjectNameChanged
func callbackQModbusDevice_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtSerialBus_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQModbusDevice_TimerEvent
func callbackQModbusDevice_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQModbusDeviceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QModbusDevice) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusDevice_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQModbusDevice_MetaObject
func callbackQModbusDevice_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQModbusDeviceFromPointer(ptr).MetaObjectDefault())
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
		runtime.SetFinalizer(ptr, nil)
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

func NewQModbusDeviceIdentification() *QModbusDeviceIdentification {
	var tmpValue = NewQModbusDeviceIdentificationFromPointer(C.QModbusDeviceIdentification_NewQModbusDeviceIdentification())
	runtime.SetFinalizer(tmpValue, (*QModbusDeviceIdentification).DestroyQModbusDeviceIdentification)
	return tmpValue
}

func (ptr *QModbusDeviceIdentification) Insert(objectId uint, value core.QByteArray_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusDeviceIdentification_Insert(ptr.Pointer(), C.uint(uint32(objectId)), core.PointerFromQByteArray(value)) != 0
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

func (ptr *QModbusDeviceIdentification) ConformityLevel() QModbusDeviceIdentification__ConformityLevel {
	if ptr.Pointer() != nil {
		return QModbusDeviceIdentification__ConformityLevel(C.QModbusDeviceIdentification_ConformityLevel(ptr.Pointer()))
	}
	return 0
}

func (ptr *QModbusDeviceIdentification) Value(objectId uint) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QModbusDeviceIdentification_Value(ptr.Pointer(), C.uint(uint32(objectId))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QModbusDeviceIdentification) ObjectIds() []int {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtSerialBus_PackedList) []int {
			var out = make([]int, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQModbusDeviceIdentificationFromPointer(l.data).__objectIds_atList(i)
			}
			return out
		}(C.QModbusDeviceIdentification_ObjectIds(ptr.Pointer()))
	}
	return make([]int, 0)
}

func (ptr *QModbusDeviceIdentification) Contains(objectId uint) bool {
	if ptr.Pointer() != nil {
		return C.QModbusDeviceIdentification_Contains(ptr.Pointer(), C.uint(uint32(objectId))) != 0
	}
	return false
}

func (ptr *QModbusDeviceIdentification) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QModbusDeviceIdentification_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QModbusDeviceIdentification) __objectIds_atList(i int) int {
	if ptr.Pointer() != nil {
		return int(int32(C.QModbusDeviceIdentification___objectIds_atList(ptr.Pointer(), C.int(int32(i)))))
	}
	return 0
}

func (ptr *QModbusDeviceIdentification) __objectIds_setList(i int) {
	if ptr.Pointer() != nil {
		C.QModbusDeviceIdentification___objectIds_setList(ptr.Pointer(), C.int(int32(i)))
	}
}

func (ptr *QModbusDeviceIdentification) __objectIds_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QModbusDeviceIdentification___objectIds_newList(ptr.Pointer()))
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
		qt.DisconnectAllSignals(ptr.Pointer(), "")
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
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

func (ptr *QModbusPdu) SetData(data core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusPdu_SetData(ptr.Pointer(), core.PointerFromQByteArray(data))
	}
}

//export callbackQModbusPdu_SetFunctionCode
func callbackQModbusPdu_SetFunctionCode(ptr unsafe.Pointer, code C.longlong) {
	if signal := qt.GetSignal(ptr, "setFunctionCode"); signal != nil {
		signal.(func(QModbusPdu__FunctionCode))(QModbusPdu__FunctionCode(code))
	} else {
		NewQModbusPduFromPointer(ptr).SetFunctionCodeDefault(QModbusPdu__FunctionCode(code))
	}
}

func (ptr *QModbusPdu) ConnectSetFunctionCode(f func(code QModbusPdu__FunctionCode)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setFunctionCode"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setFunctionCode", func(code QModbusPdu__FunctionCode) {
				signal.(func(QModbusPdu__FunctionCode))(code)
				f(code)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setFunctionCode", f)
		}
	}
}

func (ptr *QModbusPdu) DisconnectSetFunctionCode() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setFunctionCode")
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

//export callbackQModbusPdu_DestroyQModbusPdu
func callbackQModbusPdu_DestroyQModbusPdu(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QModbusPdu"); signal != nil {
		signal.(func())()
	} else {
		NewQModbusPduFromPointer(ptr).DestroyQModbusPduDefault()
	}
}

func (ptr *QModbusPdu) ConnectDestroyQModbusPdu(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QModbusPdu"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QModbusPdu", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QModbusPdu", f)
		}
	}
}

func (ptr *QModbusPdu) DisconnectDestroyQModbusPdu() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QModbusPdu")
	}
}

func (ptr *QModbusPdu) DestroyQModbusPdu() {
	if ptr.Pointer() != nil {
		C.QModbusPdu_DestroyQModbusPdu(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QModbusPdu) DestroyQModbusPduDefault() {
	if ptr.Pointer() != nil {
		C.QModbusPdu_DestroyQModbusPduDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
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

func (ptr *QModbusPdu) Data() *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QModbusPdu_Data(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
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

func (ptr *QModbusPdu) DataSize() int16 {
	if ptr.Pointer() != nil {
		return int16(C.QModbusPdu_DataSize(ptr.Pointer()))
	}
	return 0
}

func (ptr *QModbusPdu) Size() int16 {
	if ptr.Pointer() != nil {
		return int16(C.QModbusPdu_Size(ptr.Pointer()))
	}
	return 0
}

func (ptr *QModbusPdu) __encode_vector_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QModbusPdu___encode_vector_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QModbusPdu) __encode_vector_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusPdu___encode_vector_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QModbusPdu) __encode_vector_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QModbusPdu___encode_vector_newList2(ptr.Pointer()))
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

//go:generate stringer -type=QModbusReply__ReplyType
//QModbusReply::ReplyType
type QModbusReply__ReplyType int64

const (
	QModbusReply__Raw    QModbusReply__ReplyType = QModbusReply__ReplyType(0)
	QModbusReply__Common QModbusReply__ReplyType = QModbusReply__ReplyType(1)
)

func NewQModbusReply(ty QModbusReply__ReplyType, serverAddress int, parent core.QObject_ITF) *QModbusReply {
	var tmpValue = NewQModbusReplyFromPointer(C.QModbusReply_NewQModbusReply(C.longlong(ty), C.int(int32(serverAddress)), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQModbusReply_ErrorOccurred
func callbackQModbusReply_ErrorOccurred(ptr unsafe.Pointer, error C.longlong) {
	if signal := qt.GetSignal(ptr, "errorOccurred"); signal != nil {
		signal.(func(QModbusDevice__Error))(QModbusDevice__Error(error))
	}

}

func (ptr *QModbusReply) ConnectErrorOccurred(f func(error QModbusDevice__Error)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "errorOccurred") {
			C.QModbusReply_ConnectErrorOccurred(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "errorOccurred"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "errorOccurred", func(error QModbusDevice__Error) {
				signal.(func(QModbusDevice__Error))(error)
				f(error)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "errorOccurred", f)
		}
	}
}

func (ptr *QModbusReply) DisconnectErrorOccurred() {
	if ptr.Pointer() != nil {
		C.QModbusReply_DisconnectErrorOccurred(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "errorOccurred")
	}
}

func (ptr *QModbusReply) ErrorOccurred(error QModbusDevice__Error) {
	if ptr.Pointer() != nil {
		C.QModbusReply_ErrorOccurred(ptr.Pointer(), C.longlong(error))
	}
}

//export callbackQModbusReply_Finished
func callbackQModbusReply_Finished(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "finished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QModbusReply) ConnectFinished(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "finished") {
			C.QModbusReply_ConnectFinished(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "finished"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "finished", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "finished", f)
		}
	}
}

func (ptr *QModbusReply) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QModbusReply_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "finished")
	}
}

func (ptr *QModbusReply) Finished() {
	if ptr.Pointer() != nil {
		C.QModbusReply_Finished(ptr.Pointer())
	}
}

func (ptr *QModbusReply) Result() *QModbusDataUnit {
	if ptr.Pointer() != nil {
		var tmpValue = NewQModbusDataUnitFromPointer(C.QModbusReply_Result(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QModbusDataUnit).DestroyQModbusDataUnit)
		return tmpValue
	}
	return nil
}

func (ptr *QModbusReply) Error() QModbusDevice__Error {
	if ptr.Pointer() != nil {
		return QModbusDevice__Error(C.QModbusReply_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QModbusReply) RawResult() *QModbusResponse {
	if ptr.Pointer() != nil {
		var tmpValue = NewQModbusResponseFromPointer(C.QModbusReply_RawResult(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QModbusResponse).DestroyQModbusResponse)
		return tmpValue
	}
	return nil
}

func (ptr *QModbusReply) ErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QModbusReply_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QModbusReply) Type() QModbusReply__ReplyType {
	if ptr.Pointer() != nil {
		return QModbusReply__ReplyType(C.QModbusReply_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QModbusReply) IsFinished() bool {
	if ptr.Pointer() != nil {
		return C.QModbusReply_IsFinished(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QModbusReply) ServerAddress() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QModbusReply_ServerAddress(ptr.Pointer())))
	}
	return 0
}

func (ptr *QModbusReply) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QModbusReply___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QModbusReply) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusReply___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QModbusReply) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QModbusReply___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QModbusReply) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QModbusReply___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QModbusReply) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusReply___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QModbusReply) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QModbusReply___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QModbusReply) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QModbusReply___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QModbusReply) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusReply___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QModbusReply) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.QModbusReply___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QModbusReply) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QModbusReply___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QModbusReply) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusReply___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QModbusReply) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QModbusReply___findChildren_newList(ptr.Pointer()))
}

func (ptr *QModbusReply) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QModbusReply___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QModbusReply) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusReply___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QModbusReply) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QModbusReply___children_newList(ptr.Pointer()))
}

//export callbackQModbusReply_Event
func callbackQModbusReply_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusReplyFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QModbusReply) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusReply_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQModbusReply_EventFilter
func callbackQModbusReply_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusReplyFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QModbusReply) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QModbusReply_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQModbusReply_ChildEvent
func callbackQModbusReply_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQModbusReplyFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QModbusReply) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusReply_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQModbusReply_ConnectNotify
func callbackQModbusReply_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQModbusReplyFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QModbusReply) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusReply_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQModbusReply_CustomEvent
func callbackQModbusReply_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQModbusReplyFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QModbusReply) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusReply_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQModbusReply_DeleteLater
func callbackQModbusReply_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQModbusReplyFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QModbusReply) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QModbusReply_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQModbusReply_Destroyed
func callbackQModbusReply_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQModbusReply_DisconnectNotify
func callbackQModbusReply_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQModbusReplyFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QModbusReply) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusReply_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQModbusReply_ObjectNameChanged
func callbackQModbusReply_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtSerialBus_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQModbusReply_TimerEvent
func callbackQModbusReply_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQModbusReplyFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QModbusReply) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QModbusReply_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQModbusReply_MetaObject
func callbackQModbusReply_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQModbusReplyFromPointer(ptr).MetaObjectDefault())
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
		qt.DisconnectAllSignals(ptr.Pointer(), "")
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func NewQModbusRequest() *QModbusRequest {
	return NewQModbusRequestFromPointer(C.QModbusRequest_NewQModbusRequest())
}

func NewQModbusRequest3(code QModbusPdu__FunctionCode, data core.QByteArray_ITF) *QModbusRequest {
	return NewQModbusRequestFromPointer(C.QModbusRequest_NewQModbusRequest3(C.longlong(code), core.PointerFromQByteArray(data)))
}

func NewQModbusRequest2(pdu QModbusPdu_ITF) *QModbusRequest {
	return NewQModbusRequestFromPointer(C.QModbusRequest_NewQModbusRequest2(PointerFromQModbusPdu(pdu)))
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
		qt.DisconnectAllSignals(ptr.Pointer(), "")
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func NewQModbusResponse() *QModbusResponse {
	return NewQModbusResponseFromPointer(C.QModbusResponse_NewQModbusResponse())
}

func NewQModbusResponse3(code QModbusPdu__FunctionCode, data core.QByteArray_ITF) *QModbusResponse {
	return NewQModbusResponseFromPointer(C.QModbusResponse_NewQModbusResponse3(C.longlong(code), core.PointerFromQByteArray(data)))
}

func NewQModbusResponse2(pdu QModbusPdu_ITF) *QModbusResponse {
	return NewQModbusResponseFromPointer(C.QModbusResponse_NewQModbusResponse2(PointerFromQModbusPdu(pdu)))
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
func NewQModbusRtuSerialMaster(parent core.QObject_ITF) *QModbusRtuSerialMaster {
	var tmpValue = NewQModbusRtuSerialMasterFromPointer(C.QModbusRtuSerialMaster_NewQModbusRtuSerialMaster(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQModbusRtuSerialMaster_Open
func callbackQModbusRtuSerialMaster_Open(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "open"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusRtuSerialMasterFromPointer(ptr).OpenDefault())))
}

func (ptr *QModbusRtuSerialMaster) ConnectOpen(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "open"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "open", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "open", f)
		}
	}
}

func (ptr *QModbusRtuSerialMaster) DisconnectOpen() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "open")
	}
}

func (ptr *QModbusRtuSerialMaster) Open() bool {
	if ptr.Pointer() != nil {
		return C.QModbusRtuSerialMaster_Open(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QModbusRtuSerialMaster) OpenDefault() bool {
	if ptr.Pointer() != nil {
		return C.QModbusRtuSerialMaster_OpenDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQModbusRtuSerialMaster_Close
func callbackQModbusRtuSerialMaster_Close(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		signal.(func())()
	} else {
		NewQModbusRtuSerialMasterFromPointer(ptr).CloseDefault()
	}
}

func (ptr *QModbusRtuSerialMaster) ConnectClose(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "close"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "close", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "close", f)
		}
	}
}

func (ptr *QModbusRtuSerialMaster) DisconnectClose() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "close")
	}
}

func (ptr *QModbusRtuSerialMaster) Close() {
	if ptr.Pointer() != nil {
		C.QModbusRtuSerialMaster_Close(ptr.Pointer())
	}
}

func (ptr *QModbusRtuSerialMaster) CloseDefault() {
	if ptr.Pointer() != nil {
		C.QModbusRtuSerialMaster_CloseDefault(ptr.Pointer())
	}
}

func (ptr *QModbusRtuSerialMaster) SetInterFrameDelay(microseconds int) {
	if ptr.Pointer() != nil {
		C.QModbusRtuSerialMaster_SetInterFrameDelay(ptr.Pointer(), C.int(int32(microseconds)))
	}
}

func (ptr *QModbusRtuSerialMaster) InterFrameDelay() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QModbusRtuSerialMaster_InterFrameDelay(ptr.Pointer())))
	}
	return 0
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
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQModbusRtuSerialSlave_Open
func callbackQModbusRtuSerialSlave_Open(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "open"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusRtuSerialSlaveFromPointer(ptr).OpenDefault())))
}

func (ptr *QModbusRtuSerialSlave) ConnectOpen(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "open"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "open", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "open", f)
		}
	}
}

func (ptr *QModbusRtuSerialSlave) DisconnectOpen() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "open")
	}
}

func (ptr *QModbusRtuSerialSlave) Open() bool {
	if ptr.Pointer() != nil {
		return C.QModbusRtuSerialSlave_Open(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QModbusRtuSerialSlave) OpenDefault() bool {
	if ptr.Pointer() != nil {
		return C.QModbusRtuSerialSlave_OpenDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQModbusRtuSerialSlave_Close
func callbackQModbusRtuSerialSlave_Close(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		signal.(func())()
	} else {
		NewQModbusRtuSerialSlaveFromPointer(ptr).CloseDefault()
	}
}

func (ptr *QModbusRtuSerialSlave) ConnectClose(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "close"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "close", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "close", f)
		}
	}
}

func (ptr *QModbusRtuSerialSlave) DisconnectClose() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "close")
	}
}

func (ptr *QModbusRtuSerialSlave) Close() {
	if ptr.Pointer() != nil {
		C.QModbusRtuSerialSlave_Close(ptr.Pointer())
	}
}

func (ptr *QModbusRtuSerialSlave) CloseDefault() {
	if ptr.Pointer() != nil {
		C.QModbusRtuSerialSlave_CloseDefault(ptr.Pointer())
	}
}

func (ptr *QModbusRtuSerialSlave) DestroyQModbusRtuSerialSlave() {
	if ptr.Pointer() != nil {
		C.QModbusRtuSerialSlave_DestroyQModbusRtuSerialSlave(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
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

//export callbackQModbusServer_ProcessPrivateRequest
func callbackQModbusServer_ProcessPrivateRequest(ptr unsafe.Pointer, request unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "processPrivateRequest"); signal != nil {
		return PointerFromQModbusResponse(signal.(func(*QModbusPdu) *QModbusResponse)(NewQModbusPduFromPointer(request)))
	}

	return PointerFromQModbusResponse(NewQModbusServerFromPointer(ptr).ProcessPrivateRequestDefault(NewQModbusPduFromPointer(request)))
}

func (ptr *QModbusServer) ConnectProcessPrivateRequest(f func(request *QModbusPdu) *QModbusResponse) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "processPrivateRequest"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "processPrivateRequest", func(request *QModbusPdu) *QModbusResponse {
				signal.(func(*QModbusPdu) *QModbusResponse)(request)
				return f(request)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "processPrivateRequest", f)
		}
	}
}

func (ptr *QModbusServer) DisconnectProcessPrivateRequest() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "processPrivateRequest")
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
	if signal := qt.GetSignal(ptr, "processRequest"); signal != nil {
		return PointerFromQModbusResponse(signal.(func(*QModbusPdu) *QModbusResponse)(NewQModbusPduFromPointer(request)))
	}

	return PointerFromQModbusResponse(NewQModbusServerFromPointer(ptr).ProcessRequestDefault(NewQModbusPduFromPointer(request)))
}

func (ptr *QModbusServer) ConnectProcessRequest(f func(request *QModbusPdu) *QModbusResponse) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "processRequest"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "processRequest", func(request *QModbusPdu) *QModbusResponse {
				signal.(func(*QModbusPdu) *QModbusResponse)(request)
				return f(request)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "processRequest", f)
		}
	}
}

func (ptr *QModbusServer) DisconnectProcessRequest() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "processRequest")
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

func NewQModbusServer(parent core.QObject_ITF) *QModbusServer {
	var tmpValue = NewQModbusServerFromPointer(C.QModbusServer_NewQModbusServer(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
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

//export callbackQModbusServer_SetValue
func callbackQModbusServer_SetValue(ptr unsafe.Pointer, option C.int, newValue unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "setValue"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int, *core.QVariant) bool)(int(int32(option)), core.NewQVariantFromPointer(newValue)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusServerFromPointer(ptr).SetValueDefault(int(int32(option)), core.NewQVariantFromPointer(newValue)))))
}

func (ptr *QModbusServer) ConnectSetValue(f func(option int, newValue *core.QVariant) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "setValue"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "setValue", func(option int, newValue *core.QVariant) bool {
				signal.(func(int, *core.QVariant) bool)(option, newValue)
				return f(option, newValue)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "setValue", f)
		}
	}
}

func (ptr *QModbusServer) DisconnectSetValue() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "setValue")
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

//export callbackQModbusServer_WriteData
func callbackQModbusServer_WriteData(ptr unsafe.Pointer, newData unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "writeData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QModbusDataUnit) bool)(NewQModbusDataUnitFromPointer(newData)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusServerFromPointer(ptr).WriteDataDefault(NewQModbusDataUnitFromPointer(newData)))))
}

func (ptr *QModbusServer) ConnectWriteData(f func(newData *QModbusDataUnit) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "writeData"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "writeData", func(newData *QModbusDataUnit) bool {
				signal.(func(*QModbusDataUnit) bool)(newData)
				return f(newData)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "writeData", f)
		}
	}
}

func (ptr *QModbusServer) DisconnectWriteData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "writeData")
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

//export callbackQModbusServer_DataWritten
func callbackQModbusServer_DataWritten(ptr unsafe.Pointer, regist C.longlong, address C.int, size C.int) {
	if signal := qt.GetSignal(ptr, "dataWritten"); signal != nil {
		signal.(func(QModbusDataUnit__RegisterType, int, int))(QModbusDataUnit__RegisterType(regist), int(int32(address)), int(int32(size)))
	}

}

func (ptr *QModbusServer) ConnectDataWritten(f func(regist QModbusDataUnit__RegisterType, address int, size int)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "dataWritten") {
			C.QModbusServer_ConnectDataWritten(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "dataWritten"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "dataWritten", func(regist QModbusDataUnit__RegisterType, address int, size int) {
				signal.(func(QModbusDataUnit__RegisterType, int, int))(regist, address, size)
				f(regist, address, size)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "dataWritten", f)
		}
	}
}

func (ptr *QModbusServer) DisconnectDataWritten() {
	if ptr.Pointer() != nil {
		C.QModbusServer_DisconnectDataWritten(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "dataWritten")
	}
}

func (ptr *QModbusServer) DataWritten(regist QModbusDataUnit__RegisterType, address int, size int) {
	if ptr.Pointer() != nil {
		C.QModbusServer_DataWritten(ptr.Pointer(), C.longlong(regist), C.int(int32(address)), C.int(int32(size)))
	}
}

func (ptr *QModbusServer) SetServerAddress(serverAddress int) {
	if ptr.Pointer() != nil {
		C.QModbusServer_SetServerAddress(ptr.Pointer(), C.int(int32(serverAddress)))
	}
}

//export callbackQModbusServer_Value
func callbackQModbusServer_Value(ptr unsafe.Pointer, option C.int) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "value"); signal != nil {
		return core.PointerFromQVariant(signal.(func(int) *core.QVariant)(int(int32(option))))
	}

	return core.PointerFromQVariant(NewQModbusServerFromPointer(ptr).ValueDefault(int(int32(option))))
}

func (ptr *QModbusServer) ConnectValue(f func(option int) *core.QVariant) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "value"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "value", func(option int) *core.QVariant {
				signal.(func(int) *core.QVariant)(option)
				return f(option)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "value", f)
		}
	}
}

func (ptr *QModbusServer) DisconnectValue() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "value")
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

//export callbackQModbusServer_ProcessesBroadcast
func callbackQModbusServer_ProcessesBroadcast(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "processesBroadcast"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusServerFromPointer(ptr).ProcessesBroadcastDefault())))
}

func (ptr *QModbusServer) ConnectProcessesBroadcast(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "processesBroadcast"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "processesBroadcast", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "processesBroadcast", f)
		}
	}
}

func (ptr *QModbusServer) DisconnectProcessesBroadcast() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "processesBroadcast")
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
	if signal := qt.GetSignal(ptr, "readData"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*QModbusDataUnit) bool)(NewQModbusDataUnitFromPointer(newData)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusServerFromPointer(ptr).ReadDataDefault(NewQModbusDataUnitFromPointer(newData)))))
}

func (ptr *QModbusServer) ConnectReadData(f func(newData *QModbusDataUnit) bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "readData"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "readData", func(newData *QModbusDataUnit) bool {
				signal.(func(*QModbusDataUnit) bool)(newData)
				return f(newData)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "readData", f)
		}
	}
}

func (ptr *QModbusServer) DisconnectReadData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "readData")
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

//export callbackQModbusServer_Open
func callbackQModbusServer_Open(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "open"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusServerFromPointer(ptr).OpenDefault())))
}

func (ptr *QModbusServer) Open() bool {
	if ptr.Pointer() != nil {
		return C.QModbusServer_Open(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QModbusServer) OpenDefault() bool {
	if ptr.Pointer() != nil {
		return C.QModbusServer_OpenDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQModbusServer_Close
func callbackQModbusServer_Close(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		signal.(func())()
	} else {
		NewQModbusServerFromPointer(ptr).CloseDefault()
	}
}

func (ptr *QModbusServer) Close() {
	if ptr.Pointer() != nil {
		C.QModbusServer_Close(ptr.Pointer())
	}
}

func (ptr *QModbusServer) CloseDefault() {
	if ptr.Pointer() != nil {
		C.QModbusServer_CloseDefault(ptr.Pointer())
	}
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
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQModbusTcpClient_Open
func callbackQModbusTcpClient_Open(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "open"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusTcpClientFromPointer(ptr).OpenDefault())))
}

func (ptr *QModbusTcpClient) ConnectOpen(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "open"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "open", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "open", f)
		}
	}
}

func (ptr *QModbusTcpClient) DisconnectOpen() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "open")
	}
}

func (ptr *QModbusTcpClient) Open() bool {
	if ptr.Pointer() != nil {
		return C.QModbusTcpClient_Open(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QModbusTcpClient) OpenDefault() bool {
	if ptr.Pointer() != nil {
		return C.QModbusTcpClient_OpenDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQModbusTcpClient_Close
func callbackQModbusTcpClient_Close(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		signal.(func())()
	} else {
		NewQModbusTcpClientFromPointer(ptr).CloseDefault()
	}
}

func (ptr *QModbusTcpClient) ConnectClose(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "close"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "close", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "close", f)
		}
	}
}

func (ptr *QModbusTcpClient) DisconnectClose() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "close")
	}
}

func (ptr *QModbusTcpClient) Close() {
	if ptr.Pointer() != nil {
		C.QModbusTcpClient_Close(ptr.Pointer())
	}
}

func (ptr *QModbusTcpClient) CloseDefault() {
	if ptr.Pointer() != nil {
		C.QModbusTcpClient_CloseDefault(ptr.Pointer())
	}
}

func (ptr *QModbusTcpClient) DestroyQModbusTcpClient() {
	if ptr.Pointer() != nil {
		C.QModbusTcpClient_DestroyQModbusTcpClient(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
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
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQModbusTcpServer_Open
func callbackQModbusTcpServer_Open(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "open"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQModbusTcpServerFromPointer(ptr).OpenDefault())))
}

func (ptr *QModbusTcpServer) ConnectOpen(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "open"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "open", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "open", f)
		}
	}
}

func (ptr *QModbusTcpServer) DisconnectOpen() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "open")
	}
}

func (ptr *QModbusTcpServer) Open() bool {
	if ptr.Pointer() != nil {
		return C.QModbusTcpServer_Open(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QModbusTcpServer) OpenDefault() bool {
	if ptr.Pointer() != nil {
		return C.QModbusTcpServer_OpenDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQModbusTcpServer_Close
func callbackQModbusTcpServer_Close(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		signal.(func())()
	} else {
		NewQModbusTcpServerFromPointer(ptr).CloseDefault()
	}
}

func (ptr *QModbusTcpServer) ConnectClose(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "close"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "close", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "close", f)
		}
	}
}

func (ptr *QModbusTcpServer) DisconnectClose() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "close")
	}
}

func (ptr *QModbusTcpServer) Close() {
	if ptr.Pointer() != nil {
		C.QModbusTcpServer_Close(ptr.Pointer())
	}
}

func (ptr *QModbusTcpServer) CloseDefault() {
	if ptr.Pointer() != nil {
		C.QModbusTcpServer_CloseDefault(ptr.Pointer())
	}
}

func (ptr *QModbusTcpServer) DestroyQModbusTcpServer() {
	if ptr.Pointer() != nil {
		C.QModbusTcpServer_DestroyQModbusTcpServer(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}
