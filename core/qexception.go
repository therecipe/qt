package core

//#include "qexception.h"
import "C"
import (
	"unsafe"
)

type QException struct {
	ptr unsafe.Pointer
}

type QException_ITF interface {
	QException_PTR() *QException
}

func (p *QException) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QException) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQException(ptr QException_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QException_PTR().Pointer()
	}
	return nil
}

func NewQExceptionFromPointer(ptr unsafe.Pointer) *QException {
	var n = new(QException)
	n.SetPointer(ptr)
	return n
}

func (ptr *QException) QException_PTR() *QException {
	return ptr
}

func (ptr *QException) Clone() *QException {
	if ptr.Pointer() != nil {
		return NewQExceptionFromPointer(C.QException_Clone(ptr.Pointer()))
	}
	return nil
}

func (ptr *QException) Raise() {
	if ptr.Pointer() != nil {
		C.QException_Raise(ptr.Pointer())
	}
}
