package location

//#include "location.h"
import "C"
import (
	"unsafe"
)

type QPlaceSearchRequest struct {
	ptr unsafe.Pointer
}

type QPlaceSearchRequest_ITF interface {
	QPlaceSearchRequest_PTR() *QPlaceSearchRequest
}

func (p *QPlaceSearchRequest) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPlaceSearchRequest) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPlaceSearchRequest(ptr QPlaceSearchRequest_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceSearchRequest_PTR().Pointer()
	}
	return nil
}

func NewQPlaceSearchRequestFromPointer(ptr unsafe.Pointer) *QPlaceSearchRequest {
	var n = new(QPlaceSearchRequest)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPlaceSearchRequest) QPlaceSearchRequest_PTR() *QPlaceSearchRequest {
	return ptr
}

//QPlaceSearchRequest::RelevanceHint
type QPlaceSearchRequest__RelevanceHint int64

const (
	QPlaceSearchRequest__UnspecifiedHint      = QPlaceSearchRequest__RelevanceHint(0)
	QPlaceSearchRequest__DistanceHint         = QPlaceSearchRequest__RelevanceHint(1)
	QPlaceSearchRequest__LexicalPlaceNameHint = QPlaceSearchRequest__RelevanceHint(2)
)
