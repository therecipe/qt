// +build !minimal

package bluetooth

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "bluetooth.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/network"
	"runtime"
	"strings"
	"unsafe"
)

func cGoUnpackString(s C.struct_QtBluetooth_PackedString) string {
	if len := int(s.len); len == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}

type OSXBluetooth struct {
	ptr unsafe.Pointer
}

type OSXBluetooth_ITF interface {
	OSXBluetooth_PTR() *OSXBluetooth
}

func (ptr *OSXBluetooth) OSXBluetooth_PTR() *OSXBluetooth {
	return ptr
}

func (ptr *OSXBluetooth) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *OSXBluetooth) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromOSXBluetooth(ptr OSXBluetooth_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.OSXBluetooth_PTR().Pointer()
	}
	return nil
}

func NewOSXBluetoothFromPointer(ptr unsafe.Pointer) *OSXBluetooth {
	var n = new(OSXBluetooth)
	n.SetPointer(ptr)
	return n
}

func (ptr *OSXBluetooth) DestroyOSXBluetooth() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//go:generate stringer -type=OSXBluetooth__CentralManagerState
//OSXBluetooth::CentralManagerState
type OSXBluetooth__CentralManagerState int64

const (
	OSXBluetooth__CentralManagerIdle          OSXBluetooth__CentralManagerState = OSXBluetooth__CentralManagerState(0)
	OSXBluetooth__CentralManagerUpdating      OSXBluetooth__CentralManagerState = OSXBluetooth__CentralManagerState(1)
	OSXBluetooth__CentralManagerConnecting    OSXBluetooth__CentralManagerState = OSXBluetooth__CentralManagerState(2)
	OSXBluetooth__CentralManagerDiscovering   OSXBluetooth__CentralManagerState = OSXBluetooth__CentralManagerState(3)
	OSXBluetooth__CentralManagerDisconnecting OSXBluetooth__CentralManagerState = OSXBluetooth__CentralManagerState(4)
)

//go:generate stringer -type=OSXBluetooth__OBEXRequest
//OSXBluetooth::OBEXRequest
type OSXBluetooth__OBEXRequest int64

const (
	OSXBluetooth__OBEXNoop       OSXBluetooth__OBEXRequest = OSXBluetooth__OBEXRequest(0)
	OSXBluetooth__OBEXConnect    OSXBluetooth__OBEXRequest = OSXBluetooth__OBEXRequest(1)
	OSXBluetooth__OBEXDisconnect OSXBluetooth__OBEXRequest = OSXBluetooth__OBEXRequest(2)
	OSXBluetooth__OBEXPut        OSXBluetooth__OBEXRequest = OSXBluetooth__OBEXRequest(3)
	OSXBluetooth__OBEXGet        OSXBluetooth__OBEXRequest = OSXBluetooth__OBEXRequest(4)
	OSXBluetooth__OBEXSetPath    OSXBluetooth__OBEXRequest = OSXBluetooth__OBEXRequest(5)
	OSXBluetooth__OBEXAbort      OSXBluetooth__OBEXRequest = OSXBluetooth__OBEXRequest(6)
)

func (ptr *OSXBluetooth) __extract_services_uuids_atList(i int) *QBluetoothUuid {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothUuidFromPointer(C.OSXBluetooth___extract_services_uuids_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

func (ptr *OSXBluetooth) __extract_services_uuids_setList(i QBluetoothUuid_ITF) {
	if ptr.Pointer() != nil {
		C.OSXBluetooth___extract_services_uuids_setList(ptr.Pointer(), PointerFromQBluetoothUuid(i))
	}
}

func (ptr *OSXBluetooth) __extract_services_uuids_newList() unsafe.Pointer {
	return unsafe.Pointer(C.OSXBluetooth___extract_services_uuids_newList(ptr.Pointer()))
}

type QBluetooth struct {
	ptr unsafe.Pointer
}

type QBluetooth_ITF interface {
	QBluetooth_PTR() *QBluetooth
}

func (ptr *QBluetooth) QBluetooth_PTR() *QBluetooth {
	return ptr
}

func (ptr *QBluetooth) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QBluetooth) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
	}
}

func PointerFromQBluetooth(ptr QBluetooth_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetooth_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothFromPointer(ptr unsafe.Pointer) *QBluetooth {
	var n = new(QBluetooth)
	n.SetPointer(ptr)
	return n
}

func (ptr *QBluetooth) DestroyQBluetooth() {
	if ptr != nil {
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//go:generate stringer -type=QBluetooth__AttAccessConstraint
//QBluetooth::AttAccessConstraint
type QBluetooth__AttAccessConstraint int64

const (
	QBluetooth__AttAuthorizationRequired  QBluetooth__AttAccessConstraint = QBluetooth__AttAccessConstraint(0x1)
	QBluetooth__AttAuthenticationRequired QBluetooth__AttAccessConstraint = QBluetooth__AttAccessConstraint(0x2)
	QBluetooth__AttEncryptionRequired     QBluetooth__AttAccessConstraint = QBluetooth__AttAccessConstraint(0x4)
)

//go:generate stringer -type=QBluetooth__Security
//QBluetooth::Security
type QBluetooth__Security int64

const (
	QBluetooth__NoSecurity     QBluetooth__Security = QBluetooth__Security(0x00)
	QBluetooth__Authorization  QBluetooth__Security = QBluetooth__Security(0x01)
	QBluetooth__Authentication QBluetooth__Security = QBluetooth__Security(0x02)
	QBluetooth__Encryption     QBluetooth__Security = QBluetooth__Security(0x04)
	QBluetooth__Secure         QBluetooth__Security = QBluetooth__Security(0x08)
)

type QBluetoothAddress struct {
	ptr unsafe.Pointer
}

type QBluetoothAddress_ITF interface {
	QBluetoothAddress_PTR() *QBluetoothAddress
}

func (ptr *QBluetoothAddress) QBluetoothAddress_PTR() *QBluetoothAddress {
	return ptr
}

func (ptr *QBluetoothAddress) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QBluetoothAddress) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
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
func NewQBluetoothAddress() *QBluetoothAddress {
	var tmpValue = NewQBluetoothAddressFromPointer(C.QBluetoothAddress_NewQBluetoothAddress())
	runtime.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
	return tmpValue
}

func NewQBluetoothAddress4(other QBluetoothAddress_ITF) *QBluetoothAddress {
	var tmpValue = NewQBluetoothAddressFromPointer(C.QBluetoothAddress_NewQBluetoothAddress4(PointerFromQBluetoothAddress(other)))
	runtime.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
	return tmpValue
}

func NewQBluetoothAddress3(address string) *QBluetoothAddress {
	var addressC *C.char
	if address != "" {
		addressC = C.CString(address)
		defer C.free(unsafe.Pointer(addressC))
	}
	var tmpValue = NewQBluetoothAddressFromPointer(C.QBluetoothAddress_NewQBluetoothAddress3(C.struct_QtBluetooth_PackedString{data: addressC, len: C.longlong(len(address))}))
	runtime.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
	return tmpValue
}

func NewQBluetoothAddress2(address uint64) *QBluetoothAddress {
	var tmpValue = NewQBluetoothAddressFromPointer(C.QBluetoothAddress_NewQBluetoothAddress2(C.ulonglong(address)))
	runtime.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
	return tmpValue
}

func (ptr *QBluetoothAddress) Clear() {
	if ptr.Pointer() != nil {
		C.QBluetoothAddress_Clear(ptr.Pointer())
	}
}

func (ptr *QBluetoothAddress) DestroyQBluetoothAddress() {
	if ptr.Pointer() != nil {
		C.QBluetoothAddress_DestroyQBluetoothAddress(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QBluetoothAddress) ToString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QBluetoothAddress_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothAddress) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothAddress_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothAddress) ToUInt64() uint64 {
	if ptr.Pointer() != nil {
		return uint64(C.QBluetoothAddress_ToUInt64(ptr.Pointer()))
	}
	return 0
}

type QBluetoothDeviceDiscoveryAgent struct {
	core.QObject
}

type QBluetoothDeviceDiscoveryAgent_ITF interface {
	core.QObject_ITF
	QBluetoothDeviceDiscoveryAgent_PTR() *QBluetoothDeviceDiscoveryAgent
}

func (ptr *QBluetoothDeviceDiscoveryAgent) QBluetoothDeviceDiscoveryAgent_PTR() *QBluetoothDeviceDiscoveryAgent {
	return ptr
}

func (ptr *QBluetoothDeviceDiscoveryAgent) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QBluetoothDeviceDiscoveryAgent) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
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

//go:generate stringer -type=QBluetoothDeviceDiscoveryAgent__DiscoveryMethod
//QBluetoothDeviceDiscoveryAgent::DiscoveryMethod
type QBluetoothDeviceDiscoveryAgent__DiscoveryMethod int64

const (
	QBluetoothDeviceDiscoveryAgent__NoMethod        QBluetoothDeviceDiscoveryAgent__DiscoveryMethod = QBluetoothDeviceDiscoveryAgent__DiscoveryMethod(0x0)
	QBluetoothDeviceDiscoveryAgent__ClassicMethod   QBluetoothDeviceDiscoveryAgent__DiscoveryMethod = QBluetoothDeviceDiscoveryAgent__DiscoveryMethod(0x01)
	QBluetoothDeviceDiscoveryAgent__LowEnergyMethod QBluetoothDeviceDiscoveryAgent__DiscoveryMethod = QBluetoothDeviceDiscoveryAgent__DiscoveryMethod(0x02)
)

//go:generate stringer -type=QBluetoothDeviceDiscoveryAgent__Error
//QBluetoothDeviceDiscoveryAgent::Error
type QBluetoothDeviceDiscoveryAgent__Error int64

const (
	QBluetoothDeviceDiscoveryAgent__NoError                      QBluetoothDeviceDiscoveryAgent__Error = QBluetoothDeviceDiscoveryAgent__Error(0)
	QBluetoothDeviceDiscoveryAgent__InputOutputError             QBluetoothDeviceDiscoveryAgent__Error = QBluetoothDeviceDiscoveryAgent__Error(1)
	QBluetoothDeviceDiscoveryAgent__PoweredOffError              QBluetoothDeviceDiscoveryAgent__Error = QBluetoothDeviceDiscoveryAgent__Error(2)
	QBluetoothDeviceDiscoveryAgent__InvalidBluetoothAdapterError QBluetoothDeviceDiscoveryAgent__Error = QBluetoothDeviceDiscoveryAgent__Error(3)
	QBluetoothDeviceDiscoveryAgent__UnsupportedPlatformError     QBluetoothDeviceDiscoveryAgent__Error = QBluetoothDeviceDiscoveryAgent__Error(4)
	QBluetoothDeviceDiscoveryAgent__UnsupportedDiscoveryMethod   QBluetoothDeviceDiscoveryAgent__Error = QBluetoothDeviceDiscoveryAgent__Error(5)
	QBluetoothDeviceDiscoveryAgent__UnknownError                 QBluetoothDeviceDiscoveryAgent__Error = QBluetoothDeviceDiscoveryAgent__Error(100)
)

//go:generate stringer -type=QBluetoothDeviceDiscoveryAgent__InquiryType
//QBluetoothDeviceDiscoveryAgent::InquiryType
type QBluetoothDeviceDiscoveryAgent__InquiryType int64

const (
	QBluetoothDeviceDiscoveryAgent__GeneralUnlimitedInquiry QBluetoothDeviceDiscoveryAgent__InquiryType = QBluetoothDeviceDiscoveryAgent__InquiryType(0)
	QBluetoothDeviceDiscoveryAgent__LimitedInquiry          QBluetoothDeviceDiscoveryAgent__InquiryType = QBluetoothDeviceDiscoveryAgent__InquiryType(1)
)

func QBluetoothDeviceDiscoveryAgent_SupportedDiscoveryMethods() QBluetoothDeviceDiscoveryAgent__DiscoveryMethod {
	return QBluetoothDeviceDiscoveryAgent__DiscoveryMethod(C.QBluetoothDeviceDiscoveryAgent_QBluetoothDeviceDiscoveryAgent_SupportedDiscoveryMethods())
}

func (ptr *QBluetoothDeviceDiscoveryAgent) SupportedDiscoveryMethods() QBluetoothDeviceDiscoveryAgent__DiscoveryMethod {
	return QBluetoothDeviceDiscoveryAgent__DiscoveryMethod(C.QBluetoothDeviceDiscoveryAgent_QBluetoothDeviceDiscoveryAgent_SupportedDiscoveryMethods())
}

func NewQBluetoothDeviceDiscoveryAgent(parent core.QObject_ITF) *QBluetoothDeviceDiscoveryAgent {
	var tmpValue = NewQBluetoothDeviceDiscoveryAgentFromPointer(C.QBluetoothDeviceDiscoveryAgent_NewQBluetoothDeviceDiscoveryAgent(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQBluetoothDeviceDiscoveryAgent2(deviceAdapter QBluetoothAddress_ITF, parent core.QObject_ITF) *QBluetoothDeviceDiscoveryAgent {
	var tmpValue = NewQBluetoothDeviceDiscoveryAgentFromPointer(C.QBluetoothDeviceDiscoveryAgent_NewQBluetoothDeviceDiscoveryAgent2(PointerFromQBluetoothAddress(deviceAdapter), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQBluetoothDeviceDiscoveryAgent_Canceled
func callbackQBluetoothDeviceDiscoveryAgent_Canceled(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "canceled"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectCanceled(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "canceled") {
			C.QBluetoothDeviceDiscoveryAgent_ConnectCanceled(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "canceled"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "canceled", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "canceled", f)
		}
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectCanceled() {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_DisconnectCanceled(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "canceled")
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) Canceled() {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_Canceled(ptr.Pointer())
	}
}

//export callbackQBluetoothDeviceDiscoveryAgent_DeviceDiscovered
func callbackQBluetoothDeviceDiscoveryAgent_DeviceDiscovered(ptr unsafe.Pointer, info unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deviceDiscovered"); signal != nil {
		signal.(func(*QBluetoothDeviceInfo))(NewQBluetoothDeviceInfoFromPointer(info))
	}

}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectDeviceDiscovered(f func(info *QBluetoothDeviceInfo)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "deviceDiscovered") {
			C.QBluetoothDeviceDiscoveryAgent_ConnectDeviceDiscovered(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "deviceDiscovered"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "deviceDiscovered", func(info *QBluetoothDeviceInfo) {
				signal.(func(*QBluetoothDeviceInfo))(info)
				f(info)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "deviceDiscovered", f)
		}
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectDeviceDiscovered() {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_DisconnectDeviceDiscovered(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "deviceDiscovered")
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DeviceDiscovered(info QBluetoothDeviceInfo_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_DeviceDiscovered(ptr.Pointer(), PointerFromQBluetoothDeviceInfo(info))
	}
}

//export callbackQBluetoothDeviceDiscoveryAgent_Error2
func callbackQBluetoothDeviceDiscoveryAgent_Error2(ptr unsafe.Pointer, error C.longlong) {
	if signal := qt.GetSignal(ptr, "error2"); signal != nil {
		signal.(func(QBluetoothDeviceDiscoveryAgent__Error))(QBluetoothDeviceDiscoveryAgent__Error(error))
	}

}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectError2(f func(error QBluetoothDeviceDiscoveryAgent__Error)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "error2") {
			C.QBluetoothDeviceDiscoveryAgent_ConnectError2(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "error2"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "error2", func(error QBluetoothDeviceDiscoveryAgent__Error) {
				signal.(func(QBluetoothDeviceDiscoveryAgent__Error))(error)
				f(error)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "error2", f)
		}
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectError2() {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "error2")
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) Error2(error QBluetoothDeviceDiscoveryAgent__Error) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_Error2(ptr.Pointer(), C.longlong(error))
	}
}

//export callbackQBluetoothDeviceDiscoveryAgent_Finished
func callbackQBluetoothDeviceDiscoveryAgent_Finished(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "finished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectFinished(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "finished") {
			C.QBluetoothDeviceDiscoveryAgent_ConnectFinished(ptr.Pointer())
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

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "finished")
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) Finished() {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_Finished(ptr.Pointer())
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) SetInquiryType(ty QBluetoothDeviceDiscoveryAgent__InquiryType) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_SetInquiryType(ptr.Pointer(), C.longlong(ty))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) SetLowEnergyDiscoveryTimeout(timeout int) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_SetLowEnergyDiscoveryTimeout(ptr.Pointer(), C.int(int32(timeout)))
	}
}

//export callbackQBluetoothDeviceDiscoveryAgent_Start
func callbackQBluetoothDeviceDiscoveryAgent_Start(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "start"); signal != nil {
		signal.(func())()
	} else {
		NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr).StartDefault()
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectStart(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "start"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "start", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "start", f)
		}
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectStart() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "start")
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) Start() {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_Start(ptr.Pointer())
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) StartDefault() {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_StartDefault(ptr.Pointer())
	}
}

//export callbackQBluetoothDeviceDiscoveryAgent_Start2
func callbackQBluetoothDeviceDiscoveryAgent_Start2(ptr unsafe.Pointer, methods C.longlong) {
	if signal := qt.GetSignal(ptr, "start2"); signal != nil {
		signal.(func(QBluetoothDeviceDiscoveryAgent__DiscoveryMethod))(QBluetoothDeviceDiscoveryAgent__DiscoveryMethod(methods))
	} else {
		NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr).Start2Default(QBluetoothDeviceDiscoveryAgent__DiscoveryMethod(methods))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectStart2(f func(methods QBluetoothDeviceDiscoveryAgent__DiscoveryMethod)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "start2"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "start2", func(methods QBluetoothDeviceDiscoveryAgent__DiscoveryMethod) {
				signal.(func(QBluetoothDeviceDiscoveryAgent__DiscoveryMethod))(methods)
				f(methods)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "start2", f)
		}
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectStart2() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "start2")
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) Start2(methods QBluetoothDeviceDiscoveryAgent__DiscoveryMethod) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_Start2(ptr.Pointer(), C.longlong(methods))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) Start2Default(methods QBluetoothDeviceDiscoveryAgent__DiscoveryMethod) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_Start2Default(ptr.Pointer(), C.longlong(methods))
	}
}

//export callbackQBluetoothDeviceDiscoveryAgent_Stop
func callbackQBluetoothDeviceDiscoveryAgent_Stop(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "stop"); signal != nil {
		signal.(func())()
	} else {
		NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr).StopDefault()
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectStop(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "stop"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "stop", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stop", f)
		}
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectStop() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "stop")
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) Stop() {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_Stop(ptr.Pointer())
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) StopDefault() {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_StopDefault(ptr.Pointer())
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DestroyQBluetoothDeviceDiscoveryAgent() {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_DestroyQBluetoothDeviceDiscoveryAgent(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) Error() QBluetoothDeviceDiscoveryAgent__Error {
	if ptr.Pointer() != nil {
		return QBluetoothDeviceDiscoveryAgent__Error(C.QBluetoothDeviceDiscoveryAgent_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothDeviceDiscoveryAgent) InquiryType() QBluetoothDeviceDiscoveryAgent__InquiryType {
	if ptr.Pointer() != nil {
		return QBluetoothDeviceDiscoveryAgent__InquiryType(C.QBluetoothDeviceDiscoveryAgent_InquiryType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DiscoveredDevices() []*QBluetoothDeviceInfo {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtBluetooth_PackedList) []*QBluetoothDeviceInfo {
			var out = make([]*QBluetoothDeviceInfo, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQBluetoothDeviceDiscoveryAgentFromPointer(l.data).__discoveredDevices_atList(i)
			}
			return out
		}(C.QBluetoothDeviceDiscoveryAgent_DiscoveredDevices(ptr.Pointer()))
	}
	return make([]*QBluetoothDeviceInfo, 0)
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QBluetoothDeviceDiscoveryAgent_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothDeviceDiscoveryAgent) IsActive() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothDeviceDiscoveryAgent_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothDeviceDiscoveryAgent) LowEnergyDiscoveryTimeout() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QBluetoothDeviceDiscoveryAgent_LowEnergyDiscoveryTimeout(ptr.Pointer())))
	}
	return 0
}

func (ptr *QBluetoothDeviceDiscoveryAgent) __discoveredDevices_atList(i int) *QBluetoothDeviceInfo {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothDeviceInfoFromPointer(C.QBluetoothDeviceDiscoveryAgent___discoveredDevices_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QBluetoothDeviceInfo).DestroyQBluetoothDeviceInfo)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothDeviceDiscoveryAgent) __discoveredDevices_setList(i QBluetoothDeviceInfo_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent___discoveredDevices_setList(ptr.Pointer(), PointerFromQBluetoothDeviceInfo(i))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) __discoveredDevices_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QBluetoothDeviceDiscoveryAgent___discoveredDevices_newList(ptr.Pointer()))
}

func (ptr *QBluetoothDeviceDiscoveryAgent) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QBluetoothDeviceDiscoveryAgent___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothDeviceDiscoveryAgent) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QBluetoothDeviceDiscoveryAgent___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QBluetoothDeviceDiscoveryAgent) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QBluetoothDeviceDiscoveryAgent___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothDeviceDiscoveryAgent) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QBluetoothDeviceDiscoveryAgent___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QBluetoothDeviceDiscoveryAgent) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QBluetoothDeviceDiscoveryAgent___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothDeviceDiscoveryAgent) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.QBluetoothDeviceDiscoveryAgent___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QBluetoothDeviceDiscoveryAgent) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QBluetoothDeviceDiscoveryAgent___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothDeviceDiscoveryAgent) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QBluetoothDeviceDiscoveryAgent___findChildren_newList(ptr.Pointer()))
}

func (ptr *QBluetoothDeviceDiscoveryAgent) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QBluetoothDeviceDiscoveryAgent___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothDeviceDiscoveryAgent) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QBluetoothDeviceDiscoveryAgent___children_newList(ptr.Pointer()))
}

//export callbackQBluetoothDeviceDiscoveryAgent_Event
func callbackQBluetoothDeviceDiscoveryAgent_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QBluetoothDeviceDiscoveryAgent) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothDeviceDiscoveryAgent_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQBluetoothDeviceDiscoveryAgent_EventFilter
func callbackQBluetoothDeviceDiscoveryAgent_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QBluetoothDeviceDiscoveryAgent) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothDeviceDiscoveryAgent_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQBluetoothDeviceDiscoveryAgent_ChildEvent
func callbackQBluetoothDeviceDiscoveryAgent_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQBluetoothDeviceDiscoveryAgent_ConnectNotify
func callbackQBluetoothDeviceDiscoveryAgent_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBluetoothDeviceDiscoveryAgent_CustomEvent
func callbackQBluetoothDeviceDiscoveryAgent_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQBluetoothDeviceDiscoveryAgent_DeleteLater
func callbackQBluetoothDeviceDiscoveryAgent_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQBluetoothDeviceDiscoveryAgent_Destroyed
func callbackQBluetoothDeviceDiscoveryAgent_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQBluetoothDeviceDiscoveryAgent_DisconnectNotify
func callbackQBluetoothDeviceDiscoveryAgent_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBluetoothDeviceDiscoveryAgent_ObjectNameChanged
func callbackQBluetoothDeviceDiscoveryAgent_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtBluetooth_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQBluetoothDeviceDiscoveryAgent_TimerEvent
func callbackQBluetoothDeviceDiscoveryAgent_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQBluetoothDeviceDiscoveryAgent_MetaObject
func callbackQBluetoothDeviceDiscoveryAgent_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QBluetoothDeviceDiscoveryAgent) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QBluetoothDeviceDiscoveryAgent_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QBluetoothDeviceInfo struct {
	ptr unsafe.Pointer
}

type QBluetoothDeviceInfo_ITF interface {
	QBluetoothDeviceInfo_PTR() *QBluetoothDeviceInfo
}

func (ptr *QBluetoothDeviceInfo) QBluetoothDeviceInfo_PTR() *QBluetoothDeviceInfo {
	return ptr
}

func (ptr *QBluetoothDeviceInfo) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QBluetoothDeviceInfo) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
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

//go:generate stringer -type=QBluetoothDeviceInfo__CoreConfiguration
//QBluetoothDeviceInfo::CoreConfiguration
type QBluetoothDeviceInfo__CoreConfiguration int64

const (
	QBluetoothDeviceInfo__UnknownCoreConfiguration              QBluetoothDeviceInfo__CoreConfiguration = QBluetoothDeviceInfo__CoreConfiguration(0x0)
	QBluetoothDeviceInfo__LowEnergyCoreConfiguration            QBluetoothDeviceInfo__CoreConfiguration = QBluetoothDeviceInfo__CoreConfiguration(0x01)
	QBluetoothDeviceInfo__BaseRateCoreConfiguration             QBluetoothDeviceInfo__CoreConfiguration = QBluetoothDeviceInfo__CoreConfiguration(0x02)
	QBluetoothDeviceInfo__BaseRateAndLowEnergyCoreConfiguration QBluetoothDeviceInfo__CoreConfiguration = QBluetoothDeviceInfo__CoreConfiguration(0x03)
)

//go:generate stringer -type=QBluetoothDeviceInfo__DataCompleteness
//QBluetoothDeviceInfo::DataCompleteness
type QBluetoothDeviceInfo__DataCompleteness int64

const (
	QBluetoothDeviceInfo__DataComplete    QBluetoothDeviceInfo__DataCompleteness = QBluetoothDeviceInfo__DataCompleteness(0)
	QBluetoothDeviceInfo__DataIncomplete  QBluetoothDeviceInfo__DataCompleteness = QBluetoothDeviceInfo__DataCompleteness(1)
	QBluetoothDeviceInfo__DataUnavailable QBluetoothDeviceInfo__DataCompleteness = QBluetoothDeviceInfo__DataCompleteness(2)
)

//go:generate stringer -type=QBluetoothDeviceInfo__MajorDeviceClass
//QBluetoothDeviceInfo::MajorDeviceClass
type QBluetoothDeviceInfo__MajorDeviceClass int64

const (
	QBluetoothDeviceInfo__MiscellaneousDevice QBluetoothDeviceInfo__MajorDeviceClass = QBluetoothDeviceInfo__MajorDeviceClass(0)
	QBluetoothDeviceInfo__ComputerDevice      QBluetoothDeviceInfo__MajorDeviceClass = QBluetoothDeviceInfo__MajorDeviceClass(1)
	QBluetoothDeviceInfo__PhoneDevice         QBluetoothDeviceInfo__MajorDeviceClass = QBluetoothDeviceInfo__MajorDeviceClass(2)
	QBluetoothDeviceInfo__LANAccessDevice     QBluetoothDeviceInfo__MajorDeviceClass = QBluetoothDeviceInfo__MajorDeviceClass(3)
	QBluetoothDeviceInfo__AudioVideoDevice    QBluetoothDeviceInfo__MajorDeviceClass = QBluetoothDeviceInfo__MajorDeviceClass(4)
	QBluetoothDeviceInfo__PeripheralDevice    QBluetoothDeviceInfo__MajorDeviceClass = QBluetoothDeviceInfo__MajorDeviceClass(5)
	QBluetoothDeviceInfo__ImagingDevice       QBluetoothDeviceInfo__MajorDeviceClass = QBluetoothDeviceInfo__MajorDeviceClass(6)
	QBluetoothDeviceInfo__WearableDevice      QBluetoothDeviceInfo__MajorDeviceClass = QBluetoothDeviceInfo__MajorDeviceClass(7)
	QBluetoothDeviceInfo__ToyDevice           QBluetoothDeviceInfo__MajorDeviceClass = QBluetoothDeviceInfo__MajorDeviceClass(8)
	QBluetoothDeviceInfo__HealthDevice        QBluetoothDeviceInfo__MajorDeviceClass = QBluetoothDeviceInfo__MajorDeviceClass(9)
	QBluetoothDeviceInfo__UncategorizedDevice QBluetoothDeviceInfo__MajorDeviceClass = QBluetoothDeviceInfo__MajorDeviceClass(31)
)

//go:generate stringer -type=QBluetoothDeviceInfo__MinorAudioVideoClass
//QBluetoothDeviceInfo::MinorAudioVideoClass
type QBluetoothDeviceInfo__MinorAudioVideoClass int64

const (
	QBluetoothDeviceInfo__UncategorizedAudioVideoDevice QBluetoothDeviceInfo__MinorAudioVideoClass = QBluetoothDeviceInfo__MinorAudioVideoClass(0)
	QBluetoothDeviceInfo__WearableHeadsetDevice         QBluetoothDeviceInfo__MinorAudioVideoClass = QBluetoothDeviceInfo__MinorAudioVideoClass(1)
	QBluetoothDeviceInfo__HandsFreeDevice               QBluetoothDeviceInfo__MinorAudioVideoClass = QBluetoothDeviceInfo__MinorAudioVideoClass(2)
	QBluetoothDeviceInfo__Microphone                    QBluetoothDeviceInfo__MinorAudioVideoClass = QBluetoothDeviceInfo__MinorAudioVideoClass(4)
	QBluetoothDeviceInfo__Loudspeaker                   QBluetoothDeviceInfo__MinorAudioVideoClass = QBluetoothDeviceInfo__MinorAudioVideoClass(5)
	QBluetoothDeviceInfo__Headphones                    QBluetoothDeviceInfo__MinorAudioVideoClass = QBluetoothDeviceInfo__MinorAudioVideoClass(6)
	QBluetoothDeviceInfo__PortableAudioDevice           QBluetoothDeviceInfo__MinorAudioVideoClass = QBluetoothDeviceInfo__MinorAudioVideoClass(7)
	QBluetoothDeviceInfo__CarAudio                      QBluetoothDeviceInfo__MinorAudioVideoClass = QBluetoothDeviceInfo__MinorAudioVideoClass(8)
	QBluetoothDeviceInfo__SetTopBox                     QBluetoothDeviceInfo__MinorAudioVideoClass = QBluetoothDeviceInfo__MinorAudioVideoClass(9)
	QBluetoothDeviceInfo__HiFiAudioDevice               QBluetoothDeviceInfo__MinorAudioVideoClass = QBluetoothDeviceInfo__MinorAudioVideoClass(10)
	QBluetoothDeviceInfo__Vcr                           QBluetoothDeviceInfo__MinorAudioVideoClass = QBluetoothDeviceInfo__MinorAudioVideoClass(11)
	QBluetoothDeviceInfo__VideoCamera                   QBluetoothDeviceInfo__MinorAudioVideoClass = QBluetoothDeviceInfo__MinorAudioVideoClass(12)
	QBluetoothDeviceInfo__Camcorder                     QBluetoothDeviceInfo__MinorAudioVideoClass = QBluetoothDeviceInfo__MinorAudioVideoClass(13)
	QBluetoothDeviceInfo__VideoMonitor                  QBluetoothDeviceInfo__MinorAudioVideoClass = QBluetoothDeviceInfo__MinorAudioVideoClass(14)
	QBluetoothDeviceInfo__VideoDisplayAndLoudspeaker    QBluetoothDeviceInfo__MinorAudioVideoClass = QBluetoothDeviceInfo__MinorAudioVideoClass(15)
	QBluetoothDeviceInfo__VideoConferencing             QBluetoothDeviceInfo__MinorAudioVideoClass = QBluetoothDeviceInfo__MinorAudioVideoClass(16)
	QBluetoothDeviceInfo__GamingDevice                  QBluetoothDeviceInfo__MinorAudioVideoClass = QBluetoothDeviceInfo__MinorAudioVideoClass(18)
)

//go:generate stringer -type=QBluetoothDeviceInfo__MinorComputerClass
//QBluetoothDeviceInfo::MinorComputerClass
type QBluetoothDeviceInfo__MinorComputerClass int64

const (
	QBluetoothDeviceInfo__UncategorizedComputer     QBluetoothDeviceInfo__MinorComputerClass = QBluetoothDeviceInfo__MinorComputerClass(0)
	QBluetoothDeviceInfo__DesktopComputer           QBluetoothDeviceInfo__MinorComputerClass = QBluetoothDeviceInfo__MinorComputerClass(1)
	QBluetoothDeviceInfo__ServerComputer            QBluetoothDeviceInfo__MinorComputerClass = QBluetoothDeviceInfo__MinorComputerClass(2)
	QBluetoothDeviceInfo__LaptopComputer            QBluetoothDeviceInfo__MinorComputerClass = QBluetoothDeviceInfo__MinorComputerClass(3)
	QBluetoothDeviceInfo__HandheldClamShellComputer QBluetoothDeviceInfo__MinorComputerClass = QBluetoothDeviceInfo__MinorComputerClass(4)
	QBluetoothDeviceInfo__HandheldComputer          QBluetoothDeviceInfo__MinorComputerClass = QBluetoothDeviceInfo__MinorComputerClass(5)
	QBluetoothDeviceInfo__WearableComputer          QBluetoothDeviceInfo__MinorComputerClass = QBluetoothDeviceInfo__MinorComputerClass(6)
)

//go:generate stringer -type=QBluetoothDeviceInfo__MinorHealthClass
//QBluetoothDeviceInfo::MinorHealthClass
type QBluetoothDeviceInfo__MinorHealthClass int64

const (
	QBluetoothDeviceInfo__UncategorizedHealthDevice  QBluetoothDeviceInfo__MinorHealthClass = QBluetoothDeviceInfo__MinorHealthClass(0)
	QBluetoothDeviceInfo__HealthBloodPressureMonitor QBluetoothDeviceInfo__MinorHealthClass = QBluetoothDeviceInfo__MinorHealthClass(0x1)
	QBluetoothDeviceInfo__HealthThermometer          QBluetoothDeviceInfo__MinorHealthClass = QBluetoothDeviceInfo__MinorHealthClass(0x2)
	QBluetoothDeviceInfo__HealthWeightScale          QBluetoothDeviceInfo__MinorHealthClass = QBluetoothDeviceInfo__MinorHealthClass(0x3)
	QBluetoothDeviceInfo__HealthGlucoseMeter         QBluetoothDeviceInfo__MinorHealthClass = QBluetoothDeviceInfo__MinorHealthClass(0x4)
	QBluetoothDeviceInfo__HealthPulseOximeter        QBluetoothDeviceInfo__MinorHealthClass = QBluetoothDeviceInfo__MinorHealthClass(0x5)
	QBluetoothDeviceInfo__HealthDataDisplay          QBluetoothDeviceInfo__MinorHealthClass = QBluetoothDeviceInfo__MinorHealthClass(0x7)
	QBluetoothDeviceInfo__HealthStepCounter          QBluetoothDeviceInfo__MinorHealthClass = QBluetoothDeviceInfo__MinorHealthClass(0x8)
)

//go:generate stringer -type=QBluetoothDeviceInfo__MinorImagingClass
//QBluetoothDeviceInfo::MinorImagingClass
type QBluetoothDeviceInfo__MinorImagingClass int64

const (
	QBluetoothDeviceInfo__UncategorizedImagingDevice QBluetoothDeviceInfo__MinorImagingClass = QBluetoothDeviceInfo__MinorImagingClass(0)
	QBluetoothDeviceInfo__ImageDisplay               QBluetoothDeviceInfo__MinorImagingClass = QBluetoothDeviceInfo__MinorImagingClass(0x04)
	QBluetoothDeviceInfo__ImageCamera                QBluetoothDeviceInfo__MinorImagingClass = QBluetoothDeviceInfo__MinorImagingClass(0x08)
	QBluetoothDeviceInfo__ImageScanner               QBluetoothDeviceInfo__MinorImagingClass = QBluetoothDeviceInfo__MinorImagingClass(0x10)
	QBluetoothDeviceInfo__ImagePrinter               QBluetoothDeviceInfo__MinorImagingClass = QBluetoothDeviceInfo__MinorImagingClass(0x20)
)

//go:generate stringer -type=QBluetoothDeviceInfo__MinorMiscellaneousClass
//QBluetoothDeviceInfo::MinorMiscellaneousClass
type QBluetoothDeviceInfo__MinorMiscellaneousClass int64

const (
	QBluetoothDeviceInfo__UncategorizedMiscellaneous QBluetoothDeviceInfo__MinorMiscellaneousClass = QBluetoothDeviceInfo__MinorMiscellaneousClass(0)
)

//go:generate stringer -type=QBluetoothDeviceInfo__MinorNetworkClass
//QBluetoothDeviceInfo::MinorNetworkClass
type QBluetoothDeviceInfo__MinorNetworkClass int64

const (
	QBluetoothDeviceInfo__NetworkFullService     QBluetoothDeviceInfo__MinorNetworkClass = QBluetoothDeviceInfo__MinorNetworkClass(0x00)
	QBluetoothDeviceInfo__NetworkLoadFactorOne   QBluetoothDeviceInfo__MinorNetworkClass = QBluetoothDeviceInfo__MinorNetworkClass(0x08)
	QBluetoothDeviceInfo__NetworkLoadFactorTwo   QBluetoothDeviceInfo__MinorNetworkClass = QBluetoothDeviceInfo__MinorNetworkClass(0x10)
	QBluetoothDeviceInfo__NetworkLoadFactorThree QBluetoothDeviceInfo__MinorNetworkClass = QBluetoothDeviceInfo__MinorNetworkClass(0x18)
	QBluetoothDeviceInfo__NetworkLoadFactorFour  QBluetoothDeviceInfo__MinorNetworkClass = QBluetoothDeviceInfo__MinorNetworkClass(0x20)
	QBluetoothDeviceInfo__NetworkLoadFactorFive  QBluetoothDeviceInfo__MinorNetworkClass = QBluetoothDeviceInfo__MinorNetworkClass(0x28)
	QBluetoothDeviceInfo__NetworkLoadFactorSix   QBluetoothDeviceInfo__MinorNetworkClass = QBluetoothDeviceInfo__MinorNetworkClass(0x30)
	QBluetoothDeviceInfo__NetworkNoService       QBluetoothDeviceInfo__MinorNetworkClass = QBluetoothDeviceInfo__MinorNetworkClass(0x38)
)

//go:generate stringer -type=QBluetoothDeviceInfo__MinorPeripheralClass
//QBluetoothDeviceInfo::MinorPeripheralClass
type QBluetoothDeviceInfo__MinorPeripheralClass int64

const (
	QBluetoothDeviceInfo__UncategorizedPeripheral              QBluetoothDeviceInfo__MinorPeripheralClass = QBluetoothDeviceInfo__MinorPeripheralClass(0)
	QBluetoothDeviceInfo__KeyboardPeripheral                   QBluetoothDeviceInfo__MinorPeripheralClass = QBluetoothDeviceInfo__MinorPeripheralClass(0x10)
	QBluetoothDeviceInfo__PointingDevicePeripheral             QBluetoothDeviceInfo__MinorPeripheralClass = QBluetoothDeviceInfo__MinorPeripheralClass(0x20)
	QBluetoothDeviceInfo__KeyboardWithPointingDevicePeripheral QBluetoothDeviceInfo__MinorPeripheralClass = QBluetoothDeviceInfo__MinorPeripheralClass(0x30)
	QBluetoothDeviceInfo__JoystickPeripheral                   QBluetoothDeviceInfo__MinorPeripheralClass = QBluetoothDeviceInfo__MinorPeripheralClass(0x01)
	QBluetoothDeviceInfo__GamepadPeripheral                    QBluetoothDeviceInfo__MinorPeripheralClass = QBluetoothDeviceInfo__MinorPeripheralClass(0x02)
	QBluetoothDeviceInfo__RemoteControlPeripheral              QBluetoothDeviceInfo__MinorPeripheralClass = QBluetoothDeviceInfo__MinorPeripheralClass(0x03)
	QBluetoothDeviceInfo__SensingDevicePeripheral              QBluetoothDeviceInfo__MinorPeripheralClass = QBluetoothDeviceInfo__MinorPeripheralClass(0x04)
	QBluetoothDeviceInfo__DigitizerTabletPeripheral            QBluetoothDeviceInfo__MinorPeripheralClass = QBluetoothDeviceInfo__MinorPeripheralClass(0x05)
	QBluetoothDeviceInfo__CardReaderPeripheral                 QBluetoothDeviceInfo__MinorPeripheralClass = QBluetoothDeviceInfo__MinorPeripheralClass(0x06)
)

//go:generate stringer -type=QBluetoothDeviceInfo__MinorPhoneClass
//QBluetoothDeviceInfo::MinorPhoneClass
type QBluetoothDeviceInfo__MinorPhoneClass int64

const (
	QBluetoothDeviceInfo__UncategorizedPhone            QBluetoothDeviceInfo__MinorPhoneClass = QBluetoothDeviceInfo__MinorPhoneClass(0)
	QBluetoothDeviceInfo__CellularPhone                 QBluetoothDeviceInfo__MinorPhoneClass = QBluetoothDeviceInfo__MinorPhoneClass(1)
	QBluetoothDeviceInfo__CordlessPhone                 QBluetoothDeviceInfo__MinorPhoneClass = QBluetoothDeviceInfo__MinorPhoneClass(2)
	QBluetoothDeviceInfo__SmartPhone                    QBluetoothDeviceInfo__MinorPhoneClass = QBluetoothDeviceInfo__MinorPhoneClass(3)
	QBluetoothDeviceInfo__WiredModemOrVoiceGatewayPhone QBluetoothDeviceInfo__MinorPhoneClass = QBluetoothDeviceInfo__MinorPhoneClass(4)
	QBluetoothDeviceInfo__CommonIsdnAccessPhone         QBluetoothDeviceInfo__MinorPhoneClass = QBluetoothDeviceInfo__MinorPhoneClass(5)
)

//go:generate stringer -type=QBluetoothDeviceInfo__MinorToyClass
//QBluetoothDeviceInfo::MinorToyClass
type QBluetoothDeviceInfo__MinorToyClass int64

const (
	QBluetoothDeviceInfo__UncategorizedToy QBluetoothDeviceInfo__MinorToyClass = QBluetoothDeviceInfo__MinorToyClass(0)
	QBluetoothDeviceInfo__ToyRobot         QBluetoothDeviceInfo__MinorToyClass = QBluetoothDeviceInfo__MinorToyClass(1)
	QBluetoothDeviceInfo__ToyVehicle       QBluetoothDeviceInfo__MinorToyClass = QBluetoothDeviceInfo__MinorToyClass(2)
	QBluetoothDeviceInfo__ToyDoll          QBluetoothDeviceInfo__MinorToyClass = QBluetoothDeviceInfo__MinorToyClass(3)
	QBluetoothDeviceInfo__ToyController    QBluetoothDeviceInfo__MinorToyClass = QBluetoothDeviceInfo__MinorToyClass(4)
	QBluetoothDeviceInfo__ToyGame          QBluetoothDeviceInfo__MinorToyClass = QBluetoothDeviceInfo__MinorToyClass(5)
)

//go:generate stringer -type=QBluetoothDeviceInfo__MinorWearableClass
//QBluetoothDeviceInfo::MinorWearableClass
type QBluetoothDeviceInfo__MinorWearableClass int64

const (
	QBluetoothDeviceInfo__UncategorizedWearableDevice QBluetoothDeviceInfo__MinorWearableClass = QBluetoothDeviceInfo__MinorWearableClass(0)
	QBluetoothDeviceInfo__WearableWristWatch          QBluetoothDeviceInfo__MinorWearableClass = QBluetoothDeviceInfo__MinorWearableClass(1)
	QBluetoothDeviceInfo__WearablePager               QBluetoothDeviceInfo__MinorWearableClass = QBluetoothDeviceInfo__MinorWearableClass(2)
	QBluetoothDeviceInfo__WearableJacket              QBluetoothDeviceInfo__MinorWearableClass = QBluetoothDeviceInfo__MinorWearableClass(3)
	QBluetoothDeviceInfo__WearableHelmet              QBluetoothDeviceInfo__MinorWearableClass = QBluetoothDeviceInfo__MinorWearableClass(4)
	QBluetoothDeviceInfo__WearableGlasses             QBluetoothDeviceInfo__MinorWearableClass = QBluetoothDeviceInfo__MinorWearableClass(5)
)

//go:generate stringer -type=QBluetoothDeviceInfo__ServiceClass
//QBluetoothDeviceInfo::ServiceClass
type QBluetoothDeviceInfo__ServiceClass int64

const (
	QBluetoothDeviceInfo__NoService             QBluetoothDeviceInfo__ServiceClass = QBluetoothDeviceInfo__ServiceClass(0x0000)
	QBluetoothDeviceInfo__PositioningService    QBluetoothDeviceInfo__ServiceClass = QBluetoothDeviceInfo__ServiceClass(0x0001)
	QBluetoothDeviceInfo__NetworkingService     QBluetoothDeviceInfo__ServiceClass = QBluetoothDeviceInfo__ServiceClass(0x0002)
	QBluetoothDeviceInfo__RenderingService      QBluetoothDeviceInfo__ServiceClass = QBluetoothDeviceInfo__ServiceClass(0x0004)
	QBluetoothDeviceInfo__CapturingService      QBluetoothDeviceInfo__ServiceClass = QBluetoothDeviceInfo__ServiceClass(0x0008)
	QBluetoothDeviceInfo__ObjectTransferService QBluetoothDeviceInfo__ServiceClass = QBluetoothDeviceInfo__ServiceClass(0x0010)
	QBluetoothDeviceInfo__AudioService          QBluetoothDeviceInfo__ServiceClass = QBluetoothDeviceInfo__ServiceClass(0x0020)
	QBluetoothDeviceInfo__TelephonyService      QBluetoothDeviceInfo__ServiceClass = QBluetoothDeviceInfo__ServiceClass(0x0040)
	QBluetoothDeviceInfo__InformationService    QBluetoothDeviceInfo__ServiceClass = QBluetoothDeviceInfo__ServiceClass(0x0080)
	QBluetoothDeviceInfo__AllServices           QBluetoothDeviceInfo__ServiceClass = QBluetoothDeviceInfo__ServiceClass(0x07ff)
)

func NewQBluetoothDeviceInfo() *QBluetoothDeviceInfo {
	var tmpValue = NewQBluetoothDeviceInfoFromPointer(C.QBluetoothDeviceInfo_NewQBluetoothDeviceInfo())
	runtime.SetFinalizer(tmpValue, (*QBluetoothDeviceInfo).DestroyQBluetoothDeviceInfo)
	return tmpValue
}

func NewQBluetoothDeviceInfo2(address QBluetoothAddress_ITF, name string, classOfDevice uint) *QBluetoothDeviceInfo {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	var tmpValue = NewQBluetoothDeviceInfoFromPointer(C.QBluetoothDeviceInfo_NewQBluetoothDeviceInfo2(PointerFromQBluetoothAddress(address), C.struct_QtBluetooth_PackedString{data: nameC, len: C.longlong(len(name))}, C.uint(uint32(classOfDevice))))
	runtime.SetFinalizer(tmpValue, (*QBluetoothDeviceInfo).DestroyQBluetoothDeviceInfo)
	return tmpValue
}

func NewQBluetoothDeviceInfo4(other QBluetoothDeviceInfo_ITF) *QBluetoothDeviceInfo {
	var tmpValue = NewQBluetoothDeviceInfoFromPointer(C.QBluetoothDeviceInfo_NewQBluetoothDeviceInfo4(PointerFromQBluetoothDeviceInfo(other)))
	runtime.SetFinalizer(tmpValue, (*QBluetoothDeviceInfo).DestroyQBluetoothDeviceInfo)
	return tmpValue
}

func NewQBluetoothDeviceInfo3(uuid QBluetoothUuid_ITF, name string, classOfDevice uint) *QBluetoothDeviceInfo {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	var tmpValue = NewQBluetoothDeviceInfoFromPointer(C.QBluetoothDeviceInfo_NewQBluetoothDeviceInfo3(PointerFromQBluetoothUuid(uuid), C.struct_QtBluetooth_PackedString{data: nameC, len: C.longlong(len(name))}, C.uint(uint32(classOfDevice))))
	runtime.SetFinalizer(tmpValue, (*QBluetoothDeviceInfo).DestroyQBluetoothDeviceInfo)
	return tmpValue
}

func (ptr *QBluetoothDeviceInfo) SetCached(cached bool) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceInfo_SetCached(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(cached))))
	}
}

func (ptr *QBluetoothDeviceInfo) SetCoreConfigurations(coreConfigs QBluetoothDeviceInfo__CoreConfiguration) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceInfo_SetCoreConfigurations(ptr.Pointer(), C.longlong(coreConfigs))
	}
}

func (ptr *QBluetoothDeviceInfo) SetDeviceUuid(uuid QBluetoothUuid_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceInfo_SetDeviceUuid(ptr.Pointer(), PointerFromQBluetoothUuid(uuid))
	}
}

func (ptr *QBluetoothDeviceInfo) SetRssi(sign int16) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceInfo_SetRssi(ptr.Pointer(), C.short(sign))
	}
}

func (ptr *QBluetoothDeviceInfo) SetServiceUuids(uuids []*QBluetoothUuid, completeness QBluetoothDeviceInfo__DataCompleteness) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceInfo_SetServiceUuids(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewQBluetoothDeviceInfoFromPointer(NewQBluetoothDeviceInfoFromPointer(nil).__setServiceUuids_uuids_newList())
			for _, v := range uuids {
				tmpList.__setServiceUuids_uuids_setList(v)
			}
			return tmpList.Pointer()
		}(), C.longlong(completeness))
	}
}

func (ptr *QBluetoothDeviceInfo) DestroyQBluetoothDeviceInfo() {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceInfo_DestroyQBluetoothDeviceInfo(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QBluetoothDeviceInfo) ServiceUuidsCompleteness() QBluetoothDeviceInfo__DataCompleteness {
	if ptr.Pointer() != nil {
		return QBluetoothDeviceInfo__DataCompleteness(C.QBluetoothDeviceInfo_ServiceUuidsCompleteness(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothDeviceInfo) MajorDeviceClass() QBluetoothDeviceInfo__MajorDeviceClass {
	if ptr.Pointer() != nil {
		return QBluetoothDeviceInfo__MajorDeviceClass(C.QBluetoothDeviceInfo_MajorDeviceClass(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothDeviceInfo) Address() *QBluetoothAddress {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothAddressFromPointer(C.QBluetoothDeviceInfo_Address(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothDeviceInfo) CoreConfigurations() QBluetoothDeviceInfo__CoreConfiguration {
	if ptr.Pointer() != nil {
		return QBluetoothDeviceInfo__CoreConfiguration(C.QBluetoothDeviceInfo_CoreConfigurations(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothDeviceInfo) DeviceUuid() *QBluetoothUuid {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothUuidFromPointer(C.QBluetoothDeviceInfo_DeviceUuid(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothDeviceInfo) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QBluetoothDeviceInfo_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothDeviceInfo) ServiceClasses() QBluetoothDeviceInfo__ServiceClass {
	if ptr.Pointer() != nil {
		return QBluetoothDeviceInfo__ServiceClass(C.QBluetoothDeviceInfo_ServiceClasses(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothDeviceInfo) IsCached() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothDeviceInfo_IsCached(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothDeviceInfo) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothDeviceInfo_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothDeviceInfo) Rssi() int16 {
	if ptr.Pointer() != nil {
		return int16(C.QBluetoothDeviceInfo_Rssi(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothDeviceInfo) MinorDeviceClass() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QBluetoothDeviceInfo_MinorDeviceClass(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothDeviceInfo) __setServiceUuids_uuids_atList(i int) *QBluetoothUuid {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothUuidFromPointer(C.QBluetoothDeviceInfo___setServiceUuids_uuids_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothDeviceInfo) __setServiceUuids_uuids_setList(i QBluetoothUuid_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceInfo___setServiceUuids_uuids_setList(ptr.Pointer(), PointerFromQBluetoothUuid(i))
	}
}

func (ptr *QBluetoothDeviceInfo) __setServiceUuids_uuids_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QBluetoothDeviceInfo___setServiceUuids_uuids_newList(ptr.Pointer()))
}

func (ptr *QBluetoothDeviceInfo) __serviceUuids_atList(i int) *QBluetoothUuid {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothUuidFromPointer(C.QBluetoothDeviceInfo___serviceUuids_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothDeviceInfo) __serviceUuids_setList(i QBluetoothUuid_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceInfo___serviceUuids_setList(ptr.Pointer(), PointerFromQBluetoothUuid(i))
	}
}

func (ptr *QBluetoothDeviceInfo) __serviceUuids_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QBluetoothDeviceInfo___serviceUuids_newList(ptr.Pointer()))
}

type QBluetoothHostInfo struct {
	ptr unsafe.Pointer
}

type QBluetoothHostInfo_ITF interface {
	QBluetoothHostInfo_PTR() *QBluetoothHostInfo
}

func (ptr *QBluetoothHostInfo) QBluetoothHostInfo_PTR() *QBluetoothHostInfo {
	return ptr
}

func (ptr *QBluetoothHostInfo) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QBluetoothHostInfo) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
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
func NewQBluetoothHostInfo() *QBluetoothHostInfo {
	var tmpValue = NewQBluetoothHostInfoFromPointer(C.QBluetoothHostInfo_NewQBluetoothHostInfo())
	runtime.SetFinalizer(tmpValue, (*QBluetoothHostInfo).DestroyQBluetoothHostInfo)
	return tmpValue
}

func NewQBluetoothHostInfo2(other QBluetoothHostInfo_ITF) *QBluetoothHostInfo {
	var tmpValue = NewQBluetoothHostInfoFromPointer(C.QBluetoothHostInfo_NewQBluetoothHostInfo2(PointerFromQBluetoothHostInfo(other)))
	runtime.SetFinalizer(tmpValue, (*QBluetoothHostInfo).DestroyQBluetoothHostInfo)
	return tmpValue
}

func (ptr *QBluetoothHostInfo) SetAddress(address QBluetoothAddress_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothHostInfo_SetAddress(ptr.Pointer(), PointerFromQBluetoothAddress(address))
	}
}

func (ptr *QBluetoothHostInfo) SetName(name string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QBluetoothHostInfo_SetName(ptr.Pointer(), C.struct_QtBluetooth_PackedString{data: nameC, len: C.longlong(len(name))})
	}
}

func (ptr *QBluetoothHostInfo) DestroyQBluetoothHostInfo() {
	if ptr.Pointer() != nil {
		C.QBluetoothHostInfo_DestroyQBluetoothHostInfo(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QBluetoothHostInfo) Address() *QBluetoothAddress {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothAddressFromPointer(C.QBluetoothHostInfo_Address(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothHostInfo) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QBluetoothHostInfo_Name(ptr.Pointer()))
	}
	return ""
}

type QBluetoothLocalDevice struct {
	core.QObject
}

type QBluetoothLocalDevice_ITF interface {
	core.QObject_ITF
	QBluetoothLocalDevice_PTR() *QBluetoothLocalDevice
}

func (ptr *QBluetoothLocalDevice) QBluetoothLocalDevice_PTR() *QBluetoothLocalDevice {
	return ptr
}

func (ptr *QBluetoothLocalDevice) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QBluetoothLocalDevice) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
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

//go:generate stringer -type=QBluetoothLocalDevice__Error
//QBluetoothLocalDevice::Error
type QBluetoothLocalDevice__Error int64

const (
	QBluetoothLocalDevice__NoError      QBluetoothLocalDevice__Error = QBluetoothLocalDevice__Error(0)
	QBluetoothLocalDevice__PairingError QBluetoothLocalDevice__Error = QBluetoothLocalDevice__Error(1)
	QBluetoothLocalDevice__UnknownError QBluetoothLocalDevice__Error = QBluetoothLocalDevice__Error(100)
)

//go:generate stringer -type=QBluetoothLocalDevice__HostMode
//QBluetoothLocalDevice::HostMode
type QBluetoothLocalDevice__HostMode int64

const (
	QBluetoothLocalDevice__HostPoweredOff                 QBluetoothLocalDevice__HostMode = QBluetoothLocalDevice__HostMode(0)
	QBluetoothLocalDevice__HostConnectable                QBluetoothLocalDevice__HostMode = QBluetoothLocalDevice__HostMode(1)
	QBluetoothLocalDevice__HostDiscoverable               QBluetoothLocalDevice__HostMode = QBluetoothLocalDevice__HostMode(2)
	QBluetoothLocalDevice__HostDiscoverableLimitedInquiry QBluetoothLocalDevice__HostMode = QBluetoothLocalDevice__HostMode(3)
)

//go:generate stringer -type=QBluetoothLocalDevice__Pairing
//QBluetoothLocalDevice::Pairing
type QBluetoothLocalDevice__Pairing int64

const (
	QBluetoothLocalDevice__Unpaired         QBluetoothLocalDevice__Pairing = QBluetoothLocalDevice__Pairing(0)
	QBluetoothLocalDevice__Paired           QBluetoothLocalDevice__Pairing = QBluetoothLocalDevice__Pairing(1)
	QBluetoothLocalDevice__AuthorizedPaired QBluetoothLocalDevice__Pairing = QBluetoothLocalDevice__Pairing(2)
)

func NewQBluetoothLocalDevice(parent core.QObject_ITF) *QBluetoothLocalDevice {
	var tmpValue = NewQBluetoothLocalDeviceFromPointer(C.QBluetoothLocalDevice_NewQBluetoothLocalDevice(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQBluetoothLocalDevice2(address QBluetoothAddress_ITF, parent core.QObject_ITF) *QBluetoothLocalDevice {
	var tmpValue = NewQBluetoothLocalDeviceFromPointer(C.QBluetoothLocalDevice_NewQBluetoothLocalDevice2(PointerFromQBluetoothAddress(address), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QBluetoothLocalDevice_AllDevices() []*QBluetoothHostInfo {
	return func(l C.struct_QtBluetooth_PackedList) []*QBluetoothHostInfo {
		var out = make([]*QBluetoothHostInfo, int(l.len))
		for i := 0; i < int(l.len); i++ {
			out[i] = NewQBluetoothLocalDeviceFromPointer(l.data).__allDevices_atList(i)
		}
		return out
	}(C.QBluetoothLocalDevice_QBluetoothLocalDevice_AllDevices())
}

func (ptr *QBluetoothLocalDevice) AllDevices() []*QBluetoothHostInfo {
	return func(l C.struct_QtBluetooth_PackedList) []*QBluetoothHostInfo {
		var out = make([]*QBluetoothHostInfo, int(l.len))
		for i := 0; i < int(l.len); i++ {
			out[i] = NewQBluetoothLocalDeviceFromPointer(l.data).__allDevices_atList(i)
		}
		return out
	}(C.QBluetoothLocalDevice_QBluetoothLocalDevice_AllDevices())
}

//export callbackQBluetoothLocalDevice_DeviceConnected
func callbackQBluetoothLocalDevice_DeviceConnected(ptr unsafe.Pointer, address unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deviceConnected"); signal != nil {
		signal.(func(*QBluetoothAddress))(NewQBluetoothAddressFromPointer(address))
	}

}

func (ptr *QBluetoothLocalDevice) ConnectDeviceConnected(f func(address *QBluetoothAddress)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "deviceConnected") {
			C.QBluetoothLocalDevice_ConnectDeviceConnected(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "deviceConnected"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "deviceConnected", func(address *QBluetoothAddress) {
				signal.(func(*QBluetoothAddress))(address)
				f(address)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "deviceConnected", f)
		}
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectDeviceConnected() {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DisconnectDeviceConnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "deviceConnected")
	}
}

func (ptr *QBluetoothLocalDevice) DeviceConnected(address QBluetoothAddress_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DeviceConnected(ptr.Pointer(), PointerFromQBluetoothAddress(address))
	}
}

//export callbackQBluetoothLocalDevice_DeviceDisconnected
func callbackQBluetoothLocalDevice_DeviceDisconnected(ptr unsafe.Pointer, address unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deviceDisconnected"); signal != nil {
		signal.(func(*QBluetoothAddress))(NewQBluetoothAddressFromPointer(address))
	}

}

func (ptr *QBluetoothLocalDevice) ConnectDeviceDisconnected(f func(address *QBluetoothAddress)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "deviceDisconnected") {
			C.QBluetoothLocalDevice_ConnectDeviceDisconnected(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "deviceDisconnected"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "deviceDisconnected", func(address *QBluetoothAddress) {
				signal.(func(*QBluetoothAddress))(address)
				f(address)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "deviceDisconnected", f)
		}
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectDeviceDisconnected() {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DisconnectDeviceDisconnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "deviceDisconnected")
	}
}

func (ptr *QBluetoothLocalDevice) DeviceDisconnected(address QBluetoothAddress_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DeviceDisconnected(ptr.Pointer(), PointerFromQBluetoothAddress(address))
	}
}

//export callbackQBluetoothLocalDevice_Error
func callbackQBluetoothLocalDevice_Error(ptr unsafe.Pointer, error C.longlong) {
	if signal := qt.GetSignal(ptr, "error"); signal != nil {
		signal.(func(QBluetoothLocalDevice__Error))(QBluetoothLocalDevice__Error(error))
	}

}

func (ptr *QBluetoothLocalDevice) ConnectError(f func(error QBluetoothLocalDevice__Error)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "error") {
			C.QBluetoothLocalDevice_ConnectError(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "error"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "error", func(error QBluetoothLocalDevice__Error) {
				signal.(func(QBluetoothLocalDevice__Error))(error)
				f(error)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "error", f)
		}
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectError() {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DisconnectError(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "error")
	}
}

func (ptr *QBluetoothLocalDevice) Error(error QBluetoothLocalDevice__Error) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_Error(ptr.Pointer(), C.longlong(error))
	}
}

//export callbackQBluetoothLocalDevice_HostModeStateChanged
func callbackQBluetoothLocalDevice_HostModeStateChanged(ptr unsafe.Pointer, state C.longlong) {
	if signal := qt.GetSignal(ptr, "hostModeStateChanged"); signal != nil {
		signal.(func(QBluetoothLocalDevice__HostMode))(QBluetoothLocalDevice__HostMode(state))
	}

}

func (ptr *QBluetoothLocalDevice) ConnectHostModeStateChanged(f func(state QBluetoothLocalDevice__HostMode)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "hostModeStateChanged") {
			C.QBluetoothLocalDevice_ConnectHostModeStateChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "hostModeStateChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "hostModeStateChanged", func(state QBluetoothLocalDevice__HostMode) {
				signal.(func(QBluetoothLocalDevice__HostMode))(state)
				f(state)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hostModeStateChanged", f)
		}
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectHostModeStateChanged() {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DisconnectHostModeStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "hostModeStateChanged")
	}
}

func (ptr *QBluetoothLocalDevice) HostModeStateChanged(state QBluetoothLocalDevice__HostMode) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_HostModeStateChanged(ptr.Pointer(), C.longlong(state))
	}
}

//export callbackQBluetoothLocalDevice_PairingConfirmation
func callbackQBluetoothLocalDevice_PairingConfirmation(ptr unsafe.Pointer, accept C.char) {
	if signal := qt.GetSignal(ptr, "pairingConfirmation"); signal != nil {
		signal.(func(bool))(int8(accept) != 0)
	} else {
		NewQBluetoothLocalDeviceFromPointer(ptr).PairingConfirmationDefault(int8(accept) != 0)
	}
}

func (ptr *QBluetoothLocalDevice) ConnectPairingConfirmation(f func(accept bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "pairingConfirmation"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "pairingConfirmation", func(accept bool) {
				signal.(func(bool))(accept)
				f(accept)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "pairingConfirmation", f)
		}
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectPairingConfirmation() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "pairingConfirmation")
	}
}

func (ptr *QBluetoothLocalDevice) PairingConfirmation(accept bool) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_PairingConfirmation(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(accept))))
	}
}

func (ptr *QBluetoothLocalDevice) PairingConfirmationDefault(accept bool) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_PairingConfirmationDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(accept))))
	}
}

//export callbackQBluetoothLocalDevice_PairingDisplayConfirmation
func callbackQBluetoothLocalDevice_PairingDisplayConfirmation(ptr unsafe.Pointer, address unsafe.Pointer, pin C.struct_QtBluetooth_PackedString) {
	if signal := qt.GetSignal(ptr, "pairingDisplayConfirmation"); signal != nil {
		signal.(func(*QBluetoothAddress, string))(NewQBluetoothAddressFromPointer(address), cGoUnpackString(pin))
	}

}

func (ptr *QBluetoothLocalDevice) ConnectPairingDisplayConfirmation(f func(address *QBluetoothAddress, pin string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "pairingDisplayConfirmation") {
			C.QBluetoothLocalDevice_ConnectPairingDisplayConfirmation(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "pairingDisplayConfirmation"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "pairingDisplayConfirmation", func(address *QBluetoothAddress, pin string) {
				signal.(func(*QBluetoothAddress, string))(address, pin)
				f(address, pin)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "pairingDisplayConfirmation", f)
		}
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectPairingDisplayConfirmation() {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DisconnectPairingDisplayConfirmation(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "pairingDisplayConfirmation")
	}
}

func (ptr *QBluetoothLocalDevice) PairingDisplayConfirmation(address QBluetoothAddress_ITF, pin string) {
	if ptr.Pointer() != nil {
		var pinC *C.char
		if pin != "" {
			pinC = C.CString(pin)
			defer C.free(unsafe.Pointer(pinC))
		}
		C.QBluetoothLocalDevice_PairingDisplayConfirmation(ptr.Pointer(), PointerFromQBluetoothAddress(address), C.struct_QtBluetooth_PackedString{data: pinC, len: C.longlong(len(pin))})
	}
}

//export callbackQBluetoothLocalDevice_PairingDisplayPinCode
func callbackQBluetoothLocalDevice_PairingDisplayPinCode(ptr unsafe.Pointer, address unsafe.Pointer, pin C.struct_QtBluetooth_PackedString) {
	if signal := qt.GetSignal(ptr, "pairingDisplayPinCode"); signal != nil {
		signal.(func(*QBluetoothAddress, string))(NewQBluetoothAddressFromPointer(address), cGoUnpackString(pin))
	}

}

func (ptr *QBluetoothLocalDevice) ConnectPairingDisplayPinCode(f func(address *QBluetoothAddress, pin string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "pairingDisplayPinCode") {
			C.QBluetoothLocalDevice_ConnectPairingDisplayPinCode(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "pairingDisplayPinCode"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "pairingDisplayPinCode", func(address *QBluetoothAddress, pin string) {
				signal.(func(*QBluetoothAddress, string))(address, pin)
				f(address, pin)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "pairingDisplayPinCode", f)
		}
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectPairingDisplayPinCode() {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DisconnectPairingDisplayPinCode(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "pairingDisplayPinCode")
	}
}

func (ptr *QBluetoothLocalDevice) PairingDisplayPinCode(address QBluetoothAddress_ITF, pin string) {
	if ptr.Pointer() != nil {
		var pinC *C.char
		if pin != "" {
			pinC = C.CString(pin)
			defer C.free(unsafe.Pointer(pinC))
		}
		C.QBluetoothLocalDevice_PairingDisplayPinCode(ptr.Pointer(), PointerFromQBluetoothAddress(address), C.struct_QtBluetooth_PackedString{data: pinC, len: C.longlong(len(pin))})
	}
}

//export callbackQBluetoothLocalDevice_PairingFinished
func callbackQBluetoothLocalDevice_PairingFinished(ptr unsafe.Pointer, address unsafe.Pointer, pairing C.longlong) {
	if signal := qt.GetSignal(ptr, "pairingFinished"); signal != nil {
		signal.(func(*QBluetoothAddress, QBluetoothLocalDevice__Pairing))(NewQBluetoothAddressFromPointer(address), QBluetoothLocalDevice__Pairing(pairing))
	}

}

func (ptr *QBluetoothLocalDevice) ConnectPairingFinished(f func(address *QBluetoothAddress, pairing QBluetoothLocalDevice__Pairing)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "pairingFinished") {
			C.QBluetoothLocalDevice_ConnectPairingFinished(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "pairingFinished"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "pairingFinished", func(address *QBluetoothAddress, pairing QBluetoothLocalDevice__Pairing) {
				signal.(func(*QBluetoothAddress, QBluetoothLocalDevice__Pairing))(address, pairing)
				f(address, pairing)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "pairingFinished", f)
		}
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectPairingFinished() {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DisconnectPairingFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "pairingFinished")
	}
}

func (ptr *QBluetoothLocalDevice) PairingFinished(address QBluetoothAddress_ITF, pairing QBluetoothLocalDevice__Pairing) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_PairingFinished(ptr.Pointer(), PointerFromQBluetoothAddress(address), C.longlong(pairing))
	}
}

func (ptr *QBluetoothLocalDevice) PowerOn() {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_PowerOn(ptr.Pointer())
	}
}

func (ptr *QBluetoothLocalDevice) RequestPairing(address QBluetoothAddress_ITF, pairing QBluetoothLocalDevice__Pairing) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_RequestPairing(ptr.Pointer(), PointerFromQBluetoothAddress(address), C.longlong(pairing))
	}
}

func (ptr *QBluetoothLocalDevice) SetHostMode(mode QBluetoothLocalDevice__HostMode) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_SetHostMode(ptr.Pointer(), C.longlong(mode))
	}
}

//export callbackQBluetoothLocalDevice_DestroyQBluetoothLocalDevice
func callbackQBluetoothLocalDevice_DestroyQBluetoothLocalDevice(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QBluetoothLocalDevice"); signal != nil {
		signal.(func())()
	} else {
		NewQBluetoothLocalDeviceFromPointer(ptr).DestroyQBluetoothLocalDeviceDefault()
	}
}

func (ptr *QBluetoothLocalDevice) ConnectDestroyQBluetoothLocalDevice(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QBluetoothLocalDevice"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QBluetoothLocalDevice", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QBluetoothLocalDevice", f)
		}
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectDestroyQBluetoothLocalDevice() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QBluetoothLocalDevice")
	}
}

func (ptr *QBluetoothLocalDevice) DestroyQBluetoothLocalDevice() {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DestroyQBluetoothLocalDevice(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QBluetoothLocalDevice) DestroyQBluetoothLocalDeviceDefault() {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DestroyQBluetoothLocalDeviceDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QBluetoothLocalDevice) HostMode() QBluetoothLocalDevice__HostMode {
	if ptr.Pointer() != nil {
		return QBluetoothLocalDevice__HostMode(C.QBluetoothLocalDevice_HostMode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothLocalDevice) PairingStatus(address QBluetoothAddress_ITF) QBluetoothLocalDevice__Pairing {
	if ptr.Pointer() != nil {
		return QBluetoothLocalDevice__Pairing(C.QBluetoothLocalDevice_PairingStatus(ptr.Pointer(), PointerFromQBluetoothAddress(address)))
	}
	return 0
}

func (ptr *QBluetoothLocalDevice) Address() *QBluetoothAddress {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothAddressFromPointer(C.QBluetoothLocalDevice_Address(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothLocalDevice) ConnectedDevices() []*QBluetoothAddress {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtBluetooth_PackedList) []*QBluetoothAddress {
			var out = make([]*QBluetoothAddress, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQBluetoothLocalDeviceFromPointer(l.data).__connectedDevices_atList(i)
			}
			return out
		}(C.QBluetoothLocalDevice_ConnectedDevices(ptr.Pointer()))
	}
	return make([]*QBluetoothAddress, 0)
}

func (ptr *QBluetoothLocalDevice) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QBluetoothLocalDevice_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothLocalDevice) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothLocalDevice_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothLocalDevice) __allDevices_atList(i int) *QBluetoothHostInfo {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothHostInfoFromPointer(C.QBluetoothLocalDevice___allDevices_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QBluetoothHostInfo).DestroyQBluetoothHostInfo)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothLocalDevice) __allDevices_setList(i QBluetoothHostInfo_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice___allDevices_setList(ptr.Pointer(), PointerFromQBluetoothHostInfo(i))
	}
}

func (ptr *QBluetoothLocalDevice) __allDevices_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QBluetoothLocalDevice___allDevices_newList(ptr.Pointer()))
}

func (ptr *QBluetoothLocalDevice) __connectedDevices_atList(i int) *QBluetoothAddress {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothAddressFromPointer(C.QBluetoothLocalDevice___connectedDevices_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothLocalDevice) __connectedDevices_setList(i QBluetoothAddress_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice___connectedDevices_setList(ptr.Pointer(), PointerFromQBluetoothAddress(i))
	}
}

func (ptr *QBluetoothLocalDevice) __connectedDevices_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QBluetoothLocalDevice___connectedDevices_newList(ptr.Pointer()))
}

func (ptr *QBluetoothLocalDevice) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QBluetoothLocalDevice___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothLocalDevice) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QBluetoothLocalDevice) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QBluetoothLocalDevice___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QBluetoothLocalDevice) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QBluetoothLocalDevice___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothLocalDevice) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QBluetoothLocalDevice) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QBluetoothLocalDevice___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QBluetoothLocalDevice) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QBluetoothLocalDevice___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothLocalDevice) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QBluetoothLocalDevice) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.QBluetoothLocalDevice___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QBluetoothLocalDevice) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QBluetoothLocalDevice___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothLocalDevice) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QBluetoothLocalDevice) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QBluetoothLocalDevice___findChildren_newList(ptr.Pointer()))
}

func (ptr *QBluetoothLocalDevice) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QBluetoothLocalDevice___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothLocalDevice) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QBluetoothLocalDevice) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QBluetoothLocalDevice___children_newList(ptr.Pointer()))
}

//export callbackQBluetoothLocalDevice_Event
func callbackQBluetoothLocalDevice_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothLocalDeviceFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QBluetoothLocalDevice) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothLocalDevice_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQBluetoothLocalDevice_EventFilter
func callbackQBluetoothLocalDevice_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothLocalDeviceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QBluetoothLocalDevice) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothLocalDevice_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQBluetoothLocalDevice_ChildEvent
func callbackQBluetoothLocalDevice_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQBluetoothLocalDeviceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QBluetoothLocalDevice) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQBluetoothLocalDevice_ConnectNotify
func callbackQBluetoothLocalDevice_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBluetoothLocalDeviceFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBluetoothLocalDevice) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBluetoothLocalDevice_CustomEvent
func callbackQBluetoothLocalDevice_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQBluetoothLocalDeviceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QBluetoothLocalDevice) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQBluetoothLocalDevice_DeleteLater
func callbackQBluetoothLocalDevice_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQBluetoothLocalDeviceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QBluetoothLocalDevice) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQBluetoothLocalDevice_Destroyed
func callbackQBluetoothLocalDevice_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQBluetoothLocalDevice_DisconnectNotify
func callbackQBluetoothLocalDevice_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBluetoothLocalDeviceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBluetoothLocalDevice_ObjectNameChanged
func callbackQBluetoothLocalDevice_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtBluetooth_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQBluetoothLocalDevice_TimerEvent
func callbackQBluetoothLocalDevice_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQBluetoothLocalDeviceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QBluetoothLocalDevice) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQBluetoothLocalDevice_MetaObject
func callbackQBluetoothLocalDevice_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQBluetoothLocalDeviceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QBluetoothLocalDevice) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QBluetoothLocalDevice_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QBluetoothServer struct {
	core.QObject
}

type QBluetoothServer_ITF interface {
	core.QObject_ITF
	QBluetoothServer_PTR() *QBluetoothServer
}

func (ptr *QBluetoothServer) QBluetoothServer_PTR() *QBluetoothServer {
	return ptr
}

func (ptr *QBluetoothServer) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QBluetoothServer) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
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

//go:generate stringer -type=QBluetoothServer__Error
//QBluetoothServer::Error
type QBluetoothServer__Error int64

const (
	QBluetoothServer__NoError                       QBluetoothServer__Error = QBluetoothServer__Error(0)
	QBluetoothServer__UnknownError                  QBluetoothServer__Error = QBluetoothServer__Error(1)
	QBluetoothServer__PoweredOffError               QBluetoothServer__Error = QBluetoothServer__Error(2)
	QBluetoothServer__InputOutputError              QBluetoothServer__Error = QBluetoothServer__Error(3)
	QBluetoothServer__ServiceAlreadyRegisteredError QBluetoothServer__Error = QBluetoothServer__Error(4)
	QBluetoothServer__UnsupportedProtocolError      QBluetoothServer__Error = QBluetoothServer__Error(5)
)

func NewQBluetoothServer(serverType QBluetoothServiceInfo__Protocol, parent core.QObject_ITF) *QBluetoothServer {
	var tmpValue = NewQBluetoothServerFromPointer(C.QBluetoothServer_NewQBluetoothServer(C.longlong(serverType), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQBluetoothServer_Error2
func callbackQBluetoothServer_Error2(ptr unsafe.Pointer, error C.longlong) {
	if signal := qt.GetSignal(ptr, "error2"); signal != nil {
		signal.(func(QBluetoothServer__Error))(QBluetoothServer__Error(error))
	}

}

func (ptr *QBluetoothServer) ConnectError2(f func(error QBluetoothServer__Error)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "error2") {
			C.QBluetoothServer_ConnectError2(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "error2"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "error2", func(error QBluetoothServer__Error) {
				signal.(func(QBluetoothServer__Error))(error)
				f(error)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "error2", f)
		}
	}
}

func (ptr *QBluetoothServer) DisconnectError2() {
	if ptr.Pointer() != nil {
		C.QBluetoothServer_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "error2")
	}
}

func (ptr *QBluetoothServer) Error2(error QBluetoothServer__Error) {
	if ptr.Pointer() != nil {
		C.QBluetoothServer_Error2(ptr.Pointer(), C.longlong(error))
	}
}

//export callbackQBluetoothServer_NewConnection
func callbackQBluetoothServer_NewConnection(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "newConnection"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBluetoothServer) ConnectNewConnection(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "newConnection") {
			C.QBluetoothServer_ConnectNewConnection(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "newConnection"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "newConnection", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "newConnection", f)
		}
	}
}

func (ptr *QBluetoothServer) DisconnectNewConnection() {
	if ptr.Pointer() != nil {
		C.QBluetoothServer_DisconnectNewConnection(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "newConnection")
	}
}

func (ptr *QBluetoothServer) NewConnection() {
	if ptr.Pointer() != nil {
		C.QBluetoothServer_NewConnection(ptr.Pointer())
	}
}

func (ptr *QBluetoothServer) Listen2(uuid QBluetoothUuid_ITF, serviceName string) *QBluetoothServiceInfo {
	if ptr.Pointer() != nil {
		var serviceNameC *C.char
		if serviceName != "" {
			serviceNameC = C.CString(serviceName)
			defer C.free(unsafe.Pointer(serviceNameC))
		}
		var tmpValue = NewQBluetoothServiceInfoFromPointer(C.QBluetoothServer_Listen2(ptr.Pointer(), PointerFromQBluetoothUuid(uuid), C.struct_QtBluetooth_PackedString{data: serviceNameC, len: C.longlong(len(serviceName))}))
		runtime.SetFinalizer(tmpValue, (*QBluetoothServiceInfo).DestroyQBluetoothServiceInfo)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothServer) NextPendingConnection() *QBluetoothSocket {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothSocketFromPointer(C.QBluetoothServer_NextPendingConnection(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothServer) Listen(address QBluetoothAddress_ITF, port uint16) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothServer_Listen(ptr.Pointer(), PointerFromQBluetoothAddress(address), C.ushort(port)) != 0
	}
	return false
}

func (ptr *QBluetoothServer) Close() {
	if ptr.Pointer() != nil {
		C.QBluetoothServer_Close(ptr.Pointer())
	}
}

func (ptr *QBluetoothServer) SetMaxPendingConnections(numConnections int) {
	if ptr.Pointer() != nil {
		C.QBluetoothServer_SetMaxPendingConnections(ptr.Pointer(), C.int(int32(numConnections)))
	}
}

func (ptr *QBluetoothServer) SetSecurityFlags(security QBluetooth__Security) {
	if ptr.Pointer() != nil {
		C.QBluetoothServer_SetSecurityFlags(ptr.Pointer(), C.longlong(security))
	}
}

func (ptr *QBluetoothServer) DestroyQBluetoothServer() {
	if ptr.Pointer() != nil {
		C.QBluetoothServer_DestroyQBluetoothServer(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QBluetoothServer) Error() QBluetoothServer__Error {
	if ptr.Pointer() != nil {
		return QBluetoothServer__Error(C.QBluetoothServer_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothServer) SecurityFlags() QBluetooth__Security {
	if ptr.Pointer() != nil {
		return QBluetooth__Security(C.QBluetoothServer_SecurityFlags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothServer) ServerAddress() *QBluetoothAddress {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothAddressFromPointer(C.QBluetoothServer_ServerAddress(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothServer) ServerType() QBluetoothServiceInfo__Protocol {
	if ptr.Pointer() != nil {
		return QBluetoothServiceInfo__Protocol(C.QBluetoothServer_ServerType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothServer) HasPendingConnections() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothServer_HasPendingConnections(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothServer) IsListening() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothServer_IsListening(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothServer) MaxPendingConnections() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QBluetoothServer_MaxPendingConnections(ptr.Pointer())))
	}
	return 0
}

func (ptr *QBluetoothServer) ServerPort() uint16 {
	if ptr.Pointer() != nil {
		return uint16(C.QBluetoothServer_ServerPort(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothServer) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QBluetoothServer___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothServer) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServer___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QBluetoothServer) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QBluetoothServer___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QBluetoothServer) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QBluetoothServer___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothServer) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServer___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QBluetoothServer) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QBluetoothServer___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QBluetoothServer) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QBluetoothServer___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothServer) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServer___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QBluetoothServer) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.QBluetoothServer___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QBluetoothServer) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QBluetoothServer___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothServer) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServer___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QBluetoothServer) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QBluetoothServer___findChildren_newList(ptr.Pointer()))
}

func (ptr *QBluetoothServer) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QBluetoothServer___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothServer) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServer___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QBluetoothServer) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QBluetoothServer___children_newList(ptr.Pointer()))
}

//export callbackQBluetoothServer_Event
func callbackQBluetoothServer_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothServerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QBluetoothServer) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothServer_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQBluetoothServer_EventFilter
func callbackQBluetoothServer_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothServerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QBluetoothServer) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothServer_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQBluetoothServer_ChildEvent
func callbackQBluetoothServer_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQBluetoothServerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QBluetoothServer) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServer_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQBluetoothServer_ConnectNotify
func callbackQBluetoothServer_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBluetoothServerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBluetoothServer) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServer_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBluetoothServer_CustomEvent
func callbackQBluetoothServer_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQBluetoothServerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QBluetoothServer) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServer_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQBluetoothServer_DeleteLater
func callbackQBluetoothServer_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQBluetoothServerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QBluetoothServer) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QBluetoothServer_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQBluetoothServer_Destroyed
func callbackQBluetoothServer_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQBluetoothServer_DisconnectNotify
func callbackQBluetoothServer_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBluetoothServerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBluetoothServer) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServer_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBluetoothServer_ObjectNameChanged
func callbackQBluetoothServer_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtBluetooth_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQBluetoothServer_TimerEvent
func callbackQBluetoothServer_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQBluetoothServerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QBluetoothServer) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServer_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQBluetoothServer_MetaObject
func callbackQBluetoothServer_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQBluetoothServerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QBluetoothServer) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QBluetoothServer_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QBluetoothServiceDiscoveryAgent struct {
	core.QObject
}

type QBluetoothServiceDiscoveryAgent_ITF interface {
	core.QObject_ITF
	QBluetoothServiceDiscoveryAgent_PTR() *QBluetoothServiceDiscoveryAgent
}

func (ptr *QBluetoothServiceDiscoveryAgent) QBluetoothServiceDiscoveryAgent_PTR() *QBluetoothServiceDiscoveryAgent {
	return ptr
}

func (ptr *QBluetoothServiceDiscoveryAgent) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QBluetoothServiceDiscoveryAgent) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
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

//go:generate stringer -type=QBluetoothServiceDiscoveryAgent__DiscoveryMode
//QBluetoothServiceDiscoveryAgent::DiscoveryMode
type QBluetoothServiceDiscoveryAgent__DiscoveryMode int64

const (
	QBluetoothServiceDiscoveryAgent__MinimalDiscovery QBluetoothServiceDiscoveryAgent__DiscoveryMode = QBluetoothServiceDiscoveryAgent__DiscoveryMode(0)
	QBluetoothServiceDiscoveryAgent__FullDiscovery    QBluetoothServiceDiscoveryAgent__DiscoveryMode = QBluetoothServiceDiscoveryAgent__DiscoveryMode(1)
)

//go:generate stringer -type=QBluetoothServiceDiscoveryAgent__Error
//QBluetoothServiceDiscoveryAgent::Error
type QBluetoothServiceDiscoveryAgent__Error int64

const (
	QBluetoothServiceDiscoveryAgent__NoError                      QBluetoothServiceDiscoveryAgent__Error = QBluetoothServiceDiscoveryAgent__Error(QBluetoothDeviceDiscoveryAgent__NoError)
	QBluetoothServiceDiscoveryAgent__InputOutputError             QBluetoothServiceDiscoveryAgent__Error = QBluetoothServiceDiscoveryAgent__Error(QBluetoothDeviceDiscoveryAgent__InputOutputError)
	QBluetoothServiceDiscoveryAgent__PoweredOffError              QBluetoothServiceDiscoveryAgent__Error = QBluetoothServiceDiscoveryAgent__Error(QBluetoothDeviceDiscoveryAgent__PoweredOffError)
	QBluetoothServiceDiscoveryAgent__InvalidBluetoothAdapterError QBluetoothServiceDiscoveryAgent__Error = QBluetoothServiceDiscoveryAgent__Error(QBluetoothDeviceDiscoveryAgent__InvalidBluetoothAdapterError)
	QBluetoothServiceDiscoveryAgent__UnknownError                 QBluetoothServiceDiscoveryAgent__Error = QBluetoothServiceDiscoveryAgent__Error(QBluetoothDeviceDiscoveryAgent__UnknownError)
)

func NewQBluetoothServiceDiscoveryAgent(parent core.QObject_ITF) *QBluetoothServiceDiscoveryAgent {
	var tmpValue = NewQBluetoothServiceDiscoveryAgentFromPointer(C.QBluetoothServiceDiscoveryAgent_NewQBluetoothServiceDiscoveryAgent(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQBluetoothServiceDiscoveryAgent2(deviceAdapter QBluetoothAddress_ITF, parent core.QObject_ITF) *QBluetoothServiceDiscoveryAgent {
	var tmpValue = NewQBluetoothServiceDiscoveryAgentFromPointer(C.QBluetoothServiceDiscoveryAgent_NewQBluetoothServiceDiscoveryAgent2(PointerFromQBluetoothAddress(deviceAdapter), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QBluetoothServiceDiscoveryAgent) SetRemoteAddress(address QBluetoothAddress_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothServiceDiscoveryAgent_SetRemoteAddress(ptr.Pointer(), PointerFromQBluetoothAddress(address)) != 0
	}
	return false
}

//export callbackQBluetoothServiceDiscoveryAgent_Canceled
func callbackQBluetoothServiceDiscoveryAgent_Canceled(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "canceled"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectCanceled(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "canceled") {
			C.QBluetoothServiceDiscoveryAgent_ConnectCanceled(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "canceled"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "canceled", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "canceled", f)
		}
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectCanceled() {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_DisconnectCanceled(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "canceled")
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) Canceled() {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_Canceled(ptr.Pointer())
	}
}

//export callbackQBluetoothServiceDiscoveryAgent_Clear
func callbackQBluetoothServiceDiscoveryAgent_Clear(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "clear"); signal != nil {
		signal.(func())()
	} else {
		NewQBluetoothServiceDiscoveryAgentFromPointer(ptr).ClearDefault()
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectClear(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "clear"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "clear", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clear", f)
		}
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectClear() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "clear")
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) Clear() {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_Clear(ptr.Pointer())
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) ClearDefault() {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_ClearDefault(ptr.Pointer())
	}
}

//export callbackQBluetoothServiceDiscoveryAgent_Error2
func callbackQBluetoothServiceDiscoveryAgent_Error2(ptr unsafe.Pointer, error C.longlong) {
	if signal := qt.GetSignal(ptr, "error2"); signal != nil {
		signal.(func(QBluetoothServiceDiscoveryAgent__Error))(QBluetoothServiceDiscoveryAgent__Error(error))
	}

}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectError2(f func(error QBluetoothServiceDiscoveryAgent__Error)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "error2") {
			C.QBluetoothServiceDiscoveryAgent_ConnectError2(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "error2"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "error2", func(error QBluetoothServiceDiscoveryAgent__Error) {
				signal.(func(QBluetoothServiceDiscoveryAgent__Error))(error)
				f(error)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "error2", f)
		}
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectError2() {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "error2")
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) Error2(error QBluetoothServiceDiscoveryAgent__Error) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_Error2(ptr.Pointer(), C.longlong(error))
	}
}

//export callbackQBluetoothServiceDiscoveryAgent_Finished
func callbackQBluetoothServiceDiscoveryAgent_Finished(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "finished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectFinished(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "finished") {
			C.QBluetoothServiceDiscoveryAgent_ConnectFinished(ptr.Pointer())
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

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "finished")
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) Finished() {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_Finished(ptr.Pointer())
	}
}

//export callbackQBluetoothServiceDiscoveryAgent_ServiceDiscovered
func callbackQBluetoothServiceDiscoveryAgent_ServiceDiscovered(ptr unsafe.Pointer, info unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "serviceDiscovered"); signal != nil {
		signal.(func(*QBluetoothServiceInfo))(NewQBluetoothServiceInfoFromPointer(info))
	}

}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectServiceDiscovered(f func(info *QBluetoothServiceInfo)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "serviceDiscovered") {
			C.QBluetoothServiceDiscoveryAgent_ConnectServiceDiscovered(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "serviceDiscovered"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "serviceDiscovered", func(info *QBluetoothServiceInfo) {
				signal.(func(*QBluetoothServiceInfo))(info)
				f(info)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "serviceDiscovered", f)
		}
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectServiceDiscovered() {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_DisconnectServiceDiscovered(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "serviceDiscovered")
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) ServiceDiscovered(info QBluetoothServiceInfo_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_ServiceDiscovered(ptr.Pointer(), PointerFromQBluetoothServiceInfo(info))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) SetUuidFilter2(uuid QBluetoothUuid_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_SetUuidFilter2(ptr.Pointer(), PointerFromQBluetoothUuid(uuid))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) SetUuidFilter(uuids []*QBluetoothUuid) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_SetUuidFilter(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewQBluetoothServiceDiscoveryAgentFromPointer(NewQBluetoothServiceDiscoveryAgentFromPointer(nil).__setUuidFilter_uuids_newList())
			for _, v := range uuids {
				tmpList.__setUuidFilter_uuids_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

//export callbackQBluetoothServiceDiscoveryAgent_Start
func callbackQBluetoothServiceDiscoveryAgent_Start(ptr unsafe.Pointer, mode C.longlong) {
	if signal := qt.GetSignal(ptr, "start"); signal != nil {
		signal.(func(QBluetoothServiceDiscoveryAgent__DiscoveryMode))(QBluetoothServiceDiscoveryAgent__DiscoveryMode(mode))
	} else {
		NewQBluetoothServiceDiscoveryAgentFromPointer(ptr).StartDefault(QBluetoothServiceDiscoveryAgent__DiscoveryMode(mode))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectStart(f func(mode QBluetoothServiceDiscoveryAgent__DiscoveryMode)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "start"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "start", func(mode QBluetoothServiceDiscoveryAgent__DiscoveryMode) {
				signal.(func(QBluetoothServiceDiscoveryAgent__DiscoveryMode))(mode)
				f(mode)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "start", f)
		}
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectStart() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "start")
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) Start(mode QBluetoothServiceDiscoveryAgent__DiscoveryMode) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_Start(ptr.Pointer(), C.longlong(mode))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) StartDefault(mode QBluetoothServiceDiscoveryAgent__DiscoveryMode) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_StartDefault(ptr.Pointer(), C.longlong(mode))
	}
}

//export callbackQBluetoothServiceDiscoveryAgent_Stop
func callbackQBluetoothServiceDiscoveryAgent_Stop(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "stop"); signal != nil {
		signal.(func())()
	} else {
		NewQBluetoothServiceDiscoveryAgentFromPointer(ptr).StopDefault()
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectStop(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "stop"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "stop", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stop", f)
		}
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectStop() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "stop")
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) Stop() {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_Stop(ptr.Pointer())
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) StopDefault() {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_StopDefault(ptr.Pointer())
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DestroyQBluetoothServiceDiscoveryAgent() {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_DestroyQBluetoothServiceDiscoveryAgent(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) Error() QBluetoothServiceDiscoveryAgent__Error {
	if ptr.Pointer() != nil {
		return QBluetoothServiceDiscoveryAgent__Error(C.QBluetoothServiceDiscoveryAgent_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothServiceDiscoveryAgent) RemoteAddress() *QBluetoothAddress {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothAddressFromPointer(C.QBluetoothServiceDiscoveryAgent_RemoteAddress(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothServiceDiscoveryAgent) DiscoveredServices() []*QBluetoothServiceInfo {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtBluetooth_PackedList) []*QBluetoothServiceInfo {
			var out = make([]*QBluetoothServiceInfo, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQBluetoothServiceDiscoveryAgentFromPointer(l.data).__discoveredServices_atList(i)
			}
			return out
		}(C.QBluetoothServiceDiscoveryAgent_DiscoveredServices(ptr.Pointer()))
	}
	return make([]*QBluetoothServiceInfo, 0)
}

func (ptr *QBluetoothServiceDiscoveryAgent) UuidFilter() []*QBluetoothUuid {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtBluetooth_PackedList) []*QBluetoothUuid {
			var out = make([]*QBluetoothUuid, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQBluetoothServiceDiscoveryAgentFromPointer(l.data).__uuidFilter_atList(i)
			}
			return out
		}(C.QBluetoothServiceDiscoveryAgent_UuidFilter(ptr.Pointer()))
	}
	return make([]*QBluetoothUuid, 0)
}

func (ptr *QBluetoothServiceDiscoveryAgent) ErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QBluetoothServiceDiscoveryAgent_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothServiceDiscoveryAgent) IsActive() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothServiceDiscoveryAgent_IsActive(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothServiceDiscoveryAgent) __setUuidFilter_uuids_atList(i int) *QBluetoothUuid {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothUuidFromPointer(C.QBluetoothServiceDiscoveryAgent___setUuidFilter_uuids_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothServiceDiscoveryAgent) __setUuidFilter_uuids_setList(i QBluetoothUuid_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent___setUuidFilter_uuids_setList(ptr.Pointer(), PointerFromQBluetoothUuid(i))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) __setUuidFilter_uuids_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QBluetoothServiceDiscoveryAgent___setUuidFilter_uuids_newList(ptr.Pointer()))
}

func (ptr *QBluetoothServiceDiscoveryAgent) __discoveredServices_atList(i int) *QBluetoothServiceInfo {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothServiceInfoFromPointer(C.QBluetoothServiceDiscoveryAgent___discoveredServices_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QBluetoothServiceInfo).DestroyQBluetoothServiceInfo)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothServiceDiscoveryAgent) __discoveredServices_setList(i QBluetoothServiceInfo_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent___discoveredServices_setList(ptr.Pointer(), PointerFromQBluetoothServiceInfo(i))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) __discoveredServices_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QBluetoothServiceDiscoveryAgent___discoveredServices_newList(ptr.Pointer()))
}

func (ptr *QBluetoothServiceDiscoveryAgent) __uuidFilter_atList(i int) *QBluetoothUuid {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothUuidFromPointer(C.QBluetoothServiceDiscoveryAgent___uuidFilter_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothServiceDiscoveryAgent) __uuidFilter_setList(i QBluetoothUuid_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent___uuidFilter_setList(ptr.Pointer(), PointerFromQBluetoothUuid(i))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) __uuidFilter_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QBluetoothServiceDiscoveryAgent___uuidFilter_newList(ptr.Pointer()))
}

func (ptr *QBluetoothServiceDiscoveryAgent) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QBluetoothServiceDiscoveryAgent___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothServiceDiscoveryAgent) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QBluetoothServiceDiscoveryAgent___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QBluetoothServiceDiscoveryAgent) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QBluetoothServiceDiscoveryAgent___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothServiceDiscoveryAgent) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QBluetoothServiceDiscoveryAgent___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QBluetoothServiceDiscoveryAgent) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QBluetoothServiceDiscoveryAgent___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothServiceDiscoveryAgent) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.QBluetoothServiceDiscoveryAgent___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QBluetoothServiceDiscoveryAgent) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QBluetoothServiceDiscoveryAgent___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothServiceDiscoveryAgent) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QBluetoothServiceDiscoveryAgent___findChildren_newList(ptr.Pointer()))
}

func (ptr *QBluetoothServiceDiscoveryAgent) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QBluetoothServiceDiscoveryAgent___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothServiceDiscoveryAgent) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QBluetoothServiceDiscoveryAgent___children_newList(ptr.Pointer()))
}

//export callbackQBluetoothServiceDiscoveryAgent_Event
func callbackQBluetoothServiceDiscoveryAgent_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothServiceDiscoveryAgentFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QBluetoothServiceDiscoveryAgent) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothServiceDiscoveryAgent_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQBluetoothServiceDiscoveryAgent_EventFilter
func callbackQBluetoothServiceDiscoveryAgent_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothServiceDiscoveryAgentFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QBluetoothServiceDiscoveryAgent) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothServiceDiscoveryAgent_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQBluetoothServiceDiscoveryAgent_ChildEvent
func callbackQBluetoothServiceDiscoveryAgent_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQBluetoothServiceDiscoveryAgentFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQBluetoothServiceDiscoveryAgent_ConnectNotify
func callbackQBluetoothServiceDiscoveryAgent_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBluetoothServiceDiscoveryAgentFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBluetoothServiceDiscoveryAgent_CustomEvent
func callbackQBluetoothServiceDiscoveryAgent_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQBluetoothServiceDiscoveryAgentFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQBluetoothServiceDiscoveryAgent_DeleteLater
func callbackQBluetoothServiceDiscoveryAgent_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQBluetoothServiceDiscoveryAgentFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQBluetoothServiceDiscoveryAgent_Destroyed
func callbackQBluetoothServiceDiscoveryAgent_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQBluetoothServiceDiscoveryAgent_DisconnectNotify
func callbackQBluetoothServiceDiscoveryAgent_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBluetoothServiceDiscoveryAgentFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBluetoothServiceDiscoveryAgent_ObjectNameChanged
func callbackQBluetoothServiceDiscoveryAgent_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtBluetooth_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQBluetoothServiceDiscoveryAgent_TimerEvent
func callbackQBluetoothServiceDiscoveryAgent_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQBluetoothServiceDiscoveryAgentFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQBluetoothServiceDiscoveryAgent_MetaObject
func callbackQBluetoothServiceDiscoveryAgent_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQBluetoothServiceDiscoveryAgentFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QBluetoothServiceDiscoveryAgent) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QBluetoothServiceDiscoveryAgent_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QBluetoothServiceInfo struct {
	ptr unsafe.Pointer
}

type QBluetoothServiceInfo_ITF interface {
	QBluetoothServiceInfo_PTR() *QBluetoothServiceInfo
}

func (ptr *QBluetoothServiceInfo) QBluetoothServiceInfo_PTR() *QBluetoothServiceInfo {
	return ptr
}

func (ptr *QBluetoothServiceInfo) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QBluetoothServiceInfo) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
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

//go:generate stringer -type=QBluetoothServiceInfo__AttributeId
//QBluetoothServiceInfo::AttributeId
type QBluetoothServiceInfo__AttributeId int64

var (
	QBluetoothServiceInfo__ServiceRecordHandle              QBluetoothServiceInfo__AttributeId = QBluetoothServiceInfo__AttributeId(0x0000)
	QBluetoothServiceInfo__ServiceClassIds                  QBluetoothServiceInfo__AttributeId = QBluetoothServiceInfo__AttributeId(0x0001)
	QBluetoothServiceInfo__ServiceRecordState               QBluetoothServiceInfo__AttributeId = QBluetoothServiceInfo__AttributeId(0x0002)
	QBluetoothServiceInfo__ServiceId                        QBluetoothServiceInfo__AttributeId = QBluetoothServiceInfo__AttributeId(0x0003)
	QBluetoothServiceInfo__ProtocolDescriptorList           QBluetoothServiceInfo__AttributeId = QBluetoothServiceInfo__AttributeId(0x0004)
	QBluetoothServiceInfo__BrowseGroupList                  QBluetoothServiceInfo__AttributeId = QBluetoothServiceInfo__AttributeId(0x0005)
	QBluetoothServiceInfo__LanguageBaseAttributeIdList      QBluetoothServiceInfo__AttributeId = QBluetoothServiceInfo__AttributeId(0x0006)
	QBluetoothServiceInfo__ServiceInfoTimeToLive            QBluetoothServiceInfo__AttributeId = QBluetoothServiceInfo__AttributeId(0x0007)
	QBluetoothServiceInfo__ServiceAvailability              QBluetoothServiceInfo__AttributeId = QBluetoothServiceInfo__AttributeId(0x0008)
	QBluetoothServiceInfo__BluetoothProfileDescriptorList   QBluetoothServiceInfo__AttributeId = QBluetoothServiceInfo__AttributeId(0x0009)
	QBluetoothServiceInfo__DocumentationUrl                 QBluetoothServiceInfo__AttributeId = QBluetoothServiceInfo__AttributeId(0x000A)
	QBluetoothServiceInfo__ClientExecutableUrl              QBluetoothServiceInfo__AttributeId = QBluetoothServiceInfo__AttributeId(0x000B)
	QBluetoothServiceInfo__IconUrl                          QBluetoothServiceInfo__AttributeId = QBluetoothServiceInfo__AttributeId(0x000C)
	QBluetoothServiceInfo__AdditionalProtocolDescriptorList QBluetoothServiceInfo__AttributeId = QBluetoothServiceInfo__AttributeId(0x000D)
	QBluetoothServiceInfo__PrimaryLanguageBase              QBluetoothServiceInfo__AttributeId = QBluetoothServiceInfo__AttributeId(0x0100)
	QBluetoothServiceInfo__ServiceName                      QBluetoothServiceInfo__AttributeId = QBluetoothServiceInfo__AttributeId(C.QBluetoothServiceInfo_ServiceName_Type())
	QBluetoothServiceInfo__ServiceDescription               QBluetoothServiceInfo__AttributeId = QBluetoothServiceInfo__AttributeId(C.QBluetoothServiceInfo_ServiceDescription_Type())
	QBluetoothServiceInfo__ServiceProvider                  QBluetoothServiceInfo__AttributeId = QBluetoothServiceInfo__AttributeId(C.QBluetoothServiceInfo_ServiceProvider_Type())
)

//go:generate stringer -type=QBluetoothServiceInfo__Protocol
//QBluetoothServiceInfo::Protocol
type QBluetoothServiceInfo__Protocol int64

const (
	QBluetoothServiceInfo__UnknownProtocol QBluetoothServiceInfo__Protocol = QBluetoothServiceInfo__Protocol(0)
	QBluetoothServiceInfo__L2capProtocol   QBluetoothServiceInfo__Protocol = QBluetoothServiceInfo__Protocol(1)
	QBluetoothServiceInfo__RfcommProtocol  QBluetoothServiceInfo__Protocol = QBluetoothServiceInfo__Protocol(2)
)

func NewQBluetoothServiceInfo() *QBluetoothServiceInfo {
	var tmpValue = NewQBluetoothServiceInfoFromPointer(C.QBluetoothServiceInfo_NewQBluetoothServiceInfo())
	runtime.SetFinalizer(tmpValue, (*QBluetoothServiceInfo).DestroyQBluetoothServiceInfo)
	return tmpValue
}

func NewQBluetoothServiceInfo2(other QBluetoothServiceInfo_ITF) *QBluetoothServiceInfo {
	var tmpValue = NewQBluetoothServiceInfoFromPointer(C.QBluetoothServiceInfo_NewQBluetoothServiceInfo2(PointerFromQBluetoothServiceInfo(other)))
	runtime.SetFinalizer(tmpValue, (*QBluetoothServiceInfo).DestroyQBluetoothServiceInfo)
	return tmpValue
}

func (ptr *QBluetoothServiceInfo) RegisterService(localAdapter QBluetoothAddress_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothServiceInfo_RegisterService(ptr.Pointer(), PointerFromQBluetoothAddress(localAdapter)) != 0
	}
	return false
}

func (ptr *QBluetoothServiceInfo) UnregisterService() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothServiceInfo_UnregisterService(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothServiceInfo) RemoveAttribute(attributeId uint16) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceInfo_RemoveAttribute(ptr.Pointer(), C.ushort(attributeId))
	}
}

func (ptr *QBluetoothServiceInfo) SetAttribute2(attributeId uint16, value QBluetoothUuid_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceInfo_SetAttribute2(ptr.Pointer(), C.ushort(attributeId), PointerFromQBluetoothUuid(value))
	}
}

func (ptr *QBluetoothServiceInfo) SetAttribute(attributeId uint16, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceInfo_SetAttribute(ptr.Pointer(), C.ushort(attributeId), core.PointerFromQVariant(value))
	}
}

func (ptr *QBluetoothServiceInfo) SetDevice(device QBluetoothDeviceInfo_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceInfo_SetDevice(ptr.Pointer(), PointerFromQBluetoothDeviceInfo(device))
	}
}

func (ptr *QBluetoothServiceInfo) SetServiceAvailability(availability string) {
	if ptr.Pointer() != nil {
		var availabilityC *C.char
		if availability != "" {
			availabilityC = C.CString(availability)
			defer C.free(unsafe.Pointer(availabilityC))
		}
		C.QBluetoothServiceInfo_SetServiceAvailability(ptr.Pointer(), availabilityC)
	}
}

func (ptr *QBluetoothServiceInfo) SetServiceDescription(description string) {
	if ptr.Pointer() != nil {
		var descriptionC *C.char
		if description != "" {
			descriptionC = C.CString(description)
			defer C.free(unsafe.Pointer(descriptionC))
		}
		C.QBluetoothServiceInfo_SetServiceDescription(ptr.Pointer(), C.struct_QtBluetooth_PackedString{data: descriptionC, len: C.longlong(len(description))})
	}
}

func (ptr *QBluetoothServiceInfo) SetServiceName(name string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QBluetoothServiceInfo_SetServiceName(ptr.Pointer(), C.struct_QtBluetooth_PackedString{data: nameC, len: C.longlong(len(name))})
	}
}

func (ptr *QBluetoothServiceInfo) SetServiceProvider(provider string) {
	if ptr.Pointer() != nil {
		var providerC *C.char
		if provider != "" {
			providerC = C.CString(provider)
			defer C.free(unsafe.Pointer(providerC))
		}
		C.QBluetoothServiceInfo_SetServiceProvider(ptr.Pointer(), C.struct_QtBluetooth_PackedString{data: providerC, len: C.longlong(len(provider))})
	}
}

func (ptr *QBluetoothServiceInfo) SetServiceUuid(uuid QBluetoothUuid_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceInfo_SetServiceUuid(ptr.Pointer(), PointerFromQBluetoothUuid(uuid))
	}
}

func (ptr *QBluetoothServiceInfo) DestroyQBluetoothServiceInfo() {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceInfo_DestroyQBluetoothServiceInfo(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QBluetoothServiceInfo) Device() *QBluetoothDeviceInfo {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothDeviceInfoFromPointer(C.QBluetoothServiceInfo_Device(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothDeviceInfo).DestroyQBluetoothDeviceInfo)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothServiceInfo) SocketProtocol() QBluetoothServiceInfo__Protocol {
	if ptr.Pointer() != nil {
		return QBluetoothServiceInfo__Protocol(C.QBluetoothServiceInfo_SocketProtocol(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothServiceInfo) ServiceUuid() *QBluetoothUuid {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothUuidFromPointer(C.QBluetoothServiceInfo_ServiceUuid(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothServiceInfo) ServiceClassUuids() []*QBluetoothUuid {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtBluetooth_PackedList) []*QBluetoothUuid {
			var out = make([]*QBluetoothUuid, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQBluetoothServiceInfoFromPointer(l.data).__serviceClassUuids_atList(i)
			}
			return out
		}(C.QBluetoothServiceInfo_ServiceClassUuids(ptr.Pointer()))
	}
	return make([]*QBluetoothUuid, 0)
}

func (ptr *QBluetoothServiceInfo) ServiceDescription() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QBluetoothServiceInfo_ServiceDescription(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothServiceInfo) ServiceName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QBluetoothServiceInfo_ServiceName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothServiceInfo) ServiceProvider() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QBluetoothServiceInfo_ServiceProvider(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothServiceInfo) Attribute(attributeId uint16) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QBluetoothServiceInfo_Attribute(ptr.Pointer(), C.ushort(attributeId)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothServiceInfo) Contains(attributeId uint16) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothServiceInfo_Contains(ptr.Pointer(), C.ushort(attributeId)) != 0
	}
	return false
}

func (ptr *QBluetoothServiceInfo) IsComplete() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothServiceInfo_IsComplete(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothServiceInfo) IsRegistered() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothServiceInfo_IsRegistered(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothServiceInfo) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothServiceInfo_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothServiceInfo) ProtocolServiceMultiplexer() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QBluetoothServiceInfo_ProtocolServiceMultiplexer(ptr.Pointer())))
	}
	return 0
}

func (ptr *QBluetoothServiceInfo) ServerChannel() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QBluetoothServiceInfo_ServerChannel(ptr.Pointer())))
	}
	return 0
}

func (ptr *QBluetoothServiceInfo) ServiceAvailability() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QBluetoothServiceInfo_ServiceAvailability(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothServiceInfo) __serviceClassUuids_atList(i int) *QBluetoothUuid {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothUuidFromPointer(C.QBluetoothServiceInfo___serviceClassUuids_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothServiceInfo) __serviceClassUuids_setList(i QBluetoothUuid_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceInfo___serviceClassUuids_setList(ptr.Pointer(), PointerFromQBluetoothUuid(i))
	}
}

func (ptr *QBluetoothServiceInfo) __serviceClassUuids_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QBluetoothServiceInfo___serviceClassUuids_newList(ptr.Pointer()))
}

type QBluetoothSocket struct {
	core.QIODevice
}

type QBluetoothSocket_ITF interface {
	core.QIODevice_ITF
	QBluetoothSocket_PTR() *QBluetoothSocket
}

func (ptr *QBluetoothSocket) QBluetoothSocket_PTR() *QBluetoothSocket {
	return ptr
}

func (ptr *QBluetoothSocket) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QIODevice_PTR().Pointer()
	}
	return nil
}

func (ptr *QBluetoothSocket) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QIODevice_PTR().SetPointer(p)
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

//go:generate stringer -type=QBluetoothSocket__SocketError
//QBluetoothSocket::SocketError
type QBluetoothSocket__SocketError int64

const (
	QBluetoothSocket__NoSocketError            QBluetoothSocket__SocketError = QBluetoothSocket__SocketError(-2)
	QBluetoothSocket__UnknownSocketError       QBluetoothSocket__SocketError = QBluetoothSocket__SocketError(network.QAbstractSocket__UnknownSocketError)
	QBluetoothSocket__HostNotFoundError        QBluetoothSocket__SocketError = QBluetoothSocket__SocketError(network.QAbstractSocket__HostNotFoundError)
	QBluetoothSocket__ServiceNotFoundError     QBluetoothSocket__SocketError = QBluetoothSocket__SocketError(network.QAbstractSocket__SocketAddressNotAvailableError)
	QBluetoothSocket__NetworkError             QBluetoothSocket__SocketError = QBluetoothSocket__SocketError(network.QAbstractSocket__NetworkError)
	QBluetoothSocket__UnsupportedProtocolError QBluetoothSocket__SocketError = QBluetoothSocket__SocketError(8)
	QBluetoothSocket__OperationError           QBluetoothSocket__SocketError = QBluetoothSocket__SocketError(network.QAbstractSocket__OperationError)
)

//go:generate stringer -type=QBluetoothSocket__SocketState
//QBluetoothSocket::SocketState
type QBluetoothSocket__SocketState int64

const (
	QBluetoothSocket__UnconnectedState   QBluetoothSocket__SocketState = QBluetoothSocket__SocketState(network.QAbstractSocket__UnconnectedState)
	QBluetoothSocket__ServiceLookupState QBluetoothSocket__SocketState = QBluetoothSocket__SocketState(network.QAbstractSocket__HostLookupState)
	QBluetoothSocket__ConnectingState    QBluetoothSocket__SocketState = QBluetoothSocket__SocketState(network.QAbstractSocket__ConnectingState)
	QBluetoothSocket__ConnectedState     QBluetoothSocket__SocketState = QBluetoothSocket__SocketState(network.QAbstractSocket__ConnectedState)
	QBluetoothSocket__BoundState         QBluetoothSocket__SocketState = QBluetoothSocket__SocketState(network.QAbstractSocket__BoundState)
	QBluetoothSocket__ClosingState       QBluetoothSocket__SocketState = QBluetoothSocket__SocketState(network.QAbstractSocket__ClosingState)
	QBluetoothSocket__ListeningState     QBluetoothSocket__SocketState = QBluetoothSocket__SocketState(network.QAbstractSocket__ListeningState)
)

//export callbackQBluetoothSocket_Connected
func callbackQBluetoothSocket_Connected(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connected"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBluetoothSocket) ConnectConnected(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "connected") {
			C.QBluetoothSocket_ConnectConnected(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "connected"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "connected", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "connected", f)
		}
	}
}

func (ptr *QBluetoothSocket) DisconnectConnected() {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DisconnectConnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "connected")
	}
}

func (ptr *QBluetoothSocket) Connected() {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_Connected(ptr.Pointer())
	}
}

//export callbackQBluetoothSocket_Disconnected
func callbackQBluetoothSocket_Disconnected(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnected"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QBluetoothSocket) ConnectDisconnected(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "disconnected") {
			C.QBluetoothSocket_ConnectDisconnected(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "disconnected"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "disconnected", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "disconnected", f)
		}
	}
}

func (ptr *QBluetoothSocket) DisconnectDisconnected() {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DisconnectDisconnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "disconnected")
	}
}

func (ptr *QBluetoothSocket) Disconnected() {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_Disconnected(ptr.Pointer())
	}
}

//export callbackQBluetoothSocket_Error2
func callbackQBluetoothSocket_Error2(ptr unsafe.Pointer, error C.longlong) {
	if signal := qt.GetSignal(ptr, "error2"); signal != nil {
		signal.(func(QBluetoothSocket__SocketError))(QBluetoothSocket__SocketError(error))
	}

}

func (ptr *QBluetoothSocket) ConnectError2(f func(error QBluetoothSocket__SocketError)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "error2") {
			C.QBluetoothSocket_ConnectError2(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "error2"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "error2", func(error QBluetoothSocket__SocketError) {
				signal.(func(QBluetoothSocket__SocketError))(error)
				f(error)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "error2", f)
		}
	}
}

func (ptr *QBluetoothSocket) DisconnectError2() {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "error2")
	}
}

func (ptr *QBluetoothSocket) Error2(error QBluetoothSocket__SocketError) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_Error2(ptr.Pointer(), C.longlong(error))
	}
}

//export callbackQBluetoothSocket_StateChanged
func callbackQBluetoothSocket_StateChanged(ptr unsafe.Pointer, state C.longlong) {
	if signal := qt.GetSignal(ptr, "stateChanged"); signal != nil {
		signal.(func(QBluetoothSocket__SocketState))(QBluetoothSocket__SocketState(state))
	}

}

func (ptr *QBluetoothSocket) ConnectStateChanged(f func(state QBluetoothSocket__SocketState)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "stateChanged") {
			C.QBluetoothSocket_ConnectStateChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "stateChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "stateChanged", func(state QBluetoothSocket__SocketState) {
				signal.(func(QBluetoothSocket__SocketState))(state)
				f(state)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stateChanged", f)
		}
	}
}

func (ptr *QBluetoothSocket) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "stateChanged")
	}
}

func (ptr *QBluetoothSocket) StateChanged(state QBluetoothSocket__SocketState) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_StateChanged(ptr.Pointer(), C.longlong(state))
	}
}

func NewQBluetoothSocket(socketType QBluetoothServiceInfo__Protocol, parent core.QObject_ITF) *QBluetoothSocket {
	var tmpValue = NewQBluetoothSocketFromPointer(C.QBluetoothSocket_NewQBluetoothSocket(C.longlong(socketType), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQBluetoothSocket2(parent core.QObject_ITF) *QBluetoothSocket {
	var tmpValue = NewQBluetoothSocketFromPointer(C.QBluetoothSocket_NewQBluetoothSocket2(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QBluetoothSocket) SetSocketDescriptor(socketDescriptor int, socketType QBluetoothServiceInfo__Protocol, socketState QBluetoothSocket__SocketState, openMode core.QIODevice__OpenModeFlag) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_SetSocketDescriptor(ptr.Pointer(), C.int(int32(socketDescriptor)), C.longlong(socketType), C.longlong(socketState), C.longlong(openMode)) != 0
	}
	return false
}

//export callbackQBluetoothSocket_ReadData
func callbackQBluetoothSocket_ReadData(ptr unsafe.Pointer, data C.struct_QtBluetooth_PackedString, maxSize C.longlong) C.longlong {
	if signal := qt.GetSignal(ptr, "readData"); signal != nil {
		var retS = cGoUnpackString(data)
		var ret = C.longlong(signal.(func(*string, int64) int64)(&retS, int64(maxSize)))
		if ret > 0 {
			C.memcpy(unsafe.Pointer(data.data), unsafe.Pointer(C.CString(retS)), C.size_t(ret))
		}
		return ret
	}
	var retS = cGoUnpackString(data)
	var ret = C.longlong(NewQBluetoothSocketFromPointer(ptr).ReadDataDefault(&retS, int64(maxSize)))
	if ret > 0 {
		C.memcpy(unsafe.Pointer(data.data), unsafe.Pointer(C.CString(retS)), C.size_t(ret))
	}
	return ret
}

func (ptr *QBluetoothSocket) ConnectReadData(f func(data *string, maxSize int64) int64) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "readData"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "readData", func(data *string, maxSize int64) int64 {
				signal.(func(*string, int64) int64)(data, maxSize)
				return f(data, maxSize)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "readData", f)
		}
	}
}

func (ptr *QBluetoothSocket) DisconnectReadData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "readData")
	}
}

func (ptr *QBluetoothSocket) ReadData(data *string, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		var dataC = C.CString(strings.Repeat("0", int(maxSize)))
		defer C.free(unsafe.Pointer(dataC))
		var ret = int64(C.QBluetoothSocket_ReadData(ptr.Pointer(), dataC, C.longlong(maxSize)))
		if ret > 0 {
			*data = C.GoStringN(dataC, C.int(ret))
		}
		return ret
	}
	return 0
}

func (ptr *QBluetoothSocket) ReadDataDefault(data *string, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		var dataC = C.CString(strings.Repeat("0", int(maxSize)))
		defer C.free(unsafe.Pointer(dataC))
		var ret = int64(C.QBluetoothSocket_ReadDataDefault(ptr.Pointer(), dataC, C.longlong(maxSize)))
		if ret > 0 {
			*data = C.GoStringN(dataC, C.int(ret))
		}
		return ret
	}
	return 0
}

//export callbackQBluetoothSocket_WriteData
func callbackQBluetoothSocket_WriteData(ptr unsafe.Pointer, data C.struct_QtBluetooth_PackedString, maxSize C.longlong) C.longlong {
	if signal := qt.GetSignal(ptr, "writeData"); signal != nil {
		return C.longlong(signal.(func(string, int64) int64)(cGoUnpackString(data), int64(maxSize)))
	}

	return C.longlong(NewQBluetoothSocketFromPointer(ptr).WriteDataDefault(cGoUnpackString(data), int64(maxSize)))
}

func (ptr *QBluetoothSocket) ConnectWriteData(f func(data string, maxSize int64) int64) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "writeData"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "writeData", func(data string, maxSize int64) int64 {
				signal.(func(string, int64) int64)(data, maxSize)
				return f(data, maxSize)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "writeData", f)
		}
	}
}

func (ptr *QBluetoothSocket) DisconnectWriteData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "writeData")
	}
}

func (ptr *QBluetoothSocket) WriteData(data string, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if data != "" {
			dataC = C.CString(data)
			defer C.free(unsafe.Pointer(dataC))
		}
		return int64(C.QBluetoothSocket_WriteData(ptr.Pointer(), dataC, C.longlong(maxSize)))
	}
	return 0
}

func (ptr *QBluetoothSocket) WriteDataDefault(data string, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if data != "" {
			dataC = C.CString(data)
			defer C.free(unsafe.Pointer(dataC))
		}
		return int64(C.QBluetoothSocket_WriteDataDefault(ptr.Pointer(), dataC, C.longlong(maxSize)))
	}
	return 0
}

func (ptr *QBluetoothSocket) Abort() {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_Abort(ptr.Pointer())
	}
}

//export callbackQBluetoothSocket_Close
func callbackQBluetoothSocket_Close(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		signal.(func())()
	} else {
		NewQBluetoothSocketFromPointer(ptr).CloseDefault()
	}
}

func (ptr *QBluetoothSocket) CloseDefault() {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_CloseDefault(ptr.Pointer())
	}
}

func (ptr *QBluetoothSocket) ConnectToService2(address QBluetoothAddress_ITF, uuid QBluetoothUuid_ITF, openMode core.QIODevice__OpenModeFlag) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_ConnectToService2(ptr.Pointer(), PointerFromQBluetoothAddress(address), PointerFromQBluetoothUuid(uuid), C.longlong(openMode))
	}
}

func (ptr *QBluetoothSocket) ConnectToService3(address QBluetoothAddress_ITF, port uint16, openMode core.QIODevice__OpenModeFlag) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_ConnectToService3(ptr.Pointer(), PointerFromQBluetoothAddress(address), C.ushort(port), C.longlong(openMode))
	}
}

func (ptr *QBluetoothSocket) ConnectToService(service QBluetoothServiceInfo_ITF, openMode core.QIODevice__OpenModeFlag) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_ConnectToService(ptr.Pointer(), PointerFromQBluetoothServiceInfo(service), C.longlong(openMode))
	}
}

func (ptr *QBluetoothSocket) DisconnectFromService() {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DisconnectFromService(ptr.Pointer())
	}
}

func (ptr *QBluetoothSocket) DoDeviceDiscovery(service QBluetoothServiceInfo_ITF, openMode core.QIODevice__OpenModeFlag) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DoDeviceDiscovery(ptr.Pointer(), PointerFromQBluetoothServiceInfo(service), C.longlong(openMode))
	}
}

func (ptr *QBluetoothSocket) SetPreferredSecurityFlags(flags QBluetooth__Security) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_SetPreferredSecurityFlags(ptr.Pointer(), C.longlong(flags))
	}
}

func (ptr *QBluetoothSocket) SetSocketError(error_ QBluetoothSocket__SocketError) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_SetSocketError(ptr.Pointer(), C.longlong(error_))
	}
}

func (ptr *QBluetoothSocket) SetSocketState(state QBluetoothSocket__SocketState) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_SetSocketState(ptr.Pointer(), C.longlong(state))
	}
}

//export callbackQBluetoothSocket_DestroyQBluetoothSocket
func callbackQBluetoothSocket_DestroyQBluetoothSocket(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QBluetoothSocket"); signal != nil {
		signal.(func())()
	} else {
		NewQBluetoothSocketFromPointer(ptr).DestroyQBluetoothSocketDefault()
	}
}

func (ptr *QBluetoothSocket) ConnectDestroyQBluetoothSocket(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QBluetoothSocket"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "~QBluetoothSocket", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QBluetoothSocket", f)
		}
	}
}

func (ptr *QBluetoothSocket) DisconnectDestroyQBluetoothSocket() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QBluetoothSocket")
	}
}

func (ptr *QBluetoothSocket) DestroyQBluetoothSocket() {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DestroyQBluetoothSocket(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QBluetoothSocket) DestroyQBluetoothSocketDefault() {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DestroyQBluetoothSocketDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QBluetoothSocket) PreferredSecurityFlags() QBluetooth__Security {
	if ptr.Pointer() != nil {
		return QBluetooth__Security(C.QBluetoothSocket_PreferredSecurityFlags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothSocket) LocalAddress() *QBluetoothAddress {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothAddressFromPointer(C.QBluetoothSocket_LocalAddress(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothSocket) PeerAddress() *QBluetoothAddress {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothAddressFromPointer(C.QBluetoothSocket_PeerAddress(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothSocket) SocketType() QBluetoothServiceInfo__Protocol {
	if ptr.Pointer() != nil {
		return QBluetoothServiceInfo__Protocol(C.QBluetoothSocket_SocketType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothSocket) LocalName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QBluetoothSocket_LocalName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothSocket) PeerName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QBluetoothSocket_PeerName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothSocket) Error() QBluetoothSocket__SocketError {
	if ptr.Pointer() != nil {
		return QBluetoothSocket__SocketError(C.QBluetoothSocket_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothSocket) State() QBluetoothSocket__SocketState {
	if ptr.Pointer() != nil {
		return QBluetoothSocket__SocketState(C.QBluetoothSocket_State(ptr.Pointer()))
	}
	return 0
}

//export callbackQBluetoothSocket_CanReadLine
func callbackQBluetoothSocket_CanReadLine(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canReadLine"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothSocketFromPointer(ptr).CanReadLineDefault())))
}

func (ptr *QBluetoothSocket) CanReadLineDefault() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_CanReadLineDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQBluetoothSocket_IsSequential
func callbackQBluetoothSocket_IsSequential(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isSequential"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothSocketFromPointer(ptr).IsSequentialDefault())))
}

func (ptr *QBluetoothSocket) IsSequentialDefault() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_IsSequentialDefault(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothSocket) SocketDescriptor() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QBluetoothSocket_SocketDescriptor(ptr.Pointer())))
	}
	return 0
}

//export callbackQBluetoothSocket_BytesAvailable
func callbackQBluetoothSocket_BytesAvailable(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "bytesAvailable"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQBluetoothSocketFromPointer(ptr).BytesAvailableDefault())
}

func (ptr *QBluetoothSocket) BytesAvailableDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QBluetoothSocket_BytesAvailableDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQBluetoothSocket_BytesToWrite
func callbackQBluetoothSocket_BytesToWrite(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "bytesToWrite"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQBluetoothSocketFromPointer(ptr).BytesToWriteDefault())
}

func (ptr *QBluetoothSocket) BytesToWriteDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QBluetoothSocket_BytesToWriteDefault(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothSocket) LocalPort() uint16 {
	if ptr.Pointer() != nil {
		return uint16(C.QBluetoothSocket_LocalPort(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothSocket) PeerPort() uint16 {
	if ptr.Pointer() != nil {
		return uint16(C.QBluetoothSocket_PeerPort(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothSocket) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QBluetoothSocket___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothSocket) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QBluetoothSocket) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QBluetoothSocket___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QBluetoothSocket) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QBluetoothSocket___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothSocket) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QBluetoothSocket) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QBluetoothSocket___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QBluetoothSocket) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QBluetoothSocket___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothSocket) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QBluetoothSocket) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.QBluetoothSocket___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QBluetoothSocket) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QBluetoothSocket___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothSocket) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QBluetoothSocket) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QBluetoothSocket___findChildren_newList(ptr.Pointer()))
}

func (ptr *QBluetoothSocket) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QBluetoothSocket___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothSocket) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QBluetoothSocket) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QBluetoothSocket___children_newList(ptr.Pointer()))
}

//export callbackQBluetoothSocket_Open
func callbackQBluetoothSocket_Open(ptr unsafe.Pointer, mode C.longlong) C.char {
	if signal := qt.GetSignal(ptr, "open"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(core.QIODevice__OpenModeFlag) bool)(core.QIODevice__OpenModeFlag(mode)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothSocketFromPointer(ptr).OpenDefault(core.QIODevice__OpenModeFlag(mode)))))
}

func (ptr *QBluetoothSocket) OpenDefault(mode core.QIODevice__OpenModeFlag) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_OpenDefault(ptr.Pointer(), C.longlong(mode)) != 0
	}
	return false
}

//export callbackQBluetoothSocket_Reset
func callbackQBluetoothSocket_Reset(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "reset"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothSocketFromPointer(ptr).ResetDefault())))
}

func (ptr *QBluetoothSocket) ResetDefault() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_ResetDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQBluetoothSocket_Seek
func callbackQBluetoothSocket_Seek(ptr unsafe.Pointer, pos C.longlong) C.char {
	if signal := qt.GetSignal(ptr, "seek"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int64) bool)(int64(pos)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothSocketFromPointer(ptr).SeekDefault(int64(pos)))))
}

func (ptr *QBluetoothSocket) SeekDefault(pos int64) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_SeekDefault(ptr.Pointer(), C.longlong(pos)) != 0
	}
	return false
}

//export callbackQBluetoothSocket_WaitForBytesWritten
func callbackQBluetoothSocket_WaitForBytesWritten(ptr unsafe.Pointer, msecs C.int) C.char {
	if signal := qt.GetSignal(ptr, "waitForBytesWritten"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothSocketFromPointer(ptr).WaitForBytesWrittenDefault(int(int32(msecs))))))
}

func (ptr *QBluetoothSocket) WaitForBytesWrittenDefault(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_WaitForBytesWrittenDefault(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

//export callbackQBluetoothSocket_WaitForReadyRead
func callbackQBluetoothSocket_WaitForReadyRead(ptr unsafe.Pointer, msecs C.int) C.char {
	if signal := qt.GetSignal(ptr, "waitForReadyRead"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(int) bool)(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothSocketFromPointer(ptr).WaitForReadyReadDefault(int(int32(msecs))))))
}

func (ptr *QBluetoothSocket) WaitForReadyReadDefault(msecs int) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_WaitForReadyReadDefault(ptr.Pointer(), C.int(int32(msecs))) != 0
	}
	return false
}

//export callbackQBluetoothSocket_ReadLineData
func callbackQBluetoothSocket_ReadLineData(ptr unsafe.Pointer, data C.struct_QtBluetooth_PackedString, maxSize C.longlong) C.longlong {
	if signal := qt.GetSignal(ptr, "readLineData"); signal != nil {
		return C.longlong(signal.(func(string, int64) int64)(cGoUnpackString(data), int64(maxSize)))
	}

	return C.longlong(NewQBluetoothSocketFromPointer(ptr).ReadLineDataDefault(cGoUnpackString(data), int64(maxSize)))
}

func (ptr *QBluetoothSocket) ReadLineDataDefault(data string, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if data != "" {
			dataC = C.CString(data)
			defer C.free(unsafe.Pointer(dataC))
		}
		return int64(C.QBluetoothSocket_ReadLineDataDefault(ptr.Pointer(), dataC, C.longlong(maxSize)))
	}
	return 0
}

//export callbackQBluetoothSocket_AboutToClose
func callbackQBluetoothSocket_AboutToClose(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "aboutToClose"); signal != nil {
		signal.(func())()
	}

}

//export callbackQBluetoothSocket_BytesWritten
func callbackQBluetoothSocket_BytesWritten(ptr unsafe.Pointer, bytes C.longlong) {
	if signal := qt.GetSignal(ptr, "bytesWritten"); signal != nil {
		signal.(func(int64))(int64(bytes))
	}

}

//export callbackQBluetoothSocket_ChannelBytesWritten
func callbackQBluetoothSocket_ChannelBytesWritten(ptr unsafe.Pointer, channel C.int, bytes C.longlong) {
	if signal := qt.GetSignal(ptr, "channelBytesWritten"); signal != nil {
		signal.(func(int, int64))(int(int32(channel)), int64(bytes))
	}

}

//export callbackQBluetoothSocket_ChannelReadyRead
func callbackQBluetoothSocket_ChannelReadyRead(ptr unsafe.Pointer, channel C.int) {
	if signal := qt.GetSignal(ptr, "channelReadyRead"); signal != nil {
		signal.(func(int))(int(int32(channel)))
	}

}

//export callbackQBluetoothSocket_ReadChannelFinished
func callbackQBluetoothSocket_ReadChannelFinished(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "readChannelFinished"); signal != nil {
		signal.(func())()
	}

}

//export callbackQBluetoothSocket_ReadyRead
func callbackQBluetoothSocket_ReadyRead(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "readyRead"); signal != nil {
		signal.(func())()
	}

}

//export callbackQBluetoothSocket_AtEnd
func callbackQBluetoothSocket_AtEnd(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "atEnd"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothSocketFromPointer(ptr).AtEndDefault())))
}

func (ptr *QBluetoothSocket) AtEndDefault() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_AtEndDefault(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQBluetoothSocket_Pos
func callbackQBluetoothSocket_Pos(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "pos"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQBluetoothSocketFromPointer(ptr).PosDefault())
}

func (ptr *QBluetoothSocket) PosDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QBluetoothSocket_PosDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQBluetoothSocket_Size
func callbackQBluetoothSocket_Size(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "size"); signal != nil {
		return C.longlong(signal.(func() int64)())
	}

	return C.longlong(NewQBluetoothSocketFromPointer(ptr).SizeDefault())
}

func (ptr *QBluetoothSocket) SizeDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QBluetoothSocket_SizeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQBluetoothSocket_Event
func callbackQBluetoothSocket_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothSocketFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QBluetoothSocket) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQBluetoothSocket_EventFilter
func callbackQBluetoothSocket_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothSocketFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QBluetoothSocket) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothSocket_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQBluetoothSocket_ChildEvent
func callbackQBluetoothSocket_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQBluetoothSocketFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QBluetoothSocket) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQBluetoothSocket_ConnectNotify
func callbackQBluetoothSocket_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBluetoothSocketFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBluetoothSocket) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBluetoothSocket_CustomEvent
func callbackQBluetoothSocket_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQBluetoothSocketFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QBluetoothSocket) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQBluetoothSocket_DeleteLater
func callbackQBluetoothSocket_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQBluetoothSocketFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QBluetoothSocket) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQBluetoothSocket_Destroyed
func callbackQBluetoothSocket_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQBluetoothSocket_DisconnectNotify
func callbackQBluetoothSocket_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBluetoothSocketFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBluetoothSocket) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBluetoothSocket_ObjectNameChanged
func callbackQBluetoothSocket_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtBluetooth_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQBluetoothSocket_TimerEvent
func callbackQBluetoothSocket_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQBluetoothSocketFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QBluetoothSocket) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQBluetoothSocket_MetaObject
func callbackQBluetoothSocket_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQBluetoothSocketFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QBluetoothSocket) MetaObjectDefault() *core.QMetaObject {
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

func (ptr *QBluetoothTransferManager) QBluetoothTransferManager_PTR() *QBluetoothTransferManager {
	return ptr
}

func (ptr *QBluetoothTransferManager) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QBluetoothTransferManager) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
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
func NewQBluetoothTransferManager(parent core.QObject_ITF) *QBluetoothTransferManager {
	var tmpValue = NewQBluetoothTransferManagerFromPointer(C.QBluetoothTransferManager_NewQBluetoothTransferManager(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QBluetoothTransferManager) Put(request QBluetoothTransferRequest_ITF, data core.QIODevice_ITF) *QBluetoothTransferReply {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothTransferReplyFromPointer(C.QBluetoothTransferManager_Put(ptr.Pointer(), PointerFromQBluetoothTransferRequest(request), core.PointerFromQIODevice(data)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQBluetoothTransferManager_Finished
func callbackQBluetoothTransferManager_Finished(ptr unsafe.Pointer, reply unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "finished"); signal != nil {
		signal.(func(*QBluetoothTransferReply))(NewQBluetoothTransferReplyFromPointer(reply))
	}

}

func (ptr *QBluetoothTransferManager) ConnectFinished(f func(reply *QBluetoothTransferReply)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "finished") {
			C.QBluetoothTransferManager_ConnectFinished(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "finished"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "finished", func(reply *QBluetoothTransferReply) {
				signal.(func(*QBluetoothTransferReply))(reply)
				f(reply)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "finished", f)
		}
	}
}

func (ptr *QBluetoothTransferManager) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "finished")
	}
}

func (ptr *QBluetoothTransferManager) Finished(reply QBluetoothTransferReply_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_Finished(ptr.Pointer(), PointerFromQBluetoothTransferReply(reply))
	}
}

func (ptr *QBluetoothTransferManager) DestroyQBluetoothTransferManager() {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_DestroyQBluetoothTransferManager(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QBluetoothTransferManager) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QBluetoothTransferManager___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothTransferManager) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QBluetoothTransferManager) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QBluetoothTransferManager___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QBluetoothTransferManager) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QBluetoothTransferManager___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothTransferManager) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QBluetoothTransferManager) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QBluetoothTransferManager___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QBluetoothTransferManager) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QBluetoothTransferManager___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothTransferManager) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QBluetoothTransferManager) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.QBluetoothTransferManager___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QBluetoothTransferManager) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QBluetoothTransferManager___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothTransferManager) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QBluetoothTransferManager) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QBluetoothTransferManager___findChildren_newList(ptr.Pointer()))
}

func (ptr *QBluetoothTransferManager) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QBluetoothTransferManager___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothTransferManager) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QBluetoothTransferManager) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QBluetoothTransferManager___children_newList(ptr.Pointer()))
}

//export callbackQBluetoothTransferManager_Event
func callbackQBluetoothTransferManager_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothTransferManagerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QBluetoothTransferManager) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothTransferManager_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQBluetoothTransferManager_EventFilter
func callbackQBluetoothTransferManager_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothTransferManagerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QBluetoothTransferManager) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothTransferManager_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQBluetoothTransferManager_ChildEvent
func callbackQBluetoothTransferManager_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQBluetoothTransferManagerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QBluetoothTransferManager) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQBluetoothTransferManager_ConnectNotify
func callbackQBluetoothTransferManager_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBluetoothTransferManagerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBluetoothTransferManager) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBluetoothTransferManager_CustomEvent
func callbackQBluetoothTransferManager_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQBluetoothTransferManagerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QBluetoothTransferManager) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQBluetoothTransferManager_DeleteLater
func callbackQBluetoothTransferManager_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQBluetoothTransferManagerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QBluetoothTransferManager) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQBluetoothTransferManager_Destroyed
func callbackQBluetoothTransferManager_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQBluetoothTransferManager_DisconnectNotify
func callbackQBluetoothTransferManager_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBluetoothTransferManagerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBluetoothTransferManager) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBluetoothTransferManager_ObjectNameChanged
func callbackQBluetoothTransferManager_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtBluetooth_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQBluetoothTransferManager_TimerEvent
func callbackQBluetoothTransferManager_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQBluetoothTransferManagerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QBluetoothTransferManager) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQBluetoothTransferManager_MetaObject
func callbackQBluetoothTransferManager_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQBluetoothTransferManagerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QBluetoothTransferManager) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QBluetoothTransferManager_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QBluetoothTransferReply struct {
	core.QObject
}

type QBluetoothTransferReply_ITF interface {
	core.QObject_ITF
	QBluetoothTransferReply_PTR() *QBluetoothTransferReply
}

func (ptr *QBluetoothTransferReply) QBluetoothTransferReply_PTR() *QBluetoothTransferReply {
	return ptr
}

func (ptr *QBluetoothTransferReply) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QBluetoothTransferReply) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
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

//go:generate stringer -type=QBluetoothTransferReply__TransferError
//QBluetoothTransferReply::TransferError
type QBluetoothTransferReply__TransferError int64

const (
	QBluetoothTransferReply__NoError                   QBluetoothTransferReply__TransferError = QBluetoothTransferReply__TransferError(0)
	QBluetoothTransferReply__UnknownError              QBluetoothTransferReply__TransferError = QBluetoothTransferReply__TransferError(1)
	QBluetoothTransferReply__FileNotFoundError         QBluetoothTransferReply__TransferError = QBluetoothTransferReply__TransferError(2)
	QBluetoothTransferReply__HostNotFoundError         QBluetoothTransferReply__TransferError = QBluetoothTransferReply__TransferError(3)
	QBluetoothTransferReply__UserCanceledTransferError QBluetoothTransferReply__TransferError = QBluetoothTransferReply__TransferError(4)
	QBluetoothTransferReply__IODeviceNotReadableError  QBluetoothTransferReply__TransferError = QBluetoothTransferReply__TransferError(5)
	QBluetoothTransferReply__ResourceBusyError         QBluetoothTransferReply__TransferError = QBluetoothTransferReply__TransferError(6)
	QBluetoothTransferReply__SessionError              QBluetoothTransferReply__TransferError = QBluetoothTransferReply__TransferError(7)
)

func NewQBluetoothTransferReply(parent core.QObject_ITF) *QBluetoothTransferReply {
	var tmpValue = NewQBluetoothTransferReplyFromPointer(C.QBluetoothTransferReply_NewQBluetoothTransferReply(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQBluetoothTransferReply_Abort
func callbackQBluetoothTransferReply_Abort(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "abort"); signal != nil {
		signal.(func())()
	} else {
		NewQBluetoothTransferReplyFromPointer(ptr).AbortDefault()
	}
}

func (ptr *QBluetoothTransferReply) ConnectAbort(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "abort"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "abort", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "abort", f)
		}
	}
}

func (ptr *QBluetoothTransferReply) DisconnectAbort() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "abort")
	}
}

func (ptr *QBluetoothTransferReply) Abort() {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_Abort(ptr.Pointer())
	}
}

func (ptr *QBluetoothTransferReply) AbortDefault() {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_AbortDefault(ptr.Pointer())
	}
}

//export callbackQBluetoothTransferReply_Error2
func callbackQBluetoothTransferReply_Error2(ptr unsafe.Pointer, errorType C.longlong) {
	if signal := qt.GetSignal(ptr, "error2"); signal != nil {
		signal.(func(QBluetoothTransferReply__TransferError))(QBluetoothTransferReply__TransferError(errorType))
	}

}

func (ptr *QBluetoothTransferReply) ConnectError2(f func(errorType QBluetoothTransferReply__TransferError)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "error2") {
			C.QBluetoothTransferReply_ConnectError2(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "error2"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "error2", func(errorType QBluetoothTransferReply__TransferError) {
				signal.(func(QBluetoothTransferReply__TransferError))(errorType)
				f(errorType)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "error2", f)
		}
	}
}

func (ptr *QBluetoothTransferReply) DisconnectError2() {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "error2")
	}
}

func (ptr *QBluetoothTransferReply) Error2(errorType QBluetoothTransferReply__TransferError) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_Error2(ptr.Pointer(), C.longlong(errorType))
	}
}

//export callbackQBluetoothTransferReply_Finished
func callbackQBluetoothTransferReply_Finished(ptr unsafe.Pointer, reply unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "finished"); signal != nil {
		signal.(func(*QBluetoothTransferReply))(NewQBluetoothTransferReplyFromPointer(reply))
	}

}

func (ptr *QBluetoothTransferReply) ConnectFinished(f func(reply *QBluetoothTransferReply)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "finished") {
			C.QBluetoothTransferReply_ConnectFinished(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "finished"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "finished", func(reply *QBluetoothTransferReply) {
				signal.(func(*QBluetoothTransferReply))(reply)
				f(reply)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "finished", f)
		}
	}
}

func (ptr *QBluetoothTransferReply) DisconnectFinished() {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_DisconnectFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "finished")
	}
}

func (ptr *QBluetoothTransferReply) Finished(reply QBluetoothTransferReply_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_Finished(ptr.Pointer(), PointerFromQBluetoothTransferReply(reply))
	}
}

func (ptr *QBluetoothTransferReply) SetManager(manager QBluetoothTransferManager_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_SetManager(ptr.Pointer(), PointerFromQBluetoothTransferManager(manager))
	}
}

func (ptr *QBluetoothTransferReply) SetRequest(request QBluetoothTransferRequest_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_SetRequest(ptr.Pointer(), PointerFromQBluetoothTransferRequest(request))
	}
}

//export callbackQBluetoothTransferReply_TransferProgress
func callbackQBluetoothTransferReply_TransferProgress(ptr unsafe.Pointer, bytesTransferred C.longlong, bytesTotal C.longlong) {
	if signal := qt.GetSignal(ptr, "transferProgress"); signal != nil {
		signal.(func(int64, int64))(int64(bytesTransferred), int64(bytesTotal))
	}

}

func (ptr *QBluetoothTransferReply) ConnectTransferProgress(f func(bytesTransferred int64, bytesTotal int64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "transferProgress") {
			C.QBluetoothTransferReply_ConnectTransferProgress(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "transferProgress"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "transferProgress", func(bytesTransferred int64, bytesTotal int64) {
				signal.(func(int64, int64))(bytesTransferred, bytesTotal)
				f(bytesTransferred, bytesTotal)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "transferProgress", f)
		}
	}
}

func (ptr *QBluetoothTransferReply) DisconnectTransferProgress() {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_DisconnectTransferProgress(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "transferProgress")
	}
}

func (ptr *QBluetoothTransferReply) TransferProgress(bytesTransferred int64, bytesTotal int64) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_TransferProgress(ptr.Pointer(), C.longlong(bytesTransferred), C.longlong(bytesTotal))
	}
}

func (ptr *QBluetoothTransferReply) DestroyQBluetoothTransferReply() {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_DestroyQBluetoothTransferReply(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QBluetoothTransferReply) Manager() *QBluetoothTransferManager {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothTransferManagerFromPointer(C.QBluetoothTransferReply_Manager(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothTransferReply) Request() *QBluetoothTransferRequest {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothTransferRequestFromPointer(C.QBluetoothTransferReply_Request(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothTransferRequest).DestroyQBluetoothTransferRequest)
		return tmpValue
	}
	return nil
}

//export callbackQBluetoothTransferReply_ErrorString
func callbackQBluetoothTransferReply_ErrorString(ptr unsafe.Pointer) C.struct_QtBluetooth_PackedString {
	if signal := qt.GetSignal(ptr, "errorString"); signal != nil {
		tempVal := signal.(func() string)()
		return C.struct_QtBluetooth_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_QtBluetooth_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QBluetoothTransferReply) ConnectErrorString(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "errorString"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "errorString", func() string {
				signal.(func() string)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "errorString", f)
		}
	}
}

func (ptr *QBluetoothTransferReply) DisconnectErrorString() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "errorString")
	}
}

func (ptr *QBluetoothTransferReply) ErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QBluetoothTransferReply_ErrorString(ptr.Pointer()))
	}
	return ""
}

//export callbackQBluetoothTransferReply_Error
func callbackQBluetoothTransferReply_Error(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "error"); signal != nil {
		return C.longlong(signal.(func() QBluetoothTransferReply__TransferError)())
	}

	return C.longlong(0)
}

func (ptr *QBluetoothTransferReply) ConnectError(f func() QBluetoothTransferReply__TransferError) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "error"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "error", func() QBluetoothTransferReply__TransferError {
				signal.(func() QBluetoothTransferReply__TransferError)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "error", f)
		}
	}
}

func (ptr *QBluetoothTransferReply) DisconnectError() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "error")
	}
}

func (ptr *QBluetoothTransferReply) Error() QBluetoothTransferReply__TransferError {
	if ptr.Pointer() != nil {
		return QBluetoothTransferReply__TransferError(C.QBluetoothTransferReply_Error(ptr.Pointer()))
	}
	return 0
}

//export callbackQBluetoothTransferReply_IsFinished
func callbackQBluetoothTransferReply_IsFinished(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isFinished"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QBluetoothTransferReply) ConnectIsFinished(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "isFinished"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "isFinished", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "isFinished", f)
		}
	}
}

func (ptr *QBluetoothTransferReply) DisconnectIsFinished() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "isFinished")
	}
}

func (ptr *QBluetoothTransferReply) IsFinished() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothTransferReply_IsFinished(ptr.Pointer()) != 0
	}
	return false
}

//export callbackQBluetoothTransferReply_IsRunning
func callbackQBluetoothTransferReply_IsRunning(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isRunning"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func() bool)())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QBluetoothTransferReply) ConnectIsRunning(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "isRunning"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "isRunning", func() bool {
				signal.(func() bool)()
				return f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "isRunning", f)
		}
	}
}

func (ptr *QBluetoothTransferReply) DisconnectIsRunning() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "isRunning")
	}
}

func (ptr *QBluetoothTransferReply) IsRunning() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothTransferReply_IsRunning(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QBluetoothTransferReply) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QBluetoothTransferReply___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothTransferReply) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QBluetoothTransferReply) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QBluetoothTransferReply___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QBluetoothTransferReply) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QBluetoothTransferReply___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothTransferReply) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QBluetoothTransferReply) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QBluetoothTransferReply___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QBluetoothTransferReply) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QBluetoothTransferReply___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothTransferReply) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QBluetoothTransferReply) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.QBluetoothTransferReply___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QBluetoothTransferReply) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QBluetoothTransferReply___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothTransferReply) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QBluetoothTransferReply) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QBluetoothTransferReply___findChildren_newList(ptr.Pointer()))
}

func (ptr *QBluetoothTransferReply) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QBluetoothTransferReply___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothTransferReply) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QBluetoothTransferReply) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QBluetoothTransferReply___children_newList(ptr.Pointer()))
}

//export callbackQBluetoothTransferReply_Event
func callbackQBluetoothTransferReply_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothTransferReplyFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QBluetoothTransferReply) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothTransferReply_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQBluetoothTransferReply_EventFilter
func callbackQBluetoothTransferReply_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothTransferReplyFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QBluetoothTransferReply) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothTransferReply_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQBluetoothTransferReply_ChildEvent
func callbackQBluetoothTransferReply_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQBluetoothTransferReplyFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QBluetoothTransferReply) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQBluetoothTransferReply_ConnectNotify
func callbackQBluetoothTransferReply_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBluetoothTransferReplyFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBluetoothTransferReply) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBluetoothTransferReply_CustomEvent
func callbackQBluetoothTransferReply_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQBluetoothTransferReplyFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QBluetoothTransferReply) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQBluetoothTransferReply_DeleteLater
func callbackQBluetoothTransferReply_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQBluetoothTransferReplyFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QBluetoothTransferReply) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQBluetoothTransferReply_Destroyed
func callbackQBluetoothTransferReply_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQBluetoothTransferReply_DisconnectNotify
func callbackQBluetoothTransferReply_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBluetoothTransferReplyFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBluetoothTransferReply) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBluetoothTransferReply_ObjectNameChanged
func callbackQBluetoothTransferReply_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtBluetooth_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQBluetoothTransferReply_TimerEvent
func callbackQBluetoothTransferReply_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQBluetoothTransferReplyFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QBluetoothTransferReply) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQBluetoothTransferReply_MetaObject
func callbackQBluetoothTransferReply_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQBluetoothTransferReplyFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QBluetoothTransferReply) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QBluetoothTransferReply_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QBluetoothTransferRequest struct {
	ptr unsafe.Pointer
}

type QBluetoothTransferRequest_ITF interface {
	QBluetoothTransferRequest_PTR() *QBluetoothTransferRequest
}

func (ptr *QBluetoothTransferRequest) QBluetoothTransferRequest_PTR() *QBluetoothTransferRequest {
	return ptr
}

func (ptr *QBluetoothTransferRequest) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QBluetoothTransferRequest) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
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

//go:generate stringer -type=QBluetoothTransferRequest__Attribute
//QBluetoothTransferRequest::Attribute
type QBluetoothTransferRequest__Attribute int64

const (
	QBluetoothTransferRequest__DescriptionAttribute QBluetoothTransferRequest__Attribute = QBluetoothTransferRequest__Attribute(0)
	QBluetoothTransferRequest__TimeAttribute        QBluetoothTransferRequest__Attribute = QBluetoothTransferRequest__Attribute(1)
	QBluetoothTransferRequest__TypeAttribute        QBluetoothTransferRequest__Attribute = QBluetoothTransferRequest__Attribute(2)
	QBluetoothTransferRequest__LengthAttribute      QBluetoothTransferRequest__Attribute = QBluetoothTransferRequest__Attribute(3)
	QBluetoothTransferRequest__NameAttribute        QBluetoothTransferRequest__Attribute = QBluetoothTransferRequest__Attribute(4)
)

func NewQBluetoothTransferRequest(address QBluetoothAddress_ITF) *QBluetoothTransferRequest {
	var tmpValue = NewQBluetoothTransferRequestFromPointer(C.QBluetoothTransferRequest_NewQBluetoothTransferRequest(PointerFromQBluetoothAddress(address)))
	runtime.SetFinalizer(tmpValue, (*QBluetoothTransferRequest).DestroyQBluetoothTransferRequest)
	return tmpValue
}

func NewQBluetoothTransferRequest2(other QBluetoothTransferRequest_ITF) *QBluetoothTransferRequest {
	var tmpValue = NewQBluetoothTransferRequestFromPointer(C.QBluetoothTransferRequest_NewQBluetoothTransferRequest2(PointerFromQBluetoothTransferRequest(other)))
	runtime.SetFinalizer(tmpValue, (*QBluetoothTransferRequest).DestroyQBluetoothTransferRequest)
	return tmpValue
}

func (ptr *QBluetoothTransferRequest) SetAttribute(code QBluetoothTransferRequest__Attribute, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferRequest_SetAttribute(ptr.Pointer(), C.longlong(code), core.PointerFromQVariant(value))
	}
}

func (ptr *QBluetoothTransferRequest) DestroyQBluetoothTransferRequest() {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferRequest_DestroyQBluetoothTransferRequest(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QBluetoothTransferRequest) Address() *QBluetoothAddress {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothAddressFromPointer(C.QBluetoothTransferRequest_Address(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothTransferRequest) Attribute(code QBluetoothTransferRequest__Attribute, defaultValue core.QVariant_ITF) *core.QVariant {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQVariantFromPointer(C.QBluetoothTransferRequest_Attribute(ptr.Pointer(), C.longlong(code), core.PointerFromQVariant(defaultValue)))
		runtime.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

type QBluetoothUuid struct {
	core.QUuid
}

type QBluetoothUuid_ITF interface {
	core.QUuid_ITF
	QBluetoothUuid_PTR() *QBluetoothUuid
}

func (ptr *QBluetoothUuid) QBluetoothUuid_PTR() *QBluetoothUuid {
	return ptr
}

func (ptr *QBluetoothUuid) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QUuid_PTR().Pointer()
	}
	return nil
}

func (ptr *QBluetoothUuid) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QUuid_PTR().SetPointer(p)
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

//go:generate stringer -type=QBluetoothUuid__CharacteristicType
//QBluetoothUuid::CharacteristicType
type QBluetoothUuid__CharacteristicType int64

const (
	QBluetoothUuid__DeviceName                                    QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a00)
	QBluetoothUuid__Appearance                                    QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a01)
	QBluetoothUuid__PeripheralPrivacyFlag                         QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a02)
	QBluetoothUuid__ReconnectionAddress                           QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a03)
	QBluetoothUuid__PeripheralPreferredConnectionParameters       QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a04)
	QBluetoothUuid__ServiceChanged                                QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a05)
	QBluetoothUuid__AlertLevel                                    QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a06)
	QBluetoothUuid__TxPowerLevel                                  QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a07)
	QBluetoothUuid__DateTime                                      QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a08)
	QBluetoothUuid__DayOfWeek                                     QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a09)
	QBluetoothUuid__DayDateTime                                   QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a0a)
	QBluetoothUuid__ExactTime256                                  QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a0c)
	QBluetoothUuid__DSTOffset                                     QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a0d)
	QBluetoothUuid__TimeZone                                      QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a0e)
	QBluetoothUuid__LocalTimeInformation                          QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a0f)
	QBluetoothUuid__TimeWithDST                                   QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a11)
	QBluetoothUuid__TimeAccuracy                                  QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a12)
	QBluetoothUuid__TimeSource                                    QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a13)
	QBluetoothUuid__ReferenceTimeInformation                      QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a14)
	QBluetoothUuid__TimeUpdateControlPoint                        QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a16)
	QBluetoothUuid__TimeUpdateState                               QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a17)
	QBluetoothUuid__GlucoseMeasurement                            QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a18)
	QBluetoothUuid__BatteryLevel                                  QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a19)
	QBluetoothUuid__TemperatureMeasurement                        QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a1c)
	QBluetoothUuid__TemperatureType                               QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a1d)
	QBluetoothUuid__IntermediateTemperature                       QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a1e)
	QBluetoothUuid__MeasurementInterval                           QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a21)
	QBluetoothUuid__BootKeyboardInputReport                       QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a22)
	QBluetoothUuid__SystemID                                      QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a23)
	QBluetoothUuid__ModelNumberString                             QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a24)
	QBluetoothUuid__SerialNumberString                            QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a25)
	QBluetoothUuid__FirmwareRevisionString                        QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a26)
	QBluetoothUuid__HardwareRevisionString                        QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a27)
	QBluetoothUuid__SoftwareRevisionString                        QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a28)
	QBluetoothUuid__ManufacturerNameString                        QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a29)
	QBluetoothUuid__IEEE1107320601RegulatoryCertificationDataList QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a2a)
	QBluetoothUuid__CurrentTime                                   QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a2b)
	QBluetoothUuid__MagneticDeclination                           QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a2c)
	QBluetoothUuid__ScanRefresh                                   QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a31)
	QBluetoothUuid__BootKeyboardOutputReport                      QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a32)
	QBluetoothUuid__BootMouseInputReport                          QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a33)
	QBluetoothUuid__GlucoseMeasurementContext                     QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a34)
	QBluetoothUuid__BloodPressureMeasurement                      QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a35)
	QBluetoothUuid__IntermediateCuffPressure                      QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a36)
	QBluetoothUuid__HeartRateMeasurement                          QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a37)
	QBluetoothUuid__BodySensorLocation                            QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a38)
	QBluetoothUuid__HeartRateControlPoint                         QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a39)
	QBluetoothUuid__AlertStatus                                   QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a3f)
	QBluetoothUuid__RingerControlPoint                            QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a40)
	QBluetoothUuid__RingerSetting                                 QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a41)
	QBluetoothUuid__AlertCategoryIDBitMask                        QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a42)
	QBluetoothUuid__AlertCategoryID                               QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a43)
	QBluetoothUuid__AlertNotificationControlPoint                 QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a44)
	QBluetoothUuid__UnreadAlertStatus                             QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a45)
	QBluetoothUuid__NewAlert                                      QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a46)
	QBluetoothUuid__SupportedNewAlertCategory                     QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a47)
	QBluetoothUuid__SupportedUnreadAlertCategory                  QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a48)
	QBluetoothUuid__BloodPressureFeature                          QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a49)
	QBluetoothUuid__HIDInformation                                QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a4a)
	QBluetoothUuid__ReportMap                                     QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a4b)
	QBluetoothUuid__HIDControlPoint                               QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a4c)
	QBluetoothUuid__Report                                        QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a4d)
	QBluetoothUuid__ProtocolMode                                  QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a4e)
	QBluetoothUuid__ScanIntervalWindow                            QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a4f)
	QBluetoothUuid__PnPID                                         QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a50)
	QBluetoothUuid__GlucoseFeature                                QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a51)
	QBluetoothUuid__RecordAccessControlPoint                      QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a52)
	QBluetoothUuid__RSCMeasurement                                QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a53)
	QBluetoothUuid__RSCFeature                                    QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a54)
	QBluetoothUuid__SCControlPoint                                QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a55)
	QBluetoothUuid__CSCMeasurement                                QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a5b)
	QBluetoothUuid__CSCFeature                                    QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a5c)
	QBluetoothUuid__SensorLocation                                QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a5d)
	QBluetoothUuid__CyclingPowerMeasurement                       QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a63)
	QBluetoothUuid__CyclingPowerVector                            QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a64)
	QBluetoothUuid__CyclingPowerFeature                           QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a65)
	QBluetoothUuid__CyclingPowerControlPoint                      QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a66)
	QBluetoothUuid__LocationAndSpeed                              QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a67)
	QBluetoothUuid__Navigation                                    QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a68)
	QBluetoothUuid__PositionQuality                               QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a69)
	QBluetoothUuid__LNFeature                                     QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a6a)
	QBluetoothUuid__LNControlPoint                                QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a6b)
	QBluetoothUuid__Elevation                                     QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a6c)
	QBluetoothUuid__Pressure                                      QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a6d)
	QBluetoothUuid__Temperature                                   QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a6e)
	QBluetoothUuid__Humidity                                      QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a6f)
	QBluetoothUuid__TrueWindSpeed                                 QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a70)
	QBluetoothUuid__TrueWindDirection                             QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a71)
	QBluetoothUuid__ApparentWindSpeed                             QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a72)
	QBluetoothUuid__ApparentWindDirection                         QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a73)
	QBluetoothUuid__GustFactor                                    QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a74)
	QBluetoothUuid__PollenConcentration                           QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a75)
	QBluetoothUuid__UVIndex                                       QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a76)
	QBluetoothUuid__Irradiance                                    QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a77)
	QBluetoothUuid__Rainfall                                      QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a78)
	QBluetoothUuid__WindChill                                     QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a79)
	QBluetoothUuid__HeatIndex                                     QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a7a)
	QBluetoothUuid__DewPoint                                      QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a7b)
	QBluetoothUuid__DescriptorValueChanged                        QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a7d)
	QBluetoothUuid__AerobicHeartRateLowerLimit                    QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a7e)
	QBluetoothUuid__AerobicThreshold                              QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a7f)
	QBluetoothUuid__Age                                           QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a80)
	QBluetoothUuid__AnaerobicHeartRateLowerLimit                  QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a81)
	QBluetoothUuid__AnaerobicHeartRateUpperLimit                  QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a82)
	QBluetoothUuid__AnaerobicThreshold                            QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a83)
	QBluetoothUuid__AerobicHeartRateUpperLimit                    QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a84)
	QBluetoothUuid__DateOfBirth                                   QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a85)
	QBluetoothUuid__DateOfThresholdAssessment                     QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a86)
	QBluetoothUuid__EmailAddress                                  QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a87)
	QBluetoothUuid__FatBurnHeartRateLowerLimit                    QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a88)
	QBluetoothUuid__FatBurnHeartRateUpperLimit                    QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a89)
	QBluetoothUuid__FirstName                                     QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a8a)
	QBluetoothUuid__FiveZoneHeartRateLimits                       QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a8b)
	QBluetoothUuid__Gender                                        QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a8c)
	QBluetoothUuid__HeartRateMax                                  QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a8d)
	QBluetoothUuid__Height                                        QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a8e)
	QBluetoothUuid__HipCircumference                              QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a8f)
	QBluetoothUuid__LastName                                      QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a90)
	QBluetoothUuid__MaximumRecommendedHeartRate                   QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a91)
	QBluetoothUuid__RestingHeartRate                              QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a92)
	QBluetoothUuid__SportTypeForAerobicAnaerobicThresholds        QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a93)
	QBluetoothUuid__ThreeZoneHeartRateLimits                      QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a94)
	QBluetoothUuid__TwoZoneHeartRateLimits                        QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a95)
	QBluetoothUuid__VO2Max                                        QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a96)
	QBluetoothUuid__WaistCircumference                            QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a97)
	QBluetoothUuid__Weight                                        QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a98)
	QBluetoothUuid__DatabaseChangeIncrement                       QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a99)
	QBluetoothUuid__UserIndex                                     QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a9a)
	QBluetoothUuid__BodyCompositionFeature                        QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a9b)
	QBluetoothUuid__BodyCompositionMeasurement                    QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a9c)
	QBluetoothUuid__WeightMeasurement                             QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a9d)
	QBluetoothUuid__WeightScaleFeature                            QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a9e)
	QBluetoothUuid__UserControlPoint                              QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2a9f)
	QBluetoothUuid__MagneticFluxDensity2D                         QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2aa0)
	QBluetoothUuid__MagneticFluxDensity3D                         QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2aa1)
	QBluetoothUuid__Language                                      QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2aa2)
	QBluetoothUuid__BarometricPressureTrend                       QBluetoothUuid__CharacteristicType = QBluetoothUuid__CharacteristicType(0x2aa3)
)

//go:generate stringer -type=QBluetoothUuid__DescriptorType
//QBluetoothUuid::DescriptorType
type QBluetoothUuid__DescriptorType int64

const (
	QBluetoothUuid__UnknownDescriptorType              QBluetoothUuid__DescriptorType = QBluetoothUuid__DescriptorType(0x0)
	QBluetoothUuid__CharacteristicExtendedProperties   QBluetoothUuid__DescriptorType = QBluetoothUuid__DescriptorType(0x2900)
	QBluetoothUuid__CharacteristicUserDescription      QBluetoothUuid__DescriptorType = QBluetoothUuid__DescriptorType(0x2901)
	QBluetoothUuid__ClientCharacteristicConfiguration  QBluetoothUuid__DescriptorType = QBluetoothUuid__DescriptorType(0x2902)
	QBluetoothUuid__ServerCharacteristicConfiguration  QBluetoothUuid__DescriptorType = QBluetoothUuid__DescriptorType(0x2903)
	QBluetoothUuid__CharacteristicPresentationFormat   QBluetoothUuid__DescriptorType = QBluetoothUuid__DescriptorType(0x2904)
	QBluetoothUuid__CharacteristicAggregateFormat      QBluetoothUuid__DescriptorType = QBluetoothUuid__DescriptorType(0x2905)
	QBluetoothUuid__ValidRange                         QBluetoothUuid__DescriptorType = QBluetoothUuid__DescriptorType(0x2906)
	QBluetoothUuid__ExternalReportReference            QBluetoothUuid__DescriptorType = QBluetoothUuid__DescriptorType(0x2907)
	QBluetoothUuid__ReportReference                    QBluetoothUuid__DescriptorType = QBluetoothUuid__DescriptorType(0x2908)
	QBluetoothUuid__EnvironmentalSensingConfiguration  QBluetoothUuid__DescriptorType = QBluetoothUuid__DescriptorType(0x290b)
	QBluetoothUuid__EnvironmentalSensingMeasurement    QBluetoothUuid__DescriptorType = QBluetoothUuid__DescriptorType(0x290c)
	QBluetoothUuid__EnvironmentalSensingTriggerSetting QBluetoothUuid__DescriptorType = QBluetoothUuid__DescriptorType(0x290d)
)

//go:generate stringer -type=QBluetoothUuid__ProtocolUuid
//QBluetoothUuid::ProtocolUuid
type QBluetoothUuid__ProtocolUuid int64

const (
	QBluetoothUuid__Sdp                    QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x0001)
	QBluetoothUuid__Udp                    QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x0002)
	QBluetoothUuid__Rfcomm                 QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x0003)
	QBluetoothUuid__Tcp                    QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x0004)
	QBluetoothUuid__TcsBin                 QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x0005)
	QBluetoothUuid__TcsAt                  QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x0006)
	QBluetoothUuid__Att                    QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x0007)
	QBluetoothUuid__Obex                   QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x0008)
	QBluetoothUuid__Ip                     QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x0009)
	QBluetoothUuid__Ftp                    QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x000A)
	QBluetoothUuid__Http                   QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x000C)
	QBluetoothUuid__Wsp                    QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x000E)
	QBluetoothUuid__Bnep                   QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x000F)
	QBluetoothUuid__Upnp                   QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x0010)
	QBluetoothUuid__Hidp                   QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x0011)
	QBluetoothUuid__HardcopyControlChannel QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x0012)
	QBluetoothUuid__HardcopyDataChannel    QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x0014)
	QBluetoothUuid__HardcopyNotification   QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x0016)
	QBluetoothUuid__Avctp                  QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x0017)
	QBluetoothUuid__Avdtp                  QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x0019)
	QBluetoothUuid__Cmtp                   QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x001B)
	QBluetoothUuid__UdiCPlain              QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x001D)
	QBluetoothUuid__McapControlChannel     QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x001E)
	QBluetoothUuid__McapDataChannel        QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x001F)
	QBluetoothUuid__L2cap                  QBluetoothUuid__ProtocolUuid = QBluetoothUuid__ProtocolUuid(0x0100)
)

//go:generate stringer -type=QBluetoothUuid__ServiceClassUuid
//QBluetoothUuid::ServiceClassUuid
type QBluetoothUuid__ServiceClassUuid int64

const (
	QBluetoothUuid__ServiceDiscoveryServer                QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1000)
	QBluetoothUuid__BrowseGroupDescriptor                 QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1001)
	QBluetoothUuid__PublicBrowseGroup                     QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1002)
	QBluetoothUuid__SerialPort                            QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1101)
	QBluetoothUuid__LANAccessUsingPPP                     QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1102)
	QBluetoothUuid__DialupNetworking                      QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1103)
	QBluetoothUuid__IrMCSync                              QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1104)
	QBluetoothUuid__ObexObjectPush                        QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1105)
	QBluetoothUuid__OBEXFileTransfer                      QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1106)
	QBluetoothUuid__IrMCSyncCommand                       QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1107)
	QBluetoothUuid__Headset                               QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1108)
	QBluetoothUuid__AudioSource                           QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x110a)
	QBluetoothUuid__AudioSink                             QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x110b)
	QBluetoothUuid__AV_RemoteControlTarget                QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x110c)
	QBluetoothUuid__AdvancedAudioDistribution             QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x110d)
	QBluetoothUuid__AV_RemoteControl                      QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x110e)
	QBluetoothUuid__AV_RemoteControlController            QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x110f)
	QBluetoothUuid__HeadsetAG                             QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1112)
	QBluetoothUuid__PANU                                  QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1115)
	QBluetoothUuid__NAP                                   QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1116)
	QBluetoothUuid__GN                                    QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1117)
	QBluetoothUuid__DirectPrinting                        QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1118)
	QBluetoothUuid__ReferencePrinting                     QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1119)
	QBluetoothUuid__BasicImage                            QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x111a)
	QBluetoothUuid__ImagingResponder                      QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x111b)
	QBluetoothUuid__ImagingAutomaticArchive               QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x111c)
	QBluetoothUuid__ImagingReferenceObjects               QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x111d)
	QBluetoothUuid__Handsfree                             QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x111e)
	QBluetoothUuid__HandsfreeAudioGateway                 QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x111f)
	QBluetoothUuid__DirectPrintingReferenceObjectsService QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1120)
	QBluetoothUuid__ReflectedUI                           QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1121)
	QBluetoothUuid__BasicPrinting                         QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1122)
	QBluetoothUuid__PrintingStatus                        QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1123)
	QBluetoothUuid__HumanInterfaceDeviceService           QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1124)
	QBluetoothUuid__HardcopyCableReplacement              QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1125)
	QBluetoothUuid__HCRPrint                              QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1126)
	QBluetoothUuid__HCRScan                               QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1127)
	QBluetoothUuid__SIMAccess                             QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x112d)
	QBluetoothUuid__PhonebookAccessPCE                    QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x112e)
	QBluetoothUuid__PhonebookAccessPSE                    QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x112f)
	QBluetoothUuid__PhonebookAccess                       QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1130)
	QBluetoothUuid__HeadsetHS                             QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1131)
	QBluetoothUuid__MessageAccessServer                   QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1132)
	QBluetoothUuid__MessageNotificationServer             QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1133)
	QBluetoothUuid__MessageAccessProfile                  QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1134)
	QBluetoothUuid__GNSS                                  QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1135)
	QBluetoothUuid__GNSSServer                            QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1136)
	QBluetoothUuid__Display3D                             QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1137)
	QBluetoothUuid__Glasses3D                             QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1138)
	QBluetoothUuid__Synchronization3D                     QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1139)
	QBluetoothUuid__MPSProfile                            QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x113a)
	QBluetoothUuid__MPSService                            QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x113b)
	QBluetoothUuid__PnPInformation                        QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1200)
	QBluetoothUuid__GenericNetworking                     QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1201)
	QBluetoothUuid__GenericFileTransfer                   QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1202)
	QBluetoothUuid__GenericAudio                          QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1203)
	QBluetoothUuid__GenericTelephony                      QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1204)
	QBluetoothUuid__VideoSource                           QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1303)
	QBluetoothUuid__VideoSink                             QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1304)
	QBluetoothUuid__VideoDistribution                     QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1305)
	QBluetoothUuid__HDP                                   QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1400)
	QBluetoothUuid__HDPSource                             QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1401)
	QBluetoothUuid__HDPSink                               QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1402)
	QBluetoothUuid__GenericAccess                         QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1800)
	QBluetoothUuid__GenericAttribute                      QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1801)
	QBluetoothUuid__ImmediateAlert                        QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1802)
	QBluetoothUuid__LinkLoss                              QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1803)
	QBluetoothUuid__TxPower                               QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1804)
	QBluetoothUuid__CurrentTimeService                    QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1805)
	QBluetoothUuid__ReferenceTimeUpdateService            QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1806)
	QBluetoothUuid__NextDSTChangeService                  QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1807)
	QBluetoothUuid__Glucose                               QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1808)
	QBluetoothUuid__HealthThermometer                     QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1809)
	QBluetoothUuid__DeviceInformation                     QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x180a)
	QBluetoothUuid__HeartRate                             QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x180d)
	QBluetoothUuid__PhoneAlertStatusService               QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x180e)
	QBluetoothUuid__BatteryService                        QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x180f)
	QBluetoothUuid__BloodPressure                         QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1810)
	QBluetoothUuid__AlertNotificationService              QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1811)
	QBluetoothUuid__HumanInterfaceDevice                  QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1812)
	QBluetoothUuid__ScanParameters                        QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1813)
	QBluetoothUuid__RunningSpeedAndCadence                QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1814)
	QBluetoothUuid__CyclingSpeedAndCadence                QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1816)
	QBluetoothUuid__CyclingPower                          QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1818)
	QBluetoothUuid__LocationAndNavigation                 QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x1819)
	QBluetoothUuid__EnvironmentalSensing                  QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x181a)
	QBluetoothUuid__BodyComposition                       QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x181b)
	QBluetoothUuid__UserData                              QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x181c)
	QBluetoothUuid__WeightScale                           QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x181d)
	QBluetoothUuid__BondManagement                        QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x181e)
	QBluetoothUuid__ContinuousGlucoseMonitoring           QBluetoothUuid__ServiceClassUuid = QBluetoothUuid__ServiceClassUuid(0x181f)
)

func NewQBluetoothUuid() *QBluetoothUuid {
	var tmpValue = NewQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid())
	runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
	return tmpValue
}

func NewQBluetoothUuid4(uuid QBluetoothUuid__CharacteristicType) *QBluetoothUuid {
	var tmpValue = NewQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid4(C.longlong(uuid)))
	runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
	return tmpValue
}

func NewQBluetoothUuid5(uuid QBluetoothUuid__DescriptorType) *QBluetoothUuid {
	var tmpValue = NewQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid5(C.longlong(uuid)))
	runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
	return tmpValue
}

func NewQBluetoothUuid2(uuid QBluetoothUuid__ProtocolUuid) *QBluetoothUuid {
	var tmpValue = NewQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid2(C.longlong(uuid)))
	runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
	return tmpValue
}

func NewQBluetoothUuid3(uuid QBluetoothUuid__ServiceClassUuid) *QBluetoothUuid {
	var tmpValue = NewQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid3(C.longlong(uuid)))
	runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
	return tmpValue
}

func NewQBluetoothUuid10(uuid QBluetoothUuid_ITF) *QBluetoothUuid {
	var tmpValue = NewQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid10(PointerFromQBluetoothUuid(uuid)))
	runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
	return tmpValue
}

func NewQBluetoothUuid9(uuid string) *QBluetoothUuid {
	var uuidC *C.char
	if uuid != "" {
		uuidC = C.CString(uuid)
		defer C.free(unsafe.Pointer(uuidC))
	}
	var tmpValue = NewQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid9(C.struct_QtBluetooth_PackedString{data: uuidC, len: C.longlong(len(uuid))}))
	runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
	return tmpValue
}

func NewQBluetoothUuid11(uuid core.QUuid_ITF) *QBluetoothUuid {
	var tmpValue = NewQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid11(core.PointerFromQUuid(uuid)))
	runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
	return tmpValue
}

func NewQBluetoothUuid6(uuid uint16) *QBluetoothUuid {
	var tmpValue = NewQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid6(C.ushort(uuid)))
	runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
	return tmpValue
}

func NewQBluetoothUuid7(uuid uint) *QBluetoothUuid {
	var tmpValue = NewQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid7(C.uint(uint32(uuid))))
	runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
	return tmpValue
}

func QBluetoothUuid_CharacteristicToString(uuid QBluetoothUuid__CharacteristicType) string {
	return cGoUnpackString(C.QBluetoothUuid_QBluetoothUuid_CharacteristicToString(C.longlong(uuid)))
}

func (ptr *QBluetoothUuid) CharacteristicToString(uuid QBluetoothUuid__CharacteristicType) string {
	return cGoUnpackString(C.QBluetoothUuid_QBluetoothUuid_CharacteristicToString(C.longlong(uuid)))
}

func QBluetoothUuid_DescriptorToString(uuid QBluetoothUuid__DescriptorType) string {
	return cGoUnpackString(C.QBluetoothUuid_QBluetoothUuid_DescriptorToString(C.longlong(uuid)))
}

func (ptr *QBluetoothUuid) DescriptorToString(uuid QBluetoothUuid__DescriptorType) string {
	return cGoUnpackString(C.QBluetoothUuid_QBluetoothUuid_DescriptorToString(C.longlong(uuid)))
}

func QBluetoothUuid_ProtocolToString(uuid QBluetoothUuid__ProtocolUuid) string {
	return cGoUnpackString(C.QBluetoothUuid_QBluetoothUuid_ProtocolToString(C.longlong(uuid)))
}

func (ptr *QBluetoothUuid) ProtocolToString(uuid QBluetoothUuid__ProtocolUuid) string {
	return cGoUnpackString(C.QBluetoothUuid_QBluetoothUuid_ProtocolToString(C.longlong(uuid)))
}

func QBluetoothUuid_ServiceClassToString(uuid QBluetoothUuid__ServiceClassUuid) string {
	return cGoUnpackString(C.QBluetoothUuid_QBluetoothUuid_ServiceClassToString(C.longlong(uuid)))
}

func (ptr *QBluetoothUuid) ServiceClassToString(uuid QBluetoothUuid__ServiceClassUuid) string {
	return cGoUnpackString(C.QBluetoothUuid_QBluetoothUuid_ServiceClassToString(C.longlong(uuid)))
}

func (ptr *QBluetoothUuid) DestroyQBluetoothUuid() {
	if ptr.Pointer() != nil {
		C.QBluetoothUuid_DestroyQBluetoothUuid(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QBluetoothUuid) MinimumSize() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QBluetoothUuid_MinimumSize(ptr.Pointer())))
	}
	return 0
}

func (ptr *QBluetoothUuid) ToUInt16(ok bool) uint16 {
	if ptr.Pointer() != nil {
		return uint16(C.QBluetoothUuid_ToUInt16(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(ok)))))
	}
	return 0
}

func (ptr *QBluetoothUuid) ToUInt32(ok bool) uint {
	if ptr.Pointer() != nil {
		return uint(uint32(C.QBluetoothUuid_ToUInt32(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(ok))))))
	}
	return 0
}

type QLowEnergyAdvertisingData struct {
	ptr unsafe.Pointer
}

type QLowEnergyAdvertisingData_ITF interface {
	QLowEnergyAdvertisingData_PTR() *QLowEnergyAdvertisingData
}

func (ptr *QLowEnergyAdvertisingData) QLowEnergyAdvertisingData_PTR() *QLowEnergyAdvertisingData {
	return ptr
}

func (ptr *QLowEnergyAdvertisingData) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QLowEnergyAdvertisingData) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
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

//go:generate stringer -type=QLowEnergyAdvertisingData__Discoverability
//QLowEnergyAdvertisingData::Discoverability
type QLowEnergyAdvertisingData__Discoverability int64

const (
	QLowEnergyAdvertisingData__DiscoverabilityNone    QLowEnergyAdvertisingData__Discoverability = QLowEnergyAdvertisingData__Discoverability(0)
	QLowEnergyAdvertisingData__DiscoverabilityLimited QLowEnergyAdvertisingData__Discoverability = QLowEnergyAdvertisingData__Discoverability(1)
	QLowEnergyAdvertisingData__DiscoverabilityGeneral QLowEnergyAdvertisingData__Discoverability = QLowEnergyAdvertisingData__Discoverability(2)
)

func NewQLowEnergyAdvertisingData() *QLowEnergyAdvertisingData {
	var tmpValue = NewQLowEnergyAdvertisingDataFromPointer(C.QLowEnergyAdvertisingData_NewQLowEnergyAdvertisingData())
	runtime.SetFinalizer(tmpValue, (*QLowEnergyAdvertisingData).DestroyQLowEnergyAdvertisingData)
	return tmpValue
}

func NewQLowEnergyAdvertisingData2(other QLowEnergyAdvertisingData_ITF) *QLowEnergyAdvertisingData {
	var tmpValue = NewQLowEnergyAdvertisingDataFromPointer(C.QLowEnergyAdvertisingData_NewQLowEnergyAdvertisingData2(PointerFromQLowEnergyAdvertisingData(other)))
	runtime.SetFinalizer(tmpValue, (*QLowEnergyAdvertisingData).DestroyQLowEnergyAdvertisingData)
	return tmpValue
}

func QLowEnergyAdvertisingData_InvalidManufacturerId() uint16 {
	return uint16(C.QLowEnergyAdvertisingData_QLowEnergyAdvertisingData_InvalidManufacturerId())
}

func (ptr *QLowEnergyAdvertisingData) InvalidManufacturerId() uint16 {
	return uint16(C.QLowEnergyAdvertisingData_QLowEnergyAdvertisingData_InvalidManufacturerId())
}

func (ptr *QLowEnergyAdvertisingData) SetDiscoverability(mode QLowEnergyAdvertisingData__Discoverability) {
	if ptr.Pointer() != nil {
		C.QLowEnergyAdvertisingData_SetDiscoverability(ptr.Pointer(), C.longlong(mode))
	}
}

func (ptr *QLowEnergyAdvertisingData) SetIncludePowerLevel(doInclude bool) {
	if ptr.Pointer() != nil {
		C.QLowEnergyAdvertisingData_SetIncludePowerLevel(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(doInclude))))
	}
}

func (ptr *QLowEnergyAdvertisingData) SetLocalName(name string) {
	if ptr.Pointer() != nil {
		var nameC *C.char
		if name != "" {
			nameC = C.CString(name)
			defer C.free(unsafe.Pointer(nameC))
		}
		C.QLowEnergyAdvertisingData_SetLocalName(ptr.Pointer(), C.struct_QtBluetooth_PackedString{data: nameC, len: C.longlong(len(name))})
	}
}

func (ptr *QLowEnergyAdvertisingData) SetManufacturerData(id uint16, data core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyAdvertisingData_SetManufacturerData(ptr.Pointer(), C.ushort(id), core.PointerFromQByteArray(data))
	}
}

func (ptr *QLowEnergyAdvertisingData) SetRawData(data core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyAdvertisingData_SetRawData(ptr.Pointer(), core.PointerFromQByteArray(data))
	}
}

func (ptr *QLowEnergyAdvertisingData) SetServices(services []*QBluetoothUuid) {
	if ptr.Pointer() != nil {
		C.QLowEnergyAdvertisingData_SetServices(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewQLowEnergyAdvertisingDataFromPointer(NewQLowEnergyAdvertisingDataFromPointer(nil).__setServices_services_newList())
			for _, v := range services {
				tmpList.__setServices_services_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QLowEnergyAdvertisingData) Swap(other QLowEnergyAdvertisingData_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyAdvertisingData_Swap(ptr.Pointer(), PointerFromQLowEnergyAdvertisingData(other))
	}
}

func (ptr *QLowEnergyAdvertisingData) DestroyQLowEnergyAdvertisingData() {
	if ptr.Pointer() != nil {
		C.QLowEnergyAdvertisingData_DestroyQLowEnergyAdvertisingData(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QLowEnergyAdvertisingData) Discoverability() QLowEnergyAdvertisingData__Discoverability {
	if ptr.Pointer() != nil {
		return QLowEnergyAdvertisingData__Discoverability(C.QLowEnergyAdvertisingData_Discoverability(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyAdvertisingData) ManufacturerData() *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QLowEnergyAdvertisingData_ManufacturerData(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyAdvertisingData) RawData() *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QLowEnergyAdvertisingData_RawData(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyAdvertisingData) Services() []*QBluetoothUuid {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtBluetooth_PackedList) []*QBluetoothUuid {
			var out = make([]*QBluetoothUuid, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQLowEnergyAdvertisingDataFromPointer(l.data).__services_atList(i)
			}
			return out
		}(C.QLowEnergyAdvertisingData_Services(ptr.Pointer()))
	}
	return make([]*QBluetoothUuid, 0)
}

func (ptr *QLowEnergyAdvertisingData) LocalName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QLowEnergyAdvertisingData_LocalName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLowEnergyAdvertisingData) IncludePowerLevel() bool {
	if ptr.Pointer() != nil {
		return C.QLowEnergyAdvertisingData_IncludePowerLevel(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLowEnergyAdvertisingData) ManufacturerId() uint16 {
	if ptr.Pointer() != nil {
		return uint16(C.QLowEnergyAdvertisingData_ManufacturerId(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyAdvertisingData) __setServices_services_atList(i int) *QBluetoothUuid {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothUuidFromPointer(C.QLowEnergyAdvertisingData___setServices_services_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyAdvertisingData) __setServices_services_setList(i QBluetoothUuid_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyAdvertisingData___setServices_services_setList(ptr.Pointer(), PointerFromQBluetoothUuid(i))
	}
}

func (ptr *QLowEnergyAdvertisingData) __setServices_services_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QLowEnergyAdvertisingData___setServices_services_newList(ptr.Pointer()))
}

func (ptr *QLowEnergyAdvertisingData) __services_atList(i int) *QBluetoothUuid {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothUuidFromPointer(C.QLowEnergyAdvertisingData___services_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyAdvertisingData) __services_setList(i QBluetoothUuid_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyAdvertisingData___services_setList(ptr.Pointer(), PointerFromQBluetoothUuid(i))
	}
}

func (ptr *QLowEnergyAdvertisingData) __services_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QLowEnergyAdvertisingData___services_newList(ptr.Pointer()))
}

type QLowEnergyAdvertisingParameters struct {
	ptr unsafe.Pointer
}

type QLowEnergyAdvertisingParameters_ITF interface {
	QLowEnergyAdvertisingParameters_PTR() *QLowEnergyAdvertisingParameters
}

func (ptr *QLowEnergyAdvertisingParameters) QLowEnergyAdvertisingParameters_PTR() *QLowEnergyAdvertisingParameters {
	return ptr
}

func (ptr *QLowEnergyAdvertisingParameters) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QLowEnergyAdvertisingParameters) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
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

//go:generate stringer -type=QLowEnergyAdvertisingParameters__FilterPolicy
//QLowEnergyAdvertisingParameters::FilterPolicy
type QLowEnergyAdvertisingParameters__FilterPolicy int64

const (
	QLowEnergyAdvertisingParameters__IgnoreWhiteList                      QLowEnergyAdvertisingParameters__FilterPolicy = QLowEnergyAdvertisingParameters__FilterPolicy(0x00)
	QLowEnergyAdvertisingParameters__UseWhiteListForScanning              QLowEnergyAdvertisingParameters__FilterPolicy = QLowEnergyAdvertisingParameters__FilterPolicy(0x01)
	QLowEnergyAdvertisingParameters__UseWhiteListForConnecting            QLowEnergyAdvertisingParameters__FilterPolicy = QLowEnergyAdvertisingParameters__FilterPolicy(0x02)
	QLowEnergyAdvertisingParameters__UseWhiteListForScanningAndConnecting QLowEnergyAdvertisingParameters__FilterPolicy = QLowEnergyAdvertisingParameters__FilterPolicy(0x03)
)

//go:generate stringer -type=QLowEnergyAdvertisingParameters__Mode
//QLowEnergyAdvertisingParameters::Mode
type QLowEnergyAdvertisingParameters__Mode int64

const (
	QLowEnergyAdvertisingParameters__AdvInd        QLowEnergyAdvertisingParameters__Mode = QLowEnergyAdvertisingParameters__Mode(0x0)
	QLowEnergyAdvertisingParameters__AdvScanInd    QLowEnergyAdvertisingParameters__Mode = QLowEnergyAdvertisingParameters__Mode(0x2)
	QLowEnergyAdvertisingParameters__AdvNonConnInd QLowEnergyAdvertisingParameters__Mode = QLowEnergyAdvertisingParameters__Mode(0x3)
)

func NewQLowEnergyAdvertisingParameters() *QLowEnergyAdvertisingParameters {
	var tmpValue = NewQLowEnergyAdvertisingParametersFromPointer(C.QLowEnergyAdvertisingParameters_NewQLowEnergyAdvertisingParameters())
	runtime.SetFinalizer(tmpValue, (*QLowEnergyAdvertisingParameters).DestroyQLowEnergyAdvertisingParameters)
	return tmpValue
}

func NewQLowEnergyAdvertisingParameters2(other QLowEnergyAdvertisingParameters_ITF) *QLowEnergyAdvertisingParameters {
	var tmpValue = NewQLowEnergyAdvertisingParametersFromPointer(C.QLowEnergyAdvertisingParameters_NewQLowEnergyAdvertisingParameters2(PointerFromQLowEnergyAdvertisingParameters(other)))
	runtime.SetFinalizer(tmpValue, (*QLowEnergyAdvertisingParameters).DestroyQLowEnergyAdvertisingParameters)
	return tmpValue
}

func (ptr *QLowEnergyAdvertisingParameters) SetInterval(minimum uint16, maximum uint16) {
	if ptr.Pointer() != nil {
		C.QLowEnergyAdvertisingParameters_SetInterval(ptr.Pointer(), C.ushort(minimum), C.ushort(maximum))
	}
}

func (ptr *QLowEnergyAdvertisingParameters) SetMode(mode QLowEnergyAdvertisingParameters__Mode) {
	if ptr.Pointer() != nil {
		C.QLowEnergyAdvertisingParameters_SetMode(ptr.Pointer(), C.longlong(mode))
	}
}

func (ptr *QLowEnergyAdvertisingParameters) Swap(other QLowEnergyAdvertisingParameters_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyAdvertisingParameters_Swap(ptr.Pointer(), PointerFromQLowEnergyAdvertisingParameters(other))
	}
}

func (ptr *QLowEnergyAdvertisingParameters) DestroyQLowEnergyAdvertisingParameters() {
	if ptr.Pointer() != nil {
		C.QLowEnergyAdvertisingParameters_DestroyQLowEnergyAdvertisingParameters(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QLowEnergyAdvertisingParameters) FilterPolicy() QLowEnergyAdvertisingParameters__FilterPolicy {
	if ptr.Pointer() != nil {
		return QLowEnergyAdvertisingParameters__FilterPolicy(C.QLowEnergyAdvertisingParameters_FilterPolicy(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyAdvertisingParameters) Mode() QLowEnergyAdvertisingParameters__Mode {
	if ptr.Pointer() != nil {
		return QLowEnergyAdvertisingParameters__Mode(C.QLowEnergyAdvertisingParameters_Mode(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyAdvertisingParameters) MaximumInterval() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QLowEnergyAdvertisingParameters_MaximumInterval(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLowEnergyAdvertisingParameters) MinimumInterval() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QLowEnergyAdvertisingParameters_MinimumInterval(ptr.Pointer())))
	}
	return 0
}

type QLowEnergyCharacteristic struct {
	ptr unsafe.Pointer
}

type QLowEnergyCharacteristic_ITF interface {
	QLowEnergyCharacteristic_PTR() *QLowEnergyCharacteristic
}

func (ptr *QLowEnergyCharacteristic) QLowEnergyCharacteristic_PTR() *QLowEnergyCharacteristic {
	return ptr
}

func (ptr *QLowEnergyCharacteristic) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QLowEnergyCharacteristic) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
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

//go:generate stringer -type=QLowEnergyCharacteristic__PropertyType
//QLowEnergyCharacteristic::PropertyType
type QLowEnergyCharacteristic__PropertyType int64

const (
	QLowEnergyCharacteristic__Unknown          QLowEnergyCharacteristic__PropertyType = QLowEnergyCharacteristic__PropertyType(0x00)
	QLowEnergyCharacteristic__Broadcasting     QLowEnergyCharacteristic__PropertyType = QLowEnergyCharacteristic__PropertyType(0x01)
	QLowEnergyCharacteristic__Read             QLowEnergyCharacteristic__PropertyType = QLowEnergyCharacteristic__PropertyType(0x02)
	QLowEnergyCharacteristic__WriteNoResponse  QLowEnergyCharacteristic__PropertyType = QLowEnergyCharacteristic__PropertyType(0x04)
	QLowEnergyCharacteristic__Write            QLowEnergyCharacteristic__PropertyType = QLowEnergyCharacteristic__PropertyType(0x08)
	QLowEnergyCharacteristic__Notify           QLowEnergyCharacteristic__PropertyType = QLowEnergyCharacteristic__PropertyType(0x10)
	QLowEnergyCharacteristic__Indicate         QLowEnergyCharacteristic__PropertyType = QLowEnergyCharacteristic__PropertyType(0x20)
	QLowEnergyCharacteristic__WriteSigned      QLowEnergyCharacteristic__PropertyType = QLowEnergyCharacteristic__PropertyType(0x40)
	QLowEnergyCharacteristic__ExtendedProperty QLowEnergyCharacteristic__PropertyType = QLowEnergyCharacteristic__PropertyType(0x80)
)

func NewQLowEnergyCharacteristic() *QLowEnergyCharacteristic {
	var tmpValue = NewQLowEnergyCharacteristicFromPointer(C.QLowEnergyCharacteristic_NewQLowEnergyCharacteristic())
	runtime.SetFinalizer(tmpValue, (*QLowEnergyCharacteristic).DestroyQLowEnergyCharacteristic)
	return tmpValue
}

func NewQLowEnergyCharacteristic2(other QLowEnergyCharacteristic_ITF) *QLowEnergyCharacteristic {
	var tmpValue = NewQLowEnergyCharacteristicFromPointer(C.QLowEnergyCharacteristic_NewQLowEnergyCharacteristic2(PointerFromQLowEnergyCharacteristic(other)))
	runtime.SetFinalizer(tmpValue, (*QLowEnergyCharacteristic).DestroyQLowEnergyCharacteristic)
	return tmpValue
}

func (ptr *QLowEnergyCharacteristic) DestroyQLowEnergyCharacteristic() {
	if ptr.Pointer() != nil {
		C.QLowEnergyCharacteristic_DestroyQLowEnergyCharacteristic(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QLowEnergyCharacteristic) Uuid() *QBluetoothUuid {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothUuidFromPointer(C.QLowEnergyCharacteristic_Uuid(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyCharacteristic) Value() *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QLowEnergyCharacteristic_Value(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyCharacteristic) Descriptors() []*QLowEnergyDescriptor {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtBluetooth_PackedList) []*QLowEnergyDescriptor {
			var out = make([]*QLowEnergyDescriptor, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQLowEnergyCharacteristicFromPointer(l.data).__descriptors_atList(i)
			}
			return out
		}(C.QLowEnergyCharacteristic_Descriptors(ptr.Pointer()))
	}
	return make([]*QLowEnergyDescriptor, 0)
}

func (ptr *QLowEnergyCharacteristic) Properties() QLowEnergyCharacteristic__PropertyType {
	if ptr.Pointer() != nil {
		return QLowEnergyCharacteristic__PropertyType(C.QLowEnergyCharacteristic_Properties(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyCharacteristic) Descriptor(uuid QBluetoothUuid_ITF) *QLowEnergyDescriptor {
	if ptr.Pointer() != nil {
		var tmpValue = NewQLowEnergyDescriptorFromPointer(C.QLowEnergyCharacteristic_Descriptor(ptr.Pointer(), PointerFromQBluetoothUuid(uuid)))
		runtime.SetFinalizer(tmpValue, (*QLowEnergyDescriptor).DestroyQLowEnergyDescriptor)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyCharacteristic) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QLowEnergyCharacteristic_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLowEnergyCharacteristic) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QLowEnergyCharacteristic_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLowEnergyCharacteristic) __descriptors_atList(i int) *QLowEnergyDescriptor {
	if ptr.Pointer() != nil {
		var tmpValue = NewQLowEnergyDescriptorFromPointer(C.QLowEnergyCharacteristic___descriptors_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QLowEnergyDescriptor).DestroyQLowEnergyDescriptor)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyCharacteristic) __descriptors_setList(i QLowEnergyDescriptor_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyCharacteristic___descriptors_setList(ptr.Pointer(), PointerFromQLowEnergyDescriptor(i))
	}
}

func (ptr *QLowEnergyCharacteristic) __descriptors_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QLowEnergyCharacteristic___descriptors_newList(ptr.Pointer()))
}

type QLowEnergyCharacteristicData struct {
	ptr unsafe.Pointer
}

type QLowEnergyCharacteristicData_ITF interface {
	QLowEnergyCharacteristicData_PTR() *QLowEnergyCharacteristicData
}

func (ptr *QLowEnergyCharacteristicData) QLowEnergyCharacteristicData_PTR() *QLowEnergyCharacteristicData {
	return ptr
}

func (ptr *QLowEnergyCharacteristicData) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QLowEnergyCharacteristicData) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
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
func NewQLowEnergyCharacteristicData() *QLowEnergyCharacteristicData {
	var tmpValue = NewQLowEnergyCharacteristicDataFromPointer(C.QLowEnergyCharacteristicData_NewQLowEnergyCharacteristicData())
	runtime.SetFinalizer(tmpValue, (*QLowEnergyCharacteristicData).DestroyQLowEnergyCharacteristicData)
	return tmpValue
}

func NewQLowEnergyCharacteristicData2(other QLowEnergyCharacteristicData_ITF) *QLowEnergyCharacteristicData {
	var tmpValue = NewQLowEnergyCharacteristicDataFromPointer(C.QLowEnergyCharacteristicData_NewQLowEnergyCharacteristicData2(PointerFromQLowEnergyCharacteristicData(other)))
	runtime.SetFinalizer(tmpValue, (*QLowEnergyCharacteristicData).DestroyQLowEnergyCharacteristicData)
	return tmpValue
}

func (ptr *QLowEnergyCharacteristicData) AddDescriptor(descriptor QLowEnergyDescriptorData_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyCharacteristicData_AddDescriptor(ptr.Pointer(), PointerFromQLowEnergyDescriptorData(descriptor))
	}
}

func (ptr *QLowEnergyCharacteristicData) SetDescriptors(descriptors []*QLowEnergyDescriptorData) {
	if ptr.Pointer() != nil {
		C.QLowEnergyCharacteristicData_SetDescriptors(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewQLowEnergyCharacteristicDataFromPointer(NewQLowEnergyCharacteristicDataFromPointer(nil).__setDescriptors_descriptors_newList())
			for _, v := range descriptors {
				tmpList.__setDescriptors_descriptors_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QLowEnergyCharacteristicData) SetProperties(properties QLowEnergyCharacteristic__PropertyType) {
	if ptr.Pointer() != nil {
		C.QLowEnergyCharacteristicData_SetProperties(ptr.Pointer(), C.longlong(properties))
	}
}

func (ptr *QLowEnergyCharacteristicData) SetReadConstraints(constraints QBluetooth__AttAccessConstraint) {
	if ptr.Pointer() != nil {
		C.QLowEnergyCharacteristicData_SetReadConstraints(ptr.Pointer(), C.longlong(constraints))
	}
}

func (ptr *QLowEnergyCharacteristicData) SetUuid(uuid QBluetoothUuid_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyCharacteristicData_SetUuid(ptr.Pointer(), PointerFromQBluetoothUuid(uuid))
	}
}

func (ptr *QLowEnergyCharacteristicData) SetValue(value core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyCharacteristicData_SetValue(ptr.Pointer(), core.PointerFromQByteArray(value))
	}
}

func (ptr *QLowEnergyCharacteristicData) SetValueLength(minimum int, maximum int) {
	if ptr.Pointer() != nil {
		C.QLowEnergyCharacteristicData_SetValueLength(ptr.Pointer(), C.int(int32(minimum)), C.int(int32(maximum)))
	}
}

func (ptr *QLowEnergyCharacteristicData) SetWriteConstraints(constraints QBluetooth__AttAccessConstraint) {
	if ptr.Pointer() != nil {
		C.QLowEnergyCharacteristicData_SetWriteConstraints(ptr.Pointer(), C.longlong(constraints))
	}
}

func (ptr *QLowEnergyCharacteristicData) Swap(other QLowEnergyCharacteristicData_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyCharacteristicData_Swap(ptr.Pointer(), PointerFromQLowEnergyCharacteristicData(other))
	}
}

func (ptr *QLowEnergyCharacteristicData) DestroyQLowEnergyCharacteristicData() {
	if ptr.Pointer() != nil {
		C.QLowEnergyCharacteristicData_DestroyQLowEnergyCharacteristicData(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QLowEnergyCharacteristicData) ReadConstraints() QBluetooth__AttAccessConstraint {
	if ptr.Pointer() != nil {
		return QBluetooth__AttAccessConstraint(C.QLowEnergyCharacteristicData_ReadConstraints(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyCharacteristicData) WriteConstraints() QBluetooth__AttAccessConstraint {
	if ptr.Pointer() != nil {
		return QBluetooth__AttAccessConstraint(C.QLowEnergyCharacteristicData_WriteConstraints(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyCharacteristicData) Uuid() *QBluetoothUuid {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothUuidFromPointer(C.QLowEnergyCharacteristicData_Uuid(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyCharacteristicData) Value() *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QLowEnergyCharacteristicData_Value(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyCharacteristicData) Descriptors() []*QLowEnergyDescriptorData {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtBluetooth_PackedList) []*QLowEnergyDescriptorData {
			var out = make([]*QLowEnergyDescriptorData, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQLowEnergyCharacteristicDataFromPointer(l.data).__descriptors_atList(i)
			}
			return out
		}(C.QLowEnergyCharacteristicData_Descriptors(ptr.Pointer()))
	}
	return make([]*QLowEnergyDescriptorData, 0)
}

func (ptr *QLowEnergyCharacteristicData) Properties() QLowEnergyCharacteristic__PropertyType {
	if ptr.Pointer() != nil {
		return QLowEnergyCharacteristic__PropertyType(C.QLowEnergyCharacteristicData_Properties(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyCharacteristicData) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QLowEnergyCharacteristicData_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLowEnergyCharacteristicData) MaximumValueLength() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QLowEnergyCharacteristicData_MaximumValueLength(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLowEnergyCharacteristicData) MinimumValueLength() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QLowEnergyCharacteristicData_MinimumValueLength(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLowEnergyCharacteristicData) __setDescriptors_descriptors_atList(i int) *QLowEnergyDescriptorData {
	if ptr.Pointer() != nil {
		var tmpValue = NewQLowEnergyDescriptorDataFromPointer(C.QLowEnergyCharacteristicData___setDescriptors_descriptors_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QLowEnergyDescriptorData).DestroyQLowEnergyDescriptorData)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyCharacteristicData) __setDescriptors_descriptors_setList(i QLowEnergyDescriptorData_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyCharacteristicData___setDescriptors_descriptors_setList(ptr.Pointer(), PointerFromQLowEnergyDescriptorData(i))
	}
}

func (ptr *QLowEnergyCharacteristicData) __setDescriptors_descriptors_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QLowEnergyCharacteristicData___setDescriptors_descriptors_newList(ptr.Pointer()))
}

func (ptr *QLowEnergyCharacteristicData) __descriptors_atList(i int) *QLowEnergyDescriptorData {
	if ptr.Pointer() != nil {
		var tmpValue = NewQLowEnergyDescriptorDataFromPointer(C.QLowEnergyCharacteristicData___descriptors_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QLowEnergyDescriptorData).DestroyQLowEnergyDescriptorData)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyCharacteristicData) __descriptors_setList(i QLowEnergyDescriptorData_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyCharacteristicData___descriptors_setList(ptr.Pointer(), PointerFromQLowEnergyDescriptorData(i))
	}
}

func (ptr *QLowEnergyCharacteristicData) __descriptors_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QLowEnergyCharacteristicData___descriptors_newList(ptr.Pointer()))
}

type QLowEnergyConnectionParameters struct {
	ptr unsafe.Pointer
}

type QLowEnergyConnectionParameters_ITF interface {
	QLowEnergyConnectionParameters_PTR() *QLowEnergyConnectionParameters
}

func (ptr *QLowEnergyConnectionParameters) QLowEnergyConnectionParameters_PTR() *QLowEnergyConnectionParameters {
	return ptr
}

func (ptr *QLowEnergyConnectionParameters) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QLowEnergyConnectionParameters) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
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
func NewQLowEnergyConnectionParameters() *QLowEnergyConnectionParameters {
	var tmpValue = NewQLowEnergyConnectionParametersFromPointer(C.QLowEnergyConnectionParameters_NewQLowEnergyConnectionParameters())
	runtime.SetFinalizer(tmpValue, (*QLowEnergyConnectionParameters).DestroyQLowEnergyConnectionParameters)
	return tmpValue
}

func NewQLowEnergyConnectionParameters2(other QLowEnergyConnectionParameters_ITF) *QLowEnergyConnectionParameters {
	var tmpValue = NewQLowEnergyConnectionParametersFromPointer(C.QLowEnergyConnectionParameters_NewQLowEnergyConnectionParameters2(PointerFromQLowEnergyConnectionParameters(other)))
	runtime.SetFinalizer(tmpValue, (*QLowEnergyConnectionParameters).DestroyQLowEnergyConnectionParameters)
	return tmpValue
}

func (ptr *QLowEnergyConnectionParameters) SetIntervalRange(minimum float64, maximum float64) {
	if ptr.Pointer() != nil {
		C.QLowEnergyConnectionParameters_SetIntervalRange(ptr.Pointer(), C.double(minimum), C.double(maximum))
	}
}

func (ptr *QLowEnergyConnectionParameters) SetLatency(latency int) {
	if ptr.Pointer() != nil {
		C.QLowEnergyConnectionParameters_SetLatency(ptr.Pointer(), C.int(int32(latency)))
	}
}

func (ptr *QLowEnergyConnectionParameters) SetSupervisionTimeout(timeout int) {
	if ptr.Pointer() != nil {
		C.QLowEnergyConnectionParameters_SetSupervisionTimeout(ptr.Pointer(), C.int(int32(timeout)))
	}
}

func (ptr *QLowEnergyConnectionParameters) Swap(other QLowEnergyConnectionParameters_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyConnectionParameters_Swap(ptr.Pointer(), PointerFromQLowEnergyConnectionParameters(other))
	}
}

func (ptr *QLowEnergyConnectionParameters) DestroyQLowEnergyConnectionParameters() {
	if ptr.Pointer() != nil {
		C.QLowEnergyConnectionParameters_DestroyQLowEnergyConnectionParameters(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QLowEnergyConnectionParameters) MaximumInterval() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QLowEnergyConnectionParameters_MaximumInterval(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyConnectionParameters) MinimumInterval() float64 {
	if ptr.Pointer() != nil {
		return float64(C.QLowEnergyConnectionParameters_MinimumInterval(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyConnectionParameters) Latency() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QLowEnergyConnectionParameters_Latency(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLowEnergyConnectionParameters) SupervisionTimeout() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QLowEnergyConnectionParameters_SupervisionTimeout(ptr.Pointer())))
	}
	return 0
}

type QLowEnergyController struct {
	core.QObject
}

type QLowEnergyController_ITF interface {
	core.QObject_ITF
	QLowEnergyController_PTR() *QLowEnergyController
}

func (ptr *QLowEnergyController) QLowEnergyController_PTR() *QLowEnergyController {
	return ptr
}

func (ptr *QLowEnergyController) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QLowEnergyController) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
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

//go:generate stringer -type=QLowEnergyController__ControllerState
//QLowEnergyController::ControllerState
type QLowEnergyController__ControllerState int64

const (
	QLowEnergyController__UnconnectedState QLowEnergyController__ControllerState = QLowEnergyController__ControllerState(0)
	QLowEnergyController__ConnectingState  QLowEnergyController__ControllerState = QLowEnergyController__ControllerState(1)
	QLowEnergyController__ConnectedState   QLowEnergyController__ControllerState = QLowEnergyController__ControllerState(2)
	QLowEnergyController__DiscoveringState QLowEnergyController__ControllerState = QLowEnergyController__ControllerState(3)
	QLowEnergyController__DiscoveredState  QLowEnergyController__ControllerState = QLowEnergyController__ControllerState(4)
	QLowEnergyController__ClosingState     QLowEnergyController__ControllerState = QLowEnergyController__ControllerState(5)
	QLowEnergyController__AdvertisingState QLowEnergyController__ControllerState = QLowEnergyController__ControllerState(6)
)

//go:generate stringer -type=QLowEnergyController__Error
//QLowEnergyController::Error
type QLowEnergyController__Error int64

const (
	QLowEnergyController__NoError                      QLowEnergyController__Error = QLowEnergyController__Error(0)
	QLowEnergyController__UnknownError                 QLowEnergyController__Error = QLowEnergyController__Error(1)
	QLowEnergyController__UnknownRemoteDeviceError     QLowEnergyController__Error = QLowEnergyController__Error(2)
	QLowEnergyController__NetworkError                 QLowEnergyController__Error = QLowEnergyController__Error(3)
	QLowEnergyController__InvalidBluetoothAdapterError QLowEnergyController__Error = QLowEnergyController__Error(4)
	QLowEnergyController__ConnectionError              QLowEnergyController__Error = QLowEnergyController__Error(5)
	QLowEnergyController__AdvertisingError             QLowEnergyController__Error = QLowEnergyController__Error(6)
)

//go:generate stringer -type=QLowEnergyController__RemoteAddressType
//QLowEnergyController::RemoteAddressType
type QLowEnergyController__RemoteAddressType int64

const (
	QLowEnergyController__PublicAddress QLowEnergyController__RemoteAddressType = QLowEnergyController__RemoteAddressType(0)
	QLowEnergyController__RandomAddress QLowEnergyController__RemoteAddressType = QLowEnergyController__RemoteAddressType(1)
)

//go:generate stringer -type=QLowEnergyController__Role
//QLowEnergyController::Role
type QLowEnergyController__Role int64

const (
	QLowEnergyController__CentralRole    QLowEnergyController__Role = QLowEnergyController__Role(0)
	QLowEnergyController__PeripheralRole QLowEnergyController__Role = QLowEnergyController__Role(1)
)

func QLowEnergyController_CreateCentral(remoteDevice QBluetoothDeviceInfo_ITF, parent core.QObject_ITF) *QLowEnergyController {
	var tmpValue = NewQLowEnergyControllerFromPointer(C.QLowEnergyController_QLowEnergyController_CreateCentral(PointerFromQBluetoothDeviceInfo(remoteDevice), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QLowEnergyController) CreateCentral(remoteDevice QBluetoothDeviceInfo_ITF, parent core.QObject_ITF) *QLowEnergyController {
	var tmpValue = NewQLowEnergyControllerFromPointer(C.QLowEnergyController_QLowEnergyController_CreateCentral(PointerFromQBluetoothDeviceInfo(remoteDevice), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QLowEnergyController_CreatePeripheral(parent core.QObject_ITF) *QLowEnergyController {
	var tmpValue = NewQLowEnergyControllerFromPointer(C.QLowEnergyController_QLowEnergyController_CreatePeripheral(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QLowEnergyController) CreatePeripheral(parent core.QObject_ITF) *QLowEnergyController {
	var tmpValue = NewQLowEnergyControllerFromPointer(C.QLowEnergyController_QLowEnergyController_CreatePeripheral(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QLowEnergyController) AddService(service QLowEnergyServiceData_ITF, parent core.QObject_ITF) *QLowEnergyService {
	if ptr.Pointer() != nil {
		var tmpValue = NewQLowEnergyServiceFromPointer(C.QLowEnergyController_AddService(ptr.Pointer(), PointerFromQLowEnergyServiceData(service), core.PointerFromQObject(parent)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyController) CreateServiceObject(serviceUuid QBluetoothUuid_ITF, parent core.QObject_ITF) *QLowEnergyService {
	if ptr.Pointer() != nil {
		var tmpValue = NewQLowEnergyServiceFromPointer(C.QLowEnergyController_CreateServiceObject(ptr.Pointer(), PointerFromQBluetoothUuid(serviceUuid), core.PointerFromQObject(parent)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyController) ConnectToDevice() {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_ConnectToDevice(ptr.Pointer())
	}
}

//export callbackQLowEnergyController_Connected
func callbackQLowEnergyController_Connected(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connected"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QLowEnergyController) ConnectConnected(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "connected") {
			C.QLowEnergyController_ConnectConnected(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "connected"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "connected", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "connected", f)
		}
	}
}

func (ptr *QLowEnergyController) DisconnectConnected() {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_DisconnectConnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "connected")
	}
}

func (ptr *QLowEnergyController) Connected() {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_Connected(ptr.Pointer())
	}
}

//export callbackQLowEnergyController_ConnectionUpdated
func callbackQLowEnergyController_ConnectionUpdated(ptr unsafe.Pointer, newParameters unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectionUpdated"); signal != nil {
		signal.(func(*QLowEnergyConnectionParameters))(NewQLowEnergyConnectionParametersFromPointer(newParameters))
	}

}

func (ptr *QLowEnergyController) ConnectConnectionUpdated(f func(newParameters *QLowEnergyConnectionParameters)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "connectionUpdated") {
			C.QLowEnergyController_ConnectConnectionUpdated(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "connectionUpdated"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "connectionUpdated", func(newParameters *QLowEnergyConnectionParameters) {
				signal.(func(*QLowEnergyConnectionParameters))(newParameters)
				f(newParameters)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "connectionUpdated", f)
		}
	}
}

func (ptr *QLowEnergyController) DisconnectConnectionUpdated() {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_DisconnectConnectionUpdated(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "connectionUpdated")
	}
}

func (ptr *QLowEnergyController) ConnectionUpdated(newParameters QLowEnergyConnectionParameters_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_ConnectionUpdated(ptr.Pointer(), PointerFromQLowEnergyConnectionParameters(newParameters))
	}
}

func (ptr *QLowEnergyController) DisconnectFromDevice() {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_DisconnectFromDevice(ptr.Pointer())
	}
}

//export callbackQLowEnergyController_Disconnected
func callbackQLowEnergyController_Disconnected(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnected"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QLowEnergyController) ConnectDisconnected(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "disconnected") {
			C.QLowEnergyController_ConnectDisconnected(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "disconnected"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "disconnected", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "disconnected", f)
		}
	}
}

func (ptr *QLowEnergyController) DisconnectDisconnected() {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_DisconnectDisconnected(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "disconnected")
	}
}

func (ptr *QLowEnergyController) Disconnected() {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_Disconnected(ptr.Pointer())
	}
}

func (ptr *QLowEnergyController) DiscoverServices() {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_DiscoverServices(ptr.Pointer())
	}
}

//export callbackQLowEnergyController_DiscoveryFinished
func callbackQLowEnergyController_DiscoveryFinished(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "discoveryFinished"); signal != nil {
		signal.(func())()
	}

}

func (ptr *QLowEnergyController) ConnectDiscoveryFinished(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "discoveryFinished") {
			C.QLowEnergyController_ConnectDiscoveryFinished(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "discoveryFinished"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "discoveryFinished", func() {
				signal.(func())()
				f()
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "discoveryFinished", f)
		}
	}
}

func (ptr *QLowEnergyController) DisconnectDiscoveryFinished() {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_DisconnectDiscoveryFinished(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "discoveryFinished")
	}
}

func (ptr *QLowEnergyController) DiscoveryFinished() {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_DiscoveryFinished(ptr.Pointer())
	}
}

//export callbackQLowEnergyController_Error2
func callbackQLowEnergyController_Error2(ptr unsafe.Pointer, newError C.longlong) {
	if signal := qt.GetSignal(ptr, "error2"); signal != nil {
		signal.(func(QLowEnergyController__Error))(QLowEnergyController__Error(newError))
	}

}

func (ptr *QLowEnergyController) ConnectError2(f func(newError QLowEnergyController__Error)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "error2") {
			C.QLowEnergyController_ConnectError2(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "error2"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "error2", func(newError QLowEnergyController__Error) {
				signal.(func(QLowEnergyController__Error))(newError)
				f(newError)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "error2", f)
		}
	}
}

func (ptr *QLowEnergyController) DisconnectError2() {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "error2")
	}
}

func (ptr *QLowEnergyController) Error2(newError QLowEnergyController__Error) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_Error2(ptr.Pointer(), C.longlong(newError))
	}
}

func (ptr *QLowEnergyController) RequestConnectionUpdate(parameters QLowEnergyConnectionParameters_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_RequestConnectionUpdate(ptr.Pointer(), PointerFromQLowEnergyConnectionParameters(parameters))
	}
}

//export callbackQLowEnergyController_ServiceDiscovered
func callbackQLowEnergyController_ServiceDiscovered(ptr unsafe.Pointer, newService unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "serviceDiscovered"); signal != nil {
		signal.(func(*QBluetoothUuid))(NewQBluetoothUuidFromPointer(newService))
	}

}

func (ptr *QLowEnergyController) ConnectServiceDiscovered(f func(newService *QBluetoothUuid)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "serviceDiscovered") {
			C.QLowEnergyController_ConnectServiceDiscovered(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "serviceDiscovered"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "serviceDiscovered", func(newService *QBluetoothUuid) {
				signal.(func(*QBluetoothUuid))(newService)
				f(newService)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "serviceDiscovered", f)
		}
	}
}

func (ptr *QLowEnergyController) DisconnectServiceDiscovered() {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_DisconnectServiceDiscovered(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "serviceDiscovered")
	}
}

func (ptr *QLowEnergyController) ServiceDiscovered(newService QBluetoothUuid_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_ServiceDiscovered(ptr.Pointer(), PointerFromQBluetoothUuid(newService))
	}
}

func (ptr *QLowEnergyController) SetRemoteAddressType(ty QLowEnergyController__RemoteAddressType) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_SetRemoteAddressType(ptr.Pointer(), C.longlong(ty))
	}
}

func (ptr *QLowEnergyController) StartAdvertising(parameters QLowEnergyAdvertisingParameters_ITF, advertisingData QLowEnergyAdvertisingData_ITF, scanResponseData QLowEnergyAdvertisingData_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_StartAdvertising(ptr.Pointer(), PointerFromQLowEnergyAdvertisingParameters(parameters), PointerFromQLowEnergyAdvertisingData(advertisingData), PointerFromQLowEnergyAdvertisingData(scanResponseData))
	}
}

//export callbackQLowEnergyController_StateChanged
func callbackQLowEnergyController_StateChanged(ptr unsafe.Pointer, state C.longlong) {
	if signal := qt.GetSignal(ptr, "stateChanged"); signal != nil {
		signal.(func(QLowEnergyController__ControllerState))(QLowEnergyController__ControllerState(state))
	}

}

func (ptr *QLowEnergyController) ConnectStateChanged(f func(state QLowEnergyController__ControllerState)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "stateChanged") {
			C.QLowEnergyController_ConnectStateChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "stateChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "stateChanged", func(state QLowEnergyController__ControllerState) {
				signal.(func(QLowEnergyController__ControllerState))(state)
				f(state)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stateChanged", f)
		}
	}
}

func (ptr *QLowEnergyController) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "stateChanged")
	}
}

func (ptr *QLowEnergyController) StateChanged(state QLowEnergyController__ControllerState) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_StateChanged(ptr.Pointer(), C.longlong(state))
	}
}

func (ptr *QLowEnergyController) StopAdvertising() {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_StopAdvertising(ptr.Pointer())
	}
}

func (ptr *QLowEnergyController) DestroyQLowEnergyController() {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_DestroyQLowEnergyController(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QLowEnergyController) State() QLowEnergyController__ControllerState {
	if ptr.Pointer() != nil {
		return QLowEnergyController__ControllerState(C.QLowEnergyController_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyController) Error() QLowEnergyController__Error {
	if ptr.Pointer() != nil {
		return QLowEnergyController__Error(C.QLowEnergyController_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyController) LocalAddress() *QBluetoothAddress {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothAddressFromPointer(C.QLowEnergyController_LocalAddress(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyController) RemoteAddress() *QBluetoothAddress {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothAddressFromPointer(C.QLowEnergyController_RemoteAddress(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyController) RemoteDeviceUuid() *QBluetoothUuid {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothUuidFromPointer(C.QLowEnergyController_RemoteDeviceUuid(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyController) Services() []*QBluetoothUuid {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtBluetooth_PackedList) []*QBluetoothUuid {
			var out = make([]*QBluetoothUuid, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQLowEnergyControllerFromPointer(l.data).__services_atList(i)
			}
			return out
		}(C.QLowEnergyController_Services(ptr.Pointer()))
	}
	return make([]*QBluetoothUuid, 0)
}

func (ptr *QLowEnergyController) ErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QLowEnergyController_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLowEnergyController) RemoteName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QLowEnergyController_RemoteName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLowEnergyController) RemoteAddressType() QLowEnergyController__RemoteAddressType {
	if ptr.Pointer() != nil {
		return QLowEnergyController__RemoteAddressType(C.QLowEnergyController_RemoteAddressType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyController) Role() QLowEnergyController__Role {
	if ptr.Pointer() != nil {
		return QLowEnergyController__Role(C.QLowEnergyController_Role(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyController) __services_atList(i int) *QBluetoothUuid {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothUuidFromPointer(C.QLowEnergyController___services_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyController) __services_setList(i QBluetoothUuid_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController___services_setList(ptr.Pointer(), PointerFromQBluetoothUuid(i))
	}
}

func (ptr *QLowEnergyController) __services_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QLowEnergyController___services_newList(ptr.Pointer()))
}

func (ptr *QLowEnergyController) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QLowEnergyController___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyController) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QLowEnergyController) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QLowEnergyController___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QLowEnergyController) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QLowEnergyController___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyController) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QLowEnergyController) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QLowEnergyController___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QLowEnergyController) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QLowEnergyController___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyController) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QLowEnergyController) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.QLowEnergyController___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QLowEnergyController) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QLowEnergyController___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyController) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QLowEnergyController) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QLowEnergyController___findChildren_newList(ptr.Pointer()))
}

func (ptr *QLowEnergyController) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QLowEnergyController___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyController) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QLowEnergyController) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QLowEnergyController___children_newList(ptr.Pointer()))
}

//export callbackQLowEnergyController_Event
func callbackQLowEnergyController_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLowEnergyControllerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QLowEnergyController) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QLowEnergyController_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQLowEnergyController_EventFilter
func callbackQLowEnergyController_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLowEnergyControllerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QLowEnergyController) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QLowEnergyController_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQLowEnergyController_ChildEvent
func callbackQLowEnergyController_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQLowEnergyControllerFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QLowEnergyController) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQLowEnergyController_ConnectNotify
func callbackQLowEnergyController_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQLowEnergyControllerFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QLowEnergyController) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQLowEnergyController_CustomEvent
func callbackQLowEnergyController_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQLowEnergyControllerFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QLowEnergyController) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQLowEnergyController_DeleteLater
func callbackQLowEnergyController_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQLowEnergyControllerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QLowEnergyController) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQLowEnergyController_Destroyed
func callbackQLowEnergyController_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQLowEnergyController_DisconnectNotify
func callbackQLowEnergyController_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQLowEnergyControllerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QLowEnergyController) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQLowEnergyController_ObjectNameChanged
func callbackQLowEnergyController_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtBluetooth_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQLowEnergyController_TimerEvent
func callbackQLowEnergyController_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQLowEnergyControllerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QLowEnergyController) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQLowEnergyController_MetaObject
func callbackQLowEnergyController_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQLowEnergyControllerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QLowEnergyController) MetaObjectDefault() *core.QMetaObject {
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

func (ptr *QLowEnergyDescriptor) QLowEnergyDescriptor_PTR() *QLowEnergyDescriptor {
	return ptr
}

func (ptr *QLowEnergyDescriptor) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QLowEnergyDescriptor) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
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
func NewQLowEnergyDescriptor() *QLowEnergyDescriptor {
	var tmpValue = NewQLowEnergyDescriptorFromPointer(C.QLowEnergyDescriptor_NewQLowEnergyDescriptor())
	runtime.SetFinalizer(tmpValue, (*QLowEnergyDescriptor).DestroyQLowEnergyDescriptor)
	return tmpValue
}

func NewQLowEnergyDescriptor2(other QLowEnergyDescriptor_ITF) *QLowEnergyDescriptor {
	var tmpValue = NewQLowEnergyDescriptorFromPointer(C.QLowEnergyDescriptor_NewQLowEnergyDescriptor2(PointerFromQLowEnergyDescriptor(other)))
	runtime.SetFinalizer(tmpValue, (*QLowEnergyDescriptor).DestroyQLowEnergyDescriptor)
	return tmpValue
}

func (ptr *QLowEnergyDescriptor) DestroyQLowEnergyDescriptor() {
	if ptr.Pointer() != nil {
		C.QLowEnergyDescriptor_DestroyQLowEnergyDescriptor(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QLowEnergyDescriptor) Uuid() *QBluetoothUuid {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothUuidFromPointer(C.QLowEnergyDescriptor_Uuid(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyDescriptor) Type() QBluetoothUuid__DescriptorType {
	if ptr.Pointer() != nil {
		return QBluetoothUuid__DescriptorType(C.QLowEnergyDescriptor_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyDescriptor) Value() *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QLowEnergyDescriptor_Value(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyDescriptor) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QLowEnergyDescriptor_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLowEnergyDescriptor) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QLowEnergyDescriptor_IsValid(ptr.Pointer()) != 0
	}
	return false
}

type QLowEnergyDescriptorData struct {
	ptr unsafe.Pointer
}

type QLowEnergyDescriptorData_ITF interface {
	QLowEnergyDescriptorData_PTR() *QLowEnergyDescriptorData
}

func (ptr *QLowEnergyDescriptorData) QLowEnergyDescriptorData_PTR() *QLowEnergyDescriptorData {
	return ptr
}

func (ptr *QLowEnergyDescriptorData) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QLowEnergyDescriptorData) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
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
func NewQLowEnergyDescriptorData() *QLowEnergyDescriptorData {
	var tmpValue = NewQLowEnergyDescriptorDataFromPointer(C.QLowEnergyDescriptorData_NewQLowEnergyDescriptorData())
	runtime.SetFinalizer(tmpValue, (*QLowEnergyDescriptorData).DestroyQLowEnergyDescriptorData)
	return tmpValue
}

func NewQLowEnergyDescriptorData2(uuid QBluetoothUuid_ITF, value core.QByteArray_ITF) *QLowEnergyDescriptorData {
	var tmpValue = NewQLowEnergyDescriptorDataFromPointer(C.QLowEnergyDescriptorData_NewQLowEnergyDescriptorData2(PointerFromQBluetoothUuid(uuid), core.PointerFromQByteArray(value)))
	runtime.SetFinalizer(tmpValue, (*QLowEnergyDescriptorData).DestroyQLowEnergyDescriptorData)
	return tmpValue
}

func NewQLowEnergyDescriptorData3(other QLowEnergyDescriptorData_ITF) *QLowEnergyDescriptorData {
	var tmpValue = NewQLowEnergyDescriptorDataFromPointer(C.QLowEnergyDescriptorData_NewQLowEnergyDescriptorData3(PointerFromQLowEnergyDescriptorData(other)))
	runtime.SetFinalizer(tmpValue, (*QLowEnergyDescriptorData).DestroyQLowEnergyDescriptorData)
	return tmpValue
}

func (ptr *QLowEnergyDescriptorData) SetReadPermissions(readable bool, constraints QBluetooth__AttAccessConstraint) {
	if ptr.Pointer() != nil {
		C.QLowEnergyDescriptorData_SetReadPermissions(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(readable))), C.longlong(constraints))
	}
}

func (ptr *QLowEnergyDescriptorData) SetUuid(uuid QBluetoothUuid_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyDescriptorData_SetUuid(ptr.Pointer(), PointerFromQBluetoothUuid(uuid))
	}
}

func (ptr *QLowEnergyDescriptorData) SetValue(value core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyDescriptorData_SetValue(ptr.Pointer(), core.PointerFromQByteArray(value))
	}
}

func (ptr *QLowEnergyDescriptorData) SetWritePermissions(writable bool, constraints QBluetooth__AttAccessConstraint) {
	if ptr.Pointer() != nil {
		C.QLowEnergyDescriptorData_SetWritePermissions(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(writable))), C.longlong(constraints))
	}
}

func (ptr *QLowEnergyDescriptorData) Swap(other QLowEnergyDescriptorData_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyDescriptorData_Swap(ptr.Pointer(), PointerFromQLowEnergyDescriptorData(other))
	}
}

func (ptr *QLowEnergyDescriptorData) DestroyQLowEnergyDescriptorData() {
	if ptr.Pointer() != nil {
		C.QLowEnergyDescriptorData_DestroyQLowEnergyDescriptorData(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QLowEnergyDescriptorData) ReadConstraints() QBluetooth__AttAccessConstraint {
	if ptr.Pointer() != nil {
		return QBluetooth__AttAccessConstraint(C.QLowEnergyDescriptorData_ReadConstraints(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyDescriptorData) WriteConstraints() QBluetooth__AttAccessConstraint {
	if ptr.Pointer() != nil {
		return QBluetooth__AttAccessConstraint(C.QLowEnergyDescriptorData_WriteConstraints(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyDescriptorData) Uuid() *QBluetoothUuid {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothUuidFromPointer(C.QLowEnergyDescriptorData_Uuid(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyDescriptorData) Value() *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QLowEnergyDescriptorData_Value(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyDescriptorData) IsReadable() bool {
	if ptr.Pointer() != nil {
		return C.QLowEnergyDescriptorData_IsReadable(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLowEnergyDescriptorData) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QLowEnergyDescriptorData_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLowEnergyDescriptorData) IsWritable() bool {
	if ptr.Pointer() != nil {
		return C.QLowEnergyDescriptorData_IsWritable(ptr.Pointer()) != 0
	}
	return false
}

type QLowEnergyService struct {
	core.QObject
}

type QLowEnergyService_ITF interface {
	core.QObject_ITF
	QLowEnergyService_PTR() *QLowEnergyService
}

func (ptr *QLowEnergyService) QLowEnergyService_PTR() *QLowEnergyService {
	return ptr
}

func (ptr *QLowEnergyService) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.QObject_PTR().Pointer()
	}
	return nil
}

func (ptr *QLowEnergyService) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.QObject_PTR().SetPointer(p)
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

//go:generate stringer -type=QLowEnergyService__ServiceError
//QLowEnergyService::ServiceError
type QLowEnergyService__ServiceError int64

const (
	QLowEnergyService__NoError                  QLowEnergyService__ServiceError = QLowEnergyService__ServiceError(0)
	QLowEnergyService__OperationError           QLowEnergyService__ServiceError = QLowEnergyService__ServiceError(1)
	QLowEnergyService__CharacteristicWriteError QLowEnergyService__ServiceError = QLowEnergyService__ServiceError(2)
	QLowEnergyService__DescriptorWriteError     QLowEnergyService__ServiceError = QLowEnergyService__ServiceError(3)
	QLowEnergyService__UnknownError             QLowEnergyService__ServiceError = QLowEnergyService__ServiceError(4)
	QLowEnergyService__CharacteristicReadError  QLowEnergyService__ServiceError = QLowEnergyService__ServiceError(5)
	QLowEnergyService__DescriptorReadError      QLowEnergyService__ServiceError = QLowEnergyService__ServiceError(6)
)

//go:generate stringer -type=QLowEnergyService__ServiceState
//QLowEnergyService::ServiceState
type QLowEnergyService__ServiceState int64

const (
	QLowEnergyService__InvalidService      QLowEnergyService__ServiceState = QLowEnergyService__ServiceState(0)
	QLowEnergyService__DiscoveryRequired   QLowEnergyService__ServiceState = QLowEnergyService__ServiceState(1)
	QLowEnergyService__DiscoveringServices QLowEnergyService__ServiceState = QLowEnergyService__ServiceState(2)
	QLowEnergyService__ServiceDiscovered   QLowEnergyService__ServiceState = QLowEnergyService__ServiceState(3)
	QLowEnergyService__LocalService        QLowEnergyService__ServiceState = QLowEnergyService__ServiceState(4)
)

//go:generate stringer -type=QLowEnergyService__ServiceType
//QLowEnergyService::ServiceType
type QLowEnergyService__ServiceType int64

const (
	QLowEnergyService__PrimaryService  QLowEnergyService__ServiceType = QLowEnergyService__ServiceType(0x0001)
	QLowEnergyService__IncludedService QLowEnergyService__ServiceType = QLowEnergyService__ServiceType(0x0002)
)

//go:generate stringer -type=QLowEnergyService__WriteMode
//QLowEnergyService::WriteMode
type QLowEnergyService__WriteMode int64

const (
	QLowEnergyService__WriteWithResponse    QLowEnergyService__WriteMode = QLowEnergyService__WriteMode(0)
	QLowEnergyService__WriteWithoutResponse QLowEnergyService__WriteMode = QLowEnergyService__WriteMode(1)
	QLowEnergyService__WriteSigned          QLowEnergyService__WriteMode = QLowEnergyService__WriteMode(2)
)

//export callbackQLowEnergyService_CharacteristicChanged
func callbackQLowEnergyService_CharacteristicChanged(ptr unsafe.Pointer, characteristic unsafe.Pointer, newValue unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "characteristicChanged"); signal != nil {
		signal.(func(*QLowEnergyCharacteristic, *core.QByteArray))(NewQLowEnergyCharacteristicFromPointer(characteristic), core.NewQByteArrayFromPointer(newValue))
	}

}

func (ptr *QLowEnergyService) ConnectCharacteristicChanged(f func(characteristic *QLowEnergyCharacteristic, newValue *core.QByteArray)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "characteristicChanged") {
			C.QLowEnergyService_ConnectCharacteristicChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "characteristicChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "characteristicChanged", func(characteristic *QLowEnergyCharacteristic, newValue *core.QByteArray) {
				signal.(func(*QLowEnergyCharacteristic, *core.QByteArray))(characteristic, newValue)
				f(characteristic, newValue)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "characteristicChanged", f)
		}
	}
}

func (ptr *QLowEnergyService) DisconnectCharacteristicChanged() {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_DisconnectCharacteristicChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "characteristicChanged")
	}
}

func (ptr *QLowEnergyService) CharacteristicChanged(characteristic QLowEnergyCharacteristic_ITF, newValue core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_CharacteristicChanged(ptr.Pointer(), PointerFromQLowEnergyCharacteristic(characteristic), core.PointerFromQByteArray(newValue))
	}
}

//export callbackQLowEnergyService_CharacteristicRead
func callbackQLowEnergyService_CharacteristicRead(ptr unsafe.Pointer, characteristic unsafe.Pointer, value unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "characteristicRead"); signal != nil {
		signal.(func(*QLowEnergyCharacteristic, *core.QByteArray))(NewQLowEnergyCharacteristicFromPointer(characteristic), core.NewQByteArrayFromPointer(value))
	}

}

func (ptr *QLowEnergyService) ConnectCharacteristicRead(f func(characteristic *QLowEnergyCharacteristic, value *core.QByteArray)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "characteristicRead") {
			C.QLowEnergyService_ConnectCharacteristicRead(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "characteristicRead"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "characteristicRead", func(characteristic *QLowEnergyCharacteristic, value *core.QByteArray) {
				signal.(func(*QLowEnergyCharacteristic, *core.QByteArray))(characteristic, value)
				f(characteristic, value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "characteristicRead", f)
		}
	}
}

func (ptr *QLowEnergyService) DisconnectCharacteristicRead() {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_DisconnectCharacteristicRead(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "characteristicRead")
	}
}

func (ptr *QLowEnergyService) CharacteristicRead(characteristic QLowEnergyCharacteristic_ITF, value core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_CharacteristicRead(ptr.Pointer(), PointerFromQLowEnergyCharacteristic(characteristic), core.PointerFromQByteArray(value))
	}
}

//export callbackQLowEnergyService_CharacteristicWritten
func callbackQLowEnergyService_CharacteristicWritten(ptr unsafe.Pointer, characteristic unsafe.Pointer, newValue unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "characteristicWritten"); signal != nil {
		signal.(func(*QLowEnergyCharacteristic, *core.QByteArray))(NewQLowEnergyCharacteristicFromPointer(characteristic), core.NewQByteArrayFromPointer(newValue))
	}

}

func (ptr *QLowEnergyService) ConnectCharacteristicWritten(f func(characteristic *QLowEnergyCharacteristic, newValue *core.QByteArray)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "characteristicWritten") {
			C.QLowEnergyService_ConnectCharacteristicWritten(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "characteristicWritten"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "characteristicWritten", func(characteristic *QLowEnergyCharacteristic, newValue *core.QByteArray) {
				signal.(func(*QLowEnergyCharacteristic, *core.QByteArray))(characteristic, newValue)
				f(characteristic, newValue)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "characteristicWritten", f)
		}
	}
}

func (ptr *QLowEnergyService) DisconnectCharacteristicWritten() {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_DisconnectCharacteristicWritten(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "characteristicWritten")
	}
}

func (ptr *QLowEnergyService) CharacteristicWritten(characteristic QLowEnergyCharacteristic_ITF, newValue core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_CharacteristicWritten(ptr.Pointer(), PointerFromQLowEnergyCharacteristic(characteristic), core.PointerFromQByteArray(newValue))
	}
}

//export callbackQLowEnergyService_DescriptorRead
func callbackQLowEnergyService_DescriptorRead(ptr unsafe.Pointer, descriptor unsafe.Pointer, value unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "descriptorRead"); signal != nil {
		signal.(func(*QLowEnergyDescriptor, *core.QByteArray))(NewQLowEnergyDescriptorFromPointer(descriptor), core.NewQByteArrayFromPointer(value))
	}

}

func (ptr *QLowEnergyService) ConnectDescriptorRead(f func(descriptor *QLowEnergyDescriptor, value *core.QByteArray)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "descriptorRead") {
			C.QLowEnergyService_ConnectDescriptorRead(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "descriptorRead"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "descriptorRead", func(descriptor *QLowEnergyDescriptor, value *core.QByteArray) {
				signal.(func(*QLowEnergyDescriptor, *core.QByteArray))(descriptor, value)
				f(descriptor, value)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "descriptorRead", f)
		}
	}
}

func (ptr *QLowEnergyService) DisconnectDescriptorRead() {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_DisconnectDescriptorRead(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "descriptorRead")
	}
}

func (ptr *QLowEnergyService) DescriptorRead(descriptor QLowEnergyDescriptor_ITF, value core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_DescriptorRead(ptr.Pointer(), PointerFromQLowEnergyDescriptor(descriptor), core.PointerFromQByteArray(value))
	}
}

//export callbackQLowEnergyService_DescriptorWritten
func callbackQLowEnergyService_DescriptorWritten(ptr unsafe.Pointer, descriptor unsafe.Pointer, newValue unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "descriptorWritten"); signal != nil {
		signal.(func(*QLowEnergyDescriptor, *core.QByteArray))(NewQLowEnergyDescriptorFromPointer(descriptor), core.NewQByteArrayFromPointer(newValue))
	}

}

func (ptr *QLowEnergyService) ConnectDescriptorWritten(f func(descriptor *QLowEnergyDescriptor, newValue *core.QByteArray)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "descriptorWritten") {
			C.QLowEnergyService_ConnectDescriptorWritten(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "descriptorWritten"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "descriptorWritten", func(descriptor *QLowEnergyDescriptor, newValue *core.QByteArray) {
				signal.(func(*QLowEnergyDescriptor, *core.QByteArray))(descriptor, newValue)
				f(descriptor, newValue)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "descriptorWritten", f)
		}
	}
}

func (ptr *QLowEnergyService) DisconnectDescriptorWritten() {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_DisconnectDescriptorWritten(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "descriptorWritten")
	}
}

func (ptr *QLowEnergyService) DescriptorWritten(descriptor QLowEnergyDescriptor_ITF, newValue core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_DescriptorWritten(ptr.Pointer(), PointerFromQLowEnergyDescriptor(descriptor), core.PointerFromQByteArray(newValue))
	}
}

func (ptr *QLowEnergyService) DiscoverDetails() {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_DiscoverDetails(ptr.Pointer())
	}
}

//export callbackQLowEnergyService_Error2
func callbackQLowEnergyService_Error2(ptr unsafe.Pointer, newError C.longlong) {
	if signal := qt.GetSignal(ptr, "error2"); signal != nil {
		signal.(func(QLowEnergyService__ServiceError))(QLowEnergyService__ServiceError(newError))
	}

}

func (ptr *QLowEnergyService) ConnectError2(f func(newError QLowEnergyService__ServiceError)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "error2") {
			C.QLowEnergyService_ConnectError2(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "error2"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "error2", func(newError QLowEnergyService__ServiceError) {
				signal.(func(QLowEnergyService__ServiceError))(newError)
				f(newError)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "error2", f)
		}
	}
}

func (ptr *QLowEnergyService) DisconnectError2() {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_DisconnectError2(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "error2")
	}
}

func (ptr *QLowEnergyService) Error2(newError QLowEnergyService__ServiceError) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_Error2(ptr.Pointer(), C.longlong(newError))
	}
}

func (ptr *QLowEnergyService) ReadCharacteristic(characteristic QLowEnergyCharacteristic_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_ReadCharacteristic(ptr.Pointer(), PointerFromQLowEnergyCharacteristic(characteristic))
	}
}

func (ptr *QLowEnergyService) ReadDescriptor(descriptor QLowEnergyDescriptor_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_ReadDescriptor(ptr.Pointer(), PointerFromQLowEnergyDescriptor(descriptor))
	}
}

//export callbackQLowEnergyService_StateChanged
func callbackQLowEnergyService_StateChanged(ptr unsafe.Pointer, newState C.longlong) {
	if signal := qt.GetSignal(ptr, "stateChanged"); signal != nil {
		signal.(func(QLowEnergyService__ServiceState))(QLowEnergyService__ServiceState(newState))
	}

}

func (ptr *QLowEnergyService) ConnectStateChanged(f func(newState QLowEnergyService__ServiceState)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "stateChanged") {
			C.QLowEnergyService_ConnectStateChanged(ptr.Pointer())
		}

		if signal := qt.LendSignal(ptr.Pointer(), "stateChanged"); signal != nil {
			qt.ConnectSignal(ptr.Pointer(), "stateChanged", func(newState QLowEnergyService__ServiceState) {
				signal.(func(QLowEnergyService__ServiceState))(newState)
				f(newState)
			})
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stateChanged", f)
		}
	}
}

func (ptr *QLowEnergyService) DisconnectStateChanged() {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_DisconnectStateChanged(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "stateChanged")
	}
}

func (ptr *QLowEnergyService) StateChanged(newState QLowEnergyService__ServiceState) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_StateChanged(ptr.Pointer(), C.longlong(newState))
	}
}

func (ptr *QLowEnergyService) WriteCharacteristic(characteristic QLowEnergyCharacteristic_ITF, newValue core.QByteArray_ITF, mode QLowEnergyService__WriteMode) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_WriteCharacteristic(ptr.Pointer(), PointerFromQLowEnergyCharacteristic(characteristic), core.PointerFromQByteArray(newValue), C.longlong(mode))
	}
}

func (ptr *QLowEnergyService) WriteDescriptor(descriptor QLowEnergyDescriptor_ITF, newValue core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_WriteDescriptor(ptr.Pointer(), PointerFromQLowEnergyDescriptor(descriptor), core.PointerFromQByteArray(newValue))
	}
}

func (ptr *QLowEnergyService) DestroyQLowEnergyService() {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_DestroyQLowEnergyService(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QLowEnergyService) ServiceUuid() *QBluetoothUuid {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothUuidFromPointer(C.QLowEnergyService_ServiceUuid(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyService) IncludedServices() []*QBluetoothUuid {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtBluetooth_PackedList) []*QBluetoothUuid {
			var out = make([]*QBluetoothUuid, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQLowEnergyServiceFromPointer(l.data).__includedServices_atList(i)
			}
			return out
		}(C.QLowEnergyService_IncludedServices(ptr.Pointer()))
	}
	return make([]*QBluetoothUuid, 0)
}

func (ptr *QLowEnergyService) Characteristics() []*QLowEnergyCharacteristic {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtBluetooth_PackedList) []*QLowEnergyCharacteristic {
			var out = make([]*QLowEnergyCharacteristic, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQLowEnergyServiceFromPointer(l.data).__characteristics_atList(i)
			}
			return out
		}(C.QLowEnergyService_Characteristics(ptr.Pointer()))
	}
	return make([]*QLowEnergyCharacteristic, 0)
}

func (ptr *QLowEnergyService) Characteristic(uuid QBluetoothUuid_ITF) *QLowEnergyCharacteristic {
	if ptr.Pointer() != nil {
		var tmpValue = NewQLowEnergyCharacteristicFromPointer(C.QLowEnergyService_Characteristic(ptr.Pointer(), PointerFromQBluetoothUuid(uuid)))
		runtime.SetFinalizer(tmpValue, (*QLowEnergyCharacteristic).DestroyQLowEnergyCharacteristic)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyService) State() QLowEnergyService__ServiceState {
	if ptr.Pointer() != nil {
		return QLowEnergyService__ServiceState(C.QLowEnergyService_State(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyService) Type() QLowEnergyService__ServiceType {
	if ptr.Pointer() != nil {
		return QLowEnergyService__ServiceType(C.QLowEnergyService_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyService) ServiceName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QLowEnergyService_ServiceName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLowEnergyService) Error() QLowEnergyService__ServiceError {
	if ptr.Pointer() != nil {
		return QLowEnergyService__ServiceError(C.QLowEnergyService_Error(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyService) Contains(characteristic QLowEnergyCharacteristic_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QLowEnergyService_Contains(ptr.Pointer(), PointerFromQLowEnergyCharacteristic(characteristic)) != 0
	}
	return false
}

func (ptr *QLowEnergyService) Contains2(descriptor QLowEnergyDescriptor_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QLowEnergyService_Contains2(ptr.Pointer(), PointerFromQLowEnergyDescriptor(descriptor)) != 0
	}
	return false
}

func (ptr *QLowEnergyService) __includedServices_atList(i int) *QBluetoothUuid {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothUuidFromPointer(C.QLowEnergyService___includedServices_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyService) __includedServices_setList(i QBluetoothUuid_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService___includedServices_setList(ptr.Pointer(), PointerFromQBluetoothUuid(i))
	}
}

func (ptr *QLowEnergyService) __includedServices_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QLowEnergyService___includedServices_newList(ptr.Pointer()))
}

func (ptr *QLowEnergyService) __characteristics_atList(i int) *QLowEnergyCharacteristic {
	if ptr.Pointer() != nil {
		var tmpValue = NewQLowEnergyCharacteristicFromPointer(C.QLowEnergyService___characteristics_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QLowEnergyCharacteristic).DestroyQLowEnergyCharacteristic)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyService) __characteristics_setList(i QLowEnergyCharacteristic_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService___characteristics_setList(ptr.Pointer(), PointerFromQLowEnergyCharacteristic(i))
	}
}

func (ptr *QLowEnergyService) __characteristics_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QLowEnergyService___characteristics_newList(ptr.Pointer()))
}

func (ptr *QLowEnergyService) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQByteArrayFromPointer(C.QLowEnergyService___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyService) __dynamicPropertyNames_setList(i core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService___dynamicPropertyNames_setList(ptr.Pointer(), core.PointerFromQByteArray(i))
	}
}

func (ptr *QLowEnergyService) __dynamicPropertyNames_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QLowEnergyService___dynamicPropertyNames_newList(ptr.Pointer()))
}

func (ptr *QLowEnergyService) __findChildren_atList2(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QLowEnergyService___findChildren_atList2(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyService) __findChildren_setList2(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService___findChildren_setList2(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QLowEnergyService) __findChildren_newList2() unsafe.Pointer {
	return unsafe.Pointer(C.QLowEnergyService___findChildren_newList2(ptr.Pointer()))
}

func (ptr *QLowEnergyService) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QLowEnergyService___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyService) __findChildren_setList3(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService___findChildren_setList3(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QLowEnergyService) __findChildren_newList3() unsafe.Pointer {
	return unsafe.Pointer(C.QLowEnergyService___findChildren_newList3(ptr.Pointer()))
}

func (ptr *QLowEnergyService) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QLowEnergyService___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyService) __findChildren_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService___findChildren_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QLowEnergyService) __findChildren_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QLowEnergyService___findChildren_newList(ptr.Pointer()))
}

func (ptr *QLowEnergyService) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		var tmpValue = core.NewQObjectFromPointer(C.QLowEnergyService___children_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyService) __children_setList(i core.QObject_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService___children_setList(ptr.Pointer(), core.PointerFromQObject(i))
	}
}

func (ptr *QLowEnergyService) __children_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QLowEnergyService___children_newList(ptr.Pointer()))
}

//export callbackQLowEnergyService_Event
func callbackQLowEnergyService_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QEvent) bool)(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLowEnergyServiceFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QLowEnergyService) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QLowEnergyService_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e)) != 0
	}
	return false
}

//export callbackQLowEnergyService_EventFilter
func callbackQLowEnergyService_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt(signal.(func(*core.QObject, *core.QEvent) bool)(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLowEnergyServiceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QLowEnergyService) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QLowEnergyService_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event)) != 0
	}
	return false
}

//export callbackQLowEnergyService_ChildEvent
func callbackQLowEnergyService_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		signal.(func(*core.QChildEvent))(core.NewQChildEventFromPointer(event))
	} else {
		NewQLowEnergyServiceFromPointer(ptr).ChildEventDefault(core.NewQChildEventFromPointer(event))
	}
}

func (ptr *QLowEnergyService) ChildEventDefault(event core.QChildEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_ChildEventDefault(ptr.Pointer(), core.PointerFromQChildEvent(event))
	}
}

//export callbackQLowEnergyService_ConnectNotify
func callbackQLowEnergyService_ConnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQLowEnergyServiceFromPointer(ptr).ConnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QLowEnergyService) ConnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_ConnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQLowEnergyService_CustomEvent
func callbackQLowEnergyService_CustomEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "customEvent"); signal != nil {
		signal.(func(*core.QEvent))(core.NewQEventFromPointer(event))
	} else {
		NewQLowEnergyServiceFromPointer(ptr).CustomEventDefault(core.NewQEventFromPointer(event))
	}
}

func (ptr *QLowEnergyService) CustomEventDefault(event core.QEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_CustomEventDefault(ptr.Pointer(), core.PointerFromQEvent(event))
	}
}

//export callbackQLowEnergyService_DeleteLater
func callbackQLowEnergyService_DeleteLater(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deleteLater"); signal != nil {
		signal.(func())()
	} else {
		NewQLowEnergyServiceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QLowEnergyService) DeleteLaterDefault() {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_DeleteLaterDefault(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

//export callbackQLowEnergyService_Destroyed
func callbackQLowEnergyService_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		signal.(func(*core.QObject))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQLowEnergyService_DisconnectNotify
func callbackQLowEnergyService_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		signal.(func(*core.QMetaMethod))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQLowEnergyServiceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QLowEnergyService) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQLowEnergyService_ObjectNameChanged
func callbackQLowEnergyService_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtBluetooth_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		signal.(func(string))(cGoUnpackString(objectName))
	}

}

//export callbackQLowEnergyService_TimerEvent
func callbackQLowEnergyService_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		signal.(func(*core.QTimerEvent))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQLowEnergyServiceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QLowEnergyService) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
}

//export callbackQLowEnergyService_MetaObject
func callbackQLowEnergyService_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject(signal.(func() *core.QMetaObject)())
	}

	return core.PointerFromQMetaObject(NewQLowEnergyServiceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QLowEnergyService) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QLowEnergyService_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

type QLowEnergyServiceData struct {
	ptr unsafe.Pointer
}

type QLowEnergyServiceData_ITF interface {
	QLowEnergyServiceData_PTR() *QLowEnergyServiceData
}

func (ptr *QLowEnergyServiceData) QLowEnergyServiceData_PTR() *QLowEnergyServiceData {
	return ptr
}

func (ptr *QLowEnergyServiceData) Pointer() unsafe.Pointer {
	if ptr != nil {
		return ptr.ptr
	}
	return nil
}

func (ptr *QLowEnergyServiceData) SetPointer(p unsafe.Pointer) {
	if ptr != nil {
		ptr.ptr = p
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

//go:generate stringer -type=QLowEnergyServiceData__ServiceType
//QLowEnergyServiceData::ServiceType
type QLowEnergyServiceData__ServiceType int64

const (
	QLowEnergyServiceData__ServiceTypePrimary   QLowEnergyServiceData__ServiceType = QLowEnergyServiceData__ServiceType(0x2800)
	QLowEnergyServiceData__ServiceTypeSecondary QLowEnergyServiceData__ServiceType = QLowEnergyServiceData__ServiceType(0x2801)
)

func NewQLowEnergyServiceData() *QLowEnergyServiceData {
	var tmpValue = NewQLowEnergyServiceDataFromPointer(C.QLowEnergyServiceData_NewQLowEnergyServiceData())
	runtime.SetFinalizer(tmpValue, (*QLowEnergyServiceData).DestroyQLowEnergyServiceData)
	return tmpValue
}

func NewQLowEnergyServiceData2(other QLowEnergyServiceData_ITF) *QLowEnergyServiceData {
	var tmpValue = NewQLowEnergyServiceDataFromPointer(C.QLowEnergyServiceData_NewQLowEnergyServiceData2(PointerFromQLowEnergyServiceData(other)))
	runtime.SetFinalizer(tmpValue, (*QLowEnergyServiceData).DestroyQLowEnergyServiceData)
	return tmpValue
}

func (ptr *QLowEnergyServiceData) AddCharacteristic(characteristic QLowEnergyCharacteristicData_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyServiceData_AddCharacteristic(ptr.Pointer(), PointerFromQLowEnergyCharacteristicData(characteristic))
	}
}

func (ptr *QLowEnergyServiceData) AddIncludedService(service QLowEnergyService_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyServiceData_AddIncludedService(ptr.Pointer(), PointerFromQLowEnergyService(service))
	}
}

func (ptr *QLowEnergyServiceData) SetCharacteristics(characteristics []*QLowEnergyCharacteristicData) {
	if ptr.Pointer() != nil {
		C.QLowEnergyServiceData_SetCharacteristics(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewQLowEnergyServiceDataFromPointer(NewQLowEnergyServiceDataFromPointer(nil).__setCharacteristics_characteristics_newList())
			for _, v := range characteristics {
				tmpList.__setCharacteristics_characteristics_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QLowEnergyServiceData) SetIncludedServices(services []*QLowEnergyService) {
	if ptr.Pointer() != nil {
		C.QLowEnergyServiceData_SetIncludedServices(ptr.Pointer(), func() unsafe.Pointer {
			var tmpList = NewQLowEnergyServiceDataFromPointer(NewQLowEnergyServiceDataFromPointer(nil).__setIncludedServices_services_newList())
			for _, v := range services {
				tmpList.__setIncludedServices_services_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QLowEnergyServiceData) SetType(ty QLowEnergyServiceData__ServiceType) {
	if ptr.Pointer() != nil {
		C.QLowEnergyServiceData_SetType(ptr.Pointer(), C.longlong(ty))
	}
}

func (ptr *QLowEnergyServiceData) SetUuid(uuid QBluetoothUuid_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyServiceData_SetUuid(ptr.Pointer(), PointerFromQBluetoothUuid(uuid))
	}
}

func (ptr *QLowEnergyServiceData) Swap(other QLowEnergyServiceData_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyServiceData_Swap(ptr.Pointer(), PointerFromQLowEnergyServiceData(other))
	}
}

func (ptr *QLowEnergyServiceData) DestroyQLowEnergyServiceData() {
	if ptr.Pointer() != nil {
		C.QLowEnergyServiceData_DestroyQLowEnergyServiceData(ptr.Pointer())
		ptr.SetPointer(nil)
		runtime.SetFinalizer(ptr, nil)
	}
}

func (ptr *QLowEnergyServiceData) Uuid() *QBluetoothUuid {
	if ptr.Pointer() != nil {
		var tmpValue = NewQBluetoothUuidFromPointer(C.QLowEnergyServiceData_Uuid(ptr.Pointer()))
		runtime.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyServiceData) Characteristics() []*QLowEnergyCharacteristicData {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtBluetooth_PackedList) []*QLowEnergyCharacteristicData {
			var out = make([]*QLowEnergyCharacteristicData, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQLowEnergyServiceDataFromPointer(l.data).__characteristics_atList(i)
			}
			return out
		}(C.QLowEnergyServiceData_Characteristics(ptr.Pointer()))
	}
	return make([]*QLowEnergyCharacteristicData, 0)
}

func (ptr *QLowEnergyServiceData) IncludedServices() []*QLowEnergyService {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtBluetooth_PackedList) []*QLowEnergyService {
			var out = make([]*QLowEnergyService, int(l.len))
			for i := 0; i < int(l.len); i++ {
				out[i] = NewQLowEnergyServiceDataFromPointer(l.data).__includedServices_atList(i)
			}
			return out
		}(C.QLowEnergyServiceData_IncludedServices(ptr.Pointer()))
	}
	return make([]*QLowEnergyService, 0)
}

func (ptr *QLowEnergyServiceData) Type() QLowEnergyServiceData__ServiceType {
	if ptr.Pointer() != nil {
		return QLowEnergyServiceData__ServiceType(C.QLowEnergyServiceData_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyServiceData) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QLowEnergyServiceData_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QLowEnergyServiceData) __setCharacteristics_characteristics_atList(i int) *QLowEnergyCharacteristicData {
	if ptr.Pointer() != nil {
		var tmpValue = NewQLowEnergyCharacteristicDataFromPointer(C.QLowEnergyServiceData___setCharacteristics_characteristics_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QLowEnergyCharacteristicData).DestroyQLowEnergyCharacteristicData)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyServiceData) __setCharacteristics_characteristics_setList(i QLowEnergyCharacteristicData_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyServiceData___setCharacteristics_characteristics_setList(ptr.Pointer(), PointerFromQLowEnergyCharacteristicData(i))
	}
}

func (ptr *QLowEnergyServiceData) __setCharacteristics_characteristics_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QLowEnergyServiceData___setCharacteristics_characteristics_newList(ptr.Pointer()))
}

func (ptr *QLowEnergyServiceData) __setIncludedServices_services_atList(i int) *QLowEnergyService {
	if ptr.Pointer() != nil {
		var tmpValue = NewQLowEnergyServiceFromPointer(C.QLowEnergyServiceData___setIncludedServices_services_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyServiceData) __setIncludedServices_services_setList(i QLowEnergyService_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyServiceData___setIncludedServices_services_setList(ptr.Pointer(), PointerFromQLowEnergyService(i))
	}
}

func (ptr *QLowEnergyServiceData) __setIncludedServices_services_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QLowEnergyServiceData___setIncludedServices_services_newList(ptr.Pointer()))
}

func (ptr *QLowEnergyServiceData) __characteristics_atList(i int) *QLowEnergyCharacteristicData {
	if ptr.Pointer() != nil {
		var tmpValue = NewQLowEnergyCharacteristicDataFromPointer(C.QLowEnergyServiceData___characteristics_atList(ptr.Pointer(), C.int(int32(i))))
		runtime.SetFinalizer(tmpValue, (*QLowEnergyCharacteristicData).DestroyQLowEnergyCharacteristicData)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyServiceData) __characteristics_setList(i QLowEnergyCharacteristicData_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyServiceData___characteristics_setList(ptr.Pointer(), PointerFromQLowEnergyCharacteristicData(i))
	}
}

func (ptr *QLowEnergyServiceData) __characteristics_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QLowEnergyServiceData___characteristics_newList(ptr.Pointer()))
}

func (ptr *QLowEnergyServiceData) __includedServices_atList(i int) *QLowEnergyService {
	if ptr.Pointer() != nil {
		var tmpValue = NewQLowEnergyServiceFromPointer(C.QLowEnergyServiceData___includedServices_atList(ptr.Pointer(), C.int(int32(i))))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyServiceData) __includedServices_setList(i QLowEnergyService_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyServiceData___includedServices_setList(ptr.Pointer(), PointerFromQLowEnergyService(i))
	}
}

func (ptr *QLowEnergyServiceData) __includedServices_newList() unsafe.Pointer {
	return unsafe.Pointer(C.QLowEnergyServiceData___includedServices_newList(ptr.Pointer()))
}
