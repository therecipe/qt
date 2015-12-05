package bluetooth

//#include "bluetooth.h"
import "C"
import (
	"log"
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothHostInfo::QBluetoothHostInfo")
		}
	}()

	return NewQBluetoothHostInfoFromPointer(C.QBluetoothHostInfo_NewQBluetoothHostInfo())
}

func NewQBluetoothHostInfo2(other QBluetoothHostInfo_ITF) *QBluetoothHostInfo {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothHostInfo::QBluetoothHostInfo")
		}
	}()

	return NewQBluetoothHostInfoFromPointer(C.QBluetoothHostInfo_NewQBluetoothHostInfo2(PointerFromQBluetoothHostInfo(other)))
}

func (ptr *QBluetoothHostInfo) Name() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothHostInfo::name")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothHostInfo_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QBluetoothHostInfo) SetAddress(address QBluetoothAddress_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothHostInfo::setAddress")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBluetoothHostInfo_SetAddress(ptr.Pointer(), PointerFromQBluetoothAddress(address))
	}
}

func (ptr *QBluetoothHostInfo) SetName(name string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothHostInfo::setName")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBluetoothHostInfo_SetName(ptr.Pointer(), C.CString(name))
	}
}

func (ptr *QBluetoothHostInfo) DestroyQBluetoothHostInfo() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothHostInfo::~QBluetoothHostInfo")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBluetoothHostInfo_DestroyQBluetoothHostInfo(ptr.Pointer())
	}
}
