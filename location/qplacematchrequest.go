package location

//#include "location.h"
import "C"
import (
	"unsafe"
)

type QPlaceMatchRequest struct {
	ptr unsafe.Pointer
}

type QPlaceMatchRequest_ITF interface {
	QPlaceMatchRequest_PTR() *QPlaceMatchRequest
}

func (p *QPlaceMatchRequest) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPlaceMatchRequest) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPlaceMatchRequest(ptr QPlaceMatchRequest_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceMatchRequest_PTR().Pointer()
	}
	return nil
}

func NewQPlaceMatchRequestFromPointer(ptr unsafe.Pointer) *QPlaceMatchRequest {
	var n = new(QPlaceMatchRequest)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPlaceMatchRequest) QPlaceMatchRequest_PTR() *QPlaceMatchRequest {
	return ptr
}
