package core

//#include "qsequentialiterable.h"
import "C"
import (
	"unsafe"
)

type QSequentialIterable struct {
	ptr unsafe.Pointer
}

type QSequentialIterableITF interface {
	QSequentialIterablePTR() *QSequentialIterable
}

func (p *QSequentialIterable) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSequentialIterable) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSequentialIterable(ptr QSequentialIterableITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSequentialIterablePTR().Pointer()
	}
	return nil
}

func QSequentialIterableFromPointer(ptr unsafe.Pointer) *QSequentialIterable {
	var n = new(QSequentialIterable)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSequentialIterable) QSequentialIterablePTR() *QSequentialIterable {
	return ptr
}

func (ptr *QSequentialIterable) At(idx int) string {
	if ptr.Pointer() != nil {
		return C.GoString(C.QSequentialIterable_At(C.QtObjectPtr(ptr.Pointer()), C.int(idx)))
	}
	return ""
}

func (ptr *QSequentialIterable) CanReverseIterate() bool {
	if ptr.Pointer() != nil {
		return C.QSequentialIterable_CanReverseIterate(C.QtObjectPtr(ptr.Pointer())) != 0
	}
	return false
}

func (ptr *QSequentialIterable) Size() int {
	if ptr.Pointer() != nil {
		return int(C.QSequentialIterable_Size(C.QtObjectPtr(ptr.Pointer())))
	}
	return 0
}
