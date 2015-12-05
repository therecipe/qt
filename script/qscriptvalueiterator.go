package script

//#include "script.h"
import "C"
import (
	"unsafe"
)

type QScriptValueIterator struct {
	ptr unsafe.Pointer
}

type QScriptValueIterator_ITF interface {
	QScriptValueIterator_PTR() *QScriptValueIterator
}

func (p *QScriptValueIterator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QScriptValueIterator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQScriptValueIterator(ptr QScriptValueIterator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptValueIterator_PTR().Pointer()
	}
	return nil
}

func NewQScriptValueIteratorFromPointer(ptr unsafe.Pointer) *QScriptValueIterator {
	var n = new(QScriptValueIterator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QScriptValueIterator) QScriptValueIterator_PTR() *QScriptValueIterator {
	return ptr
}
