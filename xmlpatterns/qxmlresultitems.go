package xmlpatterns

//#include "xmlpatterns.h"
import "C"
import (
	"github.com/therecipe/qt"
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
	for len(n.ObjectNameAbs()) < len("QXmlResultItems_") {
		n.SetObjectNameAbs("QXmlResultItems_" + qt.Identifier())
	}
	return n
}

func (ptr *QXmlResultItems) QXmlResultItems_PTR() *QXmlResultItems {
	return ptr
}

func NewQXmlResultItems() *QXmlResultItems {
	defer qt.Recovering("QXmlResultItems::QXmlResultItems")

	return NewQXmlResultItemsFromPointer(C.QXmlResultItems_NewQXmlResultItems())
}

func (ptr *QXmlResultItems) HasError() bool {
	defer qt.Recovering("QXmlResultItems::hasError")

	if ptr.Pointer() != nil {
		return C.QXmlResultItems_HasError(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QXmlResultItems) DestroyQXmlResultItems() {
	defer qt.Recovering("QXmlResultItems::~QXmlResultItems")

	if ptr.Pointer() != nil {
		C.QXmlResultItems_DestroyQXmlResultItems(ptr.Pointer())
	}
}

func (ptr *QXmlResultItems) ObjectNameAbs() string {
	defer qt.Recovering("QXmlResultItems::objectNameAbs")

	if ptr.Pointer() != nil {
		return C.GoString(C.QXmlResultItems_ObjectNameAbs(ptr.Pointer()))
	}
	return ""
}

func (ptr *QXmlResultItems) SetObjectNameAbs(name string) {
	defer qt.Recovering("QXmlResultItems::setObjectNameAbs")

	if ptr.Pointer() != nil {
		C.QXmlResultItems_SetObjectNameAbs(ptr.Pointer(), C.CString(name))
	}
}
