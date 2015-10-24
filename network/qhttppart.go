package network

//#include "qhttppart.h"
import "C"
import (
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QHttpPart struct {
	ptr unsafe.Pointer
}

type QHttpPartITF interface {
	QHttpPartPTR() *QHttpPart
}

func (p *QHttpPart) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QHttpPart) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQHttpPart(ptr QHttpPartITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHttpPartPTR().Pointer()
	}
	return nil
}

func QHttpPartFromPointer(ptr unsafe.Pointer) *QHttpPart {
	var n = new(QHttpPart)
	n.SetPointer(ptr)
	return n
}

func (ptr *QHttpPart) QHttpPartPTR() *QHttpPart {
	return ptr
}

func NewQHttpPart() *QHttpPart {
	return QHttpPartFromPointer(unsafe.Pointer(C.QHttpPart_NewQHttpPart()))
}

func NewQHttpPart2(other QHttpPartITF) *QHttpPart {
	return QHttpPartFromPointer(unsafe.Pointer(C.QHttpPart_NewQHttpPart2(C.QtObjectPtr(PointerFromQHttpPart(other)))))
}

func (ptr *QHttpPart) SetBody(body core.QByteArrayITF) {
	if ptr.Pointer() != nil {
		C.QHttpPart_SetBody(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(body)))
	}
}

func (ptr *QHttpPart) SetBodyDevice(device core.QIODeviceITF) {
	if ptr.Pointer() != nil {
		C.QHttpPart_SetBodyDevice(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQIODevice(device)))
	}
}

func (ptr *QHttpPart) SetHeader(header QNetworkRequest__KnownHeaders, value string) {
	if ptr.Pointer() != nil {
		C.QHttpPart_SetHeader(C.QtObjectPtr(ptr.Pointer()), C.int(header), C.CString(value))
	}
}

func (ptr *QHttpPart) SetRawHeader(headerName core.QByteArrayITF, headerValue core.QByteArrayITF) {
	if ptr.Pointer() != nil {
		C.QHttpPart_SetRawHeader(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(headerName)), C.QtObjectPtr(core.PointerFromQByteArray(headerValue)))
	}
}

func (ptr *QHttpPart) Swap(other QHttpPartITF) {
	if ptr.Pointer() != nil {
		C.QHttpPart_Swap(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQHttpPart(other)))
	}
}

func (ptr *QHttpPart) DestroyQHttpPart() {
	if ptr.Pointer() != nil {
		C.QHttpPart_DestroyQHttpPart(C.QtObjectPtr(ptr.Pointer()))
	}
}
