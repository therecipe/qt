package location

//#include "location.h"
import "C"
import (
	"unsafe"
)

type QPlaceSearchResult struct {
	ptr unsafe.Pointer
}

type QPlaceSearchResult_ITF interface {
	QPlaceSearchResult_PTR() *QPlaceSearchResult
}

func (p *QPlaceSearchResult) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPlaceSearchResult) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPlaceSearchResult(ptr QPlaceSearchResult_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceSearchResult_PTR().Pointer()
	}
	return nil
}

func NewQPlaceSearchResultFromPointer(ptr unsafe.Pointer) *QPlaceSearchResult {
	var n = new(QPlaceSearchResult)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPlaceSearchResult) QPlaceSearchResult_PTR() *QPlaceSearchResult {
	return ptr
}

//QPlaceSearchResult::SearchResultType
type QPlaceSearchResult__SearchResultType int64

const (
	QPlaceSearchResult__UnknownSearchResult  = QPlaceSearchResult__SearchResultType(0)
	QPlaceSearchResult__PlaceResult          = QPlaceSearchResult__SearchResultType(1)
	QPlaceSearchResult__ProposedSearchResult = QPlaceSearchResult__SearchResultType(2)
)
