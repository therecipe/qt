package bluetooth

//#include "qbluetoothtransferrequest.h"
import "C"
import (
	"unsafe"
)

type QBluetoothTransferRequest struct {
	ptr unsafe.Pointer
}

type QBluetoothTransferRequestITF interface {
	QBluetoothTransferRequestPTR() *QBluetoothTransferRequest
}

func (p *QBluetoothTransferRequest) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QBluetoothTransferRequest) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQBluetoothTransferRequest(ptr QBluetoothTransferRequestITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QBluetoothTransferRequestPTR().Pointer()
	}
	return nil
}

func QBluetoothTransferRequestFromPointer(ptr unsafe.Pointer) *QBluetoothTransferRequest {
	var n = new(QBluetoothTransferRequest)
	n.SetPointer(ptr)
	return n
}

func (ptr *QBluetoothTransferRequest) QBluetoothTransferRequestPTR() *QBluetoothTransferRequest {
	return ptr
}

//QBluetoothTransferRequest::Attribute
type QBluetoothTransferRequest__Attribute int

var (
	QBluetoothTransferRequest__DescriptionAttribute = QBluetoothTransferRequest__Attribute(0)
	QBluetoothTransferRequest__TimeAttribute        = QBluetoothTransferRequest__Attribute(1)
	QBluetoothTransferRequest__TypeAttribute        = QBluetoothTransferRequest__Attribute(2)
	QBluetoothTransferRequest__LengthAttribute      = QBluetoothTransferRequest__Attribute(3)
	QBluetoothTransferRequest__NameAttribute        = QBluetoothTransferRequest__Attribute(4)
)

func NewQBluetoothTransferRequest(address QBluetoothAddressITF) *QBluetoothTransferRequest {
	return QBluetoothTransferRequestFromPointer(unsafe.Pointer(C.QBluetoothTransferRequest_NewQBluetoothTransferRequest(C.QtObjectPtr(PointerFromQBluetoothAddress(address)))))
}

func NewQBluetoothTransferRequest2(other QBluetoothTransferRequestITF) *QBluetoothTransferRequest {
	return QBluetoothTransferRequestFromPointer(unsafe.Pointer(C.QBluetoothTransferRequest_NewQBluetoothTransferRequest2(C.QtObjectPtr(PointerFromQBluetoothTransferRequest(other)))))
}

func (ptr *QBluetoothTransferRequest) Attribute(code QBluetoothTransferRequest__Attribute, defaultValue string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QBluetoothTransferRequest_Attribute(C.QtObjectPtr(ptr.Pointer()), C.int(code), C.CString(defaultValue)))
	}
	return ""
}

func (ptr *QBluetoothTransferRequest) SetAttribute(code QBluetoothTransferRequest__Attribute, value string) {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferRequest_SetAttribute(C.QtObjectPtr(ptr.Pointer()), C.int(code), C.CString(value))
	}
}

func (ptr *QBluetoothTransferRequest) DestroyQBluetoothTransferRequest() {
	if ptr.Pointer() != nil {
		C.QBluetoothTransferRequest_DestroyQBluetoothTransferRequest(C.QtObjectPtr(ptr.Pointer()))
	}
}
