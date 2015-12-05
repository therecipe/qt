package location

//#include "location.h"
import "C"
import (
	"unsafe"
)

type QPlaceResult struct {
	QPlaceSearchResult
}

type QPlaceResult_ITF interface {
	QPlaceSearchResult_ITF
	QPlaceResult_PTR() *QPlaceResult
}

func PointerFromQPlaceResult(ptr QPlaceResult_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceResult_PTR().Pointer()
	}
	return nil
}

func NewQPlaceResultFromPointer(ptr unsafe.Pointer) *QPlaceResult {
	var n = new(QPlaceResult)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPlaceResult) QPlaceResult_PTR() *QPlaceResult {
	return ptr
}
