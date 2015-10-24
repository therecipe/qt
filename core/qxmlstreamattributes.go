package core

//#include "qxmlstreamattributes.h"
import "C"
import (
	"unsafe"
)

type QXmlStreamAttributes struct {
	QVector
}

type QXmlStreamAttributesITF interface {
	QVectorITF
	QXmlStreamAttributesPTR() *QXmlStreamAttributes
}

func PointerFromQXmlStreamAttributes(ptr QXmlStreamAttributesITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlStreamAttributesPTR().Pointer()
	}
	return nil
}

func QXmlStreamAttributesFromPointer(ptr unsafe.Pointer) *QXmlStreamAttributes {
	var n = new(QXmlStreamAttributes)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlStreamAttributes) QXmlStreamAttributesPTR() *QXmlStreamAttributes {
	return ptr
}

func NewQXmlStreamAttributes() *QXmlStreamAttributes {
	return QXmlStreamAttributesFromPointer(unsafe.Pointer(C.QXmlStreamAttributes_NewQXmlStreamAttributes()))
}

func (ptr *QXmlStreamAttributes) Append(namespaceUri string, name string, value string) {
	if ptr.Pointer() != nil {
		C.QXmlStreamAttributes_Append(C.QtObjectPtr(ptr.Pointer()), C.CString(namespaceUri), C.CString(name), C.CString(value))
	}
}

func (ptr *QXmlStreamAttributes) Append2(qualifiedName string, value string) {
	if ptr.Pointer() != nil {
		C.QXmlStreamAttributes_Append2(C.QtObjectPtr(ptr.Pointer()), C.CString(qualifiedName), C.CString(value))
	}
}

func (ptr *QXmlStreamAttributes) HasAttribute2(qualifiedName QLatin1StringITF) bool {
	if ptr.Pointer() != nil {
		return C.QXmlStreamAttributes_HasAttribute2(C.QtObjectPtr(ptr.Pointer()), C.QtObjectPtr(PointerFromQLatin1String(qualifiedName))) != 0
	}
	return false
}

func (ptr *QXmlStreamAttributes) HasAttribute3(namespaceUri string, name string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlStreamAttributes_HasAttribute3(C.QtObjectPtr(ptr.Pointer()), C.CString(namespaceUri), C.CString(name)) != 0
	}
	return false
}

func (ptr *QXmlStreamAttributes) HasAttribute(qualifiedName string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlStreamAttributes_HasAttribute(C.QtObjectPtr(ptr.Pointer()), C.CString(qualifiedName)) != 0
	}
	return false
}
