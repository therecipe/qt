package location

//#include "qplace.h"
import "C"
import (
	"unsafe"
)

type QPlace struct {
	ptr unsafe.Pointer
}

type QPlaceITF interface {
	QPlacePTR() *QPlace
}

func (p *QPlace) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPlace) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPlace(ptr QPlaceITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlacePTR().Pointer()
	}
	return nil
}

func QPlaceFromPointer(ptr unsafe.Pointer) *QPlace {
	var n = new(QPlace)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPlace) QPlacePTR() *QPlace {
	return ptr
}
