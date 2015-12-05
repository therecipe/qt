package network

//#include "network.h"
import "C"
import (
	"log"
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
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkInterface::QNetworkInterface")
		}
	}()

	return NewQNetworkInterfaceFromPointer(C.QNetworkInterface_NewQNetworkInterface())
}

func NewQNetworkInterface2(other QNetworkInterface_ITF) *QNetworkInterface {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkInterface::QNetworkInterface")
		}
	}()

	return NewQNetworkInterfaceFromPointer(C.QNetworkInterface_NewQNetworkInterface2(PointerFromQNetworkInterface(other)))
}

func (ptr *QNetworkInterface) Flags() QNetworkInterface__InterfaceFlag {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkInterface::flags")
		}
	}()

	if ptr.Pointer() != nil {
		return QNetworkInterface__InterfaceFlag(C.QNetworkInterface_Flags(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkInterface) HardwareAddress() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkInterface::hardwareAddress")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkInterface_HardwareAddress(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkInterface) HumanReadableName() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkInterface::humanReadableName")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkInterface_HumanReadableName(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkInterface) Index() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkInterface::index")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QNetworkInterface_Index(ptr.Pointer()))
	}
	return 0
}

func (ptr *QNetworkInterface) IsValid() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkInterface::isValid")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QNetworkInterface_IsValid(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QNetworkInterface) Name() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkInterface::name")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QNetworkInterface_Name(ptr.Pointer()))
	}
	return ""
}

func (ptr *QNetworkInterface) Swap(other QNetworkInterface_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkInterface::swap")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNetworkInterface_Swap(ptr.Pointer(), PointerFromQNetworkInterface(other))
	}
}

func (ptr *QNetworkInterface) DestroyQNetworkInterface() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QNetworkInterface::~QNetworkInterface")
		}
	}()

	if ptr.Pointer() != nil {
		C.QNetworkInterface_DestroyQNetworkInterface(ptr.Pointer())
	}
}
