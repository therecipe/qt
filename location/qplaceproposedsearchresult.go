package location

//#include "qplaceproposedsearchresult.h"
import "C"
import (
	"unsafe"
)

type QPlaceProposedSearchResult struct {
	QPlaceSearchResult
}

type QPlaceProposedSearchResultITF interface {
	QPlaceSearchResultITF
	QPlaceProposedSearchResultPTR() *QPlaceProposedSearchResult
}

func PointerFromQPlaceProposedSearchResult(ptr QPlaceProposedSearchResultITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceProposedSearchResultPTR().Pointer()
	}
	return nil
}

func QPlaceProposedSearchResultFromPointer(ptr unsafe.Pointer) *QPlaceProposedSearchResult {
	var n = new(QPlaceProposedSearchResult)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPlaceProposedSearchResult) QPlaceProposedSearchResultPTR() *QPlaceProposedSearchResult {
	return ptr
}
