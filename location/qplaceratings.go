package location

//#include "qplaceratings.h"
import "C"
import (
	"unsafe"
)

type QPlaceRatings struct {
	ptr unsafe.Pointer
}

type QPlaceRatingsITF interface {
	QPlaceRatingsPTR() *QPlaceRatings
}

func (p *QPlaceRatings) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPlaceRatings) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPlaceRatings(ptr QPlaceRatingsITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceRatingsPTR().Pointer()
	}
	return nil
}

func QPlaceRatingsFromPointer(ptr unsafe.Pointer) *QPlaceRatings {
	var n = new(QPlaceRatings)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPlaceRatings) QPlaceRatingsPTR() *QPlaceRatings {
	return ptr
}
