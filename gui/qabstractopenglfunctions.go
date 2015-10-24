package gui

//#include "qabstractopenglfunctions.h"
import "C"
import (
	"unsafe"
)

type QAbstractOpenGLFunctions struct {
	ptr unsafe.Pointer
}

type QAbstractOpenGLFunctionsITF interface {
	QAbstractOpenGLFunctionsPTR() *QAbstractOpenGLFunctions
}

func (p *QAbstractOpenGLFunctions) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAbstractOpenGLFunctions) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQAbstractOpenGLFunctions(ptr QAbstractOpenGLFunctionsITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractOpenGLFunctionsPTR().Pointer()
	}
	return nil
}

func QAbstractOpenGLFunctionsFromPointer(ptr unsafe.Pointer) *QAbstractOpenGLFunctions {
	var n = new(QAbstractOpenGLFunctions)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAbstractOpenGLFunctions) QAbstractOpenGLFunctionsPTR() *QAbstractOpenGLFunctions {
	return ptr
}
