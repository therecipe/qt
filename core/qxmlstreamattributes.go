package core

//#include "qxmlstreamattributes.h"
import "C"
import (
	"unsafe"
)

type QXmlStreamAttributes struct {
	QVector
}

type QXmlStreamAttributes_ITF interface {
	QVector_ITF
	QXmlStreamAttributes_PTR() *QXmlStreamAttributes
}

func PointerFromQXmlStreamAttributes(ptr QXmlStreamAttributes_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlStreamAttributes_PTR().Pointer()
	}
	return nil
}

func NewQXmlStreamAttributesFromPointer(ptr unsafe.Pointer) *QXmlStreamAttributes {
	var n = new(QXmlStreamAttributes)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlStreamAttributes) QXmlStreamAttributes_PTR() *QXmlStreamAttributes {
	return ptr
}

func NewQXmlStreamAttributes() *QXmlStreamAttributes {
	return NewQXmlStreamAttributesFromPointer(C.QXmlStreamAttributes_NewQXmlStreamAttributes())
}

func (ptr *QXmlStreamAttributes) Append(namespaceUri string, name string, value string) {
	if ptr.Pointer() != nil {
		C.QXmlStreamAttributes_Append(ptr.Pointer(), C.CString(namespaceUri), C.CString(name), C.CString(value))
	}
}

func (ptr *QXmlStreamAttributes) Append2(qualifiedName string, value string) {
	if ptr.Pointer() != nil {
		C.QXmlStreamAttributes_Append2(ptr.Pointer(), C.CString(qualifiedName), C.CString(value))
	}
}

func (ptr *QXmlStreamAttributes) HasAttribute2(qualifiedName QLatin1String_ITF) bool {
	if ptr.Pointer() != nil {
		return C.QXmlStreamAttributes_HasAttribute2(ptr.Pointer(), PointerFromQLatin1String(qualifiedName)) != 0
	}
	return false
}

func (ptr *QXmlStreamAttributes) HasAttribute3(namespaceUri string, name string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlStreamAttributes_HasAttribute3(ptr.Pointer(), C.CString(namespaceUri), C.CString(name)) != 0
	}
	return false
}

func (ptr *QXmlStreamAttributes) HasAttribute(qualifiedName string) bool {
	if ptr.Pointer() != nil {
		return C.QXmlStreamAttributes_HasAttribute(ptr.Pointer(), C.CString(qualifiedName)) != 0
	}
	return false
}

func (ptr *QXmlStreamAttributes) Value3(namespaceUri QLatin1String_ITF, name QLatin1String_ITF) *QStringRef {
	if ptr.Pointer() != nil {
		return NewQStringRefFromPointer(C.QXmlStreamAttributes_Value3(ptr.Pointer(), PointerFromQLatin1String(namespaceUri), PointerFromQLatin1String(name)))
	}
	return nil
}

func (ptr *QXmlStreamAttributes) Value5(qualifiedName QLatin1String_ITF) *QStringRef {
	if ptr.Pointer() != nil {
		return NewQStringRefFromPointer(C.QXmlStreamAttributes_Value5(ptr.Pointer(), PointerFromQLatin1String(qualifiedName)))
	}
	return nil
}

func (ptr *QXmlStreamAttributes) Value2(namespaceUri string, name QLatin1String_ITF) *QStringRef {
	if ptr.Pointer() != nil {
		return NewQStringRefFromPointer(C.QXmlStreamAttributes_Value2(ptr.Pointer(), C.CString(namespaceUri), PointerFromQLatin1String(name)))
	}
	return nil
}

func (ptr *QXmlStreamAttributes) Value(namespaceUri string, name string) *QStringRef {
	if ptr.Pointer() != nil {
		return NewQStringRefFromPointer(C.QXmlStreamAttributes_Value(ptr.Pointer(), C.CString(namespaceUri), C.CString(name)))
	}
	return nil
}

func (ptr *QXmlStreamAttributes) Value4(qualifiedName string) *QStringRef {
	if ptr.Pointer() != nil {
		return NewQStringRefFromPointer(C.QXmlStreamAttributes_Value4(ptr.Pointer(), C.CString(qualifiedName)))
	}
	return nil
}
