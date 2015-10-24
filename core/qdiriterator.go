package core

//#include "qdiriterator.h"
import "C"
import (
	"unsafe"
)

type QDirIterator struct {
	ptr unsafe.Pointer
}

type QDirIteratorITF interface {
	QDirIteratorPTR() *QDirIterator
}

func (p *QDirIterator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QDirIterator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQDirIterator(ptr QDirIteratorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QDirIteratorPTR().Pointer()
	}
	return nil
}

func QDirIteratorFromPointer(ptr unsafe.Pointer) *QDirIterator {
	var n = new(QDirIterator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QDirIterator) QDirIteratorPTR() *QDirIterator {
	return ptr
}

//QDirIterator::IteratorFlag
type QDirIterator__IteratorFlag int

var (
	QDirIterator__NoIteratorFlags = QDirIterator__IteratorFlag(0x0)
	QDirIterator__FollowSymlinks  = QDirIterator__IteratorFlag(0x1)
	QDirIterator__Subdirectories  = QDirIterator__IteratorFlag(0x2)
)
