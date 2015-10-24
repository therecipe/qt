package location

//#include "qplacecategory.h"
import "C"
import (
	"unsafe"
)

type QPlaceCategory struct {
	ptr unsafe.Pointer
}

type QPlaceCategoryITF interface {
	QPlaceCategoryPTR() *QPlaceCategory
}

func (p *QPlaceCategory) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPlaceCategory) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPlaceCategory(ptr QPlaceCategoryITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceCategoryPTR().Pointer()
	}
	return nil
}

func QPlaceCategoryFromPointer(ptr unsafe.Pointer) *QPlaceCategory {
	var n = new(QPlaceCategory)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPlaceCategory) QPlaceCategoryPTR() *QPlaceCategory {
	return ptr
}
