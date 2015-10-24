package xml

//#include "qxmlentityresolver.h"
import "C"
import (
	"unsafe"
)

type QXmlEntityResolver struct {
	ptr unsafe.Pointer
}

type QXmlEntityResolverITF interface {
	QXmlEntityResolverPTR() *QXmlEntityResolver
}

func (p *QXmlEntityResolver) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlEntityResolver) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlEntityResolver(ptr QXmlEntityResolverITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlEntityResolverPTR().Pointer()
	}
	return nil
}

func QXmlEntityResolverFromPointer(ptr unsafe.Pointer) *QXmlEntityResolver {
	var n = new(QXmlEntityResolver)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlEntityResolver) QXmlEntityResolverPTR() *QXmlEntityResolver {
	return ptr
}

func (ptr *QXmlEntityResolver) ErrorString() string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlEntityResolver_ErrorString(C.QtObjectPtr(ptr.Pointer())))
	}
	return ""
}

func (ptr *QXmlEntityResolver) DestroyQXmlEntityResolver() {
	if ptr.Pointer() != nil {
		C.QXmlEntityResolver_DestroyQXmlEntityResolver(C.QtObjectPtr(ptr.Pointer()))
	}
}
