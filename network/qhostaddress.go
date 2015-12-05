package network

//#include "network.h"
import "C"
import (
	"log"
	"unsafe"
)

type QHostAddress struct {
	ptr unsafe.Pointer
}

type QHostAddress_ITF interface {
	QHostAddress_PTR() *QHostAddress
}

func (p *QHostAddress) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QHostAddress) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQHostAddress(ptr QHostAddress_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHostAddress_PTR().Pointer()
	}
	return nil
}

func NewQHostAddressFromPointer(ptr unsafe.Pointer) *QHostAddress {
	var n = new(QHostAddress)
	n.SetPointer(ptr)
	return n
}

func (ptr *QHostAddress) QHostAddress_PTR() *QHostAddress {
	return ptr
}

//QHostAddress::SpecialAddress
type QHostAddress__SpecialAddress int64

const (
	QHostAddress__Null          = QHostAddress__SpecialAddress(0)
	QHostAddress__Broadcast     = QHostAddress__SpecialAddress(1)
	QHostAddress__LocalHost     = QHostAddress__SpecialAddress(2)
	QHostAddress__LocalHostIPv6 = QHostAddress__SpecialAddress(3)
	QHostAddress__Any           = QHostAddress__SpecialAddress(4)
	QHostAddress__AnyIPv6       = QHostAddress__SpecialAddress(5)
	QHostAddress__AnyIPv4       = QHostAddress__SpecialAddress(6)
)

func NewQHostAddress() *QHostAddress {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHostAddress::QHostAddress")
		}
	}()

	return NewQHostAddressFromPointer(C.QHostAddress_NewQHostAddress())
}

func NewQHostAddress9(address QHostAddress__SpecialAddress) *QHostAddress {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHostAddress::QHostAddress")
		}
	}()

	return NewQHostAddressFromPointer(C.QHostAddress_NewQHostAddress9(C.int(address)))
}

func NewQHostAddress8(address QHostAddress_ITF) *QHostAddress {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHostAddress::QHostAddress")
		}
	}()

	return NewQHostAddressFromPointer(C.QHostAddress_NewQHostAddress8(PointerFromQHostAddress(address)))
}

func NewQHostAddress7(address string) *QHostAddress {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHostAddress::QHostAddress")
		}
	}()

	return NewQHostAddressFromPointer(C.QHostAddress_NewQHostAddress7(C.CString(address)))
}

func (ptr *QHostAddress) Clear() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHostAddress::clear")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHostAddress_Clear(ptr.Pointer())
	}
}

func (ptr *QHostAddress) IsInSubnet(subnet QHostAddress_ITF, netmask int) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHostAddress::isInSubnet")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QHostAddress_IsInSubnet(ptr.Pointer(), PointerFromQHostAddress(subnet), C.int(netmask)) != 0
	}
	return false
}

func (ptr *QHostAddress) IsLoopback() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHostAddress::isLoopback")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QHostAddress_IsLoopback(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHostAddress) IsNull() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHostAddress::isNull")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QHostAddress_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QHostAddress) Protocol() QAbstractSocket__NetworkLayerProtocol {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHostAddress::protocol")
		}
	}()

	if ptr.Pointer() != nil {
		return QAbstractSocket__NetworkLayerProtocol(C.QHostAddress_Protocol(ptr.Pointer()))
	}
	return 0
}

func (ptr *QHostAddress) ScopeId() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHostAddress::scopeId")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QHostAddress_ScopeId(ptr.Pointer()))
	}
	return ""
}

func (ptr *QHostAddress) SetAddress5(address string) bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHostAddress::setAddress")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QHostAddress_SetAddress5(ptr.Pointer(), C.CString(address)) != 0
	}
	return false
}

func (ptr *QHostAddress) SetScopeId(id string) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHostAddress::setScopeId")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHostAddress_SetScopeId(ptr.Pointer(), C.CString(id))
	}
}

func (ptr *QHostAddress) ToString() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHostAddress::toString")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QHostAddress_ToString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QHostAddress) DestroyQHostAddress() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QHostAddress::~QHostAddress")
		}
	}()

	if ptr.Pointer() != nil {
		C.QHostAddress_DestroyQHostAddress(ptr.Pointer())
	}
}
