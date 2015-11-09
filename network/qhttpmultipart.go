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

type QHttpMultiPart_ITF interface {
	core.QObject_ITF
	QHttpMultiPart_PTR() *QHttpMultiPart
}

func PointerFromQHttpMultiPart(ptr QHttpMultiPart_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QHttpMultiPart_PTR().Pointer()
	}
	return nil
}

func NewQHttpMultiPartFromPointer(ptr unsafe.Pointer) *QHttpMultiPart {
	var n = new(QHttpMultiPart)
	n.SetPointer(ptr)
	if len(n.ObjectName()) == 0 {
		n.SetObjectName("QHttpMultiPart_" + qt.RandomIdentifier())
	}
	return n
}

func (ptr *QHttpMultiPart) QHttpMultiPart_PTR() *QHttpMultiPart {
	return ptr
}

//QHttpMultiPart::ContentType
type QHttpMultiPart__ContentType int64

const (
	QHttpMultiPart__MixedType       = QHttpMultiPart__ContentType(0)
	QHttpMultiPart__RelatedType     = QHttpMultiPart__ContentType(1)
	QHttpMultiPart__FormDataType    = QHttpMultiPart__ContentType(2)
	QHttpMultiPart__AlternativeType = QHttpMultiPart__ContentType(3)
)

func NewQHttpMultiPart2(contentType QHttpMultiPart__ContentType, parent core.QObject_ITF) *QHttpMultiPart {
	return NewQHttpMultiPartFromPointer(C.QHttpMultiPart_NewQHttpMultiPart2(C.int(contentType), core.PointerFromQObject(parent)))
}

func NewQHttpMultiPart(parent core.QObject_ITF) *QHttpMultiPart {
	return NewQHttpMultiPartFromPointer(C.QHttpMultiPart_NewQHttpMultiPart(core.PointerFromQObject(parent)))
}

func (ptr *QHttpMultiPart) Append(httpPart QHttpPart_ITF) {
	if ptr.Pointer() != nil {
		C.QHttpMultiPart_Append(ptr.Pointer(), PointerFromQHttpPart(httpPart))
	}
}

func (ptr *QHttpMultiPart) Boundary() *core.QByteArray {
	if ptr.Pointer() != nil {
		return core.NewQByteArrayFromPointer(C.QHttpMultiPart_Boundary(ptr.Pointer()))
	}
	return nil
}

func (ptr *QHttpMultiPart) SetBoundary(boundary core.QByteArray_ITF) {
	if ptr.Pointer() != nil {
		C.QHttpMultiPart_SetBoundary(ptr.Pointer(), core.PointerFromQByteArray(boundary))
	}
}

func (ptr *QHttpMultiPart) SetContentType(contentType QHttpMultiPart__ContentType) {
	if ptr.Pointer() != nil {
		C.QHttpMultiPart_SetContentType(ptr.Pointer(), C.int(contentType))
	}
}

func (ptr *QHttpMultiPart) DestroyQHttpMultiPart() {
	if ptr.Pointer() != nil {
		C.QHttpMultiPart_DestroyQHttpMultiPart(ptr.Pointer())
		ptr.SetPointer(nil)
	}
}
