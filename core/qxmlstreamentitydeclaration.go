package core

//#include "qxmlstreamentitydeclaration.h"
import "C"
import (
	"unsafe"
)

type QXmlStreamEntityDeclaration struct {
	ptr unsafe.Pointer
}

type QXmlStreamEntityDeclaration_ITF interface {
	QXmlStreamEntityDeclaration_PTR() *QXmlStreamEntityDeclaration
}

func (p *QXmlStreamEntityDeclaration) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlStreamEntityDeclaration) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlStreamEntityDeclaration(ptr QXmlStreamEntityDeclaration_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlStreamEntityDeclaration_PTR().Pointer()
	}
	return nil
}

func NewQXmlStreamEntityDeclarationFromPointer(ptr unsafe.Pointer) *QXmlStreamEntityDeclaration {
	var n = new(QXmlStreamEntityDeclaration)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlStreamEntityDeclaration) QXmlStreamEntityDeclaration_PTR() *QXmlStreamEntityDeclaration {
	return ptr
}

func NewQXmlStreamEntityDeclaration() *QXmlStreamEntityDeclaration {
	return NewQXmlStreamEntityDeclarationFromPointer(C.QXmlStreamEntityDeclaration_NewQXmlStreamEntityDeclaration())
}

func NewQXmlStreamEntityDeclaration2(other QXmlStreamEntityDeclaration_ITF) *QXmlStreamEntityDeclaration {
	return NewQXmlStreamEntityDeclarationFromPointer(C.QXmlStreamEntityDeclaration_NewQXmlStreamEntityDeclaration2(PointerFromQXmlStreamEntityDeclaration(other)))
}

func (ptr *QXmlStreamEntityDeclaration) Name() *QStringRef {
	if ptr.Pointer() != nil {
		return NewQStringRefFromPointer(C.QXmlStreamEntityDeclaration_Name(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlStreamEntityDeclaration) NotationName() *QStringRef {
	if ptr.Pointer() != nil {
		return NewQStringRefFromPointer(C.QXmlStreamEntityDeclaration_NotationName(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlStreamEntityDeclaration) PublicId() *QStringRef {
	if ptr.Pointer() != nil {
		return NewQStringRefFromPointer(C.QXmlStreamEntityDeclaration_PublicId(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlStreamEntityDeclaration) SystemId() *QStringRef {
	if ptr.Pointer() != nil {
		return NewQStringRefFromPointer(C.QXmlStreamEntityDeclaration_SystemId(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlStreamEntityDeclaration) Value() *QStringRef {
	if ptr.Pointer() != nil {
		return NewQStringRefFromPointer(C.QXmlStreamEntityDeclaration_Value(ptr.Pointer()))
	}
	return nil
}

func (ptr *QXmlStreamEntityDeclaration) DestroyQXmlStreamEntityDeclaration() {
	if ptr.Pointer() != nil {
		C.QXmlStreamEntityDeclaration_DestroyQXmlStreamEntityDeclaration(ptr.Pointer())
	}
}
