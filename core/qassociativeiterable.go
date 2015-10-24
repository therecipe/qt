package core

//#include "qassociativeiterable.h"
import "C"
import (
	"unsafe"
)

type QAssociativeIterable struct {
	ptr unsafe.Pointer
}

type QAssociativeIterableITF interface {
	QAssociativeIterablePTR() *QAssociativeIterable
}

func (p *QAssociativeIterable) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAssociativeIterable) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQAssociativeIterable(ptr QAssociativeIterableITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAssociativeIterablePTR().Pointer()
	}
	return nil
}

func QAssociativeIterableFromPointer(ptr unsafe.Pointer) *QAssociativeIterable {
	var n = new(QAssociativeIterable)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAssociativeIterable) QAssociativeIterablePTR() *QAssociativeIterable {
	return ptr
}

func (ptr *QAssociativeIterable) Size() int {
	if ptr.Pointer() != nil {
		return int(C.QAssociativeIterable_Size(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}

func (ptr *QAssociativeIterable) Value(key string) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QAssociativeIterable_Value(C.QtObjectPtr(ptr.Pointer()), C.CString(key)))
	}
	return ""
}
