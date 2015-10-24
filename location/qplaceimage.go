package location

//#include "qplaceimage.h"
import "C"
import (
	"unsafe"
)

type QPlaceImage struct {
	QPlaceContent
}

type QPlaceImageITF interface {
	QPlaceContentITF
	QPlaceImagePTR() *QPlaceImage
}

func PointerFromQPlaceImage(ptr QPlaceImageITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceImagePTR().Pointer()
	}
	return nil
}

func QPlaceImageFromPointer(ptr unsafe.Pointer) *QPlaceImage {
	var n = new(QPlaceImage)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPlaceImage) QPlaceImagePTR() *QPlaceImage {
	return ptr
}
