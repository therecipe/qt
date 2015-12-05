package xml

//#include "xml.h"
import "C"
import (
	"log"
	"unsafe"
)

type QXmlEntityResolver struct {
	ptr unsafe.Pointer
}

type QXmlEntityResolver_ITF interface {
	QXmlEntityResolver_PTR() *QXmlEntityResolver
}

func (p *QXmlEntityResolver) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlEntityResolver) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlEntityResolver(ptr QXmlEntityResolver_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlEntityResolver_PTR().Pointer()
	}
	return nil
}

func NewQXmlEntityResolverFromPointer(ptr unsafe.Pointer) *QXmlEntityResolver {
	var n = new(QXmlEntityResolver)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlEntityResolver) QXmlEntityResolver_PTR() *QXmlEntityResolver {
	return ptr
}

func (ptr *QXmlEntityResolver) ErrorString() string {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlEntityResolver::errorString")
		}
	}()

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlEntityResolver_ErrorString(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlEntityResolver) DestroyQXmlEntityResolver() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlEntityResolver::~QXmlEntityResolver")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlEntityResolver_DestroyQXmlEntityResolver(ptr.Pointer())
	}
}
