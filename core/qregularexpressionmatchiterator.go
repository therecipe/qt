package core

//#include "qregularexpressionmatchiterator.h"
import "C"
import (
	"unsafe"
)

type QRegularExpressionMatchIterator struct {
	ptr unsafe.Pointer
}

type QRegularExpressionMatchIterator_ITF interface {
	QRegularExpressionMatchIterator_PTR() *QRegularExpressionMatchIterator
}

func (p *QRegularExpressionMatchIterator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QRegularExpressionMatchIterator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQRegularExpressionMatchIterator(ptr QRegularExpressionMatchIterator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRegularExpressionMatchIterator_PTR().Pointer()
	}
	return nil
}

func NewQRegularExpressionMatchIteratorFromPointer(ptr unsafe.Pointer) *QRegularExpressionMatchIterator {
	var n = new(QRegularExpressionMatchIterator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QRegularExpressionMatchIterator) QRegularExpressionMatchIterator_PTR() *QRegularExpressionMatchIterator {
	return ptr
}
