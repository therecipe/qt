package location

//#include "location.h"
import "C"
import (
	"unsafe"
)

type QPlaceCategory struct {
	ptr unsafe.Pointer
}

type QPlaceCategory_ITF interface {
	QPlaceCategory_PTR() *QPlaceCategory
}

func (p *QPlaceCategory) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPlaceCategory) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPlaceCategory(ptr QPlaceCategory_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceCategory_PTR().Pointer()
	}
	return nil
}

func NewQPlaceCategoryFromPointer(ptr unsafe.Pointer) *QPlaceCategory {
	var n = new(QPlaceCategory)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPlaceCategory) QPlaceCategory_PTR() *QPlaceCategory {
	return ptr
}
