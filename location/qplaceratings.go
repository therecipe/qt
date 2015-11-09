package location

//#include "qplaceratings.h"
import "C"
import (
	"unsafe"
)

type QPlaceRatings struct {
	ptr unsafe.Pointer
}

type QPlaceRatings_ITF interface {
	QPlaceRatings_PTR() *QPlaceRatings
}

func (p *QPlaceRatings) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPlaceRatings) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPlaceRatings(ptr QPlaceRatings_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceRatings_PTR().Pointer()
	}
	return nil
}

func NewQPlaceRatingsFromPointer(ptr unsafe.Pointer) *QPlaceRatings {
	var n = new(QPlaceRatings)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPlaceRatings) QPlaceRatings_PTR() *QPlaceRatings {
	return ptr
}
