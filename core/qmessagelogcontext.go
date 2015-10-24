package core

//#include "qmessagelogcontext.h"
import "C"
import (
	"unsafe"
)

type QMessageLogContext struct {
	ptr unsafe.Pointer
}

type QMessageLogContextITF interface {
	QMessageLogContextPTR() *QMessageLogContext
}

func (p *QMessageLogContext) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QMessageLogContext) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQMessageLogContext(ptr QMessageLogContextITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QMessageLogContextPTR().Pointer()
	}
	return nil
}

func QMessageLogContextFromPointer(ptr unsafe.Pointer) *QMessageLogContext {
	var n = new(QMessageLogContext)
	n.SetPointer(ptr)
	return n
}

func (ptr *QMessageLogContext) QMessageLogContextPTR() *QMessageLogContext {
	return ptr
}
