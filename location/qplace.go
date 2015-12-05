package location

//#include "location.h"
import "C"
import (
	"unsafe"
)

type QPlace struct {
	ptr unsafe.Pointer
}

type QPlace_ITF interface {
	QPlace_PTR() *QPlace
}

func (p *QPlace) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPlace) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPlace(ptr QPlace_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlace_PTR().Pointer()
	}
	return nil
}

func NewQPlaceFromPointer(ptr unsafe.Pointer) *QPlace {
	var n = new(QPlace)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPlace) QPlace_PTR() *QPlace {
	return ptr
}
