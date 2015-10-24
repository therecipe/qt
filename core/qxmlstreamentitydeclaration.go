package core

//#include "qxmlstreamentitydeclaration.h"
import "C"
import (
	"unsafe"
)

type QXmlStreamEntityDeclaration struct {
	ptr unsafe.Pointer
}

type QXmlStreamEntityDeclarationITF interface {
	QXmlStreamEntityDeclarationPTR() *QXmlStreamEntityDeclaration
}

func (p *QXmlStreamEntityDeclaration) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlStreamEntityDeclaration) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlStreamEntityDeclaration(ptr QXmlStreamEntityDeclarationITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlStreamEntityDeclarationPTR().Pointer()
	}
	return nil
}

func QXmlStreamEntityDeclarationFromPointer(ptr unsafe.Pointer) *QXmlStreamEntityDeclaration {
	var n = new(QXmlStreamEntityDeclaration)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlStreamEntityDeclaration) QXmlStreamEntityDeclarationPTR() *QXmlStreamEntityDeclaration {
	return ptr
}

func NewQXmlStreamEntityDeclaration() *QXmlStreamEntityDeclaration {
	return QXmlStreamEntityDeclarationFromPointer(unsafe.Pointer(C.QXmlStreamEntityDeclaration_NewQXmlStreamEntityDeclaration()))
}

func NewQXmlStreamEntityDeclaration2(other QXmlStreamEntityDeclarationITF) *QXmlStreamEntityDeclaration {
	return QXmlStreamEntityDeclarationFromPointer(unsafe.Pointer(C.QXmlStreamEntityDeclaration_NewQXmlStreamEntityDeclaration2(C.QtObjectPtr(PointerFromQXmlStreamEntityDeclaration(other)))))
}

func (ptr *QXmlStreamEntityDeclaration) DestroyQXmlStreamEntityDeclaration() {
	if ptr.Pointer() != nil {
		C.QXmlStreamEntityDeclaration_DestroyQXmlStreamEntityDeclaration(C.QtObjectPtr(ptr.Pointer()))
	}
}
