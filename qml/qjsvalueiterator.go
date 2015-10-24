package qml

//#include "qjsvalueiterator.h"
import "C"
import (
	"unsafe"
)

type QJSValueIterator struct {
	ptr unsafe.Pointer
}

type QJSValueIteratorITF interface {
	QJSValueIteratorPTR() *QJSValueIterator
}

func (p *QJSValueIterator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QJSValueIterator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQJSValueIterator(ptr QJSValueIteratorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QJSValueIteratorPTR().Pointer()
	}
	return nil
}

func QJSValueIteratorFromPointer(ptr unsafe.Pointer) *QJSValueIterator {
	var n = new(QJSValueIterator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QJSValueIterator) QJSValueIteratorPTR() *QJSValueIterator {
	return ptr
}
