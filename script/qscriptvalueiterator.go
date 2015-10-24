package script

//#include "qscriptvalueiterator.h"
import "C"
import (
	"unsafe"
)

type QScriptValueIterator struct {
	ptr unsafe.Pointer
}

type QScriptValueIteratorITF interface {
	QScriptValueIteratorPTR() *QScriptValueIterator
}

func (p *QScriptValueIterator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QScriptValueIterator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQScriptValueIterator(ptr QScriptValueIteratorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptValueIteratorPTR().Pointer()
	}
	return nil
}

func QScriptValueIteratorFromPointer(ptr unsafe.Pointer) *QScriptValueIterator {
	var n = new(QScriptValueIterator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QScriptValueIterator) QScriptValueIteratorPTR() *QScriptValueIterator {
	return ptr
}
