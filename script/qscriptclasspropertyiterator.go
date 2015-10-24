package script

//#include "qscriptclasspropertyiterator.h"
import "C"
import (
	"unsafe"
)

type QScriptClassPropertyIterator struct {
	ptr unsafe.Pointer
}

type QScriptClassPropertyIteratorITF interface {
	QScriptClassPropertyIteratorPTR() *QScriptClassPropertyIterator
}

func (p *QScriptClassPropertyIterator) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QScriptClassPropertyIterator) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQScriptClassPropertyIterator(ptr QScriptClassPropertyIteratorITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QScriptClassPropertyIteratorPTR().Pointer()
	}
	return nil
}

func QScriptClassPropertyIteratorFromPointer(ptr unsafe.Pointer) *QScriptClassPropertyIterator {
	var n = new(QScriptClassPropertyIterator)
	n.SetPointer(ptr)
	return n
}

func (ptr *QScriptClassPropertyIterator) QScriptClassPropertyIteratorPTR() *QScriptClassPropertyIterator {
	return ptr
}
