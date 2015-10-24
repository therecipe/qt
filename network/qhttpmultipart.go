package network

//#include "qhttpmultipart.h"
import "C"
import (
	"github.com/therecipe/qt"
	"github.com/therecipe/qt/core"
	"unsafe"
)

type QHttpMultiPart struct {
	core.QObject
}

type QHttpMultiPartITF interface {
	core.QObjectITF
	QHttpMultiPartPTR() *QHttpMultiPart
}

func PointerFromQHttpMultiPart(ptr QHttpMultiPartITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHttpMultiPartPTR().Pointer()
	}
	return nil
}

func QHttpMultiPartFromPointer(ptr unsafe.Pointer) *QHttpMultiPart {
	var n = new(QHttpMultiPart)
	n.SetPointer(ptr)
	if n.ObjectName() == "" {
		n.SetObjectName("QHttpMultiPart_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QHttpMultiPart) QHttpMultiPartPTR() *QHttpMultiPart {
	return ptr
}

//QHttpMultiPart::ContentType
type QHttpMultiPart__ContentType int

var (
	QHttpMultiPart__MixedType       = QHttpMultiPart__ContentType(0)
	QHttpMultiPart__RelatedType     = QHttpMultiPart__ContentType(1)
	QHttpMultiPart__FormDataType    = QHttpMultiPart__ContentType(2)
	QHttpMultiPart__AlternativeType = QHttpMultiPart__ContentType(3)
)

func NewQHttpMultiPart2(contentType QHttpMultiPart__ContentType, parent core.QObjectITF) *QHttpMultiPart {
	return QHttpMultiPartFromPointer(unsafe.Pointer(C.QHttpMultiPart_NewQHttpMultiPart2(C.int(contentType), C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func NewQHttpMultiPart(parent core.QObjectITF) *QHttpMultiPart {
	return QHttpMultiPartFromPointer(unsafe.Pointer(C.QHttpMultiPart_NewQHttpMultiPart(C.QtObjectPtr(core.PointerFromQObject(parent)))))
}

func (ptr *QHttpMultiPart) Append(httpPart QHttpPartITF) {
	if ptr.Pointer() != nil {
		C.QHttpMultiPart_Append(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQHttpPart(httpPart)))
	}
}

func (ptr *QHttpMultiPart) SetBoundary(boundary core.QByteArrayITF) {
	if ptr.Pointer() != nil {
		C.QHttpMultiPart_SetBoundary(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(core.PointerFromQByteArray(boundary)))
	}
}

func (ptr *QHttpMultiPart) SetContentType(contentType QHttpMultiPart__ContentType) {
	if ptr.Pointer() != nil {
		C.QHttpMultiPart_SetContentType(C.QtObjectPtr(ptr.Pointer()), C.int(contentType))
	}
}

func (ptr *QHttpMultiPart) DestroyQHttpMultiPart() {
	if ptr.Pointer() != nil {
		C.QHttpMultiPart_DestroyQHttpMultiPart(C.QtObjectPtr(ptr.Pointer()))
		ptr.SetPointer(nil)
	}
}
