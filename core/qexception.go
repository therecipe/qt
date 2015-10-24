package core

//#include "qexception.h"
import "C"
import (
	"unsafe"
)

type QException struct {
	ptr unsafe.Pointer
}

type QExceptionITF interface {
	QExceptionPTR() *QException
}

func (p *QException) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QException) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQException(ptr QExceptionITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QExceptionPTR().Pointer()
	}
	return nil
}

func QExceptionFromPointer(ptr unsafe.Pointer) *QException {
	var n = new(QException)
	n.SetPointer(ptr)
	return n
}

func (ptr *QException) QExceptionPTR() *QException {
	return ptr
}

func (ptr *QException) Clone() *QException {
	if ptr.Pointer() != nil {
		return QExceptionFromPointer(unsafe.Pointer(C.QException_Clone(C.QtObjectPtr(ptr.Pointer()))))
	}
	return nil
}

func (ptr *QException) Raise() {
	if ptr.Pointer() != nil {
		C.QException_Raise(C.QtObjectPtr(ptr.Pointer()))
	}
}
