package core

//#include "qxmlstreamnamespacedeclaration.h"
import "C"
import (
	"unsafe"
)

type QXmlStreamNamespaceDeclaration struct {
	ptr unsafe.Pointer
}

type QXmlStreamNamespaceDeclaration_ITF interface {
	QXmlStreamNamespaceDeclaration_PTR() *QXmlStreamNamespaceDeclaration
}

func (p *QXmlStreamNamespaceDeclaration) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlStreamNamespaceDeclaration) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlStreamNamespaceDeclaration(ptr QXmlStreamNamespaceDeclaration_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlStreamNamespaceDeclaration_PTR().Pointer()
	}
	return nil
}

func NewQXmlStreamNamespaceDeclarationFromPointer(ptr unsafe.Pointer) *QXmlStreamNamespaceDeclaration {
	var n = new(QXmlStreamNamespaceDeclaration)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlStreamNamespaceDeclaration) QXmlStreamNamespaceDeclaration_PTR() *QXmlStreamNamespaceDeclaration {
	return ptr
}

func NewQXmlStreamNamespaceDeclaration() *QXmlStreamNamespaceDeclaration {
	return NewQXmlStreamNamespaceDeclarationFromPointer(C.QXmlStreamNamespaceDeclaration_NewQXmlStreamNamespaceDeclaration())
}

func NewQXmlStreamNamespaceDeclaration3(prefix string, namespaceUri string) *QXmlStreamNamespaceDeclaration {
	return NewQXmlStreamNamespaceDeclarationFromPointer(C.QXmlStreamNamespaceDeclaration_NewQXmlStreamNamespaceDeclaration3(C.CString(prefix), C.CString(namespaceUri)))
}

func NewQXmlStreamNamespaceDeclaration2(other QXmlStreamNamespaceDeclaration_ITF) *QXmlStreamNamespaceDeclaration {
	return NewQXmlStreamNamespaceDeclarationFromPointer(C.QXmlStreamNamespaceDeclaration_NewQXmlStreamNamespaceDeclaration2(PointerFromQXmlStreamNamespaceDeclaration(other)))
}

func (ptr *QXmlStreamNamespaceDeclaration) NamespaceUri() *QStringRef {
	if ptr.Pointer() != nil {
		return NewQStringRefFromPointer(C.QXmlStreamNamespaceDeclaration_NamespaceUri(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlStreamNamespaceDeclaration) Prefix() *QStringRef {
	if ptr.Pointer() != nil {
		return NewQStringRefFromPointer(C.QXmlStreamNamespaceDeclaration_Prefix(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlStreamNamespaceDeclaration) DestroyQXmlStreamNamespaceDeclaration() {
	if ptr.Pointer() != nil {
		C.QXmlStreamNamespaceDeclaration_DestroyQXmlStreamNamespaceDeclaration(ptr.Pointer())
	}
}
