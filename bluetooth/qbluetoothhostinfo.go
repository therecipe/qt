package bluetooth

//#include "bluetooth.h"
import "C"
import (
	"github.com/therecipe/qt"
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
	defer qt.Recovering("QBluetoothHostInfo::QBluetoothHostInfo")

	return NewQBluetoothHostInfoFromPointer(C.QBluetoothHostInfo_NewQBluetoothHostInfo())
}

func NewQBluetoothHostInfo2(other QBluetoothHostInfo_ITF) *QBluetoothHostInfo {
	defer qt.Recovering("QBluetoothHostInfo::QBluetoothHostInfo")

	return NewQBluetoothHostInfoFromPointer(C.QBluetoothHostInfo_NewQBluetoothHostInfo2(PointerFromQBluetoothHostInfo(other)))
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
		C.QBluetoothHostInfo_SetName(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QBluetoothHostInfo) DestroyQBluetoothHostInfo() {
	defer qt.Recovering("QBluetoothHostInfo::~QBluetoothHostInfo")

	if ptr.Pointer() != nil {
		C.QBluetoothHostInfo_DestroyQBluetoothHostInfo(ptr.Pointer())
	}
}
