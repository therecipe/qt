package location

//#include "qplacesupplier.h"
import "C"
import (
	"unsafe"
)

type QPlaceSupplier struct {
	ptr unsafe.Pointer
}

type QPlaceSupplierITF interface {
	QPlaceSupplierPTR() *QPlaceSupplier
}

func (p *QPlaceSupplier) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QPlaceSupplier) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQPlaceSupplier(ptr QPlaceSupplierITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QPlaceSupplierPTR().Pointer()
	}
	return nil
}

func QPlaceSupplierFromPointer(ptr unsafe.Pointer) *QPlaceSupplier {
	var n = new(QPlaceSupplier)
	n.SetPointer(ptr)
	return n
}

func (ptr *QPlaceSupplier) QPlaceSupplierPTR() *QPlaceSupplier {
	return ptr
}
