package script

//#include "qscriptclasspropertyiterator.h"
import "C"
import (
	"unsafe"
)

type QScriptClassPropertyIterator struct {
	ptr unsafe.Pointer
}

type QScriptClassPropertyIterator_ITF interface {
	QScriptClassPropertyIterator_PTR() *QScriptClassPropertyIterator
}

func (p *QScriptClassPropertyIterator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QScriptClassPropertyIterator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQScriptClassPropertyIterator(ptr QScriptClassPropertyIterator_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptClassPropertyIterator_PTR().Pointer()
	}
	return nil
}

func NewQScriptClassPropertyIteratorFromPointer(ptr unsafe.Pointer) *QScriptClassPropertyIterator {
	var n = new(QScriptClassPropertyIterator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QScriptClassPropertyIterator) QScriptClassPropertyIterator_PTR() *QScriptClassPropertyIterator {
	return ptr
}
