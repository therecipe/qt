package core

//#include "qunhandledexception.h"
import "C"
import (
	"unsafe"
)

type QUnhandledException struct {
	QException
}

type QUnhandledExceptionITF interface {
	QExceptionITF
	QUnhandledExceptionPTR() *QUnhandledException
}

func PointerFromQUnhandledException(ptr QUnhandledExceptionITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QUnhandledExceptionPTR().Pointer()
	}
	return nil
}

func QUnhandledExceptionFromPointer(ptr unsafe.Pointer) *QUnhandledException {
	var n = new(QUnhandledException)
	n.SetPointer(ptr)
	return n
}

func (ptr *QUnhandledException) QUnhandledExceptionPTR() *QUnhandledException {
	return ptr
}
