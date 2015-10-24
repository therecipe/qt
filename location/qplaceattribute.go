package location

//#include "qplaceattribute.h"
import "C"
import (
	"unsafe"
)

type QPlaceAttribute struct {
	ptr unsafe.Pointer
}

type QPlaceAttributeITF interface {
	QPlaceAttributePTR() *QPlaceAttribute
}

func (p *QPlaceAttribute) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPlaceAttribute) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPlaceAttribute(ptr QPlaceAttributeITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceAttributePTR().Pointer()
	}
	return nil
}

func QPlaceAttributeFromPointer(ptr unsafe.Pointer) *QPlaceAttribute {
	var n = new(QPlaceAttribute)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPlaceAttribute) QPlaceAttributePTR() *QPlaceAttribute {
	return ptr
}
