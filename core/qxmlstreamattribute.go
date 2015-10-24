package core

//#include "qxmlstreamattribute.h"
import "C"
import (
	"unsafe"
)

type QXmlStreamAttribute struct {
	ptr unsafe.Pointer
}

type QXmlStreamAttributeITF interface {
	QXmlStreamAttributePTR() *QXmlStreamAttribute
}

func (p *QXmlStreamAttribute) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlStreamAttribute) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlStreamAttribute(ptr QXmlStreamAttributeITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlStreamAttributePTR().Pointer()
	}
	return nil
}

func QXmlStreamAttributeFromPointer(ptr unsafe.Pointer) *QXmlStreamAttribute {
	var n = new(QXmlStreamAttribute)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlStreamAttribute) QXmlStreamAttributePTR() *QXmlStreamAttribute {
	return ptr
}

func NewQXmlStreamAttribute() *QXmlStreamAttribute {
	return QXmlStreamAttributeFromPointer(unsafe.Pointer(C.QXmlStreamAttribute_NewQXmlStreamAttribute()))
}

func NewQXmlStreamAttribute3(namespaceUri string, name string, value string) *QXmlStreamAttribute {
	return QXmlStreamAttributeFromPointer(unsafe.Pointer(C.QXmlStreamAttribute_NewQXmlStreamAttribute3(C.CString(namespaceUri), C.CString(name), C.CString(value))))
}

func NewQXmlStreamAttribute2(qualifiedName string, value string) *QXmlStreamAttribute {
	return QXmlStreamAttributeFromPointer(unsafe.Pointer(C.QXmlStreamAttribute_NewQXmlStreamAttribute2(C.CString(qualifiedName), C.CString(value))))
}

func NewQXmlStreamAttribute4(other QXmlStreamAttributeITF) *QXmlStreamAttribute {
	return QXmlStreamAttributeFromPointer(unsafe.Pointer(C.QXmlStreamAttribute_NewQXmlStreamAttribute4(C.QtObjectPtr(PointerFromQXmlStreamAttribute(other)))))
}

func (ptr *QXmlStreamAttribute) IsDefault() bool {
	if ptr.Pointer() != nil {
		return C.QXmlStreamAttribute_IsDefault(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QXmlStreamAttribute) DestroyQXmlStreamAttribute() {
	if ptr.Pointer() != nil {
		C.QXmlStreamAttribute_DestroyQXmlStreamAttribute(C.QtObjectPtr(ptr.Pointer()))
	}
}
