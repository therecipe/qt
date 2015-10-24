package location

//#include "qplacesearchresult.h"
import "C"
import (
	"unsafe"
)

type QPlaceSearchResult struct {
	ptr unsafe.Pointer
}

type QPlaceSearchResultITF interface {
	QPlaceSearchResultPTR() *QPlaceSearchResult
}

func (p *QPlaceSearchResult) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPlaceSearchResult) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPlaceSearchResult(ptr QPlaceSearchResultITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceSearchResultPTR().Pointer()
	}
	return nil
}

func QPlaceSearchResultFromPointer(ptr unsafe.Pointer) *QPlaceSearchResult {
	var n = new(QPlaceSearchResult)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPlaceSearchResult) QPlaceSearchResultPTR() *QPlaceSearchResult {
	return ptr
}

//QPlaceSearchResult::SearchResultType
type QPlaceSearchResult__SearchResultType int

var (
	QPlaceSearchResult__UnknownSearchResult  = QPlaceSearchResult__SearchResultType(0)
	QPlaceSearchResult__PlaceResult          = QPlaceSearchResult__SearchResultType(1)
	QPlaceSearchResult__ProposedSearchResult = QPlaceSearchResult__SearchResultType(2)
)
