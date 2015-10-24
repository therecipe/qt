package location

//#include "qplacesearchrequest.h"
import "C"
import (
	"unsafe"
)

type QPlaceSearchRequest struct {
	ptr unsafe.Pointer
}

type QPlaceSearchRequestITF interface {
	QPlaceSearchRequestPTR() *QPlaceSearchRequest
}

func (p *QPlaceSearchRequest) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPlaceSearchRequest) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPlaceSearchRequest(ptr QPlaceSearchRequestITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceSearchRequestPTR().Pointer()
	}
	return nil
}

func QPlaceSearchRequestFromPointer(ptr unsafe.Pointer) *QPlaceSearchRequest {
	var n = new(QPlaceSearchRequest)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPlaceSearchRequest) QPlaceSearchRequestPTR() *QPlaceSearchRequest {
	return ptr
}

//QPlaceSearchRequest::RelevanceHint
type QPlaceSearchRequest__RelevanceHint int

var (
	QPlaceSearchRequest__UnspecifiedHint      = QPlaceSearchRequest__RelevanceHint(0)
	QPlaceSearchRequest__DistanceHint         = QPlaceSearchRequest__RelevanceHint(1)
	QPlaceSearchRequest__LexicalPlaceNameHint = QPlaceSearchRequest__RelevanceHint(2)
)
