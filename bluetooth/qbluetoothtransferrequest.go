package bluetooth

//#include "bluetooth.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"log"
	"unsafe"
)

type QBluetoothTransferRequest struct {
	ptr unsafe.Pointer
}

type QBluetoothTransferRequest_ITF interface {
	QBluetoothTransferRequest_PTR() *QBluetoothTransferRequest
}

func (p *QBluetoothTransferRequest) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QBluetoothTransferRequest) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQBluetoothTransferRequest(ptr QBluetoothTransferRequest_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothTransferRequest_PTR().Pointer()
	}
	return nil
}

func NewQBluetoothTransferRequestFromPointer(ptr unsafe.Pointer) *QBluetoothTransferRequest {
	var n = new(QBluetoothTransferRequest)
	n.SetPointer(ptr)
	return n
}

func (ptr *QBluetoothTransferRequest) QBluetoothTransferRequest_PTR() *QBluetoothTransferRequest {
	return ptr
}

//QBluetoothTransferRequest::Attribute
type QBluetoothTransferRequest__Attribute int64

const (
	QBluetoothTransferRequest__DescriptionAttribute = QBluetoothTransferRequest__Attribute(0)
	QBluetoothTransferRequest__TimeAttribute        = QBluetoothTransferRequest__Attribute(1)
	QBluetoothTransferRequest__TypeAttribute        = QBluetoothTransferRequest__Attribute(2)
	QBluetoothTransferRequest__LengthAttribute      = QBluetoothTransferRequest__Attribute(3)
	QBluetoothTransferRequest__NameAttribute        = QBluetoothTransferRequest__Attribute(4)
)

func NewQBluetoothTransferRequest(address QBluetoothAddress_ITF) *QBluetoothTransferRequest {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothTransferRequest::QBluetoothTransferRequest")
		}
	}()

	return NewQBluetoothTransferRequestFromPointer(C.QBluetoothTransferRequest_NewQBluetoothTransferRequest(PointerFromQBluetoothAddress(address)))
}

func NewQBluetoothTransferRequest2(other QBluetoothTransferRequest_ITF) *QBluetoothTransferRequest {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothTransferRequest::QBluetoothTransferRequest")
		}
	}()

	return NewQBluetoothTransferRequestFromPointer(C.QBluetoothTransferRequest_NewQBluetoothTransferRequest2(PointerFromQBluetoothTransferRequest(other)))
}

func (ptr *QBluetoothTransferRequest) Attribute(code QBluetoothTransferRequest__Attribute, defaultValue core.QVariant_ITF) *core.QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothTransferRequest::attribute")
		}
	}()

	if ptr.Pointer() != nil {
		return core.NewQVariantFromPointer(C.QBluetoothTransferRequest_Attribute(ptr.Pointer(), C.int(code), core.PointerFromQVariant(defaultValue)))
	}
	return nil
}

func (ptr *QBluetoothTransferRequest) SetAttribute(code QBluetoothTransferRequest__Attribute, value core.QVariant_ITF) {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothTransferRequest::setAttribute")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBluetoothTransferRequest_SetAttribute(ptr.Pointer(), C.int(code), core.PointerFromQVariant(value))
	}
}

func (ptr *QBluetoothTransferRequest) DestroyQBluetoothTransferRequest() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QBluetoothTransferRequest::~QBluetoothTransferRequest")
		}
	}()

	if ptr.Pointer() != nil {
		C.QBluetoothTransferRequest_DestroyQBluetoothTransferRequest(ptr.Pointer())
	}
}
