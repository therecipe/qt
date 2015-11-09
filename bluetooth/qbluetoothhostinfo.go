package bluetooth

//#include "qbluetoothhostinfo.h"
import "C"
import (
	"unsafe"
)

type QBluetoothHostInfo struct {
	ptr unsafe.Pointer
}

type QBluetoothHostInfo_ITF interface {
	QBluetoothHostInfo_PTR() *QBluetoothHostInfo
}

func (p *QBluetoothHostInfo) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QBluetoothHostInfo) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
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

func (ptr *QBluetoothHostInfo) QBluetoothHostInfo_PTR() *QBluetoothHostInfo {
	return ptr
}

func NewQBluetoothHostInfo() *QBluetoothHostInfo {
	return NewQBluetoothHostInfoFromPointer(C.QBluetoothHostInfo_NewQBluetoothHostInfo())
}

func NewQBluetoothHostInfo2(other QBluetoothHostInfo_ITF) *QBluetoothHostInfo {
	return NewQBluetoothHostInfoFromPointer(C.QBluetoothHostInfo_NewQBluetoothHostInfo2(PointerFromQBluetoothHostInfo(other)))
}

func (ptr *QBluetoothHostInfo) Name() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothHostInfo_Name(ptr.Pointer()))
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
		C.QBluetoothHostInfo_SetName(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QBluetoothHostInfo) DestroyQBluetoothHostInfo() {
	if ptr.Pointer() != nil {
		C.QBluetoothHostInfo_DestroyQBluetoothHostInfo(ptr.Pointer())
	}
}
