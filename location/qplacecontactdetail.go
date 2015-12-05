package location

//#include "location.h"
import "C"
import (
	"unsafe"
)

type QPlaceContactDetail struct {
	ptr unsafe.Pointer
}

type QPlaceContactDetail_ITF interface {
	QPlaceContactDetail_PTR() *QPlaceContactDetail
}

func (p *QPlaceContactDetail) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPlaceContactDetail) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPlaceContactDetail(ptr QPlaceContactDetail_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceContactDetail_PTR().Pointer()
	}
	return nil
}

func NewQPlaceContactDetailFromPointer(ptr unsafe.Pointer) *QPlaceContactDetail {
	var n = new(QPlaceContactDetail)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPlaceContactDetail) QPlaceContactDetail_PTR() *QPlaceContactDetail {
	return ptr
}
