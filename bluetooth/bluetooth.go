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
	"reflect"
	"strings"
	"unsafe"
)

func cGoFreePacked(ptr unsafe.Pointer) { core.NewQByteArrayFromPointer(ptr).DestroyQByteArray() }
func cGoUnpackString(s C.struct_QtBluetooth_PackedString) string {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtBluetooth_PackedString) []byte {
	defer cGoFreePacked(s.ptr)
	if int(s.len) == -1 {
		gs := C.GoString(s.data)
		return *(*[]byte)(unsafe.Pointer(&gs))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}
func unpackStringList(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}
	return strings.Split(s, "¡¦!")
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

func NewQBluetoothFromPointer(ptr unsafe.Pointer) (n *QBluetooth) {
	n = new(QBluetooth)
	n.SetPointer(ptr)
	return
}
func (ptr *QBluetooth) DestroyQBluetooth() {
	if ptr != nil {
		qt.SetFinalizer(ptr, nil)

		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
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

func NewQBluetoothAddressFromPointer(ptr unsafe.Pointer) (n *QBluetoothAddress) {
	n = new(QBluetoothAddress)
	n.SetPointer(ptr)
	return
}
func NewQBluetoothAddress() *QBluetoothAddress {
	tmpValue := NewQBluetoothAddressFromPointer(C.QBluetoothAddress_NewQBluetoothAddress())
	qt.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
	return tmpValue
}

func NewQBluetoothAddress2(address uint64) *QBluetoothAddress {
	tmpValue := NewQBluetoothAddressFromPointer(C.QBluetoothAddress_NewQBluetoothAddress2(C.ulonglong(address)))
	qt.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
	return tmpValue
}

func NewQBluetoothAddress3(address string) *QBluetoothAddress {
	var addressC *C.char
	if address != "" {
		addressC = C.CString(address)
		defer C.free(unsafe.Pointer(addressC))
	}
	tmpValue := NewQBluetoothAddressFromPointer(C.QBluetoothAddress_NewQBluetoothAddress3(C.struct_QtBluetooth_PackedString{data: addressC, len: C.longlong(len(address))}))
	qt.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
	return tmpValue
}

func NewQBluetoothAddress4(other QBluetoothAddress_ITF) *QBluetoothAddress {
	tmpValue := NewQBluetoothAddressFromPointer(C.QBluetoothAddress_NewQBluetoothAddress4(PointerFromQBluetoothAddress(other)))
	qt.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
	return tmpValue
}

func (ptr *QBluetoothAddress) Clear() {
	if ptr.Pointer() != nil {
		C.QBluetoothAddress_Clear(ptr.Pointer())
	}
}

func (ptr *QBluetoothAddress) IsNull() bool {
	if ptr.Pointer() != nil {
		return int8(C.QBluetoothAddress_IsNull(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QBluetoothAddress) ToString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QBluetoothAddress_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothAddress) ToUInt64() uint64 {
	if ptr.Pointer() != nil {
		return uint64(C.QBluetoothAddress_ToUInt64(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothAddress) DestroyQBluetoothAddress() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QBluetoothAddress_DestroyQBluetoothAddress(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
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

func NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr unsafe.Pointer) (n *QBluetoothDeviceDiscoveryAgent) {
	n = new(QBluetoothDeviceDiscoveryAgent)
	n.SetPointer(ptr)
	return
}

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

//go:generate stringer -type=QBluetoothDeviceDiscoveryAgent__DiscoveryMethod
//QBluetoothDeviceDiscoveryAgent::DiscoveryMethod
type QBluetoothDeviceDiscoveryAgent__DiscoveryMethod int64

const (
	QBluetoothDeviceDiscoveryAgent__NoMethod        QBluetoothDeviceDiscoveryAgent__DiscoveryMethod = QBluetoothDeviceDiscoveryAgent__DiscoveryMethod(0x0)
	QBluetoothDeviceDiscoveryAgent__ClassicMethod   QBluetoothDeviceDiscoveryAgent__DiscoveryMethod = QBluetoothDeviceDiscoveryAgent__DiscoveryMethod(0x01)
	QBluetoothDeviceDiscoveryAgent__LowEnergyMethod QBluetoothDeviceDiscoveryAgent__DiscoveryMethod = QBluetoothDeviceDiscoveryAgent__DiscoveryMethod(0x02)
)

func NewQBluetoothDeviceDiscoveryAgent(parent core.QObject_ITF) *QBluetoothDeviceDiscoveryAgent {
	tmpValue := NewQBluetoothDeviceDiscoveryAgentFromPointer(C.QBluetoothDeviceDiscoveryAgent_NewQBluetoothDeviceDiscoveryAgent(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQBluetoothDeviceDiscoveryAgent2(deviceAdapter QBluetoothAddress_ITF, parent core.QObject_ITF) *QBluetoothDeviceDiscoveryAgent {
	tmpValue := NewQBluetoothDeviceDiscoveryAgentFromPointer(C.QBluetoothDeviceDiscoveryAgent_NewQBluetoothDeviceDiscoveryAgent2(PointerFromQBluetoothAddress(deviceAdapter), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQBluetoothDeviceDiscoveryAgent_Canceled
func callbackQBluetoothDeviceDiscoveryAgent_Canceled(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "canceled"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectCanceled(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "canceled") {
			C.QBluetoothDeviceDiscoveryAgent_ConnectCanceled(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "canceled")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "canceled"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "canceled", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "canceled", unsafe.Pointer(&f))
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
		(*(*func(*QBluetoothDeviceInfo))(signal))(NewQBluetoothDeviceInfoFromPointer(info))
	}

}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectDeviceDiscovered(f func(info *QBluetoothDeviceInfo)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "deviceDiscovered") {
			C.QBluetoothDeviceDiscoveryAgent_ConnectDeviceDiscovered(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "deviceDiscovered")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "deviceDiscovered"); signal != nil {
			f := func(info *QBluetoothDeviceInfo) {
				(*(*func(*QBluetoothDeviceInfo))(signal))(info)
				f(info)
			}
			qt.ConnectSignal(ptr.Pointer(), "deviceDiscovered", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "deviceDiscovered", unsafe.Pointer(&f))
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

//export callbackQBluetoothDeviceDiscoveryAgent_DeviceUpdated
func callbackQBluetoothDeviceDiscoveryAgent_DeviceUpdated(ptr unsafe.Pointer, info unsafe.Pointer, updatedFields C.longlong) {
	if signal := qt.GetSignal(ptr, "deviceUpdated"); signal != nil {
		(*(*func(*QBluetoothDeviceInfo, QBluetoothDeviceInfo__Field))(signal))(NewQBluetoothDeviceInfoFromPointer(info), QBluetoothDeviceInfo__Field(updatedFields))
	}

}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectDeviceUpdated(f func(info *QBluetoothDeviceInfo, updatedFields QBluetoothDeviceInfo__Field)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "deviceUpdated") {
			C.QBluetoothDeviceDiscoveryAgent_ConnectDeviceUpdated(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "deviceUpdated")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "deviceUpdated"); signal != nil {
			f := func(info *QBluetoothDeviceInfo, updatedFields QBluetoothDeviceInfo__Field) {
				(*(*func(*QBluetoothDeviceInfo, QBluetoothDeviceInfo__Field))(signal))(info, updatedFields)
				f(info, updatedFields)
			}
			qt.ConnectSignal(ptr.Pointer(), "deviceUpdated", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "deviceUpdated", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectDeviceUpdated() {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_DisconnectDeviceUpdated(ptr.Pointer())
		qt.DisconnectSignal(ptr.Pointer(), "deviceUpdated")
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DeviceUpdated(info QBluetoothDeviceInfo_ITF, updatedFields QBluetoothDeviceInfo__Field) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_DeviceUpdated(ptr.Pointer(), PointerFromQBluetoothDeviceInfo(info), C.longlong(updatedFields))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DiscoveredDevices() []*QBluetoothDeviceInfo {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtBluetooth_PackedList) []*QBluetoothDeviceInfo {
			out := make([]*QBluetoothDeviceInfo, int(l.len))
			tmpList := NewQBluetoothDeviceDiscoveryAgentFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__discoveredDevices_atList(i)
			}
			return out
		}(C.QBluetoothDeviceDiscoveryAgent_DiscoveredDevices(ptr.Pointer()))
	}
	return make([]*QBluetoothDeviceInfo, 0)
}

func (ptr *QBluetoothDeviceDiscoveryAgent) Error() QBluetoothDeviceDiscoveryAgent__Error {
	if ptr.Pointer() != nil {
		return QBluetoothDeviceDiscoveryAgent__Error(C.QBluetoothDeviceDiscoveryAgent_Error(ptr.Pointer()))
	}
	return 0
}

//export callbackQBluetoothDeviceDiscoveryAgent_Error2
func callbackQBluetoothDeviceDiscoveryAgent_Error2(ptr unsafe.Pointer, error C.longlong) {
	if signal := qt.GetSignal(ptr, "error2"); signal != nil {
		(*(*func(QBluetoothDeviceDiscoveryAgent__Error))(signal))(QBluetoothDeviceDiscoveryAgent__Error(error))
	}

}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectError2(f func(error QBluetoothDeviceDiscoveryAgent__Error)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "error2") {
			C.QBluetoothDeviceDiscoveryAgent_ConnectError2(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "error")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "error2"); signal != nil {
			f := func(error QBluetoothDeviceDiscoveryAgent__Error) {
				(*(*func(QBluetoothDeviceDiscoveryAgent__Error))(signal))(error)
				f(error)
			}
			qt.ConnectSignal(ptr.Pointer(), "error2", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "error2", unsafe.Pointer(&f))
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

func (ptr *QBluetoothDeviceDiscoveryAgent) ErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QBluetoothDeviceDiscoveryAgent_ErrorString(ptr.Pointer()))
	}
	return ""
}

//export callbackQBluetoothDeviceDiscoveryAgent_Finished
func callbackQBluetoothDeviceDiscoveryAgent_Finished(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "finished"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectFinished(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "finished") {
			C.QBluetoothDeviceDiscoveryAgent_ConnectFinished(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "finished")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "finished"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "finished", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "finished", unsafe.Pointer(&f))
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

func (ptr *QBluetoothDeviceDiscoveryAgent) InquiryType() QBluetoothDeviceDiscoveryAgent__InquiryType {
	if ptr.Pointer() != nil {
		return QBluetoothDeviceDiscoveryAgent__InquiryType(C.QBluetoothDeviceDiscoveryAgent_InquiryType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothDeviceDiscoveryAgent) IsActive() bool {
	if ptr.Pointer() != nil {
		return int8(C.QBluetoothDeviceDiscoveryAgent_IsActive(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QBluetoothDeviceDiscoveryAgent) LowEnergyDiscoveryTimeout() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QBluetoothDeviceDiscoveryAgent_LowEnergyDiscoveryTimeout(ptr.Pointer())))
	}
	return 0
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
		(*(*func())(signal))()
	} else {
		NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr).StartDefault()
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectStart(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "start"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "start", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "start", unsafe.Pointer(&f))
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
		(*(*func(QBluetoothDeviceDiscoveryAgent__DiscoveryMethod))(signal))(QBluetoothDeviceDiscoveryAgent__DiscoveryMethod(methods))
	} else {
		NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr).Start2Default(QBluetoothDeviceDiscoveryAgent__DiscoveryMethod(methods))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectStart2(f func(methods QBluetoothDeviceDiscoveryAgent__DiscoveryMethod)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "start2"); signal != nil {
			f := func(methods QBluetoothDeviceDiscoveryAgent__DiscoveryMethod) {
				(*(*func(QBluetoothDeviceDiscoveryAgent__DiscoveryMethod))(signal))(methods)
				f(methods)
			}
			qt.ConnectSignal(ptr.Pointer(), "start2", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "start2", unsafe.Pointer(&f))
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
		(*(*func())(signal))()
	} else {
		NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr).StopDefault()
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectStop(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "stop"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "stop", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stop", unsafe.Pointer(&f))
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

func QBluetoothDeviceDiscoveryAgent_SupportedDiscoveryMethods() QBluetoothDeviceDiscoveryAgent__DiscoveryMethod {
	return QBluetoothDeviceDiscoveryAgent__DiscoveryMethod(C.QBluetoothDeviceDiscoveryAgent_QBluetoothDeviceDiscoveryAgent_SupportedDiscoveryMethods())
}

func (ptr *QBluetoothDeviceDiscoveryAgent) SupportedDiscoveryMethods() QBluetoothDeviceDiscoveryAgent__DiscoveryMethod {
	return QBluetoothDeviceDiscoveryAgent__DiscoveryMethod(C.QBluetoothDeviceDiscoveryAgent_QBluetoothDeviceDiscoveryAgent_SupportedDiscoveryMethods())
}

//export callbackQBluetoothDeviceDiscoveryAgent_DestroyQBluetoothDeviceDiscoveryAgent
func callbackQBluetoothDeviceDiscoveryAgent_DestroyQBluetoothDeviceDiscoveryAgent(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QBluetoothDeviceDiscoveryAgent"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr).DestroyQBluetoothDeviceDiscoveryAgentDefault()
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) ConnectDestroyQBluetoothDeviceDiscoveryAgent(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QBluetoothDeviceDiscoveryAgent"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QBluetoothDeviceDiscoveryAgent", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QBluetoothDeviceDiscoveryAgent", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectDestroyQBluetoothDeviceDiscoveryAgent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QBluetoothDeviceDiscoveryAgent")
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DestroyQBluetoothDeviceDiscoveryAgent() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QBluetoothDeviceDiscoveryAgent_DestroyQBluetoothDeviceDiscoveryAgent(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DestroyQBluetoothDeviceDiscoveryAgentDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QBluetoothDeviceDiscoveryAgent_DestroyQBluetoothDeviceDiscoveryAgentDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) __discoveredDevices_atList(i int) *QBluetoothDeviceInfo {
	if ptr.Pointer() != nil {
		tmpValue := NewQBluetoothDeviceInfoFromPointer(C.QBluetoothDeviceDiscoveryAgent___discoveredDevices_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QBluetoothDeviceInfo).DestroyQBluetoothDeviceInfo)
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
	return C.QBluetoothDeviceDiscoveryAgent___discoveredDevices_newList(ptr.Pointer())
}

func (ptr *QBluetoothDeviceDiscoveryAgent) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QBluetoothDeviceDiscoveryAgent___children_atList(ptr.Pointer(), C.int(int32(i))))
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
	return C.QBluetoothDeviceDiscoveryAgent___children_newList(ptr.Pointer())
}

func (ptr *QBluetoothDeviceDiscoveryAgent) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QBluetoothDeviceDiscoveryAgent___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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
	return C.QBluetoothDeviceDiscoveryAgent___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QBluetoothDeviceDiscoveryAgent) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QBluetoothDeviceDiscoveryAgent___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
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
	return C.QBluetoothDeviceDiscoveryAgent___findChildren_newList(ptr.Pointer())
}

func (ptr *QBluetoothDeviceDiscoveryAgent) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QBluetoothDeviceDiscoveryAgent___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
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
	return C.QBluetoothDeviceDiscoveryAgent___findChildren_newList3(ptr.Pointer())
}

//export callbackQBluetoothDeviceDiscoveryAgent_ChildEvent
func callbackQBluetoothDeviceDiscoveryAgent_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
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
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
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
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
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
		(*(*func())(signal))()
	} else {
		NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QBluetoothDeviceDiscoveryAgent_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQBluetoothDeviceDiscoveryAgent_Destroyed
func callbackQBluetoothDeviceDiscoveryAgent_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQBluetoothDeviceDiscoveryAgent_DisconnectNotify
func callbackQBluetoothDeviceDiscoveryAgent_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBluetoothDeviceDiscoveryAgent_Event
func callbackQBluetoothDeviceDiscoveryAgent_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QBluetoothDeviceDiscoveryAgent) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QBluetoothDeviceDiscoveryAgent_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQBluetoothDeviceDiscoveryAgent_EventFilter
func callbackQBluetoothDeviceDiscoveryAgent_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QBluetoothDeviceDiscoveryAgent) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QBluetoothDeviceDiscoveryAgent_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQBluetoothDeviceDiscoveryAgent_MetaObject
func callbackQBluetoothDeviceDiscoveryAgent_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QBluetoothDeviceDiscoveryAgent) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QBluetoothDeviceDiscoveryAgent_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQBluetoothDeviceDiscoveryAgent_ObjectNameChanged
func callbackQBluetoothDeviceDiscoveryAgent_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtBluetooth_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQBluetoothDeviceDiscoveryAgent_TimerEvent
func callbackQBluetoothDeviceDiscoveryAgent_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQBluetoothDeviceDiscoveryAgentFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QBluetoothDeviceDiscoveryAgent) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceDiscoveryAgent_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
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

func NewQBluetoothDeviceInfoFromPointer(ptr unsafe.Pointer) (n *QBluetoothDeviceInfo) {
	n = new(QBluetoothDeviceInfo)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QBluetoothDeviceInfo__MajorDeviceClass
//QBluetoothDeviceInfo::MajorDeviceClass
type QBluetoothDeviceInfo__MajorDeviceClass int64

const (
	QBluetoothDeviceInfo__MiscellaneousDevice QBluetoothDeviceInfo__MajorDeviceClass = QBluetoothDeviceInfo__MajorDeviceClass(0)
	QBluetoothDeviceInfo__ComputerDevice      QBluetoothDeviceInfo__MajorDeviceClass = QBluetoothDeviceInfo__MajorDeviceClass(1)
	QBluetoothDeviceInfo__PhoneDevice         QBluetoothDeviceInfo__MajorDeviceClass = QBluetoothDeviceInfo__MajorDeviceClass(2)
	QBluetoothDeviceInfo__LANAccessDevice     QBluetoothDeviceInfo__MajorDeviceClass = QBluetoothDeviceInfo__MajorDeviceClass(3)
	QBluetoothDeviceInfo__NetworkDevice       QBluetoothDeviceInfo__MajorDeviceClass = QBluetoothDeviceInfo__MajorDeviceClass(3)
	QBluetoothDeviceInfo__AudioVideoDevice    QBluetoothDeviceInfo__MajorDeviceClass = QBluetoothDeviceInfo__MajorDeviceClass(4)
	QBluetoothDeviceInfo__PeripheralDevice    QBluetoothDeviceInfo__MajorDeviceClass = QBluetoothDeviceInfo__MajorDeviceClass(5)
	QBluetoothDeviceInfo__ImagingDevice       QBluetoothDeviceInfo__MajorDeviceClass = QBluetoothDeviceInfo__MajorDeviceClass(6)
	QBluetoothDeviceInfo__WearableDevice      QBluetoothDeviceInfo__MajorDeviceClass = QBluetoothDeviceInfo__MajorDeviceClass(7)
	QBluetoothDeviceInfo__ToyDevice           QBluetoothDeviceInfo__MajorDeviceClass = QBluetoothDeviceInfo__MajorDeviceClass(8)
	QBluetoothDeviceInfo__HealthDevice        QBluetoothDeviceInfo__MajorDeviceClass = QBluetoothDeviceInfo__MajorDeviceClass(9)
	QBluetoothDeviceInfo__UncategorizedDevice QBluetoothDeviceInfo__MajorDeviceClass = QBluetoothDeviceInfo__MajorDeviceClass(31)
)

//go:generate stringer -type=QBluetoothDeviceInfo__MinorMiscellaneousClass
//QBluetoothDeviceInfo::MinorMiscellaneousClass
type QBluetoothDeviceInfo__MinorMiscellaneousClass int64

const (
	QBluetoothDeviceInfo__UncategorizedMiscellaneous QBluetoothDeviceInfo__MinorMiscellaneousClass = QBluetoothDeviceInfo__MinorMiscellaneousClass(0)
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

//go:generate stringer -type=QBluetoothDeviceInfo__DataCompleteness
//QBluetoothDeviceInfo::DataCompleteness
type QBluetoothDeviceInfo__DataCompleteness int64

const (
	QBluetoothDeviceInfo__DataComplete    QBluetoothDeviceInfo__DataCompleteness = QBluetoothDeviceInfo__DataCompleteness(0)
	QBluetoothDeviceInfo__DataIncomplete  QBluetoothDeviceInfo__DataCompleteness = QBluetoothDeviceInfo__DataCompleteness(1)
	QBluetoothDeviceInfo__DataUnavailable QBluetoothDeviceInfo__DataCompleteness = QBluetoothDeviceInfo__DataCompleteness(2)
)

//go:generate stringer -type=QBluetoothDeviceInfo__Field
//QBluetoothDeviceInfo::Field
type QBluetoothDeviceInfo__Field int64

const (
	QBluetoothDeviceInfo__None             QBluetoothDeviceInfo__Field = QBluetoothDeviceInfo__Field(0x0000)
	QBluetoothDeviceInfo__RSSI             QBluetoothDeviceInfo__Field = QBluetoothDeviceInfo__Field(0x0001)
	QBluetoothDeviceInfo__ManufacturerData QBluetoothDeviceInfo__Field = QBluetoothDeviceInfo__Field(0x0002)
	QBluetoothDeviceInfo__All              QBluetoothDeviceInfo__Field = QBluetoothDeviceInfo__Field(0x7fff)
)

//go:generate stringer -type=QBluetoothDeviceInfo__CoreConfiguration
//QBluetoothDeviceInfo::CoreConfiguration
type QBluetoothDeviceInfo__CoreConfiguration int64

const (
	QBluetoothDeviceInfo__UnknownCoreConfiguration              QBluetoothDeviceInfo__CoreConfiguration = QBluetoothDeviceInfo__CoreConfiguration(0x0)
	QBluetoothDeviceInfo__LowEnergyCoreConfiguration            QBluetoothDeviceInfo__CoreConfiguration = QBluetoothDeviceInfo__CoreConfiguration(0x01)
	QBluetoothDeviceInfo__BaseRateCoreConfiguration             QBluetoothDeviceInfo__CoreConfiguration = QBluetoothDeviceInfo__CoreConfiguration(0x02)
	QBluetoothDeviceInfo__BaseRateAndLowEnergyCoreConfiguration QBluetoothDeviceInfo__CoreConfiguration = QBluetoothDeviceInfo__CoreConfiguration(0x03)
)

func NewQBluetoothDeviceInfo() *QBluetoothDeviceInfo {
	tmpValue := NewQBluetoothDeviceInfoFromPointer(C.QBluetoothDeviceInfo_NewQBluetoothDeviceInfo())
	qt.SetFinalizer(tmpValue, (*QBluetoothDeviceInfo).DestroyQBluetoothDeviceInfo)
	return tmpValue
}

func NewQBluetoothDeviceInfo2(address QBluetoothAddress_ITF, name string, classOfDevice uint) *QBluetoothDeviceInfo {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := NewQBluetoothDeviceInfoFromPointer(C.QBluetoothDeviceInfo_NewQBluetoothDeviceInfo2(PointerFromQBluetoothAddress(address), C.struct_QtBluetooth_PackedString{data: nameC, len: C.longlong(len(name))}, C.uint(uint32(classOfDevice))))
	qt.SetFinalizer(tmpValue, (*QBluetoothDeviceInfo).DestroyQBluetoothDeviceInfo)
	return tmpValue
}

func NewQBluetoothDeviceInfo3(uuid QBluetoothUuid_ITF, name string, classOfDevice uint) *QBluetoothDeviceInfo {
	var nameC *C.char
	if name != "" {
		nameC = C.CString(name)
		defer C.free(unsafe.Pointer(nameC))
	}
	tmpValue := NewQBluetoothDeviceInfoFromPointer(C.QBluetoothDeviceInfo_NewQBluetoothDeviceInfo3(PointerFromQBluetoothUuid(uuid), C.struct_QtBluetooth_PackedString{data: nameC, len: C.longlong(len(name))}, C.uint(uint32(classOfDevice))))
	qt.SetFinalizer(tmpValue, (*QBluetoothDeviceInfo).DestroyQBluetoothDeviceInfo)
	return tmpValue
}

func NewQBluetoothDeviceInfo4(other QBluetoothDeviceInfo_ITF) *QBluetoothDeviceInfo {
	tmpValue := NewQBluetoothDeviceInfoFromPointer(C.QBluetoothDeviceInfo_NewQBluetoothDeviceInfo4(PointerFromQBluetoothDeviceInfo(other)))
	qt.SetFinalizer(tmpValue, (*QBluetoothDeviceInfo).DestroyQBluetoothDeviceInfo)
	return tmpValue
}

func (ptr *QBluetoothDeviceInfo) Address() *QBluetoothAddress {
	if ptr.Pointer() != nil {
		tmpValue := NewQBluetoothAddressFromPointer(C.QBluetoothDeviceInfo_Address(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
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
		tmpValue := NewQBluetoothUuidFromPointer(C.QBluetoothDeviceInfo_DeviceUuid(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothDeviceInfo) IsCached() bool {
	if ptr.Pointer() != nil {
		return int8(C.QBluetoothDeviceInfo_IsCached(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QBluetoothDeviceInfo) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QBluetoothDeviceInfo_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QBluetoothDeviceInfo) MajorDeviceClass() QBluetoothDeviceInfo__MajorDeviceClass {
	if ptr.Pointer() != nil {
		return QBluetoothDeviceInfo__MajorDeviceClass(C.QBluetoothDeviceInfo_MajorDeviceClass(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothDeviceInfo) ManufacturerIds() []uint16 {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtBluetooth_PackedList) []uint16 {
			out := make([]uint16, int(l.len))
			tmpList := NewQBluetoothDeviceInfoFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__manufacturerIds_atList(i)
			}
			return out
		}(C.QBluetoothDeviceInfo_ManufacturerIds(ptr.Pointer()))
	}
	return make([]uint16, 0)
}

func (ptr *QBluetoothDeviceInfo) MinorDeviceClass() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QBluetoothDeviceInfo_MinorDeviceClass(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothDeviceInfo) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QBluetoothDeviceInfo_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothDeviceInfo) Rssi() int16 {
	if ptr.Pointer() != nil {
		return int16(C.QBluetoothDeviceInfo_Rssi(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothDeviceInfo) ServiceClasses() QBluetoothDeviceInfo__ServiceClass {
	if ptr.Pointer() != nil {
		return QBluetoothDeviceInfo__ServiceClass(C.QBluetoothDeviceInfo_ServiceClasses(ptr.Pointer()))
	}
	return 0
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

func (ptr *QBluetoothDeviceInfo) SetManufacturerData(manufacturerId uint16, data core.QByteArray_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QBluetoothDeviceInfo_SetManufacturerData(ptr.Pointer(), C.ushort(manufacturerId), core.PointerFromQByteArray(data))) != 0
	}
	return false
}

func (ptr *QBluetoothDeviceInfo) SetRssi(sign int16) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceInfo_SetRssi(ptr.Pointer(), C.short(sign))
	}
}

func (ptr *QBluetoothDeviceInfo) SetServiceUuids2(uuids []*QBluetoothUuid) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceInfo_SetServiceUuids2(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQBluetoothDeviceInfoFromPointer(NewQBluetoothDeviceInfoFromPointer(nil).__setServiceUuids_uuids_newList2())
			for _, v := range uuids {
				tmpList.__setServiceUuids_uuids_setList2(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QBluetoothDeviceInfo) DestroyQBluetoothDeviceInfo() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QBluetoothDeviceInfo_DestroyQBluetoothDeviceInfo(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothDeviceInfo) __manufacturerIds_atList(i int) uint16 {
	if ptr.Pointer() != nil {
		return uint16(C.QBluetoothDeviceInfo___manufacturerIds_atList(ptr.Pointer(), C.int(int32(i))))
	}
	return 0
}

func (ptr *QBluetoothDeviceInfo) __manufacturerIds_setList(i uint16) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceInfo___manufacturerIds_setList(ptr.Pointer(), C.ushort(i))
	}
}

func (ptr *QBluetoothDeviceInfo) __manufacturerIds_newList() unsafe.Pointer {
	return C.QBluetoothDeviceInfo___manufacturerIds_newList(ptr.Pointer())
}

func (ptr *QBluetoothDeviceInfo) __serviceUuids_atList(i int) *QBluetoothUuid {
	if ptr.Pointer() != nil {
		tmpValue := NewQBluetoothUuidFromPointer(C.QBluetoothDeviceInfo___serviceUuids_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
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
	return C.QBluetoothDeviceInfo___serviceUuids_newList(ptr.Pointer())
}

func (ptr *QBluetoothDeviceInfo) __setServiceUuids_uuids_atList(i int) *QBluetoothUuid {
	if ptr.Pointer() != nil {
		tmpValue := NewQBluetoothUuidFromPointer(C.QBluetoothDeviceInfo___setServiceUuids_uuids_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
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
	return C.QBluetoothDeviceInfo___setServiceUuids_uuids_newList(ptr.Pointer())
}

func (ptr *QBluetoothDeviceInfo) __setServiceUuids_uuids_atList2(i int) *QBluetoothUuid {
	if ptr.Pointer() != nil {
		tmpValue := NewQBluetoothUuidFromPointer(C.QBluetoothDeviceInfo___setServiceUuids_uuids_atList2(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothDeviceInfo) __setServiceUuids_uuids_setList2(i QBluetoothUuid_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceInfo___setServiceUuids_uuids_setList2(ptr.Pointer(), PointerFromQBluetoothUuid(i))
	}
}

func (ptr *QBluetoothDeviceInfo) __setServiceUuids_uuids_newList2() unsafe.Pointer {
	return C.QBluetoothDeviceInfo___setServiceUuids_uuids_newList2(ptr.Pointer())
}

func (ptr *QBluetoothDeviceInfo) ____manufacturerData_keyList_atList2(i int) uint16 {
	if ptr.Pointer() != nil {
		return uint16(C.QBluetoothDeviceInfo_____manufacturerData_keyList_atList2(ptr.Pointer(), C.int(int32(i))))
	}
	return 0
}

func (ptr *QBluetoothDeviceInfo) ____manufacturerData_keyList_setList2(i uint16) {
	if ptr.Pointer() != nil {
		C.QBluetoothDeviceInfo_____manufacturerData_keyList_setList2(ptr.Pointer(), C.ushort(i))
	}
}

func (ptr *QBluetoothDeviceInfo) ____manufacturerData_keyList_newList2() unsafe.Pointer {
	return C.QBluetoothDeviceInfo_____manufacturerData_keyList_newList2(ptr.Pointer())
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

func NewQBluetoothHostInfoFromPointer(ptr unsafe.Pointer) (n *QBluetoothHostInfo) {
	n = new(QBluetoothHostInfo)
	n.SetPointer(ptr)
	return
}
func NewQBluetoothHostInfo() *QBluetoothHostInfo {
	tmpValue := NewQBluetoothHostInfoFromPointer(C.QBluetoothHostInfo_NewQBluetoothHostInfo())
	qt.SetFinalizer(tmpValue, (*QBluetoothHostInfo).DestroyQBluetoothHostInfo)
	return tmpValue
}

func NewQBluetoothHostInfo2(other QBluetoothHostInfo_ITF) *QBluetoothHostInfo {
	tmpValue := NewQBluetoothHostInfoFromPointer(C.QBluetoothHostInfo_NewQBluetoothHostInfo2(PointerFromQBluetoothHostInfo(other)))
	qt.SetFinalizer(tmpValue, (*QBluetoothHostInfo).DestroyQBluetoothHostInfo)
	return tmpValue
}

func (ptr *QBluetoothHostInfo) Address() *QBluetoothAddress {
	if ptr.Pointer() != nil {
		tmpValue := NewQBluetoothAddressFromPointer(C.QBluetoothHostInfo_Address(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
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

		qt.SetFinalizer(ptr, nil)
		C.QBluetoothHostInfo_DestroyQBluetoothHostInfo(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
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

func NewQBluetoothLocalDeviceFromPointer(ptr unsafe.Pointer) (n *QBluetoothLocalDevice) {
	n = new(QBluetoothLocalDevice)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QBluetoothLocalDevice__Pairing
//QBluetoothLocalDevice::Pairing
type QBluetoothLocalDevice__Pairing int64

const (
	QBluetoothLocalDevice__Unpaired         QBluetoothLocalDevice__Pairing = QBluetoothLocalDevice__Pairing(0)
	QBluetoothLocalDevice__Paired           QBluetoothLocalDevice__Pairing = QBluetoothLocalDevice__Pairing(1)
	QBluetoothLocalDevice__AuthorizedPaired QBluetoothLocalDevice__Pairing = QBluetoothLocalDevice__Pairing(2)
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

//go:generate stringer -type=QBluetoothLocalDevice__Error
//QBluetoothLocalDevice::Error
type QBluetoothLocalDevice__Error int64

const (
	QBluetoothLocalDevice__NoError      QBluetoothLocalDevice__Error = QBluetoothLocalDevice__Error(0)
	QBluetoothLocalDevice__PairingError QBluetoothLocalDevice__Error = QBluetoothLocalDevice__Error(1)
	QBluetoothLocalDevice__UnknownError QBluetoothLocalDevice__Error = QBluetoothLocalDevice__Error(100)
)

func NewQBluetoothLocalDevice(parent core.QObject_ITF) *QBluetoothLocalDevice {
	tmpValue := NewQBluetoothLocalDeviceFromPointer(C.QBluetoothLocalDevice_NewQBluetoothLocalDevice(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQBluetoothLocalDevice2(address QBluetoothAddress_ITF, parent core.QObject_ITF) *QBluetoothLocalDevice {
	tmpValue := NewQBluetoothLocalDeviceFromPointer(C.QBluetoothLocalDevice_NewQBluetoothLocalDevice2(PointerFromQBluetoothAddress(address), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QBluetoothLocalDevice) Address() *QBluetoothAddress {
	if ptr.Pointer() != nil {
		tmpValue := NewQBluetoothAddressFromPointer(C.QBluetoothLocalDevice_Address(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
		return tmpValue
	}
	return nil
}

func QBluetoothLocalDevice_AllDevices() []*QBluetoothHostInfo {
	return func(l C.struct_QtBluetooth_PackedList) []*QBluetoothHostInfo {
		out := make([]*QBluetoothHostInfo, int(l.len))
		tmpList := NewQBluetoothLocalDeviceFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__allDevices_atList(i)
		}
		return out
	}(C.QBluetoothLocalDevice_QBluetoothLocalDevice_AllDevices())
}

func (ptr *QBluetoothLocalDevice) AllDevices() []*QBluetoothHostInfo {
	return func(l C.struct_QtBluetooth_PackedList) []*QBluetoothHostInfo {
		out := make([]*QBluetoothHostInfo, int(l.len))
		tmpList := NewQBluetoothLocalDeviceFromPointer(l.data)
		for i := 0; i < len(out); i++ {
			out[i] = tmpList.__allDevices_atList(i)
		}
		return out
	}(C.QBluetoothLocalDevice_QBluetoothLocalDevice_AllDevices())
}

func (ptr *QBluetoothLocalDevice) ConnectedDevices() []*QBluetoothAddress {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtBluetooth_PackedList) []*QBluetoothAddress {
			out := make([]*QBluetoothAddress, int(l.len))
			tmpList := NewQBluetoothLocalDeviceFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__connectedDevices_atList(i)
			}
			return out
		}(C.QBluetoothLocalDevice_ConnectedDevices(ptr.Pointer()))
	}
	return make([]*QBluetoothAddress, 0)
}

//export callbackQBluetoothLocalDevice_DeviceConnected
func callbackQBluetoothLocalDevice_DeviceConnected(ptr unsafe.Pointer, address unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "deviceConnected"); signal != nil {
		(*(*func(*QBluetoothAddress))(signal))(NewQBluetoothAddressFromPointer(address))
	}

}

func (ptr *QBluetoothLocalDevice) ConnectDeviceConnected(f func(address *QBluetoothAddress)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "deviceConnected") {
			C.QBluetoothLocalDevice_ConnectDeviceConnected(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "deviceConnected")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "deviceConnected"); signal != nil {
			f := func(address *QBluetoothAddress) {
				(*(*func(*QBluetoothAddress))(signal))(address)
				f(address)
			}
			qt.ConnectSignal(ptr.Pointer(), "deviceConnected", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "deviceConnected", unsafe.Pointer(&f))
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
		(*(*func(*QBluetoothAddress))(signal))(NewQBluetoothAddressFromPointer(address))
	}

}

func (ptr *QBluetoothLocalDevice) ConnectDeviceDisconnected(f func(address *QBluetoothAddress)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "deviceDisconnected") {
			C.QBluetoothLocalDevice_ConnectDeviceDisconnected(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "deviceDisconnected")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "deviceDisconnected"); signal != nil {
			f := func(address *QBluetoothAddress) {
				(*(*func(*QBluetoothAddress))(signal))(address)
				f(address)
			}
			qt.ConnectSignal(ptr.Pointer(), "deviceDisconnected", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "deviceDisconnected", unsafe.Pointer(&f))
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
		(*(*func(QBluetoothLocalDevice__Error))(signal))(QBluetoothLocalDevice__Error(error))
	}

}

func (ptr *QBluetoothLocalDevice) ConnectError(f func(error QBluetoothLocalDevice__Error)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "error") {
			C.QBluetoothLocalDevice_ConnectError(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "error")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "error"); signal != nil {
			f := func(error QBluetoothLocalDevice__Error) {
				(*(*func(QBluetoothLocalDevice__Error))(signal))(error)
				f(error)
			}
			qt.ConnectSignal(ptr.Pointer(), "error", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "error", unsafe.Pointer(&f))
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

func (ptr *QBluetoothLocalDevice) HostMode() QBluetoothLocalDevice__HostMode {
	if ptr.Pointer() != nil {
		return QBluetoothLocalDevice__HostMode(C.QBluetoothLocalDevice_HostMode(ptr.Pointer()))
	}
	return 0
}

//export callbackQBluetoothLocalDevice_HostModeStateChanged
func callbackQBluetoothLocalDevice_HostModeStateChanged(ptr unsafe.Pointer, state C.longlong) {
	if signal := qt.GetSignal(ptr, "hostModeStateChanged"); signal != nil {
		(*(*func(QBluetoothLocalDevice__HostMode))(signal))(QBluetoothLocalDevice__HostMode(state))
	}

}

func (ptr *QBluetoothLocalDevice) ConnectHostModeStateChanged(f func(state QBluetoothLocalDevice__HostMode)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "hostModeStateChanged") {
			C.QBluetoothLocalDevice_ConnectHostModeStateChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "hostModeStateChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "hostModeStateChanged"); signal != nil {
			f := func(state QBluetoothLocalDevice__HostMode) {
				(*(*func(QBluetoothLocalDevice__HostMode))(signal))(state)
				f(state)
			}
			qt.ConnectSignal(ptr.Pointer(), "hostModeStateChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "hostModeStateChanged", unsafe.Pointer(&f))
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

func (ptr *QBluetoothLocalDevice) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QBluetoothLocalDevice_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QBluetoothLocalDevice) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QBluetoothLocalDevice_Name(ptr.Pointer()))
	}
	return ""
}

//export callbackQBluetoothLocalDevice_PairingConfirmation
func callbackQBluetoothLocalDevice_PairingConfirmation(ptr unsafe.Pointer, confirmation C.char) {
	if signal := qt.GetSignal(ptr, "pairingConfirmation"); signal != nil {
		(*(*func(bool))(signal))(int8(confirmation) != 0)
	} else {
		NewQBluetoothLocalDeviceFromPointer(ptr).PairingConfirmationDefault(int8(confirmation) != 0)
	}
}

func (ptr *QBluetoothLocalDevice) ConnectPairingConfirmation(f func(confirmation bool)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "pairingConfirmation"); signal != nil {
			f := func(confirmation bool) {
				(*(*func(bool))(signal))(confirmation)
				f(confirmation)
			}
			qt.ConnectSignal(ptr.Pointer(), "pairingConfirmation", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "pairingConfirmation", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectPairingConfirmation() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "pairingConfirmation")
	}
}

func (ptr *QBluetoothLocalDevice) PairingConfirmation(confirmation bool) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_PairingConfirmation(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(confirmation))))
	}
}

func (ptr *QBluetoothLocalDevice) PairingConfirmationDefault(confirmation bool) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_PairingConfirmationDefault(ptr.Pointer(), C.char(int8(qt.GoBoolToInt(confirmation))))
	}
}

//export callbackQBluetoothLocalDevice_PairingDisplayConfirmation
func callbackQBluetoothLocalDevice_PairingDisplayConfirmation(ptr unsafe.Pointer, address unsafe.Pointer, pin C.struct_QtBluetooth_PackedString) {
	if signal := qt.GetSignal(ptr, "pairingDisplayConfirmation"); signal != nil {
		(*(*func(*QBluetoothAddress, string))(signal))(NewQBluetoothAddressFromPointer(address), cGoUnpackString(pin))
	}

}

func (ptr *QBluetoothLocalDevice) ConnectPairingDisplayConfirmation(f func(address *QBluetoothAddress, pin string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "pairingDisplayConfirmation") {
			C.QBluetoothLocalDevice_ConnectPairingDisplayConfirmation(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "pairingDisplayConfirmation")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "pairingDisplayConfirmation"); signal != nil {
			f := func(address *QBluetoothAddress, pin string) {
				(*(*func(*QBluetoothAddress, string))(signal))(address, pin)
				f(address, pin)
			}
			qt.ConnectSignal(ptr.Pointer(), "pairingDisplayConfirmation", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "pairingDisplayConfirmation", unsafe.Pointer(&f))
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
		(*(*func(*QBluetoothAddress, string))(signal))(NewQBluetoothAddressFromPointer(address), cGoUnpackString(pin))
	}

}

func (ptr *QBluetoothLocalDevice) ConnectPairingDisplayPinCode(f func(address *QBluetoothAddress, pin string)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "pairingDisplayPinCode") {
			C.QBluetoothLocalDevice_ConnectPairingDisplayPinCode(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "pairingDisplayPinCode")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "pairingDisplayPinCode"); signal != nil {
			f := func(address *QBluetoothAddress, pin string) {
				(*(*func(*QBluetoothAddress, string))(signal))(address, pin)
				f(address, pin)
			}
			qt.ConnectSignal(ptr.Pointer(), "pairingDisplayPinCode", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "pairingDisplayPinCode", unsafe.Pointer(&f))
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
		(*(*func(*QBluetoothAddress, QBluetoothLocalDevice__Pairing))(signal))(NewQBluetoothAddressFromPointer(address), QBluetoothLocalDevice__Pairing(pairing))
	}

}

func (ptr *QBluetoothLocalDevice) ConnectPairingFinished(f func(address *QBluetoothAddress, pairing QBluetoothLocalDevice__Pairing)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "pairingFinished") {
			C.QBluetoothLocalDevice_ConnectPairingFinished(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "pairingFinished")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "pairingFinished"); signal != nil {
			f := func(address *QBluetoothAddress, pairing QBluetoothLocalDevice__Pairing) {
				(*(*func(*QBluetoothAddress, QBluetoothLocalDevice__Pairing))(signal))(address, pairing)
				f(address, pairing)
			}
			qt.ConnectSignal(ptr.Pointer(), "pairingFinished", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "pairingFinished", unsafe.Pointer(&f))
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

func (ptr *QBluetoothLocalDevice) PairingStatus(address QBluetoothAddress_ITF) QBluetoothLocalDevice__Pairing {
	if ptr.Pointer() != nil {
		return QBluetoothLocalDevice__Pairing(C.QBluetoothLocalDevice_PairingStatus(ptr.Pointer(), PointerFromQBluetoothAddress(address)))
	}
	return 0
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
		(*(*func())(signal))()
	} else {
		NewQBluetoothLocalDeviceFromPointer(ptr).DestroyQBluetoothLocalDeviceDefault()
	}
}

func (ptr *QBluetoothLocalDevice) ConnectDestroyQBluetoothLocalDevice(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QBluetoothLocalDevice"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QBluetoothLocalDevice", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QBluetoothLocalDevice", unsafe.Pointer(&f))
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

		qt.SetFinalizer(ptr, nil)
		C.QBluetoothLocalDevice_DestroyQBluetoothLocalDevice(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothLocalDevice) DestroyQBluetoothLocalDeviceDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QBluetoothLocalDevice_DestroyQBluetoothLocalDeviceDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothLocalDevice) __allDevices_atList(i int) *QBluetoothHostInfo {
	if ptr.Pointer() != nil {
		tmpValue := NewQBluetoothHostInfoFromPointer(C.QBluetoothLocalDevice___allDevices_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QBluetoothHostInfo).DestroyQBluetoothHostInfo)
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
	return C.QBluetoothLocalDevice___allDevices_newList(ptr.Pointer())
}

func (ptr *QBluetoothLocalDevice) __connectedDevices_atList(i int) *QBluetoothAddress {
	if ptr.Pointer() != nil {
		tmpValue := NewQBluetoothAddressFromPointer(C.QBluetoothLocalDevice___connectedDevices_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
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
	return C.QBluetoothLocalDevice___connectedDevices_newList(ptr.Pointer())
}

func (ptr *QBluetoothLocalDevice) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QBluetoothLocalDevice___children_atList(ptr.Pointer(), C.int(int32(i))))
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
	return C.QBluetoothLocalDevice___children_newList(ptr.Pointer())
}

func (ptr *QBluetoothLocalDevice) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QBluetoothLocalDevice___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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
	return C.QBluetoothLocalDevice___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QBluetoothLocalDevice) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QBluetoothLocalDevice___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
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
	return C.QBluetoothLocalDevice___findChildren_newList(ptr.Pointer())
}

func (ptr *QBluetoothLocalDevice) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QBluetoothLocalDevice___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
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
	return C.QBluetoothLocalDevice___findChildren_newList3(ptr.Pointer())
}

//export callbackQBluetoothLocalDevice_ChildEvent
func callbackQBluetoothLocalDevice_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
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
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
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
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
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
		(*(*func())(signal))()
	} else {
		NewQBluetoothLocalDeviceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QBluetoothLocalDevice) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QBluetoothLocalDevice_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQBluetoothLocalDevice_Destroyed
func callbackQBluetoothLocalDevice_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQBluetoothLocalDevice_DisconnectNotify
func callbackQBluetoothLocalDevice_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBluetoothLocalDeviceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBluetoothLocalDevice) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBluetoothLocalDevice_Event
func callbackQBluetoothLocalDevice_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothLocalDeviceFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QBluetoothLocalDevice) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QBluetoothLocalDevice_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQBluetoothLocalDevice_EventFilter
func callbackQBluetoothLocalDevice_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothLocalDeviceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QBluetoothLocalDevice) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QBluetoothLocalDevice_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQBluetoothLocalDevice_MetaObject
func callbackQBluetoothLocalDevice_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQBluetoothLocalDeviceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QBluetoothLocalDevice) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QBluetoothLocalDevice_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQBluetoothLocalDevice_ObjectNameChanged
func callbackQBluetoothLocalDevice_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtBluetooth_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQBluetoothLocalDevice_TimerEvent
func callbackQBluetoothLocalDevice_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQBluetoothLocalDeviceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QBluetoothLocalDevice) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothLocalDevice_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
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

func NewQBluetoothServerFromPointer(ptr unsafe.Pointer) (n *QBluetoothServer) {
	n = new(QBluetoothServer)
	n.SetPointer(ptr)
	return
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
	tmpValue := NewQBluetoothServerFromPointer(C.QBluetoothServer_NewQBluetoothServer(C.longlong(serverType), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QBluetoothServer) Close() {
	if ptr.Pointer() != nil {
		C.QBluetoothServer_Close(ptr.Pointer())
	}
}

func (ptr *QBluetoothServer) Error() QBluetoothServer__Error {
	if ptr.Pointer() != nil {
		return QBluetoothServer__Error(C.QBluetoothServer_Error(ptr.Pointer()))
	}
	return 0
}

//export callbackQBluetoothServer_Error2
func callbackQBluetoothServer_Error2(ptr unsafe.Pointer, error C.longlong) {
	if signal := qt.GetSignal(ptr, "error2"); signal != nil {
		(*(*func(QBluetoothServer__Error))(signal))(QBluetoothServer__Error(error))
	}

}

func (ptr *QBluetoothServer) ConnectError2(f func(error QBluetoothServer__Error)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "error2") {
			C.QBluetoothServer_ConnectError2(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "error")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "error2"); signal != nil {
			f := func(error QBluetoothServer__Error) {
				(*(*func(QBluetoothServer__Error))(signal))(error)
				f(error)
			}
			qt.ConnectSignal(ptr.Pointer(), "error2", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "error2", unsafe.Pointer(&f))
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

func (ptr *QBluetoothServer) HasPendingConnections() bool {
	if ptr.Pointer() != nil {
		return int8(C.QBluetoothServer_HasPendingConnections(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QBluetoothServer) IsListening() bool {
	if ptr.Pointer() != nil {
		return int8(C.QBluetoothServer_IsListening(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QBluetoothServer) Listen(address QBluetoothAddress_ITF, port uint16) bool {
	if ptr.Pointer() != nil {
		return int8(C.QBluetoothServer_Listen(ptr.Pointer(), PointerFromQBluetoothAddress(address), C.ushort(port))) != 0
	}
	return false
}

func (ptr *QBluetoothServer) Listen2(uuid QBluetoothUuid_ITF, serviceName string) *QBluetoothServiceInfo {
	if ptr.Pointer() != nil {
		var serviceNameC *C.char
		if serviceName != "" {
			serviceNameC = C.CString(serviceName)
			defer C.free(unsafe.Pointer(serviceNameC))
		}
		tmpValue := NewQBluetoothServiceInfoFromPointer(C.QBluetoothServer_Listen2(ptr.Pointer(), PointerFromQBluetoothUuid(uuid), C.struct_QtBluetooth_PackedString{data: serviceNameC, len: C.longlong(len(serviceName))}))
		qt.SetFinalizer(tmpValue, (*QBluetoothServiceInfo).DestroyQBluetoothServiceInfo)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothServer) MaxPendingConnections() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QBluetoothServer_MaxPendingConnections(ptr.Pointer())))
	}
	return 0
}

//export callbackQBluetoothServer_NewConnection
func callbackQBluetoothServer_NewConnection(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "newConnection"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QBluetoothServer) ConnectNewConnection(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "newConnection") {
			C.QBluetoothServer_ConnectNewConnection(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "newConnection")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "newConnection"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "newConnection", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "newConnection", unsafe.Pointer(&f))
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

func (ptr *QBluetoothServer) NextPendingConnection() *QBluetoothSocket {
	if ptr.Pointer() != nil {
		tmpValue := NewQBluetoothSocketFromPointer(C.QBluetoothServer_NextPendingConnection(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothServer) SecurityFlags() QBluetooth__Security {
	if ptr.Pointer() != nil {
		return QBluetooth__Security(C.QBluetoothServer_SecurityFlags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothServer) ServerAddress() *QBluetoothAddress {
	if ptr.Pointer() != nil {
		tmpValue := NewQBluetoothAddressFromPointer(C.QBluetoothServer_ServerAddress(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothServer) ServerPort() uint16 {
	if ptr.Pointer() != nil {
		return uint16(C.QBluetoothServer_ServerPort(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothServer) ServerType() QBluetoothServiceInfo__Protocol {
	if ptr.Pointer() != nil {
		return QBluetoothServiceInfo__Protocol(C.QBluetoothServer_ServerType(ptr.Pointer()))
	}
	return 0
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

//export callbackQBluetoothServer_DestroyQBluetoothServer
func callbackQBluetoothServer_DestroyQBluetoothServer(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QBluetoothServer"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQBluetoothServerFromPointer(ptr).DestroyQBluetoothServerDefault()
	}
}

func (ptr *QBluetoothServer) ConnectDestroyQBluetoothServer(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QBluetoothServer"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QBluetoothServer", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QBluetoothServer", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBluetoothServer) DisconnectDestroyQBluetoothServer() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QBluetoothServer")
	}
}

func (ptr *QBluetoothServer) DestroyQBluetoothServer() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QBluetoothServer_DestroyQBluetoothServer(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothServer) DestroyQBluetoothServerDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QBluetoothServer_DestroyQBluetoothServerDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothServer) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QBluetoothServer___children_atList(ptr.Pointer(), C.int(int32(i))))
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
	return C.QBluetoothServer___children_newList(ptr.Pointer())
}

func (ptr *QBluetoothServer) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QBluetoothServer___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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
	return C.QBluetoothServer___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QBluetoothServer) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QBluetoothServer___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
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
	return C.QBluetoothServer___findChildren_newList(ptr.Pointer())
}

func (ptr *QBluetoothServer) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QBluetoothServer___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
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
	return C.QBluetoothServer___findChildren_newList3(ptr.Pointer())
}

//export callbackQBluetoothServer_ChildEvent
func callbackQBluetoothServer_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
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
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
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
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
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
		(*(*func())(signal))()
	} else {
		NewQBluetoothServerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QBluetoothServer) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QBluetoothServer_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQBluetoothServer_Destroyed
func callbackQBluetoothServer_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQBluetoothServer_DisconnectNotify
func callbackQBluetoothServer_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBluetoothServerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBluetoothServer) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServer_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBluetoothServer_Event
func callbackQBluetoothServer_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothServerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QBluetoothServer) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QBluetoothServer_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQBluetoothServer_EventFilter
func callbackQBluetoothServer_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothServerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QBluetoothServer) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QBluetoothServer_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQBluetoothServer_MetaObject
func callbackQBluetoothServer_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQBluetoothServerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QBluetoothServer) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QBluetoothServer_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQBluetoothServer_ObjectNameChanged
func callbackQBluetoothServer_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtBluetooth_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQBluetoothServer_TimerEvent
func callbackQBluetoothServer_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQBluetoothServerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QBluetoothServer) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServer_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
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

func NewQBluetoothServiceDiscoveryAgentFromPointer(ptr unsafe.Pointer) (n *QBluetoothServiceDiscoveryAgent) {
	n = new(QBluetoothServiceDiscoveryAgent)
	n.SetPointer(ptr)
	return
}

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

//go:generate stringer -type=QBluetoothServiceDiscoveryAgent__DiscoveryMode
//QBluetoothServiceDiscoveryAgent::DiscoveryMode
type QBluetoothServiceDiscoveryAgent__DiscoveryMode int64

const (
	QBluetoothServiceDiscoveryAgent__MinimalDiscovery QBluetoothServiceDiscoveryAgent__DiscoveryMode = QBluetoothServiceDiscoveryAgent__DiscoveryMode(0)
	QBluetoothServiceDiscoveryAgent__FullDiscovery    QBluetoothServiceDiscoveryAgent__DiscoveryMode = QBluetoothServiceDiscoveryAgent__DiscoveryMode(1)
)

func NewQBluetoothServiceDiscoveryAgent(parent core.QObject_ITF) *QBluetoothServiceDiscoveryAgent {
	tmpValue := NewQBluetoothServiceDiscoveryAgentFromPointer(C.QBluetoothServiceDiscoveryAgent_NewQBluetoothServiceDiscoveryAgent(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQBluetoothServiceDiscoveryAgent2(deviceAdapter QBluetoothAddress_ITF, parent core.QObject_ITF) *QBluetoothServiceDiscoveryAgent {
	tmpValue := NewQBluetoothServiceDiscoveryAgentFromPointer(C.QBluetoothServiceDiscoveryAgent_NewQBluetoothServiceDiscoveryAgent2(PointerFromQBluetoothAddress(deviceAdapter), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQBluetoothServiceDiscoveryAgent_Canceled
func callbackQBluetoothServiceDiscoveryAgent_Canceled(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "canceled"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectCanceled(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "canceled") {
			C.QBluetoothServiceDiscoveryAgent_ConnectCanceled(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "canceled")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "canceled"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "canceled", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "canceled", unsafe.Pointer(&f))
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
		(*(*func())(signal))()
	} else {
		NewQBluetoothServiceDiscoveryAgentFromPointer(ptr).ClearDefault()
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectClear(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "clear"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "clear", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "clear", unsafe.Pointer(&f))
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

func (ptr *QBluetoothServiceDiscoveryAgent) DiscoveredServices() []*QBluetoothServiceInfo {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtBluetooth_PackedList) []*QBluetoothServiceInfo {
			out := make([]*QBluetoothServiceInfo, int(l.len))
			tmpList := NewQBluetoothServiceDiscoveryAgentFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__discoveredServices_atList(i)
			}
			return out
		}(C.QBluetoothServiceDiscoveryAgent_DiscoveredServices(ptr.Pointer()))
	}
	return make([]*QBluetoothServiceInfo, 0)
}

func (ptr *QBluetoothServiceDiscoveryAgent) Error() QBluetoothServiceDiscoveryAgent__Error {
	if ptr.Pointer() != nil {
		return QBluetoothServiceDiscoveryAgent__Error(C.QBluetoothServiceDiscoveryAgent_Error(ptr.Pointer()))
	}
	return 0
}

//export callbackQBluetoothServiceDiscoveryAgent_Error2
func callbackQBluetoothServiceDiscoveryAgent_Error2(ptr unsafe.Pointer, error C.longlong) {
	if signal := qt.GetSignal(ptr, "error2"); signal != nil {
		(*(*func(QBluetoothServiceDiscoveryAgent__Error))(signal))(QBluetoothServiceDiscoveryAgent__Error(error))
	}

}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectError2(f func(error QBluetoothServiceDiscoveryAgent__Error)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "error2") {
			C.QBluetoothServiceDiscoveryAgent_ConnectError2(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "error")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "error2"); signal != nil {
			f := func(error QBluetoothServiceDiscoveryAgent__Error) {
				(*(*func(QBluetoothServiceDiscoveryAgent__Error))(signal))(error)
				f(error)
			}
			qt.ConnectSignal(ptr.Pointer(), "error2", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "error2", unsafe.Pointer(&f))
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

func (ptr *QBluetoothServiceDiscoveryAgent) ErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QBluetoothServiceDiscoveryAgent_ErrorString(ptr.Pointer()))
	}
	return ""
}

//export callbackQBluetoothServiceDiscoveryAgent_Finished
func callbackQBluetoothServiceDiscoveryAgent_Finished(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "finished"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectFinished(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "finished") {
			C.QBluetoothServiceDiscoveryAgent_ConnectFinished(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "finished")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "finished"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "finished", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "finished", unsafe.Pointer(&f))
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

func (ptr *QBluetoothServiceDiscoveryAgent) IsActive() bool {
	if ptr.Pointer() != nil {
		return int8(C.QBluetoothServiceDiscoveryAgent_IsActive(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QBluetoothServiceDiscoveryAgent) RemoteAddress() *QBluetoothAddress {
	if ptr.Pointer() != nil {
		tmpValue := NewQBluetoothAddressFromPointer(C.QBluetoothServiceDiscoveryAgent_RemoteAddress(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
		return tmpValue
	}
	return nil
}

//export callbackQBluetoothServiceDiscoveryAgent_ServiceDiscovered
func callbackQBluetoothServiceDiscoveryAgent_ServiceDiscovered(ptr unsafe.Pointer, info unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "serviceDiscovered"); signal != nil {
		(*(*func(*QBluetoothServiceInfo))(signal))(NewQBluetoothServiceInfoFromPointer(info))
	}

}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectServiceDiscovered(f func(info *QBluetoothServiceInfo)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "serviceDiscovered") {
			C.QBluetoothServiceDiscoveryAgent_ConnectServiceDiscovered(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "serviceDiscovered")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "serviceDiscovered"); signal != nil {
			f := func(info *QBluetoothServiceInfo) {
				(*(*func(*QBluetoothServiceInfo))(signal))(info)
				f(info)
			}
			qt.ConnectSignal(ptr.Pointer(), "serviceDiscovered", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "serviceDiscovered", unsafe.Pointer(&f))
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

func (ptr *QBluetoothServiceDiscoveryAgent) SetRemoteAddress(address QBluetoothAddress_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QBluetoothServiceDiscoveryAgent_SetRemoteAddress(ptr.Pointer(), PointerFromQBluetoothAddress(address))) != 0
	}
	return false
}

func (ptr *QBluetoothServiceDiscoveryAgent) SetUuidFilter(uuids []*QBluetoothUuid) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_SetUuidFilter(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQBluetoothServiceDiscoveryAgentFromPointer(NewQBluetoothServiceDiscoveryAgentFromPointer(nil).__setUuidFilter_uuids_newList())
			for _, v := range uuids {
				tmpList.__setUuidFilter_uuids_setList(v)
			}
			return tmpList.Pointer()
		}())
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) SetUuidFilter2(uuid QBluetoothUuid_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_SetUuidFilter2(ptr.Pointer(), PointerFromQBluetoothUuid(uuid))
	}
}

//export callbackQBluetoothServiceDiscoveryAgent_Start
func callbackQBluetoothServiceDiscoveryAgent_Start(ptr unsafe.Pointer, mode C.longlong) {
	if signal := qt.GetSignal(ptr, "start"); signal != nil {
		(*(*func(QBluetoothServiceDiscoveryAgent__DiscoveryMode))(signal))(QBluetoothServiceDiscoveryAgent__DiscoveryMode(mode))
	} else {
		NewQBluetoothServiceDiscoveryAgentFromPointer(ptr).StartDefault(QBluetoothServiceDiscoveryAgent__DiscoveryMode(mode))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectStart(f func(mode QBluetoothServiceDiscoveryAgent__DiscoveryMode)) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "start"); signal != nil {
			f := func(mode QBluetoothServiceDiscoveryAgent__DiscoveryMode) {
				(*(*func(QBluetoothServiceDiscoveryAgent__DiscoveryMode))(signal))(mode)
				f(mode)
			}
			qt.ConnectSignal(ptr.Pointer(), "start", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "start", unsafe.Pointer(&f))
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
		(*(*func())(signal))()
	} else {
		NewQBluetoothServiceDiscoveryAgentFromPointer(ptr).StopDefault()
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectStop(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "stop"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "stop", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stop", unsafe.Pointer(&f))
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

func (ptr *QBluetoothServiceDiscoveryAgent) UuidFilter() []*QBluetoothUuid {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtBluetooth_PackedList) []*QBluetoothUuid {
			out := make([]*QBluetoothUuid, int(l.len))
			tmpList := NewQBluetoothServiceDiscoveryAgentFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__uuidFilter_atList(i)
			}
			return out
		}(C.QBluetoothServiceDiscoveryAgent_UuidFilter(ptr.Pointer()))
	}
	return make([]*QBluetoothUuid, 0)
}

//export callbackQBluetoothServiceDiscoveryAgent_DestroyQBluetoothServiceDiscoveryAgent
func callbackQBluetoothServiceDiscoveryAgent_DestroyQBluetoothServiceDiscoveryAgent(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QBluetoothServiceDiscoveryAgent"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQBluetoothServiceDiscoveryAgentFromPointer(ptr).DestroyQBluetoothServiceDiscoveryAgentDefault()
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) ConnectDestroyQBluetoothServiceDiscoveryAgent(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QBluetoothServiceDiscoveryAgent"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QBluetoothServiceDiscoveryAgent", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QBluetoothServiceDiscoveryAgent", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectDestroyQBluetoothServiceDiscoveryAgent() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QBluetoothServiceDiscoveryAgent")
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DestroyQBluetoothServiceDiscoveryAgent() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QBluetoothServiceDiscoveryAgent_DestroyQBluetoothServiceDiscoveryAgent(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DestroyQBluetoothServiceDiscoveryAgentDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QBluetoothServiceDiscoveryAgent_DestroyQBluetoothServiceDiscoveryAgentDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) __discoveredServices_atList(i int) *QBluetoothServiceInfo {
	if ptr.Pointer() != nil {
		tmpValue := NewQBluetoothServiceInfoFromPointer(C.QBluetoothServiceDiscoveryAgent___discoveredServices_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QBluetoothServiceInfo).DestroyQBluetoothServiceInfo)
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
	return C.QBluetoothServiceDiscoveryAgent___discoveredServices_newList(ptr.Pointer())
}

func (ptr *QBluetoothServiceDiscoveryAgent) __setUuidFilter_uuids_atList(i int) *QBluetoothUuid {
	if ptr.Pointer() != nil {
		tmpValue := NewQBluetoothUuidFromPointer(C.QBluetoothServiceDiscoveryAgent___setUuidFilter_uuids_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
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
	return C.QBluetoothServiceDiscoveryAgent___setUuidFilter_uuids_newList(ptr.Pointer())
}

func (ptr *QBluetoothServiceDiscoveryAgent) __uuidFilter_atList(i int) *QBluetoothUuid {
	if ptr.Pointer() != nil {
		tmpValue := NewQBluetoothUuidFromPointer(C.QBluetoothServiceDiscoveryAgent___uuidFilter_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
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
	return C.QBluetoothServiceDiscoveryAgent___uuidFilter_newList(ptr.Pointer())
}

func (ptr *QBluetoothServiceDiscoveryAgent) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QBluetoothServiceDiscoveryAgent___children_atList(ptr.Pointer(), C.int(int32(i))))
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
	return C.QBluetoothServiceDiscoveryAgent___children_newList(ptr.Pointer())
}

func (ptr *QBluetoothServiceDiscoveryAgent) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QBluetoothServiceDiscoveryAgent___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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
	return C.QBluetoothServiceDiscoveryAgent___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QBluetoothServiceDiscoveryAgent) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QBluetoothServiceDiscoveryAgent___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
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
	return C.QBluetoothServiceDiscoveryAgent___findChildren_newList(ptr.Pointer())
}

func (ptr *QBluetoothServiceDiscoveryAgent) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QBluetoothServiceDiscoveryAgent___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
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
	return C.QBluetoothServiceDiscoveryAgent___findChildren_newList3(ptr.Pointer())
}

//export callbackQBluetoothServiceDiscoveryAgent_ChildEvent
func callbackQBluetoothServiceDiscoveryAgent_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
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
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
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
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
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
		(*(*func())(signal))()
	} else {
		NewQBluetoothServiceDiscoveryAgentFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QBluetoothServiceDiscoveryAgent_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQBluetoothServiceDiscoveryAgent_Destroyed
func callbackQBluetoothServiceDiscoveryAgent_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQBluetoothServiceDiscoveryAgent_DisconnectNotify
func callbackQBluetoothServiceDiscoveryAgent_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBluetoothServiceDiscoveryAgentFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBluetoothServiceDiscoveryAgent_Event
func callbackQBluetoothServiceDiscoveryAgent_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothServiceDiscoveryAgentFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QBluetoothServiceDiscoveryAgent) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QBluetoothServiceDiscoveryAgent_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQBluetoothServiceDiscoveryAgent_EventFilter
func callbackQBluetoothServiceDiscoveryAgent_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothServiceDiscoveryAgentFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QBluetoothServiceDiscoveryAgent) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QBluetoothServiceDiscoveryAgent_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQBluetoothServiceDiscoveryAgent_MetaObject
func callbackQBluetoothServiceDiscoveryAgent_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQBluetoothServiceDiscoveryAgentFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QBluetoothServiceDiscoveryAgent) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QBluetoothServiceDiscoveryAgent_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQBluetoothServiceDiscoveryAgent_ObjectNameChanged
func callbackQBluetoothServiceDiscoveryAgent_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtBluetooth_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQBluetoothServiceDiscoveryAgent_TimerEvent
func callbackQBluetoothServiceDiscoveryAgent_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQBluetoothServiceDiscoveryAgentFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QBluetoothServiceDiscoveryAgent) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceDiscoveryAgent_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
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

func NewQBluetoothServiceInfoFromPointer(ptr unsafe.Pointer) (n *QBluetoothServiceInfo) {
	n = new(QBluetoothServiceInfo)
	n.SetPointer(ptr)
	return
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
	tmpValue := NewQBluetoothServiceInfoFromPointer(C.QBluetoothServiceInfo_NewQBluetoothServiceInfo())
	qt.SetFinalizer(tmpValue, (*QBluetoothServiceInfo).DestroyQBluetoothServiceInfo)
	return tmpValue
}

func NewQBluetoothServiceInfo2(other QBluetoothServiceInfo_ITF) *QBluetoothServiceInfo {
	tmpValue := NewQBluetoothServiceInfoFromPointer(C.QBluetoothServiceInfo_NewQBluetoothServiceInfo2(PointerFromQBluetoothServiceInfo(other)))
	qt.SetFinalizer(tmpValue, (*QBluetoothServiceInfo).DestroyQBluetoothServiceInfo)
	return tmpValue
}

func (ptr *QBluetoothServiceInfo) Attribute(attributeId uint16) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QBluetoothServiceInfo_Attribute(ptr.Pointer(), C.ushort(attributeId)))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothServiceInfo) Contains(attributeId uint16) bool {
	if ptr.Pointer() != nil {
		return int8(C.QBluetoothServiceInfo_Contains(ptr.Pointer(), C.ushort(attributeId))) != 0
	}
	return false
}

func (ptr *QBluetoothServiceInfo) Device() *QBluetoothDeviceInfo {
	if ptr.Pointer() != nil {
		tmpValue := NewQBluetoothDeviceInfoFromPointer(C.QBluetoothServiceInfo_Device(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QBluetoothDeviceInfo).DestroyQBluetoothDeviceInfo)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothServiceInfo) IsComplete() bool {
	if ptr.Pointer() != nil {
		return int8(C.QBluetoothServiceInfo_IsComplete(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QBluetoothServiceInfo) IsRegistered() bool {
	if ptr.Pointer() != nil {
		return int8(C.QBluetoothServiceInfo_IsRegistered(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QBluetoothServiceInfo) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QBluetoothServiceInfo_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QBluetoothServiceInfo) ProtocolServiceMultiplexer() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QBluetoothServiceInfo_ProtocolServiceMultiplexer(ptr.Pointer())))
	}
	return 0
}

func (ptr *QBluetoothServiceInfo) RegisterService(localAdapter QBluetoothAddress_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QBluetoothServiceInfo_RegisterService(ptr.Pointer(), PointerFromQBluetoothAddress(localAdapter))) != 0
	}
	return false
}

func (ptr *QBluetoothServiceInfo) RemoveAttribute(attributeId uint16) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceInfo_RemoveAttribute(ptr.Pointer(), C.ushort(attributeId))
	}
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

func (ptr *QBluetoothServiceInfo) ServiceClassUuids() []*QBluetoothUuid {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtBluetooth_PackedList) []*QBluetoothUuid {
			out := make([]*QBluetoothUuid, int(l.len))
			tmpList := NewQBluetoothServiceInfoFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__serviceClassUuids_atList(i)
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

func (ptr *QBluetoothServiceInfo) ServiceUuid() *QBluetoothUuid {
	if ptr.Pointer() != nil {
		tmpValue := NewQBluetoothUuidFromPointer(C.QBluetoothServiceInfo_ServiceUuid(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothServiceInfo) SetAttribute(attributeId uint16, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceInfo_SetAttribute(ptr.Pointer(), C.ushort(attributeId), core.PointerFromQVariant(value))
	}
}

func (ptr *QBluetoothServiceInfo) SetAttribute2(attributeId uint16, value QBluetoothUuid_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothServiceInfo_SetAttribute2(ptr.Pointer(), C.ushort(attributeId), PointerFromQBluetoothUuid(value))
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

func (ptr *QBluetoothServiceInfo) SocketProtocol() QBluetoothServiceInfo__Protocol {
	if ptr.Pointer() != nil {
		return QBluetoothServiceInfo__Protocol(C.QBluetoothServiceInfo_SocketProtocol(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothServiceInfo) UnregisterService() bool {
	if ptr.Pointer() != nil {
		return int8(C.QBluetoothServiceInfo_UnregisterService(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QBluetoothServiceInfo) DestroyQBluetoothServiceInfo() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QBluetoothServiceInfo_DestroyQBluetoothServiceInfo(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothServiceInfo) __serviceClassUuids_atList(i int) *QBluetoothUuid {
	if ptr.Pointer() != nil {
		tmpValue := NewQBluetoothUuidFromPointer(C.QBluetoothServiceInfo___serviceClassUuids_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
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
	return C.QBluetoothServiceInfo___serviceClassUuids_newList(ptr.Pointer())
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

func NewQBluetoothSocketFromPointer(ptr unsafe.Pointer) (n *QBluetoothSocket) {
	n = new(QBluetoothSocket)
	n.SetPointer(ptr)
	return
}

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

//go:generate stringer -type=QBluetoothSocket__SocketError
//QBluetoothSocket::SocketError
type QBluetoothSocket__SocketError int64

const (
	QBluetoothSocket__NoSocketError            QBluetoothSocket__SocketError = QBluetoothSocket__SocketError(-2)
	QBluetoothSocket__UnknownSocketError       QBluetoothSocket__SocketError = QBluetoothSocket__SocketError(network.QAbstractSocket__UnknownSocketError)
	QBluetoothSocket__RemoteHostClosedError    QBluetoothSocket__SocketError = QBluetoothSocket__SocketError(network.QAbstractSocket__RemoteHostClosedError)
	QBluetoothSocket__HostNotFoundError        QBluetoothSocket__SocketError = QBluetoothSocket__SocketError(network.QAbstractSocket__HostNotFoundError)
	QBluetoothSocket__ServiceNotFoundError     QBluetoothSocket__SocketError = QBluetoothSocket__SocketError(network.QAbstractSocket__SocketAddressNotAvailableError)
	QBluetoothSocket__NetworkError             QBluetoothSocket__SocketError = QBluetoothSocket__SocketError(network.QAbstractSocket__NetworkError)
	QBluetoothSocket__UnsupportedProtocolError QBluetoothSocket__SocketError = QBluetoothSocket__SocketError(8)
	QBluetoothSocket__OperationError           QBluetoothSocket__SocketError = QBluetoothSocket__SocketError(network.QAbstractSocket__OperationError)
)

func NewQBluetoothSocket(socketType QBluetoothServiceInfo__Protocol, parent core.QObject_ITF) *QBluetoothSocket {
	tmpValue := NewQBluetoothSocketFromPointer(C.QBluetoothSocket_NewQBluetoothSocket(C.longlong(socketType), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func NewQBluetoothSocket2(parent core.QObject_ITF) *QBluetoothSocket {
	tmpValue := NewQBluetoothSocketFromPointer(C.QBluetoothSocket_NewQBluetoothSocket2(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QBluetoothSocket) Abort() {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_Abort(ptr.Pointer())
	}
}

//export callbackQBluetoothSocket_BytesAvailable
func callbackQBluetoothSocket_BytesAvailable(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "bytesAvailable"); signal != nil {
		return C.longlong((*(*func() int64)(signal))())
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
		return C.longlong((*(*func() int64)(signal))())
	}

	return C.longlong(NewQBluetoothSocketFromPointer(ptr).BytesToWriteDefault())
}

func (ptr *QBluetoothSocket) BytesToWriteDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QBluetoothSocket_BytesToWriteDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQBluetoothSocket_CanReadLine
func callbackQBluetoothSocket_CanReadLine(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "canReadLine"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothSocketFromPointer(ptr).CanReadLineDefault())))
}

func (ptr *QBluetoothSocket) CanReadLineDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QBluetoothSocket_CanReadLineDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQBluetoothSocket_Close
func callbackQBluetoothSocket_Close(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "close"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQBluetoothSocketFromPointer(ptr).CloseDefault()
	}
}

func (ptr *QBluetoothSocket) CloseDefault() {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_CloseDefault(ptr.Pointer())
	}
}

func (ptr *QBluetoothSocket) ConnectToService(service QBluetoothServiceInfo_ITF, openMode core.QIODevice__OpenModeFlag) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_ConnectToService(ptr.Pointer(), PointerFromQBluetoothServiceInfo(service), C.longlong(openMode))
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

//export callbackQBluetoothSocket_Connected
func callbackQBluetoothSocket_Connected(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "connected"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QBluetoothSocket) ConnectConnected(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "connected") {
			C.QBluetoothSocket_ConnectConnected(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "connected")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "connected"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "connected", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "connected", unsafe.Pointer(&f))
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

func (ptr *QBluetoothSocket) DisconnectFromService() {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DisconnectFromService(ptr.Pointer())
	}
}

//export callbackQBluetoothSocket_Disconnected
func callbackQBluetoothSocket_Disconnected(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnected"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QBluetoothSocket) ConnectDisconnected(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "disconnected") {
			C.QBluetoothSocket_ConnectDisconnected(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "disconnected")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "disconnected"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "disconnected", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "disconnected", unsafe.Pointer(&f))
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

func (ptr *QBluetoothSocket) DoDeviceDiscovery(service QBluetoothServiceInfo_ITF, openMode core.QIODevice__OpenModeFlag) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DoDeviceDiscovery(ptr.Pointer(), PointerFromQBluetoothServiceInfo(service), C.longlong(openMode))
	}
}

func (ptr *QBluetoothSocket) Error() QBluetoothSocket__SocketError {
	if ptr.Pointer() != nil {
		return QBluetoothSocket__SocketError(C.QBluetoothSocket_Error(ptr.Pointer()))
	}
	return 0
}

//export callbackQBluetoothSocket_Error2
func callbackQBluetoothSocket_Error2(ptr unsafe.Pointer, error C.longlong) {
	if signal := qt.GetSignal(ptr, "error2"); signal != nil {
		(*(*func(QBluetoothSocket__SocketError))(signal))(QBluetoothSocket__SocketError(error))
	}

}

func (ptr *QBluetoothSocket) ConnectError2(f func(error QBluetoothSocket__SocketError)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "error2") {
			C.QBluetoothSocket_ConnectError2(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "error")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "error2"); signal != nil {
			f := func(error QBluetoothSocket__SocketError) {
				(*(*func(QBluetoothSocket__SocketError))(signal))(error)
				f(error)
			}
			qt.ConnectSignal(ptr.Pointer(), "error2", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "error2", unsafe.Pointer(&f))
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

//export callbackQBluetoothSocket_IsSequential
func callbackQBluetoothSocket_IsSequential(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isSequential"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothSocketFromPointer(ptr).IsSequentialDefault())))
}

func (ptr *QBluetoothSocket) IsSequentialDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QBluetoothSocket_IsSequentialDefault(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QBluetoothSocket) LocalAddress() *QBluetoothAddress {
	if ptr.Pointer() != nil {
		tmpValue := NewQBluetoothAddressFromPointer(C.QBluetoothSocket_LocalAddress(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothSocket) LocalName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QBluetoothSocket_LocalName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothSocket) LocalPort() uint16 {
	if ptr.Pointer() != nil {
		return uint16(C.QBluetoothSocket_LocalPort(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothSocket) PeerAddress() *QBluetoothAddress {
	if ptr.Pointer() != nil {
		tmpValue := NewQBluetoothAddressFromPointer(C.QBluetoothSocket_PeerAddress(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothSocket) PeerName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QBluetoothSocket_PeerName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothSocket) PeerPort() uint16 {
	if ptr.Pointer() != nil {
		return uint16(C.QBluetoothSocket_PeerPort(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothSocket) PreferredSecurityFlags() QBluetooth__Security {
	if ptr.Pointer() != nil {
		return QBluetooth__Security(C.QBluetoothSocket_PreferredSecurityFlags(ptr.Pointer()))
	}
	return 0
}

//export callbackQBluetoothSocket_ReadData
func callbackQBluetoothSocket_ReadData(ptr unsafe.Pointer, data C.struct_QtBluetooth_PackedString, maxSize C.longlong) C.longlong {
	if signal := qt.GetSignal(ptr, "readData"); signal != nil {
		retS := cGoUnpackString(data)
		ret := C.longlong((*(*func(*string, int64) int64)(signal))(&retS, int64(maxSize)))
		if ret > 0 {
			C.memcpy(unsafe.Pointer(data.data), unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&retS)).Data), C.size_t(ret))
		}
		return ret
	}
	retS := cGoUnpackString(data)
	ret := C.longlong(NewQBluetoothSocketFromPointer(ptr).ReadDataDefault(&retS, int64(maxSize)))
	if ret > 0 {
		C.memcpy(unsafe.Pointer(data.data), unsafe.Pointer((*reflect.StringHeader)(unsafe.Pointer(&retS)).Data), C.size_t(ret))
	}
	return ret
}

func (ptr *QBluetoothSocket) ConnectReadData(f func(data *string, maxSize int64) int64) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "readData"); signal != nil {
			f := func(data *string, maxSize int64) int64 {
				(*(*func(*string, int64) int64)(signal))(data, maxSize)
				return f(data, maxSize)
			}
			qt.ConnectSignal(ptr.Pointer(), "readData", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "readData", unsafe.Pointer(&f))
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
		dataC := C.CString(strings.Repeat("0", int(maxSize)))
		defer C.free(unsafe.Pointer(dataC))
		ret := int64(C.QBluetoothSocket_ReadData(ptr.Pointer(), dataC, C.longlong(maxSize)))
		if ret > 0 {
			*data = C.GoStringN(dataC, C.int(ret))
		}
		return ret
	}
	return 0
}

func (ptr *QBluetoothSocket) ReadDataDefault(data *string, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		dataC := C.CString(strings.Repeat("0", int(maxSize)))
		defer C.free(unsafe.Pointer(dataC))
		ret := int64(C.QBluetoothSocket_ReadDataDefault(ptr.Pointer(), dataC, C.longlong(maxSize)))
		if ret > 0 {
			*data = C.GoStringN(dataC, C.int(ret))
		}
		return ret
	}
	return 0
}

func (ptr *QBluetoothSocket) SetPreferredSecurityFlags(flags QBluetooth__Security) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_SetPreferredSecurityFlags(ptr.Pointer(), C.longlong(flags))
	}
}

func (ptr *QBluetoothSocket) SetSocketDescriptor(socketDescriptor int, socketType QBluetoothServiceInfo__Protocol, socketState QBluetoothSocket__SocketState, openMode core.QIODevice__OpenModeFlag) bool {
	if ptr.Pointer() != nil {
		return int8(C.QBluetoothSocket_SetSocketDescriptor(ptr.Pointer(), C.int(int32(socketDescriptor)), C.longlong(socketType), C.longlong(socketState), C.longlong(openMode))) != 0
	}
	return false
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

func (ptr *QBluetoothSocket) SocketDescriptor() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QBluetoothSocket_SocketDescriptor(ptr.Pointer())))
	}
	return 0
}

func (ptr *QBluetoothSocket) SocketType() QBluetoothServiceInfo__Protocol {
	if ptr.Pointer() != nil {
		return QBluetoothServiceInfo__Protocol(C.QBluetoothSocket_SocketType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QBluetoothSocket) State() QBluetoothSocket__SocketState {
	if ptr.Pointer() != nil {
		return QBluetoothSocket__SocketState(C.QBluetoothSocket_State(ptr.Pointer()))
	}
	return 0
}

//export callbackQBluetoothSocket_StateChanged
func callbackQBluetoothSocket_StateChanged(ptr unsafe.Pointer, state C.longlong) {
	if signal := qt.GetSignal(ptr, "stateChanged"); signal != nil {
		(*(*func(QBluetoothSocket__SocketState))(signal))(QBluetoothSocket__SocketState(state))
	}

}

func (ptr *QBluetoothSocket) ConnectStateChanged(f func(state QBluetoothSocket__SocketState)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "stateChanged") {
			C.QBluetoothSocket_ConnectStateChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "stateChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "stateChanged"); signal != nil {
			f := func(state QBluetoothSocket__SocketState) {
				(*(*func(QBluetoothSocket__SocketState))(signal))(state)
				f(state)
			}
			qt.ConnectSignal(ptr.Pointer(), "stateChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stateChanged", unsafe.Pointer(&f))
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

//export callbackQBluetoothSocket_WriteData
func callbackQBluetoothSocket_WriteData(ptr unsafe.Pointer, data C.struct_QtBluetooth_PackedString, maxSize C.longlong) C.longlong {
	if signal := qt.GetSignal(ptr, "writeData"); signal != nil {
		return C.longlong((*(*func([]byte, int64) int64)(signal))(cGoUnpackBytes(data), int64(maxSize)))
	}

	return C.longlong(NewQBluetoothSocketFromPointer(ptr).WriteDataDefault(cGoUnpackBytes(data), int64(maxSize)))
}

func (ptr *QBluetoothSocket) ConnectWriteData(f func(data []byte, maxSize int64) int64) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "writeData"); signal != nil {
			f := func(data []byte, maxSize int64) int64 {
				(*(*func([]byte, int64) int64)(signal))(data, maxSize)
				return f(data, maxSize)
			}
			qt.ConnectSignal(ptr.Pointer(), "writeData", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "writeData", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBluetoothSocket) DisconnectWriteData() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "writeData")
	}
}

func (ptr *QBluetoothSocket) WriteData(data []byte, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if len(data) != 0 {
			dataC = (*C.char)(unsafe.Pointer(&data[0]))
		}
		return int64(C.QBluetoothSocket_WriteData(ptr.Pointer(), dataC, C.longlong(maxSize)))
	}
	return 0
}

func (ptr *QBluetoothSocket) WriteDataDefault(data []byte, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if len(data) != 0 {
			dataC = (*C.char)(unsafe.Pointer(&data[0]))
		}
		return int64(C.QBluetoothSocket_WriteDataDefault(ptr.Pointer(), dataC, C.longlong(maxSize)))
	}
	return 0
}

//export callbackQBluetoothSocket_DestroyQBluetoothSocket
func callbackQBluetoothSocket_DestroyQBluetoothSocket(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QBluetoothSocket"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQBluetoothSocketFromPointer(ptr).DestroyQBluetoothSocketDefault()
	}
}

func (ptr *QBluetoothSocket) ConnectDestroyQBluetoothSocket(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QBluetoothSocket"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QBluetoothSocket", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QBluetoothSocket", unsafe.Pointer(&f))
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

		qt.SetFinalizer(ptr, nil)
		C.QBluetoothSocket_DestroyQBluetoothSocket(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothSocket) DestroyQBluetoothSocketDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QBluetoothSocket_DestroyQBluetoothSocketDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothSocket) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QBluetoothSocket___children_atList(ptr.Pointer(), C.int(int32(i))))
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
	return C.QBluetoothSocket___children_newList(ptr.Pointer())
}

func (ptr *QBluetoothSocket) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QBluetoothSocket___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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
	return C.QBluetoothSocket___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QBluetoothSocket) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QBluetoothSocket___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
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
	return C.QBluetoothSocket___findChildren_newList(ptr.Pointer())
}

func (ptr *QBluetoothSocket) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QBluetoothSocket___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
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
	return C.QBluetoothSocket___findChildren_newList3(ptr.Pointer())
}

//export callbackQBluetoothSocket_AboutToClose
func callbackQBluetoothSocket_AboutToClose(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "aboutToClose"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackQBluetoothSocket_AtEnd
func callbackQBluetoothSocket_AtEnd(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "atEnd"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothSocketFromPointer(ptr).AtEndDefault())))
}

func (ptr *QBluetoothSocket) AtEndDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QBluetoothSocket_AtEndDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQBluetoothSocket_BytesWritten
func callbackQBluetoothSocket_BytesWritten(ptr unsafe.Pointer, bytes C.longlong) {
	if signal := qt.GetSignal(ptr, "bytesWritten"); signal != nil {
		(*(*func(int64))(signal))(int64(bytes))
	}

}

//export callbackQBluetoothSocket_ChannelBytesWritten
func callbackQBluetoothSocket_ChannelBytesWritten(ptr unsafe.Pointer, channel C.int, bytes C.longlong) {
	if signal := qt.GetSignal(ptr, "channelBytesWritten"); signal != nil {
		(*(*func(int, int64))(signal))(int(int32(channel)), int64(bytes))
	}

}

//export callbackQBluetoothSocket_ChannelReadyRead
func callbackQBluetoothSocket_ChannelReadyRead(ptr unsafe.Pointer, channel C.int) {
	if signal := qt.GetSignal(ptr, "channelReadyRead"); signal != nil {
		(*(*func(int))(signal))(int(int32(channel)))
	}

}

//export callbackQBluetoothSocket_Open
func callbackQBluetoothSocket_Open(ptr unsafe.Pointer, mode C.longlong) C.char {
	if signal := qt.GetSignal(ptr, "open"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(core.QIODevice__OpenModeFlag) bool)(signal))(core.QIODevice__OpenModeFlag(mode)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothSocketFromPointer(ptr).OpenDefault(core.QIODevice__OpenModeFlag(mode)))))
}

func (ptr *QBluetoothSocket) OpenDefault(mode core.QIODevice__OpenModeFlag) bool {
	if ptr.Pointer() != nil {
		return int8(C.QBluetoothSocket_OpenDefault(ptr.Pointer(), C.longlong(mode))) != 0
	}
	return false
}

//export callbackQBluetoothSocket_Pos
func callbackQBluetoothSocket_Pos(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "pos"); signal != nil {
		return C.longlong((*(*func() int64)(signal))())
	}

	return C.longlong(NewQBluetoothSocketFromPointer(ptr).PosDefault())
}

func (ptr *QBluetoothSocket) PosDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QBluetoothSocket_PosDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQBluetoothSocket_ReadChannelFinished
func callbackQBluetoothSocket_ReadChannelFinished(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "readChannelFinished"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackQBluetoothSocket_ReadLineData
func callbackQBluetoothSocket_ReadLineData(ptr unsafe.Pointer, data C.struct_QtBluetooth_PackedString, maxSize C.longlong) C.longlong {
	if signal := qt.GetSignal(ptr, "readLineData"); signal != nil {
		return C.longlong((*(*func([]byte, int64) int64)(signal))(cGoUnpackBytes(data), int64(maxSize)))
	}

	return C.longlong(NewQBluetoothSocketFromPointer(ptr).ReadLineDataDefault(cGoUnpackBytes(data), int64(maxSize)))
}

func (ptr *QBluetoothSocket) ReadLineDataDefault(data []byte, maxSize int64) int64 {
	if ptr.Pointer() != nil {
		var dataC *C.char
		if len(data) != 0 {
			dataC = (*C.char)(unsafe.Pointer(&data[0]))
		}
		return int64(C.QBluetoothSocket_ReadLineDataDefault(ptr.Pointer(), dataC, C.longlong(maxSize)))
	}
	return 0
}

//export callbackQBluetoothSocket_ReadyRead
func callbackQBluetoothSocket_ReadyRead(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "readyRead"); signal != nil {
		(*(*func())(signal))()
	}

}

//export callbackQBluetoothSocket_Reset
func callbackQBluetoothSocket_Reset(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "reset"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothSocketFromPointer(ptr).ResetDefault())))
}

func (ptr *QBluetoothSocket) ResetDefault() bool {
	if ptr.Pointer() != nil {
		return int8(C.QBluetoothSocket_ResetDefault(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQBluetoothSocket_Seek
func callbackQBluetoothSocket_Seek(ptr unsafe.Pointer, pos C.longlong) C.char {
	if signal := qt.GetSignal(ptr, "seek"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int64) bool)(signal))(int64(pos)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothSocketFromPointer(ptr).SeekDefault(int64(pos)))))
}

func (ptr *QBluetoothSocket) SeekDefault(pos int64) bool {
	if ptr.Pointer() != nil {
		return int8(C.QBluetoothSocket_SeekDefault(ptr.Pointer(), C.longlong(pos))) != 0
	}
	return false
}

//export callbackQBluetoothSocket_Size
func callbackQBluetoothSocket_Size(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "size"); signal != nil {
		return C.longlong((*(*func() int64)(signal))())
	}

	return C.longlong(NewQBluetoothSocketFromPointer(ptr).SizeDefault())
}

func (ptr *QBluetoothSocket) SizeDefault() int64 {
	if ptr.Pointer() != nil {
		return int64(C.QBluetoothSocket_SizeDefault(ptr.Pointer()))
	}
	return 0
}

//export callbackQBluetoothSocket_WaitForBytesWritten
func callbackQBluetoothSocket_WaitForBytesWritten(ptr unsafe.Pointer, msecs C.int) C.char {
	if signal := qt.GetSignal(ptr, "waitForBytesWritten"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int) bool)(signal))(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothSocketFromPointer(ptr).WaitForBytesWrittenDefault(int(int32(msecs))))))
}

func (ptr *QBluetoothSocket) WaitForBytesWrittenDefault(msecs int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QBluetoothSocket_WaitForBytesWrittenDefault(ptr.Pointer(), C.int(int32(msecs)))) != 0
	}
	return false
}

//export callbackQBluetoothSocket_WaitForReadyRead
func callbackQBluetoothSocket_WaitForReadyRead(ptr unsafe.Pointer, msecs C.int) C.char {
	if signal := qt.GetSignal(ptr, "waitForReadyRead"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(int) bool)(signal))(int(int32(msecs))))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothSocketFromPointer(ptr).WaitForReadyReadDefault(int(int32(msecs))))))
}

func (ptr *QBluetoothSocket) WaitForReadyReadDefault(msecs int) bool {
	if ptr.Pointer() != nil {
		return int8(C.QBluetoothSocket_WaitForReadyReadDefault(ptr.Pointer(), C.int(int32(msecs)))) != 0
	}
	return false
}

//export callbackQBluetoothSocket_ChildEvent
func callbackQBluetoothSocket_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
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
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
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
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
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
		(*(*func())(signal))()
	} else {
		NewQBluetoothSocketFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QBluetoothSocket) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QBluetoothSocket_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQBluetoothSocket_Destroyed
func callbackQBluetoothSocket_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQBluetoothSocket_DisconnectNotify
func callbackQBluetoothSocket_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBluetoothSocketFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBluetoothSocket) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBluetoothSocket_Event
func callbackQBluetoothSocket_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothSocketFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QBluetoothSocket) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QBluetoothSocket_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQBluetoothSocket_EventFilter
func callbackQBluetoothSocket_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothSocketFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QBluetoothSocket) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QBluetoothSocket_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQBluetoothSocket_MetaObject
func callbackQBluetoothSocket_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQBluetoothSocketFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QBluetoothSocket) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QBluetoothSocket_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQBluetoothSocket_ObjectNameChanged
func callbackQBluetoothSocket_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtBluetooth_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQBluetoothSocket_TimerEvent
func callbackQBluetoothSocket_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQBluetoothSocketFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QBluetoothSocket) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothSocket_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
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

func NewQBluetoothTransferManagerFromPointer(ptr unsafe.Pointer) (n *QBluetoothTransferManager) {
	n = new(QBluetoothTransferManager)
	n.SetPointer(ptr)
	return
}
func NewQBluetoothTransferManager(parent core.QObject_ITF) *QBluetoothTransferManager {
	tmpValue := NewQBluetoothTransferManagerFromPointer(C.QBluetoothTransferManager_NewQBluetoothTransferManager(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQBluetoothTransferManager_Finished
func callbackQBluetoothTransferManager_Finished(ptr unsafe.Pointer, reply unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "finished"); signal != nil {
		(*(*func(*QBluetoothTransferReply))(signal))(NewQBluetoothTransferReplyFromPointer(reply))
	}

}

func (ptr *QBluetoothTransferManager) ConnectFinished(f func(reply *QBluetoothTransferReply)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "finished") {
			C.QBluetoothTransferManager_ConnectFinished(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "finished")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "finished"); signal != nil {
			f := func(reply *QBluetoothTransferReply) {
				(*(*func(*QBluetoothTransferReply))(signal))(reply)
				f(reply)
			}
			qt.ConnectSignal(ptr.Pointer(), "finished", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "finished", unsafe.Pointer(&f))
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

func (ptr *QBluetoothTransferManager) Put(request QBluetoothTransferRequest_ITF, data core.QIODevice_ITF) *QBluetoothTransferReply {
	if ptr.Pointer() != nil {
		tmpValue := NewQBluetoothTransferReplyFromPointer(C.QBluetoothTransferManager_Put(ptr.Pointer(), PointerFromQBluetoothTransferRequest(request), core.PointerFromQIODevice(data)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

//export callbackQBluetoothTransferManager_DestroyQBluetoothTransferManager
func callbackQBluetoothTransferManager_DestroyQBluetoothTransferManager(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QBluetoothTransferManager"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQBluetoothTransferManagerFromPointer(ptr).DestroyQBluetoothTransferManagerDefault()
	}
}

func (ptr *QBluetoothTransferManager) ConnectDestroyQBluetoothTransferManager(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QBluetoothTransferManager"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QBluetoothTransferManager", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QBluetoothTransferManager", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBluetoothTransferManager) DisconnectDestroyQBluetoothTransferManager() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QBluetoothTransferManager")
	}
}

func (ptr *QBluetoothTransferManager) DestroyQBluetoothTransferManager() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QBluetoothTransferManager_DestroyQBluetoothTransferManager(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothTransferManager) DestroyQBluetoothTransferManagerDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QBluetoothTransferManager_DestroyQBluetoothTransferManagerDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothTransferManager) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QBluetoothTransferManager___children_atList(ptr.Pointer(), C.int(int32(i))))
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
	return C.QBluetoothTransferManager___children_newList(ptr.Pointer())
}

func (ptr *QBluetoothTransferManager) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QBluetoothTransferManager___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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
	return C.QBluetoothTransferManager___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QBluetoothTransferManager) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QBluetoothTransferManager___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
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
	return C.QBluetoothTransferManager___findChildren_newList(ptr.Pointer())
}

func (ptr *QBluetoothTransferManager) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QBluetoothTransferManager___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
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
	return C.QBluetoothTransferManager___findChildren_newList3(ptr.Pointer())
}

//export callbackQBluetoothTransferManager_ChildEvent
func callbackQBluetoothTransferManager_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
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
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
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
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
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
		(*(*func())(signal))()
	} else {
		NewQBluetoothTransferManagerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QBluetoothTransferManager) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QBluetoothTransferManager_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQBluetoothTransferManager_Destroyed
func callbackQBluetoothTransferManager_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQBluetoothTransferManager_DisconnectNotify
func callbackQBluetoothTransferManager_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBluetoothTransferManagerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBluetoothTransferManager) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBluetoothTransferManager_Event
func callbackQBluetoothTransferManager_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothTransferManagerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QBluetoothTransferManager) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QBluetoothTransferManager_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQBluetoothTransferManager_EventFilter
func callbackQBluetoothTransferManager_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothTransferManagerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QBluetoothTransferManager) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QBluetoothTransferManager_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQBluetoothTransferManager_MetaObject
func callbackQBluetoothTransferManager_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQBluetoothTransferManagerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QBluetoothTransferManager) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QBluetoothTransferManager_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQBluetoothTransferManager_ObjectNameChanged
func callbackQBluetoothTransferManager_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtBluetooth_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQBluetoothTransferManager_TimerEvent
func callbackQBluetoothTransferManager_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQBluetoothTransferManagerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QBluetoothTransferManager) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferManager_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
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

func NewQBluetoothTransferReplyFromPointer(ptr unsafe.Pointer) (n *QBluetoothTransferReply) {
	n = new(QBluetoothTransferReply)
	n.SetPointer(ptr)
	return
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
	tmpValue := NewQBluetoothTransferReplyFromPointer(C.QBluetoothTransferReply_NewQBluetoothTransferReply(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

//export callbackQBluetoothTransferReply_Abort
func callbackQBluetoothTransferReply_Abort(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "abort"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQBluetoothTransferReplyFromPointer(ptr).AbortDefault()
	}
}

func (ptr *QBluetoothTransferReply) ConnectAbort(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "abort"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "abort", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "abort", unsafe.Pointer(&f))
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

//export callbackQBluetoothTransferReply_Error
func callbackQBluetoothTransferReply_Error(ptr unsafe.Pointer) C.longlong {
	if signal := qt.GetSignal(ptr, "error"); signal != nil {
		return C.longlong((*(*func() QBluetoothTransferReply__TransferError)(signal))())
	}

	return C.longlong(0)
}

func (ptr *QBluetoothTransferReply) ConnectError(f func() QBluetoothTransferReply__TransferError) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "error"); signal != nil {
			f := func() QBluetoothTransferReply__TransferError {
				(*(*func() QBluetoothTransferReply__TransferError)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "error", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "error", unsafe.Pointer(&f))
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

//export callbackQBluetoothTransferReply_Error2
func callbackQBluetoothTransferReply_Error2(ptr unsafe.Pointer, errorType C.longlong) {
	if signal := qt.GetSignal(ptr, "error2"); signal != nil {
		(*(*func(QBluetoothTransferReply__TransferError))(signal))(QBluetoothTransferReply__TransferError(errorType))
	}

}

func (ptr *QBluetoothTransferReply) ConnectError2(f func(errorType QBluetoothTransferReply__TransferError)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "error2") {
			C.QBluetoothTransferReply_ConnectError2(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "error")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "error2"); signal != nil {
			f := func(errorType QBluetoothTransferReply__TransferError) {
				(*(*func(QBluetoothTransferReply__TransferError))(signal))(errorType)
				f(errorType)
			}
			qt.ConnectSignal(ptr.Pointer(), "error2", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "error2", unsafe.Pointer(&f))
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

//export callbackQBluetoothTransferReply_ErrorString
func callbackQBluetoothTransferReply_ErrorString(ptr unsafe.Pointer) C.struct_QtBluetooth_PackedString {
	if signal := qt.GetSignal(ptr, "errorString"); signal != nil {
		tempVal := (*(*func() string)(signal))()
		return C.struct_QtBluetooth_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
	}
	tempVal := ""
	return C.struct_QtBluetooth_PackedString{data: C.CString(tempVal), len: C.longlong(len(tempVal))}
}

func (ptr *QBluetoothTransferReply) ConnectErrorString(f func() string) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "errorString"); signal != nil {
			f := func() string {
				(*(*func() string)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "errorString", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "errorString", unsafe.Pointer(&f))
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

//export callbackQBluetoothTransferReply_Finished
func callbackQBluetoothTransferReply_Finished(ptr unsafe.Pointer, reply unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "finished"); signal != nil {
		(*(*func(*QBluetoothTransferReply))(signal))(NewQBluetoothTransferReplyFromPointer(reply))
	}

}

func (ptr *QBluetoothTransferReply) ConnectFinished(f func(reply *QBluetoothTransferReply)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "finished") {
			C.QBluetoothTransferReply_ConnectFinished(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "finished")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "finished"); signal != nil {
			f := func(reply *QBluetoothTransferReply) {
				(*(*func(*QBluetoothTransferReply))(signal))(reply)
				f(reply)
			}
			qt.ConnectSignal(ptr.Pointer(), "finished", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "finished", unsafe.Pointer(&f))
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

//export callbackQBluetoothTransferReply_IsFinished
func callbackQBluetoothTransferReply_IsFinished(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isFinished"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QBluetoothTransferReply) ConnectIsFinished(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "isFinished"); signal != nil {
			f := func() bool {
				(*(*func() bool)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "isFinished", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "isFinished", unsafe.Pointer(&f))
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
		return int8(C.QBluetoothTransferReply_IsFinished(ptr.Pointer())) != 0
	}
	return false
}

//export callbackQBluetoothTransferReply_IsRunning
func callbackQBluetoothTransferReply_IsRunning(ptr unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "isRunning"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func() bool)(signal))())))
	}

	return C.char(int8(qt.GoBoolToInt(false)))
}

func (ptr *QBluetoothTransferReply) ConnectIsRunning(f func() bool) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "isRunning"); signal != nil {
			f := func() bool {
				(*(*func() bool)(signal))()
				return f()
			}
			qt.ConnectSignal(ptr.Pointer(), "isRunning", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "isRunning", unsafe.Pointer(&f))
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
		return int8(C.QBluetoothTransferReply_IsRunning(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QBluetoothTransferReply) Manager() *QBluetoothTransferManager {
	if ptr.Pointer() != nil {
		tmpValue := NewQBluetoothTransferManagerFromPointer(C.QBluetoothTransferReply_Manager(ptr.Pointer()))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothTransferReply) Request() *QBluetoothTransferRequest {
	if ptr.Pointer() != nil {
		tmpValue := NewQBluetoothTransferRequestFromPointer(C.QBluetoothTransferReply_Request(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QBluetoothTransferRequest).DestroyQBluetoothTransferRequest)
		return tmpValue
	}
	return nil
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
		(*(*func(int64, int64))(signal))(int64(bytesTransferred), int64(bytesTotal))
	}

}

func (ptr *QBluetoothTransferReply) ConnectTransferProgress(f func(bytesTransferred int64, bytesTotal int64)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "transferProgress") {
			C.QBluetoothTransferReply_ConnectTransferProgress(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "transferProgress")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "transferProgress"); signal != nil {
			f := func(bytesTransferred int64, bytesTotal int64) {
				(*(*func(int64, int64))(signal))(bytesTransferred, bytesTotal)
				f(bytesTransferred, bytesTotal)
			}
			qt.ConnectSignal(ptr.Pointer(), "transferProgress", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "transferProgress", unsafe.Pointer(&f))
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

//export callbackQBluetoothTransferReply_DestroyQBluetoothTransferReply
func callbackQBluetoothTransferReply_DestroyQBluetoothTransferReply(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QBluetoothTransferReply"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQBluetoothTransferReplyFromPointer(ptr).DestroyQBluetoothTransferReplyDefault()
	}
}

func (ptr *QBluetoothTransferReply) ConnectDestroyQBluetoothTransferReply(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QBluetoothTransferReply"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QBluetoothTransferReply", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QBluetoothTransferReply", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QBluetoothTransferReply) DisconnectDestroyQBluetoothTransferReply() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QBluetoothTransferReply")
	}
}

func (ptr *QBluetoothTransferReply) DestroyQBluetoothTransferReply() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QBluetoothTransferReply_DestroyQBluetoothTransferReply(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothTransferReply) DestroyQBluetoothTransferReplyDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QBluetoothTransferReply_DestroyQBluetoothTransferReplyDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QBluetoothTransferReply) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QBluetoothTransferReply___children_atList(ptr.Pointer(), C.int(int32(i))))
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
	return C.QBluetoothTransferReply___children_newList(ptr.Pointer())
}

func (ptr *QBluetoothTransferReply) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QBluetoothTransferReply___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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
	return C.QBluetoothTransferReply___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QBluetoothTransferReply) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QBluetoothTransferReply___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
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
	return C.QBluetoothTransferReply___findChildren_newList(ptr.Pointer())
}

func (ptr *QBluetoothTransferReply) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QBluetoothTransferReply___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
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
	return C.QBluetoothTransferReply___findChildren_newList3(ptr.Pointer())
}

//export callbackQBluetoothTransferReply_ChildEvent
func callbackQBluetoothTransferReply_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
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
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
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
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
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
		(*(*func())(signal))()
	} else {
		NewQBluetoothTransferReplyFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QBluetoothTransferReply) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QBluetoothTransferReply_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQBluetoothTransferReply_Destroyed
func callbackQBluetoothTransferReply_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQBluetoothTransferReply_DisconnectNotify
func callbackQBluetoothTransferReply_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQBluetoothTransferReplyFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QBluetoothTransferReply) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQBluetoothTransferReply_Event
func callbackQBluetoothTransferReply_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothTransferReplyFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QBluetoothTransferReply) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QBluetoothTransferReply_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQBluetoothTransferReply_EventFilter
func callbackQBluetoothTransferReply_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQBluetoothTransferReplyFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QBluetoothTransferReply) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QBluetoothTransferReply_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQBluetoothTransferReply_MetaObject
func callbackQBluetoothTransferReply_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQBluetoothTransferReplyFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QBluetoothTransferReply) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QBluetoothTransferReply_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQBluetoothTransferReply_ObjectNameChanged
func callbackQBluetoothTransferReply_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtBluetooth_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQBluetoothTransferReply_TimerEvent
func callbackQBluetoothTransferReply_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQBluetoothTransferReplyFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QBluetoothTransferReply) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferReply_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
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

func NewQBluetoothTransferRequestFromPointer(ptr unsafe.Pointer) (n *QBluetoothTransferRequest) {
	n = new(QBluetoothTransferRequest)
	n.SetPointer(ptr)
	return
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
	tmpValue := NewQBluetoothTransferRequestFromPointer(C.QBluetoothTransferRequest_NewQBluetoothTransferRequest(PointerFromQBluetoothAddress(address)))
	qt.SetFinalizer(tmpValue, (*QBluetoothTransferRequest).DestroyQBluetoothTransferRequest)
	return tmpValue
}

func NewQBluetoothTransferRequest2(other QBluetoothTransferRequest_ITF) *QBluetoothTransferRequest {
	tmpValue := NewQBluetoothTransferRequestFromPointer(C.QBluetoothTransferRequest_NewQBluetoothTransferRequest2(PointerFromQBluetoothTransferRequest(other)))
	qt.SetFinalizer(tmpValue, (*QBluetoothTransferRequest).DestroyQBluetoothTransferRequest)
	return tmpValue
}

func (ptr *QBluetoothTransferRequest) Address() *QBluetoothAddress {
	if ptr.Pointer() != nil {
		tmpValue := NewQBluetoothAddressFromPointer(C.QBluetoothTransferRequest_Address(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothTransferRequest) Attribute(code QBluetoothTransferRequest__Attribute, defaultValue core.QVariant_ITF) *core.QVariant {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQVariantFromPointer(C.QBluetoothTransferRequest_Attribute(ptr.Pointer(), C.longlong(code), core.PointerFromQVariant(defaultValue)))
		qt.SetFinalizer(tmpValue, (*core.QVariant).DestroyQVariant)
		return tmpValue
	}
	return nil
}

func (ptr *QBluetoothTransferRequest) SetAttribute(code QBluetoothTransferRequest__Attribute, value core.QVariant_ITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferRequest_SetAttribute(ptr.Pointer(), C.longlong(code), core.PointerFromQVariant(value))
	}
}

func (ptr *QBluetoothTransferRequest) DestroyQBluetoothTransferRequest() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QBluetoothTransferRequest_DestroyQBluetoothTransferRequest(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
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

func NewQBluetoothUuidFromPointer(ptr unsafe.Pointer) (n *QBluetoothUuid) {
	n = new(QBluetoothUuid)
	n.SetPointer(ptr)
	return
}

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

func NewQBluetoothUuid() *QBluetoothUuid {
	tmpValue := NewQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid())
	qt.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
	return tmpValue
}

func NewQBluetoothUuid2(uuid QBluetoothUuid__ProtocolUuid) *QBluetoothUuid {
	tmpValue := NewQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid2(C.longlong(uuid)))
	qt.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
	return tmpValue
}

func NewQBluetoothUuid3(uuid QBluetoothUuid__ServiceClassUuid) *QBluetoothUuid {
	tmpValue := NewQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid3(C.longlong(uuid)))
	qt.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
	return tmpValue
}

func NewQBluetoothUuid4(uuid QBluetoothUuid__CharacteristicType) *QBluetoothUuid {
	tmpValue := NewQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid4(C.longlong(uuid)))
	qt.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
	return tmpValue
}

func NewQBluetoothUuid5(uuid QBluetoothUuid__DescriptorType) *QBluetoothUuid {
	tmpValue := NewQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid5(C.longlong(uuid)))
	qt.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
	return tmpValue
}

func NewQBluetoothUuid6(uuid uint16) *QBluetoothUuid {
	tmpValue := NewQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid6(C.ushort(uuid)))
	qt.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
	return tmpValue
}

func NewQBluetoothUuid7(uuid uint) *QBluetoothUuid {
	tmpValue := NewQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid7(C.uint(uint32(uuid))))
	qt.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
	return tmpValue
}

func NewQBluetoothUuid9(uuid string) *QBluetoothUuid {
	var uuidC *C.char
	if uuid != "" {
		uuidC = C.CString(uuid)
		defer C.free(unsafe.Pointer(uuidC))
	}
	tmpValue := NewQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid9(C.struct_QtBluetooth_PackedString{data: uuidC, len: C.longlong(len(uuid))}))
	qt.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
	return tmpValue
}

func NewQBluetoothUuid10(uuid QBluetoothUuid_ITF) *QBluetoothUuid {
	tmpValue := NewQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid10(PointerFromQBluetoothUuid(uuid)))
	qt.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
	return tmpValue
}

func NewQBluetoothUuid11(uuid core.QUuid_ITF) *QBluetoothUuid {
	tmpValue := NewQBluetoothUuidFromPointer(C.QBluetoothUuid_NewQBluetoothUuid11(core.PointerFromQUuid(uuid)))
	qt.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
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

func (ptr *QBluetoothUuid) MinimumSize() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QBluetoothUuid_MinimumSize(ptr.Pointer())))
	}
	return 0
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

func (ptr *QBluetoothUuid) ToUInt16(ok *bool) uint16 {
	if ptr.Pointer() != nil {
		var okC C.char
		if ok != nil {
			okC = C.char(int8(qt.GoBoolToInt(*ok)))
			defer func() { *ok = int8(okC) != 0 }()
		}
		return uint16(C.QBluetoothUuid_ToUInt16(ptr.Pointer(), &okC))
	}
	return 0
}

func (ptr *QBluetoothUuid) ToUInt32(ok *bool) uint {
	if ptr.Pointer() != nil {
		var okC C.char
		if ok != nil {
			okC = C.char(int8(qt.GoBoolToInt(*ok)))
			defer func() { *ok = int8(okC) != 0 }()
		}
		return uint(uint32(C.QBluetoothUuid_ToUInt32(ptr.Pointer(), &okC)))
	}
	return 0
}

func (ptr *QBluetoothUuid) DestroyQBluetoothUuid() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QBluetoothUuid_DestroyQBluetoothUuid(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
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

func NewQLowEnergyAdvertisingDataFromPointer(ptr unsafe.Pointer) (n *QLowEnergyAdvertisingData) {
	n = new(QLowEnergyAdvertisingData)
	n.SetPointer(ptr)
	return
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
	tmpValue := NewQLowEnergyAdvertisingDataFromPointer(C.QLowEnergyAdvertisingData_NewQLowEnergyAdvertisingData())
	qt.SetFinalizer(tmpValue, (*QLowEnergyAdvertisingData).DestroyQLowEnergyAdvertisingData)
	return tmpValue
}

func NewQLowEnergyAdvertisingData2(other QLowEnergyAdvertisingData_ITF) *QLowEnergyAdvertisingData {
	tmpValue := NewQLowEnergyAdvertisingDataFromPointer(C.QLowEnergyAdvertisingData_NewQLowEnergyAdvertisingData2(PointerFromQLowEnergyAdvertisingData(other)))
	qt.SetFinalizer(tmpValue, (*QLowEnergyAdvertisingData).DestroyQLowEnergyAdvertisingData)
	return tmpValue
}

func (ptr *QLowEnergyAdvertisingData) Discoverability() QLowEnergyAdvertisingData__Discoverability {
	if ptr.Pointer() != nil {
		return QLowEnergyAdvertisingData__Discoverability(C.QLowEnergyAdvertisingData_Discoverability(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyAdvertisingData) IncludePowerLevel() bool {
	if ptr.Pointer() != nil {
		return int8(C.QLowEnergyAdvertisingData_IncludePowerLevel(ptr.Pointer())) != 0
	}
	return false
}

func QLowEnergyAdvertisingData_InvalidManufacturerId() uint16 {
	return uint16(C.QLowEnergyAdvertisingData_QLowEnergyAdvertisingData_InvalidManufacturerId())
}

func (ptr *QLowEnergyAdvertisingData) InvalidManufacturerId() uint16 {
	return uint16(C.QLowEnergyAdvertisingData_QLowEnergyAdvertisingData_InvalidManufacturerId())
}

func (ptr *QLowEnergyAdvertisingData) LocalName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QLowEnergyAdvertisingData_LocalName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLowEnergyAdvertisingData) ManufacturerId() uint16 {
	if ptr.Pointer() != nil {
		return uint16(C.QLowEnergyAdvertisingData_ManufacturerId(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyAdvertisingData) RawData() *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QLowEnergyAdvertisingData_RawData(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyAdvertisingData) Services() []*QBluetoothUuid {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtBluetooth_PackedList) []*QBluetoothUuid {
			out := make([]*QBluetoothUuid, int(l.len))
			tmpList := NewQLowEnergyAdvertisingDataFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__services_atList(i)
			}
			return out
		}(C.QLowEnergyAdvertisingData_Services(ptr.Pointer()))
	}
	return make([]*QBluetoothUuid, 0)
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
			tmpList := NewQLowEnergyAdvertisingDataFromPointer(NewQLowEnergyAdvertisingDataFromPointer(nil).__setServices_services_newList())
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

		qt.SetFinalizer(ptr, nil)
		C.QLowEnergyAdvertisingData_DestroyQLowEnergyAdvertisingData(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QLowEnergyAdvertisingData) __services_atList(i int) *QBluetoothUuid {
	if ptr.Pointer() != nil {
		tmpValue := NewQBluetoothUuidFromPointer(C.QLowEnergyAdvertisingData___services_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
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
	return C.QLowEnergyAdvertisingData___services_newList(ptr.Pointer())
}

func (ptr *QLowEnergyAdvertisingData) __setServices_services_atList(i int) *QBluetoothUuid {
	if ptr.Pointer() != nil {
		tmpValue := NewQBluetoothUuidFromPointer(C.QLowEnergyAdvertisingData___setServices_services_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
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
	return C.QLowEnergyAdvertisingData___setServices_services_newList(ptr.Pointer())
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

func NewQLowEnergyAdvertisingParametersFromPointer(ptr unsafe.Pointer) (n *QLowEnergyAdvertisingParameters) {
	n = new(QLowEnergyAdvertisingParameters)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QLowEnergyAdvertisingParameters__Mode
//QLowEnergyAdvertisingParameters::Mode
type QLowEnergyAdvertisingParameters__Mode int64

const (
	QLowEnergyAdvertisingParameters__AdvInd        QLowEnergyAdvertisingParameters__Mode = QLowEnergyAdvertisingParameters__Mode(0x0)
	QLowEnergyAdvertisingParameters__AdvScanInd    QLowEnergyAdvertisingParameters__Mode = QLowEnergyAdvertisingParameters__Mode(0x2)
	QLowEnergyAdvertisingParameters__AdvNonConnInd QLowEnergyAdvertisingParameters__Mode = QLowEnergyAdvertisingParameters__Mode(0x3)
)

//go:generate stringer -type=QLowEnergyAdvertisingParameters__FilterPolicy
//QLowEnergyAdvertisingParameters::FilterPolicy
type QLowEnergyAdvertisingParameters__FilterPolicy int64

const (
	QLowEnergyAdvertisingParameters__IgnoreWhiteList                      QLowEnergyAdvertisingParameters__FilterPolicy = QLowEnergyAdvertisingParameters__FilterPolicy(0x00)
	QLowEnergyAdvertisingParameters__UseWhiteListForScanning              QLowEnergyAdvertisingParameters__FilterPolicy = QLowEnergyAdvertisingParameters__FilterPolicy(0x01)
	QLowEnergyAdvertisingParameters__UseWhiteListForConnecting            QLowEnergyAdvertisingParameters__FilterPolicy = QLowEnergyAdvertisingParameters__FilterPolicy(0x02)
	QLowEnergyAdvertisingParameters__UseWhiteListForScanningAndConnecting QLowEnergyAdvertisingParameters__FilterPolicy = QLowEnergyAdvertisingParameters__FilterPolicy(0x03)
)

func NewQLowEnergyAdvertisingParameters() *QLowEnergyAdvertisingParameters {
	tmpValue := NewQLowEnergyAdvertisingParametersFromPointer(C.QLowEnergyAdvertisingParameters_NewQLowEnergyAdvertisingParameters())
	qt.SetFinalizer(tmpValue, (*QLowEnergyAdvertisingParameters).DestroyQLowEnergyAdvertisingParameters)
	return tmpValue
}

func NewQLowEnergyAdvertisingParameters2(other QLowEnergyAdvertisingParameters_ITF) *QLowEnergyAdvertisingParameters {
	tmpValue := NewQLowEnergyAdvertisingParametersFromPointer(C.QLowEnergyAdvertisingParameters_NewQLowEnergyAdvertisingParameters2(PointerFromQLowEnergyAdvertisingParameters(other)))
	qt.SetFinalizer(tmpValue, (*QLowEnergyAdvertisingParameters).DestroyQLowEnergyAdvertisingParameters)
	return tmpValue
}

func (ptr *QLowEnergyAdvertisingParameters) FilterPolicy() QLowEnergyAdvertisingParameters__FilterPolicy {
	if ptr.Pointer() != nil {
		return QLowEnergyAdvertisingParameters__FilterPolicy(C.QLowEnergyAdvertisingParameters_FilterPolicy(ptr.Pointer()))
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

func (ptr *QLowEnergyAdvertisingParameters) Mode() QLowEnergyAdvertisingParameters__Mode {
	if ptr.Pointer() != nil {
		return QLowEnergyAdvertisingParameters__Mode(C.QLowEnergyAdvertisingParameters_Mode(ptr.Pointer()))
	}
	return 0
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

		qt.SetFinalizer(ptr, nil)
		C.QLowEnergyAdvertisingParameters_DestroyQLowEnergyAdvertisingParameters(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
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

func NewQLowEnergyCharacteristicFromPointer(ptr unsafe.Pointer) (n *QLowEnergyCharacteristic) {
	n = new(QLowEnergyCharacteristic)
	n.SetPointer(ptr)
	return
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
	tmpValue := NewQLowEnergyCharacteristicFromPointer(C.QLowEnergyCharacteristic_NewQLowEnergyCharacteristic())
	qt.SetFinalizer(tmpValue, (*QLowEnergyCharacteristic).DestroyQLowEnergyCharacteristic)
	return tmpValue
}

func NewQLowEnergyCharacteristic2(other QLowEnergyCharacteristic_ITF) *QLowEnergyCharacteristic {
	tmpValue := NewQLowEnergyCharacteristicFromPointer(C.QLowEnergyCharacteristic_NewQLowEnergyCharacteristic2(PointerFromQLowEnergyCharacteristic(other)))
	qt.SetFinalizer(tmpValue, (*QLowEnergyCharacteristic).DestroyQLowEnergyCharacteristic)
	return tmpValue
}

func (ptr *QLowEnergyCharacteristic) Descriptor(uuid QBluetoothUuid_ITF) *QLowEnergyDescriptor {
	if ptr.Pointer() != nil {
		tmpValue := NewQLowEnergyDescriptorFromPointer(C.QLowEnergyCharacteristic_Descriptor(ptr.Pointer(), PointerFromQBluetoothUuid(uuid)))
		qt.SetFinalizer(tmpValue, (*QLowEnergyDescriptor).DestroyQLowEnergyDescriptor)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyCharacteristic) Descriptors() []*QLowEnergyDescriptor {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtBluetooth_PackedList) []*QLowEnergyDescriptor {
			out := make([]*QLowEnergyDescriptor, int(l.len))
			tmpList := NewQLowEnergyCharacteristicFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__descriptors_atList(i)
			}
			return out
		}(C.QLowEnergyCharacteristic_Descriptors(ptr.Pointer()))
	}
	return make([]*QLowEnergyDescriptor, 0)
}

func (ptr *QLowEnergyCharacteristic) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QLowEnergyCharacteristic_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLowEnergyCharacteristic) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QLowEnergyCharacteristic_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLowEnergyCharacteristic) Properties() QLowEnergyCharacteristic__PropertyType {
	if ptr.Pointer() != nil {
		return QLowEnergyCharacteristic__PropertyType(C.QLowEnergyCharacteristic_Properties(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyCharacteristic) Uuid() *QBluetoothUuid {
	if ptr.Pointer() != nil {
		tmpValue := NewQBluetoothUuidFromPointer(C.QLowEnergyCharacteristic_Uuid(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyCharacteristic) Value() *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QLowEnergyCharacteristic_Value(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyCharacteristic) DestroyQLowEnergyCharacteristic() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QLowEnergyCharacteristic_DestroyQLowEnergyCharacteristic(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QLowEnergyCharacteristic) __descriptors_atList(i int) *QLowEnergyDescriptor {
	if ptr.Pointer() != nil {
		tmpValue := NewQLowEnergyDescriptorFromPointer(C.QLowEnergyCharacteristic___descriptors_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QLowEnergyDescriptor).DestroyQLowEnergyDescriptor)
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
	return C.QLowEnergyCharacteristic___descriptors_newList(ptr.Pointer())
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

func NewQLowEnergyCharacteristicDataFromPointer(ptr unsafe.Pointer) (n *QLowEnergyCharacteristicData) {
	n = new(QLowEnergyCharacteristicData)
	n.SetPointer(ptr)
	return
}
func NewQLowEnergyCharacteristicData() *QLowEnergyCharacteristicData {
	tmpValue := NewQLowEnergyCharacteristicDataFromPointer(C.QLowEnergyCharacteristicData_NewQLowEnergyCharacteristicData())
	qt.SetFinalizer(tmpValue, (*QLowEnergyCharacteristicData).DestroyQLowEnergyCharacteristicData)
	return tmpValue
}

func NewQLowEnergyCharacteristicData2(other QLowEnergyCharacteristicData_ITF) *QLowEnergyCharacteristicData {
	tmpValue := NewQLowEnergyCharacteristicDataFromPointer(C.QLowEnergyCharacteristicData_NewQLowEnergyCharacteristicData2(PointerFromQLowEnergyCharacteristicData(other)))
	qt.SetFinalizer(tmpValue, (*QLowEnergyCharacteristicData).DestroyQLowEnergyCharacteristicData)
	return tmpValue
}

func (ptr *QLowEnergyCharacteristicData) AddDescriptor(descriptor QLowEnergyDescriptorData_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyCharacteristicData_AddDescriptor(ptr.Pointer(), PointerFromQLowEnergyDescriptorData(descriptor))
	}
}

func (ptr *QLowEnergyCharacteristicData) Descriptors() []*QLowEnergyDescriptorData {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtBluetooth_PackedList) []*QLowEnergyDescriptorData {
			out := make([]*QLowEnergyDescriptorData, int(l.len))
			tmpList := NewQLowEnergyCharacteristicDataFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__descriptors_atList(i)
			}
			return out
		}(C.QLowEnergyCharacteristicData_Descriptors(ptr.Pointer()))
	}
	return make([]*QLowEnergyDescriptorData, 0)
}

func (ptr *QLowEnergyCharacteristicData) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QLowEnergyCharacteristicData_IsValid(ptr.Pointer())) != 0
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

func (ptr *QLowEnergyCharacteristicData) Properties() QLowEnergyCharacteristic__PropertyType {
	if ptr.Pointer() != nil {
		return QLowEnergyCharacteristic__PropertyType(C.QLowEnergyCharacteristicData_Properties(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyCharacteristicData) ReadConstraints() QBluetooth__AttAccessConstraint {
	if ptr.Pointer() != nil {
		return QBluetooth__AttAccessConstraint(C.QLowEnergyCharacteristicData_ReadConstraints(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyCharacteristicData) SetDescriptors(descriptors []*QLowEnergyDescriptorData) {
	if ptr.Pointer() != nil {
		C.QLowEnergyCharacteristicData_SetDescriptors(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQLowEnergyCharacteristicDataFromPointer(NewQLowEnergyCharacteristicDataFromPointer(nil).__setDescriptors_descriptors_newList())
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

func (ptr *QLowEnergyCharacteristicData) Uuid() *QBluetoothUuid {
	if ptr.Pointer() != nil {
		tmpValue := NewQBluetoothUuidFromPointer(C.QLowEnergyCharacteristicData_Uuid(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyCharacteristicData) Value() *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QLowEnergyCharacteristicData_Value(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyCharacteristicData) WriteConstraints() QBluetooth__AttAccessConstraint {
	if ptr.Pointer() != nil {
		return QBluetooth__AttAccessConstraint(C.QLowEnergyCharacteristicData_WriteConstraints(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyCharacteristicData) DestroyQLowEnergyCharacteristicData() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QLowEnergyCharacteristicData_DestroyQLowEnergyCharacteristicData(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QLowEnergyCharacteristicData) __descriptors_atList(i int) *QLowEnergyDescriptorData {
	if ptr.Pointer() != nil {
		tmpValue := NewQLowEnergyDescriptorDataFromPointer(C.QLowEnergyCharacteristicData___descriptors_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QLowEnergyDescriptorData).DestroyQLowEnergyDescriptorData)
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
	return C.QLowEnergyCharacteristicData___descriptors_newList(ptr.Pointer())
}

func (ptr *QLowEnergyCharacteristicData) __setDescriptors_descriptors_atList(i int) *QLowEnergyDescriptorData {
	if ptr.Pointer() != nil {
		tmpValue := NewQLowEnergyDescriptorDataFromPointer(C.QLowEnergyCharacteristicData___setDescriptors_descriptors_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QLowEnergyDescriptorData).DestroyQLowEnergyDescriptorData)
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
	return C.QLowEnergyCharacteristicData___setDescriptors_descriptors_newList(ptr.Pointer())
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

func NewQLowEnergyConnectionParametersFromPointer(ptr unsafe.Pointer) (n *QLowEnergyConnectionParameters) {
	n = new(QLowEnergyConnectionParameters)
	n.SetPointer(ptr)
	return
}
func NewQLowEnergyConnectionParameters() *QLowEnergyConnectionParameters {
	tmpValue := NewQLowEnergyConnectionParametersFromPointer(C.QLowEnergyConnectionParameters_NewQLowEnergyConnectionParameters())
	qt.SetFinalizer(tmpValue, (*QLowEnergyConnectionParameters).DestroyQLowEnergyConnectionParameters)
	return tmpValue
}

func NewQLowEnergyConnectionParameters2(other QLowEnergyConnectionParameters_ITF) *QLowEnergyConnectionParameters {
	tmpValue := NewQLowEnergyConnectionParametersFromPointer(C.QLowEnergyConnectionParameters_NewQLowEnergyConnectionParameters2(PointerFromQLowEnergyConnectionParameters(other)))
	qt.SetFinalizer(tmpValue, (*QLowEnergyConnectionParameters).DestroyQLowEnergyConnectionParameters)
	return tmpValue
}

func (ptr *QLowEnergyConnectionParameters) Latency() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QLowEnergyConnectionParameters_Latency(ptr.Pointer())))
	}
	return 0
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

func (ptr *QLowEnergyConnectionParameters) SupervisionTimeout() int {
	if ptr.Pointer() != nil {
		return int(int32(C.QLowEnergyConnectionParameters_SupervisionTimeout(ptr.Pointer())))
	}
	return 0
}

func (ptr *QLowEnergyConnectionParameters) Swap(other QLowEnergyConnectionParameters_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyConnectionParameters_Swap(ptr.Pointer(), PointerFromQLowEnergyConnectionParameters(other))
	}
}

func (ptr *QLowEnergyConnectionParameters) DestroyQLowEnergyConnectionParameters() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QLowEnergyConnectionParameters_DestroyQLowEnergyConnectionParameters(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
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

func NewQLowEnergyControllerFromPointer(ptr unsafe.Pointer) (n *QLowEnergyController) {
	n = new(QLowEnergyController)
	n.SetPointer(ptr)
	return
}

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
	QLowEnergyController__RemoteHostClosedError        QLowEnergyController__Error = QLowEnergyController__Error(7)
)

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

func (ptr *QLowEnergyController) AddService(service QLowEnergyServiceData_ITF, parent core.QObject_ITF) *QLowEnergyService {
	if ptr.Pointer() != nil {
		tmpValue := NewQLowEnergyServiceFromPointer(C.QLowEnergyController_AddService(ptr.Pointer(), PointerFromQLowEnergyServiceData(service), core.PointerFromQObject(parent)))
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
		(*(*func())(signal))()
	}

}

func (ptr *QLowEnergyController) ConnectConnected(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "connected") {
			C.QLowEnergyController_ConnectConnected(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "connected")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "connected"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "connected", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "connected", unsafe.Pointer(&f))
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
		(*(*func(*QLowEnergyConnectionParameters))(signal))(NewQLowEnergyConnectionParametersFromPointer(newParameters))
	}

}

func (ptr *QLowEnergyController) ConnectConnectionUpdated(f func(newParameters *QLowEnergyConnectionParameters)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "connectionUpdated") {
			C.QLowEnergyController_ConnectConnectionUpdated(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "connectionUpdated")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "connectionUpdated"); signal != nil {
			f := func(newParameters *QLowEnergyConnectionParameters) {
				(*(*func(*QLowEnergyConnectionParameters))(signal))(newParameters)
				f(newParameters)
			}
			qt.ConnectSignal(ptr.Pointer(), "connectionUpdated", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "connectionUpdated", unsafe.Pointer(&f))
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

func QLowEnergyController_CreateCentral(remoteDevice QBluetoothDeviceInfo_ITF, parent core.QObject_ITF) *QLowEnergyController {
	tmpValue := NewQLowEnergyControllerFromPointer(C.QLowEnergyController_QLowEnergyController_CreateCentral(PointerFromQBluetoothDeviceInfo(remoteDevice), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QLowEnergyController) CreateCentral(remoteDevice QBluetoothDeviceInfo_ITF, parent core.QObject_ITF) *QLowEnergyController {
	tmpValue := NewQLowEnergyControllerFromPointer(C.QLowEnergyController_QLowEnergyController_CreateCentral(PointerFromQBluetoothDeviceInfo(remoteDevice), core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func QLowEnergyController_CreatePeripheral(parent core.QObject_ITF) *QLowEnergyController {
	tmpValue := NewQLowEnergyControllerFromPointer(C.QLowEnergyController_QLowEnergyController_CreatePeripheral(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QLowEnergyController) CreatePeripheral(parent core.QObject_ITF) *QLowEnergyController {
	tmpValue := NewQLowEnergyControllerFromPointer(C.QLowEnergyController_QLowEnergyController_CreatePeripheral(core.PointerFromQObject(parent)))
	if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
		tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
	}
	return tmpValue
}

func (ptr *QLowEnergyController) CreateServiceObject(serviceUuid QBluetoothUuid_ITF, parent core.QObject_ITF) *QLowEnergyService {
	if ptr.Pointer() != nil {
		tmpValue := NewQLowEnergyServiceFromPointer(C.QLowEnergyController_CreateServiceObject(ptr.Pointer(), PointerFromQBluetoothUuid(serviceUuid), core.PointerFromQObject(parent)))
		if !qt.ExistsSignal(tmpValue.Pointer(), "destroyed") {
			tmpValue.ConnectDestroyed(func(*core.QObject) { tmpValue.SetPointer(nil) })
		}
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyController) DisconnectFromDevice() {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_DisconnectFromDevice(ptr.Pointer())
	}
}

//export callbackQLowEnergyController_Disconnected
func callbackQLowEnergyController_Disconnected(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnected"); signal != nil {
		(*(*func())(signal))()
	}

}

func (ptr *QLowEnergyController) ConnectDisconnected(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "disconnected") {
			C.QLowEnergyController_ConnectDisconnected(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "disconnected")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "disconnected"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "disconnected", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "disconnected", unsafe.Pointer(&f))
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
		(*(*func())(signal))()
	}

}

func (ptr *QLowEnergyController) ConnectDiscoveryFinished(f func()) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "discoveryFinished") {
			C.QLowEnergyController_ConnectDiscoveryFinished(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "discoveryFinished")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "discoveryFinished"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "discoveryFinished", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "discoveryFinished", unsafe.Pointer(&f))
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

func (ptr *QLowEnergyController) Error() QLowEnergyController__Error {
	if ptr.Pointer() != nil {
		return QLowEnergyController__Error(C.QLowEnergyController_Error(ptr.Pointer()))
	}
	return 0
}

//export callbackQLowEnergyController_Error2
func callbackQLowEnergyController_Error2(ptr unsafe.Pointer, newError C.longlong) {
	if signal := qt.GetSignal(ptr, "error2"); signal != nil {
		(*(*func(QLowEnergyController__Error))(signal))(QLowEnergyController__Error(newError))
	}

}

func (ptr *QLowEnergyController) ConnectError2(f func(newError QLowEnergyController__Error)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "error2") {
			C.QLowEnergyController_ConnectError2(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "error")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "error2"); signal != nil {
			f := func(newError QLowEnergyController__Error) {
				(*(*func(QLowEnergyController__Error))(signal))(newError)
				f(newError)
			}
			qt.ConnectSignal(ptr.Pointer(), "error2", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "error2", unsafe.Pointer(&f))
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

func (ptr *QLowEnergyController) ErrorString() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QLowEnergyController_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLowEnergyController) LocalAddress() *QBluetoothAddress {
	if ptr.Pointer() != nil {
		tmpValue := NewQBluetoothAddressFromPointer(C.QLowEnergyController_LocalAddress(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyController) RemoteAddress() *QBluetoothAddress {
	if ptr.Pointer() != nil {
		tmpValue := NewQBluetoothAddressFromPointer(C.QLowEnergyController_RemoteAddress(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QBluetoothAddress).DestroyQBluetoothAddress)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyController) RemoteAddressType() QLowEnergyController__RemoteAddressType {
	if ptr.Pointer() != nil {
		return QLowEnergyController__RemoteAddressType(C.QLowEnergyController_RemoteAddressType(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyController) RemoteDeviceUuid() *QBluetoothUuid {
	if ptr.Pointer() != nil {
		tmpValue := NewQBluetoothUuidFromPointer(C.QLowEnergyController_RemoteDeviceUuid(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyController) RemoteName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QLowEnergyController_RemoteName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLowEnergyController) RequestConnectionUpdate(parameters QLowEnergyConnectionParameters_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_RequestConnectionUpdate(ptr.Pointer(), PointerFromQLowEnergyConnectionParameters(parameters))
	}
}

func (ptr *QLowEnergyController) Role() QLowEnergyController__Role {
	if ptr.Pointer() != nil {
		return QLowEnergyController__Role(C.QLowEnergyController_Role(ptr.Pointer()))
	}
	return 0
}

//export callbackQLowEnergyController_ServiceDiscovered
func callbackQLowEnergyController_ServiceDiscovered(ptr unsafe.Pointer, newService unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "serviceDiscovered"); signal != nil {
		(*(*func(*QBluetoothUuid))(signal))(NewQBluetoothUuidFromPointer(newService))
	}

}

func (ptr *QLowEnergyController) ConnectServiceDiscovered(f func(newService *QBluetoothUuid)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "serviceDiscovered") {
			C.QLowEnergyController_ConnectServiceDiscovered(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "serviceDiscovered")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "serviceDiscovered"); signal != nil {
			f := func(newService *QBluetoothUuid) {
				(*(*func(*QBluetoothUuid))(signal))(newService)
				f(newService)
			}
			qt.ConnectSignal(ptr.Pointer(), "serviceDiscovered", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "serviceDiscovered", unsafe.Pointer(&f))
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

func (ptr *QLowEnergyController) Services() []*QBluetoothUuid {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtBluetooth_PackedList) []*QBluetoothUuid {
			out := make([]*QBluetoothUuid, int(l.len))
			tmpList := NewQLowEnergyControllerFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__services_atList(i)
			}
			return out
		}(C.QLowEnergyController_Services(ptr.Pointer()))
	}
	return make([]*QBluetoothUuid, 0)
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

func (ptr *QLowEnergyController) State() QLowEnergyController__ControllerState {
	if ptr.Pointer() != nil {
		return QLowEnergyController__ControllerState(C.QLowEnergyController_State(ptr.Pointer()))
	}
	return 0
}

//export callbackQLowEnergyController_StateChanged
func callbackQLowEnergyController_StateChanged(ptr unsafe.Pointer, state C.longlong) {
	if signal := qt.GetSignal(ptr, "stateChanged"); signal != nil {
		(*(*func(QLowEnergyController__ControllerState))(signal))(QLowEnergyController__ControllerState(state))
	}

}

func (ptr *QLowEnergyController) ConnectStateChanged(f func(state QLowEnergyController__ControllerState)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "stateChanged") {
			C.QLowEnergyController_ConnectStateChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "stateChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "stateChanged"); signal != nil {
			f := func(state QLowEnergyController__ControllerState) {
				(*(*func(QLowEnergyController__ControllerState))(signal))(state)
				f(state)
			}
			qt.ConnectSignal(ptr.Pointer(), "stateChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stateChanged", unsafe.Pointer(&f))
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

//export callbackQLowEnergyController_DestroyQLowEnergyController
func callbackQLowEnergyController_DestroyQLowEnergyController(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QLowEnergyController"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQLowEnergyControllerFromPointer(ptr).DestroyQLowEnergyControllerDefault()
	}
}

func (ptr *QLowEnergyController) ConnectDestroyQLowEnergyController(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QLowEnergyController"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QLowEnergyController", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QLowEnergyController", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLowEnergyController) DisconnectDestroyQLowEnergyController() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QLowEnergyController")
	}
}

func (ptr *QLowEnergyController) DestroyQLowEnergyController() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QLowEnergyController_DestroyQLowEnergyController(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QLowEnergyController) DestroyQLowEnergyControllerDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QLowEnergyController_DestroyQLowEnergyControllerDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QLowEnergyController) __services_atList(i int) *QBluetoothUuid {
	if ptr.Pointer() != nil {
		tmpValue := NewQBluetoothUuidFromPointer(C.QLowEnergyController___services_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
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
	return C.QLowEnergyController___services_newList(ptr.Pointer())
}

func (ptr *QLowEnergyController) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QLowEnergyController___children_atList(ptr.Pointer(), C.int(int32(i))))
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
	return C.QLowEnergyController___children_newList(ptr.Pointer())
}

func (ptr *QLowEnergyController) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QLowEnergyController___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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
	return C.QLowEnergyController___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QLowEnergyController) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QLowEnergyController___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
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
	return C.QLowEnergyController___findChildren_newList(ptr.Pointer())
}

func (ptr *QLowEnergyController) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QLowEnergyController___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
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
	return C.QLowEnergyController___findChildren_newList3(ptr.Pointer())
}

//export callbackQLowEnergyController_ChildEvent
func callbackQLowEnergyController_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
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
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
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
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
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
		(*(*func())(signal))()
	} else {
		NewQLowEnergyControllerFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QLowEnergyController) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QLowEnergyController_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQLowEnergyController_Destroyed
func callbackQLowEnergyController_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQLowEnergyController_DisconnectNotify
func callbackQLowEnergyController_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQLowEnergyControllerFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QLowEnergyController) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQLowEnergyController_Event
func callbackQLowEnergyController_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLowEnergyControllerFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QLowEnergyController) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QLowEnergyController_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQLowEnergyController_EventFilter
func callbackQLowEnergyController_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLowEnergyControllerFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QLowEnergyController) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QLowEnergyController_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQLowEnergyController_MetaObject
func callbackQLowEnergyController_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQLowEnergyControllerFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QLowEnergyController) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QLowEnergyController_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQLowEnergyController_ObjectNameChanged
func callbackQLowEnergyController_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtBluetooth_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQLowEnergyController_TimerEvent
func callbackQLowEnergyController_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQLowEnergyControllerFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QLowEnergyController) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyController_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
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

func NewQLowEnergyDescriptorFromPointer(ptr unsafe.Pointer) (n *QLowEnergyDescriptor) {
	n = new(QLowEnergyDescriptor)
	n.SetPointer(ptr)
	return
}
func NewQLowEnergyDescriptor() *QLowEnergyDescriptor {
	tmpValue := NewQLowEnergyDescriptorFromPointer(C.QLowEnergyDescriptor_NewQLowEnergyDescriptor())
	qt.SetFinalizer(tmpValue, (*QLowEnergyDescriptor).DestroyQLowEnergyDescriptor)
	return tmpValue
}

func NewQLowEnergyDescriptor2(other QLowEnergyDescriptor_ITF) *QLowEnergyDescriptor {
	tmpValue := NewQLowEnergyDescriptorFromPointer(C.QLowEnergyDescriptor_NewQLowEnergyDescriptor2(PointerFromQLowEnergyDescriptor(other)))
	qt.SetFinalizer(tmpValue, (*QLowEnergyDescriptor).DestroyQLowEnergyDescriptor)
	return tmpValue
}

func (ptr *QLowEnergyDescriptor) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QLowEnergyDescriptor_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLowEnergyDescriptor) Name() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QLowEnergyDescriptor_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLowEnergyDescriptor) Type() QBluetoothUuid__DescriptorType {
	if ptr.Pointer() != nil {
		return QBluetoothUuid__DescriptorType(C.QLowEnergyDescriptor_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyDescriptor) Uuid() *QBluetoothUuid {
	if ptr.Pointer() != nil {
		tmpValue := NewQBluetoothUuidFromPointer(C.QLowEnergyDescriptor_Uuid(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyDescriptor) Value() *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QLowEnergyDescriptor_Value(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyDescriptor) DestroyQLowEnergyDescriptor() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QLowEnergyDescriptor_DestroyQLowEnergyDescriptor(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
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

func NewQLowEnergyDescriptorDataFromPointer(ptr unsafe.Pointer) (n *QLowEnergyDescriptorData) {
	n = new(QLowEnergyDescriptorData)
	n.SetPointer(ptr)
	return
}
func NewQLowEnergyDescriptorData() *QLowEnergyDescriptorData {
	tmpValue := NewQLowEnergyDescriptorDataFromPointer(C.QLowEnergyDescriptorData_NewQLowEnergyDescriptorData())
	qt.SetFinalizer(tmpValue, (*QLowEnergyDescriptorData).DestroyQLowEnergyDescriptorData)
	return tmpValue
}

func NewQLowEnergyDescriptorData2(uuid QBluetoothUuid_ITF, value core.QByteArray_ITF) *QLowEnergyDescriptorData {
	tmpValue := NewQLowEnergyDescriptorDataFromPointer(C.QLowEnergyDescriptorData_NewQLowEnergyDescriptorData2(PointerFromQBluetoothUuid(uuid), core.PointerFromQByteArray(value)))
	qt.SetFinalizer(tmpValue, (*QLowEnergyDescriptorData).DestroyQLowEnergyDescriptorData)
	return tmpValue
}

func NewQLowEnergyDescriptorData3(other QLowEnergyDescriptorData_ITF) *QLowEnergyDescriptorData {
	tmpValue := NewQLowEnergyDescriptorDataFromPointer(C.QLowEnergyDescriptorData_NewQLowEnergyDescriptorData3(PointerFromQLowEnergyDescriptorData(other)))
	qt.SetFinalizer(tmpValue, (*QLowEnergyDescriptorData).DestroyQLowEnergyDescriptorData)
	return tmpValue
}

func (ptr *QLowEnergyDescriptorData) IsReadable() bool {
	if ptr.Pointer() != nil {
		return int8(C.QLowEnergyDescriptorData_IsReadable(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLowEnergyDescriptorData) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QLowEnergyDescriptorData_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLowEnergyDescriptorData) IsWritable() bool {
	if ptr.Pointer() != nil {
		return int8(C.QLowEnergyDescriptorData_IsWritable(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLowEnergyDescriptorData) ReadConstraints() QBluetooth__AttAccessConstraint {
	if ptr.Pointer() != nil {
		return QBluetooth__AttAccessConstraint(C.QLowEnergyDescriptorData_ReadConstraints(ptr.Pointer()))
	}
	return 0
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

func (ptr *QLowEnergyDescriptorData) Uuid() *QBluetoothUuid {
	if ptr.Pointer() != nil {
		tmpValue := NewQBluetoothUuidFromPointer(C.QLowEnergyDescriptorData_Uuid(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyDescriptorData) Value() *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QLowEnergyDescriptorData_Value(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyDescriptorData) WriteConstraints() QBluetooth__AttAccessConstraint {
	if ptr.Pointer() != nil {
		return QBluetooth__AttAccessConstraint(C.QLowEnergyDescriptorData_WriteConstraints(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyDescriptorData) DestroyQLowEnergyDescriptorData() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QLowEnergyDescriptorData_DestroyQLowEnergyDescriptorData(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
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

func NewQLowEnergyServiceFromPointer(ptr unsafe.Pointer) (n *QLowEnergyService) {
	n = new(QLowEnergyService)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QLowEnergyService__ServiceType
//QLowEnergyService::ServiceType
type QLowEnergyService__ServiceType int64

const (
	QLowEnergyService__PrimaryService  QLowEnergyService__ServiceType = QLowEnergyService__ServiceType(0x0001)
	QLowEnergyService__IncludedService QLowEnergyService__ServiceType = QLowEnergyService__ServiceType(0x0002)
)

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

//go:generate stringer -type=QLowEnergyService__WriteMode
//QLowEnergyService::WriteMode
type QLowEnergyService__WriteMode int64

const (
	QLowEnergyService__WriteWithResponse    QLowEnergyService__WriteMode = QLowEnergyService__WriteMode(0)
	QLowEnergyService__WriteWithoutResponse QLowEnergyService__WriteMode = QLowEnergyService__WriteMode(1)
	QLowEnergyService__WriteSigned          QLowEnergyService__WriteMode = QLowEnergyService__WriteMode(2)
)

func (ptr *QLowEnergyService) Characteristic(uuid QBluetoothUuid_ITF) *QLowEnergyCharacteristic {
	if ptr.Pointer() != nil {
		tmpValue := NewQLowEnergyCharacteristicFromPointer(C.QLowEnergyService_Characteristic(ptr.Pointer(), PointerFromQBluetoothUuid(uuid)))
		qt.SetFinalizer(tmpValue, (*QLowEnergyCharacteristic).DestroyQLowEnergyCharacteristic)
		return tmpValue
	}
	return nil
}

//export callbackQLowEnergyService_CharacteristicChanged
func callbackQLowEnergyService_CharacteristicChanged(ptr unsafe.Pointer, characteristic unsafe.Pointer, newValue unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "characteristicChanged"); signal != nil {
		(*(*func(*QLowEnergyCharacteristic, *core.QByteArray))(signal))(NewQLowEnergyCharacteristicFromPointer(characteristic), core.NewQByteArrayFromPointer(newValue))
	}

}

func (ptr *QLowEnergyService) ConnectCharacteristicChanged(f func(characteristic *QLowEnergyCharacteristic, newValue *core.QByteArray)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "characteristicChanged") {
			C.QLowEnergyService_ConnectCharacteristicChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "characteristicChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "characteristicChanged"); signal != nil {
			f := func(characteristic *QLowEnergyCharacteristic, newValue *core.QByteArray) {
				(*(*func(*QLowEnergyCharacteristic, *core.QByteArray))(signal))(characteristic, newValue)
				f(characteristic, newValue)
			}
			qt.ConnectSignal(ptr.Pointer(), "characteristicChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "characteristicChanged", unsafe.Pointer(&f))
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
		(*(*func(*QLowEnergyCharacteristic, *core.QByteArray))(signal))(NewQLowEnergyCharacteristicFromPointer(characteristic), core.NewQByteArrayFromPointer(value))
	}

}

func (ptr *QLowEnergyService) ConnectCharacteristicRead(f func(characteristic *QLowEnergyCharacteristic, value *core.QByteArray)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "characteristicRead") {
			C.QLowEnergyService_ConnectCharacteristicRead(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "characteristicRead")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "characteristicRead"); signal != nil {
			f := func(characteristic *QLowEnergyCharacteristic, value *core.QByteArray) {
				(*(*func(*QLowEnergyCharacteristic, *core.QByteArray))(signal))(characteristic, value)
				f(characteristic, value)
			}
			qt.ConnectSignal(ptr.Pointer(), "characteristicRead", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "characteristicRead", unsafe.Pointer(&f))
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
		(*(*func(*QLowEnergyCharacteristic, *core.QByteArray))(signal))(NewQLowEnergyCharacteristicFromPointer(characteristic), core.NewQByteArrayFromPointer(newValue))
	}

}

func (ptr *QLowEnergyService) ConnectCharacteristicWritten(f func(characteristic *QLowEnergyCharacteristic, newValue *core.QByteArray)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "characteristicWritten") {
			C.QLowEnergyService_ConnectCharacteristicWritten(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "characteristicWritten")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "characteristicWritten"); signal != nil {
			f := func(characteristic *QLowEnergyCharacteristic, newValue *core.QByteArray) {
				(*(*func(*QLowEnergyCharacteristic, *core.QByteArray))(signal))(characteristic, newValue)
				f(characteristic, newValue)
			}
			qt.ConnectSignal(ptr.Pointer(), "characteristicWritten", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "characteristicWritten", unsafe.Pointer(&f))
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

func (ptr *QLowEnergyService) Characteristics() []*QLowEnergyCharacteristic {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtBluetooth_PackedList) []*QLowEnergyCharacteristic {
			out := make([]*QLowEnergyCharacteristic, int(l.len))
			tmpList := NewQLowEnergyServiceFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__characteristics_atList(i)
			}
			return out
		}(C.QLowEnergyService_Characteristics(ptr.Pointer()))
	}
	return make([]*QLowEnergyCharacteristic, 0)
}

func (ptr *QLowEnergyService) Contains(characteristic QLowEnergyCharacteristic_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QLowEnergyService_Contains(ptr.Pointer(), PointerFromQLowEnergyCharacteristic(characteristic))) != 0
	}
	return false
}

func (ptr *QLowEnergyService) Contains2(descriptor QLowEnergyDescriptor_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QLowEnergyService_Contains2(ptr.Pointer(), PointerFromQLowEnergyDescriptor(descriptor))) != 0
	}
	return false
}

//export callbackQLowEnergyService_DescriptorRead
func callbackQLowEnergyService_DescriptorRead(ptr unsafe.Pointer, descriptor unsafe.Pointer, value unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "descriptorRead"); signal != nil {
		(*(*func(*QLowEnergyDescriptor, *core.QByteArray))(signal))(NewQLowEnergyDescriptorFromPointer(descriptor), core.NewQByteArrayFromPointer(value))
	}

}

func (ptr *QLowEnergyService) ConnectDescriptorRead(f func(descriptor *QLowEnergyDescriptor, value *core.QByteArray)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "descriptorRead") {
			C.QLowEnergyService_ConnectDescriptorRead(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "descriptorRead")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "descriptorRead"); signal != nil {
			f := func(descriptor *QLowEnergyDescriptor, value *core.QByteArray) {
				(*(*func(*QLowEnergyDescriptor, *core.QByteArray))(signal))(descriptor, value)
				f(descriptor, value)
			}
			qt.ConnectSignal(ptr.Pointer(), "descriptorRead", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "descriptorRead", unsafe.Pointer(&f))
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
		(*(*func(*QLowEnergyDescriptor, *core.QByteArray))(signal))(NewQLowEnergyDescriptorFromPointer(descriptor), core.NewQByteArrayFromPointer(newValue))
	}

}

func (ptr *QLowEnergyService) ConnectDescriptorWritten(f func(descriptor *QLowEnergyDescriptor, newValue *core.QByteArray)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "descriptorWritten") {
			C.QLowEnergyService_ConnectDescriptorWritten(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "descriptorWritten")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "descriptorWritten"); signal != nil {
			f := func(descriptor *QLowEnergyDescriptor, newValue *core.QByteArray) {
				(*(*func(*QLowEnergyDescriptor, *core.QByteArray))(signal))(descriptor, newValue)
				f(descriptor, newValue)
			}
			qt.ConnectSignal(ptr.Pointer(), "descriptorWritten", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "descriptorWritten", unsafe.Pointer(&f))
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

func (ptr *QLowEnergyService) Error() QLowEnergyService__ServiceError {
	if ptr.Pointer() != nil {
		return QLowEnergyService__ServiceError(C.QLowEnergyService_Error(ptr.Pointer()))
	}
	return 0
}

//export callbackQLowEnergyService_Error2
func callbackQLowEnergyService_Error2(ptr unsafe.Pointer, newError C.longlong) {
	if signal := qt.GetSignal(ptr, "error2"); signal != nil {
		(*(*func(QLowEnergyService__ServiceError))(signal))(QLowEnergyService__ServiceError(newError))
	}

}

func (ptr *QLowEnergyService) ConnectError2(f func(newError QLowEnergyService__ServiceError)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "error2") {
			C.QLowEnergyService_ConnectError2(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "error")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "error2"); signal != nil {
			f := func(newError QLowEnergyService__ServiceError) {
				(*(*func(QLowEnergyService__ServiceError))(signal))(newError)
				f(newError)
			}
			qt.ConnectSignal(ptr.Pointer(), "error2", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "error2", unsafe.Pointer(&f))
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

func (ptr *QLowEnergyService) IncludedServices() []*QBluetoothUuid {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtBluetooth_PackedList) []*QBluetoothUuid {
			out := make([]*QBluetoothUuid, int(l.len))
			tmpList := NewQLowEnergyServiceFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__includedServices_atList(i)
			}
			return out
		}(C.QLowEnergyService_IncludedServices(ptr.Pointer()))
	}
	return make([]*QBluetoothUuid, 0)
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

func (ptr *QLowEnergyService) ServiceName() string {
	if ptr.Pointer() != nil {
		return cGoUnpackString(C.QLowEnergyService_ServiceName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QLowEnergyService) ServiceUuid() *QBluetoothUuid {
	if ptr.Pointer() != nil {
		tmpValue := NewQBluetoothUuidFromPointer(C.QLowEnergyService_ServiceUuid(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
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

//export callbackQLowEnergyService_StateChanged
func callbackQLowEnergyService_StateChanged(ptr unsafe.Pointer, newState C.longlong) {
	if signal := qt.GetSignal(ptr, "stateChanged"); signal != nil {
		(*(*func(QLowEnergyService__ServiceState))(signal))(QLowEnergyService__ServiceState(newState))
	}

}

func (ptr *QLowEnergyService) ConnectStateChanged(f func(newState QLowEnergyService__ServiceState)) {
	if ptr.Pointer() != nil {

		if !qt.ExistsSignal(ptr.Pointer(), "stateChanged") {
			C.QLowEnergyService_ConnectStateChanged(ptr.Pointer(), C.longlong(qt.ConnectionType(ptr.Pointer(), "stateChanged")))
		}

		if signal := qt.LendSignal(ptr.Pointer(), "stateChanged"); signal != nil {
			f := func(newState QLowEnergyService__ServiceState) {
				(*(*func(QLowEnergyService__ServiceState))(signal))(newState)
				f(newState)
			}
			qt.ConnectSignal(ptr.Pointer(), "stateChanged", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "stateChanged", unsafe.Pointer(&f))
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

func (ptr *QLowEnergyService) Type() QLowEnergyService__ServiceType {
	if ptr.Pointer() != nil {
		return QLowEnergyService__ServiceType(C.QLowEnergyService_Type(ptr.Pointer()))
	}
	return 0
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

//export callbackQLowEnergyService_DestroyQLowEnergyService
func callbackQLowEnergyService_DestroyQLowEnergyService(ptr unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "~QLowEnergyService"); signal != nil {
		(*(*func())(signal))()
	} else {
		NewQLowEnergyServiceFromPointer(ptr).DestroyQLowEnergyServiceDefault()
	}
}

func (ptr *QLowEnergyService) ConnectDestroyQLowEnergyService(f func()) {
	if ptr.Pointer() != nil {

		if signal := qt.LendSignal(ptr.Pointer(), "~QLowEnergyService"); signal != nil {
			f := func() {
				(*(*func())(signal))()
				f()
			}
			qt.ConnectSignal(ptr.Pointer(), "~QLowEnergyService", unsafe.Pointer(&f))
		} else {
			qt.ConnectSignal(ptr.Pointer(), "~QLowEnergyService", unsafe.Pointer(&f))
		}
	}
}

func (ptr *QLowEnergyService) DisconnectDestroyQLowEnergyService() {
	if ptr.Pointer() != nil {

		qt.DisconnectSignal(ptr.Pointer(), "~QLowEnergyService")
	}
}

func (ptr *QLowEnergyService) DestroyQLowEnergyService() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QLowEnergyService_DestroyQLowEnergyService(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QLowEnergyService) DestroyQLowEnergyServiceDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QLowEnergyService_DestroyQLowEnergyServiceDefault(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QLowEnergyService) __characteristics_atList(i int) *QLowEnergyCharacteristic {
	if ptr.Pointer() != nil {
		tmpValue := NewQLowEnergyCharacteristicFromPointer(C.QLowEnergyService___characteristics_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QLowEnergyCharacteristic).DestroyQLowEnergyCharacteristic)
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
	return C.QLowEnergyService___characteristics_newList(ptr.Pointer())
}

func (ptr *QLowEnergyService) __includedServices_atList(i int) *QBluetoothUuid {
	if ptr.Pointer() != nil {
		tmpValue := NewQBluetoothUuidFromPointer(C.QLowEnergyService___includedServices_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
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
	return C.QLowEnergyService___includedServices_newList(ptr.Pointer())
}

func (ptr *QLowEnergyService) __children_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QLowEnergyService___children_atList(ptr.Pointer(), C.int(int32(i))))
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
	return C.QLowEnergyService___children_newList(ptr.Pointer())
}

func (ptr *QLowEnergyService) __dynamicPropertyNames_atList(i int) *core.QByteArray {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQByteArrayFromPointer(C.QLowEnergyService___dynamicPropertyNames_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*core.QByteArray).DestroyQByteArray)
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
	return C.QLowEnergyService___dynamicPropertyNames_newList(ptr.Pointer())
}

func (ptr *QLowEnergyService) __findChildren_atList(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QLowEnergyService___findChildren_atList(ptr.Pointer(), C.int(int32(i))))
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
	return C.QLowEnergyService___findChildren_newList(ptr.Pointer())
}

func (ptr *QLowEnergyService) __findChildren_atList3(i int) *core.QObject {
	if ptr.Pointer() != nil {
		tmpValue := core.NewQObjectFromPointer(C.QLowEnergyService___findChildren_atList3(ptr.Pointer(), C.int(int32(i))))
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
	return C.QLowEnergyService___findChildren_newList3(ptr.Pointer())
}

//export callbackQLowEnergyService_ChildEvent
func callbackQLowEnergyService_ChildEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "childEvent"); signal != nil {
		(*(*func(*core.QChildEvent))(signal))(core.NewQChildEventFromPointer(event))
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
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
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
		(*(*func(*core.QEvent))(signal))(core.NewQEventFromPointer(event))
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
		(*(*func())(signal))()
	} else {
		NewQLowEnergyServiceFromPointer(ptr).DeleteLaterDefault()
	}
}

func (ptr *QLowEnergyService) DeleteLaterDefault() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QLowEnergyService_DeleteLaterDefault(ptr.Pointer())
	}
}

//export callbackQLowEnergyService_Destroyed
func callbackQLowEnergyService_Destroyed(ptr unsafe.Pointer, obj unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "destroyed"); signal != nil {
		(*(*func(*core.QObject))(signal))(core.NewQObjectFromPointer(obj))
	}

}

//export callbackQLowEnergyService_DisconnectNotify
func callbackQLowEnergyService_DisconnectNotify(ptr unsafe.Pointer, sign unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "disconnectNotify"); signal != nil {
		(*(*func(*core.QMetaMethod))(signal))(core.NewQMetaMethodFromPointer(sign))
	} else {
		NewQLowEnergyServiceFromPointer(ptr).DisconnectNotifyDefault(core.NewQMetaMethodFromPointer(sign))
	}
}

func (ptr *QLowEnergyService) DisconnectNotifyDefault(sign core.QMetaMethod_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_DisconnectNotifyDefault(ptr.Pointer(), core.PointerFromQMetaMethod(sign))
	}
}

//export callbackQLowEnergyService_Event
func callbackQLowEnergyService_Event(ptr unsafe.Pointer, e unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "event"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QEvent) bool)(signal))(core.NewQEventFromPointer(e)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLowEnergyServiceFromPointer(ptr).EventDefault(core.NewQEventFromPointer(e)))))
}

func (ptr *QLowEnergyService) EventDefault(e core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QLowEnergyService_EventDefault(ptr.Pointer(), core.PointerFromQEvent(e))) != 0
	}
	return false
}

//export callbackQLowEnergyService_EventFilter
func callbackQLowEnergyService_EventFilter(ptr unsafe.Pointer, watched unsafe.Pointer, event unsafe.Pointer) C.char {
	if signal := qt.GetSignal(ptr, "eventFilter"); signal != nil {
		return C.char(int8(qt.GoBoolToInt((*(*func(*core.QObject, *core.QEvent) bool)(signal))(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
	}

	return C.char(int8(qt.GoBoolToInt(NewQLowEnergyServiceFromPointer(ptr).EventFilterDefault(core.NewQObjectFromPointer(watched), core.NewQEventFromPointer(event)))))
}

func (ptr *QLowEnergyService) EventFilterDefault(watched core.QObject_ITF, event core.QEvent_ITF) bool {
	if ptr.Pointer() != nil {
		return int8(C.QLowEnergyService_EventFilterDefault(ptr.Pointer(), core.PointerFromQObject(watched), core.PointerFromQEvent(event))) != 0
	}
	return false
}

//export callbackQLowEnergyService_MetaObject
func callbackQLowEnergyService_MetaObject(ptr unsafe.Pointer) unsafe.Pointer {
	if signal := qt.GetSignal(ptr, "metaObject"); signal != nil {
		return core.PointerFromQMetaObject((*(*func() *core.QMetaObject)(signal))())
	}

	return core.PointerFromQMetaObject(NewQLowEnergyServiceFromPointer(ptr).MetaObjectDefault())
}

func (ptr *QLowEnergyService) MetaObjectDefault() *core.QMetaObject {
	if ptr.Pointer() != nil {
		return core.NewQMetaObjectFromPointer(C.QLowEnergyService_MetaObjectDefault(ptr.Pointer()))
	}
	return nil
}

//export callbackQLowEnergyService_ObjectNameChanged
func callbackQLowEnergyService_ObjectNameChanged(ptr unsafe.Pointer, objectName C.struct_QtBluetooth_PackedString) {
	if signal := qt.GetSignal(ptr, "objectNameChanged"); signal != nil {
		(*(*func(string))(signal))(cGoUnpackString(objectName))
	}

}

//export callbackQLowEnergyService_TimerEvent
func callbackQLowEnergyService_TimerEvent(ptr unsafe.Pointer, event unsafe.Pointer) {
	if signal := qt.GetSignal(ptr, "timerEvent"); signal != nil {
		(*(*func(*core.QTimerEvent))(signal))(core.NewQTimerEventFromPointer(event))
	} else {
		NewQLowEnergyServiceFromPointer(ptr).TimerEventDefault(core.NewQTimerEventFromPointer(event))
	}
}

func (ptr *QLowEnergyService) TimerEventDefault(event core.QTimerEvent_ITF) {
	if ptr.Pointer() != nil {
		C.QLowEnergyService_TimerEventDefault(ptr.Pointer(), core.PointerFromQTimerEvent(event))
	}
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

func NewQLowEnergyServiceDataFromPointer(ptr unsafe.Pointer) (n *QLowEnergyServiceData) {
	n = new(QLowEnergyServiceData)
	n.SetPointer(ptr)
	return
}

//go:generate stringer -type=QLowEnergyServiceData__ServiceType
//QLowEnergyServiceData::ServiceType
type QLowEnergyServiceData__ServiceType int64

const (
	QLowEnergyServiceData__ServiceTypePrimary   QLowEnergyServiceData__ServiceType = QLowEnergyServiceData__ServiceType(0x2800)
	QLowEnergyServiceData__ServiceTypeSecondary QLowEnergyServiceData__ServiceType = QLowEnergyServiceData__ServiceType(0x2801)
)

func NewQLowEnergyServiceData() *QLowEnergyServiceData {
	tmpValue := NewQLowEnergyServiceDataFromPointer(C.QLowEnergyServiceData_NewQLowEnergyServiceData())
	qt.SetFinalizer(tmpValue, (*QLowEnergyServiceData).DestroyQLowEnergyServiceData)
	return tmpValue
}

func NewQLowEnergyServiceData2(other QLowEnergyServiceData_ITF) *QLowEnergyServiceData {
	tmpValue := NewQLowEnergyServiceDataFromPointer(C.QLowEnergyServiceData_NewQLowEnergyServiceData2(PointerFromQLowEnergyServiceData(other)))
	qt.SetFinalizer(tmpValue, (*QLowEnergyServiceData).DestroyQLowEnergyServiceData)
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

func (ptr *QLowEnergyServiceData) Characteristics() []*QLowEnergyCharacteristicData {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtBluetooth_PackedList) []*QLowEnergyCharacteristicData {
			out := make([]*QLowEnergyCharacteristicData, int(l.len))
			tmpList := NewQLowEnergyServiceDataFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__characteristics_atList(i)
			}
			return out
		}(C.QLowEnergyServiceData_Characteristics(ptr.Pointer()))
	}
	return make([]*QLowEnergyCharacteristicData, 0)
}

func (ptr *QLowEnergyServiceData) IncludedServices() []*QLowEnergyService {
	if ptr.Pointer() != nil {
		return func(l C.struct_QtBluetooth_PackedList) []*QLowEnergyService {
			out := make([]*QLowEnergyService, int(l.len))
			tmpList := NewQLowEnergyServiceDataFromPointer(l.data)
			for i := 0; i < len(out); i++ {
				out[i] = tmpList.__includedServices_atList(i)
			}
			return out
		}(C.QLowEnergyServiceData_IncludedServices(ptr.Pointer()))
	}
	return make([]*QLowEnergyService, 0)
}

func (ptr *QLowEnergyServiceData) IsValid() bool {
	if ptr.Pointer() != nil {
		return int8(C.QLowEnergyServiceData_IsValid(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QLowEnergyServiceData) SetCharacteristics(characteristics []*QLowEnergyCharacteristicData) {
	if ptr.Pointer() != nil {
		C.QLowEnergyServiceData_SetCharacteristics(ptr.Pointer(), func() unsafe.Pointer {
			tmpList := NewQLowEnergyServiceDataFromPointer(NewQLowEnergyServiceDataFromPointer(nil).__setCharacteristics_characteristics_newList())
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
			tmpList := NewQLowEnergyServiceDataFromPointer(NewQLowEnergyServiceDataFromPointer(nil).__setIncludedServices_services_newList())
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

func (ptr *QLowEnergyServiceData) Type() QLowEnergyServiceData__ServiceType {
	if ptr.Pointer() != nil {
		return QLowEnergyServiceData__ServiceType(C.QLowEnergyServiceData_Type(ptr.Pointer()))
	}
	return 0
}

func (ptr *QLowEnergyServiceData) Uuid() *QBluetoothUuid {
	if ptr.Pointer() != nil {
		tmpValue := NewQBluetoothUuidFromPointer(C.QLowEnergyServiceData_Uuid(ptr.Pointer()))
		qt.SetFinalizer(tmpValue, (*QBluetoothUuid).DestroyQBluetoothUuid)
		return tmpValue
	}
	return nil
}

func (ptr *QLowEnergyServiceData) DestroyQLowEnergyServiceData() {
	if ptr.Pointer() != nil {

		qt.SetFinalizer(ptr, nil)
		C.QLowEnergyServiceData_DestroyQLowEnergyServiceData(ptr.Pointer())
		C.free(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}

func (ptr *QLowEnergyServiceData) __characteristics_atList(i int) *QLowEnergyCharacteristicData {
	if ptr.Pointer() != nil {
		tmpValue := NewQLowEnergyCharacteristicDataFromPointer(C.QLowEnergyServiceData___characteristics_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QLowEnergyCharacteristicData).DestroyQLowEnergyCharacteristicData)
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
	return C.QLowEnergyServiceData___characteristics_newList(ptr.Pointer())
}

func (ptr *QLowEnergyServiceData) __includedServices_atList(i int) *QLowEnergyService {
	if ptr.Pointer() != nil {
		tmpValue := NewQLowEnergyServiceFromPointer(C.QLowEnergyServiceData___includedServices_atList(ptr.Pointer(), C.int(int32(i))))
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
	return C.QLowEnergyServiceData___includedServices_newList(ptr.Pointer())
}

func (ptr *QLowEnergyServiceData) __setCharacteristics_characteristics_atList(i int) *QLowEnergyCharacteristicData {
	if ptr.Pointer() != nil {
		tmpValue := NewQLowEnergyCharacteristicDataFromPointer(C.QLowEnergyServiceData___setCharacteristics_characteristics_atList(ptr.Pointer(), C.int(int32(i))))
		qt.SetFinalizer(tmpValue, (*QLowEnergyCharacteristicData).DestroyQLowEnergyCharacteristicData)
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
	return C.QLowEnergyServiceData___setCharacteristics_characteristics_newList(ptr.Pointer())
}

func (ptr *QLowEnergyServiceData) __setIncludedServices_services_atList(i int) *QLowEnergyService {
	if ptr.Pointer() != nil {
		tmpValue := NewQLowEnergyServiceFromPointer(C.QLowEnergyServiceData___setIncludedServices_services_atList(ptr.Pointer(), C.int(int32(i))))
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
	return C.QLowEnergyServiceData___setIncludedServices_services_newList(ptr.Pointer())
}

func init() {
	qt.ItfMap["bluetooth.QBluetooth_ITF"] = QBluetooth{}
	qt.EnumMap["bluetooth.QBluetooth__AttAuthorizationRequired"] = int64(QBluetooth__AttAuthorizationRequired)
	qt.EnumMap["bluetooth.QBluetooth__AttAuthenticationRequired"] = int64(QBluetooth__AttAuthenticationRequired)
	qt.EnumMap["bluetooth.QBluetooth__AttEncryptionRequired"] = int64(QBluetooth__AttEncryptionRequired)
	qt.EnumMap["bluetooth.QBluetooth__NoSecurity"] = int64(QBluetooth__NoSecurity)
	qt.EnumMap["bluetooth.QBluetooth__Authorization"] = int64(QBluetooth__Authorization)
	qt.EnumMap["bluetooth.QBluetooth__Authentication"] = int64(QBluetooth__Authentication)
	qt.EnumMap["bluetooth.QBluetooth__Encryption"] = int64(QBluetooth__Encryption)
	qt.EnumMap["bluetooth.QBluetooth__Secure"] = int64(QBluetooth__Secure)
	qt.ItfMap["bluetooth.QBluetoothAddress_ITF"] = QBluetoothAddress{}
	qt.FuncMap["bluetooth.NewQBluetoothAddress"] = NewQBluetoothAddress
	qt.FuncMap["bluetooth.NewQBluetoothAddress2"] = NewQBluetoothAddress2
	qt.FuncMap["bluetooth.NewQBluetoothAddress3"] = NewQBluetoothAddress3
	qt.FuncMap["bluetooth.NewQBluetoothAddress4"] = NewQBluetoothAddress4
	qt.ItfMap["bluetooth.QBluetoothDeviceDiscoveryAgent_ITF"] = QBluetoothDeviceDiscoveryAgent{}
	qt.FuncMap["bluetooth.NewQBluetoothDeviceDiscoveryAgent"] = NewQBluetoothDeviceDiscoveryAgent
	qt.FuncMap["bluetooth.NewQBluetoothDeviceDiscoveryAgent2"] = NewQBluetoothDeviceDiscoveryAgent2
	qt.FuncMap["bluetooth.QBluetoothDeviceDiscoveryAgent_SupportedDiscoveryMethods"] = QBluetoothDeviceDiscoveryAgent_SupportedDiscoveryMethods
	qt.EnumMap["bluetooth.QBluetoothDeviceDiscoveryAgent__NoError"] = int64(QBluetoothDeviceDiscoveryAgent__NoError)
	qt.EnumMap["bluetooth.QBluetoothDeviceDiscoveryAgent__InputOutputError"] = int64(QBluetoothDeviceDiscoveryAgent__InputOutputError)
	qt.EnumMap["bluetooth.QBluetoothDeviceDiscoveryAgent__PoweredOffError"] = int64(QBluetoothDeviceDiscoveryAgent__PoweredOffError)
	qt.EnumMap["bluetooth.QBluetoothDeviceDiscoveryAgent__InvalidBluetoothAdapterError"] = int64(QBluetoothDeviceDiscoveryAgent__InvalidBluetoothAdapterError)
	qt.EnumMap["bluetooth.QBluetoothDeviceDiscoveryAgent__UnsupportedPlatformError"] = int64(QBluetoothDeviceDiscoveryAgent__UnsupportedPlatformError)
	qt.EnumMap["bluetooth.QBluetoothDeviceDiscoveryAgent__UnsupportedDiscoveryMethod"] = int64(QBluetoothDeviceDiscoveryAgent__UnsupportedDiscoveryMethod)
	qt.EnumMap["bluetooth.QBluetoothDeviceDiscoveryAgent__UnknownError"] = int64(QBluetoothDeviceDiscoveryAgent__UnknownError)
	qt.EnumMap["bluetooth.QBluetoothDeviceDiscoveryAgent__GeneralUnlimitedInquiry"] = int64(QBluetoothDeviceDiscoveryAgent__GeneralUnlimitedInquiry)
	qt.EnumMap["bluetooth.QBluetoothDeviceDiscoveryAgent__LimitedInquiry"] = int64(QBluetoothDeviceDiscoveryAgent__LimitedInquiry)
	qt.EnumMap["bluetooth.QBluetoothDeviceDiscoveryAgent__NoMethod"] = int64(QBluetoothDeviceDiscoveryAgent__NoMethod)
	qt.EnumMap["bluetooth.QBluetoothDeviceDiscoveryAgent__ClassicMethod"] = int64(QBluetoothDeviceDiscoveryAgent__ClassicMethod)
	qt.EnumMap["bluetooth.QBluetoothDeviceDiscoveryAgent__LowEnergyMethod"] = int64(QBluetoothDeviceDiscoveryAgent__LowEnergyMethod)
	qt.ItfMap["bluetooth.QBluetoothDeviceInfo_ITF"] = QBluetoothDeviceInfo{}
	qt.FuncMap["bluetooth.NewQBluetoothDeviceInfo"] = NewQBluetoothDeviceInfo
	qt.FuncMap["bluetooth.NewQBluetoothDeviceInfo2"] = NewQBluetoothDeviceInfo2
	qt.FuncMap["bluetooth.NewQBluetoothDeviceInfo3"] = NewQBluetoothDeviceInfo3
	qt.FuncMap["bluetooth.NewQBluetoothDeviceInfo4"] = NewQBluetoothDeviceInfo4
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__MiscellaneousDevice"] = int64(QBluetoothDeviceInfo__MiscellaneousDevice)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__ComputerDevice"] = int64(QBluetoothDeviceInfo__ComputerDevice)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__PhoneDevice"] = int64(QBluetoothDeviceInfo__PhoneDevice)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__LANAccessDevice"] = int64(QBluetoothDeviceInfo__LANAccessDevice)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__NetworkDevice"] = int64(QBluetoothDeviceInfo__NetworkDevice)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__AudioVideoDevice"] = int64(QBluetoothDeviceInfo__AudioVideoDevice)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__PeripheralDevice"] = int64(QBluetoothDeviceInfo__PeripheralDevice)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__ImagingDevice"] = int64(QBluetoothDeviceInfo__ImagingDevice)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__WearableDevice"] = int64(QBluetoothDeviceInfo__WearableDevice)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__ToyDevice"] = int64(QBluetoothDeviceInfo__ToyDevice)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__HealthDevice"] = int64(QBluetoothDeviceInfo__HealthDevice)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__UncategorizedDevice"] = int64(QBluetoothDeviceInfo__UncategorizedDevice)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__UncategorizedMiscellaneous"] = int64(QBluetoothDeviceInfo__UncategorizedMiscellaneous)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__UncategorizedComputer"] = int64(QBluetoothDeviceInfo__UncategorizedComputer)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__DesktopComputer"] = int64(QBluetoothDeviceInfo__DesktopComputer)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__ServerComputer"] = int64(QBluetoothDeviceInfo__ServerComputer)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__LaptopComputer"] = int64(QBluetoothDeviceInfo__LaptopComputer)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__HandheldClamShellComputer"] = int64(QBluetoothDeviceInfo__HandheldClamShellComputer)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__HandheldComputer"] = int64(QBluetoothDeviceInfo__HandheldComputer)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__WearableComputer"] = int64(QBluetoothDeviceInfo__WearableComputer)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__UncategorizedPhone"] = int64(QBluetoothDeviceInfo__UncategorizedPhone)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__CellularPhone"] = int64(QBluetoothDeviceInfo__CellularPhone)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__CordlessPhone"] = int64(QBluetoothDeviceInfo__CordlessPhone)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__SmartPhone"] = int64(QBluetoothDeviceInfo__SmartPhone)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__WiredModemOrVoiceGatewayPhone"] = int64(QBluetoothDeviceInfo__WiredModemOrVoiceGatewayPhone)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__CommonIsdnAccessPhone"] = int64(QBluetoothDeviceInfo__CommonIsdnAccessPhone)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__NetworkFullService"] = int64(QBluetoothDeviceInfo__NetworkFullService)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__NetworkLoadFactorOne"] = int64(QBluetoothDeviceInfo__NetworkLoadFactorOne)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__NetworkLoadFactorTwo"] = int64(QBluetoothDeviceInfo__NetworkLoadFactorTwo)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__NetworkLoadFactorThree"] = int64(QBluetoothDeviceInfo__NetworkLoadFactorThree)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__NetworkLoadFactorFour"] = int64(QBluetoothDeviceInfo__NetworkLoadFactorFour)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__NetworkLoadFactorFive"] = int64(QBluetoothDeviceInfo__NetworkLoadFactorFive)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__NetworkLoadFactorSix"] = int64(QBluetoothDeviceInfo__NetworkLoadFactorSix)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__NetworkNoService"] = int64(QBluetoothDeviceInfo__NetworkNoService)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__UncategorizedAudioVideoDevice"] = int64(QBluetoothDeviceInfo__UncategorizedAudioVideoDevice)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__WearableHeadsetDevice"] = int64(QBluetoothDeviceInfo__WearableHeadsetDevice)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__HandsFreeDevice"] = int64(QBluetoothDeviceInfo__HandsFreeDevice)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__Microphone"] = int64(QBluetoothDeviceInfo__Microphone)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__Loudspeaker"] = int64(QBluetoothDeviceInfo__Loudspeaker)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__Headphones"] = int64(QBluetoothDeviceInfo__Headphones)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__PortableAudioDevice"] = int64(QBluetoothDeviceInfo__PortableAudioDevice)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__CarAudio"] = int64(QBluetoothDeviceInfo__CarAudio)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__SetTopBox"] = int64(QBluetoothDeviceInfo__SetTopBox)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__HiFiAudioDevice"] = int64(QBluetoothDeviceInfo__HiFiAudioDevice)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__Vcr"] = int64(QBluetoothDeviceInfo__Vcr)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__VideoCamera"] = int64(QBluetoothDeviceInfo__VideoCamera)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__Camcorder"] = int64(QBluetoothDeviceInfo__Camcorder)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__VideoMonitor"] = int64(QBluetoothDeviceInfo__VideoMonitor)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__VideoDisplayAndLoudspeaker"] = int64(QBluetoothDeviceInfo__VideoDisplayAndLoudspeaker)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__VideoConferencing"] = int64(QBluetoothDeviceInfo__VideoConferencing)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__GamingDevice"] = int64(QBluetoothDeviceInfo__GamingDevice)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__UncategorizedPeripheral"] = int64(QBluetoothDeviceInfo__UncategorizedPeripheral)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__KeyboardPeripheral"] = int64(QBluetoothDeviceInfo__KeyboardPeripheral)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__PointingDevicePeripheral"] = int64(QBluetoothDeviceInfo__PointingDevicePeripheral)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__KeyboardWithPointingDevicePeripheral"] = int64(QBluetoothDeviceInfo__KeyboardWithPointingDevicePeripheral)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__JoystickPeripheral"] = int64(QBluetoothDeviceInfo__JoystickPeripheral)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__GamepadPeripheral"] = int64(QBluetoothDeviceInfo__GamepadPeripheral)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__RemoteControlPeripheral"] = int64(QBluetoothDeviceInfo__RemoteControlPeripheral)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__SensingDevicePeripheral"] = int64(QBluetoothDeviceInfo__SensingDevicePeripheral)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__DigitizerTabletPeripheral"] = int64(QBluetoothDeviceInfo__DigitizerTabletPeripheral)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__CardReaderPeripheral"] = int64(QBluetoothDeviceInfo__CardReaderPeripheral)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__UncategorizedImagingDevice"] = int64(QBluetoothDeviceInfo__UncategorizedImagingDevice)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__ImageDisplay"] = int64(QBluetoothDeviceInfo__ImageDisplay)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__ImageCamera"] = int64(QBluetoothDeviceInfo__ImageCamera)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__ImageScanner"] = int64(QBluetoothDeviceInfo__ImageScanner)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__ImagePrinter"] = int64(QBluetoothDeviceInfo__ImagePrinter)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__UncategorizedWearableDevice"] = int64(QBluetoothDeviceInfo__UncategorizedWearableDevice)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__WearableWristWatch"] = int64(QBluetoothDeviceInfo__WearableWristWatch)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__WearablePager"] = int64(QBluetoothDeviceInfo__WearablePager)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__WearableJacket"] = int64(QBluetoothDeviceInfo__WearableJacket)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__WearableHelmet"] = int64(QBluetoothDeviceInfo__WearableHelmet)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__WearableGlasses"] = int64(QBluetoothDeviceInfo__WearableGlasses)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__UncategorizedToy"] = int64(QBluetoothDeviceInfo__UncategorizedToy)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__ToyRobot"] = int64(QBluetoothDeviceInfo__ToyRobot)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__ToyVehicle"] = int64(QBluetoothDeviceInfo__ToyVehicle)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__ToyDoll"] = int64(QBluetoothDeviceInfo__ToyDoll)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__ToyController"] = int64(QBluetoothDeviceInfo__ToyController)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__ToyGame"] = int64(QBluetoothDeviceInfo__ToyGame)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__UncategorizedHealthDevice"] = int64(QBluetoothDeviceInfo__UncategorizedHealthDevice)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__HealthBloodPressureMonitor"] = int64(QBluetoothDeviceInfo__HealthBloodPressureMonitor)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__HealthThermometer"] = int64(QBluetoothDeviceInfo__HealthThermometer)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__HealthWeightScale"] = int64(QBluetoothDeviceInfo__HealthWeightScale)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__HealthGlucoseMeter"] = int64(QBluetoothDeviceInfo__HealthGlucoseMeter)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__HealthPulseOximeter"] = int64(QBluetoothDeviceInfo__HealthPulseOximeter)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__HealthDataDisplay"] = int64(QBluetoothDeviceInfo__HealthDataDisplay)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__HealthStepCounter"] = int64(QBluetoothDeviceInfo__HealthStepCounter)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__NoService"] = int64(QBluetoothDeviceInfo__NoService)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__PositioningService"] = int64(QBluetoothDeviceInfo__PositioningService)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__NetworkingService"] = int64(QBluetoothDeviceInfo__NetworkingService)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__RenderingService"] = int64(QBluetoothDeviceInfo__RenderingService)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__CapturingService"] = int64(QBluetoothDeviceInfo__CapturingService)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__ObjectTransferService"] = int64(QBluetoothDeviceInfo__ObjectTransferService)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__AudioService"] = int64(QBluetoothDeviceInfo__AudioService)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__TelephonyService"] = int64(QBluetoothDeviceInfo__TelephonyService)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__InformationService"] = int64(QBluetoothDeviceInfo__InformationService)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__AllServices"] = int64(QBluetoothDeviceInfo__AllServices)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__DataComplete"] = int64(QBluetoothDeviceInfo__DataComplete)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__DataIncomplete"] = int64(QBluetoothDeviceInfo__DataIncomplete)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__DataUnavailable"] = int64(QBluetoothDeviceInfo__DataUnavailable)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__None"] = int64(QBluetoothDeviceInfo__None)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__RSSI"] = int64(QBluetoothDeviceInfo__RSSI)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__ManufacturerData"] = int64(QBluetoothDeviceInfo__ManufacturerData)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__All"] = int64(QBluetoothDeviceInfo__All)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__UnknownCoreConfiguration"] = int64(QBluetoothDeviceInfo__UnknownCoreConfiguration)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__LowEnergyCoreConfiguration"] = int64(QBluetoothDeviceInfo__LowEnergyCoreConfiguration)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__BaseRateCoreConfiguration"] = int64(QBluetoothDeviceInfo__BaseRateCoreConfiguration)
	qt.EnumMap["bluetooth.QBluetoothDeviceInfo__BaseRateAndLowEnergyCoreConfiguration"] = int64(QBluetoothDeviceInfo__BaseRateAndLowEnergyCoreConfiguration)
	qt.ItfMap["bluetooth.QBluetoothHostInfo_ITF"] = QBluetoothHostInfo{}
	qt.FuncMap["bluetooth.NewQBluetoothHostInfo"] = NewQBluetoothHostInfo
	qt.FuncMap["bluetooth.NewQBluetoothHostInfo2"] = NewQBluetoothHostInfo2
	qt.ItfMap["bluetooth.QBluetoothLocalDevice_ITF"] = QBluetoothLocalDevice{}
	qt.FuncMap["bluetooth.NewQBluetoothLocalDevice"] = NewQBluetoothLocalDevice
	qt.FuncMap["bluetooth.NewQBluetoothLocalDevice2"] = NewQBluetoothLocalDevice2
	qt.FuncMap["bluetooth.QBluetoothLocalDevice_AllDevices"] = QBluetoothLocalDevice_AllDevices
	qt.EnumMap["bluetooth.QBluetoothLocalDevice__Unpaired"] = int64(QBluetoothLocalDevice__Unpaired)
	qt.EnumMap["bluetooth.QBluetoothLocalDevice__Paired"] = int64(QBluetoothLocalDevice__Paired)
	qt.EnumMap["bluetooth.QBluetoothLocalDevice__AuthorizedPaired"] = int64(QBluetoothLocalDevice__AuthorizedPaired)
	qt.EnumMap["bluetooth.QBluetoothLocalDevice__HostPoweredOff"] = int64(QBluetoothLocalDevice__HostPoweredOff)
	qt.EnumMap["bluetooth.QBluetoothLocalDevice__HostConnectable"] = int64(QBluetoothLocalDevice__HostConnectable)
	qt.EnumMap["bluetooth.QBluetoothLocalDevice__HostDiscoverable"] = int64(QBluetoothLocalDevice__HostDiscoverable)
	qt.EnumMap["bluetooth.QBluetoothLocalDevice__HostDiscoverableLimitedInquiry"] = int64(QBluetoothLocalDevice__HostDiscoverableLimitedInquiry)
	qt.EnumMap["bluetooth.QBluetoothLocalDevice__NoError"] = int64(QBluetoothLocalDevice__NoError)
	qt.EnumMap["bluetooth.QBluetoothLocalDevice__PairingError"] = int64(QBluetoothLocalDevice__PairingError)
	qt.EnumMap["bluetooth.QBluetoothLocalDevice__UnknownError"] = int64(QBluetoothLocalDevice__UnknownError)
	qt.ItfMap["bluetooth.QBluetoothServer_ITF"] = QBluetoothServer{}
	qt.FuncMap["bluetooth.NewQBluetoothServer"] = NewQBluetoothServer
	qt.EnumMap["bluetooth.QBluetoothServer__NoError"] = int64(QBluetoothServer__NoError)
	qt.EnumMap["bluetooth.QBluetoothServer__UnknownError"] = int64(QBluetoothServer__UnknownError)
	qt.EnumMap["bluetooth.QBluetoothServer__PoweredOffError"] = int64(QBluetoothServer__PoweredOffError)
	qt.EnumMap["bluetooth.QBluetoothServer__InputOutputError"] = int64(QBluetoothServer__InputOutputError)
	qt.EnumMap["bluetooth.QBluetoothServer__ServiceAlreadyRegisteredError"] = int64(QBluetoothServer__ServiceAlreadyRegisteredError)
	qt.EnumMap["bluetooth.QBluetoothServer__UnsupportedProtocolError"] = int64(QBluetoothServer__UnsupportedProtocolError)
	qt.ItfMap["bluetooth.QBluetoothServiceDiscoveryAgent_ITF"] = QBluetoothServiceDiscoveryAgent{}
	qt.FuncMap["bluetooth.NewQBluetoothServiceDiscoveryAgent"] = NewQBluetoothServiceDiscoveryAgent
	qt.FuncMap["bluetooth.NewQBluetoothServiceDiscoveryAgent2"] = NewQBluetoothServiceDiscoveryAgent2
	qt.EnumMap["bluetooth.QBluetoothServiceDiscoveryAgent__NoError"] = int64(QBluetoothServiceDiscoveryAgent__NoError)
	qt.EnumMap["bluetooth.QBluetoothServiceDiscoveryAgent__InputOutputError"] = int64(QBluetoothServiceDiscoveryAgent__InputOutputError)
	qt.EnumMap["bluetooth.QBluetoothServiceDiscoveryAgent__PoweredOffError"] = int64(QBluetoothServiceDiscoveryAgent__PoweredOffError)
	qt.EnumMap["bluetooth.QBluetoothServiceDiscoveryAgent__InvalidBluetoothAdapterError"] = int64(QBluetoothServiceDiscoveryAgent__InvalidBluetoothAdapterError)
	qt.EnumMap["bluetooth.QBluetoothServiceDiscoveryAgent__UnknownError"] = int64(QBluetoothServiceDiscoveryAgent__UnknownError)
	qt.EnumMap["bluetooth.QBluetoothServiceDiscoveryAgent__MinimalDiscovery"] = int64(QBluetoothServiceDiscoveryAgent__MinimalDiscovery)
	qt.EnumMap["bluetooth.QBluetoothServiceDiscoveryAgent__FullDiscovery"] = int64(QBluetoothServiceDiscoveryAgent__FullDiscovery)
	qt.ItfMap["bluetooth.QBluetoothServiceInfo_ITF"] = QBluetoothServiceInfo{}
	qt.FuncMap["bluetooth.NewQBluetoothServiceInfo"] = NewQBluetoothServiceInfo
	qt.FuncMap["bluetooth.NewQBluetoothServiceInfo2"] = NewQBluetoothServiceInfo2
	qt.EnumMap["bluetooth.QBluetoothServiceInfo__ServiceRecordHandle"] = int64(QBluetoothServiceInfo__ServiceRecordHandle)
	qt.EnumMap["bluetooth.QBluetoothServiceInfo__ServiceClassIds"] = int64(QBluetoothServiceInfo__ServiceClassIds)
	qt.EnumMap["bluetooth.QBluetoothServiceInfo__ServiceRecordState"] = int64(QBluetoothServiceInfo__ServiceRecordState)
	qt.EnumMap["bluetooth.QBluetoothServiceInfo__ServiceId"] = int64(QBluetoothServiceInfo__ServiceId)
	qt.EnumMap["bluetooth.QBluetoothServiceInfo__ProtocolDescriptorList"] = int64(QBluetoothServiceInfo__ProtocolDescriptorList)
	qt.EnumMap["bluetooth.QBluetoothServiceInfo__BrowseGroupList"] = int64(QBluetoothServiceInfo__BrowseGroupList)
	qt.EnumMap["bluetooth.QBluetoothServiceInfo__LanguageBaseAttributeIdList"] = int64(QBluetoothServiceInfo__LanguageBaseAttributeIdList)
	qt.EnumMap["bluetooth.QBluetoothServiceInfo__ServiceInfoTimeToLive"] = int64(QBluetoothServiceInfo__ServiceInfoTimeToLive)
	qt.EnumMap["bluetooth.QBluetoothServiceInfo__ServiceAvailability"] = int64(QBluetoothServiceInfo__ServiceAvailability)
	qt.EnumMap["bluetooth.QBluetoothServiceInfo__BluetoothProfileDescriptorList"] = int64(QBluetoothServiceInfo__BluetoothProfileDescriptorList)
	qt.EnumMap["bluetooth.QBluetoothServiceInfo__DocumentationUrl"] = int64(QBluetoothServiceInfo__DocumentationUrl)
	qt.EnumMap["bluetooth.QBluetoothServiceInfo__ClientExecutableUrl"] = int64(QBluetoothServiceInfo__ClientExecutableUrl)
	qt.EnumMap["bluetooth.QBluetoothServiceInfo__IconUrl"] = int64(QBluetoothServiceInfo__IconUrl)
	qt.EnumMap["bluetooth.QBluetoothServiceInfo__AdditionalProtocolDescriptorList"] = int64(QBluetoothServiceInfo__AdditionalProtocolDescriptorList)
	qt.EnumMap["bluetooth.QBluetoothServiceInfo__PrimaryLanguageBase"] = int64(QBluetoothServiceInfo__PrimaryLanguageBase)
	qt.EnumMap["bluetooth.QBluetoothServiceInfo__ServiceName"] = int64(QBluetoothServiceInfo__ServiceName)
	qt.EnumMap["bluetooth.QBluetoothServiceInfo__ServiceDescription"] = int64(QBluetoothServiceInfo__ServiceDescription)
	qt.EnumMap["bluetooth.QBluetoothServiceInfo__ServiceProvider"] = int64(QBluetoothServiceInfo__ServiceProvider)
	qt.EnumMap["bluetooth.QBluetoothServiceInfo__UnknownProtocol"] = int64(QBluetoothServiceInfo__UnknownProtocol)
	qt.EnumMap["bluetooth.QBluetoothServiceInfo__L2capProtocol"] = int64(QBluetoothServiceInfo__L2capProtocol)
	qt.EnumMap["bluetooth.QBluetoothServiceInfo__RfcommProtocol"] = int64(QBluetoothServiceInfo__RfcommProtocol)
	qt.ItfMap["bluetooth.QBluetoothSocket_ITF"] = QBluetoothSocket{}
	qt.FuncMap["bluetooth.NewQBluetoothSocket"] = NewQBluetoothSocket
	qt.FuncMap["bluetooth.NewQBluetoothSocket2"] = NewQBluetoothSocket2
	qt.EnumMap["bluetooth.QBluetoothSocket__UnconnectedState"] = int64(QBluetoothSocket__UnconnectedState)
	qt.EnumMap["bluetooth.QBluetoothSocket__ServiceLookupState"] = int64(QBluetoothSocket__ServiceLookupState)
	qt.EnumMap["bluetooth.QBluetoothSocket__ConnectingState"] = int64(QBluetoothSocket__ConnectingState)
	qt.EnumMap["bluetooth.QBluetoothSocket__ConnectedState"] = int64(QBluetoothSocket__ConnectedState)
	qt.EnumMap["bluetooth.QBluetoothSocket__BoundState"] = int64(QBluetoothSocket__BoundState)
	qt.EnumMap["bluetooth.QBluetoothSocket__ClosingState"] = int64(QBluetoothSocket__ClosingState)
	qt.EnumMap["bluetooth.QBluetoothSocket__ListeningState"] = int64(QBluetoothSocket__ListeningState)
	qt.EnumMap["bluetooth.QBluetoothSocket__NoSocketError"] = int64(QBluetoothSocket__NoSocketError)
	qt.EnumMap["bluetooth.QBluetoothSocket__UnknownSocketError"] = int64(QBluetoothSocket__UnknownSocketError)
	qt.EnumMap["bluetooth.QBluetoothSocket__RemoteHostClosedError"] = int64(QBluetoothSocket__RemoteHostClosedError)
	qt.EnumMap["bluetooth.QBluetoothSocket__HostNotFoundError"] = int64(QBluetoothSocket__HostNotFoundError)
	qt.EnumMap["bluetooth.QBluetoothSocket__ServiceNotFoundError"] = int64(QBluetoothSocket__ServiceNotFoundError)
	qt.EnumMap["bluetooth.QBluetoothSocket__NetworkError"] = int64(QBluetoothSocket__NetworkError)
	qt.EnumMap["bluetooth.QBluetoothSocket__UnsupportedProtocolError"] = int64(QBluetoothSocket__UnsupportedProtocolError)
	qt.EnumMap["bluetooth.QBluetoothSocket__OperationError"] = int64(QBluetoothSocket__OperationError)
	qt.ItfMap["bluetooth.QBluetoothTransferManager_ITF"] = QBluetoothTransferManager{}
	qt.FuncMap["bluetooth.NewQBluetoothTransferManager"] = NewQBluetoothTransferManager
	qt.ItfMap["bluetooth.QBluetoothTransferReply_ITF"] = QBluetoothTransferReply{}
	qt.FuncMap["bluetooth.NewQBluetoothTransferReply"] = NewQBluetoothTransferReply
	qt.EnumMap["bluetooth.QBluetoothTransferReply__NoError"] = int64(QBluetoothTransferReply__NoError)
	qt.EnumMap["bluetooth.QBluetoothTransferReply__UnknownError"] = int64(QBluetoothTransferReply__UnknownError)
	qt.EnumMap["bluetooth.QBluetoothTransferReply__FileNotFoundError"] = int64(QBluetoothTransferReply__FileNotFoundError)
	qt.EnumMap["bluetooth.QBluetoothTransferReply__HostNotFoundError"] = int64(QBluetoothTransferReply__HostNotFoundError)
	qt.EnumMap["bluetooth.QBluetoothTransferReply__UserCanceledTransferError"] = int64(QBluetoothTransferReply__UserCanceledTransferError)
	qt.EnumMap["bluetooth.QBluetoothTransferReply__IODeviceNotReadableError"] = int64(QBluetoothTransferReply__IODeviceNotReadableError)
	qt.EnumMap["bluetooth.QBluetoothTransferReply__ResourceBusyError"] = int64(QBluetoothTransferReply__ResourceBusyError)
	qt.EnumMap["bluetooth.QBluetoothTransferReply__SessionError"] = int64(QBluetoothTransferReply__SessionError)
	qt.ItfMap["bluetooth.QBluetoothTransferRequest_ITF"] = QBluetoothTransferRequest{}
	qt.FuncMap["bluetooth.NewQBluetoothTransferRequest"] = NewQBluetoothTransferRequest
	qt.FuncMap["bluetooth.NewQBluetoothTransferRequest2"] = NewQBluetoothTransferRequest2
	qt.EnumMap["bluetooth.QBluetoothTransferRequest__DescriptionAttribute"] = int64(QBluetoothTransferRequest__DescriptionAttribute)
	qt.EnumMap["bluetooth.QBluetoothTransferRequest__TimeAttribute"] = int64(QBluetoothTransferRequest__TimeAttribute)
	qt.EnumMap["bluetooth.QBluetoothTransferRequest__TypeAttribute"] = int64(QBluetoothTransferRequest__TypeAttribute)
	qt.EnumMap["bluetooth.QBluetoothTransferRequest__LengthAttribute"] = int64(QBluetoothTransferRequest__LengthAttribute)
	qt.EnumMap["bluetooth.QBluetoothTransferRequest__NameAttribute"] = int64(QBluetoothTransferRequest__NameAttribute)
	qt.ItfMap["bluetooth.QBluetoothUuid_ITF"] = QBluetoothUuid{}
	qt.FuncMap["bluetooth.NewQBluetoothUuid"] = NewQBluetoothUuid
	qt.FuncMap["bluetooth.NewQBluetoothUuid2"] = NewQBluetoothUuid2
	qt.FuncMap["bluetooth.NewQBluetoothUuid3"] = NewQBluetoothUuid3
	qt.FuncMap["bluetooth.NewQBluetoothUuid4"] = NewQBluetoothUuid4
	qt.FuncMap["bluetooth.NewQBluetoothUuid5"] = NewQBluetoothUuid5
	qt.FuncMap["bluetooth.NewQBluetoothUuid6"] = NewQBluetoothUuid6
	qt.FuncMap["bluetooth.NewQBluetoothUuid7"] = NewQBluetoothUuid7
	qt.FuncMap["bluetooth.NewQBluetoothUuid9"] = NewQBluetoothUuid9
	qt.FuncMap["bluetooth.NewQBluetoothUuid10"] = NewQBluetoothUuid10
	qt.FuncMap["bluetooth.NewQBluetoothUuid11"] = NewQBluetoothUuid11
	qt.FuncMap["bluetooth.QBluetoothUuid_CharacteristicToString"] = QBluetoothUuid_CharacteristicToString
	qt.FuncMap["bluetooth.QBluetoothUuid_DescriptorToString"] = QBluetoothUuid_DescriptorToString
	qt.FuncMap["bluetooth.QBluetoothUuid_ProtocolToString"] = QBluetoothUuid_ProtocolToString
	qt.FuncMap["bluetooth.QBluetoothUuid_ServiceClassToString"] = QBluetoothUuid_ServiceClassToString
	qt.EnumMap["bluetooth.QBluetoothUuid__Sdp"] = int64(QBluetoothUuid__Sdp)
	qt.EnumMap["bluetooth.QBluetoothUuid__Udp"] = int64(QBluetoothUuid__Udp)
	qt.EnumMap["bluetooth.QBluetoothUuid__Rfcomm"] = int64(QBluetoothUuid__Rfcomm)
	qt.EnumMap["bluetooth.QBluetoothUuid__Tcp"] = int64(QBluetoothUuid__Tcp)
	qt.EnumMap["bluetooth.QBluetoothUuid__TcsBin"] = int64(QBluetoothUuid__TcsBin)
	qt.EnumMap["bluetooth.QBluetoothUuid__TcsAt"] = int64(QBluetoothUuid__TcsAt)
	qt.EnumMap["bluetooth.QBluetoothUuid__Att"] = int64(QBluetoothUuid__Att)
	qt.EnumMap["bluetooth.QBluetoothUuid__Obex"] = int64(QBluetoothUuid__Obex)
	qt.EnumMap["bluetooth.QBluetoothUuid__Ip"] = int64(QBluetoothUuid__Ip)
	qt.EnumMap["bluetooth.QBluetoothUuid__Ftp"] = int64(QBluetoothUuid__Ftp)
	qt.EnumMap["bluetooth.QBluetoothUuid__Http"] = int64(QBluetoothUuid__Http)
	qt.EnumMap["bluetooth.QBluetoothUuid__Wsp"] = int64(QBluetoothUuid__Wsp)
	qt.EnumMap["bluetooth.QBluetoothUuid__Bnep"] = int64(QBluetoothUuid__Bnep)
	qt.EnumMap["bluetooth.QBluetoothUuid__Upnp"] = int64(QBluetoothUuid__Upnp)
	qt.EnumMap["bluetooth.QBluetoothUuid__Hidp"] = int64(QBluetoothUuid__Hidp)
	qt.EnumMap["bluetooth.QBluetoothUuid__HardcopyControlChannel"] = int64(QBluetoothUuid__HardcopyControlChannel)
	qt.EnumMap["bluetooth.QBluetoothUuid__HardcopyDataChannel"] = int64(QBluetoothUuid__HardcopyDataChannel)
	qt.EnumMap["bluetooth.QBluetoothUuid__HardcopyNotification"] = int64(QBluetoothUuid__HardcopyNotification)
	qt.EnumMap["bluetooth.QBluetoothUuid__Avctp"] = int64(QBluetoothUuid__Avctp)
	qt.EnumMap["bluetooth.QBluetoothUuid__Avdtp"] = int64(QBluetoothUuid__Avdtp)
	qt.EnumMap["bluetooth.QBluetoothUuid__Cmtp"] = int64(QBluetoothUuid__Cmtp)
	qt.EnumMap["bluetooth.QBluetoothUuid__UdiCPlain"] = int64(QBluetoothUuid__UdiCPlain)
	qt.EnumMap["bluetooth.QBluetoothUuid__McapControlChannel"] = int64(QBluetoothUuid__McapControlChannel)
	qt.EnumMap["bluetooth.QBluetoothUuid__McapDataChannel"] = int64(QBluetoothUuid__McapDataChannel)
	qt.EnumMap["bluetooth.QBluetoothUuid__L2cap"] = int64(QBluetoothUuid__L2cap)
	qt.EnumMap["bluetooth.QBluetoothUuid__ServiceDiscoveryServer"] = int64(QBluetoothUuid__ServiceDiscoveryServer)
	qt.EnumMap["bluetooth.QBluetoothUuid__BrowseGroupDescriptor"] = int64(QBluetoothUuid__BrowseGroupDescriptor)
	qt.EnumMap["bluetooth.QBluetoothUuid__PublicBrowseGroup"] = int64(QBluetoothUuid__PublicBrowseGroup)
	qt.EnumMap["bluetooth.QBluetoothUuid__SerialPort"] = int64(QBluetoothUuid__SerialPort)
	qt.EnumMap["bluetooth.QBluetoothUuid__LANAccessUsingPPP"] = int64(QBluetoothUuid__LANAccessUsingPPP)
	qt.EnumMap["bluetooth.QBluetoothUuid__DialupNetworking"] = int64(QBluetoothUuid__DialupNetworking)
	qt.EnumMap["bluetooth.QBluetoothUuid__IrMCSync"] = int64(QBluetoothUuid__IrMCSync)
	qt.EnumMap["bluetooth.QBluetoothUuid__ObexObjectPush"] = int64(QBluetoothUuid__ObexObjectPush)
	qt.EnumMap["bluetooth.QBluetoothUuid__OBEXFileTransfer"] = int64(QBluetoothUuid__OBEXFileTransfer)
	qt.EnumMap["bluetooth.QBluetoothUuid__IrMCSyncCommand"] = int64(QBluetoothUuid__IrMCSyncCommand)
	qt.EnumMap["bluetooth.QBluetoothUuid__Headset"] = int64(QBluetoothUuid__Headset)
	qt.EnumMap["bluetooth.QBluetoothUuid__AudioSource"] = int64(QBluetoothUuid__AudioSource)
	qt.EnumMap["bluetooth.QBluetoothUuid__AudioSink"] = int64(QBluetoothUuid__AudioSink)
	qt.EnumMap["bluetooth.QBluetoothUuid__AV_RemoteControlTarget"] = int64(QBluetoothUuid__AV_RemoteControlTarget)
	qt.EnumMap["bluetooth.QBluetoothUuid__AdvancedAudioDistribution"] = int64(QBluetoothUuid__AdvancedAudioDistribution)
	qt.EnumMap["bluetooth.QBluetoothUuid__AV_RemoteControl"] = int64(QBluetoothUuid__AV_RemoteControl)
	qt.EnumMap["bluetooth.QBluetoothUuid__AV_RemoteControlController"] = int64(QBluetoothUuid__AV_RemoteControlController)
	qt.EnumMap["bluetooth.QBluetoothUuid__HeadsetAG"] = int64(QBluetoothUuid__HeadsetAG)
	qt.EnumMap["bluetooth.QBluetoothUuid__PANU"] = int64(QBluetoothUuid__PANU)
	qt.EnumMap["bluetooth.QBluetoothUuid__NAP"] = int64(QBluetoothUuid__NAP)
	qt.EnumMap["bluetooth.QBluetoothUuid__GN"] = int64(QBluetoothUuid__GN)
	qt.EnumMap["bluetooth.QBluetoothUuid__DirectPrinting"] = int64(QBluetoothUuid__DirectPrinting)
	qt.EnumMap["bluetooth.QBluetoothUuid__ReferencePrinting"] = int64(QBluetoothUuid__ReferencePrinting)
	qt.EnumMap["bluetooth.QBluetoothUuid__BasicImage"] = int64(QBluetoothUuid__BasicImage)
	qt.EnumMap["bluetooth.QBluetoothUuid__ImagingResponder"] = int64(QBluetoothUuid__ImagingResponder)
	qt.EnumMap["bluetooth.QBluetoothUuid__ImagingAutomaticArchive"] = int64(QBluetoothUuid__ImagingAutomaticArchive)
	qt.EnumMap["bluetooth.QBluetoothUuid__ImagingReferenceObjects"] = int64(QBluetoothUuid__ImagingReferenceObjects)
	qt.EnumMap["bluetooth.QBluetoothUuid__Handsfree"] = int64(QBluetoothUuid__Handsfree)
	qt.EnumMap["bluetooth.QBluetoothUuid__HandsfreeAudioGateway"] = int64(QBluetoothUuid__HandsfreeAudioGateway)
	qt.EnumMap["bluetooth.QBluetoothUuid__DirectPrintingReferenceObjectsService"] = int64(QBluetoothUuid__DirectPrintingReferenceObjectsService)
	qt.EnumMap["bluetooth.QBluetoothUuid__ReflectedUI"] = int64(QBluetoothUuid__ReflectedUI)
	qt.EnumMap["bluetooth.QBluetoothUuid__BasicPrinting"] = int64(QBluetoothUuid__BasicPrinting)
	qt.EnumMap["bluetooth.QBluetoothUuid__PrintingStatus"] = int64(QBluetoothUuid__PrintingStatus)
	qt.EnumMap["bluetooth.QBluetoothUuid__HumanInterfaceDeviceService"] = int64(QBluetoothUuid__HumanInterfaceDeviceService)
	qt.EnumMap["bluetooth.QBluetoothUuid__HardcopyCableReplacement"] = int64(QBluetoothUuid__HardcopyCableReplacement)
	qt.EnumMap["bluetooth.QBluetoothUuid__HCRPrint"] = int64(QBluetoothUuid__HCRPrint)
	qt.EnumMap["bluetooth.QBluetoothUuid__HCRScan"] = int64(QBluetoothUuid__HCRScan)
	qt.EnumMap["bluetooth.QBluetoothUuid__SIMAccess"] = int64(QBluetoothUuid__SIMAccess)
	qt.EnumMap["bluetooth.QBluetoothUuid__PhonebookAccessPCE"] = int64(QBluetoothUuid__PhonebookAccessPCE)
	qt.EnumMap["bluetooth.QBluetoothUuid__PhonebookAccessPSE"] = int64(QBluetoothUuid__PhonebookAccessPSE)
	qt.EnumMap["bluetooth.QBluetoothUuid__PhonebookAccess"] = int64(QBluetoothUuid__PhonebookAccess)
	qt.EnumMap["bluetooth.QBluetoothUuid__HeadsetHS"] = int64(QBluetoothUuid__HeadsetHS)
	qt.EnumMap["bluetooth.QBluetoothUuid__MessageAccessServer"] = int64(QBluetoothUuid__MessageAccessServer)
	qt.EnumMap["bluetooth.QBluetoothUuid__MessageNotificationServer"] = int64(QBluetoothUuid__MessageNotificationServer)
	qt.EnumMap["bluetooth.QBluetoothUuid__MessageAccessProfile"] = int64(QBluetoothUuid__MessageAccessProfile)
	qt.EnumMap["bluetooth.QBluetoothUuid__GNSS"] = int64(QBluetoothUuid__GNSS)
	qt.EnumMap["bluetooth.QBluetoothUuid__GNSSServer"] = int64(QBluetoothUuid__GNSSServer)
	qt.EnumMap["bluetooth.QBluetoothUuid__Display3D"] = int64(QBluetoothUuid__Display3D)
	qt.EnumMap["bluetooth.QBluetoothUuid__Glasses3D"] = int64(QBluetoothUuid__Glasses3D)
	qt.EnumMap["bluetooth.QBluetoothUuid__Synchronization3D"] = int64(QBluetoothUuid__Synchronization3D)
	qt.EnumMap["bluetooth.QBluetoothUuid__MPSProfile"] = int64(QBluetoothUuid__MPSProfile)
	qt.EnumMap["bluetooth.QBluetoothUuid__MPSService"] = int64(QBluetoothUuid__MPSService)
	qt.EnumMap["bluetooth.QBluetoothUuid__PnPInformation"] = int64(QBluetoothUuid__PnPInformation)
	qt.EnumMap["bluetooth.QBluetoothUuid__GenericNetworking"] = int64(QBluetoothUuid__GenericNetworking)
	qt.EnumMap["bluetooth.QBluetoothUuid__GenericFileTransfer"] = int64(QBluetoothUuid__GenericFileTransfer)
	qt.EnumMap["bluetooth.QBluetoothUuid__GenericAudio"] = int64(QBluetoothUuid__GenericAudio)
	qt.EnumMap["bluetooth.QBluetoothUuid__GenericTelephony"] = int64(QBluetoothUuid__GenericTelephony)
	qt.EnumMap["bluetooth.QBluetoothUuid__VideoSource"] = int64(QBluetoothUuid__VideoSource)
	qt.EnumMap["bluetooth.QBluetoothUuid__VideoSink"] = int64(QBluetoothUuid__VideoSink)
	qt.EnumMap["bluetooth.QBluetoothUuid__VideoDistribution"] = int64(QBluetoothUuid__VideoDistribution)
	qt.EnumMap["bluetooth.QBluetoothUuid__HDP"] = int64(QBluetoothUuid__HDP)
	qt.EnumMap["bluetooth.QBluetoothUuid__HDPSource"] = int64(QBluetoothUuid__HDPSource)
	qt.EnumMap["bluetooth.QBluetoothUuid__HDPSink"] = int64(QBluetoothUuid__HDPSink)
	qt.EnumMap["bluetooth.QBluetoothUuid__GenericAccess"] = int64(QBluetoothUuid__GenericAccess)
	qt.EnumMap["bluetooth.QBluetoothUuid__GenericAttribute"] = int64(QBluetoothUuid__GenericAttribute)
	qt.EnumMap["bluetooth.QBluetoothUuid__ImmediateAlert"] = int64(QBluetoothUuid__ImmediateAlert)
	qt.EnumMap["bluetooth.QBluetoothUuid__LinkLoss"] = int64(QBluetoothUuid__LinkLoss)
	qt.EnumMap["bluetooth.QBluetoothUuid__TxPower"] = int64(QBluetoothUuid__TxPower)
	qt.EnumMap["bluetooth.QBluetoothUuid__CurrentTimeService"] = int64(QBluetoothUuid__CurrentTimeService)
	qt.EnumMap["bluetooth.QBluetoothUuid__ReferenceTimeUpdateService"] = int64(QBluetoothUuid__ReferenceTimeUpdateService)
	qt.EnumMap["bluetooth.QBluetoothUuid__NextDSTChangeService"] = int64(QBluetoothUuid__NextDSTChangeService)
	qt.EnumMap["bluetooth.QBluetoothUuid__Glucose"] = int64(QBluetoothUuid__Glucose)
	qt.EnumMap["bluetooth.QBluetoothUuid__HealthThermometer"] = int64(QBluetoothUuid__HealthThermometer)
	qt.EnumMap["bluetooth.QBluetoothUuid__DeviceInformation"] = int64(QBluetoothUuid__DeviceInformation)
	qt.EnumMap["bluetooth.QBluetoothUuid__HeartRate"] = int64(QBluetoothUuid__HeartRate)
	qt.EnumMap["bluetooth.QBluetoothUuid__PhoneAlertStatusService"] = int64(QBluetoothUuid__PhoneAlertStatusService)
	qt.EnumMap["bluetooth.QBluetoothUuid__BatteryService"] = int64(QBluetoothUuid__BatteryService)
	qt.EnumMap["bluetooth.QBluetoothUuid__BloodPressure"] = int64(QBluetoothUuid__BloodPressure)
	qt.EnumMap["bluetooth.QBluetoothUuid__AlertNotificationService"] = int64(QBluetoothUuid__AlertNotificationService)
	qt.EnumMap["bluetooth.QBluetoothUuid__HumanInterfaceDevice"] = int64(QBluetoothUuid__HumanInterfaceDevice)
	qt.EnumMap["bluetooth.QBluetoothUuid__ScanParameters"] = int64(QBluetoothUuid__ScanParameters)
	qt.EnumMap["bluetooth.QBluetoothUuid__RunningSpeedAndCadence"] = int64(QBluetoothUuid__RunningSpeedAndCadence)
	qt.EnumMap["bluetooth.QBluetoothUuid__CyclingSpeedAndCadence"] = int64(QBluetoothUuid__CyclingSpeedAndCadence)
	qt.EnumMap["bluetooth.QBluetoothUuid__CyclingPower"] = int64(QBluetoothUuid__CyclingPower)
	qt.EnumMap["bluetooth.QBluetoothUuid__LocationAndNavigation"] = int64(QBluetoothUuid__LocationAndNavigation)
	qt.EnumMap["bluetooth.QBluetoothUuid__EnvironmentalSensing"] = int64(QBluetoothUuid__EnvironmentalSensing)
	qt.EnumMap["bluetooth.QBluetoothUuid__BodyComposition"] = int64(QBluetoothUuid__BodyComposition)
	qt.EnumMap["bluetooth.QBluetoothUuid__UserData"] = int64(QBluetoothUuid__UserData)
	qt.EnumMap["bluetooth.QBluetoothUuid__WeightScale"] = int64(QBluetoothUuid__WeightScale)
	qt.EnumMap["bluetooth.QBluetoothUuid__BondManagement"] = int64(QBluetoothUuid__BondManagement)
	qt.EnumMap["bluetooth.QBluetoothUuid__ContinuousGlucoseMonitoring"] = int64(QBluetoothUuid__ContinuousGlucoseMonitoring)
	qt.EnumMap["bluetooth.QBluetoothUuid__DeviceName"] = int64(QBluetoothUuid__DeviceName)
	qt.EnumMap["bluetooth.QBluetoothUuid__Appearance"] = int64(QBluetoothUuid__Appearance)
	qt.EnumMap["bluetooth.QBluetoothUuid__PeripheralPrivacyFlag"] = int64(QBluetoothUuid__PeripheralPrivacyFlag)
	qt.EnumMap["bluetooth.QBluetoothUuid__ReconnectionAddress"] = int64(QBluetoothUuid__ReconnectionAddress)
	qt.EnumMap["bluetooth.QBluetoothUuid__PeripheralPreferredConnectionParameters"] = int64(QBluetoothUuid__PeripheralPreferredConnectionParameters)
	qt.EnumMap["bluetooth.QBluetoothUuid__ServiceChanged"] = int64(QBluetoothUuid__ServiceChanged)
	qt.EnumMap["bluetooth.QBluetoothUuid__AlertLevel"] = int64(QBluetoothUuid__AlertLevel)
	qt.EnumMap["bluetooth.QBluetoothUuid__TxPowerLevel"] = int64(QBluetoothUuid__TxPowerLevel)
	qt.EnumMap["bluetooth.QBluetoothUuid__DateTime"] = int64(QBluetoothUuid__DateTime)
	qt.EnumMap["bluetooth.QBluetoothUuid__DayOfWeek"] = int64(QBluetoothUuid__DayOfWeek)
	qt.EnumMap["bluetooth.QBluetoothUuid__DayDateTime"] = int64(QBluetoothUuid__DayDateTime)
	qt.EnumMap["bluetooth.QBluetoothUuid__ExactTime256"] = int64(QBluetoothUuid__ExactTime256)
	qt.EnumMap["bluetooth.QBluetoothUuid__DSTOffset"] = int64(QBluetoothUuid__DSTOffset)
	qt.EnumMap["bluetooth.QBluetoothUuid__TimeZone"] = int64(QBluetoothUuid__TimeZone)
	qt.EnumMap["bluetooth.QBluetoothUuid__LocalTimeInformation"] = int64(QBluetoothUuid__LocalTimeInformation)
	qt.EnumMap["bluetooth.QBluetoothUuid__TimeWithDST"] = int64(QBluetoothUuid__TimeWithDST)
	qt.EnumMap["bluetooth.QBluetoothUuid__TimeAccuracy"] = int64(QBluetoothUuid__TimeAccuracy)
	qt.EnumMap["bluetooth.QBluetoothUuid__TimeSource"] = int64(QBluetoothUuid__TimeSource)
	qt.EnumMap["bluetooth.QBluetoothUuid__ReferenceTimeInformation"] = int64(QBluetoothUuid__ReferenceTimeInformation)
	qt.EnumMap["bluetooth.QBluetoothUuid__TimeUpdateControlPoint"] = int64(QBluetoothUuid__TimeUpdateControlPoint)
	qt.EnumMap["bluetooth.QBluetoothUuid__TimeUpdateState"] = int64(QBluetoothUuid__TimeUpdateState)
	qt.EnumMap["bluetooth.QBluetoothUuid__GlucoseMeasurement"] = int64(QBluetoothUuid__GlucoseMeasurement)
	qt.EnumMap["bluetooth.QBluetoothUuid__BatteryLevel"] = int64(QBluetoothUuid__BatteryLevel)
	qt.EnumMap["bluetooth.QBluetoothUuid__TemperatureMeasurement"] = int64(QBluetoothUuid__TemperatureMeasurement)
	qt.EnumMap["bluetooth.QBluetoothUuid__TemperatureType"] = int64(QBluetoothUuid__TemperatureType)
	qt.EnumMap["bluetooth.QBluetoothUuid__IntermediateTemperature"] = int64(QBluetoothUuid__IntermediateTemperature)
	qt.EnumMap["bluetooth.QBluetoothUuid__MeasurementInterval"] = int64(QBluetoothUuid__MeasurementInterval)
	qt.EnumMap["bluetooth.QBluetoothUuid__BootKeyboardInputReport"] = int64(QBluetoothUuid__BootKeyboardInputReport)
	qt.EnumMap["bluetooth.QBluetoothUuid__SystemID"] = int64(QBluetoothUuid__SystemID)
	qt.EnumMap["bluetooth.QBluetoothUuid__ModelNumberString"] = int64(QBluetoothUuid__ModelNumberString)
	qt.EnumMap["bluetooth.QBluetoothUuid__SerialNumberString"] = int64(QBluetoothUuid__SerialNumberString)
	qt.EnumMap["bluetooth.QBluetoothUuid__FirmwareRevisionString"] = int64(QBluetoothUuid__FirmwareRevisionString)
	qt.EnumMap["bluetooth.QBluetoothUuid__HardwareRevisionString"] = int64(QBluetoothUuid__HardwareRevisionString)
	qt.EnumMap["bluetooth.QBluetoothUuid__SoftwareRevisionString"] = int64(QBluetoothUuid__SoftwareRevisionString)
	qt.EnumMap["bluetooth.QBluetoothUuid__ManufacturerNameString"] = int64(QBluetoothUuid__ManufacturerNameString)
	qt.EnumMap["bluetooth.QBluetoothUuid__IEEE1107320601RegulatoryCertificationDataList"] = int64(QBluetoothUuid__IEEE1107320601RegulatoryCertificationDataList)
	qt.EnumMap["bluetooth.QBluetoothUuid__CurrentTime"] = int64(QBluetoothUuid__CurrentTime)
	qt.EnumMap["bluetooth.QBluetoothUuid__MagneticDeclination"] = int64(QBluetoothUuid__MagneticDeclination)
	qt.EnumMap["bluetooth.QBluetoothUuid__ScanRefresh"] = int64(QBluetoothUuid__ScanRefresh)
	qt.EnumMap["bluetooth.QBluetoothUuid__BootKeyboardOutputReport"] = int64(QBluetoothUuid__BootKeyboardOutputReport)
	qt.EnumMap["bluetooth.QBluetoothUuid__BootMouseInputReport"] = int64(QBluetoothUuid__BootMouseInputReport)
	qt.EnumMap["bluetooth.QBluetoothUuid__GlucoseMeasurementContext"] = int64(QBluetoothUuid__GlucoseMeasurementContext)
	qt.EnumMap["bluetooth.QBluetoothUuid__BloodPressureMeasurement"] = int64(QBluetoothUuid__BloodPressureMeasurement)
	qt.EnumMap["bluetooth.QBluetoothUuid__IntermediateCuffPressure"] = int64(QBluetoothUuid__IntermediateCuffPressure)
	qt.EnumMap["bluetooth.QBluetoothUuid__HeartRateMeasurement"] = int64(QBluetoothUuid__HeartRateMeasurement)
	qt.EnumMap["bluetooth.QBluetoothUuid__BodySensorLocation"] = int64(QBluetoothUuid__BodySensorLocation)
	qt.EnumMap["bluetooth.QBluetoothUuid__HeartRateControlPoint"] = int64(QBluetoothUuid__HeartRateControlPoint)
	qt.EnumMap["bluetooth.QBluetoothUuid__AlertStatus"] = int64(QBluetoothUuid__AlertStatus)
	qt.EnumMap["bluetooth.QBluetoothUuid__RingerControlPoint"] = int64(QBluetoothUuid__RingerControlPoint)
	qt.EnumMap["bluetooth.QBluetoothUuid__RingerSetting"] = int64(QBluetoothUuid__RingerSetting)
	qt.EnumMap["bluetooth.QBluetoothUuid__AlertCategoryIDBitMask"] = int64(QBluetoothUuid__AlertCategoryIDBitMask)
	qt.EnumMap["bluetooth.QBluetoothUuid__AlertCategoryID"] = int64(QBluetoothUuid__AlertCategoryID)
	qt.EnumMap["bluetooth.QBluetoothUuid__AlertNotificationControlPoint"] = int64(QBluetoothUuid__AlertNotificationControlPoint)
	qt.EnumMap["bluetooth.QBluetoothUuid__UnreadAlertStatus"] = int64(QBluetoothUuid__UnreadAlertStatus)
	qt.EnumMap["bluetooth.QBluetoothUuid__NewAlert"] = int64(QBluetoothUuid__NewAlert)
	qt.EnumMap["bluetooth.QBluetoothUuid__SupportedNewAlertCategory"] = int64(QBluetoothUuid__SupportedNewAlertCategory)
	qt.EnumMap["bluetooth.QBluetoothUuid__SupportedUnreadAlertCategory"] = int64(QBluetoothUuid__SupportedUnreadAlertCategory)
	qt.EnumMap["bluetooth.QBluetoothUuid__BloodPressureFeature"] = int64(QBluetoothUuid__BloodPressureFeature)
	qt.EnumMap["bluetooth.QBluetoothUuid__HIDInformation"] = int64(QBluetoothUuid__HIDInformation)
	qt.EnumMap["bluetooth.QBluetoothUuid__ReportMap"] = int64(QBluetoothUuid__ReportMap)
	qt.EnumMap["bluetooth.QBluetoothUuid__HIDControlPoint"] = int64(QBluetoothUuid__HIDControlPoint)
	qt.EnumMap["bluetooth.QBluetoothUuid__Report"] = int64(QBluetoothUuid__Report)
	qt.EnumMap["bluetooth.QBluetoothUuid__ProtocolMode"] = int64(QBluetoothUuid__ProtocolMode)
	qt.EnumMap["bluetooth.QBluetoothUuid__ScanIntervalWindow"] = int64(QBluetoothUuid__ScanIntervalWindow)
	qt.EnumMap["bluetooth.QBluetoothUuid__PnPID"] = int64(QBluetoothUuid__PnPID)
	qt.EnumMap["bluetooth.QBluetoothUuid__GlucoseFeature"] = int64(QBluetoothUuid__GlucoseFeature)
	qt.EnumMap["bluetooth.QBluetoothUuid__RecordAccessControlPoint"] = int64(QBluetoothUuid__RecordAccessControlPoint)
	qt.EnumMap["bluetooth.QBluetoothUuid__RSCMeasurement"] = int64(QBluetoothUuid__RSCMeasurement)
	qt.EnumMap["bluetooth.QBluetoothUuid__RSCFeature"] = int64(QBluetoothUuid__RSCFeature)
	qt.EnumMap["bluetooth.QBluetoothUuid__SCControlPoint"] = int64(QBluetoothUuid__SCControlPoint)
	qt.EnumMap["bluetooth.QBluetoothUuid__CSCMeasurement"] = int64(QBluetoothUuid__CSCMeasurement)
	qt.EnumMap["bluetooth.QBluetoothUuid__CSCFeature"] = int64(QBluetoothUuid__CSCFeature)
	qt.EnumMap["bluetooth.QBluetoothUuid__SensorLocation"] = int64(QBluetoothUuid__SensorLocation)
	qt.EnumMap["bluetooth.QBluetoothUuid__CyclingPowerMeasurement"] = int64(QBluetoothUuid__CyclingPowerMeasurement)
	qt.EnumMap["bluetooth.QBluetoothUuid__CyclingPowerVector"] = int64(QBluetoothUuid__CyclingPowerVector)
	qt.EnumMap["bluetooth.QBluetoothUuid__CyclingPowerFeature"] = int64(QBluetoothUuid__CyclingPowerFeature)
	qt.EnumMap["bluetooth.QBluetoothUuid__CyclingPowerControlPoint"] = int64(QBluetoothUuid__CyclingPowerControlPoint)
	qt.EnumMap["bluetooth.QBluetoothUuid__LocationAndSpeed"] = int64(QBluetoothUuid__LocationAndSpeed)
	qt.EnumMap["bluetooth.QBluetoothUuid__Navigation"] = int64(QBluetoothUuid__Navigation)
	qt.EnumMap["bluetooth.QBluetoothUuid__PositionQuality"] = int64(QBluetoothUuid__PositionQuality)
	qt.EnumMap["bluetooth.QBluetoothUuid__LNFeature"] = int64(QBluetoothUuid__LNFeature)
	qt.EnumMap["bluetooth.QBluetoothUuid__LNControlPoint"] = int64(QBluetoothUuid__LNControlPoint)
	qt.EnumMap["bluetooth.QBluetoothUuid__Elevation"] = int64(QBluetoothUuid__Elevation)
	qt.EnumMap["bluetooth.QBluetoothUuid__Pressure"] = int64(QBluetoothUuid__Pressure)
	qt.EnumMap["bluetooth.QBluetoothUuid__Temperature"] = int64(QBluetoothUuid__Temperature)
	qt.EnumMap["bluetooth.QBluetoothUuid__Humidity"] = int64(QBluetoothUuid__Humidity)
	qt.EnumMap["bluetooth.QBluetoothUuid__TrueWindSpeed"] = int64(QBluetoothUuid__TrueWindSpeed)
	qt.EnumMap["bluetooth.QBluetoothUuid__TrueWindDirection"] = int64(QBluetoothUuid__TrueWindDirection)
	qt.EnumMap["bluetooth.QBluetoothUuid__ApparentWindSpeed"] = int64(QBluetoothUuid__ApparentWindSpeed)
	qt.EnumMap["bluetooth.QBluetoothUuid__ApparentWindDirection"] = int64(QBluetoothUuid__ApparentWindDirection)
	qt.EnumMap["bluetooth.QBluetoothUuid__GustFactor"] = int64(QBluetoothUuid__GustFactor)
	qt.EnumMap["bluetooth.QBluetoothUuid__PollenConcentration"] = int64(QBluetoothUuid__PollenConcentration)
	qt.EnumMap["bluetooth.QBluetoothUuid__UVIndex"] = int64(QBluetoothUuid__UVIndex)
	qt.EnumMap["bluetooth.QBluetoothUuid__Irradiance"] = int64(QBluetoothUuid__Irradiance)
	qt.EnumMap["bluetooth.QBluetoothUuid__Rainfall"] = int64(QBluetoothUuid__Rainfall)
	qt.EnumMap["bluetooth.QBluetoothUuid__WindChill"] = int64(QBluetoothUuid__WindChill)
	qt.EnumMap["bluetooth.QBluetoothUuid__HeatIndex"] = int64(QBluetoothUuid__HeatIndex)
	qt.EnumMap["bluetooth.QBluetoothUuid__DewPoint"] = int64(QBluetoothUuid__DewPoint)
	qt.EnumMap["bluetooth.QBluetoothUuid__DescriptorValueChanged"] = int64(QBluetoothUuid__DescriptorValueChanged)
	qt.EnumMap["bluetooth.QBluetoothUuid__AerobicHeartRateLowerLimit"] = int64(QBluetoothUuid__AerobicHeartRateLowerLimit)
	qt.EnumMap["bluetooth.QBluetoothUuid__AerobicThreshold"] = int64(QBluetoothUuid__AerobicThreshold)
	qt.EnumMap["bluetooth.QBluetoothUuid__Age"] = int64(QBluetoothUuid__Age)
	qt.EnumMap["bluetooth.QBluetoothUuid__AnaerobicHeartRateLowerLimit"] = int64(QBluetoothUuid__AnaerobicHeartRateLowerLimit)
	qt.EnumMap["bluetooth.QBluetoothUuid__AnaerobicHeartRateUpperLimit"] = int64(QBluetoothUuid__AnaerobicHeartRateUpperLimit)
	qt.EnumMap["bluetooth.QBluetoothUuid__AnaerobicThreshold"] = int64(QBluetoothUuid__AnaerobicThreshold)
	qt.EnumMap["bluetooth.QBluetoothUuid__AerobicHeartRateUpperLimit"] = int64(QBluetoothUuid__AerobicHeartRateUpperLimit)
	qt.EnumMap["bluetooth.QBluetoothUuid__DateOfBirth"] = int64(QBluetoothUuid__DateOfBirth)
	qt.EnumMap["bluetooth.QBluetoothUuid__DateOfThresholdAssessment"] = int64(QBluetoothUuid__DateOfThresholdAssessment)
	qt.EnumMap["bluetooth.QBluetoothUuid__EmailAddress"] = int64(QBluetoothUuid__EmailAddress)
	qt.EnumMap["bluetooth.QBluetoothUuid__FatBurnHeartRateLowerLimit"] = int64(QBluetoothUuid__FatBurnHeartRateLowerLimit)
	qt.EnumMap["bluetooth.QBluetoothUuid__FatBurnHeartRateUpperLimit"] = int64(QBluetoothUuid__FatBurnHeartRateUpperLimit)
	qt.EnumMap["bluetooth.QBluetoothUuid__FirstName"] = int64(QBluetoothUuid__FirstName)
	qt.EnumMap["bluetooth.QBluetoothUuid__FiveZoneHeartRateLimits"] = int64(QBluetoothUuid__FiveZoneHeartRateLimits)
	qt.EnumMap["bluetooth.QBluetoothUuid__Gender"] = int64(QBluetoothUuid__Gender)
	qt.EnumMap["bluetooth.QBluetoothUuid__HeartRateMax"] = int64(QBluetoothUuid__HeartRateMax)
	qt.EnumMap["bluetooth.QBluetoothUuid__Height"] = int64(QBluetoothUuid__Height)
	qt.EnumMap["bluetooth.QBluetoothUuid__HipCircumference"] = int64(QBluetoothUuid__HipCircumference)
	qt.EnumMap["bluetooth.QBluetoothUuid__LastName"] = int64(QBluetoothUuid__LastName)
	qt.EnumMap["bluetooth.QBluetoothUuid__MaximumRecommendedHeartRate"] = int64(QBluetoothUuid__MaximumRecommendedHeartRate)
	qt.EnumMap["bluetooth.QBluetoothUuid__RestingHeartRate"] = int64(QBluetoothUuid__RestingHeartRate)
	qt.EnumMap["bluetooth.QBluetoothUuid__SportTypeForAerobicAnaerobicThresholds"] = int64(QBluetoothUuid__SportTypeForAerobicAnaerobicThresholds)
	qt.EnumMap["bluetooth.QBluetoothUuid__ThreeZoneHeartRateLimits"] = int64(QBluetoothUuid__ThreeZoneHeartRateLimits)
	qt.EnumMap["bluetooth.QBluetoothUuid__TwoZoneHeartRateLimits"] = int64(QBluetoothUuid__TwoZoneHeartRateLimits)
	qt.EnumMap["bluetooth.QBluetoothUuid__VO2Max"] = int64(QBluetoothUuid__VO2Max)
	qt.EnumMap["bluetooth.QBluetoothUuid__WaistCircumference"] = int64(QBluetoothUuid__WaistCircumference)
	qt.EnumMap["bluetooth.QBluetoothUuid__Weight"] = int64(QBluetoothUuid__Weight)
	qt.EnumMap["bluetooth.QBluetoothUuid__DatabaseChangeIncrement"] = int64(QBluetoothUuid__DatabaseChangeIncrement)
	qt.EnumMap["bluetooth.QBluetoothUuid__UserIndex"] = int64(QBluetoothUuid__UserIndex)
	qt.EnumMap["bluetooth.QBluetoothUuid__BodyCompositionFeature"] = int64(QBluetoothUuid__BodyCompositionFeature)
	qt.EnumMap["bluetooth.QBluetoothUuid__BodyCompositionMeasurement"] = int64(QBluetoothUuid__BodyCompositionMeasurement)
	qt.EnumMap["bluetooth.QBluetoothUuid__WeightMeasurement"] = int64(QBluetoothUuid__WeightMeasurement)
	qt.EnumMap["bluetooth.QBluetoothUuid__WeightScaleFeature"] = int64(QBluetoothUuid__WeightScaleFeature)
	qt.EnumMap["bluetooth.QBluetoothUuid__UserControlPoint"] = int64(QBluetoothUuid__UserControlPoint)
	qt.EnumMap["bluetooth.QBluetoothUuid__MagneticFluxDensity2D"] = int64(QBluetoothUuid__MagneticFluxDensity2D)
	qt.EnumMap["bluetooth.QBluetoothUuid__MagneticFluxDensity3D"] = int64(QBluetoothUuid__MagneticFluxDensity3D)
	qt.EnumMap["bluetooth.QBluetoothUuid__Language"] = int64(QBluetoothUuid__Language)
	qt.EnumMap["bluetooth.QBluetoothUuid__BarometricPressureTrend"] = int64(QBluetoothUuid__BarometricPressureTrend)
	qt.EnumMap["bluetooth.QBluetoothUuid__UnknownDescriptorType"] = int64(QBluetoothUuid__UnknownDescriptorType)
	qt.EnumMap["bluetooth.QBluetoothUuid__CharacteristicExtendedProperties"] = int64(QBluetoothUuid__CharacteristicExtendedProperties)
	qt.EnumMap["bluetooth.QBluetoothUuid__CharacteristicUserDescription"] = int64(QBluetoothUuid__CharacteristicUserDescription)
	qt.EnumMap["bluetooth.QBluetoothUuid__ClientCharacteristicConfiguration"] = int64(QBluetoothUuid__ClientCharacteristicConfiguration)
	qt.EnumMap["bluetooth.QBluetoothUuid__ServerCharacteristicConfiguration"] = int64(QBluetoothUuid__ServerCharacteristicConfiguration)
	qt.EnumMap["bluetooth.QBluetoothUuid__CharacteristicPresentationFormat"] = int64(QBluetoothUuid__CharacteristicPresentationFormat)
	qt.EnumMap["bluetooth.QBluetoothUuid__CharacteristicAggregateFormat"] = int64(QBluetoothUuid__CharacteristicAggregateFormat)
	qt.EnumMap["bluetooth.QBluetoothUuid__ValidRange"] = int64(QBluetoothUuid__ValidRange)
	qt.EnumMap["bluetooth.QBluetoothUuid__ExternalReportReference"] = int64(QBluetoothUuid__ExternalReportReference)
	qt.EnumMap["bluetooth.QBluetoothUuid__ReportReference"] = int64(QBluetoothUuid__ReportReference)
	qt.EnumMap["bluetooth.QBluetoothUuid__EnvironmentalSensingConfiguration"] = int64(QBluetoothUuid__EnvironmentalSensingConfiguration)
	qt.EnumMap["bluetooth.QBluetoothUuid__EnvironmentalSensingMeasurement"] = int64(QBluetoothUuid__EnvironmentalSensingMeasurement)
	qt.EnumMap["bluetooth.QBluetoothUuid__EnvironmentalSensingTriggerSetting"] = int64(QBluetoothUuid__EnvironmentalSensingTriggerSetting)
	qt.ItfMap["bluetooth.QLowEnergyAdvertisingData_ITF"] = QLowEnergyAdvertisingData{}
	qt.FuncMap["bluetooth.NewQLowEnergyAdvertisingData"] = NewQLowEnergyAdvertisingData
	qt.FuncMap["bluetooth.NewQLowEnergyAdvertisingData2"] = NewQLowEnergyAdvertisingData2
	qt.FuncMap["bluetooth.QLowEnergyAdvertisingData_InvalidManufacturerId"] = QLowEnergyAdvertisingData_InvalidManufacturerId
	qt.EnumMap["bluetooth.QLowEnergyAdvertisingData__DiscoverabilityNone"] = int64(QLowEnergyAdvertisingData__DiscoverabilityNone)
	qt.EnumMap["bluetooth.QLowEnergyAdvertisingData__DiscoverabilityLimited"] = int64(QLowEnergyAdvertisingData__DiscoverabilityLimited)
	qt.EnumMap["bluetooth.QLowEnergyAdvertisingData__DiscoverabilityGeneral"] = int64(QLowEnergyAdvertisingData__DiscoverabilityGeneral)
	qt.ItfMap["bluetooth.QLowEnergyAdvertisingParameters_ITF"] = QLowEnergyAdvertisingParameters{}
	qt.FuncMap["bluetooth.NewQLowEnergyAdvertisingParameters"] = NewQLowEnergyAdvertisingParameters
	qt.FuncMap["bluetooth.NewQLowEnergyAdvertisingParameters2"] = NewQLowEnergyAdvertisingParameters2
	qt.EnumMap["bluetooth.QLowEnergyAdvertisingParameters__AdvInd"] = int64(QLowEnergyAdvertisingParameters__AdvInd)
	qt.EnumMap["bluetooth.QLowEnergyAdvertisingParameters__AdvScanInd"] = int64(QLowEnergyAdvertisingParameters__AdvScanInd)
	qt.EnumMap["bluetooth.QLowEnergyAdvertisingParameters__AdvNonConnInd"] = int64(QLowEnergyAdvertisingParameters__AdvNonConnInd)
	qt.EnumMap["bluetooth.QLowEnergyAdvertisingParameters__IgnoreWhiteList"] = int64(QLowEnergyAdvertisingParameters__IgnoreWhiteList)
	qt.EnumMap["bluetooth.QLowEnergyAdvertisingParameters__UseWhiteListForScanning"] = int64(QLowEnergyAdvertisingParameters__UseWhiteListForScanning)
	qt.EnumMap["bluetooth.QLowEnergyAdvertisingParameters__UseWhiteListForConnecting"] = int64(QLowEnergyAdvertisingParameters__UseWhiteListForConnecting)
	qt.EnumMap["bluetooth.QLowEnergyAdvertisingParameters__UseWhiteListForScanningAndConnecting"] = int64(QLowEnergyAdvertisingParameters__UseWhiteListForScanningAndConnecting)
	qt.ItfMap["bluetooth.QLowEnergyCharacteristic_ITF"] = QLowEnergyCharacteristic{}
	qt.FuncMap["bluetooth.NewQLowEnergyCharacteristic"] = NewQLowEnergyCharacteristic
	qt.FuncMap["bluetooth.NewQLowEnergyCharacteristic2"] = NewQLowEnergyCharacteristic2
	qt.EnumMap["bluetooth.QLowEnergyCharacteristic__Unknown"] = int64(QLowEnergyCharacteristic__Unknown)
	qt.EnumMap["bluetooth.QLowEnergyCharacteristic__Broadcasting"] = int64(QLowEnergyCharacteristic__Broadcasting)
	qt.EnumMap["bluetooth.QLowEnergyCharacteristic__Read"] = int64(QLowEnergyCharacteristic__Read)
	qt.EnumMap["bluetooth.QLowEnergyCharacteristic__WriteNoResponse"] = int64(QLowEnergyCharacteristic__WriteNoResponse)
	qt.EnumMap["bluetooth.QLowEnergyCharacteristic__Write"] = int64(QLowEnergyCharacteristic__Write)
	qt.EnumMap["bluetooth.QLowEnergyCharacteristic__Notify"] = int64(QLowEnergyCharacteristic__Notify)
	qt.EnumMap["bluetooth.QLowEnergyCharacteristic__Indicate"] = int64(QLowEnergyCharacteristic__Indicate)
	qt.EnumMap["bluetooth.QLowEnergyCharacteristic__WriteSigned"] = int64(QLowEnergyCharacteristic__WriteSigned)
	qt.EnumMap["bluetooth.QLowEnergyCharacteristic__ExtendedProperty"] = int64(QLowEnergyCharacteristic__ExtendedProperty)
	qt.ItfMap["bluetooth.QLowEnergyCharacteristicData_ITF"] = QLowEnergyCharacteristicData{}
	qt.FuncMap["bluetooth.NewQLowEnergyCharacteristicData"] = NewQLowEnergyCharacteristicData
	qt.FuncMap["bluetooth.NewQLowEnergyCharacteristicData2"] = NewQLowEnergyCharacteristicData2
	qt.ItfMap["bluetooth.QLowEnergyConnectionParameters_ITF"] = QLowEnergyConnectionParameters{}
	qt.FuncMap["bluetooth.NewQLowEnergyConnectionParameters"] = NewQLowEnergyConnectionParameters
	qt.FuncMap["bluetooth.NewQLowEnergyConnectionParameters2"] = NewQLowEnergyConnectionParameters2
	qt.ItfMap["bluetooth.QLowEnergyController_ITF"] = QLowEnergyController{}
	qt.FuncMap["bluetooth.QLowEnergyController_CreateCentral"] = QLowEnergyController_CreateCentral
	qt.FuncMap["bluetooth.QLowEnergyController_CreatePeripheral"] = QLowEnergyController_CreatePeripheral
	qt.EnumMap["bluetooth.QLowEnergyController__NoError"] = int64(QLowEnergyController__NoError)
	qt.EnumMap["bluetooth.QLowEnergyController__UnknownError"] = int64(QLowEnergyController__UnknownError)
	qt.EnumMap["bluetooth.QLowEnergyController__UnknownRemoteDeviceError"] = int64(QLowEnergyController__UnknownRemoteDeviceError)
	qt.EnumMap["bluetooth.QLowEnergyController__NetworkError"] = int64(QLowEnergyController__NetworkError)
	qt.EnumMap["bluetooth.QLowEnergyController__InvalidBluetoothAdapterError"] = int64(QLowEnergyController__InvalidBluetoothAdapterError)
	qt.EnumMap["bluetooth.QLowEnergyController__ConnectionError"] = int64(QLowEnergyController__ConnectionError)
	qt.EnumMap["bluetooth.QLowEnergyController__AdvertisingError"] = int64(QLowEnergyController__AdvertisingError)
	qt.EnumMap["bluetooth.QLowEnergyController__RemoteHostClosedError"] = int64(QLowEnergyController__RemoteHostClosedError)
	qt.EnumMap["bluetooth.QLowEnergyController__UnconnectedState"] = int64(QLowEnergyController__UnconnectedState)
	qt.EnumMap["bluetooth.QLowEnergyController__ConnectingState"] = int64(QLowEnergyController__ConnectingState)
	qt.EnumMap["bluetooth.QLowEnergyController__ConnectedState"] = int64(QLowEnergyController__ConnectedState)
	qt.EnumMap["bluetooth.QLowEnergyController__DiscoveringState"] = int64(QLowEnergyController__DiscoveringState)
	qt.EnumMap["bluetooth.QLowEnergyController__DiscoveredState"] = int64(QLowEnergyController__DiscoveredState)
	qt.EnumMap["bluetooth.QLowEnergyController__ClosingState"] = int64(QLowEnergyController__ClosingState)
	qt.EnumMap["bluetooth.QLowEnergyController__AdvertisingState"] = int64(QLowEnergyController__AdvertisingState)
	qt.EnumMap["bluetooth.QLowEnergyController__PublicAddress"] = int64(QLowEnergyController__PublicAddress)
	qt.EnumMap["bluetooth.QLowEnergyController__RandomAddress"] = int64(QLowEnergyController__RandomAddress)
	qt.EnumMap["bluetooth.QLowEnergyController__CentralRole"] = int64(QLowEnergyController__CentralRole)
	qt.EnumMap["bluetooth.QLowEnergyController__PeripheralRole"] = int64(QLowEnergyController__PeripheralRole)
	qt.ItfMap["bluetooth.QLowEnergyDescriptor_ITF"] = QLowEnergyDescriptor{}
	qt.FuncMap["bluetooth.NewQLowEnergyDescriptor"] = NewQLowEnergyDescriptor
	qt.FuncMap["bluetooth.NewQLowEnergyDescriptor2"] = NewQLowEnergyDescriptor2
	qt.ItfMap["bluetooth.QLowEnergyDescriptorData_ITF"] = QLowEnergyDescriptorData{}
	qt.FuncMap["bluetooth.NewQLowEnergyDescriptorData"] = NewQLowEnergyDescriptorData
	qt.FuncMap["bluetooth.NewQLowEnergyDescriptorData2"] = NewQLowEnergyDescriptorData2
	qt.FuncMap["bluetooth.NewQLowEnergyDescriptorData3"] = NewQLowEnergyDescriptorData3
	qt.ItfMap["bluetooth.QLowEnergyService_ITF"] = QLowEnergyService{}
	qt.EnumMap["bluetooth.QLowEnergyService__PrimaryService"] = int64(QLowEnergyService__PrimaryService)
	qt.EnumMap["bluetooth.QLowEnergyService__IncludedService"] = int64(QLowEnergyService__IncludedService)
	qt.EnumMap["bluetooth.QLowEnergyService__NoError"] = int64(QLowEnergyService__NoError)
	qt.EnumMap["bluetooth.QLowEnergyService__OperationError"] = int64(QLowEnergyService__OperationError)
	qt.EnumMap["bluetooth.QLowEnergyService__CharacteristicWriteError"] = int64(QLowEnergyService__CharacteristicWriteError)
	qt.EnumMap["bluetooth.QLowEnergyService__DescriptorWriteError"] = int64(QLowEnergyService__DescriptorWriteError)
	qt.EnumMap["bluetooth.QLowEnergyService__UnknownError"] = int64(QLowEnergyService__UnknownError)
	qt.EnumMap["bluetooth.QLowEnergyService__CharacteristicReadError"] = int64(QLowEnergyService__CharacteristicReadError)
	qt.EnumMap["bluetooth.QLowEnergyService__DescriptorReadError"] = int64(QLowEnergyService__DescriptorReadError)
	qt.EnumMap["bluetooth.QLowEnergyService__InvalidService"] = int64(QLowEnergyService__InvalidService)
	qt.EnumMap["bluetooth.QLowEnergyService__DiscoveryRequired"] = int64(QLowEnergyService__DiscoveryRequired)
	qt.EnumMap["bluetooth.QLowEnergyService__DiscoveringServices"] = int64(QLowEnergyService__DiscoveringServices)
	qt.EnumMap["bluetooth.QLowEnergyService__ServiceDiscovered"] = int64(QLowEnergyService__ServiceDiscovered)
	qt.EnumMap["bluetooth.QLowEnergyService__LocalService"] = int64(QLowEnergyService__LocalService)
	qt.EnumMap["bluetooth.QLowEnergyService__WriteWithResponse"] = int64(QLowEnergyService__WriteWithResponse)
	qt.EnumMap["bluetooth.QLowEnergyService__WriteWithoutResponse"] = int64(QLowEnergyService__WriteWithoutResponse)
	qt.EnumMap["bluetooth.QLowEnergyService__WriteSigned"] = int64(QLowEnergyService__WriteSigned)
	qt.ItfMap["bluetooth.QLowEnergyServiceData_ITF"] = QLowEnergyServiceData{}
	qt.FuncMap["bluetooth.NewQLowEnergyServiceData"] = NewQLowEnergyServiceData
	qt.FuncMap["bluetooth.NewQLowEnergyServiceData2"] = NewQLowEnergyServiceData2
	qt.EnumMap["bluetooth.QLowEnergyServiceData__ServiceTypePrimary"] = int64(QLowEnergyServiceData__ServiceTypePrimary)
	qt.EnumMap["bluetooth.QLowEnergyServiceData__ServiceTypeSecondary"] = int64(QLowEnergyServiceData__ServiceTypeSecondary)
}
