package core

//#include "core.h"
import "C"
import (
	"github.com/therecipe/qt"
	"unsafe"
)

type QSequentialIterable struct {
	ptr unsafe.Pointer
}

type QSequentialIterable_ITF interface {
	QSequentialIterable_PTR() *QSequentialIterable
}

func (p *QSequentialIterable) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QSequentialIterable) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQSequentialIterable(ptr QSequentialIterable_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QSequentialIterable_PTR().Pointer()
	}
	return nil
}

func NewQSequentialIterableFromPointer(ptr unsafe.Pointer) *QSequentialIterable {
	var n = new(QSequentialIterable)
	n.SetPointer(ptr)
	return n
}

func (ptr *QSequentialIterable) QSequentialIterable_PTR() *QSequentialIterable {
	return ptr
}

func (ptr *QSequentialIterable) At(idx int) *QVariant {
	defer qt.Recovering("QSequentialIterable::at")

	if ptr.Pointer() != nil {
		return NewQVariantFromPointer(C.QSequentialIterable_At(ptr.Pointer(), C.int(idx)))
	}
	return nil
}

func (ptr *QSequentialIterable) CanReverseIterate() bool {
	defer qt.Recovering("QSequentialIterable::canReverseIterate")

	if ptr.Pointer() != nil {
		return C.QSequentialIterable_CanReverseIterate(ptr.Pointer()) != 0
	}
	return false
}

func (ptr *QSequentialIterable) Size() int {
	defer qt.Recovering("QSequentialIterable::size")

	if ptr.Pointer() != nil {
		return int(C.QSequentialIterable_Size(ptr.Pointer()))
	}
	return 0
}
