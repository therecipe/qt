package location

//#include "qplacecontentrequest.h"
import "C"
import (
	"unsafe"
)

type QPlaceContentRequest struct {
	ptr unsafe.Pointer
}

type QPlaceContentRequestITF interface {
	QPlaceContentRequestPTR() *QPlaceContentRequest
}

func (p *QPlaceContentRequest) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPlaceContentRequest) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPlaceContentRequest(ptr QPlaceContentRequestITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceContentRequestPTR().Pointer()
	}
	return nil
}

func QPlaceContentRequestFromPointer(ptr unsafe.Pointer) *QPlaceContentRequest {
	var n = new(QPlaceContentRequest)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPlaceContentRequest) QPlaceContentRequestPTR() *QPlaceContentRequest {
	return ptr
}
