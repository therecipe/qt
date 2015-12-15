package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
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
	for len(n.ObjectNameAbs()) < len("QXmlStreamEntityResolver_") {
		n.SetObjectNameAbs("QXmlStreamEntityResolver_" + qt.Identifier())
	}
	return n
}

func (ptr *QXmlStreamEntityResolver) QXmlStreamEntityResolver_PTR() *QXmlStreamEntityResolver {
	return ptr
}

func (ptr *QXmlStreamEntityResolver) ResolveUndeclaredEntity(name string) string {
	defer qt.Recovering("QXmlStreamEntityResolver::resolveUndeclaredEntity")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlStreamEntityResolver_ResolveUndeclaredEntity(ptr.Pointer(), C.CString(name)))
	}
	return ""
}

func (ptr *QXmlStreamEntityResolver) DestroyQXmlStreamEntityResolver() {
	defer qt.Recovering("QXmlStreamEntityResolver::~QXmlStreamEntityResolver")

	if ptr.Pointer() != nil {
		C.QXmlStreamEntityResolver_DestroyQXmlStreamEntityResolver(ptr.Pointer())
	}
}

func (ptr *QXmlStreamEntityResolver) ObjectNameAbs() string {
	defer qt.Recovering("QXmlStreamEntityResolver::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlStreamEntityResolver_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlStreamEntityResolver) SetObjectNameAbs(name string) {
	defer qt.Recovering("QXmlStreamEntityResolver::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QXmlStreamEntityResolver_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
