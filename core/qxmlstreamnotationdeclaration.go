package core

//#include "qxmlstreamnotationdeclaration.h"
import "C"
import (
	"unsafe"
)

type QXmlStreamNotationDeclaration struct {
	ptr unsafe.Pointer
}

type QXmlStreamNotationDeclarationITF interface {
	QXmlStreamNotationDeclarationPTR() *QXmlStreamNotationDeclaration
}

func (p *QXmlStreamNotationDeclaration) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlStreamNotationDeclaration) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlStreamNotationDeclaration(ptr QXmlStreamNotationDeclarationITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlStreamNotationDeclarationPTR().Pointer()
	}
	return nil
}

func QXmlStreamNotationDeclarationFromPointer(ptr unsafe.Pointer) *QXmlStreamNotationDeclaration {
	var n = new(QXmlStreamNotationDeclaration)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlStreamNotationDeclaration) QXmlStreamNotationDeclarationPTR() *QXmlStreamNotationDeclaration {
	return ptr
}

func NewQXmlStreamNotationDeclaration() *QXmlStreamNotationDeclaration {
	return QXmlStreamNotationDeclarationFromPointer(unsafe.Pointer(C.QXmlStreamNotationDeclaration_NewQXmlStreamNotationDeclaration()))
}

func NewQXmlStreamNotationDeclaration2(other QXmlStreamNotationDeclarationITF) *QXmlStreamNotationDeclaration {
	return QXmlStreamNotationDeclarationFromPointer(unsafe.Pointer(C.QXmlStreamNotationDeclaration_NewQXmlStreamNotationDeclaration2(C.QtObjectPtr(PointerFromQXmlStreamNotationDeclaration(other)))))
}

func (ptr *QXmlStreamNotationDeclaration) DestroyQXmlStreamNotationDeclaration() {
	if ptr.Pointer() != nil {
		C.QXmlStreamNotationDeclaration_DestroyQXmlStreamNotationDeclaration(C.QtObjectPtr(ptr.Pointer()))
	}
}
