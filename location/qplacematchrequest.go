package location

//#include "qplacematchrequest.h"
import "C"
import (
	"unsafe"
)

type QPlaceMatchRequest struct {
	ptr unsafe.Pointer
}

type QPlaceMatchRequestITF interface {
	QPlaceMatchRequestPTR() *QPlaceMatchRequest
}

func (p *QPlaceMatchRequest) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPlaceMatchRequest) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPlaceMatchRequest(ptr QPlaceMatchRequestITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceMatchRequestPTR().Pointer()
	}
	return nil
}

func QPlaceMatchRequestFromPointer(ptr unsafe.Pointer) *QPlaceMatchRequest {
	var n = new(QPlaceMatchRequest)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPlaceMatchRequest) QPlaceMatchRequestPTR() *QPlaceMatchRequest {
	return ptr
}
