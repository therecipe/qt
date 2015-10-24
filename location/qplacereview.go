package location

//#include "qplacereview.h"
import "C"
import (
	"unsafe"
)

type QPlaceReview struct {
	QPlaceContent
}

type QPlaceReviewITF interface {
	QPlaceContentITF
	QPlaceReviewPTR() *QPlaceReview
}

func PointerFromQPlaceReview(ptr QPlaceReviewITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceReviewPTR().Pointer()
	}
	return nil
}

func QPlaceReviewFromPointer(ptr unsafe.Pointer) *QPlaceReview {
	var n = new(QPlaceReview)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPlaceReview) QPlaceReviewPTR() *QPlaceReview {
	return ptr
}
