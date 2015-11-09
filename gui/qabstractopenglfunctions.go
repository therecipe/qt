package gui

//#include "qabstractopenglfunctions.h"
import "C"
import (
	"unsafe"
)

type QAbstractOpenGLFunctions struct {
	ptr unsafe.Pointer
}

type QAbstractOpenGLFunctions_ITF interface {
	QAbstractOpenGLFunctions_PTR() *QAbstractOpenGLFunctions
}

func (p *QAbstractOpenGLFunctions) Pointer() unsafe.Pointer {
	return p.ptr
}

func (p *QAbstractOpenGLFunctions) SetPointer(ptr unsafe.Pointer) {
	p.ptr = ptr
}

func PointerFromQAbstractOpenGLFunctions(ptr QAbstractOpenGLFunctions_ITF) unsafe.Pointer {
	if ptr != nil {
		return ptr.QAbstractOpenGLFunctions_PTR().Pointer()
	}
	return nil
}

func NewQAbstractOpenGLFunctionsFromPointer(ptr unsafe.Pointer) *QAbstractOpenGLFunctions {
	var n = new(QAbstractOpenGLFunctions)
	n.SetPointer(ptr)
	return n
}

func (ptr *QAbstractOpenGLFunctions) QAbstractOpenGLFunctions_PTR() *QAbstractOpenGLFunctions {
	return ptr
}
