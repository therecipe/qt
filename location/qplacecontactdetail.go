package location

//#include "qplacecontactdetail.h"
import "C"
import (
	"unsafe"
)

type QPlaceContactDetail struct {
	ptr unsafe.Pointer
}

type QPlaceContactDetailITF interface {
	QPlaceContactDetailPTR() *QPlaceContactDetail
}

func (p *QPlaceContactDetail) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPlaceContactDetail) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPlaceContactDetail(ptr QPlaceContactDetailITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceContactDetailPTR().Pointer()
	}
	return nil
}

func QPlaceContactDetailFromPointer(ptr unsafe.Pointer) *QPlaceContactDetail {
	var n = new(QPlaceContactDetail)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPlaceContactDetail) QPlaceContactDetailPTR() *QPlaceContactDetail {
	return ptr
}
