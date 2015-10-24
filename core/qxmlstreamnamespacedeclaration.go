package core

//#include "qxmlstreamnamespacedeclaration.h"
import "C"
import (
	"unsafe"
)

type QXmlStreamNamespaceDeclaration struct {
	ptr unsafe.Pointer
}

type QXmlStreamNamespaceDeclarationITF interface {
	QXmlStreamNamespaceDeclarationPTR() *QXmlStreamNamespaceDeclaration
}

func (p *QXmlStreamNamespaceDeclaration) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlStreamNamespaceDeclaration) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlStreamNamespaceDeclaration(ptr QXmlStreamNamespaceDeclarationITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlStreamNamespaceDeclarationPTR().Pointer()
	}
	return nil
}

func QXmlStreamNamespaceDeclarationFromPointer(ptr unsafe.Pointer) *QXmlStreamNamespaceDeclaration {
	var n = new(QXmlStreamNamespaceDeclaration)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlStreamNamespaceDeclaration) QXmlStreamNamespaceDeclarationPTR() *QXmlStreamNamespaceDeclaration {
	return ptr
}

func NewQXmlStreamNamespaceDeclaration() *QXmlStreamNamespaceDeclaration {
	return QXmlStreamNamespaceDeclarationFromPointer(unsafe.Pointer(C.QXmlStreamNamespaceDeclaration_NewQXmlStreamNamespaceDeclaration()))
}

func NewQXmlStreamNamespaceDeclaration3(prefix string, namespaceUri string) *QXmlStreamNamespaceDeclaration {
	return QXmlStreamNamespaceDeclarationFromPointer(unsafe.Pointer(C.QXmlStreamNamespaceDeclaration_NewQXmlStreamNamespaceDeclaration3(C.CString(prefix), C.CString(namespaceUri))))
}

func NewQXmlStreamNamespaceDeclaration2(other QXmlStreamNamespaceDeclarationITF) *QXmlStreamNamespaceDeclaration {
	return QXmlStreamNamespaceDeclarationFromPointer(unsafe.Pointer(C.QXmlStreamNamespaceDeclaration_NewQXmlStreamNamespaceDeclaration2(C.QtObjectPtr(PointerFromQXmlStreamNamespaceDeclaration(other)))))
}

func (ptr *QXmlStreamNamespaceDeclaration) DestroyQXmlStreamNamespaceDeclaration() {
	if ptr.Pointer() != nil {
		C.QXmlStreamNamespaceDeclaration_DestroyQXmlStreamNamespaceDeclaration(C.QtObjectPtr(ptr.Pointer()))
	}
}
