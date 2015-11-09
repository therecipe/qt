package network

//#include "qnetworkinterface.h"
import "C"
import (
	"unsafe"
)

type QNetworkInterface struct {
	ptr unsafe.Pointer
}

type QNetworkInterface_ITF interface {
	QNetworkInterface_PTR() *QNetworkInterface
}

func (p *QNetworkInterface) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QNetworkInterface) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQNetworkInterface(ptr QNetworkInterface_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkInterface_PTR().Pointer()
	}
	return nil
}

func NewQNetworkInterfaceFromPointer(ptr unsafe.Pointer) *QNetworkInterface {
	var n = new(QNetworkInterface)
	n.SetPointer(ptr)
	return n
}

func (ptr *QNetworkInterface) QNetworkInterface_PTR() *QNetworkInterface {
	return ptr
}

//QNetworkInterface::InterfaceFlag
type QNetworkInterface__InterfaceFlag int64

const (
	QNetworkInterface__IsUp           = QNetworkInterface__InterfaceFlag(0x1)
	QNetworkInterface__IsRunning      = QNetworkInterface__InterfaceFlag(0x2)
	QNetworkInterface__CanBroadcast   = QNetworkInterface__InterfaceFlag(0x4)
	QNetworkInterface__IsLoopBack     = QNetworkInterface__InterfaceFlag(0x8)
	QNetworkInterface__IsPointToPoint = QNetworkInterface__InterfaceFlag(0x10)
	QNetworkInterface__CanMulticast   = QNetworkInterface__InterfaceFlag(0x20)
)

func NewQNetworkInterface() *QNetworkInterface {
	return NewQNetworkInterfaceFromPointer(C.QNetworkInterface_NewQNetworkInterface())
}

func NewQNetworkInterface2(other QNetworkInterface_ITF) *QNetworkInterface {
	return NewQNetworkInterfaceFromPointer(C.QNetworkInterface_NewQNetworkInterface2(PointerFromQNetworkInterface(other)))
}

func (ptr *QNetworkInterface) Flags() QNetworkInterface__InterfaceFlag {
	if ptr.Pointer() != nil {
		return QNetworkInterface__InterfaceFlag(C.QNetworkInterface_Flags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkInterface) HardwareAddress() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkInterface_HardwareAddress(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkInterface) HumanReadableName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkInterface_HumanReadableName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkInterface) Index() int {
	if ptr.Pointer() != nil {
		return int(C.QNetworkInterface_Index(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkInterface) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QNetworkInterface_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkInterface) Name() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkInterface_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkInterface) Swap(other QNetworkInterface_ITF) {
	if ptr.Pointer() != nil {
		C.QNetworkInterface_Swap(ptr.Pointer(), PointerFromQNetworkInterface(other))
	}
}

func (ptr *QNetworkInterface) DestroyQNetworkInterface() {
	if ptr.Pointer() != nil {
		C.QNetworkInterface_DestroyQNetworkInterface(ptr.Pointer())
	}
}
