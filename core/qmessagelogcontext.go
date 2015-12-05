package core

//#include "core.h"
import "C"
import (
	"unsafe"
)

type QMessageLogContext struct {
	ptr unsafe.Pointer
}

type QMessageLogContext_ITF interface {
	QMessageLogContext_PTR() *QMessageLogContext
}

func (p *QMessageLogContext) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMessageLogContext) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMessageLogContext(ptr QMessageLogContext_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMessageLogContext_PTR().Pointer()
	}
	return nil
}

func NewQMessageLogContextFromPointer(ptr unsafe.Pointer) *QMessageLogContext {
	var n = new(QMessageLogContext)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMessageLogContext) QMessageLogContext_PTR() *QMessageLogContext {
	return ptr
}
