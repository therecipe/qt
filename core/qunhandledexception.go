package core

//#include "core.h"
import "C"
import (
	"unsafe"
)

type QUnhandledException struct {
	QException
}

type QUnhandledException_ITF interface {
	QException_ITF
	QUnhandledException_PTR() *QUnhandledException
}

func PointerFromQUnhandledException(ptr QUnhandledException_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QUnhandledException_PTR().Pointer()
	}
	return nil
}

func NewQUnhandledExceptionFromPointer(ptr unsafe.Pointer) *QUnhandledException {
	var n = new(QUnhandledException)
	n.SetPointer(ptr)
	return n
}

func (ptr *QUnhandledException) QUnhandledException_PTR() *QUnhandledException {
	return ptr
}
