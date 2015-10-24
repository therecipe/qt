package core

//#include "qregularexpressionmatchiterator.h"
import "C"
import (
	"unsafe"
)

type QRegularExpressionMatchIterator struct {
	ptr unsafe.Pointer
}

type QRegularExpressionMatchIteratorITF interface {
	QRegularExpressionMatchIteratorPTR() *QRegularExpressionMatchIterator
}

func (p *QRegularExpressionMatchIterator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QRegularExpressionMatchIterator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQRegularExpressionMatchIterator(ptr QRegularExpressionMatchIteratorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QRegularExpressionMatchIteratorPTR().Pointer()
	}
	return nil
}

func QRegularExpressionMatchIteratorFromPointer(ptr unsafe.Pointer) *QRegularExpressionMatchIterator {
	var n = new(QRegularExpressionMatchIterator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QRegularExpressionMatchIterator) QRegularExpressionMatchIteratorPTR() *QRegularExpressionMatchIterator {
	return ptr
}
