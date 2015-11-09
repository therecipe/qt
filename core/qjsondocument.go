package core

//#include "qjsondocument.h"
import "C"
import (
	"unsafe"
)

type QJsonDocument struct {
	ptr unsafe.Pointer
}

type QJsonDocument_ITF interface {
	QJsonDocument_PTR() *QJsonDocument
}

func (p *QJsonDocument) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QJsonDocument) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQJsonDocument(ptr QJsonDocument_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QJsonDocument_PTR().Pointer()
	}
	return nil
}

func NewQJsonDocumentFromPointer(ptr unsafe.Pointer) *QJsonDocument {
	var n = new(QJsonDocument)
	n.SetPointer(ptr)
	return n
}

func (ptr *QJsonDocument) QJsonDocument_PTR() *QJsonDocument {
	return ptr
}

//QJsonDocument::DataValidation
type QJsonDocument__DataValidation int64

const (
	QJsonDocument__Validate         = QJsonDocument__DataValidation(0)
	QJsonDocument__BypassValidation = QJsonDocument__DataValidation(1)
)

//QJsonDocument::JsonFormat
type QJsonDocument__JsonFormat int64

const (
	QJsonDocument__Indented = QJsonDocument__JsonFormat(0)
	QJsonDocument__Compact  = QJsonDocument__JsonFormat(1)
)

func NewQJsonDocument() *QJsonDocument {
	return NewQJsonDocumentFromPointer(C.QJsonDocument_NewQJsonDocument())
}

func NewQJsonDocument3(array QJsonArray_ITF) *QJsonDocument {
	return NewQJsonDocumentFromPointer(C.QJsonDocument_NewQJsonDocument3(PointerFromQJsonArray(array)))
}

func NewQJsonDocument4(other QJsonDocument_ITF) *QJsonDocument {
	return NewQJsonDocumentFromPointer(C.QJsonDocument_NewQJsonDocument4(PointerFromQJsonDocument(other)))
}

func NewQJsonDocument2(object QJsonObject_ITF) *QJsonDocument {
	return NewQJsonDocumentFromPointer(C.QJsonDocument_NewQJsonDocument2(PointerFromQJsonObject(object)))
}

func (ptr *QJsonDocument) Array() *QJsonArray {
	if ptr.Pointer() != nil {
		return NewQJsonArrayFromPointer(C.QJsonDocument_Array(ptr.Pointer()))
	}
	return nil
}

func QJsonDocument_FromBinaryData(data QByteArray_ITF, validation QJsonDocument__DataValidation) *QJsonDocument {
	return NewQJsonDocumentFromPointer(C.QJsonDocument_QJsonDocument_FromBinaryData(PointerFromQByteArray(data), C.int(validation)))
}

func QJsonDocument_FromJson(json QByteArray_ITF, error QJsonParseError_ITF) *QJsonDocument {
	return NewQJsonDocumentFromPointer(C.QJsonDocument_QJsonDocument_FromJson(PointerFromQByteArray(json), PointerFromQJsonParseError(error)))
}

func QJsonDocument_FromRawData(data string, size int, validation QJsonDocument__DataValidation) *QJsonDocument {
	return NewQJsonDocumentFromPointer(C.QJsonDocument_QJsonDocument_FromRawData(C.CString(data), C.int(size), C.int(validation)))
}

func QJsonDocument_FromVariant(variant QVariant_ITF) *QJsonDocument {
	return NewQJsonDocumentFromPointer(C.QJsonDocument_QJsonDocument_FromVariant(PointerFromQVariant(variant)))
}

func (ptr *QJsonDocument) IsArray() bool {
	if ptr.Pointer() != nil {
		return C.QJsonDocument_IsArray(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJsonDocument) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QJsonDocument_IsEmpty(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJsonDocument) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QJsonDocument_IsNull(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJsonDocument) IsObject() bool {
	if ptr.Pointer() != nil {
		return C.QJsonDocument_IsObject(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QJsonDocument) Object() *QJsonObject {
	if ptr.Pointer() != nil {
		return NewQJsonObjectFromPointer(C.QJsonDocument_Object(ptr.Pointer()))
	}
	return nil
}

func (ptr *QJsonDocument) SetArray(array QJsonArray_ITF) {
	if ptr.Pointer() != nil {
		C.QJsonDocument_SetArray(ptr.Pointer(), PointerFromQJsonArray(array))
	}
}

func (ptr *QJsonDocument) SetObject(object QJsonObject_ITF) {
	if ptr.Pointer() != nil {
		C.QJsonDocument_SetObject(ptr.Pointer(), PointerFromQJsonObject(object))
	}
}

func (ptr *QJsonDocument) ToBinaryData() *QByteArray {
	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QJsonDocument_ToBinaryData(ptr.Pointer()))
	}
	return nil
}

func (ptr *QJsonDocument) ToJson(format QJsonDocument__JsonFormat) *QByteArray {
	if ptr.Pointer() != nil {
		return NewQByteArrayFromPointer(C.QJsonDocument_ToJson(ptr.Pointer(), C.int(format)))
	}
	return nil
}

func (ptr *QJsonDocument) ToVariant() *QVariant {
	if ptr.Pointer() != nil {
		return NewQVariantFromPointer(C.QJsonDocument_ToVariant(ptr.Pointer()))
	}
	return nil
}

func (ptr *QJsonDocument) DestroyQJsonDocument() {
	if ptr.Pointer() != nil {
		C.QJsonDocument_DestroyQJsonDocument(ptr.Pointer())
	}
}
