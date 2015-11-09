package core

//#include "qxmlstreamattribute.h"
import "C"
import (
	"unsafe"
)

type QXmlStreamAttribute struct {
	ptr unsafe.Pointer
}

type QXmlStreamAttribute_ITF interface {
	QXmlStreamAttribute_PTR() *QXmlStreamAttribute
}

func (p *QXmlStreamAttribute) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlStreamAttribute) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlStreamAttribute(ptr QXmlStreamAttribute_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlStreamAttribute_PTR().Pointer()
	}
	return nil
}

func NewQXmlStreamAttributeFromPointer(ptr unsafe.Pointer) *QXmlStreamAttribute {
	var n = new(QXmlStreamAttribute)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlStreamAttribute) QXmlStreamAttribute_PTR() *QXmlStreamAttribute {
	return ptr
}

func NewQXmlStreamAttribute() *QXmlStreamAttribute {
	return NewQXmlStreamAttributeFromPointer(C.QXmlStreamAttribute_NewQXmlStreamAttribute())
}

func NewQXmlStreamAttribute3(namespaceUri string, name string, value string) *QXmlStreamAttribute {
	return NewQXmlStreamAttributeFromPointer(C.QXmlStreamAttribute_NewQXmlStreamAttribute3(C.CString(namespaceUri), C.CString(name), C.CString(value)))
}

func NewQXmlStreamAttribute2(qualifiedName string, value string) *QXmlStreamAttribute {
	return NewQXmlStreamAttributeFromPointer(C.QXmlStreamAttribute_NewQXmlStreamAttribute2(C.CString(qualifiedName), C.CString(value)))
}

func NewQXmlStreamAttribute4(other QXmlStreamAttribute_ITF) *QXmlStreamAttribute {
	return NewQXmlStreamAttributeFromPointer(C.QXmlStreamAttribute_NewQXmlStreamAttribute4(PointerFromQXmlStreamAttribute(other)))
}

func (ptr *QXmlStreamAttribute) IsDefault() bool {
	if ptr.Pointer() != nil {
		return C.QXmlStreamAttribute_IsDefault(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlStreamAttribute) Name() *QStringRef {
	if ptr.Pointer() != nil {
		return NewQStringRefFromPointer(C.QXmlStreamAttribute_Name(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlStreamAttribute) NamespaceUri() *QStringRef {
	if ptr.Pointer() != nil {
		return NewQStringRefFromPointer(C.QXmlStreamAttribute_NamespaceUri(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlStreamAttribute) Prefix() *QStringRef {
	if ptr.Pointer() != nil {
		return NewQStringRefFromPointer(C.QXmlStreamAttribute_Prefix(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlStreamAttribute) QualifiedName() *QStringRef {
	if ptr.Pointer() != nil {
		return NewQStringRefFromPointer(C.QXmlStreamAttribute_QualifiedName(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlStreamAttribute) Value() *QStringRef {
	if ptr.Pointer() != nil {
		return NewQStringRefFromPointer(C.QXmlStreamAttribute_Value(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlStreamAttribute) DestroyQXmlStreamAttribute() {
	if ptr.Pointer() != nil {
		C.QXmlStreamAttribute_DestroyQXmlStreamAttribute(ptr.Pointer())
	}
}
