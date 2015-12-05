package location

//#include "location.h"
import "C"
import (
	"unsafe"
)

type QPlaceSupplier struct {
	ptr unsafe.Pointer
}

type QPlaceSupplier_ITF interface {
	QPlaceSupplier_PTR() *QPlaceSupplier
}

func (p *QPlaceSupplier) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPlaceSupplier) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPlaceSupplier(ptr QPlaceSupplier_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceSupplier_PTR().Pointer()
	}
	return nil
}

func NewQPlaceSupplierFromPointer(ptr unsafe.Pointer) *QPlaceSupplier {
	var n = new(QPlaceSupplier)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPlaceSupplier) QPlaceSupplier_PTR() *QPlaceSupplier {
	return ptr
}
