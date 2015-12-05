package location

//#include "location.h"
import "C"
import (
	"unsafe"
)

type QPlaceContentRequest struct {
	ptr unsafe.Pointer
}

type QPlaceContentRequest_ITF interface {
	QPlaceContentRequest_PTR() *QPlaceContentRequest
}

func (p *QPlaceContentRequest) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPlaceContentRequest) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPlaceContentRequest(ptr QPlaceContentRequest_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceContentRequest_PTR().Pointer()
	}
	return nil
}

func NewQPlaceContentRequestFromPointer(ptr unsafe.Pointer) *QPlaceContentRequest {
	var n = new(QPlaceContentRequest)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPlaceContentRequest) QPlaceContentRequest_PTR() *QPlaceContentRequest {
	return ptr
}
