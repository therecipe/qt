package core

//#include "qxmlstreamentityresolver.h"
import "C"
import (
	"unsafe"
)

type QXmlStreamEntityResolver struct {
	ptr unsafe.Pointer
}

type QXmlStreamEntityResolverITF interface {
	QXmlStreamEntityResolverPTR() *QXmlStreamEntityResolver
}

func (p *QXmlStreamEntityResolver) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlStreamEntityResolver) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlStreamEntityResolver(ptr QXmlStreamEntityResolverITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlStreamEntityResolverPTR().Pointer()
	}
	return nil
}

func QXmlStreamEntityResolverFromPointer(ptr unsafe.Pointer) *QXmlStreamEntityResolver {
	var n = new(QXmlStreamEntityResolver)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlStreamEntityResolver) QXmlStreamEntityResolverPTR() *QXmlStreamEntityResolver {
	return ptr
}

func (ptr *QXmlStreamEntityResolver) ResolveUndeclaredEntity(name string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlStreamEntityResolver_ResolveUndeclaredEntity(C.QtObjectPtr(ptr.Pointer()), C.CString(name)))
	}
	return ""
}

func (ptr *QXmlStreamEntityResolver) DestroyQXmlStreamEntityResolver() {
	if ptr.Pointer() != nil {
		C.QXmlStreamEntityResolver_DestroyQXmlStreamEntityResolver(C.QtObjectPtr(ptr.Pointer()))
	}
}
