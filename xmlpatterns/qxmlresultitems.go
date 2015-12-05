package xmlpatterns

//#include "xmlpatterns.h"
import "C"
import (
	"log"
	"unsafe"
)

type QXmlResultItems struct {
	ptr unsafe.Pointer
}

type QXmlResultItems_ITF interface {
	QXmlResultItems_PTR() *QXmlResultItems
}

func (p *QXmlResultItems) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QXmlResultItems) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQXmlResultItems(ptr QXmlResultItems_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QXmlResultItems_PTR().Pointer()
	}
	return nil
}

func NewQXmlResultItemsFromPointer(ptr unsafe.Pointer) *QXmlResultItems {
	var n = new(QXmlResultItems)
	n.SetPointer(ptr)
	return n
}

func (ptr *QXmlResultItems) QXmlResultItems_PTR() *QXmlResultItems {
	return ptr
}

func NewQXmlResultItems() *QXmlResultItems {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlResultItems::QXmlResultItems")
		}
	}()

	return NewQXmlResultItemsFromPointer(C.QXmlResultItems_NewQXmlResultItems())
}

func (ptr *QXmlResultItems) HasError() bool {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlResultItems::hasError")
		}
	}()

	if ptr.Pointer() != nil {
		return C.QXmlResultItems_HasError(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlResultItems) DestroyQXmlResultItems() {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QXmlResultItems::~QXmlResultItems")
		}
	}()

	if ptr.Pointer() != nil {
		C.QXmlResultItems_DestroyQXmlResultItems(ptr.Pointer())
	}
}
