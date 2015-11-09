package location

//#include "qplaceimage.h"
import "C"
import (
	"unsafe"
)

type QPlaceImage struct {
	QPlaceContent
}

type QPlaceImage_ITF interface {
	QPlaceContent_ITF
	QPlaceImage_PTR() *QPlaceImage
}

func PointerFromQPlaceImage(ptr QPlaceImage_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceImage_PTR().Pointer()
	}
	return nil
}

func NewQPlaceImageFromPointer(ptr unsafe.Pointer) *QPlaceImage {
	var n = new(QPlaceImage)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPlaceImage) QPlaceImage_PTR() *QPlaceImage {
	return ptr
}
