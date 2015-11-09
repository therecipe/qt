package location

//#include "qplaceattribute.h"
import "C"
import (
	"unsafe"
)

type QPlaceAttribute struct {
	ptr unsafe.Pointer
}

type QPlaceAttribute_ITF interface {
	QPlaceAttribute_PTR() *QPlaceAttribute
}

func (p *QPlaceAttribute) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPlaceAttribute) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPlaceAttribute(ptr QPlaceAttribute_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceAttribute_PTR().Pointer()
	}
	return nil
}

func NewQPlaceAttributeFromPointer(ptr unsafe.Pointer) *QPlaceAttribute {
	var n = new(QPlaceAttribute)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPlaceAttribute) QPlaceAttribute_PTR() *QPlaceAttribute {
	return ptr
}
