package bluetooth

//#include "qbluetoothaddress.h"
import "C"
import (
	"unsafe"
)

type QBluetoothAddress struct {
	ptr unsafe.Pointer
}

type QBluetoothAddressITF interface {
	QBluetoothAddressPTR() *QBluetoothAddress
}

func (p *QBluetoothAddress) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QBluetoothAddress) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQBluetoothAddress(ptr QBluetoothAddressITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothAddressPTR().Pointer()
	}
	return nil
}

func QBluetoothAddressFromPointer(ptr unsafe.Pointer) *QBluetoothAddress {
	var n = new(QBluetoothAddress)
	n.SetPointer(ptr)
	return n
}

func (ptr *QBluetoothAddress) QBluetoothAddressPTR() *QBluetoothAddress {
	return ptr
}

func NewQBluetoothAddress() *QBluetoothAddress {
	return QBluetoothAddressFromPointer(unsafe.Pointer(C.QBluetoothAddress_NewQBluetoothAddress()))
}

func NewQBluetoothAddress4(other QBluetoothAddressITF) *QBluetoothAddress {
	return QBluetoothAddressFromPointer(unsafe.Pointer(C.QBluetoothAddress_NewQBluetoothAddress4(C.QtObjectPtr(PointerFromQBluetoothAddress(other)))))
}

func NewQBluetoothAddress3(address string) *QBluetoothAddress {
	return QBluetoothAddressFromPointer(unsafe.Pointer(C.QBluetoothAddress_NewQBluetoothAddress3(C.CString(address))))
}

func (ptr *QBluetoothAddress) Clear() {
	if ptr.Pointer() != nil {
		C.QBluetoothAddress_Clear(C.QtObjectPtr(ptr.Pointer()))
	}
}

func (ptr *QBluetoothAddress) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QBluetoothAddress_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QBluetoothAddress) ToString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothAddress_ToString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QBluetoothAddress) DestroyQBluetoothAddress() {
	if ptr.Pointer() != nil {
		C.QBluetoothAddress_DestroyQBluetoothAddress(C.QtObjectPtr(ptr.Pointer()))
	}
}
