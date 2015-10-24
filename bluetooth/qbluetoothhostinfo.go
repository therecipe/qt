package bluetooth

//#include "qbluetoothhostinfo.h"
import "C"
import (
	"unsafe"
)

type QBluetoothHostInfo struct {
	ptr unsafe.Pointer
}

type QBluetoothHostInfoITF interface {
	QBluetoothHostInfoPTR() *QBluetoothHostInfo
}

func (p *QBluetoothHostInfo) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QBluetoothHostInfo) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQBluetoothHostInfo(ptr QBluetoothHostInfoITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothHostInfoPTR().Pointer()
	}
	return nil
}

func QBluetoothHostInfoFromPointer(ptr unsafe.Pointer) *QBluetoothHostInfo {
	var n = new(QBluetoothHostInfo)
	n.SetPointer(ptr)
	return n
}

func (ptr *QBluetoothHostInfo) QBluetoothHostInfoPTR() *QBluetoothHostInfo {
	return ptr
}

func NewQBluetoothHostInfo() *QBluetoothHostInfo {
	return QBluetoothHostInfoFromPointer(unsafe.Pointer(C.QBluetoothHostInfo_NewQBluetoothHostInfo()))
}

func NewQBluetoothHostInfo2(other QBluetoothHostInfoITF) *QBluetoothHostInfo {
	return QBluetoothHostInfoFromPointer(unsafe.Pointer(C.QBluetoothHostInfo_NewQBluetoothHostInfo2(C.QtObjectPtr(PointerFromQBluetoothHostInfo(other)))))
}

func (ptr *QBluetoothHostInfo) Name() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothHostInfo_Name(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QBluetoothHostInfo) SetAddress(address QBluetoothAddressITF) {
	if ptr.Pointer() != nil {
		C.QBluetoothHostInfo_SetAddress(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQBluetoothAddress(address)))
	}
}

func (ptr *QBluetoothHostInfo) SetName(name string) {
	if ptr.Pointer() != nil {
		C.QBluetoothHostInfo_SetName(C.QtObjectPtr(ptr.Pointer()), C.CString(name))
	}
}

func (ptr *QBluetoothHostInfo) DestroyQBluetoothHostInfo() {
	if ptr.Pointer() != nil {
		C.QBluetoothHostInfo_DestroyQBluetoothHostInfo(C.QtObjectPtr(ptr.Pointer()))
	}
}
