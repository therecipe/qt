package qml

//#include "qjsvalueiterator.h"
import "C"
import (
	"unsafe"
)

type QJSValueIterator struct {
	ptr unsafe.Pointer
}

type QJSValueIterator_ITF interface {
	QJSValueIterator_PTR() *QJSValueIterator
}

func (p *QJSValueIterator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QJSValueIterator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQJSValueIterator(ptr QJSValueIterator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QJSValueIterator_PTR().Pointer()
	}
	return nil
}

func NewQJSValueIteratorFromPointer(ptr unsafe.Pointer) *QJSValueIterator {
	var n = new(QJSValueIterator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QJSValueIterator) QJSValueIterator_PTR() *QJSValueIterator {
	return ptr
}
