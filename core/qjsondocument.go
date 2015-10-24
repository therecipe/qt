package core

//#include "qjsondocument.h"
import "C"
import (
	"unsafe"
)

type QJsonDocument struct {
	ptr unsafe.Pointer
}

type QJsonDocumentITF interface {
	QJsonDocumentPTR() *QJsonDocument
}

func (p *QJsonDocument) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QJsonDocument) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQJsonDocument(ptr QJsonDocumentITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QJsonDocumentPTR().Pointer()
	}
	return nil
}

func QJsonDocumentFromPointer(ptr unsafe.Pointer) *QJsonDocument {
	var n = new(QJsonDocument)
	n.SetPointer(ptr)
	return n
}

func (ptr *QJsonDocument) QJsonDocumentPTR() *QJsonDocument {
	return ptr
}

//QJsonDocument::DataValidation
type QJsonDocument__DataValidation int

var (
	QJsonDocument__Validate         = QJsonDocument__DataValidation(0)
	QJsonDocument__BypassValidation = QJsonDocument__DataValidation(1)
)

//QJsonDocument::JsonFormat
type QJsonDocument__JsonFormat int

var (
	QJsonDocument__Indented = QJsonDocument__JsonFormat(0)
	QJsonDocument__Compact  = QJsonDocument__JsonFormat(1)
)

func NewQJsonDocument() *QJsonDocument {
	return QJsonDocumentFromPointer(unsafe.Pointer(C.QJsonDocument_NewQJsonDocument()))
}

func NewQJsonDocument3(array QJsonArrayITF) *QJsonDocument {
	return QJsonDocumentFromPointer(unsafe.Pointer(C.QJsonDocument_NewQJsonDocument3(C.QtObjectPtr(PointerFromQJsonArray(array)))))
}

func NewQJsonDocument4(other QJsonDocumentITF) *QJsonDocument {
	return QJsonDocumentFromPointer(unsafe.Pointer(C.QJsonDocument_NewQJsonDocument4(C.QtObjectPtr(PointerFromQJsonDocument(other)))))
}

func NewQJsonDocument2(object QJsonObjectITF) *QJsonDocument {
	return QJsonDocumentFromPointer(unsafe.Pointer(C.QJsonDocument_NewQJsonDocument2(C.QtObjectPtr(PointerFromQJsonObject(object)))))
}

func (ptr *QJsonDocument) IsArray() bool {
	if ptr.Pointer() != nil {
		return C.QJsonDocument_IsArray(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QJsonDocument) IsEmpty() bool {
	if ptr.Pointer() != nil {
		return C.QJsonDocument_IsEmpty(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QJsonDocument) IsNull() bool {
	if ptr.Pointer() != nil {
		return C.QJsonDocument_IsNull(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QJsonDocument) IsObject() bool {
	if ptr.Pointer() != nil {
		return C.QJsonDocument_IsObject(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QJsonDocument) SetArray(array QJsonArrayITF) {
	if ptr.Pointer() != nil {
		C.QJsonDocument_SetArray(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQJsonArray(array)))
	}
}

func (ptr *QJsonDocument) SetObject(object QJsonObjectITF) {
	if ptr.Pointer() != nil {
		C.QJsonDocument_SetObject(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQJsonObject(object)))
	}
}

func (ptr *QJsonDocument) ToVariant() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QJsonDocument_ToVariant(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QJsonDocument) DestroyQJsonDocument() {
	if ptr.Pointer() != nil {
		C.QJsonDocument_DestroyQJsonDocument(C.QtObjectPtr(ptr.Pointer()))
	}
}
