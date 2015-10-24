package location

//#include "qplaceresult.h"
import "C"
import (
	"unsafe"
)

type QPlaceResult struct {
	QPlaceSearchResult
}

type QPlaceResultITF interface {
	QPlaceSearchResultITF
	QPlaceResultPTR() *QPlaceResult
}

func PointerFromQPlaceResult(ptr QPlaceResultITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceResultPTR().Pointer()
	}
	return nil
}

func QPlaceResultFromPointer(ptr unsafe.Pointer) *QPlaceResult {
	var n = new(QPlaceResult)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPlaceResult) QPlaceResultPTR() *QPlaceResult {
	return ptr
}
