package core

//#include "core.h"
import "C"
import (
	"log"
	"unsafe"
)

type QXmlStreamEntityResolver struct {
	ptr unsafe.Pointer
}

type QXmlStreamEntityResolver_ITF interface {
	QXmlStreamEntityResolver_PTR() *QXmlStreamEntityResolver
}

func (p *QXmlStreamEntityResolver) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlStreamEntityResolver) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlStreamEntityResolver(ptr QXmlStreamEntityResolver_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlStreamEntityResolver_PTR().Pointer()
	}
	return nil
}

func NewQXmlStreamEntityResolverFromPointer(ptr unsafe.Pointer) *QXmlStreamEntityResolver {
	var n = new(QXmlStreamEntityResolver)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlStreamEntityResolver) QXmlStreamEntityResolver_PTR() *QXmlStreamEntityResolver {
	return ptr
}

func (ptr *QXmlStreamEntityResolver) ResolveUndeclaredEntity(name string) string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlStreamEntityResolver::resolveUndeclaredEntity")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlStreamEntityResolver_ResolveUndeclaredEntity(ptr.Pointer(), C.CString(name)))
	}
	return ""
}

func (ptr *QXmlStreamEntityResolver) DestroyQXmlStreamEntityResolver() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlStreamEntityResolver::~QXmlStreamEntityResolver")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlStreamEntityResolver_DestroyQXmlStreamEntityResolver(ptr.Pointer())
	}
}
