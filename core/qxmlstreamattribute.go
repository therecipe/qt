package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
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
	defer qt.Recovering("QXmlStreamAttribute::QXmlStreamAttribute")

	return NewQXmlStreamAttributeFromPointer(C.QXmlStreamAttribute_NewQXmlStreamAttribute())
}

func NewQXmlStreamAttribute3(namespaceUri string, name string, value string) *QXmlStreamAttribute {
	defer qt.Recovering("QXmlStreamAttribute::QXmlStreamAttribute")

	return NewQXmlStreamAttributeFromPointer(C.QXmlStreamAttribute_NewQXmlStreamAttribute3(C.CString(namespaceUri), C.CString(name), C.CString(value)))
}

func NewQXmlStreamAttribute2(qualifiedName string, value string) *QXmlStreamAttribute {
	defer qt.Recovering("QXmlStreamAttribute::QXmlStreamAttribute")

	return NewQXmlStreamAttributeFromPointer(C.QXmlStreamAttribute_NewQXmlStreamAttribute2(C.CString(qualifiedName), C.CString(value)))
}

func NewQXmlStreamAttribute4(other QXmlStreamAttribute_ITF) *QXmlStreamAttribute {
	defer qt.Recovering("QXmlStreamAttribute::QXmlStreamAttribute")

	return NewQXmlStreamAttributeFromPointer(C.QXmlStreamAttribute_NewQXmlStreamAttribute4(PointerFromQXmlStreamAttribute(other)))
}

func (ptr *QXmlStreamAttribute) IsDefault() bool {
	defer qt.Recovering("QXmlStreamAttribute::isDefault")

	if ptr.Pointer() != nil {
		return C.QXmlStreamAttribute_IsDefault(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlStreamAttribute) Name() *QStringRef {
	defer qt.Recovering("QXmlStreamAttribute::name")

	if ptr.Pointer() != nil {
		return NewQStringRefFromPointer(C.QXmlStreamAttribute_Name(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlStreamAttribute) NamespaceUri() *QStringRef {
	defer qt.Recovering("QXmlStreamAttribute::namespaceUri")

	if ptr.Pointer() != nil {
		return NewQStringRefFromPointer(C.QXmlStreamAttribute_NamespaceUri(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlStreamAttribute) Prefix() *QStringRef {
	defer qt.Recovering("QXmlStreamAttribute::prefix")

	if ptr.Pointer() != nil {
		return NewQStringRefFromPointer(C.QXmlStreamAttribute_Prefix(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlStreamAttribute) QualifiedName() *QStringRef {
	defer qt.Recovering("QXmlStreamAttribute::qualifiedName")

	if ptr.Pointer() != nil {
		return NewQStringRefFromPointer(C.QXmlStreamAttribute_QualifiedName(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlStreamAttribute) Value() *QStringRef {
	defer qt.Recovering("QXmlStreamAttribute::value")

	if ptr.Pointer() != nil {
		return NewQStringRefFromPointer(C.QXmlStreamAttribute_Value(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlStreamAttribute) DestroyQXmlStreamAttribute() {
	defer qt.Recovering("QXmlStreamAttribute::~QXmlStreamAttribute")

	if ptr.Pointer() != nil {
		C.QXmlStreamAttribute_DestroyQXmlStreamAttribute(ptr.Pointer())
	}
}
