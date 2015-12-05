package core

//#include "core.h"
import "C"
import (
	"log"
	"unsafe"
)

type QAssociativeIterable struct {
	ptr unsafe.Pointer
}

type QAssociativeIterable_ITF interface {
	QAssociativeIterable_PTR() *QAssociativeIterable
}

func (p *QAssociativeIterable) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAssociativeIterable) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQAssociativeIterable(ptr QAssociativeIterable_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAssociativeIterable_PTR().Pointer()
	}
	return nil
}

func NewQAssociativeIterableFromPointer(ptr unsafe.Pointer) *QAssociativeIterable {
	var n = new(QAssociativeIterable)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAssociativeIterable) QAssociativeIterable_PTR() *QAssociativeIterable {
	return ptr
}

func (ptr *QAssociativeIterable) Size() int {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAssociativeIterable::size")
		}
	}()

	if ptr.Pointer() != nil {
		return int(C.QAssociativeIterable_Size(ptr.Pointer()))
	}
	return 0
}

func (ptr *QAssociativeIterable) Value(key QVariant_ITF) *QVariant {
	defer func() {
		if recover() != nil {
			log.Println("recovered in QAssociativeIterable::value")
		}
	}()

	if ptr.Pointer() != nil {
		return NewQVariantFromPointer(C.QAssociativeIterable_Value(ptr.Pointer(), PointerFromQVariant(key)))
	}
	return nil
}
