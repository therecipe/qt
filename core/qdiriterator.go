package core

//#include "qdiriterator.h"
import "C"
import (
	"unsafe"
)

type QDirIterator struct {
	ptr unsafe.Pointer
}

type QDirIterator_ITF interface {
	QDirIterator_PTR() *QDirIterator
}

func (p *QDirIterator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDirIterator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDirIterator(ptr QDirIterator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDirIterator_PTR().Pointer()
	}
	return nil
}

func NewQDirIteratorFromPointer(ptr unsafe.Pointer) *QDirIterator {
	var n = new(QDirIterator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDirIterator) QDirIterator_PTR() *QDirIterator {
	return ptr
}

//QDirIterator::IteratorFlag
type QDirIterator__IteratorFlag int64

const (
	QDirIterator__NoIteratorFlags = QDirIterator__IteratorFlag(0x0)
	QDirIterator__FollowSymlinks  = QDirIterator__IteratorFlag(0x1)
	QDirIterator__Subdirectories  = QDirIterator__IteratorFlag(0x2)
)
