package network

//#include "qnetworkinterface.h"
import "C"
import (
	"unsafe"
)

type QNetworkInterface struct {
	ptr unsafe.Pointer
}

type QNetworkInterfaceITF interface {
	QNetworkInterfacePTR() *QNetworkInterface
}

func (p *QNetworkInterface) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QNetworkInterface) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQNetworkInterface(ptr QNetworkInterfaceITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QNetworkInterfacePTR().Pointer()
	}
	return nil
}

func QNetworkInterfaceFromPointer(ptr unsafe.Pointer) *QNetworkInterface {
	var n = new(QNetworkInterface)
	n.SetPointer(ptr)
	return n
}

func (ptr *QNetworkInterface) QNetworkInterfacePTR() *QNetworkInterface {
	return ptr
}

//QNetworkInterface::InterfaceFlag
type QNetworkInterface__InterfaceFlag int

var (
	QNetworkInterface__IsUp           = QNetworkInterface__InterfaceFlag(0x1)
	QNetworkInterface__IsRunning      = QNetworkInterface__InterfaceFlag(0x2)
	QNetworkInterface__CanBroadcast   = QNetworkInterface__InterfaceFlag(0x4)
	QNetworkInterface__IsLoopBack     = QNetworkInterface__InterfaceFlag(0x8)
	QNetworkInterface__IsPointToPoint = QNetworkInterface__InterfaceFlag(0x10)
	QNetworkInterface__CanMulticast   = QNetworkInterface__InterfaceFlag(0x20)
)

func NewQNetworkInterface() *QNetworkInterface {
	return QNetworkInterfaceFromPointer(unsafe.Pointer(C.QNetworkInterface_NewQNetworkInterface()))
}

func NewQNetworkInterface2(other QNetworkInterfaceITF) *QNetworkInterface {
	return QNetworkInterfaceFromPointer(unsafe.Pointer(C.QNetworkInterface_NewQNetworkInterface2(C.QtObjectPtr(PointerFromQNetworkInterface(other)))))
}

func (ptr *QNetworkInterface) Flags() QNetworkInterface__InterfaceFlag {
	if ptr.Pointer() != nil {
		return QNetworkInterface__InterfaceFlag(C.QNetworkInterface_Flags(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNetworkInterface) HardwareAddress() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkInterface_HardwareAddress(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QNetworkInterface) HumanReadableName() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkInterface_HumanReadableName(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QNetworkInterface) Index() int {
	if ptr.Pointer() != nil {
		return int(C.QNetworkInterface_Index(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QNetworkInterface) IsValid() bool {
	if ptr.Pointer() != nil {
		return C.QNetworkInterface_IsValid(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QNetworkInterface) Name() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkInterface_Name(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QNetworkInterface) Swap(other QNetworkInterfaceITF) {
	if ptr.Pointer() != nil {
		C.QNetworkInterface_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQNetworkInterface(other)))
	}
}

func (ptr *QNetworkInterface) DestroyQNetworkInterface() {
	if ptr.Pointer() != nil {
		C.QNetworkInterface_DestroyQNetworkInterface(C.QtObjectPtr(ptr.Pointer()))
	}
}
